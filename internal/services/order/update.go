package order

import (
	"context"

	"github.com/shopspring/decimal"
)

type Update struct {
	Amount decimal.Decimal
}

func (o *Order) Update(ctx context.Context, sharding string, data *Update) error {
	return nil
}
