package order

import (
	"context"

	"example/internal/query"
)

type multiCount struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	conditionOpts []ConditionOption
}

// MultiCount 获取数据总条数
func (o *Order) MultiCount() *multiCount {
	return &multiCount{
		core:          o,
		unscoped:      o.unscoped,
		conditionOpts: make([]ConditionOption, 0),
	}
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

// Do 执行获取数据总条数
func (c *multiCount) Do(ctx context.Context) (int64, error) {
	return 0, nil
}
