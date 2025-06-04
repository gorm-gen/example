package order_test

import (
	"context"
	"testing"
)

func TestCreate(t *testing.T) {
	if err := orderSvc.Create(context.Background()); err != nil {
		t.Fatal(err)
		return
	}
}

func TestShardingCreate(t *testing.T) {
	if err := orderSvc.ShardingCreate(context.Background()); err != nil {
		t.Fatal(err)
		return
	}
}
