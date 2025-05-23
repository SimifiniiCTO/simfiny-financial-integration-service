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

func newBankAccountORM(db *gorm.DB, opts ...gen.DOOption) bankAccountORM {
	_bankAccountORM := bankAccountORM{}

	_bankAccountORM.bankAccountORMDo.UseDB(db, opts...)
	_bankAccountORM.bankAccountORMDo.UseModel(&financial_integration_service_apiv1.BankAccountORM{})

	tableName := _bankAccountORM.bankAccountORMDo.TableName()
	_bankAccountORM.ALL = field.NewAsterisk(tableName)
	_bankAccountORM.Balance = field.NewFloat32(tableName, "balance")
	_bankAccountORM.BalanceLimit = field.NewUint64(tableName, "balance_limit")
	_bankAccountORM.Currency = field.NewString(tableName, "currency")
	_bankAccountORM.CurrentFunds = field.NewFloat64(tableName, "current_funds")
	_bankAccountORM.Id = field.NewUint64(tableName, "id")
	_bankAccountORM.LinkId = field.NewUint64(tableName, "link_id")
	_bankAccountORM.Name = field.NewString(tableName, "name")
	_bankAccountORM.Number = field.NewString(tableName, "number")
	_bankAccountORM.PlaidAccountId = field.NewString(tableName, "plaid_account_id")
	_bankAccountORM.Status = field.NewString(tableName, "status")
	_bankAccountORM.Subtype = field.NewString(tableName, "subtype")
	_bankAccountORM.Type = field.NewString(tableName, "type")
	_bankAccountORM.UserId = field.NewUint64(tableName, "user_id")
	_bankAccountORM.Pockets = bankAccountORMHasManyPockets{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Pockets", "financial_integration_service_apiv1.PocketORM"),
		Goals: struct {
			field.RelationField
			Forecasts struct {
				field.RelationField
			}
			Milestones struct {
				field.RelationField
				Budget struct {
					field.RelationField
					Category struct {
						field.RelationField
					}
				}
			}
		}{
			RelationField: field.NewRelation("Pockets.Goals", "financial_integration_service_apiv1.SmartGoalORM"),
			Forecasts: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pockets.Goals.Forecasts", "financial_integration_service_apiv1.ForecastORM"),
			},
			Milestones: struct {
				field.RelationField
				Budget struct {
					field.RelationField
					Category struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("Pockets.Goals.Milestones", "financial_integration_service_apiv1.MilestoneORM"),
				Budget: struct {
					field.RelationField
					Category struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Pockets.Goals.Milestones.Budget", "financial_integration_service_apiv1.BudgetORM"),
					Category: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Pockets.Goals.Milestones.Budget.Category", "financial_integration_service_apiv1.CategoryORM"),
					},
				},
			},
		},
	}

	_bankAccountORM.fillFieldMap()

	return _bankAccountORM
}

type bankAccountORM struct {
	bankAccountORMDo

	ALL            field.Asterisk
	Balance        field.Float32
	BalanceLimit   field.Uint64
	Currency       field.String
	CurrentFunds   field.Float64
	Id             field.Uint64
	LinkId         field.Uint64
	Name           field.String
	Number         field.String
	PlaidAccountId field.String
	Status         field.String
	Subtype        field.String
	Type           field.String
	UserId         field.Uint64
	Pockets        bankAccountORMHasManyPockets

	fieldMap map[string]field.Expr
}

func (b bankAccountORM) Table(newTableName string) *bankAccountORM {
	b.bankAccountORMDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bankAccountORM) As(alias string) *bankAccountORM {
	b.bankAccountORMDo.DO = *(b.bankAccountORMDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bankAccountORM) updateTableName(table string) *bankAccountORM {
	b.ALL = field.NewAsterisk(table)
	b.Balance = field.NewFloat32(table, "balance")
	b.BalanceLimit = field.NewUint64(table, "balance_limit")
	b.Currency = field.NewString(table, "currency")
	b.CurrentFunds = field.NewFloat64(table, "current_funds")
	b.Id = field.NewUint64(table, "id")
	b.LinkId = field.NewUint64(table, "link_id")
	b.Name = field.NewString(table, "name")
	b.Number = field.NewString(table, "number")
	b.PlaidAccountId = field.NewString(table, "plaid_account_id")
	b.Status = field.NewString(table, "status")
	b.Subtype = field.NewString(table, "subtype")
	b.Type = field.NewString(table, "type")
	b.UserId = field.NewUint64(table, "user_id")

	b.fillFieldMap()

	return b
}

func (b *bankAccountORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bankAccountORM) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 14)
	b.fieldMap["balance"] = b.Balance
	b.fieldMap["balance_limit"] = b.BalanceLimit
	b.fieldMap["currency"] = b.Currency
	b.fieldMap["current_funds"] = b.CurrentFunds
	b.fieldMap["id"] = b.Id
	b.fieldMap["link_id"] = b.LinkId
	b.fieldMap["name"] = b.Name
	b.fieldMap["number"] = b.Number
	b.fieldMap["plaid_account_id"] = b.PlaidAccountId
	b.fieldMap["status"] = b.Status
	b.fieldMap["subtype"] = b.Subtype
	b.fieldMap["type"] = b.Type
	b.fieldMap["user_id"] = b.UserId

}

