package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TransformPlaidInvestmentSecurity(securities []plaid.Security) ([]*schema.InvestmentSecurity, error) {
	if len(securities) == 0 {
		return nil, errors.New("invalid input argument. securities cannot be empty")
	}

	securitySet := make([]*schema.InvestmentSecurity, 0)
	for _, security := range securities {
		sec := TransformSinglePlaidInvestmentSecurity(security)

		securitySet = append(securitySet, sec)
	}

	return securitySet, nil
}

func TransformSinglePlaidInvestmentSecurity(security plaid.Security) *schema.InvestmentSecurity {
	return &schema.InvestmentSecurity{
		ClosePrice:             security.GetClosePrice(),
		ClosePriceAsOf:         security.GetClosePriceAsOf(),
		Cusip:                  security.GetCusip(),
		InstitutionId:          security.GetInstitutionId(),
		InstitutionSecurityId:  security.GetInstitutionSecurityId(),
		IsCashEquivalent:       security.GetIsCashEquivalent(),
		Isin:                   security.GetIsin(),
		IsoCurrencyCode:        security.GetIsoCurrencyCode(),
		Name:                   security.GetName(),
		ProxySecurityId:        security.GetProxySecurityId(),
		SecurityId:             security.GetSecurityId(),
		Sedol:                  security.GetSedol(),
		TickerSymbol:           security.GetTickerSymbol(),
		Type:                   security.GetType(),
		UnofficialCurrencyCode: security.GetUnofficialCurrencyCode(),
		UpdateDatetime:         security.GetUpdateDatetime().String(),
	}
}
