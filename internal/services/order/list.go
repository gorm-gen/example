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
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	var page, pageSize uint
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
		page = uint(data.Page)
		pageSize = uint(data.PageSize)
	}

	return o.orderRepo.List().
		Where(conditions...).
		Order(order.OrderIDDesc()).
		Page(page, pageSize).
		Do(ctx)
}

func (o *Order) MultiList(ctx context.Context, sharding []string, data *List) ([]*models.Order, int64, error) {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	var page, pageSize uint
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
		page = uint(data.Page)
		pageSize = uint(data.PageSize)
	}

	return o.orderRepo.ShardingList(sharding).
		Where(conditions...).
		Order(order.OrderBy(o.q.Order.Amount.Desc()), order.OrderIDDesc()).
		Page(page, pageSize).
		Do(ctx)
}
