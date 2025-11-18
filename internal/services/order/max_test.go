package order_test

import (
	"context"
	"testing"
)

func TestShardingMax(t *testing.T) {
	sum, err := orderSvc.ShardingMax(context.Background(), shardingList)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(sum)
}
