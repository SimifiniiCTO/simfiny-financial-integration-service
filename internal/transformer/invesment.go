package transformer

import (
	"github.com/plaid/plaid-go/v12/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// NewInvestmentAccount converts a Plaid account to our own investment account interface.
func NewInvestmentAccount(userID uint64, input *plaid.AccountBase) *schema.InvestmentAccount {
	return &schema.InvestmentAccount{
		Id:             0,
		UserId:         userID,
		Name:           input.Name,
		Number:         input.GetMask(),
		Type:           string(input.Type),
		Balance:        float32(input.Balances.GetAvailable()),
		CurrentFunds:   float64(input.Balances.GetCurrent()),
		BalanceLimit:   uint64(input.Balances.GetLimit()),
		PlaidAccountId: input.GetAccountId(),
		Subtype:        string(*input.GetSubtype().Ptr()),
		Holdings:       []*schema.InvesmentHolding{},
		Securities:     []*schema.InvestmentSecurity{},
	}
}

// NewInvestmentHolding converts a Plaid holding to our own investment holding interface.
func NewInvestmentHolding(input *plaid.Holding) *schema.InvesmentHolding {
	return &schema.InvesmentHolding{
		Id:                     0,
		PlaidAccountId:         input.AccountId,
		CostBasis:              float64(*input.CostBasis.Get()),
		InstitutionPrice:       float64(input.InstitutionPrice),
		InstitutionPriceAsOf:   input.GetInstitutionPriceAsOf(),
		InstitutionValue:       float64(input.InstitutionValue),
		IsoCurrencyCode:        input.GetIsoCurrencyCode(),
		Quantity:               float64(input.Quantity),
		SecurityId:             input.SecurityId,
		UnofficialCurrencyCode: input.GetUnofficialCurrencyCode(),
	}
}
