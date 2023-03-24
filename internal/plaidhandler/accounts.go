package plaidhandler

import (
	"context"

	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
)

// GetAccounts implements PlaidWrapperImpl
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
	accounts := make([]*schema.BankAccount, len(plaidAccounts))

	// Once we have our data, convert all of the results from our request to our own bank account interface.
	for i, plaidAccount := range plaidAccounts {
		accounts[i], err = transformer.NewPlaidBankAccount(userId, plaidAccount)
		if err != nil {
			p.Logger.Error("failed to convert bank account", zap.Error(err))
			return nil, err
		}
	}

	return accounts, nil
}
