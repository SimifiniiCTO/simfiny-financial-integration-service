package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"go.uber.org/zap"
)

// syncCreditAccountsHelper is a helper function that is used to sync liability accounts.
// It takes in two slices of credit accounts, one that has been synced and one that is currently
// in the database. It then compares the two slices and updates the database accordingly.
// It returns an error if there is an error.
func (t *TaskHandler) syncCreditAccountsHelper(ctx context.Context, link *apiv1.Link, syncedCreditAccounts, currentCreditAccounts []*apiv1.CreditAccount) error {
	// check the size of the synced credit accounts
	// and return if no new synced accounts are available
	if len(syncedCreditAccounts) == 0 {
		return nil
	}

	// log the current and synced accounts
	t.logger.Info("syncing liability accounts", zap.Any("syncedCreditAccounts", syncedCreditAccounts), zap.Any("currentCreditAccounts", currentCreditAccounts))

	// for the current credit accounts, we create a hashmap comprised of plaidAccountId to the respective account
	// this is done so that we can easily check if a synced credit account is already in the database
	currentCreditAccountsMap := make(map[string]*apiv1.CreditAccount)
	for _, currentCreditAccount := range currentCreditAccounts {
		currentCreditAccountsMap[currentCreditAccount.PlaidAccountId] = currentCreditAccount
	}

	// we iterate over the synced credit accounts and cross reference them with the existing credit accounts
	// if the synced credit account is not found in the existing credit account set, then we know it is a new account that must be added
	// to the database
	accountsToBeUpdated := make([]*apiv1.CreditAccountORM, 0, len(syncedCreditAccounts))
	accountsToBeAdded := make([]*apiv1.CreditAccountORM, 0, len(syncedCreditAccounts))

	// travers the synced credit accounts
	for _, syncedCreditAccount := range syncedCreditAccounts {
		// we add the synced credit account to the new accounts to be added set if the current credit accounts in the database are empty
		if len(currentCreditAccounts) == 0 {
			record, err := syncedCreditAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new credit account to set", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		}

		// if the synced credit account is not found in the current credit accounts set, then we know it is a new account that must be added
		// to the database
		if _, ok := currentCreditAccountsMap[syncedCreditAccount.PlaidAccountId]; !ok {
			record, err := syncedCreditAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new credit account to set", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		} else {
			// if the synced credit account is found in the current credit accounts set, then we know it is an existing account that must be updated
			record, err := syncedCreditAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding existing credit account to set", zap.Any("account", record))
			accountsToBeUpdated = append(accountsToBeUpdated, &record)
			continue
		}
	}

	// perform the database operations to save the credit accounts
	l := t.postgresDb.QueryOperator.LinkORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	if len(accountsToBeUpdated) > 0 {
		t.logger.Info("updating existing credit accounts", zap.Any("accounts", accountsToBeUpdated))
		// update the credit accounts in the database
		if err := l.CreditAccounts.Model(&linkOrm).Replace(accountsToBeUpdated...); err != nil {
			return err
		}
	}

	if len(accountsToBeAdded) > 0 {
		t.logger.Info("adding new credit accounts", zap.Any("accounts", accountsToBeAdded))
		// add the new credit accounts to the database
		if err := l.CreditAccounts.Model(&linkOrm).Append(accountsToBeAdded...); err != nil {
			return err
		}
	}

	return nil
}

// syncMortgageAccountsHelper is a helper function that is used to sync mortgage liability accounts.
// It takes in two slices of credit accounts, one that has been synced and one that is currently
func (t *TaskHandler) syncMortgageAccountsHelper(ctx context.Context, link *apiv1.Link, syncedMortgageAccounts, currentMortgageAccounts []*apiv1.MortgageAccount) error {
	// check the size of the synced morgage accounts
	// and return if no new synced accounts are available
	if len(syncedMortgageAccounts) == 0 {
		return nil
	}

	// log the current and synced accounts
	t.logger.Info("syncing mortgage liability accounts", zap.Any("syncedMortgageAccounts", syncedMortgageAccounts), zap.Any("currentMortgageAccounts", currentMortgageAccounts))

	// for the current mortgage accounts, we create a hashmap comprised of plaidAccountId to the respective account
	// this is done so that we can easily check if a synced mortgage account is already in the database
	currentMortgageAccountsMap := make(map[string]*apiv1.MortgageAccount)
	for _, currentMortgageAccount := range currentMortgageAccounts {
		currentMortgageAccountsMap[currentMortgageAccount.PlaidAccountId] = currentMortgageAccount
	}

	// we iterate over the synced mortgage accounts and cross reference them with the existing mortgage accounts
	// if the synced mortgage account is not found in the existing mortgage account set, then we know it is a new account that must be added
	// to the database
	accountsToBeUpdated := make([]*apiv1.MortgageAccountORM, 0, len(syncedMortgageAccounts))
	accountsToBeAdded := make([]*apiv1.MortgageAccountORM, 0, len(syncedMortgageAccounts))

	// traverse the synced mortgage accounts
	for _, syncedMortgageAccount := range syncedMortgageAccounts {
		// we add the synced mortgage account to the new accounts to be added set if the current mortgage accounts in the database are empty
		if len(currentMortgageAccounts) == 0 {
			record, err := syncedMortgageAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new mortgage account to set", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		}

		// if the synced mortgage account is not found in the current mortgage accounts set, then we know it is a new account that must be added
		// to the database
		if _, ok := currentMortgageAccountsMap[syncedMortgageAccount.PlaidAccountId]; !ok {
			record, err := syncedMortgageAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new mortgage account to set", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		} else {
			// if the synced mortgage account is found in the current mortgage accounts set, then we know it is an existing account that must be updated
			record, err := syncedMortgageAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding existing mortgage account to set", zap.Any("account", record))
			accountsToBeUpdated = append(accountsToBeUpdated, &record)
			continue
		}
	}

	// perform the database operations to save the mortgage accounts
	l := t.postgresDb.QueryOperator.LinkORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	if len(accountsToBeUpdated) > 0 {
		t.logger.Info("updating existing mortgage accounts", zap.Any("accounts", accountsToBeUpdated))
		// update the mortgage accounts in the database
		if err := l.MortgageAccounts.Model(&linkOrm).Replace(accountsToBeUpdated...); err != nil {
			return err
		}
	}

	if len(accountsToBeAdded) > 0 {
		t.logger.Info("adding new mortgage accounts", zap.Any("accounts", accountsToBeAdded))
		// add the new credit accounts to the database
		if err := l.MortgageAccounts.Model(&linkOrm).Append(accountsToBeAdded...); err != nil {
			return err
		}
	}

	return nil
}

// syncStudentLoanAccountsHelper is a helper function that is used to sync student loan liability accounts.
// It takes in two slices of student loan accounts, one that has been synced and one that is currently
func (t *TaskHandler) syncStudentLoanAccountsHelper(ctx context.Context, link *apiv1.Link, syncedStudentLoanAccounts, currentStudentLoanAccounts []*apiv1.StudentLoanAccount) error {
	// check the size of the synced morgage accounts
	// and return if no new synced accounts are available
	if len(syncedStudentLoanAccounts) == 0 {
		return nil
	}

	// log the current and synced accounts
	t.logger.Info("syncing student loan liability accounts", zap.Any("syncedStudentLoanAccounts", syncedStudentLoanAccounts), zap.Any("currentStudentLoanAccounts", currentStudentLoanAccounts))

	// for the current student loan accounts, we create a hashmap comprised of plaidAccountId to the respective account
	// this is done so that we can easily check if a synced student loan account is already in the database
	currentStudentLoanAccountsMap := make(map[string]*apiv1.StudentLoanAccount)
	for _, studentLoanAccount := range currentStudentLoanAccounts {
		currentStudentLoanAccountsMap[studentLoanAccount.PlaidAccountId] = studentLoanAccount
	}

	// we iterate over the synced student loan accounts and cross reference them with the existing student loan accounts
	// if the synced student loan account is not found in the existing student loan account set, then we know it is a new account that must be added
	// to the database
	accountsToBeUpdated := make([]*apiv1.StudentLoanAccountORM, 0, len(syncedStudentLoanAccounts))
	accountsToBeAdded := make([]*apiv1.StudentLoanAccountORM, 0, len(syncedStudentLoanAccounts))

	// traverse the synced student loan accounts
	for _, syncedStudentLoanAccount := range syncedStudentLoanAccounts {
		// we add the synced student loan account to the new accounts to be added set if the current student loan accounts in the database are empty
		if len(currentStudentLoanAccounts) == 0 {
			record, err := syncedStudentLoanAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new student loan account to set", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		}

		// if the synced student loan account is not found in the current student loan accounts set, then we know it is a new account that must be added
		// to the database
		if _, ok := currentStudentLoanAccountsMap[syncedStudentLoanAccount.PlaidAccountId]; !ok {
			record, err := syncedStudentLoanAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new student loan account to set", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		} else {
			// if the synced student loan account is found in the current student loan accounts set, then we know it is an existing account that must be updated
			record, err := syncedStudentLoanAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding existing student loan account to set", zap.Any("account", record))
			accountsToBeUpdated = append(accountsToBeUpdated, &record)
			continue
		}
	}

	// perform the database operations to save the student loan accounts
	l := t.postgresDb.QueryOperator.LinkORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	if len(accountsToBeUpdated) > 0 {
		t.logger.Info("updating existing student loan accounts", zap.Any("accounts", accountsToBeUpdated))
		// update the student loan accounts in the database
		if err := l.StudentLoanAccounts.Model(&linkOrm).Replace(accountsToBeUpdated...); err != nil {
			return err
		}
	}

	if len(accountsToBeAdded) > 0 {
		t.logger.Info("adding new student accounts", zap.Any("accounts", accountsToBeAdded))
		// add the new credit accounts to the database
		if err := l.StudentLoanAccounts.Model(&linkOrm).Append(accountsToBeAdded...); err != nil {
			return err
		}
	}

	return nil
}
