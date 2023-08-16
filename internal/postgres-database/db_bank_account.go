package database

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gen/field"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
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

func (db *Db) UpdateBankAccounts(ctx context.Context, link *schema.Link, bankAccounts []*schema.BankAccount) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-bank-accounts"); span != nil {
		defer span.End()
	}

	if len(bankAccounts) == 0 {
		return fmt.Errorf("bank accounts must be provided. got: %v", bankAccounts)
	}

	if link == nil {
		return fmt.Errorf("link must be provided. got: %v", link)
	}

	// convert link to ORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	for _, bankAccount := range bankAccounts {
		// validate the bankAccount
		if err := bankAccount.ValidateAll(); err != nil {
			return err
		}
	}

	l := db.QueryOperator.LinkORM
	bankAcctsOrm := make([]*schema.BankAccountORM, 0, len(bankAccounts))

	for _, bankAccount := range bankAccounts {
		acctOrm, err := bankAccount.ToORM(ctx)
		if err != nil {
			return err
		}

		bankAcctsOrm = append(bankAcctsOrm, &acctOrm)
	}

	if err := l.BankAccounts.Model(&linkOrm).Replace(bankAcctsOrm...); err != nil {
		return fmt.Errorf("failed to update credit accounts: %v", err)
	}

	// perform the update operation
	result, err := l.WithContext(ctx).Updates(&linkOrm)
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
		db.Logger.Info("plaid link detected ... performing uniqueness check", zap.Any("link", link))
		// here we condition on the bank account plaid account id ... since this should be unique across all connected account
		if _, err = b.WithContext(ctx).Where(b.PlaidAccountId.Eq(bankAccount.PlaidAccountId), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same plaid account id already mapped to the link already exists
			return nil, fmt.Errorf("bank account with plaid account id %s already exists for link with id %d", bankAccount.PlaidAccountId, linkID)
		}

		db.Logger.Info("plaid link detected ... completed uniqueness check")
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

