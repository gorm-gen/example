// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"example/internal/models"
)

func newArea(db *gorm.DB, opts ...gen.DOOption) area {
	_area := area{}

	_area.areaDo.UseDB(db, opts...)
	_area.areaDo.UseModel(&models.Area{})

	tableName := _area.areaDo.TableName()
	_area.ALL = field.NewAsterisk(tableName)
	_area.ID = field.NewInt(tableName, "id")
	_area.Name = field.NewString(tableName, "name")
	_area.Pid = field.NewInt(tableName, "pid")
	_area.Level = field.NewInt8(tableName, "level")
	_area.CreatedAt = field.NewTime(tableName, "created_at")
	_area.UpdatedAt = field.NewTime(tableName, "updated_at")
	_area.DeletedAt = field.NewField(tableName, "deleted_at")
	_area.Parent = areaHasOneParent{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Parent", "models.Area"),
	}

	_area.Child = areaHasManyChild{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Child", "models.Area"),
	}

	_area.fillFieldMap()

	return _area
}

// area 地区表
type area struct {
	areaDo

	ALL       field.Asterisk
	ID        field.Int    // 地区ID
	Name      field.String // 地区名
	Pid       field.Int    // 父级ID
	Level     field.Int8   // 地区等级[1:省,2:市,3:县/区]
	CreatedAt field.Time   // 创建日期
	UpdatedAt field.Time   // 更新日期
	DeletedAt field.Field  // 删除时间戳[0:未删除,非0:删除时间戳]
	Parent    areaHasOneParent

	Child areaHasManyChild

	fieldMap map[string]field.Expr
}

func (a area) Table(newTableName string) *area {
	a.areaDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a area) As(alias string) *area {
	a.areaDo.DO = *(a.areaDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *area) updateTableName(table string) *area {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.Name = field.NewString(table, "name")
	a.Pid = field.NewInt(table, "pid")
	a.Level = field.NewInt8(table, "level")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *area) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *area) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["pid"] = a.Pid
	a.fieldMap["level"] = a.Level
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt

}

func (a area) clone(db *gorm.DB) area {
	a.areaDo.ReplaceConnPool(db.Statement.ConnPool)
	a.Parent.db = db.Session(&gorm.Session{Initialized: true})
	a.Parent.db.Statement.ConnPool = db.Statement.ConnPool
	a.Child.db = db.Session(&gorm.Session{Initialized: true})
	a.Child.db.Statement.ConnPool = db.Statement.ConnPool
	return a
}

func (a area) replaceDB(db *gorm.DB) area {
	a.areaDo.ReplaceDB(db)
	a.Parent.db = db.Session(&gorm.Session{})
	a.Child.db = db.Session(&gorm.Session{})
	return a
}

type areaHasOneParent struct {
	db *gorm.DB

	field.RelationField
}

func (a areaHasOneParent) Where(conds ...field.Expr) *areaHasOneParent {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a areaHasOneParent) WithContext(ctx context.Context) *areaHasOneParent {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a areaHasOneParent) Session(session *gorm.Session) *areaHasOneParent {
	a.db = a.db.Session(session)
	return &a
}

func (a areaHasOneParent) Model(m *models.Area) *areaHasOneParentTx {
	return &areaHasOneParentTx{a.db.Model(m).Association(a.Name())}
}

func (a areaHasOneParent) Unscoped() *areaHasOneParent {
	a.db = a.db.Unscoped()
	return &a
}

type areaHasOneParentTx struct{ tx *gorm.Association }

func (a areaHasOneParentTx) Find() (result *models.Area, err error) {
	return result, a.tx.Find(&result)
}

func (a areaHasOneParentTx) Append(values ...*models.Area) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a areaHasOneParentTx) Replace(values ...*models.Area) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a areaHasOneParentTx) Delete(values ...*models.Area) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a areaHasOneParentTx) Clear() error {
	return a.tx.Clear()
}

func (a areaHasOneParentTx) Count() int64 {
	return a.tx.Count()
}

func (a areaHasOneParentTx) Unscoped() *areaHasOneParentTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type areaHasManyChild struct {
	db *gorm.DB

	field.RelationField
}

func (a areaHasManyChild) Where(conds ...field.Expr) *areaHasManyChild {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a areaHasManyChild) WithContext(ctx context.Context) *areaHasManyChild {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a areaHasManyChild) Session(session *gorm.Session) *areaHasManyChild {
	a.db = a.db.Session(session)
	return &a
}

func (a areaHasManyChild) Model(m *models.Area) *areaHasManyChildTx {
	return &areaHasManyChildTx{a.db.Model(m).Association(a.Name())}
}

func (a areaHasManyChild) Unscoped() *areaHasManyChild {
	a.db = a.db.Unscoped()
	return &a
}

type areaHasManyChildTx struct{ tx *gorm.Association }

func (a areaHasManyChildTx) Find() (result []*models.Area, err error) {
	return result, a.tx.Find(&result)
}

