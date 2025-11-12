package models

import (
	"gorm.io/gorm"
)

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.Sharding == "" {
		o.Sharding = o.OrderNo[:6]
	}
	return nil
}
