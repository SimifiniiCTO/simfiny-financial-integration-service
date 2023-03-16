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

func newInvesmentHoldingORM(db *gorm.DB, opts ...gen.DOOption) invesmentHoldingORM {
	_invesmentHoldingORM := invesmentHoldingORM{}

	_invesmentHoldingORM.invesmentHoldingORMDo.UseDB(db, opts...)
	_invesmentHoldingORM.invesmentHoldingORMDo.UseModel(&apiv1.InvesmentHoldingORM{})

	tableName := _invesmentHoldingORM.invesmentHoldingORMDo.TableName()
	_invesmentHoldingORM.ALL = field.NewAsterisk(tableName)
	_invesmentHoldingORM.CostBasis = field.NewFloat64(tableName, "cost_basis")
	_invesmentHoldingORM.Id = field.NewUint64(tableName, "id")
	_invesmentHoldingORM.InstitutionPrice = field.NewFloat64(tableName, "institution_price")
	_invesmentHoldingORM.InstitutionPriceAsOf = field.NewString(tableName, "institution_price_as_of")
	_invesmentHoldingORM.InstitutionPriceDatetime = field.NewString(tableName, "institution_price_datetime")
	_invesmentHoldingORM.InstitutionValue = field.NewFloat64(tableName, "institution_value")
	_invesmentHoldingORM.InvestmentAccountId = field.NewUint64(tableName, "investment_account_id")
	_invesmentHoldingORM.IsoCurrencyCode = field.NewString(tableName, "iso_currency_code")
	_invesmentHoldingORM.Name = field.NewString(tableName, "name")
	_invesmentHoldingORM.PlaidAccountId = field.NewString(tableName, "plaid_account_id")
	_invesmentHoldingORM.Quantity = field.NewFloat64(tableName, "quantity")
	_invesmentHoldingORM.SecurityId = field.NewString(tableName, "security_id")
	_invesmentHoldingORM.UnofficialCurrencyCode = field.NewString(tableName, "unofficial_currency_code")

	_invesmentHoldingORM.fillFieldMap()

	return _invesmentHoldingORM
}

type invesmentHoldingORM struct {
	invesmentHoldingORMDo

	ALL                      field.Asterisk
	CostBasis                field.Float64
	Id                       field.Uint64
	InstitutionPrice         field.Float64
	InstitutionPriceAsOf     field.String
	InstitutionPriceDatetime field.String
	InstitutionValue         field.Float64
	InvestmentAccountId      field.Uint64
	IsoCurrencyCode          field.String
	Name                     field.String
	PlaidAccountId           field.String
	Quantity                 field.Float64
	SecurityId               field.String
	UnofficialCurrencyCode   field.String

	fieldMap map[string]field.Expr
}

func (i invesmentHoldingORM) Table(newTableName string) *invesmentHoldingORM {
	i.invesmentHoldingORMDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i invesmentHoldingORM) As(alias string) *invesmentHoldingORM {
	i.invesmentHoldingORMDo.DO = *(i.invesmentHoldingORMDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *invesmentHoldingORM) updateTableName(table string) *invesmentHoldingORM {
	i.ALL = field.NewAsterisk(table)
	i.CostBasis = field.NewFloat64(table, "cost_basis")
	i.Id = field.NewUint64(table, "id")
	i.InstitutionPrice = field.NewFloat64(table, "institution_price")
	i.InstitutionPriceAsOf = field.NewString(table, "institution_price_as_of")
	i.InstitutionPriceDatetime = field.NewString(table, "institution_price_datetime")
	i.InstitutionValue = field.NewFloat64(table, "institution_value")
	i.InvestmentAccountId = field.NewUint64(table, "investment_account_id")
	i.IsoCurrencyCode = field.NewString(table, "iso_currency_code")
	i.Name = field.NewString(table, "name")
	i.PlaidAccountId = field.NewString(table, "plaid_account_id")
	i.Quantity = field.NewFloat64(table, "quantity")
	i.SecurityId = field.NewString(table, "security_id")
	i.UnofficialCurrencyCode = field.NewString(table, "unofficial_currency_code")

	i.fillFieldMap()

	return i
}

func (i *invesmentHoldingORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *invesmentHoldingORM) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 13)
	i.fieldMap["cost_basis"] = i.CostBasis
	i.fieldMap["id"] = i.Id
	i.fieldMap["institution_price"] = i.InstitutionPrice
	i.fieldMap["institution_price_as_of"] = i.InstitutionPriceAsOf
	i.fieldMap["institution_price_datetime"] = i.InstitutionPriceDatetime
	i.fieldMap["institution_value"] = i.InstitutionValue
	i.fieldMap["investment_account_id"] = i.InvestmentAccountId
	i.fieldMap["iso_currency_code"] = i.IsoCurrencyCode
	i.fieldMap["name"] = i.Name
	i.fieldMap["plaid_account_id"] = i.PlaidAccountId
	i.fieldMap["quantity"] = i.Quantity
	i.fieldMap["security_id"] = i.SecurityId
	i.fieldMap["unofficial_currency_code"] = i.UnofficialCurrencyCode
}

