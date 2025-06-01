package order_test

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	orderRepo "example/internal/repositories/order"
	"example/internal/services/order"
)

func TestUpdate(t *testing.T) {
	opts := make([]orderRepo.UpdateOption, 0)
	opts = append(opts, orderRepo.UpdateAmountAdd(decimal.NewFromFloat(999)))
	id := int64(2)
	err := orderSvc.Update(context.Background(), time.Now().Format("200601"), &order.Update{ID: &id}, opts...)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestShardingUpdate(t *testing.T) {
	opts := make([]orderRepo.UpdateOption, 0)
	var id int64 = 3
	opts = append(opts, orderRepo.UpdateAmountAdd(decimal.NewFromFloat(66)))
	err := orderSvc.ShardingUpdate(context.Background(), shardingList, &order.Update{ID: &id}, opts...)
	if err != nil {
		t.Error(err)
		return
	}
}
