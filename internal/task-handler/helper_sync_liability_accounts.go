package taskhandler

import (
	"context"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// syncBankAccounts is a method on the TaskHandler struct that takes a context, a user id, a link id, and an access token
// as parameters. It is used to sync all accounts for a given user. It does this by querying the database for the link,
// querying plaid for all accounts, and then performing a cross reference to determine if any accounts need to be added
// or updated.
func (th *TaskHandler) syncCreditAccounts(ctx context.Context, userId uint64, link *apiv1.Link, accessToken string) (*plaidhandler.CreditAccountSet, error) {
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

	// query plaid for all of the user's liability accounts
	creditAccountsSet, err := plaidClient.GetPlaidLiabilityAccounts(ctx, &accessToken)
	if err != nil {
		th.logger.Error("failed to get credit accounts", zap.Error(err))
	} else {
		th.logger.Info("successfully retrieved credit accounts", zap.Any("accounts", creditAccountsSet))
		creditAccounts, mortgageAccounts, studentloanAccounts := creditAccountsSet.CrediCardAccounts, creditAccountsSet.MortgageLoanAccts, creditAccountsSet.StudentLoanAccts
		if len(creditAccounts) > 0 {
			if err := th.synchronizePlaidLinkedCreditAccounts(ctx, link, creditAccountsSet.CrediCardAccounts); err != nil {
				th.logger.Error("failed to sync credit accounts", zap.Error(err))
			}
		}

		if len(studentloanAccounts) > 0 {
			if err := th.synchronizePlaidLinkedStudentAccounts(ctx, link, creditAccountsSet.StudentLoanAccts); err != nil {
				th.logger.Error("failed to sync student loan accounts", zap.Error(err))
			}
		}

		if len(mortgageAccounts) > 0 {
			if err := th.synchronizePlaidLinkedMortgageAccounts(ctx, link, creditAccountsSet.MortgageLoanAccts); err != nil {
				th.logger.Error("failed to sync mortgage loan accounts", zap.Error(err))
			}
		}
	}

	return creditAccountsSet, nil
}
