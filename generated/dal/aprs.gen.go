// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"
)

func newAprORM(db *gorm.DB, opts ...gen.DOOption) aprORM {
	_aprORM := aprORM{}

	_aprORM.aprORMDo.UseDB(db, opts...)
	_aprORM.aprORMDo.UseModel(&apiv1.AprORM{})

	tableName := _aprORM.aprORMDo.TableName()
	_aprORM.ALL = field.NewAsterisk(tableName)
	_aprORM.BalanceSubjectToApr = field.NewFloat64(tableName, "balance_subject_to_apr")
	_aprORM.CreditAccountId = field.NewUint64(tableName, "credit_account_id")
	_aprORM.Id = field.NewUint64(tableName, "id")
	_aprORM.InterestChargeAmount = field.NewFloat64(tableName, "interest_charge_amount")
	_aprORM.Percentage = field.NewFloat64(tableName, "percentage")
	_aprORM.Type = field.NewString(tableName, "type")

	_aprORM.fillFieldMap()

	return _aprORM
}

type aprORM struct {
	aprORMDo

	ALL                  field.Asterisk
	BalanceSubjectToApr  field.Float64
	CreditAccountId      field.Uint64
	Id                   field.Uint64
	InterestChargeAmount field.Float64
	Percentage           field.Float64
	Type                 field.String

	fieldMap map[string]field.Expr
}

func (a aprORM) Table(newTableName string) *aprORM {
	a.aprORMDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aprORM) As(alias string) *aprORM {
	a.aprORMDo.DO = *(a.aprORMDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aprORM) updateTableName(table string) *aprORM {
	a.ALL = field.NewAsterisk(table)
	a.BalanceSubjectToApr = field.NewFloat64(table, "balance_subject_to_apr")
	a.CreditAccountId = field.NewUint64(table, "credit_account_id")
	a.Id = field.NewUint64(table, "id")
	a.InterestChargeAmount = field.NewFloat64(table, "interest_charge_amount")
	a.Percentage = field.NewFloat64(table, "percentage")
	a.Type = field.NewString(table, "type")

	a.fillFieldMap()

	return a
}

func (a *aprORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aprORM) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["balance_subject_to_apr"] = a.BalanceSubjectToApr
	a.fieldMap["credit_account_id"] = a.CreditAccountId
	a.fieldMap["id"] = a.Id
	a.fieldMap["interest_charge_amount"] = a.InterestChargeAmount
	a.fieldMap["percentage"] = a.Percentage
	a.fieldMap["type"] = a.Type
}

func (a aprORM) clone(db *gorm.DB) aprORM {
	a.aprORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aprORM) replaceDB(db *gorm.DB) aprORM {
	a.aprORMDo.ReplaceDB(db)
	return a
}

type aprORMDo struct{ gen.DO }

