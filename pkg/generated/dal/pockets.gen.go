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

func newPocketORM(db *gorm.DB, opts ...gen.DOOption) pocketORM {
	_pocketORM := pocketORM{}

	_pocketORM.pocketORMDo.UseDB(db, opts...)
	_pocketORM.pocketORMDo.UseModel(&financial_integration_service_apiv1.PocketORM{})

	tableName := _pocketORM.pocketORMDo.TableName()
	_pocketORM.ALL = field.NewAsterisk(tableName)
	_pocketORM.BankAccountId = field.NewUint64(tableName, "bank_account_id")
	_pocketORM.Id = field.NewUint64(tableName, "id")
	_pocketORM.Type = field.NewString(tableName, "type")
	_pocketORM.Goals = pocketORMHasManyGoals{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Goals", "financial_integration_service_apiv1.SmartGoalORM"),
		Forecasts: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Goals.Forecasts", "financial_integration_service_apiv1.ForecastORM"),
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
			RelationField: field.NewRelation("Goals.Milestones", "financial_integration_service_apiv1.MilestoneORM"),
			Budget: struct {
				field.RelationField
				Category struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Goals.Milestones.Budget", "financial_integration_service_apiv1.BudgetORM"),
				Category: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Goals.Milestones.Budget.Category", "financial_integration_service_apiv1.CategoryORM"),
				},
			},
		},
	}

	_pocketORM.fillFieldMap()

	return _pocketORM
}

type pocketORM struct {
	pocketORMDo

	ALL           field.Asterisk
	BankAccountId field.Uint64
	Id            field.Uint64
	Type          field.String
	Goals         pocketORMHasManyGoals

	fieldMap map[string]field.Expr
}

func (p pocketORM) Table(newTableName string) *pocketORM {
	p.pocketORMDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pocketORM) As(alias string) *pocketORM {
	p.pocketORMDo.DO = *(p.pocketORMDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pocketORM) updateTableName(table string) *pocketORM {
	p.ALL = field.NewAsterisk(table)
	p.BankAccountId = field.NewUint64(table, "bank_account_id")
	p.Id = field.NewUint64(table, "id")
	p.Type = field.NewString(table, "type")

	p.fillFieldMap()

	return p
}

func (p *pocketORM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pocketORM) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["bank_account_id"] = p.BankAccountId
	p.fieldMap["id"] = p.Id
	p.fieldMap["type"] = p.Type

}

func (p pocketORM) clone(db *gorm.DB) pocketORM {
	p.pocketORMDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pocketORM) replaceDB(db *gorm.DB) pocketORM {
	p.pocketORMDo.ReplaceDB(db)
	return p
}

