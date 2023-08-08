package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"go.uber.org/zap"
)

// syncBankAccountsHelper is a helper function that is used to sync bank accounts.
// It takes in two slices of bank accounts, one that has been synced and one that is currently
// in the database. It then compares the two slices and updates the database accordingly.
// It returns an error if there is an error.
func (t *TaskHandler) syncBankAccountsHelperV2(ctx context.Context, link *apiv1.Link, syncedBankAccounts []*apiv1.BankAccount) error {
	// check the size of the synced bank accounts
	// and return if no new synced accounts are available
	if len(syncedBankAccounts) == 0 {
		return nil
	}

	// perform the database operations to save the credit accounts
	l := t.postgresDb.QueryOperator.LinkORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	accountsToBeUpdated := make([]*apiv1.BankAccount, 0, len(syncedBankAccounts))
	accountsToBeAdded := make([]*apiv1.BankAccountORM, 0, len(syncedBankAccounts))
	currentBankAccounts := link.GetBankAccounts()
	// if the current bank accounts are empty, then we add all the synced bank accounts to the database
	if len(currentBankAccounts) == 0 {
		t.logger.Info("no bank accounts found in database, adding all synced bank accounts")
		for _, syncedBankAccount := range syncedBankAccounts {
			record, err := syncedBankAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			accountsToBeAdded = append(accountsToBeAdded, &record)
		}

		if err := l.BankAccounts.Model(&linkOrm).Append(accountsToBeAdded...); err != nil {
			return err
		}

		return nil
	}

	// log the current and synced accounts
	t.logger.Info("syncing bank accounts", zap.Any("syncedBankAccounts", len(syncedBankAccounts)), zap.Any("currentBankAccounts", len(currentBankAccounts)))

	// for the current bank accounts, we create a hashmap comprised of plaidAccountId to the respective account
	// this is done so that we can easily check if a synced bank account is already in the database
	currentBankAccountsMap := make(map[string]*apiv1.BankAccount)
	for _, currentBankAccount := range currentBankAccounts {
		currentBankAccountsMap[currentBankAccount.PlaidAccountId] = currentBankAccount
	}

	// construct a synced accounts map
	// this is done so that we can easily check if a synced bank account is already in the database
	syncedBankAccountsMap := make(map[string]*apiv1.BankAccount)
	for _, syncedBankAccount := range syncedBankAccounts {
		syncedBankAccountsMap[syncedBankAccount.PlaidAccountId] = syncedBankAccount
	}

	// we iterate over the synced bank accounts and cross reference them with the existing bank accounts
	// if the synced bank account is not found in the existing bank account set, then we know it is a new account that must be added
	// to the database

	bankAccountIds := make([]uint64, 0, len(currentBankAccounts))
	// If an account is no longer visible in plaid that means that we won't receive updates for that account anymore. If
	// this happens, log something and mark that account as inactive. This way we can inform the user that the account
	// is no longer receiving updates.
	for _, account := range currentBankAccounts {
		if _, ok := syncedBankAccountsMap[account.PlaidAccountId]; !ok {
			t.logger.Info("account no longer visible in plaid, marking as inactive", zap.Any("account", account))
			account.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE
			// update the bank account in the database
			// TODO: maybe emit a metric here
			if err := t.postgresDb.UpdateBankAccount(ctx, account); err != nil {
				return err
			}

			bankAccountIds = append(bankAccountIds, account.Id)
		}
	}

	// if all the bank accounts are inactive, then we return
	if (len(bankAccountIds)) == len(currentBankAccounts) && len(currentBankAccounts) > 0 {
		t.logger.Warn("none of the linked bank accounts are active at plaid", zap.Uint64("link_id", link.Id))
		t.logger.Warn("please link a new bank account to continue using the service", zap.Any("bank accounts", currentBankAccounts))
		return nil
	}

	// traverse the synced bank accounts
	for _, syncedBankAccount := range syncedBankAccounts {
		// if the synced bank account is not found in the current bank accounts set, then we know it is a new account that must be added
		// to the database
		if storedBankAccount, ok := currentBankAccountsMap[syncedBankAccount.PlaidAccountId]; !ok {
			record, err := syncedBankAccount.ToORM(ctx)
			if err != nil {
				return err
			}

			t.logger.Info("adding new bank account to set since we could not find active ones", zap.Any("account", record))
			accountsToBeAdded = append(accountsToBeAdded, &record)
			continue
		} else {
			// add all fields in the new synced bank account to the initial record
			// this is done so that we can update the existing bank account in the database
			storedBankAccount.Name = syncedBankAccount.Name
			storedBankAccount.Number = syncedBankAccount.Number
			storedBankAccount.Balance = syncedBankAccount.Balance
			storedBankAccount.Currency = syncedBankAccount.Currency
			storedBankAccount.CurrentFunds = syncedBankAccount.CurrentFunds
			storedBankAccount.BalanceLimit = syncedBankAccount.BalanceLimit
			storedBankAccount.PlaidAccountId = syncedBankAccount.PlaidAccountId
			storedBankAccount.Subtype = syncedBankAccount.Subtype
			storedBankAccount.Type = syncedBankAccount.Type
			storedBankAccount.Status = syncedBankAccount.Status

			t.logger.Info("updating existing bank account to set", zap.Any("account", storedBankAccount))
			accountsToBeUpdated = append(accountsToBeUpdated, storedBankAccount)
			continue
		}
	}

	if len(accountsToBeUpdated) > 0 {
		t.logger.Info("updating existing bank accounts", zap.Any("accounts", accountsToBeUpdated))
		if err = t.postgresDb.UpdateBankAccounts(ctx, accountsToBeUpdated); err != nil {
			t.logger.Error("failed to update bank accounts", zap.Error(err))
			return err
		}
	}

	if len(accountsToBeAdded) > 0 {
		t.logger.Info("adding new bank accounts", zap.Any("accounts", accountsToBeAdded))
		// add the new bank accounts to the database
		if err := l.BankAccounts.Model(&linkOrm).Append(accountsToBeAdded...); err != nil {
			return err
		}
	}

	return nil
}

