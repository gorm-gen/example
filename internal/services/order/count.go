package order

import (
	"context"

	"example/internal/repositories/order"
)

type Count struct {
	ID      *int64
	OrderNo *string
}

func (o *Order) Count(ctx context.Context, sharding string, data *Count) (int64, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
	}
	return o.orderRepo.
		Count().
		Where(conditions...).
		Do(ctx)
}

func (o *Order) MultiCount(ctx context.Context, sharding []string, data *Count) (int64, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
	}
	count, _, err := o.orderRepo.
		ShardingCount(sharding).
		Where(conditions...).
		Do(ctx)
	return count, err
}
