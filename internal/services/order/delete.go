package order

import (
	"context"
	"time"

	"example/internal/repositories/order"
)

type Delete struct {
	ID      *int64
	UID     *int
	OrderNo *string
}

// PhysicalDelete 物理删除/永久删除
func (o *Order) PhysicalDelete(ctx context.Context, sharding string, data *Delete) error {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.UID != nil {
			conditions = append(conditions, order.ConditionUID(*data.UID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
	}
	_, err := o.orderRepo.Delete().
		Where(conditions...).
		Do(ctx)
	return err
}

// Delete 软删除
func (o *Order) Delete(ctx context.Context, sharding string, data *Delete) error {
	conditions := make([]order.ConditionOption, 0)
	conditions = append(conditions, order.ConditionSharding(sharding))
	conditions = append(conditions, order.ConditionDeletedAtIsZero())
	if data != nil {
		if data.ID != nil {
			conditions = append(conditions, order.ConditionID(*data.ID))
		}
		if data.UID != nil {
			conditions = append(conditions, order.ConditionUID(*data.UID))
		}
		if data.OrderNo != nil {
			conditions = append(conditions, order.ConditionOrderNo(*data.OrderNo))
		}
	}
	_, err := o.orderRepo.Update().
		Where(conditions...).
		Update(order.Update(o.q.Order.DeletedAt.Value(uint(time.Now().Unix())))).
		Do(ctx)
	return err
}
