package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	"github.com/pkg/errors"
)

type NewPullInvestmentHoldingsTaskPayload struct {
	UserId      uint64   `json:"user_id"`
	LinkId      uint64   `json:"link_id"`
	AccessToken string   `json:"access_token"`
	AccountIds  []string `json:"account_ids"`
}

func NewPullInvestmentHoldingsTask(userId, linkId uint64, accessToken string, accountIds []string) (*asynq.Task, error) {
	if userId == 0 {
		return nil, errors.New("invalid input argument. user id cannot be empty")
	}

	if linkId == 0 {
		return nil, errors.New("invalid input argument. link id cannot be empty")
	}

	if accessToken == "" {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	rec := &NewPullInvestmentHoldingsTaskPayload{
		UserId:      userId,
		LinkId:      linkId,
		AccessToken: accessToken,
		AccountIds:  accountIds,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskPullInvestmentHoldings.String(), payload), nil
}

func (th *TaskHandler) RunPullInvestmentHoldingsTask(ctx context.Context, t *asynq.Task) error {
	var (
		payload        = &NewPullInvestmentHoldingsTaskPayload{}
		postgresClient = th.postgresDb
		plaidClient    = th.plaidClient
	)

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	accessToken := payload.AccessToken
	accountIds := payload.AccountIds
	linkId := payload.LinkId
	userId := payload.UserId

	// get investment holdings
	investments, err := plaidClient.GetPlaidInvestmentHoldings(ctx, &accessToken, accountIds)
	if err != nil {
		return err
	}

	// get all investment holdings from the database
	link, err := postgresClient.GetLink(ctx, userId, linkId)
	if err != nil {
		return err
	}

	// get all the current investment holdings
	currentHoldings := make([]*schema.InvesmentHolding, 0)
	for _, account := range link.InvestmentAccounts {
		currentHoldings = append(currentHoldings, account.Holdings...)
	}

	// get all the new investment holdings
	newInvestmentHoldings := make([]*schema.InvesmentHolding, 0)
	updatedInvestmentHoldings := make([]*schema.InvesmentHolding, 0)
	deletedInvestmentHoldings := make([]uint64, 0)
	plaidAccountIdToInvestmentAccountId := make(map[string]uint64, 0)

	for _, investmentAccount := range link.InvestmentAccounts {
		if isInAccountIdSlice(investmentAccount.PlaidAccountId, accountIds) {
			plaidAccountIdToInvestmentAccountId[investmentAccount.PlaidAccountId] = investmentAccount.Id
		}
	}

	// get all the deleted investment holdings
	for _, currentHolding := range currentHoldings {
		// we check if the holding (updated or new) obtained from plaid matches any of the existing holdings
		// we hold
		for idx, potentiallyNewInvestmentHolding := range investments.Holdings {
			if potentiallyNewInvestmentHolding.SecurityId == currentHolding.SecurityId {
				break
			} else if idx == len(investments.Holdings)-1 {
				// this is a deleted holding
				deletedInvestmentHoldings = append(deletedInvestmentHoldings, currentHolding.Id)
			}
		}
	}

	// delete investment holdings that should be deleted from the database
	if err := postgresClient.DeleteInvestmentHoldings(ctx, deletedInvestmentHoldings...); err != nil {
		return err
	}

	for _, potentiallyNewInvestmentHolding := range investments.Holdings {
		// we check if the holding (updated or new) obtained from plaid matches any of the existing holdings
		// we hold
		for idx, currentHolding := range currentHoldings {
			if potentiallyNewInvestmentHolding.SecurityId == currentHolding.SecurityId {
				// this is an updated holding so we update the current holding
				updatedHolding := transformer.NewInvestmentHolding(&potentiallyNewInvestmentHolding)

				// assign the id and plaid account id to the updated holding
				updatedHolding.Id = currentHolding.Id
				updatedHolding.PlaidAccountId = currentHolding.PlaidAccountId

				updatedInvestmentHoldings = append(updatedInvestmentHoldings, transformer.NewInvestmentHolding(&potentiallyNewInvestmentHolding))
				break
			} else if idx == len(currentHoldings)-1 {
				// this is a new holding
				newInvestmentHoldings = append(newInvestmentHoldings, transformer.NewInvestmentHolding(&potentiallyNewInvestmentHolding))
			}
		}
	}

	// iterate over all the new investment holdings and populate a hashmap of investment account id to investment holding list
	newInvestmentAccountIdToInvestmentHolding := make(map[uint64][]*schema.InvesmentHolding, 0)
	for _, newInvestmentHolding := range newInvestmentHoldings {
		newInvestmentHoldingAcctId := newInvestmentHolding.PlaidAccountId
		// get the investment account this new holding should be tied to
		investmentAccountId := plaidAccountIdToInvestmentAccountId[newInvestmentHoldingAcctId]
		if _, ok := newInvestmentAccountIdToInvestmentHolding[investmentAccountId]; !ok {
			newInvestmentAccountIdToInvestmentHolding[investmentAccountId] = make([]*schema.InvesmentHolding, 0)
		}

		newInvestmentAccountIdToInvestmentHolding[investmentAccountId] = append(newInvestmentAccountIdToInvestmentHolding[investmentAccountId], newInvestmentHolding)
	}

	// now iterate over hashmap keys and add the new investment holding to the database
	for investmentAccountId, investmentHoldings := range newInvestmentAccountIdToInvestmentHolding {
		if _, err := postgresClient.AddInvestmentHoldings(ctx, &userId, &investmentAccountId, investmentHoldings); err != nil {
			return err
		}
	}

	// we do the same thing for the updated set of investment holdings
	updatedInvestmentAccountIdToInvestmentHolding := make(map[uint64][]*schema.InvesmentHolding, 0)
	for _, updatedInvestmentHolding := range updatedInvestmentHoldings {
		updatedInvestmentHoldingAcctId := updatedInvestmentHolding.PlaidAccountId
		// get the investment account this new holding should be tied to
		investmentAccountId := plaidAccountIdToInvestmentAccountId[updatedInvestmentHoldingAcctId]
		if _, ok := updatedInvestmentAccountIdToInvestmentHolding[investmentAccountId]; !ok {
			updatedInvestmentAccountIdToInvestmentHolding[investmentAccountId] = make([]*schema.InvesmentHolding, 0)
		}

		updatedInvestmentAccountIdToInvestmentHolding[investmentAccountId] = append(updatedInvestmentAccountIdToInvestmentHolding[investmentAccountId], updatedInvestmentHolding)
	}

	// now iterate over hashmap keys and add the new investment holding to the database
	for investmentAccountId, investmentHoldings := range updatedInvestmentAccountIdToInvestmentHolding {
		if _, err := postgresClient.UpdateInvestmentHoldings(ctx, &userId, &investmentAccountId, investmentHoldings); err != nil {
			return err
		}
	}

	// update the securities
	updatedInvestmentSecurities, err := transformer.TransformPlaidInvestmentSecurity(investments.Securities)
	if err != nil {
		return err
	}

	// partition across plaid account id
	investmentAccountIdToSecuritiesHashmap := make(map[uint64][]*schema.InvestmentSecurity, 0)
	// iterate over the holdings and associate the account id to the proper investement
	for _, holding := range investments.Holdings {
		for _, security := range updatedInvestmentSecurities {
			if holding.SecurityId == security.SecurityId {
				if investmentAccountId, ok := plaidAccountIdToInvestmentAccountId[holding.AccountId]; ok {
					if _, ok := investmentAccountIdToSecuritiesHashmap[investmentAccountId]; !ok {
						investmentAccountIdToSecuritiesHashmap[investmentAccountId] = make([]*schema.InvestmentSecurity, 0)
					}

					investmentAccountIdToSecuritiesHashmap[investmentAccountId] = append(investmentAccountIdToSecuritiesHashmap[investmentAccountId], security)
				}

			}
		}
	}

	// update the investment securities too
	for investmentAccountId, securities := range investmentAccountIdToSecuritiesHashmap {
		if err := postgresClient.UpdateSecurities(ctx, &investmentAccountId, securities); err != nil {
			return err
		}
	}

	return nil
}

func isInAccountIdSlice(accountId string, accountIds []string) bool {
	for _, acctId := range accountIds {
		if acctId == accountId {
			return true
		}
	}
	return false
}
