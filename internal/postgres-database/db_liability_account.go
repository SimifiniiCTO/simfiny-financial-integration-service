package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetCreditAccount implements DatabaseOperations
// It attempts to query the creidt account associated to a user by id as well as all associations
func (db *Db) GetCreditAccount(ctx context.Context, creditAccountId uint64) (*schema.CreditAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-bank-account"); span != nil {
		defer span.End()
	}

	if creditAccountId == 0 {
		return nil, fmt.Errorf("bank account id must be provided. got: %d", creditAccountId)
	}

	// ensure the bank account exists
	b := db.QueryOperator.CreditAccountORM
	creditAccountOrm, err := b.
		WithContext(ctx).
		Where(b.Id.Eq(creditAccountId)).
		Preload(b.Aprs).
		First()
	if err != nil {
		return nil, fmt.Errorf("bank account with id %d does not exist", creditAccountId)
	}

	// convert orm to schema
	account, err := creditAccountOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetCreditAccount implements DatabaseOperations
// It attempts to query the creidt account associated to a user by id as well as all associations
func (db *Db) GetSudentLoanAccount(ctx context.Context, studentLoanAccountId uint64) (*schema.StudentLoanAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-studentloan-account"); span != nil {
		defer span.End()
	}

	if studentLoanAccountId == 0 {
		return nil, fmt.Errorf("bank account id must be provided. got: %d", studentLoanAccountId)
	}

	// ensure the bank account exists
	b := db.QueryOperator.StudentLoanAccountORM
	studentLoanAccountOrm, err := b.
		WithContext(ctx).
		Where(b.Id.Eq(studentLoanAccountId)).
		First()
	if err != nil {
		return nil, fmt.Errorf("bank account with id %d does not exist", studentLoanAccountId)
	}

	// convert orm to schema
	account, err := studentLoanAccountOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
