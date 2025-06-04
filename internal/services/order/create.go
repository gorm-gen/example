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

func (o *Order) ShardingCreate(ctx context.Context) error {
	var mos []*models.Order
	mos = append(mos, &models.Order{
		Sharding:  "202505",
		UID:       1,
		OrderNo:   sn.Generate(),
		Status:    0,
		Amount:    decimal.NewFromInt(100),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: 0,
	})
	mos = append(mos, &models.Order{
		Sharding:  "202506",
		UID:       2,
		OrderNo:   sn.Generate(),
		Status:    0,
		Amount:    decimal.NewFromInt(100),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: 0,
	})
	mos = append(mos, &models.Order{
		Sharding:  "202505",
		UID:       3,
		OrderNo:   sn.Generate(),
		Status:    0,
		Amount:    decimal.NewFromInt(100),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: 0,
	})
	return o.orderRepo.ShardingCreate().Values(mos...).Do(ctx)
}
