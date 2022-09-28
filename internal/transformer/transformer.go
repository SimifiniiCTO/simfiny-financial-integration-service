package transformer

import (
	"context"
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/plaid"
)

type accountMetadata struct {
	accountSubtype string
	accountType    string
	accountName    string
	accountID      string
	balance        *schema.Balance
}

// TransformPlaidInvestmentResponse transforms an investment response and as a side effect updates the virtual account object
func TransformPlaidInvestmentResponse(ctx context.Context, vAcct *schema.VirtualAccount, securities *[]plaid.Security,
	holdings *[]plaid.Holding, accounts *[]plaid.AccountBase) error {
	var err error
	if securities == nil {
		return errors.New("invalid input arguments. securities object cannot be nil")
	}

	if holdings == nil {
		return errors.New("invalid input arguments. holdings object cannot be nil")
	}

	if accounts == nil {
		return errors.New("invalid input arguments. account object cannot be nil")
	}

	acctIDtoAcctObj := getAccountMetadata(accounts)

	// deconstruct and transform liabilities object
	vAcct.InvestmentAccountID, err = transformPlaidInvestmentObject(securities, holdings, acctIDtoAcctObj)
	if err != nil {
		return err
	}

	return nil
}

// TransformPlaidLiabilitiesResponse transforms liability response and as a side effect updates the virtual account object
func TransformPlaidLiabilitiesResponse(ctx context.Context, vAcct *schema.VirtualAccount, liabilities *plaid.LiabilitiesObject,
	accounts *[]plaid.AccountBase) error {
	var err error
	if liabilities == nil {
		return errors.New("invalid input arguments. liabilities object cannot be nil")
	}

	if accounts == nil {
		return errors.New("invalid input arguments. account object cannot be nil")
	}

	acctIDtoAcctObj := getAccountMetadata(accounts)

	// deconstruct and transform liabilities object
	plaidCreditAccts := liabilities.GetCredit()
	vAcct.CreditAccountID, err = transformPlaidCreditObject(&plaidCreditAccts, acctIDtoAcctObj)
	if err != nil {
		return err
	}

	plaidMortgageAccts := liabilities.GetMortgage()
	vAcct.MortgageLoanAccountID, err = transformMortgageObject(&plaidMortgageAccts, acctIDtoAcctObj)
	if err != nil {
		return err
	}

	plaidStudentLoanAccts := liabilities.GetStudent()
	vAcct.StudentLoanAccountID, err = transformStudentloanObject(&plaidStudentLoanAccts, acctIDtoAcctObj)
	if err != nil {
		return err
	}
	return nil
}

// TransformPlaidDepositResponse transforms a deposit response and as a side effect updates the virtual account object
func TransformPlaidDepositResponse(ctx context.Context, vAcct *schema.VirtualAccount, accounts *[]plaid.AccountBase) error {
	if vAcct == nil {
		return errors.New("invalid input arguments. vAcct object cannot be nil")
	}

	if accounts == nil {
		return errors.New("invalid input arguments. account object cannot be nil")
	}

	vAcct.DepositAccountID = transformPlaidAccountToDepositAccount(accounts)
	return nil
}

func getAccountMetadata(accounts *[]plaid.AccountBase) map[string]*accountMetadata {
	acctIDtoAcctObj := make(map[string]*accountMetadata)
	// iterate over account object and store account id as key and account subtype + type as value
	for _, el := range *accounts {
		balance := el.GetBalances()
		acctIDtoAcctObj[el.AccountId] = &accountMetadata{
			accountSubtype: string(el.GetSubtype()),
			accountType:    string(el.GetType()),
			accountID:      el.GetAccountId(),
			accountName:    el.GetName(),
			balance: &schema.Balance{
				AvailableFunds: float64(balance.GetAvailable()),
				CurrentFunds:   float64(balance.GetCurrent()),
				CurrencyCode:   balance.GetIsoCurrencyCode(),
				BalanceLimit:   uint64(balance.GetLimit()),
			},
		}
	}
	return acctIDtoAcctObj
}
