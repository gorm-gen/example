package orderItem

import (
	"context"
	"time"

	"github.com/go-dev-pkg/sn"

	"example/internal/models"
)

func (o *OrderItem) Create(ctx context.Context) error {
	now := time.Now()
	oi := &models.OrderItem{
		Sharding:  now.Year()*100 + int(now.Month()),
		UID:       1,
		OrderNo:   sn.Generate(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	return o.orderItemRepo.Create().Values(oi).Do(ctx)
}
