package order

import (
	"context"

	"github.com/shopspring/decimal"

	"example/internal/repositories/order"
)

func (o *Order) ShardingMax(ctx context.Context, sharding []string) (decimal.Decimal, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	conditions = append(conditions, order.ConditionUID(1))
	_max, _, err := o.orderRepo.ShardingMax(o.q.Order.Amount, sharding).Where(conditions...).Do(ctx)
	return _max, err
}
