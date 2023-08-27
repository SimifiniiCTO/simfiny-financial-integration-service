package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SimifiniiCTO/asynq"
	encryptdecrypt "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/encrypt_decrypt"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SyncAllPlatformConnectedPlaidAccountsPayload struct{}

func (t *SyncAllPlatformConnectedPlaidAccountsPayload) String() *string {
	str := "SyncAllPlatformConnectedPlaidAccountsPayload"
	return &str
}

func NewSyncAllPlatformConnectedPlaidAccounts() (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidTaskPayload{})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskSyncAllAccounts.String(), payload), nil
}

// RunSyncAllPlatformConnectedPlaidAccounts is a method on the TaskHandler struct that takes a context and a task
// as parameters. It is used to sync all accounts for all platform users. It does this by querying the database for all
// connected accounts, decrypting the access token, and calling the sync operation for each linked account a given user has.
func (th *TaskHandler) RunSyncAllPlatformConnectedPlaidAccounts(ctx context.Context, task *asynq.Task) error {
	var (
		payload SyncPlaidTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	// ensure operation finished in time
	if err := th.ExecuteBatchSync(ctx, payload.String()); err != nil {
		return err
	}

	return nil
}

func (th *TaskHandler) ExecuteBatchSync(ctx context.Context, trigger *string) error {
	// query the database for all connected accounts
	profiles, err := th.postgresDb.GetAllUserProfiles(ctx)
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		for _, link := range profile.Link {
			// TODO: might need to be more intelligent about how we perform this as it will take a long time
			// Just monitor this as much as possible
			if link.PlaidLink != nil {
				// decrypt the access token
				accessToken, err := encryptdecrypt.DecryptUserAccessToken(ctx, link.Token, th.postgresDb.Kms, th.logger)
				if err != nil {
					th.logger.Error("failed to decrypt access token", zap.Error(err))
					continue
				}

				th.syncAllAccounts(ctx, profile.UserId, link, accessToken, trigger)
			}
		}
	}

	return nil
}

// syncAllAccounts is a method on the TaskHandler struct that takes a context, a user id, a link id, and an access token
// as parameters. It is used to sync all accounts for a given user. It does this by querying the database for the link,
// querying plaid for all accounts, and then performing a cross reference to determine if any accounts need to be added
// or updated.
func (th *TaskHandler) syncAllAccounts(ctx context.Context, userId uint64, link *apiv1.Link, accessToken *string, trigger *string) {
	var (
		creditAccounts      []*apiv1.CreditAccount      = make([]*apiv1.CreditAccount, 0)
		mortgageAccounts    []*apiv1.MortgageAccount    = make([]*apiv1.MortgageAccount, 0)
		studentloanAccounts []*apiv1.StudentLoanAccount = make([]*apiv1.StudentLoanAccount, 0)
	)
	bankAccounts, err := th.syncBankAccounts(ctx, userId, link, *accessToken)
	if err != nil {
		th.logger.Error(fmt.Sprintf("failed to sync transactions for link %d. bank account sync job failed", link.Id), zap.Error(err))
	}

	liabilityAccountSet, err := th.syncCreditAccounts(ctx, userId, link, *accessToken)
	if err != nil {
		th.logger.Error(fmt.Sprintf("failed to sync transactions for link %d. liability account sync job failed", link.Id), zap.Error(err))
	}

	investmentAccounts, err := th.syncInvestmentAccountsHelper(ctx, userId, *accessToken, link)
	if err != nil {
		th.logger.Error("failed to sync investment accounts .. possibily because no investment account is tied to the following link", zap.Error(err))
	}

	th.logger.Info("extraced investment accounts now updating account balances")
	if liabilityAccountSet != nil {
		creditAccounts = liabilityAccountSet.CrediCardAccounts
		mortgageAccounts = liabilityAccountSet.MortgageLoanAccts
		studentloanAccounts = liabilityAccountSet.StudentLoanAccts
	}

	th.logger.Info("starting the account balance update logic")
	if err := th.UpdateAccountBalancesHelper(ctx, userId, bankAccounts, creditAccounts, mortgageAccounts, studentloanAccounts, investmentAccounts); err != nil {
		th.logger.Error("failed to update account balances", zap.Error(err))
	}

	th.logger.Info("Sucessfully updated account balances")
	if err := th.syncTransactions(ctx, userId, link, *accessToken, *trigger); err != nil {
		th.logger.Error(fmt.Sprintf("failed to sync transactions for link %d. transaction sync failed", link.Id), zap.Error(err))
	}
}

