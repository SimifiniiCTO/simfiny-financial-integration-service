package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"gorm.io/gen/field"
)

// UpdateMilestone implements DatabaseOperations
func (db *Db) UpdateMilestone(ctx context.Context, milestone *schema.Milestone) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-milestone"); span != nil {
		defer span.End()
	}

	if milestone == nil {
		return fmt.Errorf("milestone must be provided")
	}

	if milestone.Id == 0 {
		return fmt.Errorf("milestone id must be provided. got: %d", milestone.Id)
	}

	// ensure the milestone exists
	m := db.QueryOperator.MilestoneORM
	if _, err := m.WithContext(ctx).Where(m.Id.Eq(milestone.Id)).First(); err != nil {
		return fmt.Errorf("milestone with id %d does not exist", milestone.Id)
	}

	// convert schema to orm
	milestoneOrm, err := milestone.ToORM(ctx)
	if err != nil {
		return err
	}

	// perform the update operation
	result, err := m.WithContext(ctx).Where(m.Id.Eq(milestone.Id)).Updates(milestoneOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetMilestone implements DatabaseOperations
func (db *Db) GetMilestone(ctx context.Context, milestoneID uint64) (*schema.Milestone, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-milestone"); span != nil {
		defer span.End()
	}

	if milestoneID == 0 {
		return nil, fmt.Errorf("milestone id must be provided. got: %d", milestoneID)
	}

	// ensure the milestone exists
	m := db.QueryOperator.MilestoneORM
	milestoneOrm, err := m.
		WithContext(ctx).
		Where(m.Id.Eq(milestoneID)).
		Preload(m.Budget.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("milestone with id %d does not exist", milestoneID)
	}

	// convert orm to schema
	milestone, err := milestoneOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &milestone, nil
}

// GetMilestonesByGoalID implements DatabaseOperations
func (db *Db) GetMilestonesByGoalID(ctx context.Context, goalID uint64) ([]*schema.Milestone, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-milestones-by-goal-id"); span != nil {
		defer span.End()
	}

	if goalID == 0 {
		return nil, fmt.Errorf("goal id must be provided. got: %d", goalID)
	}

	// ensure the goal exists
	g := db.QueryOperator.SmartGoalORM
	smartGoal, err := g.
		WithContext(ctx).
		Where(g.Id.Eq(goalID)).
		Preload(g.Milestones.Budget.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("goal with id %d does not exist", goalID)
	}

	goals, err := smartGoal.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return goals.Milestones, nil
}

// DeleteMilestone implements DatabaseOperations
func (db *Db) DeleteMilestone(ctx context.Context, milestoneID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-milestone"); span != nil {
		defer span.End()
	}

	if milestoneID == 0 {
		return fmt.Errorf("milestone id must be provided. got: %d", milestoneID)
	}

	// ensure the milestone exists
	m := db.QueryOperator.MilestoneORM
	if _, err := m.WithContext(ctx).Where(m.Id.Eq(milestoneID)).First(); err != nil {
		return fmt.Errorf("milestone with id %d does not exist", milestoneID)
	}

	// perform the delete operation
	result, err := m.WithContext(ctx).Where(m.Id.Eq(milestoneID)).Select(field.AssociationFields).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// CreateMilestone implements DatabaseOperations
func (db *Db) CreateMilestone(ctx context.Context, goalID uint64, milestone *schema.Milestone) (*schema.Milestone, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-milestone"); span != nil {
		defer span.End()
	}

	if goalID == 0 {
		return nil, fmt.Errorf("goal id must be provided. got: %d", goalID)
	}

	// ensure the goal exists
	g := db.QueryOperator.SmartGoalORM
	smartGoal, err := g.WithContext(ctx).Where(g.Id.Eq(goalID)).First()
	if err != nil {
		return nil, fmt.Errorf("goal with id %d does not exist", goalID)
	}

	// create milestone orm
	milestoneOrm, err := milestone.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	if err := g.Milestones.Model(smartGoal).Append(&milestoneOrm); err != nil {
		return nil, err
	}

	// convert orm to schema
	createdMilestone, err := milestoneOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &createdMilestone, nil
}
