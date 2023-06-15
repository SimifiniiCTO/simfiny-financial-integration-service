package plaidhandler

import (
	"context"
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	"github.com/plaid/plaid-go/v12/plaid"
)

type CreditAccountSet struct {
	CrediCardAccounts []*schema.CreditAccount
	MortgageLoanAccts []*schema.MortgageAccount
	StudentLoanAccts  []*schema.StudentLoanAccount
}

func (p *PlaidWrapper) GetPlaidLiabilityAccounts(ctx context.Context, accessToken *string) (*CreditAccountSet, error) {
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
		mortgageLoanAccounts, err = transformer.TransformMortgageAccount(&mortgageLoanLiabilities, acctMetadata)
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
