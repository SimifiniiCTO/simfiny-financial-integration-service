// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

func newCategoryORM(db *gorm.DB, opts ...gen.DOOption) categoryORM {
	_categoryORM := categoryORM{}

	_categoryORM.categoryORMDo.UseDB(db, opts...)
	_categoryORM.categoryORMDo.UseModel(&apiv1.CategoryORM{})

	tableName := _categoryORM.categoryORMDo.TableName()
	_categoryORM.ALL = field.NewAsterisk(tableName)
	_categoryORM.BudgetId = field.NewUint64(tableName, "budget_id")
	_categoryORM.Description = field.NewString(tableName, "description")
	_categoryORM.Id = field.NewUint64(tableName, "id")
	_categoryORM.Name = field.NewString(tableName, "name")
	_categoryORM.Subcategories = field.NewField(tableName, "subcategories")

	_categoryORM.fillFieldMap()

	return _categoryORM
}

type categoryORM struct {
	categoryORMDo

	ALL           field.Asterisk
	BudgetId      field.Uint64
	Description   field.String
	Id            field.Uint64
	Name          field.String
	Subcategories field.Field

	fieldMap map[string]field.Expr
}

func (c categoryORM) Table(newTableName string) *categoryORM {
	c.categoryORMDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c categoryORM) As(alias string) *categoryORM {
	c.categoryORMDo.DO = *(c.categoryORMDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *categoryORM) updateTableName(table string) *categoryORM {
	c.ALL = field.NewAsterisk(table)
	c.BudgetId = field.NewUint64(table, "budget_id")
	c.Description = field.NewString(table, "description")
	c.Id = field.NewUint64(table, "id")
	c.Name = field.NewString(table, "name")
	c.Subcategories = field.NewField(table, "subcategories")

	c.fillFieldMap()

	return c
}

func (c *categoryORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *categoryORM) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["budget_id"] = c.BudgetId
	c.fieldMap["description"] = c.Description
	c.fieldMap["id"] = c.Id
	c.fieldMap["name"] = c.Name
	c.fieldMap["subcategories"] = c.Subcategories
}

func (c categoryORM) clone(db *gorm.DB) categoryORM {
	c.categoryORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c categoryORM) replaceDB(db *gorm.DB) categoryORM {
	c.categoryORMDo.ReplaceDB(db)
	return c
}

type categoryORMDo struct{ gen.DO }