func (th *TaskHandler) UpdateAccountBalancesHelper(
	ctx context.Context,
	userId uint64,
	bankAccounts []*apiv1.BankAccount,
	creditAccounts []*apiv1.CreditAccount,
	mortgageAccounts []*apiv1.MortgageAccount,
	studentLoanAccounts []*apiv1.StudentLoanAccount,
	investmentAccounts []*apiv1.InvestmentAccount) error {
	var (
		clickhouseClient = th.clickhouseDb
	)

	accountBalances := make([]*apiv1.AccountBalanceHistory, 0, len(bankAccounts)+len(creditAccounts)+len(mortgageAccounts)+len(studentLoanAccounts))

	th.logger.Info("populating bank account balance")

	// iterate over the bank accounts and create account balance history records
	for _, account := range bankAccounts {
		accountBalance := &apiv1.AccountBalanceHistory{
			Time:            timestamppb.Now(),
			AccountId:       account.PlaidAccountId,
			IsoCurrencyCode: account.Currency,
			Balance:         float64(account.Balance),
			UserId:          userId,
			Sign:            1,
			Id:              "",
		}

		accountBalances = append(accountBalances, accountBalance)
	}

	th.logger.Info("populating credit account balance")

	// iterate over the credit accounts and create account balance history records
	for _, account := range creditAccounts {
		accountBalance := &apiv1.AccountBalanceHistory{
			AccountId: account.PlaidAccountId,
			Balance:   float64(account.Balance),
			UserId:    userId,
			Sign:      1,
			Id:        "",
			Time:      timestamppb.Now(),
		}

		accountBalances = append(accountBalances, accountBalance)
	}

	th.logger.Info("populating mortgage account balance")

	// iterate over the mortgage accounts and create account balance history records
	for _, account := range mortgageAccounts {
		accountBalance := &apiv1.AccountBalanceHistory{
			AccountId: account.PlaidAccountId,
			Balance:   float64(account.OutstandingPrincipalBalance),
			UserId:    userId,
			Sign:      1,
			Id:        "",
			Time:      timestamppb.Now(),
		}

		accountBalances = append(accountBalances, accountBalance)
	}

	th.logger.Info("populating student account balance")

	// iterate over the student loan accounts and create account balance history records
	for _, account := range studentLoanAccounts {
		accountBalance := &apiv1.AccountBalanceHistory{
			AccountId: account.PlaidAccountId,
			Balance:   float64(account.GetOriginationPrincipalAmount()),
			UserId:    userId,
			Sign:      1,
			Id:        "",
			Time:      timestamppb.Now(),
		}

		accountBalances = append(accountBalances, accountBalance)
	}

	// iterate over the investment accounts and create account balance history records
	for _, account := range investmentAccounts {
		accountBalance := &apiv1.AccountBalanceHistory{
			AccountId: account.PlaidAccountId,
			Balance:   float64(account.CurrentFunds),
			UserId:    userId,
			Sign:      1,
			Id:        "",
			Time:      timestamppb.Now(),
		}

		accountBalances = append(accountBalances, accountBalance)
	}

	th.logger.Info("populated all account balances", zap.Int("count", len(accountBalances)))

	// update the account balances in the database
	if len(accountBalances) > 0 {
		th.logger.Info("updating account balances", zap.Int("count", len(accountBalances)))

		if err := clickhouseClient.AddAccountBalances(ctx, &userId, accountBalances); err != nil {
			th.logger.Error("failed to add account balances", zap.Error(err))
			return err
		}
	}

	return nil
}