func (b bankAccountORM) clone(db *gorm.DB) bankAccountORM {
	b.bankAccountORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bankAccountORM) replaceDB(db *gorm.DB) bankAccountORM {
	b.bankAccountORMDo.ReplaceDB(db)
	return b
}

type bankAccountORMHasManyPockets struct {
	db *gorm.DB

	field.RelationField

	Goals struct {
		field.RelationField
		Forecasts struct {
			field.RelationField
		}
		Milestones struct {
			field.RelationField
			Budget struct {
				field.RelationField
				Category struct {
					field.RelationField
				}
			}
		}
	}
}

func (a bankAccountORMHasManyPockets) Where(conds ...field.Expr) *bankAccountORMHasManyPockets {
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

func (a bankAccountORMHasManyPockets) WithContext(ctx context.Context) *bankAccountORMHasManyPockets {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a bankAccountORMHasManyPockets) Model(m *financial_integration_service_apiv1.BankAccountORM) *bankAccountORMHasManyPocketsTx {
	return &bankAccountORMHasManyPocketsTx{a.db.Model(m).Association(a.Name())}
}

type bankAccountORMHasManyPocketsTx struct{ tx *gorm.Association }

func (a bankAccountORMHasManyPocketsTx) Find() (result []*financial_integration_service_apiv1.PocketORM, err error) {
	return result, a.tx.Find(&result)
}

func (a bankAccountORMHasManyPocketsTx) Append(values ...*financial_integration_service_apiv1.PocketORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a bankAccountORMHasManyPocketsTx) Replace(values ...*financial_integration_service_apiv1.PocketORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a bankAccountORMHasManyPocketsTx) Delete(values ...*financial_integration_service_apiv1.PocketORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a bankAccountORMHasManyPocketsTx) Clear() error {
	return a.tx.Clear()
}

func (a bankAccountORMHasManyPocketsTx) Count() int64 {
	return a.tx.Count()
}

type bankAccountORMDo struct{ gen.DO }

type IBankAccountORMDo interface {
	gen.SubQuery
	Debug() IBankAccountORMDo
	WithContext(ctx context.Context) IBankAccountORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBankAccountORMDo
	WriteDB() IBankAccountORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBankAccountORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBankAccountORMDo
	Not(conds ...gen.Condition) IBankAccountORMDo
	Or(conds ...gen.Condition) IBankAccountORMDo
	Select(conds ...field.Expr) IBankAccountORMDo
	Where(conds ...gen.Condition) IBankAccountORMDo
	Order(conds ...field.Expr) IBankAccountORMDo
	Distinct(cols ...field.Expr) IBankAccountORMDo
	Omit(cols ...field.Expr) IBankAccountORMDo
	Join(table schema.Tabler, on ...field.Expr) IBankAccountORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBankAccountORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBankAccountORMDo
	Group(cols ...field.Expr) IBankAccountORMDo
	Having(conds ...gen.Condition) IBankAccountORMDo
	Limit(limit int) IBankAccountORMDo
	Offset(offset int) IBankAccountORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBankAccountORMDo
	Unscoped() IBankAccountORMDo
	Create(values ...*financial_integration_service_apiv1.BankAccountORM) error
	CreateInBatches(values []*financial_integration_service_apiv1.BankAccountORM, batchSize int) error
	Save(values ...*financial_integration_service_apiv1.BankAccountORM) error
	First() (*financial_integration_service_apiv1.BankAccountORM, error)
	Take() (*financial_integration_service_apiv1.BankAccountORM, error)
	Last() (*financial_integration_service_apiv1.BankAccountORM, error)
	Find() ([]*financial_integration_service_apiv1.BankAccountORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.BankAccountORM, err error)
	FindInBatches(result *[]*financial_integration_service_apiv1.BankAccountORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_integration_service_apiv1.BankAccountORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBankAccountORMDo
	Assign(attrs ...field.AssignExpr) IBankAccountORMDo
	Joins(fields ...field.RelationField) IBankAccountORMDo
	Preload(fields ...field.RelationField) IBankAccountORMDo
	FirstOrInit() (*financial_integration_service_apiv1.BankAccountORM, error)
	FirstOrCreate() (*financial_integration_service_apiv1.BankAccountORM, error)
	FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.BankAccountORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBankAccountORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result financial_integration_service_apiv1.BankAccountORM, err error)
	GetByID(id int) (result financial_integration_service_apiv1.BankAccountORM, err error)
	GetByIDs(ids []int) (result []financial_integration_service_apiv1.BankAccountORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (b bankAccountORMDo) GetByUserID(user_id int) (result financial_integration_service_apiv1.BankAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM bank_accounts ")
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
func (b bankAccountORMDo) GetByID(id int) (result financial_integration_service_apiv1.BankAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM bank_accounts ")
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
func (b bankAccountORMDo) GetByIDs(ids []int) (result []financial_integration_service_apiv1.BankAccountORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM bank_accounts ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (b bankAccountORMDo) Debug() IBankAccountORMDo {
	return b.withDO(b.DO.Debug())
}

func (b bankAccountORMDo) WithContext(ctx context.Context) IBankAccountORMDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bankAccountORMDo) ReadDB() IBankAccountORMDo {
	return b.Clauses(dbresolver.Read)
}

func (b bankAccountORMDo) WriteDB() IBankAccountORMDo {
	return b.Clauses(dbresolver.Write)
}

func (b bankAccountORMDo) Session(config *gorm.Session) IBankAccountORMDo {
	return b.withDO(b.DO.Session(config))
}

func (b bankAccountORMDo) Clauses(conds ...clause.Expression) IBankAccountORMDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bankAccountORMDo) Returning(value interface{}, columns ...string) IBankAccountORMDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bankAccountORMDo) Not(conds ...gen.Condition) IBankAccountORMDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bankAccountORMDo) Or(conds ...gen.Condition) IBankAccountORMDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bankAccountORMDo) Select(conds ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bankAccountORMDo) Where(conds ...gen.Condition) IBankAccountORMDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bankAccountORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBankAccountORMDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b bankAccountORMDo) Order(conds ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bankAccountORMDo) Distinct(cols ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bankAccountORMDo) Omit(cols ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bankAccountORMDo) Join(table schema.Tabler, on ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bankAccountORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bankAccountORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bankAccountORMDo) Group(cols ...field.Expr) IBankAccountORMDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bankAccountORMDo) Having(conds ...gen.Condition) IBankAccountORMDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bankAccountORMDo) Limit(limit int) IBankAccountORMDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bankAccountORMDo) Offset(offset int) IBankAccountORMDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bankAccountORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBankAccountORMDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bankAccountORMDo) Unscoped() IBankAccountORMDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bankAccountORMDo) Create(values ...*financial_integration_service_apiv1.BankAccountORM) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bankAccountORMDo) CreateInBatches(values []*financial_integration_service_apiv1.BankAccountORM, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bankAccountORMDo) Save(values ...*financial_integration_service_apiv1.BankAccountORM) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bankAccountORMDo) First() (*financial_integration_service_apiv1.BankAccountORM, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BankAccountORM), nil
	}
}

