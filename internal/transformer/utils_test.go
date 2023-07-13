package transformer

import (
	"time"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/plaid/plaid-go/v14/plaid"
)

func generateMultiplePlaidAccountBase(acctType plaid.AccountType, count int) []*plaid.AccountBase {
	accounts := make([]*plaid.AccountBase, 0)
	for i := 0; i < count; i++ {
		accounts = append(accounts, generateSinglePlaidAccountBase(acctType))
	}

	return accounts
}

func generateSinglePlaidAccountBase(acctType plaid.AccountType) *plaid.AccountBase {
	return &plaid.AccountBase{
		AccountId:            helper.GenerateRandomString(10),
		Balances:             plaid.AccountBalance{Available: *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))), Current: *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))), Limit: *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))), IsoCurrencyCode: *plaid.NewNullableString(pointer.StringP("USD")), UnofficialCurrencyCode: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))), LastUpdatedDatetime: *plaid.NewNullableTime(pointer.TimeP(time.Now())), AdditionalProperties: map[string]interface{}{}},
		Mask:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Name:                 helper.GenerateRandomString(10),
		OfficialName:         *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Type:                 acctType,
		Subtype:              plaid.NullableAccountSubtype{},
		VerificationStatus:   pointer.StringP(helper.GenerateRandomString(10)),
		PersistentAccountId:  pointer.StringP(helper.GenerateRandomString(10)),
		AdditionalProperties: map[string]interface{}{},
	}
}

func generateSingleCreditCardLiability() *plaid.CreditCardLiability {
	return &plaid.CreditCardLiability{
		AccountId:              *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Aprs:                   generateMultiplePlaidAPR(10),
		IsOverdue:              *plaid.NewNullableBool(pointer.BoolP(false)),
		LastPaymentAmount:      *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		LastPaymentDate:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LastStatementIssueDate: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LastStatementBalance:   *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		MinimumPaymentAmount:   *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		NextPaymentDueDate:     *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		AdditionalProperties:   map[string]interface{}{},
	}
}

func generateMultiplePlaidAPR(count int) []plaid.APR {
	aprs := make([]plaid.APR, 0)
	for i := 0; i < count; i++ {
		aprs = append(aprs, generateSinglePlaidAPR())
	}

	return aprs
}

func generateSinglePlaidAPR() plaid.APR {
	return plaid.APR{
		AprPercentage:        float64(helper.GenerateRandomId(1000, 1000000)),
		AprType:              helper.GenerateRandomString(10),
		BalanceSubjectToApr:  *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		InterestChargeAmount: *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		AdditionalProperties: map[string]interface{}{},
	}
}

func generateInvestmentHolding() *plaid.Holding {
	return &plaid.Holding{
		AccountId:                helper.GenerateRandomString(10),
		SecurityId:               helper.GenerateRandomString(10),
		InstitutionPrice:         float64(helper.GenerateRandomId(1000, 1000000)),
		InstitutionPriceAsOf:     *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		InstitutionPriceDatetime: *plaid.NewNullableTime(pointer.TimeP(time.Now())),
		InstitutionValue:         float64(helper.GenerateRandomId(1000, 1000000)),
		CostBasis:                *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		Quantity:                 float64(helper.GenerateRandomId(1000, 1000000)),
		IsoCurrencyCode:          *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		UnofficialCurrencyCode:   *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		AdditionalProperties:     map[string]interface{}{},
	}
}