type pocketORMHasManyGoals struct {
	db *gorm.DB

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

func (a pocketORMHasManyGoals) Where(conds ...field.Expr) *pocketORMHasManyGoals {
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

func (a pocketORMHasManyGoals) WithContext(ctx context.Context) *pocketORMHasManyGoals {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a pocketORMHasManyGoals) Model(m *financial_integration_service_apiv1.PocketORM) *pocketORMHasManyGoalsTx {
	return &pocketORMHasManyGoalsTx{a.db.Model(m).Association(a.Name())}
}

type pocketORMHasManyGoalsTx struct{ tx *gorm.Association }

func (a pocketORMHasManyGoalsTx) Find() (result []*financial_integration_service_apiv1.SmartGoalORM, err error) {
	return result, a.tx.Find(&result)
}

func (a pocketORMHasManyGoalsTx) Append(values ...*financial_integration_service_apiv1.SmartGoalORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a pocketORMHasManyGoalsTx) Replace(values ...*financial_integration_service_apiv1.SmartGoalORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a pocketORMHasManyGoalsTx) Delete(values ...*financial_integration_service_apiv1.SmartGoalORM) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a pocketORMHasManyGoalsTx) Clear() error {
	return a.tx.Clear()
}

func (a pocketORMHasManyGoalsTx) Count() int64 {
	return a.tx.Count()
}

type pocketORMDo struct{ gen.DO }

type IPocketORMDo interface {
	gen.SubQuery
	Debug() IPocketORMDo
	WithContext(ctx context.Context) IPocketORMDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPocketORMDo
	WriteDB() IPocketORMDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPocketORMDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPocketORMDo
	Not(conds ...gen.Condition) IPocketORMDo
	Or(conds ...gen.Condition) IPocketORMDo
	Select(conds ...field.Expr) IPocketORMDo
	Where(conds ...gen.Condition) IPocketORMDo
	Order(conds ...field.Expr) IPocketORMDo
	Distinct(cols ...field.Expr) IPocketORMDo
	Omit(cols ...field.Expr) IPocketORMDo
	Join(table schema.Tabler, on ...field.Expr) IPocketORMDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPocketORMDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPocketORMDo
	Group(cols ...field.Expr) IPocketORMDo
	Having(conds ...gen.Condition) IPocketORMDo
	Limit(limit int) IPocketORMDo
	Offset(offset int) IPocketORMDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPocketORMDo
	Unscoped() IPocketORMDo
	Create(values ...*financial_integration_service_apiv1.PocketORM) error
	CreateInBatches(values []*financial_integration_service_apiv1.PocketORM, batchSize int) error
	Save(values ...*financial_integration_service_apiv1.PocketORM) error
	First() (*financial_integration_service_apiv1.PocketORM, error)
	Take() (*financial_integration_service_apiv1.PocketORM, error)
	Last() (*financial_integration_service_apiv1.PocketORM, error)
	Find() ([]*financial_integration_service_apiv1.PocketORM, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.PocketORM, err error)
	FindInBatches(result *[]*financial_integration_service_apiv1.PocketORM, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*financial_integration_service_apiv1.PocketORM) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPocketORMDo
	Assign(attrs ...field.AssignExpr) IPocketORMDo
	Joins(fields ...field.RelationField) IPocketORMDo
	Preload(fields ...field.RelationField) IPocketORMDo
	FirstOrInit() (*financial_integration_service_apiv1.PocketORM, error)
	FirstOrCreate() (*financial_integration_service_apiv1.PocketORM, error)
	FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.PocketORM, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPocketORMDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByUserID(user_id int) (result financial_integration_service_apiv1.PocketORM, err error)
	GetByID(id int) (result financial_integration_service_apiv1.PocketORM, err error)
	GetByIDs(ids []int) (result []financial_integration_service_apiv1.PocketORM, err error)
}

// SELECT * FROM @@table
// {{where}}
//
//	user_id=@user_id
//
// {{end}}
func (p pocketORMDo) GetByUserID(user_id int) (result financial_integration_service_apiv1.PocketORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM pockets ")
	var whereSQL0 strings.Builder
	params = append(params, user_id)
	whereSQL0.WriteString("user_id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id=@id
//
// {{end}}
func (p pocketORMDo) GetByID(id int) (result financial_integration_service_apiv1.PocketORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM pockets ")
	var whereSQL0 strings.Builder
	params = append(params, id)
	whereSQL0.WriteString("id=? ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// {{where}}
//
//	id IN (@ids)
//
// {{end}}
func (p pocketORMDo) GetByIDs(ids []int) (result []financial_integration_service_apiv1.PocketORM, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM pockets ")
	var whereSQL0 strings.Builder
	params = append(params, ids)
	whereSQL0.WriteString("id IN (?) ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p pocketORMDo) Debug() IPocketORMDo {
	return p.withDO(p.DO.Debug())
}

func (p pocketORMDo) WithContext(ctx context.Context) IPocketORMDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pocketORMDo) ReadDB() IPocketORMDo {
	return p.Clauses(dbresolver.Read)
}

func (p pocketORMDo) WriteDB() IPocketORMDo {
	return p.Clauses(dbresolver.Write)
}

func (p pocketORMDo) Session(config *gorm.Session) IPocketORMDo {
	return p.withDO(p.DO.Session(config))
}

func (p pocketORMDo) Clauses(conds ...clause.Expression) IPocketORMDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pocketORMDo) Returning(value interface{}, columns ...string) IPocketORMDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pocketORMDo) Not(conds ...gen.Condition) IPocketORMDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pocketORMDo) Or(conds ...gen.Condition) IPocketORMDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pocketORMDo) Select(conds ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pocketORMDo) Where(conds ...gen.Condition) IPocketORMDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pocketORMDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPocketORMDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pocketORMDo) Order(conds ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pocketORMDo) Distinct(cols ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pocketORMDo) Omit(cols ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pocketORMDo) Join(table schema.Tabler, on ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pocketORMDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pocketORMDo) RightJoin(table schema.Tabler, on ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pocketORMDo) Group(cols ...field.Expr) IPocketORMDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pocketORMDo) Having(conds ...gen.Condition) IPocketORMDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pocketORMDo) Limit(limit int) IPocketORMDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pocketORMDo) Offset(offset int) IPocketORMDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pocketORMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPocketORMDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pocketORMDo) Unscoped() IPocketORMDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pocketORMDo) Create(values ...*financial_integration_service_apiv1.PocketORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pocketORMDo) CreateInBatches(values []*financial_integration_service_apiv1.PocketORM, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pocketORMDo) Save(values ...*financial_integration_service_apiv1.PocketORM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pocketORMDo) First() (*financial_integration_service_apiv1.PocketORM, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.PocketORM), nil
	}
}

func (p pocketORMDo) Take() (*financial_integration_service_apiv1.PocketORM, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.PocketORM), nil
	}
}

func (p pocketORMDo) Last() (*financial_integration_service_apiv1.PocketORM, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.PocketORM), nil
	}
}

func (p pocketORMDo) Find() ([]*financial_integration_service_apiv1.PocketORM, error) {
	result, err := p.DO.Find()
	return result.([]*financial_integration_service_apiv1.PocketORM), err
}

func (p pocketORMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*financial_integration_service_apiv1.PocketORM, err error) {
	buf := make([]*financial_integration_service_apiv1.PocketORM, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pocketORMDo) FindInBatches(result *[]*financial_integration_service_apiv1.PocketORM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pocketORMDo) Attrs(attrs ...field.AssignExpr) IPocketORMDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pocketORMDo) Assign(attrs ...field.AssignExpr) IPocketORMDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pocketORMDo) Joins(fields ...field.RelationField) IPocketORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pocketORMDo) Preload(fields ...field.RelationField) IPocketORMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pocketORMDo) FirstOrInit() (*financial_integration_service_apiv1.PocketORM, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.PocketORM), nil
	}
}

func (p pocketORMDo) FirstOrCreate() (*financial_integration_service_apiv1.PocketORM, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*financial_integration_service_apiv1.PocketORM), nil
	}
}

func (p pocketORMDo) FindByPage(offset int, limit int) (result []*financial_integration_service_apiv1.PocketORM, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pocketORMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pocketORMDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pocketORMDo) Delete(models ...*financial_integration_service_apiv1.PocketORM) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pocketORMDo) withDO(do gen.Dao) *pocketORMDo {
	p.DO = *do.(*gen.DO)
	return p
}
