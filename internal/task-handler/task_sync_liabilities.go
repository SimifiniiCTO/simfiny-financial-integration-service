package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SimifiniiCTO/asynq"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
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

// processAndStoreInvestmentAccount is a helper function that will process the synced investment accounts and store them
// in the database. This function will also update the existing investment accounts if they have changed.
func (th *TaskHandler) processAndStoreInvestmentAccount(ctx context.Context, link *apiv1.Link, investmentAccts []*apiv1.InvestmentAccount) error {
	if len(investmentAccts) > 0 {
		// we ensure we have a unique set of investment accounts to account for the newly synced investment accounts
		syncedInvestmentAccounts := set.New[*apiv1.InvestmentAccount](len(investmentAccts))
		syncedInvestmentAccounts.InsertSlice(investmentAccts)

		// we ensure we have a unique set of investment accounts to account for the existing investment accounts
		currentInvestmentAccounts := set.New[*apiv1.InvestmentAccount](len(link.InvestmentAccounts))
		currentInvestmentAccounts.InsertSlice(link.InvestmentAccounts)

		th.logger.Info("current investment accounts", zap.Any("accounts", currentInvestmentAccounts.Size()))
		th.logger.Info("synced investment accounts", zap.Any("accounts", syncedInvestmentAccounts.Size()))

		// create a hasmap of plaid account id and the investment account stored in the database
		// this will be used to cross reference the synced investment accounts with the existing investment accounts
		// to determine which investment accounts need to be updated and which ones need to be added
		plaidAccountIdToInvestmentAccountHashMap := make(map[string]*apiv1.InvestmentAccount, currentInvestmentAccounts.Size())
		for _, investmentAccount := range currentInvestmentAccounts.Slice() {
			plaidAccountIdToInvestmentAccountHashMap[investmentAccount.PlaidAccountId] = investmentAccount
		}

		// iterate over the syncred investment accounts and cross reference them with the existing investment accounts
		// if the synced investment account is not found in the existing investment accounts, then add it to the existing
		// investment accounts
		accountsToUpdate := set.New[*apiv1.InvestmentAccountORM](syncedInvestmentAccounts.Size())
		accountsToAdd := set.New[*apiv1.InvestmentAccountORM](syncedInvestmentAccounts.Size())

		syncedInvestmentAccounts.ForEach(func(item *apiv1.InvestmentAccount) bool {
			if currentInvestmentAccounts.Size() == 0 {
				record, err := item.ToORM(ctx)
				if err != nil {
					th.logger.Error("failed to convert investment account to orm", zap.Error(err))
					return false
				}
				accountsToAdd.Insert(&record)
				return true
			}

			if _, ok := plaidAccountIdToInvestmentAccountHashMap[item.PlaidAccountId]; ok {
				record, err := item.ToORM(ctx)
				if err != nil {
					th.logger.Error("failed to convert investment account to orm", zap.Error(err))
					return false
				}
				accountsToUpdate.Insert(&record)
				return true
			}

			record, err := item.ToORM(ctx)
			if err != nil {
				th.logger.Error("failed to convert investment account to orm", zap.Error(err))
				return false
			}

			accountsToAdd.Insert(&record)
			return true
		})

		l := th.postgresDb.QueryOperator.LinkORM
		linkOrm, err := link.ToORM(ctx)
		if err != nil {
			return err
		}

		th.logger.Info("updated investment accounts", zap.Any("accounts", accountsToUpdate))

		// update the investment accounts in the database
		if accountsToUpdate.Size() > 0 {
			if err := l.InvestmentAccounts.Model(&linkOrm).Replace(accountsToUpdate.Slice()...); err != nil {
				return err
			}
		}

		th.logger.Info("new investment accounts", zap.Any("accounts", accountsToAdd))

		// add the new investment accounts to the database
		if accountsToAdd.Size() > 0 {
			if err := l.InvestmentAccounts.Model(&linkOrm).Append(accountsToAdd.Slice()...); err != nil {
				return err
			}
		}
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
		if err := th.processAndStoreInvestmentAccount(ctx, link, investmentAccounts); err != nil {
			return nil, err
		}
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
		if len(creditCardAccountsToAdd) > 0 {
			if err := l.CreditAccounts.Model(&linkOrm).Append(creditCardAccountsToAdd...); err != nil {
				return nil, err
			}
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
