package order

import (
	"context"
	"fmt"
	"runtime/debug"

	"go.uber.org/zap"

	"example/internal/query"

	"example/internal/repositories"

	"example/internal/models"
)

type _shardingCreate struct {
	core      *Order
	tx        *query.Query
	qTx       *query.QueryTx
	unscoped  bool
	values    []*models.Order
	batchSize int
}

// ShardingCreate 分表添加数据
func (o *Order) ShardingCreate() *_shardingCreate {
	return &_shardingCreate{
		core:     o,
		unscoped: o.unscoped,
		values:   make([]*models.Order, 0),
	}
}

// Tx 设置为事务
func (c *_shardingCreate) Tx(tx *query.Query) *_shardingCreate {
	c.tx = tx
	if tx != nil {
		c.qTx = nil
	}
	return c
}

// QueryTx 设置为手动事务
func (c *_shardingCreate) QueryTx(tx *query.QueryTx) *_shardingCreate {
	c.qTx = tx
	if tx != nil {
		c.tx = nil
	}
	return c
}

func (c *_shardingCreate) Unscoped() *_shardingCreate {
	c.unscoped = true
	return c
}

func (c *_shardingCreate) Values(values ...*models.Order) *_shardingCreate {
	c.values = append(c.values, values...)
	return c
}

// BatchSize 当批量插入时指定创建的数量
func (c *_shardingCreate) BatchSize(batchSize uint) *_shardingCreate {
	c.batchSize = int(batchSize)
	return c
}

// Do 执行添加数据
func (c *_shardingCreate) do(ctx context.Context, tx *query.Query, qTx *query.QueryTx, values ...*models.Order) (err error) {
	length := len(values)
	if length == 0 {
		return nil
	}
	cq := c.core.q.Order
	if tx != nil {
		cq = tx.Order
	}
	if qTx != nil {
		cq = qTx.Order
	}
	cr := cq.WithContext(ctx)
	if c.unscoped {
		cr = cr.Unscoped()
	}
	if length > 1 && c.batchSize > 0 {
		err = cr.CreateInBatches(values, c.batchSize)
	} else {
		err = cr.Create(values...)
	}
	if err != nil {
		if repositories.IsRealErr(err) {
			c.core.logger.Error("【Order.Create】失败", zap.Error(err))
		}
		return err
	}
	return nil
}

// Do 执行添加数据
func (c *_shardingCreate) Do(ctx context.Context) (err error) {
	length := len(c.values)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return c.do(ctx, c.tx, c.qTx, c.values...)
	}
	m := make(map[string][]*models.Order, length)
	for _, value := range c.values {
		m[value.Sharding] = append(m[value.Sharding], value)
	}
	if len(m) == 1 {
		return c.do(ctx, c.tx, c.qTx, c.values...)
	}
	if c.tx != nil || c.qTx != nil {
		for _, values := range m {
			if err = c.do(ctx, c.tx, c.qTx, values...); err != nil {
				return
			}
		}
		return
	}
	return c.core.q.Transaction(func(tx *query.Query) (err error) {
		defer func() {
			if r := recover(); r != nil {
				c.core.logger.Error("【Order.ShardingCreate】执行异常", zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
				err = fmt.Errorf("recovered from panic: %v", r)
				return
			}
		}()
		for _, values := range m {
			if err = c.do(ctx, tx, nil, values...); err != nil {
				return
			}
		}
		return
	})
}
