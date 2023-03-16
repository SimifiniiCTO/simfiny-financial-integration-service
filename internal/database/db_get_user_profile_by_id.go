package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

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
	u := db.QueryOperator.UserProfileORM
	record, err := u.
		WithContext(ctx).
		Where(u.UserId.Eq(userID)).
		Preload(u.MortgageAccounts).
		Preload(u.StudentLoanAccounts).
		Preload(u.CreditAccounts.Aprs).
		Preload(u.InvestmentAccounts.Holdings).
		Preload(u.InvestmentAccounts.Securities).
		Preload(u.BankAccounts.Pockets.Goals.Forecasts).
		Preload(u.BankAccounts.Pockets.Goals.Milestones.Budget.Category).
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
