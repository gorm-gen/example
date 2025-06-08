package orderItem_test

import (
	"testing"
)

func TestShardingSuffix(t *testing.T) {
	list, err := orderItemSvc.ShardingSuffix()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(list)
}
