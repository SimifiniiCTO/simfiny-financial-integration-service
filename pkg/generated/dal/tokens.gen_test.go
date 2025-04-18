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
	err := db.AutoMigrate(&financial_integration_service_apiv1.TokenORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&financial_integration_service_apiv1.TokenORM{}) fail: %s", err)
	}
}

func Test_tokenORMQuery(t *testing.T) {
	tokenORM := newTokenORM(db)
	tokenORM = *tokenORM.As(tokenORM.TableName())
	_do := tokenORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(tokenORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <tokens> fail:", err)
		return
	}

	_, ok := tokenORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from tokenORM success")
	}

	err = _do.Create(&financial_integration_service_apiv1.TokenORM{})
	if err != nil {
		t.Error("create item in table <tokens> fail:", err)
	}

	err = _do.Save(&financial_integration_service_apiv1.TokenORM{})
	if err != nil {
		t.Error("create item in table <tokens> fail:", err)
	}

	err = _do.CreateInBatches([]*financial_integration_service_apiv1.TokenORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <tokens> fail:", err)
	}

	_, err = _do.Select(tokenORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <tokens> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <tokens> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <tokens> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <tokens> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*financial_integration_service_apiv1.TokenORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <tokens> fail:", err)
	}

	_, err = _do.Select(tokenORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <tokens> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <tokens> fail:", err)
	}

	_, err = _do.Select(tokenORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <tokens> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <tokens> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <tokens> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <tokens> fail:", err)
	}

	_, err = _do.ScanByPage(&financial_integration_service_apiv1.TokenORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <tokens> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <tokens> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <tokens> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <tokens> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <tokens> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <tokens> fail:", err)
	}
}

var TokenORMGetByUserIDTestCase = []TestCase{}

func Test_tokenORM_GetByUserID(t *testing.T) {
	tokenORM := newTokenORM(db)
	do := tokenORM.WithContext(context.Background()).Debug()

	for i, tt := range TokenORMGetByUserIDTestCase {
		t.Run("GetByUserID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByUserID(tt.Input.Args[0].(int))
			assert(t, "GetByUserID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByUserID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TokenORMGetByIDTestCase = []TestCase{}

func Test_tokenORM_GetByID(t *testing.T) {
	tokenORM := newTokenORM(db)
	do := tokenORM.WithContext(context.Background()).Debug()

	for i, tt := range TokenORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(int))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var TokenORMGetByIDsTestCase = []TestCase{}

func Test_tokenORM_GetByIDs(t *testing.T) {
	tokenORM := newTokenORM(db)
	do := tokenORM.WithContext(context.Background()).Debug()

	for i, tt := range TokenORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
