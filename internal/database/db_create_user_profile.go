package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateUserProfile implements DatabaseOperations
func (db *Db) CreateUserProfile(ctx context.Context, profile *schema.UserProfile) (*schema.UserProfile, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-profile"); span != nil {
		defer span.End()
	}

	if profile == nil {
		return nil, fmt.Errorf("nil profile")
	}

	// validate the profile
	if err := profile.ValidateAll(); err != nil {
		return nil, err
	}

	record, err := profile.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// we check if a user profile with the same user_id already exists
	u := db.QueryOperator.UserProfileORM

	// SELECT * FROM users WHERE user_id = profile.user_id;
	if _, err = u.WithContext(ctx).Where(u.UserId.Eq(record.UserId)).First(); err != nil {
		if err == gorm.ErrRecordNotFound {
			// user does not exist hence we create the profile
			// NOTE: by default, create operations will be performed in a transaction
			if err := u.WithContext(ctx).Create(&record); err != nil {
				return nil, err
			}

			profile, err := record.ToPB(ctx)
			if err != nil {
				return nil, err
			}

			return &profile, nil
		}
	}

	// user already exists
	return nil, fmt.Errorf("user profile with user_id %d already exists", profile.UserId)
}
