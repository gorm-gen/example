package order

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"

	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"example/internal/query"
)

type _shardingMax struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	genField      field.Expr
	conditionOpts []ConditionOption
	sharding      []string
	worker        chan struct{}
	writeDB       bool
	scopes        []func(gen.Dao) gen.Dao
	trace         bool
}

// ShardingMax 分表Max数据
func (s *Order) ShardingMax(genField field.Expr, sharding []string) *_shardingMax {
	return &_shardingMax{
		core:          s,
		unscoped:      s.unscoped,
		genField:      genField,
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
		scopes:        make([]func(gen.Dao) gen.Dao, 0),
	}
}

func (m *_shardingMax) Worker(worker chan struct{}) *_shardingMax {
	if worker == nil {
		return m
	}
	m.worker = worker
	return m
}

// Tx 设置为事务
func (m *_shardingMax) Tx(tx *query.Query) *_shardingMax {
	m.tx = tx
	if tx != nil {
		m.qTx = nil
	}
	return m
}

// QueryTx 设置为手动事务
func (m *_shardingMax) QueryTx(tx *query.QueryTx) *_shardingMax {
	m.qTx = tx
	if tx != nil {
		m.tx = nil
	}
	return m
}

func (m *_shardingMax) Unscoped(unscoped ...bool) *_shardingMax {
	_unscoped := true
	if len(unscoped) > 0 {
		_unscoped = unscoped[0]
	}
	m.unscoped = _unscoped
	return m
}

func (m *_shardingMax) Scopes(funcs ...func(gen.Dao) gen.Dao) *_shardingMax {
	m.scopes = append(m.scopes, funcs...)
	return m
}

func (m *_shardingMax) Where(opts ...ConditionOption) *_shardingMax {
	m.conditionOpts = append(m.conditionOpts, opts...)
	return m
}

func (m *_shardingMax) WriteDB() *_shardingMax {
	m.writeDB = true
	return m
}

func (m *_shardingMax) Trace() *_shardingMax {
	m.trace = true
	return m
}

// Do 执行分表SUM数据
func (m *_shardingMax) Do(ctx context.Context) (decimal.Decimal, map[string]decimal.Decimal, error) {
	if m.trace {
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			if tracer := opentracing.GlobalTracer(); tracer != nil {
				span := tracer.StartSpan("SQL:STask.ShardingMax", opentracing.ChildOf(parent.Context()))
				defer span.Finish()
			}
		}
	}
	_lenSharding := len(m.sharding)
	if _lenSharding == 0 {
		return decimal.Zero, nil, nil
	}
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	_condLen := len(m.conditionOpts)
	wg := sync.WaitGroup{}
	sm := sync.Map{}
	errChan := make(chan error, _lenSharding)
	endChan := make(chan struct{}, 1)
	for _, sharding := range m.sharding {
		m.worker <- struct{}{}
		wg.Add(1)
		go func(sharding string) {
			defer func() {
				if r := recover(); r != nil {
					m.core.logger.Error(fmt.Sprintf("【STask.ShardingMax.%d】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
					errChan <- fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			defer func() {
				<-m.worker
			}()
			defer wg.Done()
			_conditionOpts := make([]ConditionOption, _condLen, _condLen+1)
			copy(_conditionOpts, m.conditionOpts)
			_conditionOpts = append(_conditionOpts, ConditionSharding(sharding))
			sr := m.core.Max(m.genField)
			sr.writeDB = m.writeDB
			sum, err := sr.Tx(m.tx).
				QueryTx(m.qTx).
				Unscoped(m.unscoped).
				Scopes(m.scopes...).
				Where(_conditionOpts...).
				Do(ctx)
			if err != nil {
				errChan <- err
				return
			}
			sm.Store(sharding, sum)
			return
		}(sharding)
	}
	go func() {
		wg.Wait()
		endChan <- struct{}{}
	}()
	select {
	case <-endChan:
		_v := decimal.Zero
		_m := make(map[string]decimal.Decimal, _lenSharding)
		_n := 0
		sm.Range(func(key, value interface{}) bool {
			v := value.(decimal.Decimal)
			_m[key.(string)] = v
			if _n == 0 {
				_v = v
				_n++
			}
			if v.GreaterThan(_v) {
				_v = v
			}
			return true
		})
		return _v, _m, nil
	case err := <-errChan:
		return decimal.Zero, nil, err
	case <-ctx.Done():
		return decimal.Zero, nil, ctx.Err()
	}
}
