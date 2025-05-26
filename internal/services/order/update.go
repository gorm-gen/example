package order

import (
	"context"

	"github.com/shopspring/decimal"

	"example/internal/repositories/order"
)

type Update struct {
	Amount decimal.Decimal
}

func (o *Order) Update(ctx context.Context, sharding string, id int64, data *Update) error {
	_, err := o.orderRepo.Update().
		Where(
			order.ConditionShardingEq(sharding),
			order.ConditionID(id),
		).Update(order.UpdateAmountAdd(data.Amount)).
		Do(ctx)
	return err
}
