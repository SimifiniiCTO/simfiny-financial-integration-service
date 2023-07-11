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

	financial_integration_service_apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func newReOccuringTransactionORM(db *gorm.DB, opts ...gen.DOOption) reOccuringTransactionORM {
	_reOccuringTransactionORM := reOccuringTransactionORM{}

	_reOccuringTransactionORM.reOccuringTransactionORMDo.UseDB(db, opts...)
	_reOccuringTransactionORM.reOccuringTransactionORMDo.UseModel(&financial_integration_service_apiv1.ReOccuringTransactionORM{})

	tableName := _reOccuringTransactionORM.reOccuringTransactionORMDo.TableName()
	_reOccuringTransactionORM.ALL = field.NewAsterisk(tableName)
	_reOccuringTransactionORM.AccountId = field.NewString(tableName, "account_id")
	_reOccuringTransactionORM.AverageAmount = field.NewString(tableName, "average_amount")
	_reOccuringTransactionORM.AverageAmountIsoCurrencyCode = field.NewString(tableName, "average_amount_iso_currency_code")
	_reOccuringTransactionORM.CategoryId = field.NewString(tableName, "category_id")
	_reOccuringTransactionORM.Description = field.NewString(tableName, "description")
	_reOccuringTransactionORM.FirstDate = field.NewString(tableName, "first_date")
	_reOccuringTransactionORM.Flow = field.NewString(tableName, "flow")
	_reOccuringTransactionORM.Frequency = field.NewString(tableName, "frequency")
	_reOccuringTransactionORM.Id = field.NewString(tableName, "id")
	_reOccuringTransactionORM.IsActive = field.NewBool(tableName, "is_active")
	_reOccuringTransactionORM.LastAmount = field.NewString(tableName, "last_amount")
	_reOccuringTransactionORM.LastAmountIsoCurrencyCode = field.NewString(tableName, "last_amount_iso_currency_code")
	_reOccuringTransactionORM.LastDate = field.NewString(tableName, "last_date")
	_reOccuringTransactionORM.LinkId = field.NewUint64(tableName, "link_id")
	_reOccuringTransactionORM.MerchantName = field.NewString(tableName, "merchant_name")
	_reOccuringTransactionORM.PersonalFinanceCategoryDetailed = field.NewString(tableName, "personal_finance_category_detailed")
	_reOccuringTransactionORM.PersonalFinanceCategoryPrimary = field.NewString(tableName, "personal_finance_category_primary")
	_reOccuringTransactionORM.Sign = field.NewInt32(tableName, "sign")
	_reOccuringTransactionORM.Status = field.NewString(tableName, "status")
	_reOccuringTransactionORM.StreamId = field.NewString(tableName, "stream_id")
	_reOccuringTransactionORM.Time = field.NewTime(tableName, "time")
	_reOccuringTransactionORM.TransactionIds = field.NewString(tableName, "transaction_ids")
	_reOccuringTransactionORM.UpdatedTime = field.NewString(tableName, "updated_time")
	_reOccuringTransactionORM.UserId = field.NewUint64(tableName, "user_id")

	_reOccuringTransactionORM.fillFieldMap()

	return _reOccuringTransactionORM
}

type reOccuringTransactionORM struct {
	reOccuringTransactionORMDo

	ALL                             field.Asterisk
	AccountId                       field.String
	AverageAmount                   field.String
	AverageAmountIsoCurrencyCode    field.String
	CategoryId                      field.String
	Description                     field.String
	FirstDate                       field.String
	Flow                            field.String
	Frequency                       field.String
	Id                              field.String
	IsActive                        field.Bool
	LastAmount                      field.String
	LastAmountIsoCurrencyCode       field.String
	LastDate                        field.String
	LinkId                          field.Uint64
	MerchantName                    field.String
	PersonalFinanceCategoryDetailed field.String
	PersonalFinanceCategoryPrimary  field.String
	Sign                            field.Int32
	Status                          field.String
	StreamId                        field.String
	Time                            field.Time
	TransactionIds                  field.String
	UpdatedTime                     field.String
	UserId                          field.Uint64

	fieldMap map[string]field.Expr
}

