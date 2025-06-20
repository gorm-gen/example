// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.

package order

import (
	"strings"
	"time"

	f "github.com/gorm-gen/field"
	"github.com/gorm-gen/field/value"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"example/internal/global"

	"example/internal/query"

	"example/internal/repositories"
)

// Order 仓库/Repository
type Order struct {
	q            *query.Query
	db           *gorm.DB
	logger       *zap.Logger
	newTableName *string
}

// Option Order仓库初始化选项
type Option func(*Order)

func WithQuery(q *query.Query) Option {
	return func(o *Order) {
		o.q = q
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(o *Order) {
		o.logger = logger
	}
}

func WithDB(db *gorm.DB) Option {
	return func(o *Order) {
		o.db = db
	}
}

func WithNewTableName(newTableName string) Option {
	return func(o *Order) {
		o.newTableName = &newTableName
	}
}

// New Order仓库初始化
func New(opts ...Option) *Order {
	o := &Order{
		q:      repositories.GetQuery(),
		db:     global.DB,
		logger: global.Logger,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// ConditionOption 字段条件筛选选项
type ConditionOption func(*Order) gen.Condition

// Condition 自定义字段条件筛选
func Condition(condition gen.Condition) ConditionOption {
	return func(*Order) gen.Condition {
		return condition
	}
}

func ConditionID(v ...int64) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).ID.Eq(0)
			}
			if length == 1 {
				return o.q.Order.Table(*o.newTableName).ID.Eq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).ID.In(v...)
		}
		if length == 0 {
			return o.q.Order.ID.Eq(0)
		}
		if length == 1 {
			return o.q.Order.ID.Eq(v[0])
		}
		return o.q.Order.ID.In(v...)
	}
}

func ConditionIDNot(v ...int64) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).ID.Neq(0)
			}
			if length == 1 {
				return o.q.Order.Table(*o.newTableName).ID.Neq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).ID.NotIn(v...)
		}
		if length == 0 {
			return o.q.Order.ID.Neq(0)
		}
		if length == 1 {
			return o.q.Order.ID.Neq(v[0])
		}
		return o.q.Order.ID.NotIn(v...)
	}
}

func ConditionIDGt(v ...int64) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).ID.Gt(0)
			}
			return o.q.Order.Table(*o.newTableName).ID.Gt(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.ID.Gt(0)
		}
		return o.q.Order.ID.Gt(v[0])
	}
}

func ConditionIDGte(v ...int64) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).ID.Gte(0)
			}
			return o.q.Order.Table(*o.newTableName).ID.Gte(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.ID.Gte(0)
		}
		return o.q.Order.ID.Gte(v[0])
	}
}

func ConditionIDLt(v ...int64) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).ID.Lt(0)
			}
			return o.q.Order.Table(*o.newTableName).ID.Lt(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.ID.Lt(0)
		}
		return o.q.Order.ID.Lt(v[0])
	}
}

func ConditionIDLte(v ...int64) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).ID.Lte(0)
			}
			return o.q.Order.Table(*o.newTableName).ID.Lte(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.ID.Lte(0)
		}
		return o.q.Order.ID.Lte(v[0])
	}
}

func ConditionIDBetween(left, right int64) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).ID.Between(left, right)
		}
		return o.q.Order.ID.Between(left, right)
	}
}

func ConditionIDNotBetween(left, right int64) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).ID.NotBetween(left, right)
		}
		return o.q.Order.ID.NotBetween(left, right)
	}
}

func ConditionSharding(v ...string) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).Sharding.Eq("")
			}
			return o.q.Order.Table(*o.newTableName).Sharding.Eq(v[0])
		}
		if length == 0 {
			return o.q.Order.Sharding.Eq("")
		}
		return o.q.Order.Sharding.Eq(v[0])
	}
}

func ConditionShardingNeq(v ...string) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).Sharding.Neq("")
			}
			return o.q.Order.Table(*o.newTableName).Sharding.Neq(v[0])
		}
		if length == 0 {
			return o.q.Order.Sharding.Neq("")
		}
		return o.q.Order.Sharding.Neq(v[0])
	}
}

func ConditionShardingLike(v string) ConditionOption {
	return func(o *Order) gen.Condition {
		if !strings.Contains(v, "%") {
			v = "%" + v + "%"
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Sharding.Like(v)
		}
		return o.q.Order.Sharding.Like(v)
	}
}

func ConditionShardingNotLike(v string) ConditionOption {
	return func(o *Order) gen.Condition {
		if !strings.Contains(v, "%") {
			v = "%" + v + "%"
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Sharding.NotLike(v)
		}
		return o.q.Order.Sharding.NotLike(v)
	}
}

