package transformer

import (
	"fmt"

	"github.com/plaid/plaid-go/v12/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// NewPlaidBankAccount converts a Plaid account to our own bank account interface.
func NewPlaidBankAccount(userID uint64, bankAccount *plaid.AccountBase) (*schema.BankAccount, error) {
	if bankAccount == nil {
		return nil, fmt.Errorf("invalid input argument. bank account cannot be nil")
	}

	if bankAccount.Type != plaid.ACCOUNTTYPE_DEPOSITORY {
		return nil, fmt.Errorf("invalid input argument. bank account must be of type depository")
	}

	return &schema.BankAccount{
		Id:             0,
		UserId:         userID,
		Name:           bankAccount.GetName(),
		Number:         bankAccount.GetMask(),
		Type:           schema.BankAccountType_BANK_ACCOUNT_TYPE_PLAID,
		Balance:        float32(bankAccount.Balances.GetAvailable()),
		Currency:       bankAccount.Balances.GetIsoCurrencyCode(),
		CurrentFunds:   float64(bankAccount.Balances.GetCurrent()),
		BalanceLimit:   uint64(bankAccount.Balances.GetLimit()),
		Pockets:        []*schema.Pocket{},
		PlaidAccountId: bankAccount.AccountId,
		Subtype:        string(bankAccount.GetSubtype()),
		Status:         schema.BankAccountStatus_BANK_ACCOUNT_STATUS_ACTIVE,
	}, nil

}
