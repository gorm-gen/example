package order_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestLast(t *testing.T) {
	res, err := orderSvc.Last(context.Background(), time.Now().Format("200601"), &order.Last{})
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

func TestMultiLast(t *testing.T) {
	res, err := orderSvc.MultiLast(context.Background(), []string{"202505", "202506"}, &order.Last{})
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