func ConditionUID(v ...int) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).UID.Eq(0)
			}
			if length == 1 {
				return o.q.Order.Table(*o.newTableName).UID.Eq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UID.In(v...)
		}
		if length == 0 {
			return o.q.Order.UID.Eq(0)
		}
		if length == 1 {
			return o.q.Order.UID.Eq(v[0])
		}
		return o.q.Order.UID.In(v...)
	}
}

func ConditionUIDNot(v ...int) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).UID.Neq(0)
			}
			if length == 1 {
				return o.q.Order.Table(*o.newTableName).UID.Neq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UID.NotIn(v...)
		}
		if length == 0 {
			return o.q.Order.UID.Neq(0)
		}
		if length == 1 {
			return o.q.Order.UID.Neq(v[0])
		}
		return o.q.Order.UID.NotIn(v...)
	}
}

func ConditionUIDGt(v ...int) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).UID.Gt(0)
			}
			return o.q.Order.Table(*o.newTableName).UID.Gt(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.UID.Gt(0)
		}
		return o.q.Order.UID.Gt(v[0])
	}
}

func ConditionUIDGte(v ...int) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).UID.Gte(0)
			}
			return o.q.Order.Table(*o.newTableName).UID.Gte(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.UID.Gte(0)
		}
		return o.q.Order.UID.Gte(v[0])
	}
}

func ConditionUIDLt(v ...int) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).UID.Lt(0)
			}
			return o.q.Order.Table(*o.newTableName).UID.Lt(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.UID.Lt(0)
		}
		return o.q.Order.UID.Lt(v[0])
	}
}

func ConditionUIDLte(v ...int) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).UID.Lte(0)
			}
			return o.q.Order.Table(*o.newTableName).UID.Lte(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.UID.Lte(0)
		}
		return o.q.Order.UID.Lte(v[0])
	}
}

func ConditionUIDBetween(left, right int) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.Between(left, right)
		}
		return o.q.Order.UID.Between(left, right)
	}
}

func ConditionUIDNotBetween(left, right int) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.NotBetween(left, right)
		}
		return o.q.Order.UID.NotBetween(left, right)
	}
}

func ConditionOrderNo(v ...string) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).OrderNo.Eq("")
			}
			return o.q.Order.Table(*o.newTableName).OrderNo.Eq(v[0])
		}
		if length == 0 {
			return o.q.Order.OrderNo.Eq("")
		}
		return o.q.Order.OrderNo.Eq(v[0])
	}
}

func ConditionOrderNoNeq(v ...string) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).OrderNo.Neq("")
			}
			return o.q.Order.Table(*o.newTableName).OrderNo.Neq(v[0])
		}
		if length == 0 {
			return o.q.Order.OrderNo.Neq("")
		}
		return o.q.Order.OrderNo.Neq(v[0])
	}
}

func ConditionOrderNoLike(v string) ConditionOption {
	return func(o *Order) gen.Condition {
		if !strings.Contains(v, "%") {
			v = "%" + v + "%"
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).OrderNo.Like(v)
		}
		return o.q.Order.OrderNo.Like(v)
	}
}

func ConditionOrderNoNotLike(v string) ConditionOption {
	return func(o *Order) gen.Condition {
		if !strings.Contains(v, "%") {
			v = "%" + v + "%"
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).OrderNo.NotLike(v)
		}
		return o.q.Order.OrderNo.NotLike(v)
	}
}

func ConditionStatus(v ...int8) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).Status.Eq(0)
			}
			if length == 1 {
				return o.q.Order.Table(*o.newTableName).Status.Eq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).Status.In(v...)
		}
		if length == 0 {
			return o.q.Order.Status.Eq(0)
		}
		if length == 1 {
			return o.q.Order.Status.Eq(v[0])
		}
		return o.q.Order.Status.In(v...)
	}
}

func ConditionStatusNot(v ...int8) ConditionOption {
	return func(o *Order) gen.Condition {
		length := len(v)
		if o.newTableName != nil {
			if length == 0 {
				return o.q.Order.Table(*o.newTableName).Status.Neq(0)
			}
			if length == 1 {
				return o.q.Order.Table(*o.newTableName).Status.Neq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).Status.NotIn(v...)
		}
		if length == 0 {
			return o.q.Order.Status.Neq(0)
		}
		if length == 1 {
			return o.q.Order.Status.Neq(v[0])
		}
		return o.q.Order.Status.NotIn(v...)
	}
}

