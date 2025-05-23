package order

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"example/internal/global"
	"example/internal/query"
	"example/internal/repositories"
	"example/internal/repositories/order"
)

type Order struct {
	q         *query.Query
	db        *gorm.DB
	logger    *zap.Logger
	orderRepo *order.Order
}

func New() *Order {
	return &Order{
		q:         repositories.GetQuery(),
		db:        global.DB,
		logger:    global.Logger,
		orderRepo: order.New(order.WithNewTableName(""), order.WithUnscoped()),
	}
}
