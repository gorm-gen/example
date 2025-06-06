package order

import (
	"context"

	"example/internal/models"
	"example/internal/repositories/order"
)

type Last struct {
	ID      *int64
	UID     *int
	OrderNo *string
}

func (o *Order) Last(ctx context.Context, sharding string, data *Last) (*models.Order, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.UID != nil {
			conditions = append(conditions, order.ConditionUID(*data.UID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
	}
	return o.orderRepo.Last().
		Select(
			o.q.Order.ID,
			o.q.Order.UID,
			o.q.Order.OrderNo,
			o.q.Order.Sharding,
			o.q.Order.Sharding,
			o.q.Order.Amount,
			o.q.Order.CreatedAt,
		).Where(conditions...).
		Do(ctx)
}

func (o *Order) ShardingLast(ctx context.Context, sharding []string, data *Last) (*models.Order, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.UID != nil {
			conditions = append(conditions, order.ConditionUID(*data.UID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
	}
	return o.orderRepo.ShardingLast(sharding).
		Select(
			o.q.Order.ID,
			o.q.Order.UID,
			o.q.Order.OrderNo,
			o.q.Order.Sharding,
			o.q.Order.Sharding,
			o.q.Order.Amount,
			o.q.Order.CreatedAt,
		).Where(conditions...).
		Do(ctx)
}
