package database

import (
	"context"
	"fmt"

	"gorm.io/gen/field"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
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

	// perform the deletion operation
	result, err := b.
		WithContext(ctx).
		Where(b.Id.Eq(bankAccountID)).
		Select(field.AssociationFields).
		Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetBankAccount implements DatabaseOperations
// It attempts the bank account associated to a user by id as well as all associations
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
		Preload(b.Pockets.Goals.Milestones.Budget).
		Preload(b.Pockets.Goals.Forecasts).
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

// UpdateBankAccount implements DatabaseOperations and updates a bank account
// A bank account can be updated if the following conditions are met:
//   - the bank account exists
//   - the bank account is valid
//   - the bank account is tied to a link
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
	acct, err := b.WithContext(ctx).Where(b.Id.Eq(bankAccount.Id)).First()
	if err != nil {
		return fmt.Errorf("bank account with id %d does not exist", bankAccount.Id)
	}

	if acct == nil {
		return fmt.Errorf("bank account with id %d does not exist", bankAccount.Id)
	}

	// convert the bank account object to orm type
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

func (db *Db) UpdateBankAccounts(ctx context.Context, bankAccounts []*schema.BankAccount) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-bank-accounts"); span != nil {
		defer span.End()
	}

	if len(bankAccounts) == 0 {
		return fmt.Errorf("bank accounts must be provided. got: %v", bankAccounts)
	}

	b := db.QueryOperator.BankAccountORM
	bankAcctsOrm := make([]*schema.BankAccountORM, 0, len(bankAccounts))
	for _, bankAccount := range bankAccounts {
		acctOrm, err := bankAccount.ToORM(ctx)
		if err != nil {
			return err
		}

		bankAcctsOrm = append(bankAcctsOrm, &acctOrm)
	}

	// perform the update operation
	result, err := b.WithContext(ctx).Updates(bankAcctsOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// CreateBankAccount creates a new bank account for a given link
// A link can have a multitude of bank accounts so as part of the creation process
// we need to ensure the following conditions are met:
//   - the link exists
//   - the bank account does not already exist for the given link
//   - the bank account is valid
//
// on the conditions that the are above met, we append the new bank account to the link
func (db *Db) CreateBankAccount(ctx context.Context, linkID uint64, bankAccount *schema.BankAccount) (*schema.BankAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-profile"); span != nil {
		defer span.End()
	}

	if bankAccount == nil || linkID == 0 {
		return nil, fmt.Errorf("bank account and link id must be provided. got: %v, %d", bankAccount, linkID)
	}

	// validate the bankAccount
	if err := bankAccount.ValidateAll(); err != nil {
		return nil, err
	}

	// ensure the link exists
	l := db.QueryOperator.LinkORM
	b := db.QueryOperator.BankAccountORM
	link, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).First()
	if err != nil {
		return nil, fmt.Errorf("link with id %d does not exist", linkID)
	}

	// given a bank account can be either tied to a plaid link or a manual link, we
	// need to ensure the bank account does not already exist for both types. Hence
	// we need to check the link type prior to deciphering which execution flow to take
	if link.LinkType == schema.LinkType_LINK_TYPE_MANUAL.String() {
		// for manually linked accounts, the account number must be unique hence that is what we use to condition upon
		if _, err = b.WithContext(ctx).Where(b.Number.Eq(bankAccount.Number), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same number already mapped to the link already exists
			return nil, fmt.Errorf("bank account with number %s already exists for link with id %d", bankAccount.Number, linkID)
		}
	} else if link.LinkType == schema.LinkType_LINK_TYPE_PLAID.String() {
		// here we condition on the bank account plaid account id ... since this should be unique across all connected account
		if _, err = b.WithContext(ctx).Where(b.PlaidAccountId.Eq(bankAccount.PlaidAccountId), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same plaid account id already mapped to the link already exists
			return nil, fmt.Errorf("bank account with plaid account id %s already exists for link with id %d", bankAccount.PlaidAccountId, linkID)
		}
	}

	bankAcctORM, err := bankAccount.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// append the new bank account to the link
	if err := l.BankAccounts.Model(link).Append(&bankAcctORM); err != nil {
		return nil, err
	}

	// perform the update operation of the link token
	result, err := l.WithContext(ctx).Updates(link)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// bankAcctOrm reference at this point is auto updated with the db record id
	// as a byproduct of the update operation
	res, err := bankAcctORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
