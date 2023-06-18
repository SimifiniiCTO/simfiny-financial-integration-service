package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// GetMortgageAccount implements DatabaseOperations
// It attempts to query the mortgage account associated to a user by id as well as all associations
func (db *Db) GetMortgageAccount(ctx context.Context, MortgageAccountId uint64) (*schema.MortgageAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-bank-account"); span != nil {
		defer span.End()
	}

	if MortgageAccountId == 0 {
		return nil, fmt.Errorf("bank account id must be provided. got: %d", MortgageAccountId)
	}

	// ensure the bank account exists
	b := db.QueryOperator.MortgageAccountORM
	MortgageAccountOrm, err := b.
		WithContext(ctx).
		Where(b.Id.Eq(MortgageAccountId)).
		First()
	if err != nil {
		return nil, fmt.Errorf("bank account with id %d does not exist", MortgageAccountId)
	}

	// convert orm to schema
	account, err := MortgageAccountOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
