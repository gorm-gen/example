// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.

package repositories

import (
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"

	"example/internal/global"

	"example/internal/query"
)

var once sync.Once
var q *query.Query

func GetQuery() *query.Query {
	once.Do(func() {
		q = query.Use(global.DB)
	})
	return q
}

// IsRealErr 是否为非超时和查询不到的错误
func IsRealErr(err error) bool {
	return !errors.Is(err, gorm.ErrRecordNotFound) &&
		!errors.Is(err, context.DeadlineExceeded) &&
		!errors.Is(err, context.Canceled)
}
