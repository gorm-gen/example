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

func TestShardingList(t *testing.T) {
	sharding := shardingList
	list, count, err := orderSvc.ShardingList(context.Background(), sharding, &order.List{
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
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
	t.Log("---------------------------------")
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     5,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     6,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     7,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     8,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     9,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     10,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     11,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     12,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     13,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     14,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     15,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     16,
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
	list, count, err = orderSvc.ShardingList(context.Background(), sharding, &order.List{
		Page:     17,
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