func (a areaHasManyChildTx) Append(values ...*models.Area) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a areaHasManyChildTx) Replace(values ...*models.Area) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a areaHasManyChildTx) Delete(values ...*models.Area) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a areaHasManyChildTx) Clear() error {
	return a.tx.Clear()
}

func (a areaHasManyChildTx) Count() int64 {
	return a.tx.Count()
}

func (a areaHasManyChildTx) Unscoped() *areaHasManyChildTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type areaDo struct{ gen.DO }

type IAreaDo interface {
	gen.SubQuery
	Debug() IAreaDo
	WithContext(ctx context.Context) IAreaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAreaDo
	WriteDB() IAreaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAreaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAreaDo
	Not(conds ...gen.Condition) IAreaDo
	Or(conds ...gen.Condition) IAreaDo
	Select(conds ...field.Expr) IAreaDo
	Where(conds ...gen.Condition) IAreaDo
	Order(conds ...field.Expr) IAreaDo
	Distinct(cols ...field.Expr) IAreaDo
	Omit(cols ...field.Expr) IAreaDo
	Join(table schema.Tabler, on ...field.Expr) IAreaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAreaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAreaDo
	Group(cols ...field.Expr) IAreaDo
	Having(conds ...gen.Condition) IAreaDo
	Limit(limit int) IAreaDo
	Offset(offset int) IAreaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAreaDo
	Unscoped() IAreaDo
	Create(values ...*models.Area) error
	CreateInBatches(values []*models.Area, batchSize int) error
	Save(values ...*models.Area) error
	First() (*models.Area, error)
	Take() (*models.Area, error)
	Last() (*models.Area, error)
	Find() ([]*models.Area, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Area, err error)
	FindInBatches(result *[]*models.Area, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Area) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAreaDo
	Assign(attrs ...field.AssignExpr) IAreaDo
	Joins(fields ...field.RelationField) IAreaDo
	Preload(fields ...field.RelationField) IAreaDo
	FirstOrInit() (*models.Area, error)
	FirstOrCreate() (*models.Area, error)
	FindByPage(offset int, limit int) (result []*models.Area, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAreaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a areaDo) Debug() IAreaDo {
	return a.withDO(a.DO.Debug())
}

func (a areaDo) WithContext(ctx context.Context) IAreaDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a areaDo) ReadDB() IAreaDo {
	return a.Clauses(dbresolver.Read)
}

func (a areaDo) WriteDB() IAreaDo {
	return a.Clauses(dbresolver.Write)
}

func (a areaDo) Session(config *gorm.Session) IAreaDo {
	return a.withDO(a.DO.Session(config))
}

func (a areaDo) Clauses(conds ...clause.Expression) IAreaDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a areaDo) Returning(value interface{}, columns ...string) IAreaDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a areaDo) Not(conds ...gen.Condition) IAreaDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a areaDo) Or(conds ...gen.Condition) IAreaDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a areaDo) Select(conds ...field.Expr) IAreaDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a areaDo) Where(conds ...gen.Condition) IAreaDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a areaDo) Order(conds ...field.Expr) IAreaDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a areaDo) Distinct(cols ...field.Expr) IAreaDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a areaDo) Omit(cols ...field.Expr) IAreaDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a areaDo) Join(table schema.Tabler, on ...field.Expr) IAreaDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a areaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAreaDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a areaDo) RightJoin(table schema.Tabler, on ...field.Expr) IAreaDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a areaDo) Group(cols ...field.Expr) IAreaDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a areaDo) Having(conds ...gen.Condition) IAreaDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a areaDo) Limit(limit int) IAreaDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a areaDo) Offset(offset int) IAreaDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a areaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAreaDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a areaDo) Unscoped() IAreaDo {
	return a.withDO(a.DO.Unscoped())
}

func (a areaDo) Create(values ...*models.Area) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a areaDo) CreateInBatches(values []*models.Area, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a areaDo) Save(values ...*models.Area) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a areaDo) First() (*models.Area, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Area), nil
	}
}

func (a areaDo) Take() (*models.Area, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Area), nil
	}
}

func (a areaDo) Last() (*models.Area, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Area), nil
	}
}

func (a areaDo) Find() ([]*models.Area, error) {
	result, err := a.DO.Find()
	return result.([]*models.Area), err
}

func (a areaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Area, err error) {
	buf := make([]*models.Area, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a areaDo) FindInBatches(result *[]*models.Area, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a areaDo) Attrs(attrs ...field.AssignExpr) IAreaDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a areaDo) Assign(attrs ...field.AssignExpr) IAreaDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a areaDo) Joins(fields ...field.RelationField) IAreaDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a areaDo) Preload(fields ...field.RelationField) IAreaDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a areaDo) FirstOrInit() (*models.Area, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Area), nil
	}
}

func (a areaDo) FirstOrCreate() (*models.Area, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Area), nil
	}
}

func (a areaDo) FindByPage(offset int, limit int) (result []*models.Area, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a areaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a areaDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a areaDo) Delete(models ...*models.Area) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *areaDo) withDO(do gen.Dao) *areaDo {
	a.DO = *do.(*gen.DO)
	return a
}
