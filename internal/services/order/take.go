package order

import (
	"context"

	"example/internal/models"
	"example/internal/repositories/order"
)

type Take struct {
	OrderNo string
}

func (o *Order) Take(ctx context.Context, data *Take) (*models.Order, error) {
	return o.orderRepo.Take().
		Where(
			order.ConditionShardingEq(data.OrderNo[:6]),
			order.ConditionOrderNoEq(data.OrderNo),
			order.ConditionDeletedAtIsZero(),
		).Do(ctx)
}
