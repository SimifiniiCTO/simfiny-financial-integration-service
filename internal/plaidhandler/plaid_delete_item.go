package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/v12/plaid"
)

// DeleteItem implements PlaidWrapperImpl
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
