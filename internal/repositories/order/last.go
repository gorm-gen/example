package order

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"example/internal/query"

	"example/internal/repositories"

	"example/internal/models"
)

type multiLast struct {
	core          *Order
	tx            *query.Query
	qTx           *query.QueryTx
	lock          clause.Expression
	unscoped      bool
	selects       []field.Expr
	conditionOpts []ConditionOption
	sharding      []string
	worker        chan struct{}
}

// MultiLast 获取多表随机最后一条记录（主键降序）
func (o *Order) MultiLast(sharding []string) *multiLast {
	return &multiLast{
		core:          o,
		unscoped:      o.unscoped,
		selects:       make([]field.Expr, 0),
		conditionOpts: make([]ConditionOption, 0),
		sharding:      sharding,
		worker:        make(chan struct{}, runtime.NumCPU()),
	}
}

func (l *multiLast) Worker(worker chan struct{}) *multiLast {
	if worker == nil {
		return l
	}
	l.worker = worker
	return l
}

// Tx 设置为事务
func (l *multiLast) Tx(tx *query.Query) *multiLast {
	l.tx = tx
	if tx != nil {
		l.qTx = nil
	}
	return l
}

// QueryTx 设置为手动事务
func (l *multiLast) QueryTx(tx *query.QueryTx) *multiLast {
	l.qTx = tx
	if tx != nil {
		l.tx = nil
	}
	return l
}

func (l *multiLast) Select(field ...field.Expr) *multiLast {
	l.selects = append(l.selects, field...)
	return l
}

func (l *multiLast) ForUpdate() *multiLast {
	l.lock = clause.Locking{Strength: clause.LockingStrengthUpdate}
	return l
}

func (l *multiLast) ForUpdateSkipLocked() *multiLast {
	l.lock = clause.Locking{Strength: clause.LockingStrengthUpdate, Options: clause.LockingOptionsSkipLocked}
	return l
}

func (l *multiLast) ForUpdateNoWait() *multiLast {
	l.lock = clause.Locking{Strength: clause.LockingStrengthUpdate, Options: clause.LockingOptionsNoWait}
	return l
}

func (l *multiLast) ForShare() *multiLast {
	l.lock = clause.Locking{Strength: clause.LockingStrengthShare}
	return l
}

func (l *multiLast) ForShareSkipLocked() *multiLast {
	l.lock = clause.Locking{Strength: clause.LockingStrengthShare, Options: clause.LockingOptionsSkipLocked}
	return l
}

func (l *multiLast) ForShareNoWait() *multiLast {
	l.lock = clause.Locking{Strength: clause.LockingStrengthShare, Options: clause.LockingOptionsNoWait}
	return l
}

func (l *multiLast) Unscoped() *multiLast {
	l.unscoped = true
	return l
}

func (l *multiLast) Where(opts ...ConditionOption) *multiLast {
	l.conditionOpts = append(l.conditionOpts, opts...)
	return l
}

// Do 执行获取多表随机最后一条记录（主键降序）
func (l *multiLast) Do(ctx context.Context) (*models.Order, error) {
	if len(l.sharding) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	fq := l.core.q.Order
	if l.tx != nil {
		fq = l.tx.Order
	}
	if l.qTx != nil {
		fq = l.qTx.Order
	}
	var conditions []gen.Condition
	if len(l.conditionOpts) > 0 {
		conditions = make([]gen.Condition, 0, len(l.conditionOpts))
		for _, opt := range l.conditionOpts {
			conditions = append(conditions, opt(l.core))
		}
	}
	var fieldExpr []field.Expr
	if len(l.selects) > 0 {
		fieldExpr = make([]field.Expr, 0, len(l.selects))
		if l.core.newTableName == nil {
			fieldExpr = append(fieldExpr, l.selects...)
		} else {
			for _, v := range l.selects {
				fieldExpr = append(fieldExpr, field.NewField(*l.core.newTableName, v.ColumnName().String()))
			}
		}
	}
	wg := sync.WaitGroup{}
	endChan := make(chan struct{})
	errChan := make(chan error)
	resultChan := make(chan *models.Order)
	for _, sharding := range l.sharding {
		l.worker <- struct{}{}
		wg.Add(1)
		go func(sharding string) {
			defer func() {
				if r := recover(); r != nil {
					l.core.logger.Error(fmt.Sprintf("【Order.MultiLast.%s】执行异常", sharding), zap.Any("recover", r), zap.ByteString("debug.Stack", debug.Stack()))
					errChan <- fmt.Errorf("recovered from panic: %v", r)
				}
			}()
			defer func() {
				<-l.worker
			}()
			defer wg.Done()
			_conditions := make([]gen.Condition, len(conditions))
			copy(_conditions, conditions)
			_conditions = append(_conditions, ConditionSharding(sharding)(l.core))
			fr := fq.WithContext(ctx)
			if len(fieldExpr) > 0 {
				fr = fr.Select(fieldExpr...)
			}
			if l.unscoped {
				fr = fr.Unscoped()
			}
			if (l.tx != nil || l.qTx != nil) && l.lock != nil {
				fr = fr.Clauses(l.lock)
			}
			res, err := fr.Where(_conditions...).Last()
			if err != nil {
				if repositories.IsRealErr(err) {
					l.core.logger.Error(fmt.Sprintf("【Order.MultiLast.%s】失败", sharding), zap.Error(err))
				}
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					errChan <- err
				}
				return
			}
			resultChan <- res
			return
		}(sharding)
	}
	go func() {
		wg.Wait()
		endChan <- struct{}{}
	}()
	select {
	case res := <-resultChan:
		return res, nil
	case <-endChan:
		return nil, gorm.ErrRecordNotFound
	case err := <-errChan:
		return nil, err
	}
}
