package order_test

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"example/internal/services/order"
)

func TestUpdate(t *testing.T) {
	err := orderSvc.Update(context.Background(), time.Now().Format("200601"), 1, &order.Update{
		Amount: decimal.NewFromInt(666),
	})
	if err != nil {
		t.Error(err)
		return
	}
}
