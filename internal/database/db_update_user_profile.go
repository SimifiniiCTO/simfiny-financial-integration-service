package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// UpdateUserProfile implements DatabaseOperations
func (db *Db) UpdateUserProfile(ctx context.Context, profile *schema.UserProfile) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-profile"); span != nil {
		defer span.End()
	}

	if profile == nil {
		return fmt.Errorf("nil profile")
	}

	// validate the profile
	if err := profile.ValidateAll(); err != nil {
		return err
	}

	record, err := profile.ToORM(ctx)
	if err != nil {
		return err
	}

	// we check if a user profile with the same user_id already exists
	u := db.QueryOperator.UserProfileORM

	// SELECT * FROM users WHERE user_id = profile.user_id;
	if _, err = u.WithContext(ctx).Where(u.Id.Eq(record.Id)).First(); err != nil {
		return err
	}

	// user exists hence we update the profile
	result, err := u.WithContext(ctx).Where(u.Id.Eq(record.Id)).Updates(&record)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected")
	}

	return nil
}