func generatePlaidTransaction() *plaid.Transaction {
	return &plaid.Transaction{
		AccountId:              helper.GenerateRandomString(10),
		Amount:                 0,
		IsoCurrencyCode:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		UnofficialCurrencyCode: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Category: []string{
			helper.GenerateRandomString(10),
			helper.GenerateRandomString(10),
			helper.GenerateRandomString(10),
			helper.GenerateRandomString(10),
		},
		CategoryId:  *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		CheckNumber: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Date:        helper.GenerateRandomString(10),
		Location: plaid.Location{
			Address:              *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			City:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Region:               *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			PostalCode:           *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Country:              *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Lat:                  *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
			Lon:                  *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
			StoreNumber:          *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			AdditionalProperties: map[string]interface{}{},
		},
		Name:                helper.GenerateRandomString(10),
		MerchantName:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		OriginalDescription: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		PaymentMeta: plaid.PaymentMeta{
			ReferenceNumber:      *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			PpdId:                *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Payee:                *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			ByOrderOf:            *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Payer:                *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			PaymentMethod:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			PaymentProcessor:     *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Reason:               *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			AdditionalProperties: map[string]interface{}{},
		},
		Pending:                        false,
		PendingTransactionId:           *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		AccountOwner:                   *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		TransactionId:                  helper.GenerateRandomString(10),
		TransactionType:                pointer.StringP(helper.GenerateRandomString(10)),
		LogoUrl:                        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Website:                        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		AuthorizedDate:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		AuthorizedDatetime:             *plaid.NewNullableTime(pointer.TimeP(time.Now())),
		Datetime:                       *plaid.NewNullableTime(pointer.TimeP(time.Now())),
		PaymentChannel:                 helper.GenerateRandomString(10),
		PersonalFinanceCategory:        plaid.NullablePersonalFinanceCategory{},
		TransactionCode:                plaid.NullableTransactionCode{},
		PersonalFinanceCategoryIconUrl: pointer.StringP(helper.GenerateRandomString(10)),
		Counterparties: &[]plaid.TransactionCounterparty{
			{
				Name:    helper.GenerateRandomString(10),
				Type:    plaid.COUNTERPARTYTYPE_FINANCIAL_INSTITUTION,
				Website: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
				LogoUrl: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			},
		},
	}
}

func generateMultipleAccountMetadata(count int) []*accountMetadata {
	accts := make([]*accountMetadata, 0)
	for i := 0; i < count; i++ {
		accts = append(accts, generateAccountMetadata())
	}

	return accts
}

func generateAccountMetadata() *accountMetadata {
	return &accountMetadata{
		accountSubtype: helper.GenerateRandomString(10),
		accountType:    helper.GenerateRandomString(10),
		accountName:    helper.GenerateRandomString(10),
		accountID:      helper.GenerateRandomString(10),
		availableFunds: float64(helper.GenerateRandomId(1000, 1000000)),
		currentFunds:   float64(helper.GenerateRandomId(1000, 1000000)),
		currencyCode:   helper.GenerateRandomString(10),
		balanceLimit:   uint64(helper.GenerateRandomId(1000, 1000000)),
		accountNumber:  helper.GenerateRandomString(10),
	}
}

func generateManyPlaidCreditLiabilityAndAccountMetadata(count int) ([]plaid.CreditCardLiability, map[string]*accountMetadata) {
	acctMetadata := generateAccountMetadata()

	plaidCreditCardLiabilities := make([]plaid.CreditCardLiability, 0)
	for i := 0; i < count; i++ {
		liability := generateSingleCreditCardLiability()
		liability.AccountId = *plaid.NewNullableString(pointer.StringP(acctMetadata.accountID))
		plaidCreditCardLiabilities = append(plaidCreditCardLiabilities, *liability)
	}

	acctMetadataMap := make(map[string]*accountMetadata)
	acctMetadataMap[acctMetadata.accountID] = acctMetadata

	return plaidCreditCardLiabilities, acctMetadataMap
}

func generateSingleMortgageLiability() *plaid.MortgageLiability {
	return &plaid.MortgageLiability{
		AccountId:            helper.GenerateRandomString(10),
		AccountNumber:        helper.GenerateRandomString(10),
		CurrentLateFee:       *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		EscrowBalance:        *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		HasPmi:               *plaid.NewNullableBool(pointer.BoolP(false)),
		HasPrepaymentPenalty: *plaid.NewNullableBool(pointer.BoolP(false)),
		InterestRate: plaid.MortgageInterestRate{
			Percentage:           *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
			Type:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			AdditionalProperties: map[string]interface{}{},
		},
		LastPaymentAmount:          *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		LastPaymentDate:            *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LoanTypeDescription:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LoanTerm:                   *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		MaturityDate:               *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		NextMonthlyPayment:         *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		NextPaymentDueDate:         *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		OriginationDate:            *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		OriginationPrincipalAmount: *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		PastDueAmount:              *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		PropertyAddress: plaid.MortgagePropertyAddress{
			City:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Country:              *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			PostalCode:           *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Region:               *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Street:               *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			AdditionalProperties: map[string]interface{}{},
		},
		YtdInterestPaid:      *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		YtdPrincipalPaid:     *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		AdditionalProperties: map[string]interface{}{},
	}
}

