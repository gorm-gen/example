package order

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gen"

	"example/internal/query"

	"example/internal/repositories"
)

type multiDelete struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	conditionOpts []ConditionOption
	sharding      []string
	worker        chan struct{}
}

// MultiDelete 删除多表数据
func (o *Order) MultiDelete(sharding []string) *multiDelete {
	return &multiDelete{
		core:          o,
		unscoped:      o.unscoped,
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
	}
}

func (d *multiDelete) Worker(worker chan struct{}) *multiDelete {
	if worker == nil {
		return d
	}
	d.worker = worker
	return d
}

// Tx 设置为事务
func (d *multiDelete) Tx(tx *query.Query) *multiDelete {
	d.tx = tx
	if tx != nil {
		d.qTx = nil
	}
	return d
}

// QueryTx 设置为手动事务
func (d *multiDelete) QueryTx(tx *query.QueryTx) *multiDelete {
	d.qTx = tx
	if tx != nil {
		d.tx = nil
	}
	return d
}

func (d *multiDelete) Unscoped() *multiDelete {
	d.unscoped = true
	return d
}

func (d *multiDelete) Where(opts ...ConditionOption) *multiDelete {
	d.conditionOpts = append(d.conditionOpts, opts...)
	return d
}

// Do 执行删除多表数据
func (d *multiDelete) Do(ctx context.Context) (int64, map[string]int64, error) {
	if len(d.sharding) == 0 {
		return 0, nil, nil
	}
	dq := d.core.q.Order
	if d.tx != nil {
		dq = d.tx.Order
	}
	if d.qTx != nil {
		dq = d.qTx.Order
	}
	var conditions []gen.Condition
	if len(d.conditionOpts) > 0 {
		conditions = make([]gen.Condition, 0, len(d.conditionOpts))
		for _, opt := range d.conditionOpts {
			conditions = append(conditions, opt(d.core))
		}
	}
	m := make(map[string]int64, len(d.sharding))
	sm := sync.Map{}
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	endChan := make(chan struct{})
	for _, sharding := range d.sharding {
		d.worker <- struct{}{}
		wg.Add(1)
		go func(sharding string) {
			defer func() {
				if r := recover(); r != nil {
					d.core.logger.Error(fmt.Sprintf("【Order.MultiDelete.%s】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
					errChan <- fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			defer func() {
				<-d.worker
			}()
			defer wg.Done()
			_conditions := make([]gen.Condition, len(conditions))
			copy(_conditions, conditions)
			_conditions = append(_conditions, ConditionSharding(sharding)(d.core))
			dr := dq.WithContext(ctx)
			if d.unscoped {
				dr = dr.Unscoped()
			}
			res, err := dr.Where(_conditions...).Delete()
			if err != nil {
				if repositories.IsRealErr(err) {
					d.core.logger.Error(fmt.Sprintf("【Order.MultiDelete.%s】失败", sharding), zap.Error(err))
				}
				errChan <- err
				return
			}
			sm.Store(sharding, res.RowsAffected)
			return
		}(sharding)
	}
	go func() {
		wg.Wait()
		endChan <- struct{}{}
	}()
	select {
	case <-endChan:
		rowsAffected := int64(0)
		sm.Range(func(key, value interface{}) bool {
			v := value.(int64)
			m[key.(string)] = v
			rowsAffected += v
			return true
		})
		return rowsAffected, m, nil
	case err := <-errChan:
		return 0, nil, err
	}
}
