package order

import (
	"context"

	"github.com/shopspring/decimal"

	"example/internal/repositories/order"
)

type Sum struct {
	UID *int
}

func (o *Order) Sum(ctx context.Context, sharding string, data *Sum) (decimal.Decimal, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionShardingEq(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
	}
	if data.UID != nil {
		conditions = append(conditions, order.ConditionUID(*data.UID))
	}
	return o.orderRepo.Sum(o.q.Order.Amount).
		Where(conditions...).
		Do(ctx)
}
