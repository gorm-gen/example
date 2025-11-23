package order

import (
	"context"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"example/internal/query"

	"example/internal/repositories"
)

type _min struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	genField      field.Expr
	unscoped      bool
	conditionOpts []ConditionOption
	writeDB       bool
	scopes        []func(gen.Dao) gen.Dao
	trace         bool
}

// Min 从数据库中查询最小的值
func (s *Order) Min(genField field.Expr) *_min {
	return &_min{
		core:          s,
		unscoped:      s.unscoped,
		genField:      genField,
		conditionOpts: make([]ConditionOption, 0),
		scopes:        make([]func(gen.Dao) gen.Dao, 0),
	}
}

// Tx 设置为事务
func (m *_min) Tx(tx *query.Query) *_min {
	m.tx = tx
	if tx != nil {
		m.qTx = nil
	}
	return m
}

// QueryTx 设置为手动事务
func (m *_min) QueryTx(tx *query.QueryTx) *_min {
	m.qTx = tx
	if tx != nil {
		m.tx = nil
	}
	return m
}

func (m *_min) Unscoped(unscoped ...bool) *_min {
	_unscoped := true
	if len(unscoped) > 0 {
		_unscoped = unscoped[0]
	}
	m.unscoped = _unscoped
	return m
}

func (m *_min) Where(opts ...ConditionOption) *_min {
	m.conditionOpts = append(m.conditionOpts, opts...)
	return m
}

func (m *_min) Scopes(funcs ...func(gen.Dao) gen.Dao) *_min {
	m.scopes = append(m.scopes, funcs...)
	return m
}

func (m *_min) WriteDB() *_min {
	m.writeDB = true
	return m
}

func (m *_min) Trace() *_min {
	m.trace = true
	return m
}

type Min struct {
	Min decimal.Decimal `json:"min"`
}

// Do 执行从数据库中查询单列并扫描结果到切片
func (m *_min) Do(ctx context.Context) (decimal.Decimal, error) {
	if m.trace {
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			if tracer := opentracing.GlobalTracer(); tracer != nil {
				span := tracer.StartSpan("SQL:Order.Min", opentracing.ChildOf(parent.Context()))
				defer span.Finish()
			}
		}
	}
	mq := m.core.q.Order
	if m.tx != nil {
		mq = m.tx.Order
	}
	if m.qTx != nil {
		mq = m.qTx.Order
	}
	expr := field.NewField("", m.genField.ColumnName().String()).Min().As("min")
	mr := mq.WithContext(ctx)
	if m.core.newTableName != nil && *m.core.newTableName != "" {
		mr = mq.Table(*m.core.newTableName).WithContext(ctx)
	}
	mr = mr.Select(expr)
	if m.writeDB {
		mr = mr.WriteDB()
	}
	if m.unscoped {
		mr = mr.Unscoped()
	}
	if len(m.scopes) > 0 {
		mr = mr.Scopes(m.scopes...)
	}
	if _len := len(m.conditionOpts); _len > 0 {
		conditions := make([]gen.Condition, 0, _len)
		for _, opt := range m.conditionOpts {
			conditions = append(conditions, opt(m.core))
		}
		if len(conditions) > 0 {
			mr = mr.Where(conditions...)
		}
	}
	var data Max
	if err := mr.Scan(&data); err != nil {
		if repositories.IsRealErr(err) {
			m.core.logger.Error("【STask.Min】失败", zap.Error(err), zap.ByteString("debug.Stack", debug.Stack()))
		}
		return decimal.Zero, err
	}
	return data.Max, nil
}
