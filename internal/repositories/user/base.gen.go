// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.
// Code generated by github.com/gorm-gen/repository. DO NOT EDIT.

package user

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

// User 仓库/Repository
type User struct {
	q            *query.Query
	db           *gorm.DB
	logger       *zap.Logger
	newTableName *string
}

// Option User仓库初始化选项
type Option func(*User)

func WithQuery(q *query.Query) Option {
	return func(u *User) {
		u.q = q
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(u *User) {
		u.logger = logger
	}
}

func WithDB(db *gorm.DB) Option {
	return func(u *User) {
		u.db = db
	}
}

func WithNewTableName(newTableName string) Option {
	return func(u *User) {
		u.newTableName = &newTableName
	}
}

// New User仓库初始化
func New(opts ...Option) *User {
	u := &User{
		q:      repositories.GetQuery(),
		db:     global.DB,
		logger: global.Logger,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

// ConditionOption 字段条件筛选选项
type ConditionOption func(*User) gen.Condition

// Condition 自定义字段条件筛选
func Condition(condition gen.Condition) ConditionOption {
	return func(*User) gen.Condition {
		return condition
	}
}

func ConditionID(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		length := len(v)
		if u.newTableName != nil {
			if length == 0 {
				return u.q.User.Table(*u.newTableName).ID.Eq(0)
			}
			if length == 1 {
				return u.q.User.Table(*u.newTableName).ID.Eq(v[0])
			}
			return u.q.User.Table(*u.newTableName).ID.In(v...)
		}
		if length == 0 {
			return u.q.User.ID.Eq(0)
		}
		if length == 1 {
			return u.q.User.ID.Eq(v[0])
		}
		return u.q.User.ID.In(v...)
	}
}

func ConditionIDNot(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		length := len(v)
		if u.newTableName != nil {
			if length == 0 {
				return u.q.User.Table(*u.newTableName).ID.Neq(0)
			}
			if length == 1 {
				return u.q.User.Table(*u.newTableName).ID.Neq(v[0])
			}
			return u.q.User.Table(*u.newTableName).ID.NotIn(v...)
		}
		if length == 0 {
			return u.q.User.ID.Neq(0)
		}
		if length == 1 {
			return u.q.User.ID.Neq(v[0])
		}
		return u.q.User.ID.NotIn(v...)
	}
}

func ConditionIDGt(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).ID.Gt(0)
			}
			return u.q.User.Table(*u.newTableName).ID.Gt(v[0])
		}
		if len(v) == 0 {
			return u.q.User.ID.Gt(0)
		}
		return u.q.User.ID.Gt(v[0])
	}
}

func ConditionIDGte(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).ID.Gte(0)
			}
			return u.q.User.Table(*u.newTableName).ID.Gte(v[0])
		}
		if len(v) == 0 {
			return u.q.User.ID.Gte(0)
		}
		return u.q.User.ID.Gte(v[0])
	}
}

func ConditionIDLt(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).ID.Lt(0)
			}
			return u.q.User.Table(*u.newTableName).ID.Lt(v[0])
		}
		if len(v) == 0 {
			return u.q.User.ID.Lt(0)
		}
		return u.q.User.ID.Lt(v[0])
	}
}

func ConditionIDLte(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).ID.Lte(0)
			}
			return u.q.User.Table(*u.newTableName).ID.Lte(v[0])
		}
		if len(v) == 0 {
			return u.q.User.ID.Lte(0)
		}
		return u.q.User.ID.Lte(v[0])
	}
}

func ConditionIDBetween(left, right int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).ID.Between(left, right)
		}
		return u.q.User.ID.Between(left, right)
	}
}

func ConditionIDNotBetween(left, right int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).ID.NotBetween(left, right)
		}
		return u.q.User.ID.NotBetween(left, right)
	}
}

func ConditionName(v ...string) ConditionOption {
	return func(u *User) gen.Condition {
		length := len(v)
		if u.newTableName != nil {
			if length == 0 {
				return u.q.User.Table(*u.newTableName).Name.Eq("")
			}
			return u.q.User.Table(*u.newTableName).Name.Eq(v[0])
		}
		if length == 0 {
			return u.q.User.Name.Eq("")
		}
		return u.q.User.Name.Eq(v[0])
	}
}

