package order

import (
	"context"
	"runtime/debug"

	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"

	"example/internal/query"

	"example/internal/repositories"
)

type _pluck struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	lock          clause.Expression
	genField      field.Expr
	dest          interface{}
	scopes        []func(gen.Dao) gen.Dao
	unscoped      bool
	orderOpts     []OrderOption
	conditionOpts []ConditionOption
	writeDB       bool
}

// Pluck 从数据库中查询单列并扫描结果到切片
func (o *Order) Pluck(genField field.Expr, dest interface{}) *_pluck {
	return &_pluck{
		core:          o,
		genField:      genField,
		dest:          dest,
		scopes:        make([]func(gen.Dao) gen.Dao, 0),
		unscoped:      o.unscoped,
		orderOpts:     make([]OrderOption, 0),
		conditionOpts: make([]ConditionOption, 0),
	}
}

// Tx 设置为事务
func (p *_pluck) Tx(tx *query.Query) *_pluck {
	p.tx = tx
	if tx != nil {
		p.qTx = nil
	}
	return p
}

// QueryTx 设置为手动事务
func (p *_pluck) QueryTx(tx *query.QueryTx) *_pluck {
	p.qTx = tx
	if tx != nil {
		p.tx = nil
	}
	return p
}

func (p *_pluck) ForUpdate() *_pluck {
	p.lock = clause.Locking{Strength: clause.LockingStrengthUpdate}
	return p
}

func (p *_pluck) ForUpdateSkipLocked() *_pluck {
	p.lock = clause.Locking{Strength: clause.LockingStrengthUpdate, Options: clause.LockingOptionsSkipLocked}
	return p
}

func (p *_pluck) ForUpdateNoWait() *_pluck {
	p.lock = clause.Locking{Strength: clause.LockingStrengthUpdate, Options: clause.LockingOptionsNoWait}
	return p
}

func (p *_pluck) ForShare() *_pluck {
	p.lock = clause.Locking{Strength: clause.LockingStrengthShare}
	return p
}

func (p *_pluck) ForShareSkipLocked() *_pluck {
	p.lock = clause.Locking{Strength: clause.LockingStrengthShare, Options: clause.LockingOptionsSkipLocked}
	return p
}

func (p *_pluck) ForShareNoWait() *_pluck {
	p.lock = clause.Locking{Strength: clause.LockingStrengthShare, Options: clause.LockingOptionsNoWait}
	return p
}

func (p *_pluck) Unscoped() *_pluck {
	p.unscoped = true
	return p
}

func (p *_pluck) Order(opts ...OrderOption) *_pluck {
	p.orderOpts = append(p.orderOpts, opts...)
	return p
}

func (p *_pluck) Where(opts ...ConditionOption) *_pluck {
	p.conditionOpts = append(p.conditionOpts, opts...)
	return p
}

func (p *_pluck) WriteDB() *_pluck {
	p.writeDB = true
	return p
}

func (p *_pluck) Scopes(funcs ...func(gen.Dao) gen.Dao) *_pluck {
	p.scopes = append(p.scopes, funcs...)
	return p
}

// Do 执行获取数据列表
func (p *_pluck) Do(ctx context.Context) error {
	pq := p.core.q.Order
	if p.tx != nil {
		pq = p.tx.Order
	}
	if p.qTx != nil {
		pq = p.qTx.Order
	}
	pr := pq.WithContext(ctx)
	if p.core.newTableName != nil && *p.core.newTableName != "" {
		pr = pq.Table(*p.core.newTableName).WithContext(ctx)
	}
	if p.writeDB {
		pr = pr.WriteDB()
	}
	if p.unscoped {
		pr = pr.Unscoped()
	}
	if len(p.scopes) > 0 {
		pr = pr.Scopes(p.scopes...)
	}
	if (p.tx != nil || p.qTx != nil) && p.lock != nil {
		pr = pr.Clauses(p.lock)
	}
	if _len := len(p.conditionOpts); _len > 0 {
		conditions := make([]gen.Condition, 0, _len)
		for _, opt := range p.conditionOpts {
			conditions = append(conditions, opt(p.core))
		}
		if len(conditions) > 0 {
			pr = pr.Where(conditions...)
		}
	}
	if _len := len(p.orderOpts); _len > 0 {
		orders := make([]field.Expr, 0, _len)
		for _, opt := range p.orderOpts {
			orders = append(orders, opt(p.core))
		}
		if len(orders) > 0 {
			pr = pr.Order(orders...)
		}
	}
	column := field.NewField("", p.genField.ColumnName().String())
	if err := pr.Pluck(column, p.dest); err != nil {
		if repositories.IsRealErr(err) {
			p.core.logger.Error("【Order.Pluck】失败", zap.Error(err), zap.ByteString("debug.Stack", debug.Stack()))
		}
		return err
	}
	return nil
}
