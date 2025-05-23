// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	financial_integration_service_apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&financial_integration_service_apiv1.MilestoneORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&financial_integration_service_apiv1.MilestoneORM{}) fail: %s", err)
	}
}

func Test_milestoneORMQuery(t *testing.T) {
	milestoneORM := newMilestoneORM(db)
	milestoneORM = *milestoneORM.As(milestoneORM.TableName())
	_do := milestoneORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(milestoneORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <milestones> fail:", err)
		return
	}

	_, ok := milestoneORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from milestoneORM success")
	}

	err = _do.Create(&financial_integration_service_apiv1.MilestoneORM{})
	if err != nil {
		t.Error("create item in table <milestones> fail:", err)
	}

	err = _do.Save(&financial_integration_service_apiv1.MilestoneORM{})
	if err != nil {
		t.Error("create item in table <milestones> fail:", err)
	}

	err = _do.CreateInBatches([]*financial_integration_service_apiv1.MilestoneORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <milestones> fail:", err)
	}

	_, err = _do.Select(milestoneORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <milestones> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <milestones> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <milestones> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <milestones> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*financial_integration_service_apiv1.MilestoneORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <milestones> fail:", err)
	}

	_, err = _do.Select(milestoneORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <milestones> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <milestones> fail:", err)
	}

	_, err = _do.Select(milestoneORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <milestones> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <milestones> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <milestones> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <milestones> fail:", err)
	}

	_, err = _do.ScanByPage(&financial_integration_service_apiv1.MilestoneORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <milestones> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <milestones> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <milestones> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <milestones> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <milestones> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <milestones> fail:", err)
	}
}

var MilestoneORMGetByUserIDTestCase = []TestCase{}

func Test_milestoneORM_GetByUserID(t *testing.T) {
	milestoneORM := newMilestoneORM(db)
	do := milestoneORM.WithContext(context.Background()).Debug()

	for i, tt := range MilestoneORMGetByUserIDTestCase {
		t.Run("GetByUserID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByUserID(tt.Input.Args[0].(int))
			assert(t, "GetByUserID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByUserID", res2, tt.Expectation.Ret[1])
		})
	}
}

var MilestoneORMGetByIDTestCase = []TestCase{}

func Test_milestoneORM_GetByID(t *testing.T) {
	milestoneORM := newMilestoneORM(db)
	do := milestoneORM.WithContext(context.Background()).Debug()

	for i, tt := range MilestoneORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(int))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var MilestoneORMGetByIDsTestCase = []TestCase{}

func Test_milestoneORM_GetByIDs(t *testing.T) {
	milestoneORM := newMilestoneORM(db)
	do := milestoneORM.WithContext(context.Background()).Debug()

	for i, tt := range MilestoneORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
