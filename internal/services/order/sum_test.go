package order_test

import (
	"context"
	"testing"
)

func TestSum(t *testing.T) {
	sum, err := orderSvc.Sum(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(sum)
}
