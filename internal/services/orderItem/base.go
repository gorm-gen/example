package orderItem

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"example/internal/global"
	"example/internal/query"
	"example/internal/repositories"
	"example/internal/repositories/orderItem"
)

type OrderItem struct {
	q             *query.Query
	db            *gorm.DB
	logger        *zap.Logger
	orderItemRepo *orderItem.OrderItem
}

func New() *OrderItem {
	return &OrderItem{
		q:             repositories.GetQuery(),
		db:            global.DB,
		logger:        global.Logger,
		orderItemRepo: orderItem.New(),
	}
}
