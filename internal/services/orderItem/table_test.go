package orderItem_test

import (
	"testing"
	"time"
)

func TestTable(t *testing.T) {
	if err := orderItemSvc.Table(time.Now().Format("200601")); err != nil {
		t.Error(err)
		return
	}
}
