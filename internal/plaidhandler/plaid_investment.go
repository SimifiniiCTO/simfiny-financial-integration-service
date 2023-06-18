package plaidhandler

import (
	"context"
	"fmt"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/plaid/plaid-go/v12/plaid"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// GetInvestmentAccount is used to retrieve investment accounts from the Plaid API for a given set of account IDs
// and transform them into the schema format using the `transformer.NewInvestmentAccount` function and returns
// them as an array of `schema.InvestmentAccount` objects.
// TODO: refactor
func (p *PlaidWrapper) GetInvestmentAccount(ctx context.Context, userID uint64, accessToken string) ([]*schema.InvestmentAccount, error) {
	request := plaid.NewInvestmentsHoldingsGetRequest(accessToken)
	options := plaid.NewInvestmentHoldingsGetRequestOptions()
	request.SetOptions(*options)
	request.SetSecret(p.SecretKey)
	request.SetClientId(p.ClientID)

	resp, _, err := p.client.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	// populate a hashmap comprised of all the account per account id
	accountIDToAccountMap := hashmap.New()
	if resp.GetAccounts() != nil {
		for _, account := range resp.GetAccounts() {
			if account.Type == plaid.ACCOUNTTYPE_INVESTMENT {
				// populate hashmap with key as the accountID (plaid acct id) and the value a hashset of investment holding structs
				if _, ok := accountIDToAccountMap.Get(account.GetAccountId()); !ok {
					investmentAcct, err := transformer.NewInvestmentAccount(userID, &account)
					if err != nil {
						return nil, err
					}

					accountIDToAccountMap.Put(account.GetAccountId(), investmentAcct)
				}
			}
		}
	}

	// populate a hashmap comprised of all the holdings per account id
	accountIDToInvesmentHoldingMap := hashmap.New()
	securityIDToInvestmentHoldingMap := hashmap.New()
	securityIDToSecurityMap := hashmap.New()

	if resp.GetHoldings() != nil {
		for _, holding := range resp.GetHoldings() {
			holdingRecord := transformer.NewInvestmentHolding(&holding)
			// populate hashmap with key as the accountID (plaid acct id) and the value a hashset of investment holding structs
			if _, ok := accountIDToInvesmentHoldingMap.Get(holding.AccountId); !ok {
				v := hashset.New()
				v.Add(holdingRecord)
				accountIDToInvesmentHoldingMap.Put(holding.AccountId, v)
			} else {
				res, _ := accountIDToInvesmentHoldingMap.Get(holding.AccountId)
				value := res.(*hashset.Set)
				value.Add(holdingRecord)
				accountIDToInvesmentHoldingMap.Put(holding.AccountId, value)
			}

			// securities and holding retain a one to one mapping to one another
			if _, ok := securityIDToInvestmentHoldingMap.Get(holding.SecurityId); !ok {
				securityIDToInvestmentHoldingMap.Put(holding.SecurityId, holdingRecord)
			}
		}

		// populate a hashmap comprised of all the securities per security id
		for _, security := range resp.GetSecurities() {
			securityRecord := transformer.TransformSinglePlaidInvestmentSecurity(security)
			// populate hashmap with key as the accountID (plaid acct id) and the value a hashset of investment holding structs
			if _, ok := securityIDToSecurityMap.Get(security.SecurityId); !ok {
				securityIDToSecurityMap.Put(security.SecurityId, securityRecord)
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
			value, ok := holdings.(*hashset.Set)
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
				if currentHolding, ok := securityIDToInvestmentHoldingMap.Get(holding.SecurityId); ok {
					presentHolding, ok := currentHolding.(*schema.InvesmentHolding)
					if !ok {
						return nil, fmt.Errorf("could not convert value to *schema.InvestmentSecurity")
					}

					if security, ok := securityIDToSecurityMap.Get(presentHolding.SecurityId); ok {
						presentSecurity, ok := security.(*schema.InvestmentSecurity)
						if !ok {
							return nil, fmt.Errorf("could not convert value to *schema.InvestmentSecurity")
						}
						account.Securities = append(account.Securities, presentSecurity)
					}
				}
			}
		}

		result = append(result, account)
	}

	return result, nil
}
