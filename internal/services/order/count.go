package order

import (
	"context"
	"example/internal/repositories/order"
)

type Count struct {
	ID      *int
	OrderNo *string
}

func (o *Order) Count(ctx context.Context, sharding string, data *Count) (int64, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionShardingEq(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data.ID != nil {
		conditions = append(conditions, order.ConditionID(*data.ID))
	}
	if data.OrderNo != nil {
		conditions = append(conditions, order.ConditionOrderNoEq(*data.OrderNo))
	}
	return o.orderRepo.
		Count().
		Where(conditions...).
		Do(ctx)
}