type IAprORMDo interface {
	gen.SubQuery
	Debug() IAprORMDo
	WithContext(ctx context.Context) IAprORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAprORMDo
	WriteDB() IAprORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAprORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAprORMDo
	Not(conds ...gen.Condition) IAprORMDo
	Or(conds ...gen.Condition) IAprORMDo
	Select(conds ...field.Expr) IAprORMDo
	Where(conds ...gen.Condition) IAprORMDo
	Order(conds ...field.Expr) IAprORMDo
	Distinct(cols ...field.Expr) IAprORMDo
	Omit(cols ...field.Expr) IAprORMDo
	Join(table schema.Tabler, on ...field.Expr) IAprORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAprORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAprORMDo
	Group(cols ...field.Expr) IAprORMDo
	Having(conds ...gen.Condition) IAprORMDo
	Limit(limit int) IAprORMDo
	Offset(offset int) IAprORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAprORMDo
	Unscoped() IAprORMDo
	Create(values ...*apiv1.AprORM) error
	CreateInBatches(values []*apiv1.AprORM, batchSize int) error
	Save(values ...*apiv1.AprORM) error
	First() (*apiv1.AprORM, error)
	Take() (*apiv1.AprORM, error)
	Last() (*apiv1.AprORM, error)
	Find() ([]*apiv1.AprORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.AprORM, err error)
	FindInBatches(result *[]*apiv1.AprORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*apiv1.AprORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAprORMDo
	Assign(attrs ...field.AssignExpr) IAprORMDo
	Joins(fields ...field.RelationField) IAprORMDo
	Preload(fields ...field.RelationField) IAprORMDo
	FirstOrInit() (*apiv1.AprORM, error)
	FirstOrCreate() (*apiv1.AprORM, error)
	FindByPage(offset int, limit int) (result []*apiv1.AprORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAprORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result apiv1.AprORM, err error)
	GetByID(id int) (result apiv1.AprORM, err error)
	GetByIDs(ids []int) (result []apiv1.AprORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (a aprORMDo) GetByUserID(user_id int) (result apiv1.AprORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM aprs ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (a aprORMDo) GetByID(id int) (result apiv1.AprORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM aprs ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (a aprORMDo) GetByIDs(ids []int) (result []apiv1.AprORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM aprs ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a aprORMDo) Debug() IAprORMDo {
	return a.withDO(a.DO.Debug())
}

func (a aprORMDo) WithContext(ctx context.Context) IAprORMDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aprORMDo) ReadDB() IAprORMDo {
	return a.Clauses(dbresolver.Read)
}

func (a aprORMDo) WriteDB() IAprORMDo {
	return a.Clauses(dbresolver.Write)
}

func (a aprORMDo) Session(config *gorm.Session) IAprORMDo {
	return a.withDO(a.DO.Session(config))
}

func (a aprORMDo) Clauses(conds ...clause.Expression) IAprORMDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aprORMDo) Returning(value interface{}, columns ...string) IAprORMDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aprORMDo) Not(conds ...gen.Condition) IAprORMDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aprORMDo) Or(conds ...gen.Condition) IAprORMDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aprORMDo) Select(conds ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aprORMDo) Where(conds ...gen.Condition) IAprORMDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aprORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAprORMDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a aprORMDo) Order(conds ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aprORMDo) Distinct(cols ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aprORMDo) Omit(cols ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aprORMDo) Join(table schema.Tabler, on ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aprORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aprORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aprORMDo) Group(cols ...field.Expr) IAprORMDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aprORMDo) Having(conds ...gen.Condition) IAprORMDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aprORMDo) Limit(limit int) IAprORMDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aprORMDo) Offset(offset int) IAprORMDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aprORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAprORMDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aprORMDo) Unscoped() IAprORMDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aprORMDo) Create(values ...*apiv1.AprORM) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aprORMDo) CreateInBatches(values []*apiv1.AprORM, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aprORMDo) Save(values ...*apiv1.AprORM) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aprORMDo) First() (*apiv1.AprORM, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.AprORM), nil
	}
}

func (a aprORMDo) Take() (*apiv1.AprORM, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.AprORM), nil
	}
}

func (a aprORMDo) Last() (*apiv1.AprORM, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.AprORM), nil
	}
}

func (a aprORMDo) Find() ([]*apiv1.AprORM, error) {
	result, err := a.DO.Find()
	return result.([]*apiv1.AprORM), err
}

func (a aprORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.AprORM, err error) {
	buf := make([]*apiv1.AprORM, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aprORMDo) FindInBatches(result *[]*apiv1.AprORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aprORMDo) Attrs(attrs ...field.AssignExpr) IAprORMDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aprORMDo) Assign(attrs ...field.AssignExpr) IAprORMDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aprORMDo) Joins(fields ...field.RelationField) IAprORMDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aprORMDo) Preload(fields ...field.RelationField) IAprORMDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aprORMDo) FirstOrInit() (*apiv1.AprORM, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.AprORM), nil
	}
}

func (a aprORMDo) FirstOrCreate() (*apiv1.AprORM, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.AprORM), nil
	}
}

func (a aprORMDo) FindByPage(offset int, limit int) (result []*apiv1.AprORM, count int64, err error) {
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

func (a aprORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aprORMDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aprORMDo) Delete(models ...*apiv1.AprORM) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aprORMDo) withDO(do gen.Dao) *aprORMDo {
	a.DO = *do.(*gen.DO)
	return a
}
