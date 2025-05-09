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
	err := db.AutoMigrate(&financial_integration_service_apiv1.ReOccuringTransactionORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&financial_integration_service_apiv1.ReOccuringTransactionORM{}) fail: %s", err)
	}
}

func Test_reOccuringTransactionORMQuery(t *testing.T) {
	reOccuringTransactionORM := newReOccuringTransactionORM(db)
	reOccuringTransactionORM = *reOccuringTransactionORM.As(reOccuringTransactionORM.TableName())
	_do := reOccuringTransactionORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(reOccuringTransactionORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <re_occuring_transactions> fail:", err)
		return
	}

	_, ok := reOccuringTransactionORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from reOccuringTransactionORM success")
	}

	err = _do.Create(&financial_integration_service_apiv1.ReOccuringTransactionORM{})
	if err != nil {
		t.Error("create item in table <re_occuring_transactions> fail:", err)
	}

	err = _do.Save(&financial_integration_service_apiv1.ReOccuringTransactionORM{})
	if err != nil {
		t.Error("create item in table <re_occuring_transactions> fail:", err)
	}

	err = _do.CreateInBatches([]*financial_integration_service_apiv1.ReOccuringTransactionORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Select(reOccuringTransactionORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <re_occuring_transactions> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*financial_integration_service_apiv1.ReOccuringTransactionORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Select(reOccuringTransactionORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Select(reOccuringTransactionORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <re_occuring_transactions> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.ScanByPage(&financial_integration_service_apiv1.ReOccuringTransactionORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <re_occuring_transactions> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <re_occuring_transactions> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <re_occuring_transactions> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <re_occuring_transactions> fail:", err)
	}
}

var ReOccuringTransactionORMGetByUserIDTestCase = []TestCase{}

func Test_reOccuringTransactionORM_GetByUserID(t *testing.T) {
	reOccuringTransactionORM := newReOccuringTransactionORM(db)
	do := reOccuringTransactionORM.WithContext(context.Background()).Debug()

	for i, tt := range ReOccuringTransactionORMGetByUserIDTestCase {
		t.Run("GetByUserID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByUserID(tt.Input.Args[0].(int))
			assert(t, "GetByUserID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByUserID", res2, tt.Expectation.Ret[1])
		})
	}
}

var ReOccuringTransactionORMGetByIDTestCase = []TestCase{}

func Test_reOccuringTransactionORM_GetByID(t *testing.T) {
	reOccuringTransactionORM := newReOccuringTransactionORM(db)
	do := reOccuringTransactionORM.WithContext(context.Background()).Debug()

	for i, tt := range ReOccuringTransactionORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(int))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var ReOccuringTransactionORMGetByIDsTestCase = []TestCase{}

func Test_reOccuringTransactionORM_GetByIDs(t *testing.T) {
	reOccuringTransactionORM := newReOccuringTransactionORM(db)
	do := reOccuringTransactionORM.WithContext(context.Background()).Debug()

	for i, tt := range ReOccuringTransactionORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
