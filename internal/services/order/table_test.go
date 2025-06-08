package order_test

import (
	"testing"
	"time"
)

func TestTable(t *testing.T) {
	if err := orderSvc.Table(time.Now().Format("200601"), "../../../resources/sql/order.sql"); err != nil {
		t.Error(err)
		return
	}
}
