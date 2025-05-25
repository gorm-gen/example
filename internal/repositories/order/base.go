package order

import (
	"regexp"
)

type MultiOrder struct {
	so       *Order
	sharding []string
}

func NewMultiOrder(opts ...Option) *MultiOrder {
	return &MultiOrder{
		so: New(opts...),
	}
}

func (o *MultiOrder) Sharding(expr string) *MultiOrder {
	reg, _ := regexp.Compile(`^order_\d{6}$`)
	tables, _ := o.so.db.Migrator().GetTables()
	var sharding []string
	for _, table := range tables {
		if reg.MatchString(table) {
			sharding = append(sharding, table[6:])
		}
	}
	o.sharding = sharding
	return o
}
