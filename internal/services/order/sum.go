package order

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	"example/internal/repositories/order"
)

func (o *Order) Sum(ctx context.Context) (decimal.Decimal, error) {
	return o.orderRepo.
		Sum(o.q.Order.Amount).
		Where(order.ConditionShardingEq(time.Now().Format("200601"))).
		Do(ctx)
}
