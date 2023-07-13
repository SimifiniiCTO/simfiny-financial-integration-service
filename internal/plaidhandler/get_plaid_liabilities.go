package plaidhandler

import (
	"context"
	"errors"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/plaid/plaid-go/v14/plaid"
)

// The CreditAccountSet type contains slices of different types of accounts, including credit card,
// mortgage, and student loan accounts.
// @property {[]*schema.CreditAccount} CrediCardAccounts - This is a slice of pointers to
// `schema.CreditAccount` objects, representing a set of credit card accounts. Each
// `schema.CreditAccount` object contains information about a single credit card account, such as the
// account holder's name, account number, balance, and credit limit.
// @property {[]*schema.MortgageAccount} MortgageLoanAccts - The `MortgageLoanAccts` property is a
// slice of pointers to `schema.MortgageAccount` structs. It is a part of the `CreditAccountSet`
// struct, which is used to group together different types of credit accounts. The `MortgageLoanAccts`
// slice
// @property {[]*schema.StudentLoanAccount} StudentLoanAccts - `StudentLoanAccts` is a slice of
// pointers to `schema.StudentLoanAccount` structs. It is a property of the `CreditAccountSet` struct,
// which is used to group together different types of credit accounts. This particular slice contains
// information about student loan accounts.
type CreditAccountSet struct {
	CrediCardAccounts []*schema.CreditAccount
	MortgageLoanAccts []*schema.MortgageAccount
	StudentLoanAccts  []*schema.StudentLoanAccount
}

// GetPlaidLiabilityAccounts is used to retrieve liability accounts from the Plaid API for a given set of account IDs
func (p *PlaidWrapper) GetPlaidLiabilityAccounts(ctx context.Context, accessToken *string, accountIds ...string) (*CreditAccountSet, error) {
	var (
		err                  error
		creditCardAccounts   []*schema.CreditAccount
		mortgageLoanAccounts []*schema.MortgageAccount
		studentLoanAccounts  []*schema.StudentLoanAccount
	)

	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	request := plaid.NewLiabilitiesGetRequest(*accessToken)
	request.SetClientId(p.ClientID)
	request.SetSecret(p.SecretKey)

	if len(accountIds) > 0 {
		request.SetOptions(
			plaid.LiabilitiesGetRequestOptions{
				AccountIds: &accountIds,
			})
	}

	resp, _, err := p.client.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	accts := filterAccountsOfTwoType(resp.GetAccounts(), string(plaid.ACCOUNTTYPE_CREDIT), string(plaid.ACCOUNTTYPE_LOAN))
	acctMetadata := transformer.GetAccountMetadata(&accts)

	creditCardLiabilities := resp.GetLiabilities().Credit
	mortgageLoanLiabilities := resp.GetLiabilities().Mortgage
	studentLoanLiabilities := resp.GetLiabilities().Student

	if len(creditCardLiabilities) > 0 {
		creditCardAccounts, err = transformer.TransformPlaidCredit(&creditCardLiabilities, acctMetadata)
		if err != nil {
			return nil, err
		}
	}

	if len(mortgageLoanLiabilities) > 0 {
		mortgageLoanAccounts, err = transformer.TransformMortgageAccount(&mortgageLoanLiabilities)
		if err != nil {
			return nil, err
		}
	}

	if len(studentLoanLiabilities) > 0 {
		studentLoanAccounts, err = transformer.TransformStudentloanAccount(&studentLoanLiabilities, acctMetadata)
		if err != nil {
			return nil, err
		}
	}

	return &CreditAccountSet{
		CrediCardAccounts: creditCardAccounts,
		MortgageLoanAccts: mortgageLoanAccounts,
		StudentLoanAccts:  studentLoanAccounts,
	}, nil
}
