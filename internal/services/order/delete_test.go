package order_test

import (
	"context"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestDelete(t *testing.T) {
	id := int64(1)
	if err := orderSvc.Delete(context.Background(), time.Now().Format("200601"), &order.Delete{ID: &id}); err != nil {
		t.Error(err)
		return
	}
}

func TestPhysicalShardingDelete(t *testing.T) {
	id := int64(3)
	if err := orderSvc.PhysicalShardingDelete(context.Background(), shardingList, &order.Delete{ID: &id}); err != nil {
		t.Error(err)
		return
	}
}

func TestShardingDelete(t *testing.T) {
	id := int64(3)
	if err := orderSvc.ShardingDelete(context.Background(), shardingList, &order.Delete{ID: &id}); err != nil {
		t.Error(err)
		return
	}
}
