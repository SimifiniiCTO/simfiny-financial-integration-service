package transformer

import (
	"fmt"

	"github.com/plaid/plaid-go/v12/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// NewInvestmentAccount converts a Plaid account to our own investment account interface.
func NewInvestmentAccount(userID uint64, bankAccount *plaid.AccountBase) (*schema.InvestmentAccount, error) {
	if bankAccount == nil {
		return nil, fmt.Errorf("invalid input argument. bank account cannot be nil")
	}

	if bankAccount.Type != plaid.ACCOUNTTYPE_INVESTMENT {
		return nil, fmt.Errorf("invalid input argument. bank account must be of type investment")
	}

	return &schema.InvestmentAccount{
		Id:             0,
		UserId:         userID,
		Name:           bankAccount.Name,
		Number:         bankAccount.GetMask(),
		Type:           string(bankAccount.Type),
		Balance:        float32(bankAccount.Balances.GetAvailable()),
		CurrentFunds:   float64(bankAccount.Balances.GetCurrent()),
		BalanceLimit:   uint64(bankAccount.Balances.GetLimit()),
		PlaidAccountId: bankAccount.GetAccountId(),
		Subtype:        string(*bankAccount.GetSubtype().Ptr()),
		Holdings:       []*schema.InvesmentHolding{},
		Securities:     []*schema.InvestmentSecurity{},
	}, nil
}

// NewInvestmentHolding converts a Plaid holding to our own investment holding interface.
func NewInvestmentHolding(input *plaid.Holding) *schema.InvesmentHolding {
	return &schema.InvesmentHolding{
		Id:                       0,
		PlaidAccountId:           input.GetAccountId(),
		CostBasis:                float64(input.GetCostBasis()),
		InstitutionPrice:         float64(input.GetInstitutionPrice()),
		InstitutionPriceAsOf:     input.GetInstitutionPriceAsOf(),
		InstitutionPriceDatetime: input.GetInstitutionPriceDatetime().Format("2006-01-02 15:04:05"),
		InstitutionValue:         float64(input.GetInstitutionValue()),
		IsoCurrencyCode:          input.GetIsoCurrencyCode(),
		Quantity:                 float64(input.GetQuantity()),
		SecurityId:               input.GetSecurityId(),
		UnofficialCurrencyCode:   input.GetUnofficialCurrencyCode(),
	}
}
