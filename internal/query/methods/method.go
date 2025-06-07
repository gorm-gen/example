package methods

import (
	"gorm.io/gen"
)

type Query interface {
	// GetByID
	// SELECT * FROM @@table WHERE `id` = @id AND `deleted_at` = 0
	GetByID(id int) (gen.T, error)
}

type User interface {
	// GetCompanyName
	// SELECT `name` FROM `company` WHERE `id` = (SELECT `id` FROM @@table WHERE `id` = @id AND `deleted_at` = 0) AND `deleted_at` = 0
	GetCompanyName(id int) (gen.M, error)
}
