package order_test

import (
	"context"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestCount(t *testing.T) {
	count, err := orderSvc.Count(context.Background(), time.Now().Format("200601"), &order.Count{})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(count)
}

func TestShardingCount(t *testing.T) {
	count, err := orderSvc.ShardingCount(context.Background(), []string{"202505", "202506"}, &order.Count{})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(count)
}