type ICategoryORMDo interface {
	gen.SubQuery
	Debug() ICategoryORMDo
	WithContext(ctx context.Context) ICategoryORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICategoryORMDo
	WriteDB() ICategoryORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICategoryORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICategoryORMDo
	Not(conds ...gen.Condition) ICategoryORMDo
	Or(conds ...gen.Condition) ICategoryORMDo
	Select(conds ...field.Expr) ICategoryORMDo
	Where(conds ...gen.Condition) ICategoryORMDo
	Order(conds ...field.Expr) ICategoryORMDo
	Distinct(cols ...field.Expr) ICategoryORMDo
	Omit(cols ...field.Expr) ICategoryORMDo
	Join(table schema.Tabler, on ...field.Expr) ICategoryORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICategoryORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICategoryORMDo
	Group(cols ...field.Expr) ICategoryORMDo
	Having(conds ...gen.Condition) ICategoryORMDo
	Limit(limit int) ICategoryORMDo
	Offset(offset int) ICategoryORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICategoryORMDo
	Unscoped() ICategoryORMDo
	Create(values ...*apiv1.CategoryORM) error
	CreateInBatches(values []*apiv1.CategoryORM, batchSize int) error
	Save(values ...*apiv1.CategoryORM) error
	First() (*apiv1.CategoryORM, error)
	Take() (*apiv1.CategoryORM, error)
	Last() (*apiv1.CategoryORM, error)
	Find() ([]*apiv1.CategoryORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.CategoryORM, err error)
	FindInBatches(result *[]*apiv1.CategoryORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*apiv1.CategoryORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICategoryORMDo
	Assign(attrs ...field.AssignExpr) ICategoryORMDo
	Joins(fields ...field.RelationField) ICategoryORMDo
	Preload(fields ...field.RelationField) ICategoryORMDo
	FirstOrInit() (*apiv1.CategoryORM, error)
	FirstOrCreate() (*apiv1.CategoryORM, error)
	FindByPage(offset int, limit int) (result []*apiv1.CategoryORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICategoryORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result apiv1.CategoryORM, err error)
	GetByID(id int) (result apiv1.CategoryORM, err error)
	GetByIDs(ids []int) (result []apiv1.CategoryORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (c categoryORMDo) GetByUserID(user_id int) (result apiv1.CategoryORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM categories ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (c categoryORMDo) GetByID(id int) (result apiv1.CategoryORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM categories ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (c categoryORMDo) GetByIDs(ids []int) (result []apiv1.CategoryORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM categories ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c categoryORMDo) Debug() ICategoryORMDo {
	return c.withDO(c.DO.Debug())
}

func (c categoryORMDo) WithContext(ctx context.Context) ICategoryORMDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoryORMDo) ReadDB() ICategoryORMDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoryORMDo) WriteDB() ICategoryORMDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoryORMDo) Session(config *gorm.Session) ICategoryORMDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoryORMDo) Clauses(conds ...clause.Expression) ICategoryORMDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoryORMDo) Returning(value interface{}, columns ...string) ICategoryORMDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoryORMDo) Not(conds ...gen.Condition) ICategoryORMDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoryORMDo) Or(conds ...gen.Condition) ICategoryORMDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoryORMDo) Select(conds ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoryORMDo) Where(conds ...gen.Condition) ICategoryORMDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoryORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICategoryORMDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c categoryORMDo) Order(conds ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoryORMDo) Distinct(cols ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoryORMDo) Omit(cols ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoryORMDo) Join(table schema.Tabler, on ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoryORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoryORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoryORMDo) Group(cols ...field.Expr) ICategoryORMDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoryORMDo) Having(conds ...gen.Condition) ICategoryORMDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoryORMDo) Limit(limit int) ICategoryORMDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoryORMDo) Offset(offset int) ICategoryORMDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoryORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICategoryORMDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoryORMDo) Unscoped() ICategoryORMDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoryORMDo) Create(values ...*apiv1.CategoryORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoryORMDo) CreateInBatches(values []*apiv1.CategoryORM, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoryORMDo) Save(values ...*apiv1.CategoryORM) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoryORMDo) First() (*apiv1.CategoryORM, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.CategoryORM), nil
	}
}

func (c categoryORMDo) Take() (*apiv1.CategoryORM, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.CategoryORM), nil
	}
}

func (c categoryORMDo) Last() (*apiv1.CategoryORM, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.CategoryORM), nil
	}
}

func (c categoryORMDo) Find() ([]*apiv1.CategoryORM, error) {
	result, err := c.DO.Find()
	return result.([]*apiv1.CategoryORM), err
}

func (c categoryORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.CategoryORM, err error) {
	buf := make([]*apiv1.CategoryORM, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoryORMDo) FindInBatches(result *[]*apiv1.CategoryORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoryORMDo) Attrs(attrs ...field.AssignExpr) ICategoryORMDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoryORMDo) Assign(attrs ...field.AssignExpr) ICategoryORMDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoryORMDo) Joins(fields ...field.RelationField) ICategoryORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoryORMDo) Preload(fields ...field.RelationField) ICategoryORMDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoryORMDo) FirstOrInit() (*apiv1.CategoryORM, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.CategoryORM), nil
	}
}

func (c categoryORMDo) FirstOrCreate() (*apiv1.CategoryORM, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.CategoryORM), nil
	}
}

func (c categoryORMDo) FindByPage(offset int, limit int) (result []*apiv1.CategoryORM, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c categoryORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoryORMDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoryORMDo) Delete(models ...*apiv1.CategoryORM) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoryORMDo) withDO(do gen.Dao) *categoryORMDo {
	c.DO = *do.(*gen.DO)
	return c
}