func ConditionNameNeq(v ...string) ConditionOption {
	return func(u *User) gen.Condition {
		length := len(v)
		if u.newTableName != nil {
			if length == 0 {
				return u.q.User.Table(*u.newTableName).Name.Neq("")
			}
			return u.q.User.Table(*u.newTableName).Name.Neq(v[0])
		}
		if length == 0 {
			return u.q.User.Name.Neq("")
		}
		return u.q.User.Name.Neq(v[0])
	}
}

func ConditionNameLike(v string) ConditionOption {
	return func(u *User) gen.Condition {
		if !strings.Contains(v, "%") {
			v = "%" + v + "%"
		}
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).Name.Like(v)
		}
		return u.q.User.Name.Like(v)
	}
}

func ConditionNameNotLike(v string) ConditionOption {
	return func(u *User) gen.Condition {
		if !strings.Contains(v, "%") {
			v = "%" + v + "%"
		}
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).Name.NotLike(v)
		}
		return u.q.User.Name.NotLike(v)
	}
}

func ConditionCompanyID(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		length := len(v)
		if u.newTableName != nil {
			if length == 0 {
				return u.q.User.Table(*u.newTableName).CompanyID.Eq(0)
			}
			if length == 1 {
				return u.q.User.Table(*u.newTableName).CompanyID.Eq(v[0])
			}
			return u.q.User.Table(*u.newTableName).CompanyID.In(v...)
		}
		if length == 0 {
			return u.q.User.CompanyID.Eq(0)
		}
		if length == 1 {
			return u.q.User.CompanyID.Eq(v[0])
		}
		return u.q.User.CompanyID.In(v...)
	}
}

func ConditionCompanyIDNot(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		length := len(v)
		if u.newTableName != nil {
			if length == 0 {
				return u.q.User.Table(*u.newTableName).CompanyID.Neq(0)
			}
			if length == 1 {
				return u.q.User.Table(*u.newTableName).CompanyID.Neq(v[0])
			}
			return u.q.User.Table(*u.newTableName).CompanyID.NotIn(v...)
		}
		if length == 0 {
			return u.q.User.CompanyID.Neq(0)
		}
		if length == 1 {
			return u.q.User.CompanyID.Neq(v[0])
		}
		return u.q.User.CompanyID.NotIn(v...)
	}
}

func ConditionCompanyIDGt(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).CompanyID.Gt(0)
			}
			return u.q.User.Table(*u.newTableName).CompanyID.Gt(v[0])
		}
		if len(v) == 0 {
			return u.q.User.CompanyID.Gt(0)
		}
		return u.q.User.CompanyID.Gt(v[0])
	}
}

func ConditionCompanyIDGte(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).CompanyID.Gte(0)
			}
			return u.q.User.Table(*u.newTableName).CompanyID.Gte(v[0])
		}
		if len(v) == 0 {
			return u.q.User.CompanyID.Gte(0)
		}
		return u.q.User.CompanyID.Gte(v[0])
	}
}

func ConditionCompanyIDLt(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).CompanyID.Lt(0)
			}
			return u.q.User.Table(*u.newTableName).CompanyID.Lt(v[0])
		}
		if len(v) == 0 {
			return u.q.User.CompanyID.Lt(0)
		}
		return u.q.User.CompanyID.Lt(v[0])
	}
}

func ConditionCompanyIDLte(v ...int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).CompanyID.Lte(0)
			}
			return u.q.User.Table(*u.newTableName).CompanyID.Lte(v[0])
		}
		if len(v) == 0 {
			return u.q.User.CompanyID.Lte(0)
		}
		return u.q.User.CompanyID.Lte(v[0])
	}
}

func ConditionCompanyIDBetween(left, right int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.Between(left, right)
		}
		return u.q.User.CompanyID.Between(left, right)
	}
}

func ConditionCompanyIDNotBetween(left, right int) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.NotBetween(left, right)
		}
		return u.q.User.CompanyID.NotBetween(left, right)
	}
}

