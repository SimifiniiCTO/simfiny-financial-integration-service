package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
	"gorm.io/gen/field"
)

// DeleteGoal implements DatabaseOperations
func (db *Db) DeleteGoal(ctx context.Context, goalID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-goal"); span != nil {
		defer span.End()
	}

	if goalID == 0 {
		return fmt.Errorf("goal id must be provided. got: %d", goalID)
	}

	// ensure the goal exists
	g := db.QueryOperator.SmartGoalORM
	if _, err := g.WithContext(ctx).Where(g.Id.Eq(goalID)).First(); err != nil {
		return fmt.Errorf("goal with id %d does not exist", goalID)
	}

	result, err := g.WithContext(ctx).Where(g.Id.Eq(goalID)).Select(field.AssociationFields).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetGoal implements DatabaseOperations
func (db *Db) GetGoal(ctx context.Context, goalID uint64) (*schema.SmartGoal, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-goal"); span != nil {
		defer span.End()
	}

	if goalID == 0 {
		return nil, fmt.Errorf("goal id must be provided. got: %d", goalID)
	}

	// ensure the goal exists
	g := db.QueryOperator.SmartGoalORM
	goalOrm, err := g.
		WithContext(ctx).
		Where(g.Id.Eq(goalID)).
		Preload(g.Forecasts).
		Preload(g.Milestones.Budget.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("goal with id %d does not exist", goalID)
	}

	// convert orm to schema
	goal, err := goalOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &goal, nil
}

// GetGoalsByPocketID implements DatabaseOperations
func (db *Db) GetGoalsByPocketID(ctx context.Context, pocketID uint64) ([]*schema.SmartGoal, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-goals-by-pocket-id"); span != nil {
		defer span.End()
	}

	if pocketID == 0 {
		return nil, fmt.Errorf("pocket id must be provided. got: %d", pocketID)
	}

	// ensure the pocket exists
	p := db.QueryOperator.PocketORM
	pocketOrm, err := p.
		WithContext(ctx).
		Where(p.Id.Eq(pocketID)).
		Preload(p.Goals.Forecasts).
		Preload(p.Goals.Milestones.Budget.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("pocket with id %d does not exist", pocketID)
	}

	pocket, err := pocketOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	// return goals
	return pocket.Goals, nil
}

// UpdateGoal implements DatabaseOperations
func (db *Db) UpdateGoal(ctx context.Context, goal *schema.SmartGoal) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-goal"); span != nil {
		defer span.End()
	}

	if goal == nil {
		return fmt.Errorf("goal must be provided")
	}

	if goal.Id == 0 {
		return fmt.Errorf("goal id must be provided. got: %d", goal.Id)
	}

	// ensure the goal exists
	g := db.QueryOperator.SmartGoalORM
	_, err := g.WithContext(ctx).Where(g.Id.Eq(goal.Id)).First()
	if err != nil {
		return fmt.Errorf("goal with id %d does not exist", goal.Id)
	}

	// convert goal to orm type
	goalOrm, err := goal.ToORM(ctx)
	if err != nil {
		return err
	}

	// update goal
	result, err := g.WithContext(ctx).Where(g.Id.Eq(goal.Id)).Updates(goalOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// CreateGoal implements DatabaseOperations
func (db *Db) CreateGoal(ctx context.Context, pocketID uint64, goal *schema.SmartGoal) (*schema.SmartGoal, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-goal"); span != nil {
		defer span.End()
	}

	if goal == nil {
		return nil, fmt.Errorf("goal must be provided")
	}

	if pocketID == 0 {
		return nil, fmt.Errorf("pocket id must be provided. got: %d", pocketID)
	}

	// ensure the pocket exists
	p := db.QueryOperator.PocketORM
	pocket, err := p.WithContext(ctx).Where(p.Id.Eq(pocketID)).First()
	if err != nil {
		return nil, fmt.Errorf("pocket with id %d does not exist", pocketID)
	}

	// convert goal to orm type
	goalOrm, err := goal.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	if err := p.Goals.Model(pocket).Append(&goalOrm); err != nil {
		return nil, err
	}

	result, err := p.WithContext(ctx).Updates(pocket)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// convert orm to schema
	record, err := goalOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &record, nil
}