func generateManyPlaidMortgageLiabilityAndAccountMetadata(count int) ([]plaid.MortgageLiability, map[string]*accountMetadata) {
	acctMetadata := generateAccountMetadata()

	plaidMortgageLiabilities := make([]plaid.MortgageLiability, 0)
	for i := 0; i < count; i++ {
		liability := generateSingleMortgageLiability()
		liability.AccountId = acctMetadata.accountID
		plaidMortgageLiabilities = append(plaidMortgageLiabilities, *liability)
	}

	acctMetadataMap := make(map[string]*accountMetadata)
	acctMetadataMap[acctMetadata.accountID] = acctMetadata

	return plaidMortgageLiabilities, acctMetadataMap
}

func generateSinglePlaidStudentLoan() *plaid.StudentLoan {
	return &plaid.StudentLoan{
		AccountId:     *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		AccountNumber: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		DisbursementDates: []string{
			helper.GenerateRandomString(10),
			helper.GenerateRandomString(10),
			helper.GenerateRandomString(10),
			helper.GenerateRandomString(10),
		},
		ExpectedPayoffDate:     *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Guarantor:              *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		InterestRatePercentage: 0,
		IsOverdue:              *plaid.NewNullableBool(pointer.BoolP(false)),
		LastPaymentAmount:      *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		LastPaymentDate:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LastStatementIssueDate: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LoanName:               *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		LoanStatus: plaid.StudentLoanStatus{
			EndDate:              *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Type:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			AdditionalProperties: map[string]interface{}{},
		},
		MinimumPaymentAmount:       *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		NextPaymentDueDate:         *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		OriginationDate:            *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		OriginationPrincipalAmount: *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		OutstandingInterestAmount:  *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		PaymentReferenceNumber:     *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		PslfStatus: plaid.PSLFStatus{
			EstimatedEligibilityDate: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			PaymentsMade:             plaid.NullableInt32{},
			PaymentsRemaining:        plaid.NullableInt32{},
			AdditionalProperties:     map[string]interface{}{},
		},
		RepaymentPlan: plaid.StudentRepaymentPlan{
			Description:          *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			Type:                 *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
			AdditionalProperties: map[string]interface{}{},
		},
		SequenceNumber:       *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		ServicerAddress:      plaid.ServicerAddressData{},
		YtdInterestPaid:      *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		YtdPrincipalPaid:     *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		AdditionalProperties: map[string]interface{}{},
	}
}

func generateMultipleStudentLoanLiability(count int) []*plaid.StudentLoan {
	studentLoans := make([]*plaid.StudentLoan, 0)
	for i := 0; i < count; i++ {
		studentLoans = append(studentLoans, generateSinglePlaidStudentLoan())
	}

	return studentLoans
}

func generateManyPlaidStudentLoanAndAccountMetadata(count int) ([]plaid.StudentLoan, map[string]*accountMetadata) {
	acctMetadata := generateAccountMetadata()

	plaidStudentLoans := make([]plaid.StudentLoan, 0)
	for i := 0; i < count; i++ {
		liability := generateSinglePlaidStudentLoan()
		liability.AccountId = *plaid.NewNullableString(pointer.StringP(acctMetadata.accountID))
		plaidStudentLoans = append(plaidStudentLoans, *liability)
	}

	acctMetadataMap := make(map[string]*accountMetadata)
	acctMetadataMap[acctMetadata.accountID] = acctMetadata

	return plaidStudentLoans, acctMetadataMap
}

func generateSinglSecurity() *plaid.Security {
	return &plaid.Security{
		ClosePrice:             *plaid.NewNullableFloat64(pointer.Float64P(float64(helper.GenerateRandomId(1000, 1000000)))),
		ClosePriceAsOf:         *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Cusip:                  *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		InstitutionId:          *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		InstitutionSecurityId:  *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		IsCashEquivalent:       *plaid.NewNullableBool(pointer.BoolP(false)),
		Isin:                   *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		IsoCurrencyCode:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Name:                   *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		ProxySecurityId:        *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		SecurityId:             helper.GenerateRandomString(10),
		Sedol:                  *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		TickerSymbol:           *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		Type:                   *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		UnofficialCurrencyCode: *plaid.NewNullableString(pointer.StringP(helper.GenerateRandomString(10))),
		UpdateDatetime:         *plaid.NewNullableTime(pointer.TimeP(time.Now())),
		AdditionalProperties:   map[string]interface{}{},
	}
}

func generateMultipleSecurity(count int) []plaid.Security {
	securities := make([]plaid.Security, 0)
	for i := 0; i < count; i++ {
		securities = append(securities, *generateSinglSecurity())
	}

	return securities
}
