// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&apiv1.InvesmentHoldingORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&apiv1.InvesmentHoldingORM{}) fail: %s", err)
	}
}

func Test_invesmentHoldingORMQuery(t *testing.T) {
	invesmentHoldingORM := newInvesmentHoldingORM(db)
	invesmentHoldingORM = *invesmentHoldingORM.As(invesmentHoldingORM.TableName())
	_do := invesmentHoldingORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(invesmentHoldingORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <invesment_holdings> fail:", err)
		return
	}

	_, ok := invesmentHoldingORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from invesmentHoldingORM success")
	}

	err = _do.Create(&apiv1.InvesmentHoldingORM{})
	if err != nil {
		t.Error("create item in table <invesment_holdings> fail:", err)
	}

	err = _do.Save(&apiv1.InvesmentHoldingORM{})
	if err != nil {
		t.Error("create item in table <invesment_holdings> fail:", err)
	}

	err = _do.CreateInBatches([]*apiv1.InvesmentHoldingORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <invesment_holdings> fail:", err)
	}

	_, err = _do.Select(invesmentHoldingORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <invesment_holdings> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*apiv1.InvesmentHoldingORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Select(invesmentHoldingORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Select(invesmentHoldingORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <invesment_holdings> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.ScanByPage(&apiv1.InvesmentHoldingORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <invesment_holdings> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <invesment_holdings> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <invesment_holdings> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <invesment_holdings> fail:", err)
	}
}

var InvesmentHoldingORMGetByUserIDTestCase = []TestCase{}

func Test_invesmentHoldingORM_GetByUserID(t *testing.T) {
	invesmentHoldingORM := newInvesmentHoldingORM(db)
	do := invesmentHoldingORM.WithContext(context.Background()).Debug()

	for i, tt := range InvesmentHoldingORMGetByUserIDTestCase {
		t.Run("GetByUserID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByUserID(tt.Input.Args[0].(int))
			assert(t, "GetByUserID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByUserID", res2, tt.Expectation.Ret[1])
		})
	}
}

var InvesmentHoldingORMGetByIDTestCase = []TestCase{}

func Test_invesmentHoldingORM_GetByID(t *testing.T) {
	invesmentHoldingORM := newInvesmentHoldingORM(db)
	do := invesmentHoldingORM.WithContext(context.Background()).Debug()

	for i, tt := range InvesmentHoldingORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(int))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var InvesmentHoldingORMGetByIDsTestCase = []TestCase{}

func Test_invesmentHoldingORM_GetByIDs(t *testing.T) {
	invesmentHoldingORM := newInvesmentHoldingORM(db)
	do := invesmentHoldingORM.WithContext(context.Background()).Debug()

	for i, tt := range InvesmentHoldingORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
