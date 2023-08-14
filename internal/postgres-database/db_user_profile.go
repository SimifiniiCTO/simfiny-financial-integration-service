package database

import (
	"context"
	"fmt"

	"gorm.io/gen/field"
	"gorm.io/gorm"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
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

func (db *Db) UpdateUserProfileSubscription(ctx context.Context, userId uint64, subscription *schema.StripeSubscription) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-profile"); span != nil {
		defer span.End()
	}

	// validate the request
	if userId == 0 {
		return fmt.Errorf("invalid user_id")
	}

	if subscription == nil {
		return fmt.Errorf("nil subscription")
	}

	// validate stripe subscription
	if err := subscription.ValidateAll(); err != nil {
		return err
	}

	// convert the subscription to orm
	subscriptionRecord, err := subscription.ToORM(ctx)
	if err != nil {
		return err
	}

	u := db.QueryOperator.UserProfileORM
	record, err := u.
		WithContext(ctx).
		Where(u.UserId.Eq(userId)).
		First()
	if err != nil {
		return err
	}

	if err := u.StripeSubscriptions.Model(record).Replace(&subscriptionRecord); err != nil {
		return err
	}

	return nil
}

// GetUserProfileByUserID obtains a user profile by user_id if it exists within the database
func (db *Db) GetUserProfileByUserID(ctx context.Context, userID uint64) (*schema.UserProfile, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-profile"); span != nil {
		defer span.End()
	}

	// validate the request
	if userID == 0 {
		return nil, fmt.Errorf("invalid user_id")
	}

	// preload all dependencies
	// meaning for all connected accounts, obtain all accounts and subaccounts
	u := db.QueryOperator.UserProfileORM
	record, err := u.
		WithContext(ctx).
		Where(u.UserId.Eq(userID)).
		Preload(u.Link.BankAccounts.Pockets.Goals.Forecasts).
		Preload(u.Link.BankAccounts.Pockets.Goals.Milestones.Budget).
		Preload(u.Link.CreditAccounts.Aprs).
		Preload(u.Link.MortgageAccounts).
		Preload(u.Link.StudentLoanAccounts).
		Preload(u.Link.InvestmentAccounts.Holdings).
		Preload(u.Link.InvestmentAccounts.Securities).
		Preload(u.Link.Token).
		Preload(u.Link.PlaidLink).
		Preload(u.StripeSubscriptions).
		Preload(u.ActionableInsights).
		First()
	if err != nil {
		return nil, err
	}

	profile, err := record.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// GetAllUserProfiles obtains all user profiles presently in the database
func (db *Db) GetAllUserProfiles(ctx context.Context) ([]*schema.UserProfile, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-profiles"); span != nil {
		defer span.End()
	}

	// preload all dependencies
	// meaning for all connected accounts, obtain all accounts and subaccounts
	u := db.QueryOperator.UserProfileORM
	record, err := u.
		WithContext(ctx).
		Preload(u.Link.BankAccounts.Pockets.Goals.Forecasts).
		Preload(u.Link.BankAccounts.Pockets.Goals.Milestones.Budget).
		Preload(u.Link.CreditAccounts.Aprs).
		Preload(u.Link.MortgageAccounts).
		Preload(u.Link.StudentLoanAccounts).
		Preload(u.Link.InvestmentAccounts.Holdings).
		Preload(u.Link.InvestmentAccounts.Securities).
		Preload(u.Link.Token).
		Preload(u.Link.PlaidLink).
		Preload(u.StripeSubscriptions).
		Preload(u.ActionableInsights).
		Find()
	if err != nil {
		return nil, err
	}

	results := make([]*schema.UserProfile, 0, len(record))
	for _, r := range record {
		profile, err := r.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		results = append(results, &profile)
	}

	return results, nil
}

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

// GetUserProfileByEmail obtains a user profile by user_id if it exists within the database
func (db *Db) GetUserProfileByEmail(ctx context.Context, email string) (*schema.UserProfile, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-profile"); span != nil {
		defer span.End()
	}

	// validate the request
	if email == "" {
		return nil, fmt.Errorf("invalid stripe customer id. stripe customer id cannot be empty")
	}

	// preload all dependencies
	// meaning for all connected accounts, obtain all accounts and subaccounts
	u := db.QueryOperator.UserProfileORM
	record, err := u.
		WithContext(ctx).
		Where(u.Email.Eq(email)).
		Preload(u.Link.BankAccounts.Pockets.Goals.Forecasts).
		Preload(u.Link.BankAccounts.Pockets.Goals.Milestones.Budget).
		Preload(u.Link.CreditAccounts.Aprs).
		Preload(u.Link.MortgageAccounts).
		Preload(u.Link.StudentLoanAccounts).
		Preload(u.Link.InvestmentAccounts.Holdings).
		Preload(u.Link.InvestmentAccounts.Securities).
		Preload(u.Link.Token).
		Preload(u.Link.PlaidLink).
		Preload(u.StripeSubscriptions).
		Preload(u.ActionableInsights).
		First()
	if err != nil {
		return nil, err
	}

	profile, err := record.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
