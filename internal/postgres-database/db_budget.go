package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"gorm.io/gen/field"
)

// DeleteBudget implements DatabaseOperations
func (db *Db) DeleteBudget(ctx context.Context, budgetID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-budget"); span != nil {
		defer span.End()
	}

	if budgetID == 0 {
		return fmt.Errorf("budget id must be provided. got: %d", budgetID)
	}

	// ensure the budget exists
	b := db.QueryOperator.BudgetORM
	if _, err := b.WithContext(ctx).Where(b.Id.Eq(budgetID)).First(); err != nil {
		return fmt.Errorf("budget with id %d does not exist", budgetID)
	}

	// perform the delete operation
	result, err := b.WithContext(ctx).Where(b.Id.Eq(budgetID)).Select(field.AssociationFields).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetBudget implements DatabaseOperations
func (db *Db) GetBudget(ctx context.Context, budgetID uint64) (*schema.Budget, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-budget"); span != nil {
		defer span.End()
	}

	if budgetID == 0 {
		return nil, fmt.Errorf("budget id must be provided. got: %d", budgetID)
	}

	// ensure the budget exists
	b := db.QueryOperator.BudgetORM
	budgetOrm, err := b.
		WithContext(ctx).
		Where(b.Id.Eq(budgetID)).
		Preload(b.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("budget with id %d does not exist", budgetID)
	}

	// convert orm to schema
	budget, err := budgetOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &budget, nil
}

// UpdateBudget implements DatabaseOperations
func (db *Db) UpdateBudget(ctx context.Context, budget *schema.Budget) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-budget"); span != nil {
		defer span.End()
	}

	if budget == nil {
		return fmt.Errorf("nil budget")
	}

	// validate the budget
	if err := budget.ValidateAll(); err != nil {
		return err
	}

	// ensure the budget exists
	b := db.QueryOperator.BudgetORM
	if _, err := b.WithContext(ctx).Where(b.Id.Eq(budget.Id)).First(); err != nil {
		return fmt.Errorf("budget with id %d does not exist", budget.Id)
	}

	// convert schema to orm
	updatedBudgetORM, err := budget.ToORM(ctx)
	if err != nil {
		return err
	}

	// update the budget
	result, err := b.WithContext(ctx).Updates(updatedBudgetORM)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (db *Db) CreateBudget(ctx context.Context, milestoneId uint64, budget *schema.Budget) (*schema.Budget, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-budget"); span != nil {
		defer span.End()
	}

	if budget == nil || milestoneId == 0 {
		return nil, fmt.Errorf("budget, and milestone id must be provided. got: %v, %d", budget, milestoneId)
	}

	// validate the budget
	if err := budget.ValidateAll(); err != nil {
		return nil, err
	}

	// ensure the milestone exists
	m := db.QueryOperator.MilestoneORM
	milestone, err := m.WithContext(ctx).Where(m.Id.Eq(milestoneId)).First()
	if err != nil {
		return nil, fmt.Errorf("milestone with id %d does not exist", milestoneId)
	}

	// convert budget to orm
	budgetOrm, err := budget.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// append the new budget to the milestone
	if err := m.Budget.Model(milestone).Append(&budgetOrm); err != nil {
		return nil, err
	}

	// perform the update operation against the milestone
	result, err := m.WithContext(ctx).Updates(milestone)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// budgetOrm reference at this point is auto updated with the db record id
	// as a byproduct of the update operation
	res, err := budgetOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