func (db *Db) CreateCreditAccount(ctx context.Context, linkID uint64, creditAccount *schema.CreditAccount) (*schema.CreditAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-credit-account"); span != nil {
		defer span.End()
	}

	if creditAccount == nil || linkID == 0 {
		return nil, fmt.Errorf("credit account and link id must be provided. got: %v, %d", creditAccount, linkID)
	}

	// validate the creditAccount
	if err := creditAccount.ValidateAll(); err != nil {
		return nil, err
	}

	// ensure the link exists
	l := db.QueryOperator.LinkORM
	b := db.QueryOperator.CreditAccountORM
	link, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).First()
	if err != nil {
		return nil, fmt.Errorf("link with id %d does not exist", linkID)
	}

	// given a credit account can be either tied to a plaid link or a manual link, we
	// need to ensure the credit account does not already exist for both types. Hence
	// we need to check the link type prior to deciphering which execution flow to take
	if link.LinkType == schema.LinkType_LINK_TYPE_MANUAL.String() {
		// for manually linked accounts, the account number must be unique hence that is what we use to condition upon
		if _, err = b.WithContext(ctx).Where(b.Number.Eq(creditAccount.Number), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same number already mapped to the link already exists
			return nil, fmt.Errorf("credit account with number %s already exists for link with id %d", creditAccount.Number, linkID)
		}
	} else if link.LinkType == schema.LinkType_LINK_TYPE_PLAID.String() {
		db.Logger.Info("plaid link detected ... performing uniqueness check", zap.Any("link", link))
		// here we condition on the credit account plaid account id ... since this should be unique across all connected account
		if _, err = b.WithContext(ctx).Where(b.PlaidAccountId.Eq(creditAccount.PlaidAccountId), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same plaid account id already mapped to the link already exists
			return nil, fmt.Errorf("credit account with plaid account id %s already exists for link with id %d", creditAccount.PlaidAccountId, linkID)
		}

		db.Logger.Info("plaid link detected ... completed uniqueness check")
	}

	creditAcctORM, err := creditAccount.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// append the new credit account to the link
	if err := l.CreditAccounts.Model(link).Append(&creditAcctORM); err != nil {
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

	// creditAccountOrm reference at this point is auto updated with the db record id
	// as a byproduct of the update operation
	res, err := creditAcctORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (db *Db) UpdateCreditAccounts(ctx context.Context, link *schema.Link, creditAccounts []*schema.CreditAccount) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-credit-accounts"); span != nil {
		defer span.End()
	}

	if len(creditAccounts) == 0 {
		return fmt.Errorf("credit accounts must be provided. got: %v", creditAccounts)
	}

	if link == nil {
		return fmt.Errorf("link must be provided. got: %v", link)
	}

	// convert link to ORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	for _, account := range creditAccounts {
		// validate the account
		if err := account.ValidateAll(); err != nil {
			return err
		}

		db.Logger.Info("validated bank account", zap.Any("credit_account", account))

	}

	l := db.QueryOperator.LinkORM
	creditAcctsOrm := make([]*schema.CreditAccountORM, 0, len(creditAccounts))
	for _, creditAccount := range creditAccounts {
		acctOrm, err := creditAccount.ToORM(ctx)
		if err != nil {
			return err
		}

		creditAcctsOrm = append(creditAcctsOrm, &acctOrm)
	}

	if err := l.CreditAccounts.Model(&linkOrm).Replace(creditAcctsOrm...); err != nil {
		return fmt.Errorf("failed to update credit accounts: %v", err)
	}

	// perform the update operation
	result, err := l.WithContext(ctx).Updates(&linkOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (db *Db) CreateStudentLoantAccount(ctx context.Context, linkID uint64, studentLoanAccount *schema.StudentLoanAccount) (*schema.StudentLoanAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-credit-account"); span != nil {
		defer span.End()
	}

	if studentLoanAccount == nil || linkID == 0 {
		return nil, fmt.Errorf("student loan account and link id must be provided. got: %v, %d", studentLoanAccount, linkID)
	}

	// validate the studentLoanAccount
	if err := studentLoanAccount.ValidateAll(); err != nil {
		return nil, err
	}

	// ensure the link exists
	l := db.QueryOperator.LinkORM
	b := db.QueryOperator.StudentLoanAccountORM
	link, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).First()
	if err != nil {
		return nil, fmt.Errorf("link with id %d does not exist", linkID)
	}

	// given a student loan account can be either tied to a plaid link or a manual link, we
	// need to ensure the student loan account does not already exist for both types. Hence
	// we need to check the link type prior to deciphering which execution flow to take
	if link.LinkType == schema.LinkType_LINK_TYPE_MANUAL.String() {
		// for manually linked accounts, the account number must be unique hence that is what we use to condition upon
		if _, err = b.WithContext(ctx).Where(b.Name.Eq(studentLoanAccount.Name), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same number already mapped to the link already exists
			return nil, fmt.Errorf("student account with name %s already exists for link with id %d", studentLoanAccount.Name, linkID)
		}
	} else if link.LinkType == schema.LinkType_LINK_TYPE_PLAID.String() {
		db.Logger.Info("plaid link detected ... performing uniqueness check", zap.Any("link", link))
		// here we condition on the student account plaid account id ... since this should be unique across all connected account
		if _, err = b.WithContext(ctx).Where(b.PlaidAccountId.Eq(studentLoanAccount.PlaidAccountId), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same plaid account id already mapped to the link already exists
			return nil, fmt.Errorf("student loan account with plaid account id %s already exists for link with id %d", studentLoanAccount.PlaidAccountId, linkID)
		}

		db.Logger.Info("plaid link detected ... completed uniqueness check")
	}

	studentLoanAcctORM, err := studentLoanAccount.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// append the new credit account to the link
	if err := l.StudentLoanAccounts.Model(link).Append(&studentLoanAcctORM); err != nil {
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

	// studentLoanAcctORM reference at this point is auto updated with the db record id
	// as a byproduct of the update operation
	res, err := studentLoanAcctORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (db *Db) UpdateStudentLoanAccounts(ctx context.Context, link *schema.Link, studentLoanAccounts []*schema.StudentLoanAccount) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-student-loan-accounts"); span != nil {
		defer span.End()
	}

	if len(studentLoanAccounts) == 0 {
		return fmt.Errorf("student loan accounts must be provided. got: %v", studentLoanAccounts)
	}

	if link == nil {
		return fmt.Errorf("link must be provided. got: %v", link)
	}

	// convert link to ORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	for _, account := range studentLoanAccounts {
		// validate the account
		if err := account.ValidateAll(); err != nil {
			return err
		}

		db.Logger.Info("validated bank account", zap.Any("student_account", account))

	}

	l := db.QueryOperator.LinkORM
	studentLoanAcctsOrm := make([]*schema.StudentLoanAccountORM, 0, len(studentLoanAccounts))
	for _, studentLoanAccount := range studentLoanAccounts {
		acctOrm, err := studentLoanAccount.ToORM(ctx)
		if err != nil {
			return err
		}

		studentLoanAcctsOrm = append(studentLoanAcctsOrm, &acctOrm)
	}

	if err := l.StudentLoanAccounts.Model(&linkOrm).Replace(studentLoanAcctsOrm...); err != nil {
		return fmt.Errorf("failed to update credit accounts: %v", err)
	}

	// perform the update operation
	result, err := l.WithContext(ctx).Updates(&linkOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// mortgage accounts
func (db *Db) CreateMortgageLoantAccount(ctx context.Context, linkID uint64, mortgageLoanAccount *schema.MortgageAccount) (*schema.MortgageAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-credit-account"); span != nil {
		defer span.End()
	}

	if mortgageLoanAccount == nil || linkID == 0 {
		return nil, fmt.Errorf("mortgage loan account and link id must be provided. got: %v, %d", mortgageLoanAccount, linkID)
	}

	// validate the mortgageLoanAccount
	if err := mortgageLoanAccount.ValidateAll(); err != nil {
		return nil, err
	}

	// ensure the link exists
	l := db.QueryOperator.LinkORM
	b := db.QueryOperator.MortgageAccountORM
	link, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).First()
	if err != nil {
		return nil, fmt.Errorf("link with id %d does not exist", linkID)
	}

	// given a mortgage loan account can be either tied to a plaid link or a manual link, we
	// need to ensure the mortgage loan account does not already exist for both types. Hence
	// we need to check the link type prior to deciphering which execution flow to take
	if link.LinkType == schema.LinkType_LINK_TYPE_MANUAL.String() {
		// for manually linked accounts, the account number must be unique hence that is what we use to condition upon
		if _, err = b.WithContext(ctx).Where(b.AccountNumber.Eq(mortgageLoanAccount.AccountNumber), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same number already mapped to the link already exists
			return nil, fmt.Errorf("mortgage account with number %s already exists for link with id %d", mortgageLoanAccount.AccountNumber, linkID)
		}
	} else if link.LinkType == schema.LinkType_LINK_TYPE_PLAID.String() {
		db.Logger.Info("plaid link detected ... performing uniqueness check", zap.Any("link", link))
		// here we condition on the mortgage account plaid account id ... since this should be unique across all connected account
		if _, err = b.WithContext(ctx).Where(b.PlaidAccountId.Eq(mortgageLoanAccount.PlaidAccountId), b.LinkId.Eq(linkID)).First(); err == nil {
			// account with the same plaid account id already mapped to the link already exists
			return nil, fmt.Errorf("mortgage loan account with plaid account id %s already exists for link with id %d", mortgageLoanAccount.PlaidAccountId, linkID)
		}

		db.Logger.Info("plaid link detected ... completed uniqueness check")
	}

	mortgageLoanAcctORM, err := mortgageLoanAccount.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// append the new credit account to the link
	if err := l.MortgageAccounts.Model(link).Append(&mortgageLoanAcctORM); err != nil {
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

	// mortgageLoanAcctORM reference at this point is auto updated with the db record id
	// as a byproduct of the update operation
	res, err := mortgageLoanAcctORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (db *Db) UpdateMortgageLoanAccounts(ctx context.Context, link *schema.Link, mortgageLoanAccounts []*schema.MortgageAccount) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-mortgage-loan-accounts"); span != nil {
		defer span.End()
	}

	if len(mortgageLoanAccounts) == 0 {
		return fmt.Errorf("mortgage loan accounts must be provided. got: %v", mortgageLoanAccounts)
	}

	if link == nil {
		return fmt.Errorf("link must be provided. got: %v", link)
	}

	// convert link to ORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	for _, account := range mortgageLoanAccounts {
		// validate the account
		if err := account.ValidateAll(); err != nil {
			return err
		}

		db.Logger.Info("validated account", zap.Any("mortgage_account", account))

	}

	l := db.QueryOperator.LinkORM
	mortgageLoanAcctsOrm := make([]*schema.MortgageAccountORM, 0, len(mortgageLoanAccounts))
	for _, account := range mortgageLoanAccounts {
		acctOrm, err := account.ToORM(ctx)
		if err != nil {
			return err
		}

		mortgageLoanAcctsOrm = append(mortgageLoanAcctsOrm, &acctOrm)
	}

	if err := l.MortgageAccounts.Model(&linkOrm).Replace(mortgageLoanAcctsOrm...); err != nil {
		return fmt.Errorf("failed to update credit accounts: %v", err)
	}

	// perform the update operation
	result, err := l.WithContext(ctx).Updates(&linkOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