func ConditionBalance(v decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).Balance.Eq(value.NewDecimal(v))
		}
		return u.q.User.Balance.Eq(value.NewDecimal(v))
	}
}

func ConditionBalanceNeq(v decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).Balance.Neq(value.NewDecimal(v))
		}
		return u.q.User.Balance.Neq(value.NewDecimal(v))
	}
}

func ConditionBalanceGt(v ...decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).Balance.Gt(value.NewDecimal(decimal.Zero))
			}
			return u.q.User.Table(*u.newTableName).Balance.Gt(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return u.q.User.Balance.Gt(value.NewDecimal(decimal.Zero))
		}
		return u.q.User.Balance.Gt(value.NewDecimal(v[0]))
	}
}

func ConditionBalanceGte(v ...decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).Balance.Gte(value.NewDecimal(decimal.Zero))
			}
			return u.q.User.Table(*u.newTableName).Balance.Gte(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return u.q.User.Balance.Gte(value.NewDecimal(decimal.Zero))
		}
		return u.q.User.Balance.Gte(value.NewDecimal(v[0]))
	}
}

func ConditionBalanceLt(v ...decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).Balance.Lt(value.NewDecimal(decimal.Zero))
			}
			return u.q.User.Table(*u.newTableName).Balance.Lt(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return u.q.User.Balance.Lt(value.NewDecimal(decimal.Zero))
		}
		return u.q.User.Balance.Lt(value.NewDecimal(v[0]))
	}
}

func ConditionBalanceLte(v ...decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) == 0 {
				return u.q.User.Table(*u.newTableName).Balance.Lte(value.NewDecimal(decimal.Zero))
			}
			return u.q.User.Table(*u.newTableName).Balance.Lte(value.NewDecimal(v[0]))
		}
		if len(v) == 0 {
			return u.q.User.Balance.Lte(value.NewDecimal(decimal.Zero))
		}
		return u.q.User.Balance.Lte(value.NewDecimal(v[0]))
	}
}

func ConditionBalanceBetween(left, right decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).Between(left, right)
		}
		return f.NewDecimal(u.q.User.Balance).Between(left, right)
	}
}

func ConditionBalanceNotBetween(left, right decimal.Decimal) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).NotBetween(left, right)
		}
		return f.NewDecimal(u.q.User.Balance).NotBetween(left, right)
	}
}

func ConditionCreatedAt(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).CreatedAt.Eq(v[0])
			}
			return u.q.User.Table(*u.newTableName).CreatedAt.Eq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.CreatedAt.Eq(v[0])
		}
		return u.q.User.CreatedAt.Eq(time.Now())
	}
}

func ConditionCreatedAtNeq(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).CreatedAt.Neq(v[0])
			}
			return u.q.User.Table(*u.newTableName).CreatedAt.Neq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.CreatedAt.Neq(v[0])
		}
		return u.q.User.CreatedAt.Neq(time.Now())
	}
}

func ConditionCreatedAtGt(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).CreatedAt.Gt(v[0])
			}
			return u.q.User.Table(*u.newTableName).CreatedAt.Gt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.CreatedAt.Gt(v[0])
		}
		return u.q.User.CreatedAt.Gt(time.Now())
	}
}

func ConditionCreatedAtGte(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).CreatedAt.Gte(v[0])
			}
			return u.q.User.Table(*u.newTableName).CreatedAt.Gte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.CreatedAt.Gte(v[0])
		}
		return u.q.User.CreatedAt.Gte(time.Now())
	}
}

func ConditionCreatedAtLt(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).CreatedAt.Lt(v[0])
			}
			return u.q.User.Table(*u.newTableName).CreatedAt.Lt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.CreatedAt.Lt(v[0])
		}
		return u.q.User.CreatedAt.Lt(time.Now())
	}
}

func ConditionCreatedAtLte(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).CreatedAt.Lte(v[0])
			}
			return u.q.User.Table(*u.newTableName).CreatedAt.Lte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.CreatedAt.Lte(v[0])
		}
		return u.q.User.CreatedAt.Lte(time.Now())
	}
}

