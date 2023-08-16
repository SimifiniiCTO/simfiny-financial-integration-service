package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"go.uber.org/zap"
)

// syncCreditAccountsHelper is a helper function that is used to sync liability accounts.
// It takes in two slices of credit accounts, one that has been synced and one that is currently
// in the database. It then compares the two slices and updates the database accordingly.
// It returns an error if there is an error.
func (th *TaskHandler) synchronizePlaidLinkedCreditAccounts(ctx context.Context, link *apiv1.Link, syncedPlaidLinkedCreditAccounts []*apiv1.CreditAccount) error {
	// return immediately if there are no synced bank accounts
	if len(syncedPlaidLinkedCreditAccounts) == 0 {
		// log a warning if no accounts were found
		th.logger.Warn("no accounts found for user with the following linked account", zap.Uint64("link_id", link.GetId()))
		return nil
	}

	l := th.postgresDb.QueryOperator.LinkORM
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	// ensure we have enforced unique bank accounts
	currentCreditAccounts := set.From[*apiv1.CreditAccount](link.GetCreditAccounts())
	currentSyncedCreditAccounts := set.From[*apiv1.CreditAccount](syncedPlaidLinkedCreditAccounts)

	// we create a hashset of plaid account ids ofr the synced accounts
	// this is used to mark accounts stored in our database as inactive
	// if they are not present in the synced accounts
	currentSyncedCreditAccountIdsHashSet := set.New[string](0)
	currentSyncedCreditAccounts.ForEach(func(account *apiv1.CreditAccount) bool {
		currentSyncedCreditAccountIdsHashSet.Insert(account.PlaidAccountId)
		return true
	})

	// iterate over the current stored bank accounts and mark them as inactive
	// if they are not present in the synced accounts
	inactiveAccountIdsHashSet := set.New[string](0)
	currentCreditAccounts.ForEach(func(account *apiv1.CreditAccount) bool {
		if !currentSyncedCreditAccountIdsHashSet.Contains(account.PlaidAccountId) {
			th.logger.Info("marking credit account as inactive", zap.Any("account", account))
			account.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE

			record, err := account.ToORM(ctx)
			if err != nil {
				th.logger.Error("failed to update credit account", zap.Error(err))
				return false
			}

			if err := l.CreditAccounts.Model(&linkOrm).Replace(&record); err != nil {
				th.logger.Error("failed to update credit account", zap.Error(err))
				return false
			}

			inactiveAccountIdsHashSet.Insert(account.PlaidAccountId)
		}
		return true
	})

	// Check if all accounts are inactive
	if inactiveAccountIdsHashSet.Size() == currentCreditAccounts.Size() && currentCreditAccounts.Size() > 0 {
		th.logger.Warn("none of the linked credit accounts are active at plaid", zap.Uint64("link_id", link.Id))
		th.logger.Warn("please link a new credit account to continue using the service or update the account link", zap.Any("credit accounts", currentCreditAccounts))
		return nil
	}

	// now we need to perform a cross reference to determine if any accounts need to be added or updated
	// we create a hashset of plaid account ids ofr the current stored accounts
	// this is used to determine if we need to add or update an account
	currentCreditAccountIdsHashSet := set.New[string](0)
	currentCreditAccountPlaidIdToCreditAccountMap := make(map[string]*apiv1.CreditAccount)
	currentCreditAccounts.ForEach(func(account *apiv1.CreditAccount) bool {
		currentCreditAccountIdsHashSet.Insert(account.PlaidAccountId)
		currentCreditAccountPlaidIdToCreditAccountMap[account.PlaidAccountId] = account
		return true
	})

	// iterate over the synced credit accounts and determine if we need to add or update an account
	accountsToBeAdded := set.New[*apiv1.CreditAccount](0)
	accountsToBeUpdated := set.New[*apiv1.CreditAccount](0)

	currentSyncedCreditAccounts.ForEach(func(account *apiv1.CreditAccount) bool {
		if currentCreditAccountIdsHashSet.Contains(account.PlaidAccountId) {
			// we need to update the account
			if storedCreditAccount, ok := currentCreditAccountPlaidIdToCreditAccountMap[account.PlaidAccountId]; ok {
				copyCreditAccount(account, storedCreditAccount)
				accountsToBeUpdated.Insert(storedCreditAccount)
			} else {
				th.logger.Error("failed to find credit account in map. this should not happen", zap.Any("account", account))
			}

		} else {
			// we need to add the account
			accountToBeAdded := account
			accountsToBeAdded.Insert(accountToBeAdded)
		}
		return true
	})

	// print out the contents of the sets
	th.logger.Info("accounts to be added", zap.Any("accounts", accountsToBeAdded))
	th.logger.Info("accounts to be updated", zap.Any("accounts", accountsToBeUpdated))

	// add the new accountsToBeAdded
	if accountsToBeUpdated.Size() > 0 {
		if err := th.postgresDb.UpdateCreditAccounts(ctx, link, accountsToBeUpdated.Slice()); err != nil {
			th.logger.Error("failed to update credit accounts", zap.Error(err))
			return err
		}
	}

	if accountsToBeAdded.Size() > 0 {
		if accountsToBeUpdated.Size() > 0 {
			th.logger.Info("should not be adding new accounts")
		} else {
			accountsToBeAdded.ForEach(func(account *apiv1.CreditAccount) bool {
				if _, err := th.postgresDb.CreateCreditAccount(ctx, link.Id, account); err != nil {
					th.logger.Error("failed to create credit account", zap.Error(err))
					return false
				}

				return true
			})
		}

	}

	return nil
}

// copyBankAccount is a helper method that copies the contents of a source bank account to a destination bank account
func copyCreditAccount(src *apiv1.CreditAccount, dest *apiv1.CreditAccount) {
	if src == nil || dest == nil {
		return
	}

	dest.Name = src.Name
	dest.Number = src.Number
	dest.Type = src.Type
	dest.Balance = src.Balance
	dest.CurrentFunds = src.CurrentFunds
	dest.BalanceLimit = src.BalanceLimit
	dest.PlaidAccountId = src.PlaidAccountId
	dest.Subtype = src.Subtype
}
