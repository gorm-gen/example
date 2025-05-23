package table

import (
	"log"
	"sync"
	"time"

	"example/internal/services/order"
)

var once sync.Once

func Init() {
	once.Do(func() {
		if err := order.New().Table(time.Now().Format("200601")); err != nil {
			log.Fatal(err)
			return
		}
	})
}
