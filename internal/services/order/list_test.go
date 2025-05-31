package order_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestList(t *testing.T) {
	list, err := orderSvc.List(context.Background(), time.Now().Format("200601"), &order.List{})
	if err != nil {
		t.Error(err)
		return
	}
	var bytes []byte
	if bytes, err = json.Marshal(list); err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
}

func TestMultiList(t *testing.T) {
	sharding := []string{"202505", "202506"}
	list, count, err := orderSvc.MultiList(context.Background(), sharding, &order.List{
		Page:     1,
		PageSize: 2,
	})
	if err != nil {
		t.Error(err)
		return
	}
	var bytes []byte
	if bytes, err = json.Marshal(list); err != nil {
		t.Error(err)
		return
	}
	t.Log(count)
	t.Log("---------------------------------")
	t.Log(string(bytes))
	t.Log("---------------------------------")
	list, count, err = orderSvc.MultiList(context.Background(), sharding, &order.List{
		Page:     2,
		PageSize: 2,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if bytes, err = json.Marshal(list); err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	t.Log("---------------------------------")
	list, count, err = orderSvc.MultiList(context.Background(), sharding, &order.List{
		Page:     3,
		PageSize: 2,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if bytes, err = json.Marshal(list); err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	t.Log("---------------------------------")
	list, count, err = orderSvc.MultiList(context.Background(), sharding, &order.List{
		Page:     4,
		PageSize: 2,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if bytes, err = json.Marshal(list); err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
}
