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

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func newTransactionORM(db *gorm.DB, opts ...gen.DOOption) transactionORM {
	_transactionORM := transactionORM{}

	_transactionORM.transactionORMDo.UseDB(db, opts...)
	_transactionORM.transactionORMDo.UseModel(&apiv1.TransactionORM{})

	tableName := _transactionORM.transactionORMDo.TableName()
	_transactionORM.ALL = field.NewAsterisk(tableName)
	_transactionORM.AccountId = field.NewString(tableName, "account_id")
	_transactionORM.AccountOwner = field.NewString(tableName, "account_owner")
	_transactionORM.Amount = field.NewFloat64(tableName, "amount")
	_transactionORM.AuthorizedDate = field.NewString(tableName, "authorized_date")
	_transactionORM.AuthorizedDatetime = field.NewString(tableName, "authorized_datetime")
	_transactionORM.CategoryId = field.NewString(tableName, "category_id")
	_transactionORM.CheckNumber = field.NewString(tableName, "check_number")
	_transactionORM.Date = field.NewString(tableName, "date")
	_transactionORM.Datetime = field.NewString(tableName, "datetime")
	_transactionORM.Id = field.NewUint64(tableName, "id")
	_transactionORM.IsoCurrencyCode = field.NewString(tableName, "iso_currency_code")
	_transactionORM.LinkId = field.NewUint64(tableName, "link_id")
	_transactionORM.LocationAddress = field.NewString(tableName, "location_address")
	_transactionORM.LocationCity = field.NewString(tableName, "location_city")
	_transactionORM.LocationCountry = field.NewString(tableName, "location_country")
	_transactionORM.LocationLat = field.NewFloat64(tableName, "location_lat")
	_transactionORM.LocationLon = field.NewFloat64(tableName, "location_lon")
	_transactionORM.LocationPostalCode = field.NewString(tableName, "location_postal_code")
	_transactionORM.LocationRegion = field.NewString(tableName, "location_region")
	_transactionORM.LocationStoreNumber = field.NewString(tableName, "location_store_number")
	_transactionORM.MerchantName = field.NewString(tableName, "merchant_name")
	_transactionORM.Name = field.NewString(tableName, "name")
	_transactionORM.PaymentChannel = field.NewString(tableName, "payment_channel")
	_transactionORM.PaymentMetaByOrderOf = field.NewString(tableName, "payment_meta_by_order_of")
	_transactionORM.PaymentMetaPayee = field.NewString(tableName, "payment_meta_payee")
	_transactionORM.PaymentMetaPayer = field.NewString(tableName, "payment_meta_payer")
	_transactionORM.PaymentMetaPaymentMethod = field.NewString(tableName, "payment_meta_payment_method")
	_transactionORM.PaymentMetaPaymentProcessor = field.NewString(tableName, "payment_meta_payment_processor")
	_transactionORM.PaymentMetaPpdId = field.NewString(tableName, "payment_meta_ppd_id")
	_transactionORM.PaymentMetaReason = field.NewString(tableName, "payment_meta_reason")
	_transactionORM.PaymentMetaReferenceNumber = field.NewString(tableName, "payment_meta_reference_number")
	_transactionORM.Pending = field.NewBool(tableName, "pending")
	_transactionORM.PendingTransactionId = field.NewString(tableName, "pending_transaction_id")
	_transactionORM.PersonalFinanceCategoryDetailed = field.NewString(tableName, "personal_finance_category_detailed")
	_transactionORM.PersonalFinanceCategoryPrimary = field.NewString(tableName, "personal_finance_category_primary")
	_transactionORM.Sign = field.NewInt32(tableName, "sign")
	_transactionORM.TransactionCode = field.NewString(tableName, "transaction_code")
	_transactionORM.TransactionId = field.NewString(tableName, "transaction_id")
	_transactionORM.UnofficialCurrencyCode = field.NewString(tableName, "unofficial_currency_code")
	_transactionORM.UserId = field.NewUint64(tableName, "user_id")

	_transactionORM.fillFieldMap()

	return _transactionORM
}

