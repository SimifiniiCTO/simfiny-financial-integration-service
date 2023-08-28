package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/google/uuid"
)

func (db *Db) GetBalanceHistoryByAccountID(ctx context.Context, accountId *string) ([]*schema.AccountBalanceHistory, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-account balances-for-account-id"); span != nil {
		defer span.End()
	}

	if accountId == nil {
		return nil, fmt.Errorf("account ID cannot be nil")
	}

	result := make([]schema.AccountBalanceHistoryInternal, 0)
	if err := db.queryEngine.NewSelect().Model(&result).Where("AccountId = ?", *accountId).Scan(ctx); err != nil {
		return nil, err
	}

	data := make([]*schema.AccountBalanceHistory, 0)
	for _, record := range result {
		res, err := record.ConvertTo()
		if err != nil {
			return nil, err
		}

		data = append(data, res)
	}

	return data, nil
}

func (db *Db) AddAccountBalances(ctx context.Context, userId *uint64, accountBalances []*schema.AccountBalanceHistory) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-account-balance"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return fmt.Errorf("user ID cannot be nil")
	}

	if len(accountBalances) == 0 {
		return fmt.Errorf("transactions length must be greater than 0")
	}

	balances := make([]schema.AccountBalanceHistoryInternal, 0, len(accountBalances))
	for _, balanceObject := range accountBalances {
		if balanceObject.AccountId == "" {
			return fmt.Errorf("account balance object must have an associated account id at creation time")
		}

		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		// associate the user id to the transaction of interest
		balanceObject.UserId = *userId
		balanceObject.Id = id.String()

		// validate transactions
		if err := balanceObject.Validate(); err != nil {
			return err
		}

		db.Logger.Info("transforming balance object")

		record, err := balanceObject.ConvertToInternal()
		if err != nil {
			return err
		}

		db.Logger.Info("appending balance object")

		balances = append(balances, *record)
	}

	db.Logger.Info("inserting balance object")
	if _, err := db.queryEngine.NewInsert().Model(&balances).Exec(ctx); err != nil {
		return err
	}

	return nil
}
