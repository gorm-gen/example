// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q            = new(Query)
	Area         *area
	Classify     *classify
	Company      *company
	CreditCard   *creditCard
	IdentityCard *identityCard
	Language     *language
	Order        *order
	OrderItem    *orderItem
	User         *user
	UserLanguage *userLanguage
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Area = &Q.Area
	Classify = &Q.Classify
	Company = &Q.Company
	CreditCard = &Q.CreditCard
	IdentityCard = &Q.IdentityCard
	Language = &Q.Language
	Order = &Q.Order
	OrderItem = &Q.OrderItem
	User = &Q.User
	UserLanguage = &Q.UserLanguage
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:           db,
		Area:         newArea(db, opts...),
		Classify:     newClassify(db, opts...),
		Company:      newCompany(db, opts...),
		CreditCard:   newCreditCard(db, opts...),
		IdentityCard: newIdentityCard(db, opts...),
		Language:     newLanguage(db, opts...),
		Order:        newOrder(db, opts...),
		OrderItem:    newOrderItem(db, opts...),
		User:         newUser(db, opts...),
		UserLanguage: newUserLanguage(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Area         area
	Classify     classify
	Company      company
	CreditCard   creditCard
	IdentityCard identityCard
	Language     language
	Order        order
	OrderItem    orderItem
	User         user
	UserLanguage userLanguage
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		Area:         q.Area.clone(db),
		Classify:     q.Classify.clone(db),
		Company:      q.Company.clone(db),
		CreditCard:   q.CreditCard.clone(db),
		IdentityCard: q.IdentityCard.clone(db),
		Language:     q.Language.clone(db),
		Order:        q.Order.clone(db),
		OrderItem:    q.OrderItem.clone(db),
		User:         q.User.clone(db),
		UserLanguage: q.UserLanguage.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		Area:         q.Area.replaceDB(db),
		Classify:     q.Classify.replaceDB(db),
		Company:      q.Company.replaceDB(db),
		CreditCard:   q.CreditCard.replaceDB(db),
		IdentityCard: q.IdentityCard.replaceDB(db),
		Language:     q.Language.replaceDB(db),
		Order:        q.Order.replaceDB(db),
		OrderItem:    q.OrderItem.replaceDB(db),
		User:         q.User.replaceDB(db),
		UserLanguage: q.UserLanguage.replaceDB(db),
	}
}

type queryCtx struct {
	Area         IAreaDo
	Classify     IClassifyDo
	Company      ICompanyDo
	CreditCard   ICreditCardDo
	IdentityCard IIdentityCardDo
	Language     ILanguageDo
	Order        IOrderDo
	OrderItem    IOrderItemDo
	User         IUserDo
	UserLanguage IUserLanguageDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Area:         q.Area.WithContext(ctx),
		Classify:     q.Classify.WithContext(ctx),
		Company:      q.Company.WithContext(ctx),
		CreditCard:   q.CreditCard.WithContext(ctx),
		IdentityCard: q.IdentityCard.WithContext(ctx),
		Language:     q.Language.WithContext(ctx),
		Order:        q.Order.WithContext(ctx),
		OrderItem:    q.OrderItem.WithContext(ctx),
		User:         q.User.WithContext(ctx),
		UserLanguage: q.UserLanguage.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