type transactionORM struct {
	transactionORMDo

	ALL                             field.Asterisk
	AccountId                       field.String
	AccountOwner                    field.String
	Amount                          field.Float64
	AuthorizedDate                  field.String
	AuthorizedDatetime              field.String
	CategoryId                      field.String
	CheckNumber                     field.String
	Date                            field.String
	Datetime                        field.String
	Id                              field.Uint64
	IsoCurrencyCode                 field.String
	LinkId                          field.Uint64
	LocationAddress                 field.String
	LocationCity                    field.String
	LocationCountry                 field.String
	LocationLat                     field.Float64
	LocationLon                     field.Float64
	LocationPostalCode              field.String
	LocationRegion                  field.String
	LocationStoreNumber             field.String
	MerchantName                    field.String
	Name                            field.String
	PaymentChannel                  field.String
	PaymentMetaByOrderOf            field.String
	PaymentMetaPayee                field.String
	PaymentMetaPayer                field.String
	PaymentMetaPaymentMethod        field.String
	PaymentMetaPaymentProcessor     field.String
	PaymentMetaPpdId                field.String
	PaymentMetaReason               field.String
	PaymentMetaReferenceNumber      field.String
	Pending                         field.Bool
	PendingTransactionId            field.String
	PersonalFinanceCategoryDetailed field.String
	PersonalFinanceCategoryPrimary  field.String
	Sign                            field.Int32
	TransactionCode                 field.String
	TransactionId                   field.String
	UnofficialCurrencyCode          field.String
	UserId                          field.Uint64

	fieldMap map[string]field.Expr
}

func (t transactionORM) Table(newTableName string) *transactionORM {
	t.transactionORMDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t transactionORM) As(alias string) *transactionORM {
	t.transactionORMDo.DO = *(t.transactionORMDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *transactionORM) updateTableName(table string) *transactionORM {
	t.ALL = field.NewAsterisk(table)
	t.AccountId = field.NewString(table, "account_id")
	t.AccountOwner = field.NewString(table, "account_owner")
	t.Amount = field.NewFloat64(table, "amount")
	t.AuthorizedDate = field.NewString(table, "authorized_date")
	t.AuthorizedDatetime = field.NewString(table, "authorized_datetime")
	t.CategoryId = field.NewString(table, "category_id")
	t.CheckNumber = field.NewString(table, "check_number")
	t.Date = field.NewString(table, "date")
	t.Datetime = field.NewString(table, "datetime")
	t.Id = field.NewUint64(table, "id")
	t.IsoCurrencyCode = field.NewString(table, "iso_currency_code")
	t.LinkId = field.NewUint64(table, "link_id")
	t.LocationAddress = field.NewString(table, "location_address")
	t.LocationCity = field.NewString(table, "location_city")
	t.LocationCountry = field.NewString(table, "location_country")
	t.LocationLat = field.NewFloat64(table, "location_lat")
	t.LocationLon = field.NewFloat64(table, "location_lon")
	t.LocationPostalCode = field.NewString(table, "location_postal_code")
	t.LocationRegion = field.NewString(table, "location_region")
	t.LocationStoreNumber = field.NewString(table, "location_store_number")
	t.MerchantName = field.NewString(table, "merchant_name")
	t.Name = field.NewString(table, "name")
	t.PaymentChannel = field.NewString(table, "payment_channel")
	t.PaymentMetaByOrderOf = field.NewString(table, "payment_meta_by_order_of")
	t.PaymentMetaPayee = field.NewString(table, "payment_meta_payee")
	t.PaymentMetaPayer = field.NewString(table, "payment_meta_payer")
	t.PaymentMetaPaymentMethod = field.NewString(table, "payment_meta_payment_method")
	t.PaymentMetaPaymentProcessor = field.NewString(table, "payment_meta_payment_processor")
	t.PaymentMetaPpdId = field.NewString(table, "payment_meta_ppd_id")
	t.PaymentMetaReason = field.NewString(table, "payment_meta_reason")
	t.PaymentMetaReferenceNumber = field.NewString(table, "payment_meta_reference_number")
	t.Pending = field.NewBool(table, "pending")
	t.PendingTransactionId = field.NewString(table, "pending_transaction_id")
	t.PersonalFinanceCategoryDetailed = field.NewString(table, "personal_finance_category_detailed")
	t.PersonalFinanceCategoryPrimary = field.NewString(table, "personal_finance_category_primary")
	t.Sign = field.NewInt32(table, "sign")
	t.TransactionCode = field.NewString(table, "transaction_code")
	t.TransactionId = field.NewString(table, "transaction_id")
	t.UnofficialCurrencyCode = field.NewString(table, "unofficial_currency_code")
	t.UserId = field.NewUint64(table, "user_id")

	t.fillFieldMap()

	return t
}

func (t *transactionORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *transactionORM) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 40)
	t.fieldMap["account_id"] = t.AccountId
	t.fieldMap["account_owner"] = t.AccountOwner
	t.fieldMap["amount"] = t.Amount
	t.fieldMap["authorized_date"] = t.AuthorizedDate
	t.fieldMap["authorized_datetime"] = t.AuthorizedDatetime
	t.fieldMap["category_id"] = t.CategoryId
	t.fieldMap["check_number"] = t.CheckNumber
	t.fieldMap["date"] = t.Date
	t.fieldMap["datetime"] = t.Datetime
	t.fieldMap["id"] = t.Id
	t.fieldMap["iso_currency_code"] = t.IsoCurrencyCode
	t.fieldMap["link_id"] = t.LinkId
	t.fieldMap["location_address"] = t.LocationAddress
	t.fieldMap["location_city"] = t.LocationCity
	t.fieldMap["location_country"] = t.LocationCountry
	t.fieldMap["location_lat"] = t.LocationLat
	t.fieldMap["location_lon"] = t.LocationLon
	t.fieldMap["location_postal_code"] = t.LocationPostalCode
	t.fieldMap["location_region"] = t.LocationRegion
	t.fieldMap["location_store_number"] = t.LocationStoreNumber
	t.fieldMap["merchant_name"] = t.MerchantName
	t.fieldMap["name"] = t.Name
	t.fieldMap["payment_channel"] = t.PaymentChannel
	t.fieldMap["payment_meta_by_order_of"] = t.PaymentMetaByOrderOf
	t.fieldMap["payment_meta_payee"] = t.PaymentMetaPayee
	t.fieldMap["payment_meta_payer"] = t.PaymentMetaPayer
	t.fieldMap["payment_meta_payment_method"] = t.PaymentMetaPaymentMethod
	t.fieldMap["payment_meta_payment_processor"] = t.PaymentMetaPaymentProcessor
	t.fieldMap["payment_meta_ppd_id"] = t.PaymentMetaPpdId
	t.fieldMap["payment_meta_reason"] = t.PaymentMetaReason
	t.fieldMap["payment_meta_reference_number"] = t.PaymentMetaReferenceNumber
	t.fieldMap["pending"] = t.Pending
	t.fieldMap["pending_transaction_id"] = t.PendingTransactionId
	t.fieldMap["personal_finance_category_detailed"] = t.PersonalFinanceCategoryDetailed
	t.fieldMap["personal_finance_category_primary"] = t.PersonalFinanceCategoryPrimary
	t.fieldMap["sign"] = t.Sign
	t.fieldMap["transaction_code"] = t.TransactionCode
	t.fieldMap["transaction_id"] = t.TransactionId
	t.fieldMap["unofficial_currency_code"] = t.UnofficialCurrencyCode
	t.fieldMap["user_id"] = t.UserId
}