func (t *TaskHandler) syncBankAccountsHelper(ctx context.Context, link *apiv1.Link, syncedBankAccounts []*apiv1.BankAccount) error {
	if len(syncedBankAccounts) == 0 {
		return nil
	}

	l := t.postgresDb.QueryOperator.LinkORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	// ensure we have enforced unique bank accounts
	currentBankAccounts := set.From[*apiv1.BankAccount](link.GetBankAccounts())
	currentSyncedBankAccounts := set.From[*apiv1.BankAccount](syncedBankAccounts)
	plaidAccountIdToCurrentBankAccountMap := make(map[string]*apiv1.BankAccount)
	plaidAccountIdToSyncedBankAccountMap := make(map[string]*apiv1.BankAccount)

	// populate the plaid account id to current bank account map
	currentBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		plaidAccountIdToCurrentBankAccountMap[account.PlaidAccountId] = account
		return true
	})

	t.logger.Info("populated plaid account id to current bank account map", zap.Any("map", plaidAccountIdToCurrentBankAccountMap))

	currentSyncedBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		plaidAccountIdToSyncedBankAccountMap[account.PlaidAccountId] = account
		return true
	})

	t.logger.Info("populated plaid account id to synced bank account map", zap.Any("map", plaidAccountIdToSyncedBankAccountMap))

	var accountsToBeUpdated []*apiv1.BankAccount
	var accountsToBeAdded []*apiv1.BankAccountORM
	var inactiveAccountIds []uint64

	currentSyncedBankAccounts.ForEach(func(syncedBankAccount *apiv1.BankAccount) bool {
		storedBankAccount, exists := plaidAccountIdToCurrentBankAccountMap[syncedBankAccount.PlaidAccountId]

		if !exists {
			// New bank account
			t.logger.Info("adding new bank account to set since we could not find active ones", zap.Any("account", syncedBankAccount))

			record, err := syncedBankAccount.ToORM(ctx)
			if err != nil {
				t.logger.Error("failed to convert bank account to orm", zap.Error(err))
				return false
			}
			accountsToBeAdded = append(accountsToBeAdded, &record)
		} else {
			t.logger.Info("updating existing bank account to set", zap.Any("account", storedBankAccount))

			// Update existing bank account
			storedBankAccount.Name = syncedBankAccount.Name
			storedBankAccount.Number = syncedBankAccount.Number
			storedBankAccount.Balance = syncedBankAccount.Balance
			storedBankAccount.Currency = syncedBankAccount.Currency
			storedBankAccount.CurrentFunds = syncedBankAccount.CurrentFunds
			storedBankAccount.BalanceLimit = syncedBankAccount.BalanceLimit
			storedBankAccount.PlaidAccountId = syncedBankAccount.PlaidAccountId
			storedBankAccount.Subtype = syncedBankAccount.Subtype
			storedBankAccount.Type = syncedBankAccount.Type
			storedBankAccount.Status = syncedBankAccount.Status
			accountsToBeUpdated = append(accountsToBeUpdated, storedBankAccount)
		}

		return true
	})

	// Mark accounts as inactive if they are not in syncedBankAccounts
	currentBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		if _, exists := plaidAccountIdToSyncedBankAccountMap[account.PlaidAccountId]; !exists {
			t.logger.Info("account no longer visible in plaid, marking as inactive", zap.Any("account", account))
			account.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE

			if err := t.postgresDb.UpdateBankAccount(ctx, account); err != nil {
				t.logger.Error("failed to update bank account", zap.Error(err))
				return false
			}

			inactiveAccountIds = append(inactiveAccountIds, account.Id)
			return true
		}

		return true
	})

	// Check if all accounts are inactive
	if len(inactiveAccountIds) == currentBankAccounts.Size() && currentBankAccounts.Size() > 0 {
		t.logger.Warn("none of the linked bank accounts are active at plaid", zap.Uint64("link_id", link.Id))
		t.logger.Warn("please link a new bank account to continue using the service", zap.Any("bank accounts", currentBankAccounts))
		return nil
	}

	// Update existing bank accounts
	if len(accountsToBeUpdated) > 0 {
		t.logger.Info("updating existing bank accounts", zap.Any("accounts", accountsToBeUpdated))
		if err = t.postgresDb.UpdateBankAccounts(ctx, accountsToBeUpdated); err != nil {
			t.logger.Error("failed to update bank accounts", zap.Error(err))
			return err
		}
	}

	// Add new bank accounts
	if len(accountsToBeAdded) > 0 {
		t.logger.Info("adding new bank accounts", zap.Any("accounts", accountsToBeAdded))
		// add the new bank accounts to the database
		if err := l.BankAccounts.Model(&linkOrm).Append(accountsToBeAdded...); err != nil {
			return err
		}
	}

	return nil
}