func (b bankAccountORMDo) Take() (*financial_integration_service_apiv1.BankAccountORM, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BankAccountORM), nil
	}
}

func (b bankAccountORMDo) Last() (*financial_integration_service_apiv1.BankAccountORM, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BankAccountORM), nil
	}
}

func (b bankAccountORMDo) Find() ([]*financial_integration_service_apiv1.BankAccountORM, error) {
	result, err := b.DO.Find()
	return result.([]*financial_integration_service_apiv1.BankAccountORM), err
}

func (b bankAccountORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.BankAccountORM, err error) {
	buf := make([]*financial_integration_service_apiv1.BankAccountORM, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bankAccountORMDo) FindInBatches(result *[]*financial_integration_service_apiv1.BankAccountORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bankAccountORMDo) Attrs(attrs ...field.AssignExpr) IBankAccountORMDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bankAccountORMDo) Assign(attrs ...field.AssignExpr) IBankAccountORMDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bankAccountORMDo) Joins(fields ...field.RelationField) IBankAccountORMDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bankAccountORMDo) Preload(fields ...field.RelationField) IBankAccountORMDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bankAccountORMDo) FirstOrInit() (*financial_integration_service_apiv1.BankAccountORM, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BankAccountORM), nil
	}
}

func (b bankAccountORMDo) FirstOrCreate() (*financial_integration_service_apiv1.BankAccountORM, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.BankAccountORM), nil
	}
}

func (b bankAccountORMDo) FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.BankAccountORM, count int64, err error) {
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

func (b bankAccountORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bankAccountORMDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bankAccountORMDo) Delete(models ...*financial_integration_service_apiv1.BankAccountORM) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bankAccountORMDo) withDO(do gen.Dao) *bankAccountORMDo {
	b.DO = *do.(*gen.DO)
	return b
}
