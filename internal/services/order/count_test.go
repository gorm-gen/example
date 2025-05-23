package order_test

import (
	"context"
	"testing"
)

func TestCount(t *testing.T) {
	count, err := orderSvc.Count(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(count)
}
