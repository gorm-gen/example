package orderItem_test

import (
	"context"
	"testing"
)

func TestCreate(t *testing.T) {
	if err := orderItemSvc.Create(context.Background()); err != nil {
		t.Error(err)
		return
	}
}
