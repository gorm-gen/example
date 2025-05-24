package order

import (
	"context"

	"example/internal/models"
	"example/internal/repositories/order"
)

type Take struct {
	ID      *int
	UID     *int
	OrderNo *string
}

func (o *Order) Take(ctx context.Context, sharding string, data *Take) (*models.Order, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionShardingEq(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data.ID != nil {
		conditions = append(conditions, order.ConditionID(*data.ID))
	}
	if data.UID != nil {
		conditions = append(conditions, order.ConditionUID(*data.UID))
	}
	if data.OrderNo != nil {
		conditions = append(conditions, order.ConditionOrderNoEq(*data.OrderNo))
	}
	return o.orderRepo.Take().
		Where(conditions...).
		Do(ctx)
}
