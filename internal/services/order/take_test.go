package order_test

import (
	"context"
	"encoding/json"
	"testing"

	"example/internal/services/order"
)

func TestTake(t *testing.T) {
	res, err := orderSvc.Take(context.Background(), &order.Take{OrderNo: "202505231134365741880001"})
	if err != nil {
		t.Fatal(err)
		return
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes))
}
