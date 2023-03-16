package database

import (
	"context"
	"fmt"

	"gorm.io/gen/field"
)

// DeleteUserProfileByUserID implements DatabaseOperations
func (db *Db) DeleteUserProfileByUserID(ctx context.Context, userID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-profile"); span != nil {
		defer span.End()
	}

	// validate the request
	if userID == 0 {
		return fmt.Errorf("invalid user_id")
	}

	// we check if a user profile with the same user_id already exists
	u := db.QueryOperator.UserProfileORM
	if _, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).First(); err != nil {
		return err
	}

	// user exists hence we delete the profile
	// NOTE: soft deletion is enabled by default for this
	// model hence deletion will be performed by setting the deleted_at column
	if _, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).Select(field.AssociationFields).Delete(); err != nil {
		return err
	}

	return nil
}