func (r reOccuringTransactionORM) Table(newTableName string) *reOccuringTransactionORM {
	r.reOccuringTransactionORMDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reOccuringTransactionORM) As(alias string) *reOccuringTransactionORM {
	r.reOccuringTransactionORMDo.DO = *(r.reOccuringTransactionORMDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reOccuringTransactionORM) updateTableName(table string) *reOccuringTransactionORM {
	r.ALL = field.NewAsterisk(table)
	r.AccountId = field.NewString(table, "account_id")
	r.AverageAmount = field.NewString(table, "average_amount")
	r.AverageAmountIsoCurrencyCode = field.NewString(table, "average_amount_iso_currency_code")
	r.CategoryId = field.NewString(table, "category_id")
	r.Description = field.NewString(table, "description")
	r.FirstDate = field.NewString(table, "first_date")
	r.Flow = field.NewString(table, "flow")
	r.Frequency = field.NewString(table, "frequency")
	r.Id = field.NewString(table, "id")
	r.IsActive = field.NewBool(table, "is_active")
	r.LastAmount = field.NewString(table, "last_amount")
	r.LastAmountIsoCurrencyCode = field.NewString(table, "last_amount_iso_currency_code")
	r.LastDate = field.NewString(table, "last_date")
	r.LinkId = field.NewUint64(table, "link_id")
	r.MerchantName = field.NewString(table, "merchant_name")
	r.PersonalFinanceCategoryDetailed = field.NewString(table, "personal_finance_category_detailed")
	r.PersonalFinanceCategoryPrimary = field.NewString(table, "personal_finance_category_primary")
	r.Sign = field.NewInt32(table, "sign")
	r.Status = field.NewString(table, "status")
	r.StreamId = field.NewString(table, "stream_id")
	r.Time = field.NewTime(table, "time")
	r.TransactionIds = field.NewString(table, "transaction_ids")
	r.UpdatedTime = field.NewString(table, "updated_time")
	r.UserId = field.NewUint64(table, "user_id")

	r.fillFieldMap()

	return r
}

func (r *reOccuringTransactionORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reOccuringTransactionORM) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 24)
	r.fieldMap["account_id"] = r.AccountId
	r.fieldMap["average_amount"] = r.AverageAmount
	r.fieldMap["average_amount_iso_currency_code"] = r.AverageAmountIsoCurrencyCode
	r.fieldMap["category_id"] = r.CategoryId
	r.fieldMap["description"] = r.Description
	r.fieldMap["first_date"] = r.FirstDate
	r.fieldMap["flow"] = r.Flow
	r.fieldMap["frequency"] = r.Frequency
	r.fieldMap["id"] = r.Id
	r.fieldMap["is_active"] = r.IsActive
	r.fieldMap["last_amount"] = r.LastAmount
	r.fieldMap["last_amount_iso_currency_code"] = r.LastAmountIsoCurrencyCode
	r.fieldMap["last_date"] = r.LastDate
	r.fieldMap["link_id"] = r.LinkId
	r.fieldMap["merchant_name"] = r.MerchantName
	r.fieldMap["personal_finance_category_detailed"] = r.PersonalFinanceCategoryDetailed
	r.fieldMap["personal_finance_category_primary"] = r.PersonalFinanceCategoryPrimary
	r.fieldMap["sign"] = r.Sign
	r.fieldMap["status"] = r.Status
	r.fieldMap["stream_id"] = r.StreamId
	r.fieldMap["time"] = r.Time
	r.fieldMap["transaction_ids"] = r.TransactionIds
	r.fieldMap["updated_time"] = r.UpdatedTime
	r.fieldMap["user_id"] = r.UserId
}

func (r reOccuringTransactionORM) clone(db *gorm.DB) reOccuringTransactionORM {
	r.reOccuringTransactionORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reOccuringTransactionORM) replaceDB(db *gorm.DB) reOccuringTransactionORM {
	r.reOccuringTransactionORMDo.ReplaceDB(db)
	return r
}

type reOccuringTransactionORMDo struct{ gen.DO }

type IReOccuringTransactionORMDo interface {
	gen.SubQuery
	Debug() IReOccuringTransactionORMDo
	WithContext(ctx context.Context) IReOccuringTransactionORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReOccuringTransactionORMDo
	WriteDB() IReOccuringTransactionORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReOccuringTransactionORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReOccuringTransactionORMDo
	Not(conds ...gen.Condition) IReOccuringTransactionORMDo
	Or(conds ...gen.Condition) IReOccuringTransactionORMDo
	Select(conds ...field.Expr) IReOccuringTransactionORMDo
	Where(conds ...gen.Condition) IReOccuringTransactionORMDo
	Order(conds ...field.Expr) IReOccuringTransactionORMDo
	Distinct(cols ...field.Expr) IReOccuringTransactionORMDo
	Omit(cols ...field.Expr) IReOccuringTransactionORMDo
	Join(table schema.Tabler, on ...field.Expr) IReOccuringTransactionORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReOccuringTransactionORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReOccuringTransactionORMDo
	Group(cols ...field.Expr) IReOccuringTransactionORMDo
	Having(conds ...gen.Condition) IReOccuringTransactionORMDo
	Limit(limit int) IReOccuringTransactionORMDo
	Offset(offset int) IReOccuringTransactionORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReOccuringTransactionORMDo
	Unscoped() IReOccuringTransactionORMDo
	Create(values ...*financial_integration_service_apiv1.ReOccuringTransactionORM) error
	CreateInBatches(values []*financial_integration_service_apiv1.ReOccuringTransactionORM, batchSize int) error
	Save(values ...*financial_integration_service_apiv1.ReOccuringTransactionORM) error
	First() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error)
	Take() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error)
	Last() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error)
	Find() ([]*financial_integration_service_apiv1.ReOccuringTransactionORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.ReOccuringTransactionORM, err error)
	FindInBatches(result *[]*financial_integration_service_apiv1.ReOccuringTransactionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_integration_service_apiv1.ReOccuringTransactionORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReOccuringTransactionORMDo
	Assign(attrs ...field.AssignExpr) IReOccuringTransactionORMDo
	Joins(fields ...field.RelationField) IReOccuringTransactionORMDo
	Preload(fields ...field.RelationField) IReOccuringTransactionORMDo
	FirstOrInit() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error)
	FirstOrCreate() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error)
	FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.ReOccuringTransactionORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReOccuringTransactionORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result financial_integration_service_apiv1.ReOccuringTransactionORM, err error)
	GetByID(id int) (result financial_integration_service_apiv1.ReOccuringTransactionORM, err error)
	GetByIDs(ids []int) (result []financial_integration_service_apiv1.ReOccuringTransactionORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (r reOccuringTransactionORMDo) GetByUserID(user_id int) (result financial_integration_service_apiv1.ReOccuringTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM re_occuring_transactions ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (r reOccuringTransactionORMDo) GetByID(id int) (result financial_integration_service_apiv1.ReOccuringTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM re_occuring_transactions ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (r reOccuringTransactionORMDo) GetByIDs(ids []int) (result []financial_integration_service_apiv1.ReOccuringTransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM re_occuring_transactions ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (r reOccuringTransactionORMDo) Debug() IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Debug())
}

func (r reOccuringTransactionORMDo) WithContext(ctx context.Context) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reOccuringTransactionORMDo) ReadDB() IReOccuringTransactionORMDo {
	return r.Clauses(dbresolver.Read)
}

func (r reOccuringTransactionORMDo) WriteDB() IReOccuringTransactionORMDo {
	return r.Clauses(dbresolver.Write)
}

func (r reOccuringTransactionORMDo) Session(config *gorm.Session) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Session(config))
}