func (i invesmentHoldingORM) clone(db *gorm.DB) invesmentHoldingORM {
	i.invesmentHoldingORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i invesmentHoldingORM) replaceDB(db *gorm.DB) invesmentHoldingORM {
	i.invesmentHoldingORMDo.ReplaceDB(db)
	return i
}

type invesmentHoldingORMDo struct{ gen.DO }

type IInvesmentHoldingORMDo interface {
	gen.SubQuery
	Debug() IInvesmentHoldingORMDo
	WithContext(ctx context.Context) IInvesmentHoldingORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInvesmentHoldingORMDo
	WriteDB() IInvesmentHoldingORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInvesmentHoldingORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInvesmentHoldingORMDo
	Not(conds ...gen.Condition) IInvesmentHoldingORMDo
	Or(conds ...gen.Condition) IInvesmentHoldingORMDo
	Select(conds ...field.Expr) IInvesmentHoldingORMDo
	Where(conds ...gen.Condition) IInvesmentHoldingORMDo
	Order(conds ...field.Expr) IInvesmentHoldingORMDo
	Distinct(cols ...field.Expr) IInvesmentHoldingORMDo
	Omit(cols ...field.Expr) IInvesmentHoldingORMDo
	Join(table schema.Tabler, on ...field.Expr) IInvesmentHoldingORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInvesmentHoldingORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInvesmentHoldingORMDo
	Group(cols ...field.Expr) IInvesmentHoldingORMDo
	Having(conds ...gen.Condition) IInvesmentHoldingORMDo
	Limit(limit int) IInvesmentHoldingORMDo
	Offset(offset int) IInvesmentHoldingORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInvesmentHoldingORMDo
	Unscoped() IInvesmentHoldingORMDo
	Create(values ...*apiv1.InvesmentHoldingORM) error
	CreateInBatches(values []*apiv1.InvesmentHoldingORM, batchSize int) error
	Save(values ...*apiv1.InvesmentHoldingORM) error
	First() (*apiv1.InvesmentHoldingORM, error)
	Take() (*apiv1.InvesmentHoldingORM, error)
	Last() (*apiv1.InvesmentHoldingORM, error)
	Find() ([]*apiv1.InvesmentHoldingORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.InvesmentHoldingORM, err error)
	FindInBatches(result *[]*apiv1.InvesmentHoldingORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*apiv1.InvesmentHoldingORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInvesmentHoldingORMDo
	Assign(attrs ...field.AssignExpr) IInvesmentHoldingORMDo
	Joins(fields ...field.RelationField) IInvesmentHoldingORMDo
	Preload(fields ...field.RelationField) IInvesmentHoldingORMDo
	FirstOrInit() (*apiv1.InvesmentHoldingORM, error)
	FirstOrCreate() (*apiv1.InvesmentHoldingORM, error)
	FindByPage(offset int, limit int) (result []*apiv1.InvesmentHoldingORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInvesmentHoldingORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result apiv1.InvesmentHoldingORM, err error)
	GetByID(id int) (result apiv1.InvesmentHoldingORM, err error)
	GetByIDs(ids []int) (result []apiv1.InvesmentHoldingORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (i invesmentHoldingORMDo) GetByUserID(user_id int) (result apiv1.InvesmentHoldingORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM invesment_holdings ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (i invesmentHoldingORMDo) GetByID(id int) (result apiv1.InvesmentHoldingORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM invesment_holdings ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (i invesmentHoldingORMDo) GetByIDs(ids []int) (result []apiv1.InvesmentHoldingORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM invesment_holdings ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (i invesmentHoldingORMDo) Debug() IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Debug())
}

func (i invesmentHoldingORMDo) WithContext(ctx context.Context) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i invesmentHoldingORMDo) ReadDB() IInvesmentHoldingORMDo {
	return i.Clauses(dbresolver.Read)
}

func (i invesmentHoldingORMDo) WriteDB() IInvesmentHoldingORMDo {
	return i.Clauses(dbresolver.Write)
}

func (i invesmentHoldingORMDo) Session(config *gorm.Session) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Session(config))
}

func (i invesmentHoldingORMDo) Clauses(conds ...clause.Expression) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i invesmentHoldingORMDo) Returning(value interface{}, columns ...string) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i invesmentHoldingORMDo) Not(conds ...gen.Condition) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i invesmentHoldingORMDo) Or(conds ...gen.Condition) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i invesmentHoldingORMDo) Select(conds ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i invesmentHoldingORMDo) Where(conds ...gen.Condition) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i invesmentHoldingORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IInvesmentHoldingORMDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i invesmentHoldingORMDo) Order(conds ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i invesmentHoldingORMDo) Distinct(cols ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i invesmentHoldingORMDo) Omit(cols ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i invesmentHoldingORMDo) Join(table schema.Tabler, on ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i invesmentHoldingORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i invesmentHoldingORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i invesmentHoldingORMDo) Group(cols ...field.Expr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i invesmentHoldingORMDo) Having(conds ...gen.Condition) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i invesmentHoldingORMDo) Limit(limit int) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i invesmentHoldingORMDo) Offset(offset int) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i invesmentHoldingORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i invesmentHoldingORMDo) Unscoped() IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Unscoped())
}

