package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"go.uber.org/zap"
)

// synchronizePlaidLinkedStudentAccounts is a helper function that is used to sync student loan accounts.
// It takes in two slices of student loan accounts, one that has been synced and one that is currently
// in the database. It then compares the two slices and updates the database accordingly.
// It returns an error if there is an error.
func (th *TaskHandler) synchronizePlaidLinkedStudentAccounts(ctx context.Context, link *apiv1.Link, syncedPlaidLinkedStudentAccounts []*apiv1.StudentLoanAccount) error {
	// return immediately if there are no synced bank accounts
	if len(syncedPlaidLinkedStudentAccounts) == 0 {
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
	currentStudentLoanAccounts := set.From[*apiv1.StudentLoanAccount](link.GetStudentLoanAccounts())
	currentSyncedStudentLoanAccounts := set.From[*apiv1.StudentLoanAccount](syncedPlaidLinkedStudentAccounts)

	// we create a hashset of plaid account ids ofr the synced accounts
	// this is used to mark accounts stored in our database as inactive
	// if they are not present in the synced accounts
	currentSyncedStudentLoanAccountIdsHashSet := set.New[string](0)
	currentSyncedStudentLoanAccounts.ForEach(func(account *apiv1.StudentLoanAccount) bool {
		currentSyncedStudentLoanAccountIdsHashSet.Insert(account.PlaidAccountId)
		return true
	})

	// iterate over the current stored bank accounts and mark them as inactive
	// if they are not present in the synced accounts
	inactiveAccountIdsHashSet := set.New[string](0)
	currentStudentLoanAccounts.ForEach(func(account *apiv1.StudentLoanAccount) bool {
		if !currentSyncedStudentLoanAccountIdsHashSet.Contains(account.PlaidAccountId) {
			th.logger.Info("marking student loan account as inactive", zap.Any("account", account))
			account.Status = apiv1.BankAccountStatus_BANK_ACCOUNT_STATUS_INACTIVE

			record, err := account.ToORM(ctx)
			if err != nil {
				th.logger.Error("failed to update student loan account", zap.Error(err))
				return false
			}

			if err := l.StudentLoanAccounts.Model(&linkOrm).Replace(&record); err != nil {
				th.logger.Error("failed to update student loan account", zap.Error(err))
				return false
			}

			inactiveAccountIdsHashSet.Insert(account.PlaidAccountId)
		}
		return true
	})

	// Check if all accounts are inactive
	if inactiveAccountIdsHashSet.Size() == currentStudentLoanAccounts.Size() && currentStudentLoanAccounts.Size() > 0 {
		th.logger.Warn("none of the linked student loan accounts are active at plaid", zap.Uint64("link_id", link.Id))
		th.logger.Warn("please link a new student loan account to continue using the service or update the account link", zap.Any("student loan accounts", currentStudentLoanAccounts))
		return nil
	}

	// now we need to perform a cross reference to determine if any accounts need to be added or updated
	// we create a hashset of plaid account ids ofr the current stored accounts
	// this is used to determine if we need to add or update an account
	currentStudentLoanAccountIdsHashSet := set.New[string](0)
	currentStudentLoanAccountPlaidIdToStudentLoanAccountMap := make(map[string]*apiv1.StudentLoanAccount)
	currentStudentLoanAccounts.ForEach(func(account *apiv1.StudentLoanAccount) bool {
		currentStudentLoanAccountIdsHashSet.Insert(account.PlaidAccountId)
		currentStudentLoanAccountPlaidIdToStudentLoanAccountMap[account.PlaidAccountId] = account
		return true
	})

	// iterate over the synced credit accounts and determine if we need to add or update an account
	accountsToBeAdded := set.New[*apiv1.StudentLoanAccount](0)
	accountsToBeUpdated := set.New[*apiv1.StudentLoanAccount](0)

	currentSyncedStudentLoanAccounts.ForEach(func(account *apiv1.StudentLoanAccount) bool {
		if currentStudentLoanAccountIdsHashSet.Contains(account.PlaidAccountId) {
			// we need to update the account
			if storedStudentLoanAccount, ok := currentStudentLoanAccountPlaidIdToStudentLoanAccountMap[account.PlaidAccountId]; ok {
				copyStudentLoanAccount(account, storedStudentLoanAccount)
				accountsToBeUpdated.Insert(storedStudentLoanAccount)
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
		if err := th.postgresDb.UpdateStudentLoanAccounts(ctx, accountsToBeUpdated.Slice()); err != nil {
			th.logger.Error("failed to update student loan accounts", zap.Error(err))
			return err
		}
	}

	if accountsToBeAdded.Size() > 0 {
		if accountsToBeUpdated.Size() > 0 {
			th.logger.Info("should not be adding new accounts")
		} else {
			accountsToBeAdded.ForEach(func(account *apiv1.StudentLoanAccount) bool {
				if _, err := th.postgresDb.CreateStudentLoantAccount(ctx, link.Id, account); err != nil {
					th.logger.Error("failed to create student loan account", zap.Error(err))
					return false
				}

				return true
			})
		}

	}

	return nil
}

// copyBankAccount is a helper method that copies the contents of a source bank account to a destination bank account
func copyStudentLoanAccount(src *apiv1.StudentLoanAccount, dest *apiv1.StudentLoanAccount) {
	if src == nil || dest == nil {
		return
	}

	dest.Name = src.Name
	dest.DisbursementDates = src.DisbursementDates
	dest.ExpectedPayoffDate = src.ExpectedPayoffDate
	dest.Guarantor = src.Guarantor
	dest.InterestRatePercentage = src.InterestRatePercentage
	dest.IsOverdue = src.IsOverdue
	dest.LastPaymentAmount = src.LastPaymentAmount
	dest.LastPaymentDate = src.LastPaymentDate

	dest.LastStatementIssueDate = src.LastStatementIssueDate
	dest.LoanName = src.LoanName
	dest.LoanEndDate = src.LoanEndDate
	dest.MinimumPaymentAmount = src.MinimumPaymentAmount
	dest.NextPaymentDueDate = src.NextPaymentDueDate
	dest.OriginationPrincipalAmount = src.OriginationPrincipalAmount
	dest.OutstandingInterestAmount = src.OutstandingInterestAmount
	dest.PaymentReferenceNumber = src.PaymentReferenceNumber
	dest.SequenceNumber = src.SequenceNumber
	dest.YtdInterestPaid = src.YtdInterestPaid
	dest.YtdPrincipalPaid = src.YtdPrincipalPaid
	dest.LoanType = src.LoanType
	dest.PslfStatusEstimatedEligibilityDate = src.PslfStatusEstimatedEligibilityDate
	dest.PslfStatusPaymentsMade = src.PslfStatusPaymentsMade
	dest.PslfStatusPaymentsRemaining = src.PslfStatusPaymentsRemaining
	dest.RepaymentPlanType = src.RepaymentPlanType
	dest.RepaymentPlanDescription = src.RepaymentPlanDescription
	dest.ServicerAddressCity = src.ServicerAddressCity
	dest.ServicerAddressPostalCode = src.ServicerAddressPostalCode
	dest.ServicerAddressState = src.ServicerAddressState
	dest.ServicerAddressStreet = src.ServicerAddressStreet
	dest.ServicerAddressRegion = src.ServicerAddressRegion
	dest.ServicerAddressCountry = src.ServicerAddressCountry
	dest.UserId = src.UserId
	dest.Name = src.Name
	dest.Status = src.Status
}
