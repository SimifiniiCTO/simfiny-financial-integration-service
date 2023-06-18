package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SimifiniiCTO/asynq"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SyncPlaidLiabilityTaskPayload struct {
	UserId      uint64   `json:"user_id"`
	AccessToken string   `json:"access_token"`
	LinkId      uint64   `json:"plaid_link_item_id"`
	AccountIds  []string `json:"account_ids"`
}

func (t *SyncPlaidLiabilityTaskPayload) String() *string {
	str := fmt.Sprintf("sync_plaid_liability_account user_id: %d, access_token: %s, plaid_link_item_id: %d", t.UserId, t.AccessToken, t.LinkId)
	return &str
}

func NewSyncNewLiabilityAccountsTask(userId uint64, accessToken string, linkId uint64, accountIds []string) (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidLiabilityTaskPayload{
		UserId:      userId,
		AccessToken: accessToken,
		LinkId:      linkId,
		AccountIds:  accountIds,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskSyncNewLiabilityAccounts.String(), payload), nil
}

func (th *TaskHandler) RunSyncNewLiabilityAccountsTask(ctx context.Context, task *asynq.Task) error {
	var (
		payload SyncPlaidLiabilityTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	accts, err := th.processLiabilityAccountSyncOperation(ctx, &payload)
	if err != nil {
		return err
	}

	if len(accts) == 0 {
		th.logger.Warn("no accounts found for user", zap.Uint64("user_id", payload.UserId))
		return nil
	}

	return nil
}

func (th *TaskHandler) processLiabilityAccountSyncOperation(ctx context.Context, payload *SyncPlaidLiabilityTaskPayload) ([]*apiv1.BankAccount, error) {
	var (
		postgresClient = th.postgresDb
		plaidClient    = th.plaidClient
	)

	// query the link from the database
	link, err := postgresClient.GetLink(ctx, payload.UserId, payload.LinkId, false)
	if err != nil {
		return nil, err
	}

	if link.PlaidLink == nil {
		th.logger.Warn("provided link does not have any plaid credentials")
		return nil, errors.New("plaid link not found")
	}

	// query plaid for all of the user's accounts
	plaidBankAccounts, err := plaidClient.GetAccounts(ctx, payload.AccessToken, payload.UserId)
	if err != nil {
		return nil, err
	}

	// get the liability accounts
	// TODO: need to figure out wether to update the existing investment account or add a new one
	investmentAccounts, err := plaidClient.GetInvestmentAccount(ctx, payload.UserId, payload.AccessToken)
	if err != nil {
		return nil, err
	}

	if len(investmentAccounts) > 0 {
		th.logger.Info("found investment accounts", zap.Int("count", len(investmentAccounts)))
	}

	// get the credit accounts
	// TODO: need to figure out wether to update the existing credit account or add a new one
	creditAccountsSet, err := plaidClient.GetPlaidLiabilityAccounts(ctx, &payload.AccessToken, payload.AccountIds...)
	if err != nil {
		return nil, err
	}

	if len(creditAccountsSet.CrediCardAccounts) > 0 {
		syncedCreditAccounts := creditAccountsSet.CrediCardAccounts
		currentCreditAccounts := link.GetCreditAccounts()

		// iterate over the syncred credit accounts and cross reference them with the existing credit accounts
		// if the synced credit account is not found in the existing credit accounts, then add it to the existing
		// credit accounts
		creditCardAccountsToUpdate := make([]*apiv1.CreditAccountORM, 0, len(syncedCreditAccounts))
		creditCardAccountsToAdd := make([]*apiv1.CreditAccountORM, 0, len(syncedCreditAccounts))
		for _, syncedCreditAccount := range syncedCreditAccounts {
			found := false
			for idx, currentCreditAccount := range currentCreditAccounts {
				if syncedCreditAccount.PlaidAccountId == currentCreditAccount.PlaidAccountId {
					found = true
					break
				} else if syncedCreditAccount.PlaidAccountId != currentCreditAccount.PlaidAccountId && idx == len(currentCreditAccounts)-1 {
					acctOrm, err := syncedCreditAccount.ToORM(ctx)
					if err != nil {
						return nil, err
					}

					creditCardAccountsToAdd = append(creditCardAccountsToAdd, &acctOrm)
				}
			}

			if found {
				acctOrm, err := syncedCreditAccount.ToORM(ctx)
				if err != nil {
					return nil, err
				}
				creditCardAccountsToUpdate = append(creditCardAccountsToUpdate, &acctOrm)
			}
		}

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		// update the credit accounts in the database
		if err := l.CreditAccounts.Model(&linkOrm).Replace(creditCardAccountsToUpdate...); err != nil {
			return nil, err
		}

		// add the new credit accounts to the database
		if err := l.CreditAccounts.Model(&linkOrm).Append(creditCardAccountsToAdd...); err != nil {
			return nil, err
		}

		th.logger.Info("found credit accounts", zap.Int("count", len(creditAccountsSet.CrediCardAccounts)))
	}

	if len(creditAccountsSet.MortgageLoanAccts) > 0 {
		syncedMortgageLoanAccounts := creditAccountsSet.MortgageLoanAccts
		currentMortgageLoanAccounts := link.GetMortgageAccounts()

		// iterate over the syncred credit accounts and cross reference them with the existing credit accounts
		// if the synced credit account is not found in the existing credit accounts, then add it to the existing
		// credit accounts
		mortgageLoanAccountsToUpdate := make([]*apiv1.MortgageAccountORM, 0, len(syncedMortgageLoanAccounts))
		mortgageLoanAccountsToAdd := make([]*apiv1.MortgageAccountORM, 0, len(syncedMortgageLoanAccounts))
		for _, syncedCreditAccount := range syncedMortgageLoanAccounts {
			found := false
			for idx, currentCreditAccount := range currentMortgageLoanAccounts {
				if syncedCreditAccount.PlaidAccountId == currentCreditAccount.PlaidAccountId {
					found = true
					break
				} else if syncedCreditAccount.PlaidAccountId != currentCreditAccount.PlaidAccountId && idx == len(currentMortgageLoanAccounts)-1 {
					acctOrm, err := syncedCreditAccount.ToORM(ctx)
					if err != nil {
						return nil, err
					}

					mortgageLoanAccountsToAdd = append(mortgageLoanAccountsToAdd, &acctOrm)
				}
			}

			if found {
				acctOrm, err := syncedCreditAccount.ToORM(ctx)
				if err != nil {
					return nil, err
				}
				mortgageLoanAccountsToUpdate = append(mortgageLoanAccountsToUpdate, &acctOrm)
			}
		}

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		// update the credit accounts in the database
		if err := l.MortgageAccounts.Model(&linkOrm).Replace(mortgageLoanAccountsToUpdate...); err != nil {
			return nil, err
		}

		// add the new credit accounts to the database
		if err := l.MortgageAccounts.Model(&linkOrm).Append(mortgageLoanAccountsToAdd...); err != nil {
			return nil, err
		}

		th.logger.Info("found mortgage accounts", zap.Int("count", len(creditAccountsSet.MortgageLoanAccts)))
	}

	if len(creditAccountsSet.StudentLoanAccts) > 0 {
		syncedStudentLoanAccounts := creditAccountsSet.StudentLoanAccts
		currentStudentLoanAccounts := link.GetStudentLoanAccounts()

		// iterate over the syncred credit accounts and cross reference them with the existing credit accounts
		// if the synced credit account is not found in the existing credit accounts, then add it to the existing
		// credit accounts
		studentLoanAccountsToUpdate := make([]*apiv1.StudentLoanAccountORM, 0, len(syncedStudentLoanAccounts))
		studentLoanAccountsToAdd := make([]*apiv1.StudentLoanAccountORM, 0, len(syncedStudentLoanAccounts))
		for _, syncedCreditAccount := range syncedStudentLoanAccounts {
			found := false
			for idx, currentCreditAccount := range currentStudentLoanAccounts {
				if syncedCreditAccount.PlaidAccountId == currentCreditAccount.PlaidAccountId {
					found = true
					break
				} else if syncedCreditAccount.PlaidAccountId != currentCreditAccount.PlaidAccountId && idx == len(currentStudentLoanAccounts)-1 {
					acctOrm, err := syncedCreditAccount.ToORM(ctx)
					if err != nil {
						return nil, err
					}

					studentLoanAccountsToAdd = append(studentLoanAccountsToAdd, &acctOrm)
				}
			}

			if found {
				acctOrm, err := syncedCreditAccount.ToORM(ctx)
				if err != nil {
					return nil, err
				}
				studentLoanAccountsToUpdate = append(studentLoanAccountsToUpdate, &acctOrm)
			}
		}

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		// update the credit accounts in the database
		if err := l.StudentLoanAccounts.Model(&linkOrm).Replace(studentLoanAccountsToUpdate...); err != nil {
			return nil, err
		}

		// add the new credit accounts to the database
		if err := l.StudentLoanAccounts.Model(&linkOrm).Append(studentLoanAccountsToAdd...); err != nil {
			return nil, err
		}

		th.logger.Info("found student loan accounts", zap.Int("count", len(creditAccountsSet.StudentLoanAccts)))
	}

	return plaidBankAccounts, nil
}
