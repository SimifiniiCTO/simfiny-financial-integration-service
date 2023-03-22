package transformer

import (
	"github.com/plaid/plaid-go/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// NewPlaidBankAccount converts a Plaid account to our own bank account interface.
func NewPlaidBankAccount(userID uint64, bankAccount plaid.AccountBase) (*schema.BankAccount, error) {

	return &schema.BankAccount{
		Id:             0,
		UserId:         userID,
		Name:           bankAccount.GetName(),
		Number:         bankAccount.GetMask(),
		Type:           schema.BankAccountType_BANK_ACCOUNT_TYPE_PLAID,
		Balance:        bankAccount.Balances.GetAvailable(),
		Currency:       "USD",
		CurrentFunds:   float64(bankAccount.Balances.GetCurrent()),
		BalanceLimit:   uint64(bankAccount.Balances.GetLimit()),
		Pockets:        []*schema.Pocket{}, // TODO: auto generate default pockets
		PlaidAccountId: bankAccount.AccountId,
		Subtype:        string(bankAccount.GetSubtype()),
	}, nil

}
