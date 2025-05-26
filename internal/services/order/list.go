package order

import (
	"context"

	"example/internal/models"
	"example/internal/repositories/order"
)

type List struct {
	Page     int
	PageSize int
	ID       *int64
	OrderNo  *string
}

func (o *Order) List(ctx context.Context, sharding string, data *List) ([]*models.Order, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionShardingEq(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data.ID != nil {
		conditions = append(conditions, order.ConditionID(*data.ID))
	}
	if data.OrderNo != nil {
		conditions = append(conditions, order.ConditionOrderNoEq(*data.OrderNo))
	}
	return o.orderRepo.List().
		Where(conditions...).
		Order(order.OrderIDDesc()).
		Page(uint(data.Page), uint(data.PageSize)).
		Do(ctx)
}
