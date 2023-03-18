package plaidhandler

import (
	"time"

	"github.com/plaid/plaid-go/plaid"
)

type ItemToken struct {
	// The access token associated with the Item data is being requested for.
	AccessToken string
	// The `item_id` value of the Item associated with the returned `access_token`
	ItemId string
}

// NewItemTokenFromPlaid creates a new ItemToken from a plaid.ItemPublicTokenExchangeResponse
func NewItemTokenFromPlaid(input plaid.ItemPublicTokenExchangeResponse) (ItemToken, error) {
	return ItemToken{
		AccessToken: input.GetAccessToken(),
		ItemId:      input.GetItemId(),
	}, nil
}

type LinkTokenOptions struct {
	ClientUserID             string
	LegalName                string
	PhoneNumber              *string
	PhoneNumberVerifiedTime  *time.Time
	EmailAddress             string
	EmailAddressVerifiedTime *time.Time
	UpdateMode               bool
}

type LinkToken interface {
	Token() string
	Expiration() time.Time
}

var (
	_ LinkToken = PlaidLinkToken{}
)

type PlaidLinkToken struct {
	LinkToken string
	Expires   time.Time
}

func (p PlaidLinkToken) Token() string {
	return p.LinkToken
}

func (p PlaidLinkToken) Expiration() time.Time {
	return p.Expires
}
