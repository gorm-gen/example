package order

import (
	"context"

	"example/internal/models"
	"example/internal/repositories/order"
)

type First struct {
	ID      *int64
	UID     *int
	OrderNo *string
}

func (o *Order) First(ctx context.Context, sharding string, data *First) (*models.Order, error) {
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
	return o.orderRepo.First().
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

func (o *Order) ShardingFirst(ctx context.Context, sharding []string, data *First) (*models.Order, error) {
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
	return o.orderRepo.ShardingFirst(sharding).
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
