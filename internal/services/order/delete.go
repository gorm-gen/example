package order

import (
	"context"

	"example/internal/repositories/order"
)

type Delete struct {
	ID      *int64
	UID     *int
	OrderNo *string
}

func (o *Order) Delete(ctx context.Context, sharding string, data *Delete) error {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionShardingEq(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.UID != nil {
			conditions = append(conditions, order.ConditionUID(*data.UID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNoEq(*data.OrderNo))
		}
	}
	_, err := o.orderRepo.Delete().
		Where(conditions...).
		Do(ctx)
	return err
}
