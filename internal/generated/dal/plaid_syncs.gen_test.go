// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&apiv1.PlaidSyncORM{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&apiv1.PlaidSyncORM{}) fail: %s", err)
	}
}

func Test_plaidSyncORMQuery(t *testing.T) {
	plaidSyncORM := newPlaidSyncORM(db)
	plaidSyncORM = *plaidSyncORM.As(plaidSyncORM.TableName())
	_do := plaidSyncORM.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(plaidSyncORM.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <plaid_syncs> fail:", err)
		return
	}

	_, ok := plaidSyncORM.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from plaidSyncORM success")
	}

	err = _do.Create(&apiv1.PlaidSyncORM{})
	if err != nil {
		t.Error("create item in table <plaid_syncs> fail:", err)
	}

	err = _do.Save(&apiv1.PlaidSyncORM{})
	if err != nil {
		t.Error("create item in table <plaid_syncs> fail:", err)
	}

	err = _do.CreateInBatches([]*apiv1.PlaidSyncORM{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <plaid_syncs> fail:", err)
	}

	_, err = _do.Select(plaidSyncORM.ALL).Take()
	if err != nil {
		t.Error("Take() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <plaid_syncs> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*apiv1.PlaidSyncORM{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Select(plaidSyncORM.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Select(plaidSyncORM.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <plaid_syncs> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.ScanByPage(&apiv1.PlaidSyncORM{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <plaid_syncs> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <plaid_syncs> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <plaid_syncs> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <plaid_syncs> fail:", err)
	}
}

var PlaidSyncORMGetByUserIDTestCase = []TestCase{}

func Test_plaidSyncORM_GetByUserID(t *testing.T) {
	plaidSyncORM := newPlaidSyncORM(db)
	do := plaidSyncORM.WithContext(context.Background()).Debug()

	for i, tt := range PlaidSyncORMGetByUserIDTestCase {
		t.Run("GetByUserID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByUserID(tt.Input.Args[0].(int))
			assert(t, "GetByUserID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByUserID", res2, tt.Expectation.Ret[1])
		})
	}
}

var PlaidSyncORMGetByIDTestCase = []TestCase{}

func Test_plaidSyncORM_GetByID(t *testing.T) {
	plaidSyncORM := newPlaidSyncORM(db)
	do := plaidSyncORM.WithContext(context.Background()).Debug()

	for i, tt := range PlaidSyncORMGetByIDTestCase {
		t.Run("GetByID_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByID(tt.Input.Args[0].(int))
			assert(t, "GetByID", res1, tt.Expectation.Ret[0])
			assert(t, "GetByID", res2, tt.Expectation.Ret[1])
		})
	}
}

var PlaidSyncORMGetByIDsTestCase = []TestCase{}

func Test_plaidSyncORM_GetByIDs(t *testing.T) {
	plaidSyncORM := newPlaidSyncORM(db)
	do := plaidSyncORM.WithContext(context.Background()).Debug()

	for i, tt := range PlaidSyncORMGetByIDsTestCase {
		t.Run("GetByIDs_"+strconv.Itoa(i), func(t *testing.T) {
			res1, res2 := do.GetByIDs(tt.Input.Args[0].([]int))
			assert(t, "GetByIDs", res1, tt.Expectation.Ret[0])
			assert(t, "GetByIDs", res2, tt.Expectation.Ret[1])
		})
	}
}
