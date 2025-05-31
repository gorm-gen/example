package order

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"example/internal/query"

	"example/internal/repositories"
)

type multiUpdate struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	updateOpts    []UpdateOption
	conditionOpts []ConditionOption
	sharding      []string
	worker        chan struct{}
}

// MultiUpdate 更新多表数据
func (o *Order) MultiUpdate(sharding []string) *multiUpdate {
	return &multiUpdate{
		core:          o,
		unscoped:      o.unscoped,
		updateOpts:    make([]UpdateOption, 0),
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
	}
}

func (u *multiUpdate) Worker(worker chan struct{}) *multiUpdate {
	if worker == nil {
		return u
	}
	u.worker = worker
	return u
}

// Tx 设置为事务
func (u *multiUpdate) Tx(tx *query.Query) *multiUpdate {
	u.tx = tx
	if tx != nil {
		u.qTx = nil
	}
	return u
}

// QueryTx 设置为手动事务
func (u *multiUpdate) QueryTx(tx *query.QueryTx) *multiUpdate {
	u.qTx = tx
	if tx != nil {
		u.tx = nil
	}
	return u
}

func (u *multiUpdate) Unscoped() *multiUpdate {
	u.unscoped = true
	return u
}

func (u *multiUpdate) Update(opts ...UpdateOption) *multiUpdate {
	u.updateOpts = append(u.updateOpts, opts...)
	return u
}

func (u *multiUpdate) Where(opts ...ConditionOption) *multiUpdate {
	u.conditionOpts = append(u.conditionOpts, opts...)
	return u
}

// Do 执行更新多表数据
func (u *multiUpdate) Do(ctx context.Context) (int64, map[string]int64, error) {
	if len(u.updateOpts) == 0 || len(u.sharding) == 0 {
		return 0, nil, nil
	}
	uq := u.core.q.Order
	if u.tx != nil {
		uq = u.tx.Order
	}
	if u.qTx != nil {
		uq = u.qTx.Order
	}
	var conditions []gen.Condition
	if len(u.conditionOpts) > 0 {
		conditions = make([]gen.Condition, 0, len(u.conditionOpts))
		for _, opt := range u.conditionOpts {
			conditions = append(conditions, opt(u.core))
		}
	}
	columns := make([]field.AssignExpr, 0, len(u.updateOpts))
	for _, opt := range u.updateOpts {
		columns = append(columns, opt(u.core))
	}
	if len(columns) == 0 {
		return 0, nil, nil
	}
	m := make(map[string]int64, len(u.sharding))
	sm := sync.Map{}
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	endChan := make(chan struct{})
	for _, sharding := range u.sharding {
		u.worker <- struct{}{}
		wg.Add(1)
		go func(sharding string) {
			defer func() {
				if r := recover(); r != nil {
					u.core.logger.Error(fmt.Sprintf("【Order.MultiUpdate.%s】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
					errChan <- fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			defer func() {
				<-u.worker
			}()
			defer wg.Done()
			_conditions := make([]gen.Condition, len(conditions))
			copy(_conditions, conditions)
			_conditions = append(_conditions, ConditionSharding(sharding)(u.core))
			ur := uq.WithContext(ctx)
			if u.unscoped {
				ur = ur.Unscoped()
			}
			res, err := ur.Where(_conditions...).UpdateSimple(columns...)
			if err != nil {
				if repositories.IsRealErr(err) {
					u.core.logger.Error(fmt.Sprintf("【Order.MultiUpdate.%s】失败", sharding), zap.Error(err))
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