func (r reOccuringTransactionORMDo) Clauses(conds ...clause.Expression) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reOccuringTransactionORMDo) Returning(value interface{}, columns ...string) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reOccuringTransactionORMDo) Not(conds ...gen.Condition) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reOccuringTransactionORMDo) Or(conds ...gen.Condition) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reOccuringTransactionORMDo) Select(conds ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reOccuringTransactionORMDo) Where(conds ...gen.Condition) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reOccuringTransactionORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IReOccuringTransactionORMDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r reOccuringTransactionORMDo) Order(conds ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reOccuringTransactionORMDo) Distinct(cols ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reOccuringTransactionORMDo) Omit(cols ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reOccuringTransactionORMDo) Join(table schema.Tabler, on ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reOccuringTransactionORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reOccuringTransactionORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reOccuringTransactionORMDo) Group(cols ...field.Expr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reOccuringTransactionORMDo) Having(conds ...gen.Condition) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reOccuringTransactionORMDo) Limit(limit int) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reOccuringTransactionORMDo) Offset(offset int) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reOccuringTransactionORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reOccuringTransactionORMDo) Unscoped() IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reOccuringTransactionORMDo) Create(values ...*financial_integration_service_apiv1.ReOccuringTransactionORM) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reOccuringTransactionORMDo) CreateInBatches(values []*financial_integration_service_apiv1.ReOccuringTransactionORM, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reOccuringTransactionORMDo) Save(values ...*financial_integration_service_apiv1.ReOccuringTransactionORM) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reOccuringTransactionORMDo) First() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.ReOccuringTransactionORM), nil
	}
}

func (r reOccuringTransactionORMDo) Take() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.ReOccuringTransactionORM), nil
	}
}

func (r reOccuringTransactionORMDo) Last() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.ReOccuringTransactionORM), nil
	}
}

func (r reOccuringTransactionORMDo) Find() ([]*financial_integration_service_apiv1.ReOccuringTransactionORM, error) {
	result, err := r.DO.Find()
	return result.([]*financial_integration_service_apiv1.ReOccuringTransactionORM), err
}

func (r reOccuringTransactionORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.ReOccuringTransactionORM, err error) {
	buf := make([]*financial_integration_service_apiv1.ReOccuringTransactionORM, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reOccuringTransactionORMDo) FindInBatches(result *[]*financial_integration_service_apiv1.ReOccuringTransactionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reOccuringTransactionORMDo) Attrs(attrs ...field.AssignExpr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reOccuringTransactionORMDo) Assign(attrs ...field.AssignExpr) IReOccuringTransactionORMDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reOccuringTransactionORMDo) Joins(fields ...field.RelationField) IReOccuringTransactionORMDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reOccuringTransactionORMDo) Preload(fields ...field.RelationField) IReOccuringTransactionORMDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reOccuringTransactionORMDo) FirstOrInit() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.ReOccuringTransactionORM), nil
	}
}

func (r reOccuringTransactionORMDo) FirstOrCreate() (*financial_integration_service_apiv1.ReOccuringTransactionORM, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.ReOccuringTransactionORM), nil
	}
}

func (r reOccuringTransactionORMDo) FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.ReOccuringTransactionORM, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reOccuringTransactionORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reOccuringTransactionORMDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reOccuringTransactionORMDo) Delete(models ...*financial_integration_service_apiv1.ReOccuringTransactionORM) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reOccuringTransactionORMDo) withDO(do gen.Dao) *reOccuringTransactionORMDo {
	r.DO = *do.(*gen.DO)
	return r
}
