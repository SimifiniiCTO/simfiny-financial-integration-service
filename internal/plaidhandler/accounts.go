package plaidhandler

import (
	"context"

	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// GetAccounts is used to retrieve accounts from the Plaid API for a given access token.
func (p *PlaidWrapper) GetAccounts(ctx context.Context, accessToken string, userId uint64) ([]*schema.BankAccount, error) {
	// Build the get accounts request.
	request := p.client.PlaidApi.
		AccountsGet(ctx).
		AccountsGetRequest(plaid.AccountsGetRequest{
			ClientId:    &p.ClientID,
			Secret:      &p.SecretKey,
			AccessToken: accessToken,
		})

	// Send the request.
	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to get accounts from Plaid", zap.Error(err))
		return nil, err
	}

	plaidAccounts := result.GetAccounts()
	accounts := make([]*schema.BankAccount, 0)

	// Once we have our data, convert all of the results from our request to our own bank account interface.
	for _, plaidAccount := range plaidAccounts {
		if plaidAccount.Type == plaid.ACCOUNTTYPE_DEPOSITORY {
			acct := plaidAccount
			data, err := transformer.NewPlaidBankAccount(userId, &acct)
			if err != nil {
				p.Logger.Error("failed to convert bank account", zap.Error(err))
				return nil, err
			}

			accounts = append(accounts, data)
		}

	}

	return accounts, nil
}
