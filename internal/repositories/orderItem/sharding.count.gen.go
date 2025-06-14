// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.

package orderItem

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

type _shardingCount struct {
	core          *OrderItem
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	conditionOpts []ConditionOption
	sharding      []int
	worker        chan struct{}
	writeDB       bool
	scopes        []func(gen.Dao) gen.Dao
}

// ShardingCount 获取分表数据总记录
func (o *OrderItem) ShardingCount(sharding []int) *_shardingCount {
	return &_shardingCount{
		core:          o,
		unscoped:      o.unscoped,
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
		scopes:        make([]func(gen.Dao) gen.Dao, 0),
	}
}

func (c *_shardingCount) Worker(worker chan struct{}) *_shardingCount {
	if worker == nil {
		return c
	}
	c.worker = worker
	return c
}

// Tx 设置为事务
func (c *_shardingCount) Tx(tx *query.Query) *_shardingCount {
	c.tx = tx
	c.qTx = nil
	return c
}

// QueryTx 设置为手动事务
func (c *_shardingCount) QueryTx(tx *query.QueryTx) *_shardingCount {
	c.qTx = tx
	c.tx = nil
	return c
}

func (c *_shardingCount) Unscoped() *_shardingCount {
	c.unscoped = true
	return c
}

func (c *_shardingCount) Scopes(funcs ...func(gen.Dao) gen.Dao) *_shardingCount {
	c.scopes = append(c.scopes, funcs...)
	return c
}

func (c *_shardingCount) Where(opts ...ConditionOption) *_shardingCount {
	c.conditionOpts = append(c.conditionOpts, opts...)
	return c
}

func (c *_shardingCount) WriteDB() *_shardingCount {
	c.writeDB = true
	return c
}

// Do 执行获取分表数据总记录
func (c *_shardingCount) Do(ctx context.Context) (int64, map[int]int64, error) {
	_lenSharding := len(c.sharding)
	if _lenSharding == 0 {
		return 0, nil, nil
	}
	cq := c.core.q.OrderItem
	if c.tx != nil {
		cq = c.tx.OrderItem
	}
	if c.qTx != nil {
		cq = c.qTx.OrderItem
	}
	var conditions []gen.Condition
	if _len := len(c.conditionOpts); _len > 0 {
		conditions = make([]gen.Condition, 0, _len)
		for _, opt := range c.conditionOpts {
			conditions = append(conditions, opt(c.core))
		}
	}
	sm := sync.Map{}
	wg := sync.WaitGroup{}
	errChan := make(chan error)
	endChan := make(chan struct{})
	for _, sharding := range c.sharding {
		c.worker <- struct{}{}
		wg.Add(1)
		go func(sharding int) {
			defer func() {
				if r := recover(); r != nil {
					c.core.logger.Error(fmt.Sprintf("【OrderItem.ShardingCount.%d】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
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
			if c.writeDB {
				cr = cr.WriteDB()
			}
			if c.unscoped {
				cr = cr.Unscoped()
			}
			if len(c.scopes) > 0 {
				cr = cr.Scopes(c.scopes...)
			}
			count, err := cr.Where(_conditions...).Count()
			if err != nil {
				if repositories.IsRealErr(err) {
					c.core.logger.Error(fmt.Sprintf("【OrderItem.ShardingCount.%d】失败", sharding), zap.Error(err), zap.ByteString("debug.Stack", debug.Stack()))
				}
				errChan <- err
				return
			}
			sm.Store(sharding, count)
			return
		}(sharding)
	}
	go func() {
		wg.Wait()
		endChan <- struct{}{}
	}()
	select {
	case <-endChan:
		count := int64(0)
		m := make(map[int]int64, _lenSharding)
		sm.Range(func(key, value interface{}) bool {
			v := value.(int64)
			m[key.(int)] = v
			count += v
			return true
		})
		return count, m, nil
	case err := <-errChan:
		return 0, nil, err
	}
}
