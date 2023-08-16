package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// syncBankAccounts is a method on the TaskHandler struct that takes a context, a user id, a link id, and an access token
// as parameters. It is used to sync all accounts for a given user. It does this by querying the database for the link,
// querying plaid for all accounts, and then performing a cross reference to determine if any accounts need to be added
// or updated.
func (th *TaskHandler) syncBankAccounts(ctx context.Context, userId uint64, link *apiv1.Link, accessToken string) ([]*apiv1.BankAccount, error) {
	var (
		plaidClient = th.plaidClient
	)

	if link == nil {
		th.logger.Warn("provided link is nil")
		return nil, errors.New("link is nil")
	}

	if userId == 0 {
		th.logger.Warn("provided user id is 0")
		return nil, errors.New("user id is 0")
	}

	// TODO: from the linked account we need to be able to see all accounts that the user may have chosen to deselect from access to us and remove them
	// we need explicit support for this

	if link.PlaidLink == nil {
		th.logger.Warn("provided link does not have any plaid credentials")
		return nil, errors.New("plaid link not found")
	}

	// query plaid for all of the user's accounts
	syncBankAccounts, err := plaidClient.GetAccounts(ctx, accessToken, userId)
	if err != nil {
		return nil, err
	}

	// we only perform account sync cross referencing if the
	// number of synced accounts is greater than 0
	if len(syncBankAccounts) == 0 {
		th.logger.Warn("no accounts found for user", zap.Uint64("user_id", userId))
		return nil, nil
	} else {
		if err := th.synchronizePlaidLinkedBankAccounts(ctx, link, syncBankAccounts); err != nil {
			return nil, err
		}
	}

	return syncBankAccounts, nil
}

func (th *TaskHandler) synchronizePlaidLinkedBankAccounts(ctx context.Context, link *apiv1.Link, syncedPlaidLinkedBankAccounts []*apiv1.BankAccount) error {
	// return immediately if there are no synced bank accounts
	if len(syncedPlaidLinkedBankAccounts) == 0 {
		// log a warning if no accounts were found
		th.logger.Warn("no accounts found for user with the following linked account", zap.Uint64("link_id", link.GetId()))
		return nil
	}

	// ensure we have enforced unique bank accounts
	currentBankAccounts := set.From[*apiv1.BankAccount](link.GetBankAccounts())
	currentSyncedBankAccounts := set.From[*apiv1.BankAccount](syncedPlaidLinkedBankAccounts)

	// we create a hashset of plaid account ids ofr the synced accounts
	// this is used to mark accounts stored in our database as inactive
	// if they are not present in the synced accounts
	currentSyncedBankAccountIdsHashSet := set.New[string](0)
	currentSyncedBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		currentSyncedBankAccountIdsHashSet.Insert(account.PlaidAccountId)
		return true
	})

	// iterate over the current stored bank accounts and mark them as inactive
	// if they are not present in the synced accounts
	inactiveAccountIdsHashSet := set.New[string](0)
	currentBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		if !currentSyncedBankAccountIdsHashSet.Contains(account.PlaidAccountId) {
			th.logger.Info("marking bank account as inactive", zap.Any("account", account))
			account.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE

			// perform the bank account update operation
			if err := th.postgresDb.UpdateBankAccount(ctx, account); err != nil {
				th.logger.Error("failed to update bank account", zap.Error(err))
				return false
			}

			inactiveAccountIdsHashSet.Insert(account.PlaidAccountId)
		}
		return true
	})

	// Check if all accounts are inactive
	if inactiveAccountIdsHashSet.Size() == currentBankAccounts.Size() && currentBankAccounts.Size() > 0 {
		th.logger.Warn("none of the linked bank accounts are active at plaid", zap.Uint64("link_id", link.Id))
		th.logger.Warn("please link a new bank account to continue using the service or update the account link", zap.Any("bank accounts", currentBankAccounts))
		return nil
	}

	// now we need to perform a cross reference to determine if any accounts need to be added or updated
	// we create a hashset of plaid account ids ofr the current stored accounts
	// this is used to determine if we need to add or update an account
	currentBankAccountIdsHashSet := set.New[string](0)
	currentBankAccountPlaidIdToBankAccountMap := make(map[string]*apiv1.BankAccount)
	currentBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		currentBankAccountIdsHashSet.Insert(account.PlaidAccountId)
		currentBankAccountPlaidIdToBankAccountMap[account.PlaidAccountId] = account
		return true
	})

	// iterate over the synced bank accounts and determine if we need to add or update an account
	accountsToBeAdded := set.New[*apiv1.BankAccount](0)
	accountsToBeUpdated := set.New[*apiv1.BankAccount](0)

	currentSyncedBankAccounts.ForEach(func(account *apiv1.BankAccount) bool {
		if currentBankAccountIdsHashSet.Contains(account.PlaidAccountId) {
			// we need to update the account
			if storedBankAccount, ok := currentBankAccountPlaidIdToBankAccountMap[account.PlaidAccountId]; ok {
				copyBankAccount(account, storedBankAccount)
				accountsToBeUpdated.Insert(storedBankAccount)
			} else {
				th.logger.Error("failed to find bank account in map. this should not happen", zap.Any("account", account))
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
		if err := th.postgresDb.UpdateBankAccounts(ctx, link, accountsToBeUpdated.Slice()); err != nil {
			th.logger.Error("failed to update bank accounts", zap.Error(err))
			return err
		}
	}

	if accountsToBeAdded.Size() > 0 {
		accountsToBeAdded.ForEach(func(account *apiv1.BankAccount) bool {
			if _, err := th.postgresDb.CreateBankAccount(ctx, link.Id, account); err != nil {
				th.logger.Error("failed to create bank account", zap.Error(err))
				return false
			}

			return true
		})
	}

	return nil
}

// copyBankAccount is a helper method that copies the contents of a source bank account to a destination bank account
func copyBankAccount(src *apiv1.BankAccount, dest *apiv1.BankAccount) {
	if src == nil || dest == nil {
		return
	}

	dest.Balance = src.Balance
	dest.Currency = src.Currency
	dest.CurrentFunds = src.CurrentFunds
	dest.BalanceLimit = src.BalanceLimit
}