func ConditionStatusGt(v ...int8) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Status.Gt(0)
			}
			return o.q.Order.Table(*o.newTableName).Status.Gt(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.Status.Gt(0)
		}
		return o.q.Order.Status.Gt(v[0])
	}
}

func ConditionStatusGte(v ...int8) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Status.Gte(0)
			}
			return o.q.Order.Table(*o.newTableName).Status.Gte(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.Status.Gte(0)
		}
		return o.q.Order.Status.Gte(v[0])
	}
}

func ConditionStatusLt(v ...int8) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Status.Lt(0)
			}
			return o.q.Order.Table(*o.newTableName).Status.Lt(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.Status.Lt(0)
		}
		return o.q.Order.Status.Lt(v[0])
	}
}

func ConditionStatusLte(v ...int8) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Status.Lte(0)
			}
			return o.q.Order.Table(*o.newTableName).Status.Lte(v[0])
		}
		if len(v) == 0 {
			return o.q.Order.Status.Lte(0)
		}
		return o.q.Order.Status.Lte(v[0])
	}
}

func ConditionStatusBetween(left, right int8) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.Between(left, right)
		}
		return o.q.Order.Status.Between(left, right)
	}
}

func ConditionStatusNotBetween(left, right int8) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.NotBetween(left, right)
		}
		return o.q.Order.Status.NotBetween(left, right)
	}
}

func ConditionAmount(v decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Amount.Eq(value.NewDecimal(v))
		}
		return o.q.Order.Amount.Eq(value.NewDecimal(v))
	}
}

func ConditionAmountNeq(v decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Amount.Neq(value.NewDecimal(v))
		}
		return o.q.Order.Amount.Neq(value.NewDecimal(v))
	}
}

func ConditionAmountGt(v ...decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Amount.Gt(value.NewDecimal(decimal.Zero))
			}
			return o.q.Order.Table(*o.newTableName).Amount.Gt(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return o.q.Order.Amount.Gt(value.NewDecimal(decimal.Zero))
		}
		return o.q.Order.Amount.Gt(value.NewDecimal(v[0]))
	}
}

func ConditionAmountGte(v ...decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Amount.Gte(value.NewDecimal(decimal.Zero))
			}
			return o.q.Order.Table(*o.newTableName).Amount.Gte(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return o.q.Order.Amount.Gte(value.NewDecimal(decimal.Zero))
		}
		return o.q.Order.Amount.Gte(value.NewDecimal(v[0]))
	}
}

func ConditionAmountLt(v ...decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Amount.Lt(value.NewDecimal(decimal.Zero))
			}
			return o.q.Order.Table(*o.newTableName).Amount.Lt(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return o.q.Order.Amount.Lt(value.NewDecimal(decimal.Zero))
		}
		return o.q.Order.Amount.Lt(value.NewDecimal(v[0]))
	}
}

func ConditionAmountLte(v ...decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) == 0 {
				return o.q.Order.Table(*o.newTableName).Amount.Lte(value.NewDecimal(decimal.Zero))
			}
			return o.q.Order.Table(*o.newTableName).Amount.Lte(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return o.q.Order.Amount.Lte(value.NewDecimal(decimal.Zero))
		}
		return o.q.Order.Amount.Lte(value.NewDecimal(v[0]))
	}
}

func ConditionAmountBetween(left, right decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).Between(left, right)
		}
		return f.NewDecimal(o.q.Order.Amount).Between(left, right)
	}
}

func ConditionAmountNotBetween(left, right decimal.Decimal) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).NotBetween(left, right)
		}
		return f.NewDecimal(o.q.Order.Amount).NotBetween(left, right)
	}
}

func ConditionCreatedAt(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).CreatedAt.Eq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).CreatedAt.Eq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.CreatedAt.Eq(v[0])
		}
		return o.q.Order.CreatedAt.Eq(time.Now())
	}
}

func ConditionCreatedAtNeq(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).CreatedAt.Neq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).CreatedAt.Neq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.CreatedAt.Neq(v[0])
		}
		return o.q.Order.CreatedAt.Neq(time.Now())
	}
}

func ConditionCreatedAtGt(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).CreatedAt.Gt(v[0])
			}
			return o.q.Order.Table(*o.newTableName).CreatedAt.Gt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.CreatedAt.Gt(v[0])
		}
		return o.q.Order.CreatedAt.Gt(time.Now())
	}
}

func ConditionCreatedAtGte(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).CreatedAt.Gte(v[0])
			}
			return o.q.Order.Table(*o.newTableName).CreatedAt.Gte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.CreatedAt.Gte(v[0])
		}
		return o.q.Order.CreatedAt.Gte(time.Now())
	}
}

