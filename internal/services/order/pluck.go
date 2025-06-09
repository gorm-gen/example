package order

import (
	"context"

	"example/internal/repositories/order"
)

func (o *Order) Pluck(ctx context.Context, sharding string) (interface{}, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	var orderNos []string
	err := o.orderRepo.Pluck(o.q.Order.OrderNo, &orderNos).Where(conditions...).Order(order.OrderIDDesc()).Do(ctx)
	return orderNos, err
}
