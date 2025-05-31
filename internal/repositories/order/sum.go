package order

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"example/internal/query"

	"example/internal/repositories"
)

type multiSum struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	unscoped      bool
	genField      field.Expr
	conditionOpts []ConditionOption
	sharding      []string
	worker        chan struct{}
}

// MultiSum 多表SUM数据
func (o *Order) MultiSum(genField field.Expr, sharding []string) *multiSum {
	return &multiSum{
		core:          o,
		unscoped:      o.unscoped,
		genField:      genField,
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
	}
}

func (s *multiSum) Worker(worker chan struct{}) *multiSum {
	s.worker = worker
	return s
}

// Tx 设置为事务
func (s *multiSum) Tx(tx *query.Query) *multiSum {
	s.tx = tx
	if tx != nil {
		s.qTx = nil
	}
	return s
}

// QueryTx 设置为手动事务
func (s *multiSum) QueryTx(tx *query.QueryTx) *multiSum {
	s.qTx = tx
	if tx != nil {
		s.tx = nil
	}
	return s
}

func (s *multiSum) Unscoped() *multiSum {
	s.unscoped = true
	return s
}

func (s *multiSum) Where(opts ...ConditionOption) *multiSum {
	s.conditionOpts = append(s.conditionOpts, opts...)
	return s
}

// Do 执行多表SUM数据
func (s *multiSum) Do(ctx context.Context) (decimal.Decimal, error) {
	if len(s.sharding) == 0 {
		return decimal.Zero, nil
	}
	sq := s.core.q.Order
	if s.tx != nil {
		sq = s.tx.Order
	}
	if s.qTx != nil {
		sq = s.qTx.Order
	}
	var conditions []gen.Condition
	if len(s.conditionOpts) > 0 {
		conditions = make([]gen.Condition, 0, len(s.conditionOpts))
		for _, opt := range s.conditionOpts {
			conditions = append(conditions, opt(s.core))
		}
	}
	expr := field.NewField("", s.genField.ColumnName().String()).Sum().As("sum")
	sum := decimal.Zero
	wg := sync.WaitGroup{}
	sm := sync.Map{}
	errChan := make(chan error)
	successChan := make(chan struct{})
	for _, sharding := range s.sharding {
		s.worker <- struct{}{}
		wg.Add(1)
		go func(sharding string) {
			defer func() {
				if r := recover(); r != nil {
					s.core.logger.Error(fmt.Sprintf("【Order.MultiSum.%s】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
					errChan <- fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			defer func() {
				<-s.worker
			}()
			defer wg.Done()
			_conditions := make([]gen.Condition, len(conditions))
			copy(_conditions, conditions)
			_conditions = append(_conditions, ConditionSharding(sharding)(s.core))
			sr := sq.WithContext(ctx).Select(expr)
			if s.unscoped {
				sr = sr.Unscoped()
			}
			var data Sum
			if err := sr.Where(_conditions...).Scan(&data); err != nil {
				if repositories.IsRealErr(err) {
					s.core.logger.Error(fmt.Sprintf("【Order.MultiSum.%s】失败", sharding), zap.Error(err))
				}
				errChan <- err
				return
			}
			sm.Store(sharding, data.Sum)
			return
		}(sharding)
	}
	go func() {
		wg.Wait()
		successChan <- struct{}{}
	}()
	select {
	case <-successChan:
		sm.Range(func(key, value interface{}) bool {
			sum = sum.Add(value.(decimal.Decimal))
			return true
		})
		return sum, nil
	case err := <-errChan:
		return decimal.Zero, err
	}
}