func ConditionCreatedAtLt(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).CreatedAt.Lt(v[0])
			}
			return o.q.Order.Table(*o.newTableName).CreatedAt.Lt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.CreatedAt.Lt(v[0])
		}
		return o.q.Order.CreatedAt.Lt(time.Now())
	}
}

func ConditionCreatedAtLte(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).CreatedAt.Lte(v[0])
			}
			return o.q.Order.Table(*o.newTableName).CreatedAt.Lte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.CreatedAt.Lte(v[0])
		}
		return o.q.Order.CreatedAt.Lte(time.Now())
	}
}

func ConditionCreatedAtBetween(left, right time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).CreatedAt.Between(left, right)
		}
		return o.q.Order.CreatedAt.Between(left, right)
	}
}

func ConditionCreatedAtNotBetween(left, right time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).CreatedAt.NotBetween(left, right)
		}
		return o.q.Order.CreatedAt.NotBetween(left, right)
	}
}

func ConditionUpdatedAt(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).UpdatedAt.Eq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Eq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.UpdatedAt.Eq(v[0])
		}
		return o.q.Order.UpdatedAt.Eq(time.Now())
	}
}

func ConditionUpdatedAtNeq(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).UpdatedAt.Neq(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Neq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.UpdatedAt.Neq(v[0])
		}
		return o.q.Order.UpdatedAt.Neq(time.Now())
	}
}

func ConditionUpdatedAtGt(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).UpdatedAt.Gt(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Gt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.UpdatedAt.Gt(v[0])
		}
		return o.q.Order.UpdatedAt.Gt(time.Now())
	}
}

func ConditionUpdatedAtGte(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).UpdatedAt.Gte(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Gte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.UpdatedAt.Gte(v[0])
		}
		return o.q.Order.UpdatedAt.Gte(time.Now())
	}
}

func ConditionUpdatedAtLt(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).UpdatedAt.Lt(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Lt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.UpdatedAt.Lt(v[0])
		}
		return o.q.Order.UpdatedAt.Lt(time.Now())
	}
}

func ConditionUpdatedAtLte(v ...time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return o.q.Order.Table(*o.newTableName).UpdatedAt.Lte(v[0])
			}
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Lte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return o.q.Order.UpdatedAt.Lte(v[0])
		}
		return o.q.Order.UpdatedAt.Lte(time.Now())
	}
}

func ConditionUpdatedAtBetween(left, right time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Between(left, right)
		}
		return o.q.Order.UpdatedAt.Between(left, right)
	}
}

func ConditionUpdatedAtNotBetween(left, right time.Time) ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UpdatedAt.NotBetween(left, right)
		}
		return o.q.Order.UpdatedAt.NotBetween(left, right)
	}
}

func ConditionDeletedAtIsZero() ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.DeletedAt, f.WithTableName(*o.newTableName)).Eq(decimal.Zero)
		}
		return f.NewDecimal(o.q.Order.DeletedAt).Eq(decimal.Zero)
	}
}

func ConditionDeletedAtGtZero() ConditionOption {
	return func(o *Order) gen.Condition {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.DeletedAt, f.WithTableName(*o.newTableName)).Gt(decimal.Zero)
		}
		return f.NewDecimal(o.q.Order.DeletedAt).Gt(decimal.Zero)
	}
}

// UpdateOption 数据更新选项
type UpdateOption func(*Order) field.AssignExpr

// Update 自定义数据更新
func Update(update field.AssignExpr) UpdateOption {
	return func(*Order) field.AssignExpr {
		return update
	}
}

func UpdateSharding(v string) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Sharding.Value(v)
		}
		return o.q.Order.Sharding.Value(v)
	}
}

func UpdateUID(v int) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.Value(v)
		}
		return o.q.Order.UID.Value(v)
	}
}

// UpdateUIDAdd +=
func UpdateUIDAdd(v ...int) UpdateOption {
	return func(o *Order) field.AssignExpr {
		_v := int(1)
		if len(v) > 0 {
			_v = v[0]
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.Add(_v)
		}
		return o.q.Order.UID.Add(_v)
	}
}

// UpdateUIDSub -=
func UpdateUIDSub(v ...int) UpdateOption {
	return func(o *Order) field.AssignExpr {
		_v := int(1)
		if len(v) > 0 {
			_v = v[0]
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.Sub(_v)
		}
		return o.q.Order.UID.Sub(_v)
	}
}

// UpdateUIDMul *=
func UpdateUIDMul(v int) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.Mul(v)
		}
		return o.q.Order.UID.Mul(v)
	}
}