func ConditionCreatedAtBetween(left, right time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CreatedAt.Between(left, right)
		}
		return u.q.User.CreatedAt.Between(left, right)
	}
}

func ConditionCreatedAtNotBetween(left, right time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CreatedAt.NotBetween(left, right)
		}
		return u.q.User.CreatedAt.NotBetween(left, right)
	}
}

func ConditionUpdatedAt(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).UpdatedAt.Eq(v[0])
			}
			return u.q.User.Table(*u.newTableName).UpdatedAt.Eq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.UpdatedAt.Eq(v[0])
		}
		return u.q.User.UpdatedAt.Eq(time.Now())
	}
}

func ConditionUpdatedAtNeq(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).UpdatedAt.Neq(v[0])
			}
			return u.q.User.Table(*u.newTableName).UpdatedAt.Neq(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.UpdatedAt.Neq(v[0])
		}
		return u.q.User.UpdatedAt.Neq(time.Now())
	}
}

func ConditionUpdatedAtGt(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).UpdatedAt.Gt(v[0])
			}
			return u.q.User.Table(*u.newTableName).UpdatedAt.Gt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.UpdatedAt.Gt(v[0])
		}
		return u.q.User.UpdatedAt.Gt(time.Now())
	}
}

func ConditionUpdatedAtGte(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).UpdatedAt.Gte(v[0])
			}
			return u.q.User.Table(*u.newTableName).UpdatedAt.Gte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.UpdatedAt.Gte(v[0])
		}
		return u.q.User.UpdatedAt.Gte(time.Now())
	}
}

func ConditionUpdatedAtLt(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).UpdatedAt.Lt(v[0])
			}
			return u.q.User.Table(*u.newTableName).UpdatedAt.Lt(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.UpdatedAt.Lt(v[0])
		}
		return u.q.User.UpdatedAt.Lt(time.Now())
	}
}

func ConditionUpdatedAtLte(v ...time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			if len(v) > 0 && !v[0].IsZero() {
				return u.q.User.Table(*u.newTableName).UpdatedAt.Lte(v[0])
			}
			return u.q.User.Table(*u.newTableName).UpdatedAt.Lte(time.Now())
		}
		if len(v) > 0 && !v[0].IsZero() {
			return u.q.User.UpdatedAt.Lte(v[0])
		}
		return u.q.User.UpdatedAt.Lte(time.Now())
	}
}

func ConditionUpdatedAtBetween(left, right time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).UpdatedAt.Between(left, right)
		}
		return u.q.User.UpdatedAt.Between(left, right)
	}
}

func ConditionUpdatedAtNotBetween(left, right time.Time) ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).UpdatedAt.NotBetween(left, right)
		}
		return u.q.User.UpdatedAt.NotBetween(left, right)
	}
}

func ConditionDeletedAtIsZero() ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.DeletedAt, f.WithTableName(*u.newTableName)).Eq(decimal.Zero)
		}
		return f.NewDecimal(u.q.User.DeletedAt).Eq(decimal.Zero)
	}
}

func ConditionDeletedAtGtZero() ConditionOption {
	return func(u *User) gen.Condition {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.DeletedAt, f.WithTableName(*u.newTableName)).Gt(decimal.Zero)
		}
		return f.NewDecimal(u.q.User.DeletedAt).Gt(decimal.Zero)
	}
}

// UpdateOption 数据更新选项
type UpdateOption func(*User) field.AssignExpr

// Update 自定义数据更新
func Update(update field.AssignExpr) UpdateOption {
	return func(*User) field.AssignExpr {
		return update
	}
}

func UpdateName(v string) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).Name.Value(v)
		}
		return u.q.User.Name.Value(v)
	}
}

func UpdateCompanyID(v int) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.Value(v)
		}
		return u.q.User.CompanyID.Value(v)
	}
}

