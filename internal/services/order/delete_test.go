package order_test

import (
	"context"
	"testing"
	"time"

	"example/internal/services/order"
)

func TestDelete(t *testing.T) {
	id := int64(1)
	if err := orderSvc.Delete(context.Background(), time.Now().Format("200601"), &order.Delete{ID: &id}); err != nil {
		t.Error(err)
		return
	}
}
