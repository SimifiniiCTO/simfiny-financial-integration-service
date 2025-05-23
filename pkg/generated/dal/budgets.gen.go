// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	financial_integration_service_apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newBudgetORM(db *gorm.DB, opts ...gen.DOOption) budgetORM {
	_budgetORM := budgetORM{}

	_budgetORM.budgetORMDo.UseDB(db, opts...)
	_budgetORM.budgetORMDo.UseModel(&financial_integration_service_apiv1.BudgetORM{})

	tableName := _budgetORM.budgetORMDo.TableName()
	_budgetORM.ALL = field.NewAsterisk(tableName)
	_budgetORM.Description = field.NewString(tableName, "description")
	_budgetORM.EndDate = field.NewString(tableName, "end_date")
	_budgetORM.Id = field.NewUint64(tableName, "id")
	_budgetORM.MilestoneId = field.NewUint64(tableName, "milestone_id")
	_budgetORM.Name = field.NewString(tableName, "name")
	_budgetORM.StartDate = field.NewString(tableName, "start_date")
	_budgetORM.Category = budgetORMHasOneCategory{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Category", "financial_integration_service_apiv1.CategoryORM"),
	}

	_budgetORM.fillFieldMap()

	return _budgetORM
}

type budgetORM struct {
	budgetORMDo

	ALL         field.Asterisk
	Description field.String
	EndDate     field.String
	Id          field.Uint64
	MilestoneId field.Uint64
	Name        field.String
	StartDate   field.String
	Category    budgetORMHasOneCategory

	fieldMap map[string]field.Expr
}

func (b budgetORM) Table(newTableName string) *budgetORM {
	b.budgetORMDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b budgetORM) As(alias string) *budgetORM {
	b.budgetORMDo.DO = *(b.budgetORMDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *budgetORM) updateTableName(table string) *budgetORM {
	b.ALL = field.NewAsterisk(table)
	b.Description = field.NewString(table, "description")
	b.EndDate = field.NewString(table, "end_date")
	b.Id = field.NewUint64(table, "id")
	b.MilestoneId = field.NewUint64(table, "milestone_id")
	b.Name = field.NewString(table, "name")
	b.StartDate = field.NewString(table, "start_date")

	b.fillFieldMap()

	return b
}

func (b *budgetORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *budgetORM) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 7)
	b.fieldMap["description"] = b.Description
	b.fieldMap["end_date"] = b.EndDate
	b.fieldMap["id"] = b.Id
	b.fieldMap["milestone_id"] = b.MilestoneId
	b.fieldMap["name"] = b.Name
	b.fieldMap["start_date"] = b.StartDate

}

func (b budgetORM) clone(db *gorm.DB) budgetORM {
	b.budgetORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b budgetORM) replaceDB(db *gorm.DB) budgetORM {
	b.budgetORMDo.ReplaceDB(db)
	return b
}

type budgetORMHasOneCategory struct {
	db *gorm.DB

	field.RelationField
}