// UpdateUIDDiv /=
func UpdateUIDDiv(v int) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UID.Div(v)
		}
		return o.q.Order.UID.Div(v)
	}
}

func UpdateOrderNo(v string) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).OrderNo.Value(v)
		}
		return o.q.Order.OrderNo.Value(v)
	}
}

func UpdateStatus(v int8) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.Value(v)
		}
		return o.q.Order.Status.Value(v)
	}
}

// UpdateStatusAdd +=
func UpdateStatusAdd(v ...int8) UpdateOption {
	return func(o *Order) field.AssignExpr {
		_v := int8(1)
		if len(v) > 0 {
			_v = v[0]
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.Add(_v)
		}
		return o.q.Order.Status.Add(_v)
	}
}

// UpdateStatusSub -=
func UpdateStatusSub(v ...int8) UpdateOption {
	return func(o *Order) field.AssignExpr {
		_v := int8(1)
		if len(v) > 0 {
			_v = v[0]
		}
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.Sub(_v)
		}
		return o.q.Order.Status.Sub(_v)
	}
}

// UpdateStatusMul *=
func UpdateStatusMul(v int8) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.Mul(v)
		}
		return o.q.Order.Status.Mul(v)
	}
}

// UpdateStatusDiv /=
func UpdateStatusDiv(v int8) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).Status.Div(v)
		}
		return o.q.Order.Status.Div(v)
	}
}

func UpdateAmount(v decimal.Decimal) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).Value(v)
		}
		return f.NewDecimal(o.q.Order.Amount).Value(v)
	}
}

// UpdateAmountAdd +=
func UpdateAmountAdd(v decimal.Decimal) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).Add(v)
		}
		return f.NewDecimal(o.q.Order.Amount).Add(v)
	}
}

// UpdateAmountSub -=
func UpdateAmountSub(v decimal.Decimal) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).Sub(v)
		}
		return f.NewDecimal(o.q.Order.Amount).Sub(v)
	}
}

// UpdateAmountMul *=
func UpdateAmountMul(v decimal.Decimal) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).Mul(v)
		}
		return f.NewDecimal(o.q.Order.Amount).Mul(v)
	}
}

// UpdateAmountDiv /=
func UpdateAmountDiv(v decimal.Decimal) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return f.NewDecimal(o.q.Order.Amount, f.WithTableName(*o.newTableName)).Div(v)
		}
		return f.NewDecimal(o.q.Order.Amount).Div(v)
	}
}

func UpdateCreatedAt(v time.Time) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).CreatedAt.Value(v)
		}
		return o.q.Order.CreatedAt.Value(v)
	}
}

func UpdateUpdatedAt(v time.Time) UpdateOption {
	return func(o *Order) field.AssignExpr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Value(v)
		}
		return o.q.Order.UpdatedAt.Value(v)
	}
}

// OrderOption 数据排序选项
type OrderOption func(*Order) field.Expr

// Order 自定义数据排序
func OrderBy(order field.Expr) OrderOption {
	return func(*Order) field.Expr {
		return order
	}
}

func OrderIDAsc() OrderOption {
	return func(o *Order) field.Expr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).ID.Asc()
		}
		return o.q.Order.ID.Asc()
	}
}

func OrderIDDesc() OrderOption {
	return func(o *Order) field.Expr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).ID.Desc()
		}
		return o.q.Order.ID.Desc()
	}
}

func OrderCreatedAtAsc() OrderOption {
	return func(o *Order) field.Expr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).CreatedAt.Asc()
		}
		return o.q.Order.CreatedAt.Asc()
	}
}

func OrderCreatedAtDesc() OrderOption {
	return func(o *Order) field.Expr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).CreatedAt.Desc()
		}
		return o.q.Order.CreatedAt.Desc()
	}
}

func OrderUpdatedAtAsc() OrderOption {
	return func(o *Order) field.Expr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Asc()
		}
		return o.q.Order.UpdatedAt.Asc()
	}
}

func OrderUpdatedAtDesc() OrderOption {
	return func(o *Order) field.Expr {
		if o.newTableName != nil {
			return o.q.Order.Table(*o.newTableName).UpdatedAt.Desc()
		}
		return o.q.Order.UpdatedAt.Desc()
	}
}

// RelationOption 关联模型预加载选项
type RelationOption func(*Order) field.RelationField

// Relation 自定义关联模型预加载
func Relation(relation field.RelationField) RelationOption {
	return func(*Order) field.RelationField {
		return relation
	}
}

func RelationAll() RelationOption {
	return func(*Order) field.RelationField {
		return field.Associations
	}
}
