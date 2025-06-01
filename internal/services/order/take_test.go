package order_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestTake(t *testing.T) {
	res, err := orderSvc.Take(context.Background(), time.Now().Format("200601"), &order.Take{})
	if err != nil {
		t.Fatal(err)
		return
	}
	var bytes []byte
	if bytes, err = json.Marshal(res); err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes))
}

func TestShardingTake(t *testing.T) {
	res, err := orderSvc.ShardingTake(context.Background(), []string{"202505", "202506"}, &order.Take{})
	if err != nil {
		t.Fatal(err)
		return
	}
	var bytes []byte
	if bytes, err = json.Marshal(res); err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes))
}
