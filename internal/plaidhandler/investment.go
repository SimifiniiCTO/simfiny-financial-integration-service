package plaidhandler

import (
	"context"
	"fmt"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/plaid/plaid-go/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler/transform"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
)

// TODO: severely refactor this
func (p *PlaidWrapper) GetInvestmentAccount(ctx context.Context, userID uint64, accessToken string) ([]*schema.InvestmentAccount, error) {
	request := plaid.NewInvestmentsHoldingsGetRequest(accessToken)
	options := plaid.NewInvestmentHoldingsGetRequestOptions()
	request.SetOptions(*options)

	resp, _, err := p.client.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	// populate a hashmap comprised of all the account per account id
	accountIDToAccountMap := hashmap.New()
	if resp.GetAccounts() != nil {
		for _, account := range resp.GetAccounts() {
			// populate hashmap with key as the accountID (plaid acct id) and the value a hashset of investment holding structs
			if _, ok := accountIDToAccountMap.Get(account.GetAccountId()); !ok {
				accountIDToAccountMap.Put(account.GetAccountId(), transform.NewInvestmentAccount(userID, &account))
			}
			// NOTE: this should be duplicate free by this point
		}
	}

	// populate a hashmap comprised of all the holdings per account id
	accountIDToInvesmentHoldingMap := hashmap.New()
	securityIDToInvestmentHoldingMap := hashmap.New()
	if resp.GetHoldings() != nil {
		for _, holding := range resp.GetHoldings() {
			holdingRecord := transform.NewInvestmentHolding(&holding)
			// populate hashmap with key as the accountID (plaid acct id) and the value a hashset of investment holding structs
			if _, ok := accountIDToInvesmentHoldingMap.Get(holding.AccountId); !ok {
				v := hashset.New()
				v.Add(holdingRecord)
				accountIDToInvesmentHoldingMap.Put(holding.AccountId, v)
			} else {
				res, _ := accountIDToInvesmentHoldingMap.Get(holding.AccountId)
				value := res.(hashset.Set)
				value.Add(holdingRecord)
				accountIDToInvesmentHoldingMap.Put(holding.AccountId, value)
			}

			// securities and holding retain a one to one mapping to one another
			if _, ok := securityIDToInvestmentHoldingMap.Get(holding.SecurityId); !ok {
				securityIDToInvestmentHoldingMap.Put(holding.SecurityId, holdingRecord)
			}
		}
	}

	result := make([]*schema.InvestmentAccount, 0, len(resp.GetAccounts()))
	// iterate over the acct id to acct map and for each account, populate the holdings and securities fields
	for _, accountID := range accountIDToAccountMap.Keys() {
		res, _ := accountIDToAccountMap.Get(accountID)
		account := res.(*schema.InvestmentAccount)
		// populate the holdings
		if holdings, ok := accountIDToInvesmentHoldingMap.Get(accountID); ok {
			value, ok := holdings.(hashset.Set)
			if !ok {
				return nil, fmt.Errorf("could not convert value to hashset.Set")
			}
			data, err := pointer.ToSliceOf[*schema.InvesmentHolding](value.Values())
			if err != nil {
				return nil, err
			}

			account.Holdings = data

			// get all the securities tied to this holding for this given account
			for _, holding := range data {
				if security, ok := securityIDToInvestmentHoldingMap.Get(holding.SecurityId); ok {
					data, ok := security.(*schema.InvestmentSecurity)
					if !ok {
						return nil, fmt.Errorf("could not convert value to *schema.InvestmentSecurity")
					}
					account.Securities = append(account.Securities, data)
				}
			}
		}

		result = append(result, account)
	}

	return result, nil
}
