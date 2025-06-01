package order_test

import (
	"testing"
)

func TestShardingSuffix(t *testing.T) {
	list, err := orderSvc.ShardingSuffix()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(list)
}
