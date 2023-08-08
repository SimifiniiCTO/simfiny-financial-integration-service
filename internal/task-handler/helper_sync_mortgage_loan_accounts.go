package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"go.uber.org/zap"
)

// synchronizePlaidLinkedMortgageAccounts is a helper function that is used to sync mortgage loan accounts.
// It takes in two slices of mortgage loan accounts accounts, one that has been synced and one that is currently
// in the database. It then compares the two slices and updates the database accordingly.
// It returns an error if there is an error.
func (th *TaskHandler) synchronizePlaidLinkedMortgageAccounts(ctx context.Context, link *apiv1.Link, syncedPlaidLinkedMortgageLoanAccounts []*apiv1.MortgageAccount) error {
	// return immediately if there are no synced bank accounts
	if len(syncedPlaidLinkedMortgageLoanAccounts) == 0 {
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
	currentMortgageLoanAccounts := set.From[*apiv1.MortgageAccount](link.GetMortgageAccounts())
	currentSyncedMortgageLoanAccounts := set.From[*apiv1.MortgageAccount](syncedPlaidLinkedMortgageLoanAccounts)

	// we create a hashset of plaid account ids ofr the synced accounts
	// this is used to mark accounts stored in our database as inactive
	// if they are not present in the synced accounts
	currentSyncedMortgageLoanAccountIdsHashSet := set.New[string](0)
	currentSyncedMortgageLoanAccounts.ForEach(func(account *apiv1.MortgageAccount) bool {
		currentSyncedMortgageLoanAccountIdsHashSet.Insert(account.PlaidAccountId)
		return true
	})

	// iterate over the current stored bank accounts and mark them as inactive
	// if they are not present in the synced accounts
	inactiveAccountIdsHashSet := set.New[string](0)
	currentMortgageLoanAccounts.ForEach(func(account *apiv1.MortgageAccount) bool {
		if !currentSyncedMortgageLoanAccountIdsHashSet.Contains(account.PlaidAccountId) {
			th.logger.Info("marking credit account as inactive", zap.Any("account", account))
			account.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE

			record, err := account.ToORM(ctx)
			if err != nil {
				th.logger.Error("failed to update credit account", zap.Error(err))
				return false
			}

			if err := l.MortgageAccounts.Model(&linkOrm).Replace(&record); err != nil {
				th.logger.Error("failed to update credit account", zap.Error(err))
				return false
			}

			inactiveAccountIdsHashSet.Insert(account.PlaidAccountId)
		}
		return true
	})

	// Check if all accounts are inactive
	if inactiveAccountIdsHashSet.Size() == currentMortgageLoanAccounts.Size() && currentMortgageLoanAccounts.Size() > 0 {
		th.logger.Warn("none of the linked credit accounts are active at plaid", zap.Uint64("link_id", link.Id))
		th.logger.Warn("please link a new credit account to continue using the service or update the account link", zap.Any("credit accounts", currentMortgageLoanAccounts))
		return nil
	}

	// now we need to perform a cross reference to determine if any accounts need to be added or updated
	// we create a hashset of plaid account ids ofr the current stored accounts
	// this is used to determine if we need to add or update an account
	currentMortgageLoanAccountIdsHashSet := set.New[string](0)
	currentMortgageLoanAccountPlaidIdToMortgageLoanAccountMap := make(map[string]*apiv1.MortgageAccount)
	currentMortgageLoanAccounts.ForEach(func(account *apiv1.MortgageAccount) bool {
		currentMortgageLoanAccountIdsHashSet.Insert(account.PlaidAccountId)
		currentMortgageLoanAccountPlaidIdToMortgageLoanAccountMap[account.PlaidAccountId] = account
		return true
	})

	// iterate over the synced credit accounts and determine if we need to add or update an account
	accountsToBeAdded := set.New[*apiv1.MortgageAccount](0)
	accountsToBeUpdated := set.New[*apiv1.MortgageAccount](0)

	currentSyncedMortgageLoanAccounts.ForEach(func(account *apiv1.MortgageAccount) bool {
		if currentMortgageLoanAccountIdsHashSet.Contains(account.PlaidAccountId) {
			// we need to update the account
			if storedMortgageLoanAccount, ok := currentMortgageLoanAccountPlaidIdToMortgageLoanAccountMap[account.PlaidAccountId]; ok {
				copyMortgageLoanAccount(account, storedMortgageLoanAccount)
				accountsToBeUpdated.Insert(storedMortgageLoanAccount)
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
		if err := th.postgresDb.UpdateMortgageLoanAccounts(ctx, accountsToBeUpdated.Slice()); err != nil {
			th.logger.Error("failed to update credit accounts", zap.Error(err))
			return err
		}
	}

	if accountsToBeAdded.Size() > 0 {
		if accountsToBeUpdated.Size() > 0 {
			th.logger.Info("should not be adding new accounts")
		} else {
			accountsToBeAdded.ForEach(func(account *apiv1.MortgageAccount) bool {
				if _, err := th.postgresDb.CreateMortgageLoantAccount(ctx, link.Id, account); err != nil {
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
func copyMortgageLoanAccount(src *apiv1.MortgageAccount, dest *apiv1.MortgageAccount) {
	if src == nil || dest == nil {
		return
	}

	dest.PlaidAccountId = src.PlaidAccountId
	dest.AccountNumber = src.AccountNumber
	dest.CurrentLateFee = src.CurrentLateFee
	dest.EscrowBalance = src.EscrowBalance
	dest.HasPmi = src.HasPmi
	dest.HasPrepaymentPenalty = src.HasPrepaymentPenalty
	dest.LastPaymentAmount = src.LastPaymentAmount
	dest.LastPaymentDate = src.LastPaymentDate

	dest.LoanTerm = src.LoanTerm
	dest.LoanTypeDescription = src.LoanTypeDescription
	dest.MaturityDate = src.MaturityDate
	dest.NextMonthlyPayment = src.NextMonthlyPayment
	dest.NextPaymentDueDate = src.NextPaymentDueDate
	dest.OriginalPrincipalBalance = src.OriginalPrincipalBalance
	dest.OriginalPropertyValue = src.OriginalPropertyValue
	dest.OutstandingPrincipalBalance = src.OutstandingPrincipalBalance
	dest.PaymentAmount = src.PaymentAmount
	dest.PaymentDate = src.PaymentDate
	dest.OriginationDate = src.OriginationDate
	dest.OriginationPrincipalAmount = src.OriginationPrincipalAmount
	dest.PastDueAmount = src.PastDueAmount
	dest.YtdInterestPaid = src.YtdInterestPaid
	dest.YtdPrincipalPaid = src.YtdPrincipalPaid
	dest.PropertyAddressCity = src.PropertyAddressCity
	dest.PropertyAddressState = src.PropertyAddressState
	dest.PropertyAddressStreet = src.PropertyAddressStreet
	dest.PropertyAddressPostalCode = src.PropertyAddressPostalCode

	dest.PropertyRegion = src.PropertyRegion
	dest.PropertyCountry = src.PropertyCountry
	dest.InterestRatePercentage = src.InterestRatePercentage
	dest.InterestRateType = src.InterestRateType
	dest.Status = src.Status
}
