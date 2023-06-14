package transformer

import (
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func transformPlaidAccountToDepositAccount(accounts *[]plaid.AccountBase) []*schema.BankAccount {
	depositAccts := make([]*schema.BankAccount, 0)
	for _, el := range *accounts {
		balance := el.GetBalances()
		depositAcct := &schema.BankAccount{
			Id:             0,
			UserId:         0,
			Name:           el.GetName(),
			Number:         el.GetMask(),
			Type:           schema.BankAccountType_BANK_ACCOUNT_TYPE_PLAID,
			Balance:        float32(*el.GetBalances().Available.Get()),
			Currency:       *el.GetBalances().IsoCurrencyCode.Get(),
			CurrentFunds:   float64(balance.GetCurrent()),
			BalanceLimit:   uint64(balance.GetLimit()),
			Pockets:        []*schema.Pocket{},
			PlaidAccountId: el.GetAccountId(),
			Subtype:        string(el.GetSubtype()),
			Status:         schema.BankAccountStatus_BANK_ACCOUNT_STATUS_ACTIVE,
		}

		depositAccts = append(depositAccts, depositAcct)
	}

	return depositAccts
}