// UpdateCompanyIDAdd +=
func UpdateCompanyIDAdd(v ...int) UpdateOption {
	return func(u *User) field.AssignExpr {
		_v := int(1)
		if len(v) > 0 {
			_v = v[0]
		}
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.Add(_v)
		}
		return u.q.User.CompanyID.Add(_v)
	}
}

// UpdateCompanyIDSub -=
func UpdateCompanyIDSub(v ...int) UpdateOption {
	return func(u *User) field.AssignExpr {
		_v := int(1)
		if len(v) > 0 {
			_v = v[0]
		}
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.Sub(_v)
		}
		return u.q.User.CompanyID.Sub(_v)
	}
}

// UpdateCompanyIDMul *=
func UpdateCompanyIDMul(v int) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.Mul(v)
		}
		return u.q.User.CompanyID.Mul(v)
	}
}

// UpdateCompanyIDDiv /=
func UpdateCompanyIDDiv(v int) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CompanyID.Div(v)
		}
		return u.q.User.CompanyID.Div(v)
	}
}

func UpdateBalance(v decimal.Decimal) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).Value(v)
		}
		return f.NewDecimal(u.q.User.Balance).Value(v)
	}
}

// UpdateBalanceAdd +=
func UpdateBalanceAdd(v decimal.Decimal) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).Add(v)
		}
		return f.NewDecimal(u.q.User.Balance).Add(v)
	}
}

// UpdateBalanceSub -=
func UpdateBalanceSub(v decimal.Decimal) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).Sub(v)
		}
		return f.NewDecimal(u.q.User.Balance).Sub(v)
	}
}

// UpdateBalanceMul *=
func UpdateBalanceMul(v decimal.Decimal) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).Mul(v)
		}
		return f.NewDecimal(u.q.User.Balance).Mul(v)
	}
}

// UpdateBalanceDiv /=
func UpdateBalanceDiv(v decimal.Decimal) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return f.NewDecimal(u.q.User.Balance, f.WithTableName(*u.newTableName)).Div(v)
		}
		return f.NewDecimal(u.q.User.Balance).Div(v)
	}
}

func UpdateCreatedAt(v time.Time) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CreatedAt.Value(v)
		}
		return u.q.User.CreatedAt.Value(v)
	}
}

func UpdateUpdatedAt(v time.Time) UpdateOption {
	return func(u *User) field.AssignExpr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).UpdatedAt.Value(v)
		}
		return u.q.User.UpdatedAt.Value(v)
	}
}

// OrderOption 数据排序选项
type OrderOption func(*User) field.Expr

// Order 自定义数据排序
func OrderBy(order field.Expr) OrderOption {
	return func(*User) field.Expr {
		return order
	}
}

func OrderIDAsc() OrderOption {
	return func(u *User) field.Expr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).ID.Asc()
		}
		return u.q.User.ID.Asc()
	}
}

func OrderIDDesc() OrderOption {
	return func(u *User) field.Expr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).ID.Desc()
		}
		return u.q.User.ID.Desc()
	}
}

func OrderCreatedAtAsc() OrderOption {
	return func(u *User) field.Expr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CreatedAt.Asc()
		}
		return u.q.User.CreatedAt.Asc()
	}
}

func OrderCreatedAtDesc() OrderOption {
	return func(u *User) field.Expr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).CreatedAt.Desc()
		}
		return u.q.User.CreatedAt.Desc()
	}
}

func OrderUpdatedAtAsc() OrderOption {
	return func(u *User) field.Expr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).UpdatedAt.Asc()
		}
		return u.q.User.UpdatedAt.Asc()
	}
}

func OrderUpdatedAtDesc() OrderOption {
	return func(u *User) field.Expr {
		if u.newTableName != nil {
			return u.q.User.Table(*u.newTableName).UpdatedAt.Desc()
		}
		return u.q.User.UpdatedAt.Desc()
	}
}

// RelationOption 关联模型预加载选项
type RelationOption func(*User) field.RelationField

// Relation 自定义关联模型预加载
func Relation(relation field.RelationField) RelationOption {
	return func(*User) field.RelationField {
		return relation
	}
}

func RelationAll() RelationOption {
	return func(*User) field.RelationField {
		return field.Associations
	}
}
