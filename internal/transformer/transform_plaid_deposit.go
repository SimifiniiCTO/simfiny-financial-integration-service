package transformer

import (
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/v12/plaid"
)

func transformPlaidAccountToDepositAccount(accounts *[]plaid.AccountBase) []*schema.DepositAccount {
	depositAccts := make([]*schema.DepositAccount, 0)
	for _, el := range *accounts {
		balance := el.GetBalances()
		depositAcct := &schema.DepositAccount{
			PlaidAccountID: el.GetAccountId(),
			AccountSubtype: string(el.GetSubtype()),
			AccountType:    string(el.GetType()),
			AccountName:    el.GetName(),
			BalanceID: &schema.Balance{
				AvailableFunds: float64(balance.GetAvailable()),
				CurrentFunds:   float64(balance.GetCurrent()),
				CurrencyCode:   balance.GetIsoCurrencyCode(),
				BalanceLimit:   uint64(balance.GetLimit()),
			},
		}

		depositAccts = append(depositAccts, depositAcct)
	}

	return depositAccts
}
