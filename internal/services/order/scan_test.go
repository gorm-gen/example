package order_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	res, err := orderSvc.Scan(context.Background(), time.Now().Format("200601"))
	if err != nil {
		t.Fatal(err)
		return
	}
	var bytes []byte
	if bytes, err = json.Marshal(res); err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes))
}