func (t transactionORM) clone(db *gorm.DB) transactionORM {
	t.transactionORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t transactionORM) replaceDB(db *gorm.DB) transactionORM {
	t.transactionORMDo.ReplaceDB(db)
	return t
}

type transactionORMDo struct{ gen.DO }

type ITransactionORMDo interface {
	gen.SubQuery
	Debug() ITransactionORMDo
	WithContext(ctx context.Context) ITransactionORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITransactionORMDo
	WriteDB() ITransactionORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITransactionORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITransactionORMDo
	Not(conds ...gen.Condition) ITransactionORMDo
	Or(conds ...gen.Condition) ITransactionORMDo
	Select(conds ...field.Expr) ITransactionORMDo
	Where(conds ...gen.Condition) ITransactionORMDo
	Order(conds ...field.Expr) ITransactionORMDo
	Distinct(cols ...field.Expr) ITransactionORMDo
	Omit(cols ...field.Expr) ITransactionORMDo
	Join(table schema.Tabler, on ...field.Expr) ITransactionORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITransactionORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITransactionORMDo
	Group(cols ...field.Expr) ITransactionORMDo
	Having(conds ...gen.Condition) ITransactionORMDo
	Limit(limit int) ITransactionORMDo
	Offset(offset int) ITransactionORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITransactionORMDo
	Unscoped() ITransactionORMDo
	Create(values ...*apiv1.TransactionORM) error
	CreateInBatches(values []*apiv1.TransactionORM, batchSize int) error
	Save(values ...*apiv1.TransactionORM) error
	First() (*apiv1.TransactionORM, error)
	Take() (*apiv1.TransactionORM, error)
	Last() (*apiv1.TransactionORM, error)
	Find() ([]*apiv1.TransactionORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.TransactionORM, err error)
	FindInBatches(result *[]*apiv1.TransactionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*apiv1.TransactionORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITransactionORMDo
	Assign(attrs ...field.AssignExpr) ITransactionORMDo
	Joins(fields ...field.RelationField) ITransactionORMDo
	Preload(fields ...field.RelationField) ITransactionORMDo
	FirstOrInit() (*apiv1.TransactionORM, error)
	FirstOrCreate() (*apiv1.TransactionORM, error)
	FindByPage(offset int, limit int) (result []*apiv1.TransactionORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITransactionORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result apiv1.TransactionORM, err error)
	GetByID(id int) (result apiv1.TransactionORM, err error)
	GetByIDs(ids []int) (result []apiv1.TransactionORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (t transactionORMDo) GetByUserID(user_id int) (result apiv1.TransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM transactions ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (t transactionORMDo) GetByID(id int) (result apiv1.TransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM transactions ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (t transactionORMDo) GetByIDs(ids []int) (result []apiv1.TransactionORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM transactions ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (t transactionORMDo) Debug() ITransactionORMDo {
	return t.withDO(t.DO.Debug())
}

func (t transactionORMDo) WithContext(ctx context.Context) ITransactionORMDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t transactionORMDo) ReadDB() ITransactionORMDo {
	return t.Clauses(dbresolver.Read)
}

func (t transactionORMDo) WriteDB() ITransactionORMDo {
	return t.Clauses(dbresolver.Write)
}

func (t transactionORMDo) Session(config *gorm.Session) ITransactionORMDo {
	return t.withDO(t.DO.Session(config))
}

func (t transactionORMDo) Clauses(conds ...clause.Expression) ITransactionORMDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t transactionORMDo) Returning(value interface{}, columns ...string) ITransactionORMDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t transactionORMDo) Not(conds ...gen.Condition) ITransactionORMDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t transactionORMDo) Or(conds ...gen.Condition) ITransactionORMDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t transactionORMDo) Select(conds ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t transactionORMDo) Where(conds ...gen.Condition) ITransactionORMDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t transactionORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITransactionORMDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t transactionORMDo) Order(conds ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t transactionORMDo) Distinct(cols ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t transactionORMDo) Omit(cols ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t transactionORMDo) Join(table schema.Tabler, on ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t transactionORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t transactionORMDo) RightJoin(table schema.Tabler, on ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t transactionORMDo) Group(cols ...field.Expr) ITransactionORMDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t transactionORMDo) Having(conds ...gen.Condition) ITransactionORMDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t transactionORMDo) Limit(limit int) ITransactionORMDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t transactionORMDo) Offset(offset int) ITransactionORMDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t transactionORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITransactionORMDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t transactionORMDo) Unscoped() ITransactionORMDo {
	return t.withDO(t.DO.Unscoped())
}