func (a budgetORMHasOneCategory) Where(conds ...field.Expr) *budgetORMHasOneCategory {
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

func (a budgetORMHasOneCategory) WithContext(ctx context.Context) *budgetORMHasOneCategory {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a budgetORMHasOneCategory) Model(m *financial_integration_service_apiv1.BudgetORM) *budgetORMHasOneCategoryTx {
	return &budgetORMHasOneCategoryTx{a.db.Model(m).Association(a.Name())}
}

type budgetORMHasOneCategoryTx struct{ tx *gorm.Association }

func (a budgetORMHasOneCategoryTx) Find() (result *financial_integration_service_apiv1.CategoryORM, err error) {
	return result, a.tx.Find(&result)
}

func (a budgetORMHasOneCategoryTx) Append(values ...*financial_integration_service_apiv1.CategoryORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a budgetORMHasOneCategoryTx) Replace(values ...*financial_integration_service_apiv1.CategoryORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a budgetORMHasOneCategoryTx) Delete(values ...*financial_integration_service_apiv1.CategoryORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a budgetORMHasOneCategoryTx) Clear() error {
	return a.tx.Clear()
}

func (a budgetORMHasOneCategoryTx) Count() int64 {
	return a.tx.Count()
}

type budgetORMDo struct{ gen.DO }

type IBudgetORMDo interface {
	gen.SubQuery
	Debug() IBudgetORMDo
	WithContext(ctx context.Context) IBudgetORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBudgetORMDo
	WriteDB() IBudgetORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBudgetORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBudgetORMDo
	Not(conds ...gen.Condition) IBudgetORMDo
	Or(conds ...gen.Condition) IBudgetORMDo
	Select(conds ...field.Expr) IBudgetORMDo
	Where(conds ...gen.Condition) IBudgetORMDo
	Order(conds ...field.Expr) IBudgetORMDo
	Distinct(cols ...field.Expr) IBudgetORMDo
	Omit(cols ...field.Expr) IBudgetORMDo
	Join(table schema.Tabler, on ...field.Expr) IBudgetORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBudgetORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBudgetORMDo
	Group(cols ...field.Expr) IBudgetORMDo
	Having(conds ...gen.Condition) IBudgetORMDo
	Limit(limit int) IBudgetORMDo
	Offset(offset int) IBudgetORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBudgetORMDo
	Unscoped() IBudgetORMDo
	Create(values ...*financial_integration_service_apiv1.BudgetORM) error
	CreateInBatches(values []*financial_integration_service_apiv1.BudgetORM, batchSize int) error
	Save(values ...*financial_integration_service_apiv1.BudgetORM) error
	First() (*financial_integration_service_apiv1.BudgetORM, error)
	Take() (*financial_integration_service_apiv1.BudgetORM, error)
	Last() (*financial_integration_service_apiv1.BudgetORM, error)
	Find() ([]*financial_integration_service_apiv1.BudgetORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.BudgetORM, err error)
	FindInBatches(result *[]*financial_integration_service_apiv1.BudgetORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_integration_service_apiv1.BudgetORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBudgetORMDo
	Assign(attrs ...field.AssignExpr) IBudgetORMDo
	Joins(fields ...field.RelationField) IBudgetORMDo
	Preload(fields ...field.RelationField) IBudgetORMDo
	FirstOrInit() (*financial_integration_service_apiv1.BudgetORM, error)
	FirstOrCreate() (*financial_integration_service_apiv1.BudgetORM, error)
	FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.BudgetORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBudgetORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result financial_integration_service_apiv1.BudgetORM, err error)
	GetByID(id int) (result financial_integration_service_apiv1.BudgetORM, err error)
	GetByIDs(ids []int) (result []financial_integration_service_apiv1.BudgetORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (b budgetORMDo) GetByUserID(user_id int) (result financial_integration_service_apiv1.BudgetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM budgets ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (b budgetORMDo) GetByID(id int) (result financial_integration_service_apiv1.BudgetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM budgets ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (b budgetORMDo) GetByIDs(ids []int) (result []financial_integration_service_apiv1.BudgetORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM budgets ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (b budgetORMDo) Debug() IBudgetORMDo {
	return b.withDO(b.DO.Debug())
}

func (b budgetORMDo) WithContext(ctx context.Context) IBudgetORMDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b budgetORMDo) ReadDB() IBudgetORMDo {
	return b.Clauses(dbresolver.Read)
}

func (b budgetORMDo) WriteDB() IBudgetORMDo {
	return b.Clauses(dbresolver.Write)
}

func (b budgetORMDo) Session(config *gorm.Session) IBudgetORMDo {
	return b.withDO(b.DO.Session(config))
}

func (b budgetORMDo) Clauses(conds ...clause.Expression) IBudgetORMDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b budgetORMDo) Returning(value interface{}, columns ...string) IBudgetORMDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b budgetORMDo) Not(conds ...gen.Condition) IBudgetORMDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b budgetORMDo) Or(conds ...gen.Condition) IBudgetORMDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b budgetORMDo) Select(conds ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b budgetORMDo) Where(conds ...gen.Condition) IBudgetORMDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b budgetORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBudgetORMDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b budgetORMDo) Order(conds ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b budgetORMDo) Distinct(cols ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b budgetORMDo) Omit(cols ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b budgetORMDo) Join(table schema.Tabler, on ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b budgetORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b budgetORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b budgetORMDo) Group(cols ...field.Expr) IBudgetORMDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b budgetORMDo) Having(conds ...gen.Condition) IBudgetORMDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b budgetORMDo) Limit(limit int) IBudgetORMDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b budgetORMDo) Offset(offset int) IBudgetORMDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b budgetORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBudgetORMDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b budgetORMDo) Unscoped() IBudgetORMDo {
	return b.withDO(b.DO.Unscoped())
}

func (b budgetORMDo) Create(values ...*financial_integration_service_apiv1.BudgetORM) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b budgetORMDo) CreateInBatches(values []*financial_integration_service_apiv1.BudgetORM, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b budgetORMDo) Save(values ...*financial_integration_service_apiv1.BudgetORM) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b budgetORMDo) First() (*financial_integration_service_apiv1.BudgetORM, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BudgetORM), nil
	}
}

func (b budgetORMDo) Take() (*financial_integration_service_apiv1.BudgetORM, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BudgetORM), nil
	}
}

func (b budgetORMDo) Last() (*financial_integration_service_apiv1.BudgetORM, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BudgetORM), nil
	}
}

func (b budgetORMDo) Find() ([]*financial_integration_service_apiv1.BudgetORM, error) {
	result, err := b.DO.Find()
	return result.([]*financial_integration_service_apiv1.BudgetORM), err
}

func (b budgetORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.BudgetORM, err error) {
	buf := make([]*financial_integration_service_apiv1.BudgetORM, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b budgetORMDo) FindInBatches(result *[]*financial_integration_service_apiv1.BudgetORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b budgetORMDo) Attrs(attrs ...field.AssignExpr) IBudgetORMDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b budgetORMDo) Assign(attrs ...field.AssignExpr) IBudgetORMDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b budgetORMDo) Joins(fields ...field.RelationField) IBudgetORMDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b budgetORMDo) Preload(fields ...field.RelationField) IBudgetORMDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b budgetORMDo) FirstOrInit() (*financial_integration_service_apiv1.BudgetORM, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BudgetORM), nil
	}
}

func (b budgetORMDo) FirstOrCreate() (*financial_integration_service_apiv1.BudgetORM, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BudgetORM), nil
	}
}

func (b budgetORMDo) FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.BudgetORM, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b budgetORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b budgetORMDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b budgetORMDo) Delete(models ...*financial_integration_service_apiv1.BudgetORM) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *budgetORMDo) withDO(do gen.Dao) *budgetORMDo {
	b.DO = *do.(*gen.DO)
	return b
}
