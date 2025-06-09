package order

import (
	"context"

	"example/internal/repositories/order"
)

func (o *Order) Scan(ctx context.Context, sharding string) (interface{}, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	var list []struct {
		ID      int    `json:"id"`
		OrderNo string `json:"order_no"`
	}
	err := o.orderRepo.
		Scan(&list).Select(o.q.Order.ID, o.q.Order.OrderNo).
		Where(conditions...).
		Order(order.OrderIDDesc()).
		Do(ctx)
	return list, err
}