func (t transactionORMDo) Create(values ...*apiv1.TransactionORM) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t transactionORMDo) CreateInBatches(values []*apiv1.TransactionORM, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t transactionORMDo) Save(values ...*apiv1.TransactionORM) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t transactionORMDo) First() (*apiv1.TransactionORM, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.TransactionORM), nil
	}
}

func (t transactionORMDo) Take() (*apiv1.TransactionORM, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.TransactionORM), nil
	}
}

func (t transactionORMDo) Last() (*apiv1.TransactionORM, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.TransactionORM), nil
	}
}

func (t transactionORMDo) Find() ([]*apiv1.TransactionORM, error) {
	result, err := t.DO.Find()
	return result.([]*apiv1.TransactionORM), err
}

func (t transactionORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*apiv1.TransactionORM, err error) {
	buf := make([]*apiv1.TransactionORM, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t transactionORMDo) FindInBatches(result *[]*apiv1.TransactionORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t transactionORMDo) Attrs(attrs ...field.AssignExpr) ITransactionORMDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t transactionORMDo) Assign(attrs ...field.AssignExpr) ITransactionORMDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t transactionORMDo) Joins(fields ...field.RelationField) ITransactionORMDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t transactionORMDo) Preload(fields ...field.RelationField) ITransactionORMDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t transactionORMDo) FirstOrInit() (*apiv1.TransactionORM, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.TransactionORM), nil
	}
}

func (t transactionORMDo) FirstOrCreate() (*apiv1.TransactionORM, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*apiv1.TransactionORM), nil
	}
}

func (t transactionORMDo) FindByPage(offset int, limit int) (result []*apiv1.TransactionORM, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t transactionORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t transactionORMDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t transactionORMDo) Delete(models ...*apiv1.TransactionORM) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *transactionORMDo) withDO(do gen.Dao) *transactionORMDo {
	t.DO = *do.(*gen.DO)
	return t
}
