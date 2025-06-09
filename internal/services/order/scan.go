package order

import (
	"context"

	"gorm.io/gen/field"

	"example/internal/repositories/order"
)

func (o *Order) Scan(ctx context.Context, sharding string) (interface{}, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	var list []struct {
		ID int    `json:"id"`
		No string `json:"no"`
	}
	err := o.orderRepo.Scan(&list).
		Select(
			field.NewInt64("", o.q.Order.ID.ColumnName().String()),
			field.NewString("", o.q.Order.OrderNo.ColumnName().String()).As("no"),
		).Where(conditions...).
		Order(order.OrderIDDesc()).
		Do(ctx)
	return list, err
}
