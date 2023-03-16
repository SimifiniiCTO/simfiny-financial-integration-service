// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&apiv1.SmartGoalORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&apiv1.SmartGoalORM{}) fail: %s", err)
	}
}

func Test_smartGoalORMQuery(t *testing.T) {
	smartGoalORM := newSmartGoalORM(db)
	smartGoalORM = *smartGoalORM.As(smartGoalORM.TableName())
	_do := smartGoalORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(smartGoalORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <smart_goals> fail:", err)
		return
	}

	_, ok := smartGoalORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from smartGoalORM success")
	}

	err = _do.Create(&apiv1.SmartGoalORM{})
	if err != nil {
		t.Error("create item in table <smart_goals> fail:", err)
	}

	err = _do.Save(&apiv1.SmartGoalORM{})
	if err != nil {
		t.Error("create item in table <smart_goals> fail:", err)
	}

	err = _do.CreateInBatches([]*apiv1.SmartGoalORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <smart_goals> fail:", err)
	}

	_, err = _do.Select(smartGoalORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <smart_goals> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <smart_goals> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <smart_goals> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <smart_goals> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*apiv1.SmartGoalORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <smart_goals> fail:", err)
	}

	_, err = _do.Select(smartGoalORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <smart_goals> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <smart_goals> fail:", err)
	}

	_, err = _do.Select(smartGoalORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <smart_goals> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <smart_goals> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <smart_goals> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <smart_goals> fail:", err)
	}

	_, err = _do.ScanByPage(&apiv1.SmartGoalORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <smart_goals> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <smart_goals> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <smart_goals> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <smart_goals> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <smart_goals> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <smart_goals> fail:", err)
	}
}

var SmartGoalORMGetByUserIDTestCase = []TestCase{}

func Test_smartGoalORM_GetByUserID(t *testing.T) {
	smartGoalORM := newSmartGoalORM(db)
	do := smartGoalORM.WithContext(context.Background()).Debug()

	for i, tt := range SmartGoalORMGetByUserIDTestCase {
		t.Run("GetByUserID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByUserID(tt.Input.Args[0].(int))
			assert(t, "GetByUserID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByUserID", res2, tt.Expectation.Ret[1])
		})
	}
}

var SmartGoalORMGetByIDTestCase = []TestCase{}

func Test_smartGoalORM_GetByID(t *testing.T) {
	smartGoalORM := newSmartGoalORM(db)
	do := smartGoalORM.WithContext(context.Background()).Debug()

	for i, tt := range SmartGoalORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(int))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var SmartGoalORMGetByIDsTestCase = []TestCase{}

func Test_smartGoalORM_GetByIDs(t *testing.T) {
	smartGoalORM := newSmartGoalORM(db)
	do := smartGoalORM.WithContext(context.Background()).Debug()

	for i, tt := range SmartGoalORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
