package taskhandler

import (
	"context"

	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"go.uber.org/zap"
)

func (th *TaskHandler) syncInvestmentAccountsHelper(
	ctx context.Context,
	userId uint64,
	accessToken string,
	link *apiv1.Link) ([]*apiv1.InvestmentAccount, error) {
	var (
		err         error
		plaidClient = th.plaidClient
	)

	// based on the link, we get all the investment accounts of interest
	currentStoredInvestmentAccounts := link.InvestmentAccounts

	// using the plaid api, extract the investment accounts tied to this access token
	newSyncedInvestmentAccounts, err := plaidClient.GetInvestmentAccount(ctx, userId, accessToken)
	if err == nil {
		if len(newSyncedInvestmentAccounts) > 0 {
			th.logger.Info("found investment accounts", zap.Int("count", len(newSyncedInvestmentAccounts)))

			// store the investment accounts in the database (only the updated ones)
			if err := th.processAndStoreInvestmentAccount(ctx, link, newSyncedInvestmentAccounts); err != nil {
				return nil, err
			}
		}
	} else {
		th.logger.Error("failed to get investment accounts", zap.Error(err))
	}

	// ensure we can extract all the investment account ids which
	// we will be using to extract the investment transactions
	th.logger.Info("extracting investment account ids")
	accountIds := set.New[string](len(currentStoredInvestmentAccounts))
	for _, investmentAccount := range currentStoredInvestmentAccounts {
		th.logger.Info("adding investment account id", zap.String("account_id", investmentAccount.PlaidAccountId))
		accountIds.Insert(investmentAccount.PlaidAccountId)
	}

	th.logger.Info("extracted investment account ids", zap.Int("count", accountIds.Size()))
	if accountIds.Size() > 0 {
		th.logger.Info("extracting investment transactions")
		// extract the investment transactions for the investment accounts
		if err := th.queryAndUpdateInvestmentTransactions(ctx, accessToken, userId, link, accountIds.Slice()); err != nil {
			th.logger.Error("failed to get investment account transactions", zap.Error(err))
		}
	}

	return newSyncedInvestmentAccounts, nil
}
