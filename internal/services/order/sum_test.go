package order_test

import (
	"context"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestSum(t *testing.T) {
	sum, err := orderSvc.Sum(context.Background(), time.Now().Format("200601"), &order.Sum{})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(sum)
}

func TestMultiSum(t *testing.T) {
	sum, err := orderSvc.MultiSum(context.Background(), []string{"202505", "202506"}, &order.Sum{})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(sum)
}