func (i invesmentHoldingORMDo) Create(values ...*apiv1.InvesmentHoldingORM) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i invesmentHoldingORMDo) CreateInBatches(values []*apiv1.InvesmentHoldingORM, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i invesmentHoldingORMDo) Save(values ...*apiv1.InvesmentHoldingORM) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i invesmentHoldingORMDo) First() (*apiv1.InvesmentHoldingORM, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.InvesmentHoldingORM), nil
	}
}

func (i invesmentHoldingORMDo) Take() (*apiv1.InvesmentHoldingORM, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.InvesmentHoldingORM), nil
	}
}

func (i invesmentHoldingORMDo) Last() (*apiv1.InvesmentHoldingORM, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.InvesmentHoldingORM), nil
	}
}

func (i invesmentHoldingORMDo) Find() ([]*apiv1.InvesmentHoldingORM, error) {
	result, err := i.DO.Find()
	return result.([]*apiv1.InvesmentHoldingORM), err
}

func (i invesmentHoldingORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.InvesmentHoldingORM, err error) {
	buf := make([]*apiv1.InvesmentHoldingORM, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i invesmentHoldingORMDo) FindInBatches(result *[]*apiv1.InvesmentHoldingORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i invesmentHoldingORMDo) Attrs(attrs ...field.AssignExpr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i invesmentHoldingORMDo) Assign(attrs ...field.AssignExpr) IInvesmentHoldingORMDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i invesmentHoldingORMDo) Joins(fields ...field.RelationField) IInvesmentHoldingORMDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i invesmentHoldingORMDo) Preload(fields ...field.RelationField) IInvesmentHoldingORMDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i invesmentHoldingORMDo) FirstOrInit() (*apiv1.InvesmentHoldingORM, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.InvesmentHoldingORM), nil
	}
}

func (i invesmentHoldingORMDo) FirstOrCreate() (*apiv1.InvesmentHoldingORM, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.InvesmentHoldingORM), nil
	}
}

func (i invesmentHoldingORMDo) FindByPage(offset int, limit int) (result []*apiv1.InvesmentHoldingORM, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i invesmentHoldingORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i invesmentHoldingORMDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i invesmentHoldingORMDo) Delete(models ...*apiv1.InvesmentHoldingORM) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *invesmentHoldingORMDo) withDO(do gen.Dao) *invesmentHoldingORMDo {
	i.DO = *do.(*gen.DO)
	return i
}
