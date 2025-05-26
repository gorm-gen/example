package order

import (
	"context"

	"example/internal/repositories/order"
)

type Update struct {
	ID      *int64
	UID     *int
	OrderNo *string
}

func (o *Order) Update(ctx context.Context, sharding string, data *Update, opts ...order.UpdateOption) error {
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
	_, err := o.orderRepo.Update().
		Where(conditions...).
		Update(opts...).
		Do(ctx)
	return err
}
