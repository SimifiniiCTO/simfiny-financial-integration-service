package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// AddActionableInsight adds an actionable insight to the backend database
func (db *Db) AddActionableInsight(ctx context.Context, userId uint64, insight *schema.ActionableInsight) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-store-actionable-insight"); span != nil {
		defer span.End()
	}

	// validate the request
	if userId == 0 {
		return fmt.Errorf("invalid user_id")
	}

	// validate the actionable insight
	if err := insight.ValidateAll(); err != nil {
		return err
	}

	// convert the actionable insight to ORM
	record, err := insight.ToORM(ctx)
	if err != nil {
		return err
	}

	// query the user to check of its existence in the database
	u := db.QueryOperator.UserProfileORM
	userProfileRecord, err := u.
		WithContext(ctx).
		Where(u.UserId.Eq(userId)).
		First()
	if err != nil {
		return err
	}

	// associate the actionable insight to the user profile record
	// NOTE: a one to many relationship exists between the user profile record and the
	// actionable insight
	if err := u.ActionableInsights.Model(userProfileRecord).Append(&record); err != nil {
		return err
	}

	return nil
}

// GetActionableInsights obtains the set of actionable insights for the user from the database
func (db *Db) GetActionableInsights(ctx context.Context, userId uint64) ([]*schema.ActionableInsight, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-actionable-insights"); span != nil {
		defer span.End()
	}

	// validate the request
	if userId == 0 {
		return nil, fmt.Errorf("invalid user_id")
	}

	// we use the user profile to obtain the actionable insights sicne the actionable insights
	// are tied to the user profile
	u := db.QueryOperator.UserProfileORM
	record, err := u.
		WithContext(ctx).
		Where(u.UserId.Eq(userId)).
		Preload(u.ActionableInsights).First()
	if err != nil {
		return nil, err
	}

	if len(record.ActionableInsights) == 0 {
		return nil, fmt.Errorf("no actionable information available for user %d", userId)
	}

	insights := record.ActionableInsights
	results := make([]*schema.ActionableInsight, 0, len(insights))
	for _, r := range insights {
		profile, err := r.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		results = append(results, &profile)
	}

	return results, nil
}
