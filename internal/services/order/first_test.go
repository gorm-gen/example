package order_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestFirst(t *testing.T) {
	res, err := orderSvc.First(context.Background(), time.Now().Format("200601"), &order.First{})
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

func TestMultiFirst(t *testing.T) {
	res, err := orderSvc.MultiFirst(context.Background(), []string{"202505", "202506"}, &order.First{})
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
