package order

import (
	"context"
	"time"

	"github.com/go-dev-pkg/sn"
	"github.com/shopspring/decimal"

	"example/internal/models"
)

func (o *Order) Create(ctx context.Context) error {
	orderNo := sn.Generate()
	mo := &models.Order{
		Sharding:  orderNo[:6],
		UID:       1,
		OrderNo:   orderNo,
		Status:    0,
		Amount:    decimal.NewFromInt(100),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: 0,
	}
	return o.orderRepo.Create().Values(mo).Do(ctx)
}
