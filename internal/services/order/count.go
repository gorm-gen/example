package order

import (
	"context"
	"time"

	"example/internal/repositories/order"
)

func (o *Order) Count(ctx context.Context) (int64, error) {
	return o.orderRepo.
		Count().
		Where(
			order.ConditionShardingEq(time.Now().Format("200601")),
			order.ConditionDeletedAtIsZero(),
		).Do(ctx)
}
