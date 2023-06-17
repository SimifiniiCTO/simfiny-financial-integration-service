package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/v12/plaid"
)

// DeleteItem is used to delete an item from the Plaid API for a given access token
func (p *PlaidWrapper) DeleteItem(ctx context.Context, accessToken *string) error {
	if accessToken == nil {
		return errors.New("invalid input argument. access token cannot be empty")
	}

	request := plaid.NewItemRemoveRequest(*accessToken)
	if _, _, err := p.client.PlaidApi.
		ItemRemove(ctx).
		ItemRemoveRequest(*request).
		Execute(); err != nil {
		return err
	}

	return nil
}
