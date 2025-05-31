package order

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
	"gorm.io/gen"

	"example/internal/query"
	"example/internal/repositories"
)

type multiCount struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	conditionOpts []ConditionOption
	sharding      []string
	worker        chan struct{}
}

// MultiCount 获取多表数据总条数
func (o *Order) MultiCount(sharding []string) *multiCount {
	return &multiCount{
		core:          o,
		unscoped:      o.unscoped,
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
	}
}

func (c *multiCount) Worker(worker chan struct{}) *multiCount {
	c.worker = worker
	return c
}

// Tx 设置为事务
func (c *multiCount) Tx(tx *query.Query) *multiCount {
	c.tx = tx
	c.qTx = nil
	return c
}

// QueryTx 设置为手动事务
func (c *multiCount) QueryTx(tx *query.QueryTx) *multiCount {
	c.qTx = tx
	c.tx = nil
	return c
}

func (c *multiCount) Unscoped() *multiCount {
	c.unscoped = true
	return c
}

func (c *multiCount) Where(opts ...ConditionOption) *multiCount {
	c.conditionOpts = append(c.conditionOpts, opts...)
	return c
}

// Do 执行获取多表数据总条数
func (c *multiCount) Do(ctx context.Context) (int64, error) {
	if len(c.sharding) == 0 {
		return 0, nil
	}
	cq := c.core.q.Order
	if c.tx != nil {
		cq = c.tx.Order
	}
	if c.qTx != nil {
		cq = c.qTx.Order
	}
	var conditions []gen.Condition
	if len(c.conditionOpts) > 0 {
		conditions = make([]gen.Condition, 0, len(c.conditionOpts))
		for _, opt := range c.conditionOpts {
			conditions = append(conditions, opt(c.core))
		}
	}
	count := int64(0)
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	endChan := make(chan struct{})
	for _, sharding := range c.sharding {
		c.worker <- struct{}{}
		wg.Add(1)
		go func(sharding string) {
			defer func() {
				if r := recover(); r != nil {
					c.core.logger.Error(fmt.Sprintf("【Order.MultiCount.%s】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
					errChan <- fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			defer func() {
				<-c.worker
			}()
			defer wg.Done()
			_conditions := make([]gen.Condition, len(conditions))
			copy(_conditions, conditions)
			_conditions = append(_conditions, ConditionSharding(sharding)(c.core))
			cr := cq.WithContext(ctx)
			if c.unscoped {
				cr = cr.Unscoped()
			}
			__count, err := cr.Where(_conditions...).Count()
			if err != nil {
				if repositories.IsRealErr(err) {
					c.core.logger.Error(fmt.Sprintf("【Order.MultiCount.%s】失败", sharding), zap.Error(err))
				}
				errChan <- err
				return
			}
			atomic.AddInt64(&count, __count)
			return
		}(sharding)
	}
	go func() {
		wg.Wait()
		endChan <- struct{}{}
	}()
	select {
	case <-endChan:
		return count, nil
	case err := <-errChan:
		return 0, err
	}
}
