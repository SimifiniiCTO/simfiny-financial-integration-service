package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// DeleteBankAccount implements DatabaseOperations
func (db *Db) DeleteBankAccount(ctx context.Context, bankAccountID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-bank-account"); span != nil {
		defer span.End()
	}

	if bankAccountID == 0 {
		return fmt.Errorf("bank account id must be provided. got: %d", bankAccountID)
	}

	// ensure the bank account exists
	b := db.QueryOperator.BankAccountORM
	if _, err := b.WithContext(ctx).Where(b.Id.Eq(bankAccountID)).First(); err != nil {
		return fmt.Errorf("bank account with id %d does not exist", bankAccountID)
	}

	// perform the delete operation
	// TODO: ensure all records are deleted especiallly ones with foreign key relationships
	result, err := b.WithContext(ctx).Where(b.Id.Eq(bankAccountID)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetBankAccount implements DatabaseOperations
func (db *Db) GetBankAccount(ctx context.Context, bankAccountID uint64) (*schema.BankAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-bank-account"); span != nil {
		defer span.End()
	}

	if bankAccountID == 0 {
		return nil, fmt.Errorf("bank account id must be provided. got: %d", bankAccountID)
	}

	// ensure the bank account exists
	b := db.QueryOperator.BankAccountORM
	bankAccountOrm, err := b.
		WithContext(ctx).
		Where(b.Id.Eq(bankAccountID)).
		Preload(b.Pockets.Goals.Forecasts).
		Preload(b.Pockets.Goals.Milestones.Budget.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("bank account with id %d does not exist", bankAccountID)
	}

	// convert orm to schema
	bankAccount, err := bankAccountOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &bankAccount, nil
}

// UpdateBankAccount implements DatabaseOperations
func (db *Db) UpdateBankAccount(ctx context.Context, bankAccount *schema.BankAccount) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-bank-account"); span != nil {
		defer span.End()
	}

	if bankAccount == nil {
		return fmt.Errorf("bank account must be provided. got: %v", bankAccount)
	}

	// ensure the bank account exists
	b := db.QueryOperator.BankAccountORM
	if _, err := b.WithContext(ctx).Where(b.Id.Eq(bankAccount.Id)).First(); err != nil {
		return fmt.Errorf("bank account with id %d does not exist", bankAccount.Id)
	}

	// update the bank account
	bankAccountOrm, err := bankAccount.ToORM(ctx)
	if err != nil {
		return err
	}

	// perform the update operation
	result, err := b.WithContext(ctx).Updates(bankAccountOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// CreateBankAccount implements DatabaseOperations
func (db *Db) CreateBankAccount(ctx context.Context, userID uint64, bankAccount *schema.BankAccount) (*schema.BankAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-profile"); span != nil {
		defer span.End()
	}

	if bankAccount == nil || userID == 0 {
		return nil, fmt.Errorf("bank account and user id must be provided. got: %v, %d", bankAccount, userID)
	}

	// validate the bankAccount
	if err := bankAccount.ValidateAll(); err != nil {
		return nil, err
	}

	// ensure the user profile exists
	u := db.QueryOperator.UserProfileORM
	userProfile, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).First()
	if err != nil {
		return nil, fmt.Errorf("user profile with user_id %d does not exist", userID)
	}

	// convert bank account to orm
	bankAcctOrm, err := bankAccount.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// append the new bank account to the user profile
	if err := u.BankAccounts.Model(userProfile).Append(&bankAcctOrm); err != nil {
		return nil, err
	}

	// perform the update operation
	result, err := u.WithContext(ctx).Updates(userProfile)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// bankAcctOrm reference at this point is auto updated with the db record id
	// as a byproduct of the update operation
	res, err := bankAcctOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
