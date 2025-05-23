package order

import (
	"go.uber.org/zap"

	"example/internal/global"
	"example/internal/query"
	"example/internal/repositories"
	"example/internal/repositories/order"
)

type Order struct {
	q         *query.Query
	logger    *zap.Logger
	orderRepo *order.Order
}

func New() *Order {
	return &Order{
		q:         repositories.GetQuery(),
		logger:    global.Logger,
		orderRepo: order.New(order.WithNewTableName("")),
	}
}
