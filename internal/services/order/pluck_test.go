package order_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestPluck(t *testing.T) {
	res, err := orderSvc.Pluck(context.Background(), time.Now().Format("200601"))
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
