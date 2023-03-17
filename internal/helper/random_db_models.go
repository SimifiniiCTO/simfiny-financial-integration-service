package helper

import (
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func GenerateBugetCategory() *schema.Category {
	return &schema.Category{
		Id:          0,
		Name:        GenerateRandomString(10),
		Description: GenerateRandomString(10),
	}
}

// GenerateRandomString generates a random string of length n
func GenerateSingleBudget() *schema.Budget {
	return &schema.Budget{
		Id:          0,
		Name:        GenerateRandomString(10),
		Description: GenerateRandomString(10),
		StartDate:   GenerateRandomString(10),
		EndDate:     GenerateRandomString(10),
		Category:    GenerateBugetCategory(),
	}
}

func GenerateBudgets(budgetCount int) []*schema.Budget {
	budgets := make([]*schema.Budget, 0, budgetCount)
	for i := 0; i < budgetCount; i++ {
		budgets = append(budgets, GenerateSingleBudget())
	}
	return budgets
}

func GenerateSingleMilestone(numBudgets int) *schema.Milestone {
	return &schema.Milestone{
		Name:         GenerateRandomString(10),
		Description:  GenerateRandomString(10),
		IsCompleted:  false,
		TargetDate:   GenerateRandomString(10),
		TargetAmount: GenerateRandomString(10),
		Budget:       GenerateSingleBudget(),
	}
}

func GenerateMilestone(milestoneCount int) []*schema.Milestone {
	milestones := make([]*schema.Milestone, 0, milestoneCount)
	for i := 0; i < milestoneCount; i++ {
		milestones = append(milestones, GenerateSingleMilestone(3))
	}

	return milestones
}

func GenerateSingleForecast() *schema.Forecast {
	return &schema.Forecast{
		Id:                       0,
		ForecastedAmount:         GenerateRandomString(10),
		ForecastedCompletionDate: GenerateRandomString(10),
		VarianceAmount:           GenerateRandomString(10),
	}
}

func GenerateSingleGoal(numMilestones int) *schema.SmartGoal {
	return &schema.SmartGoal{
		Id:           0,
		UserId:       GenerateRandomString(10),
		Name:         GenerateRandomString(10),
		Description:  GenerateRandomString(10),
		IsCompleted:  false,
		GoalType:     schema.GoalType_GOAL_TYPE_DEBT,
		Duration:     GenerateRandomString(10),
		StartDate:    GenerateRandomString(10),
		EndDate:      GenerateRandomString(10),
		TargetAmount: GenerateRandomString(10),
		Milestones:   GenerateMilestone(numMilestones),
		Forecasts:    GenerateSingleForecast(),
	}
}

func GenerateGoals(goalCount int) []*schema.SmartGoal {
	goals := make([]*schema.SmartGoal, 0, goalCount)
	for i := 0; i < goalCount; i++ {
		goals = append(goals, GenerateSingleGoal(10))
	}

	return goals
}

func GenerateSinglePocket() *schema.Pocket {
	return &schema.Pocket{
		Id:    0,
		Type:  schema.PocketType_POCKET_TYPE_DEBT_REDUCTION,
		Goals: GenerateGoals(3),
	}
}

func GenerateSinglePocketWithGoals(goalCount int) *schema.Pocket {
	return &schema.Pocket{
		Id:    0,
		Type:  schema.PocketType_POCKET_TYPE_DEBT_REDUCTION,
		Goals: GenerateGoals(goalCount),
	}
}

func GeneratePockets(pocketCount int) []*schema.Pocket {
	pockets := make([]*schema.Pocket, 0, pocketCount)
	for i := 0; i < pocketCount; i++ {
		pockets = append(pockets, GenerateSinglePocket())
	}

	return pockets
}

func GenerateBankAccount() *schema.BankAccount {
	return &schema.BankAccount{
		UserId:         uint64(GenerateRandomId(100000, 3000000)),
		Name:           GenerateRandomString(10),
		Number:         GenerateRandomString(10),
		Type:           schema.BankAccountType_BANK_ACCOUNT_TYPE_MANUAL,
		Balance:        float32(GenerateRandomId(1000, 100000)),
		Currency:       GenerateRandomString(10),
		CurrentFunds:   float64(GenerateRandomId(1000, 100000)),
		BalanceLimit:   uint64(GenerateRandomId(1000, 100000)),
		Pockets:        GeneratePockets(3),
		PlaidAccountId: GenerateRandomString(10),
		Subtype:        GenerateRandomString(10),
	}
}

func GenerateBankAccounts(bankAccountCount int) []*schema.BankAccount {
	bankAccounts := make([]*schema.BankAccount, 0, bankAccountCount)
	for i := 0; i < bankAccountCount; i++ {
		bankAccounts = append(bankAccounts, GenerateBankAccount())
	}

	return bankAccounts
}

func GenerateSingleCreditAccount() *schema.CreditAccount {
	return &schema.CreditAccount{
		UserId:                 GenerateRandomString(10),
		Name:                   GenerateRandomString(10),
		Number:                 GenerateRandomString(10),
		Type:                   GenerateRandomString(10),
		Balance:                float32(GenerateRandomId(1000, 100000)),
		CurrentFunds:           float64(GenerateRandomId(1000, 100000)),
		BalanceLimit:           uint64(GenerateRandomId(1000, 100000)),
		PlaidAccountId:         GenerateRandomString(10),
		Subtype:                GenerateRandomString(10),
		IsOverdue:              false,
		LastPaymentAmount:      float64(GenerateRandomId(1000, 100000)),
		LastPaymentDate:        GenerateRandomString(10),
		LastStatementIssueDate: GenerateRandomString(10),
		MinimumAmountDueDate:   float64(GenerateRandomId(1000, 100000)),
		NextPaymentDate:        GenerateRandomString(10),
		Aprs:                   []*schema.Apr{},
		LastStatementBalance:   float64(GenerateRandomId(1000, 100000)),
		MinimumPaymentAmount:   float64(GenerateRandomId(1000, 100000)),
		NextPaymentDueDate:     GenerateRandomString(10),
	}
}

func GenerateCreditAccounts(creditAccountCount int) []*schema.CreditAccount {
	creditAccounts := make([]*schema.CreditAccount, 0, creditAccountCount)
	for i := 0; i < creditAccountCount; i++ {
		creditAccounts = append(creditAccounts, GenerateSingleCreditAccount())
	}

	return creditAccounts
}

func GenerateSingleMortgage() *schema.MortgageAccount {
	return &schema.MortgageAccount{
		PlaidAccountId:              GenerateRandomString(10),
		AccountNumber:               GenerateRandomString(10),
		CurrentLateFee:              float64(GenerateRandomId(1000, 100000)),
		EscrowBalance:               float64(GenerateRandomId(1000, 100000)),
		HasPmi:                      false,
		HasPrepaymentPenalty:        false,
		LastPaymentAmount:           float64(GenerateRandomId(1000, 100000)),
		LastPaymentDate:             GenerateRandomString(10),
		LoanTerm:                    GenerateRandomString(10),
		LoanTypeDescription:         GenerateRandomString(10),
		MaturityDate:                GenerateRandomString(10),
		NextMonthlyPayment:          float64(GenerateRandomId(1000, 100000)),
		NextPaymentDueDate:          GenerateRandomString(10),
		OriginalPrincipalBalance:    float64(GenerateRandomId(1000, 100000)),
		OriginalPropertyValue:       float64(GenerateRandomId(1000, 100000)),
		OutstandingPrincipalBalance: float64(GenerateRandomId(1000, 100000)),
		PaymentAmount:               float64(GenerateRandomId(1000, 100000)),
		PaymentDate:                 GenerateRandomString(10),
		OriginationDate:             GenerateRandomString(10),
		OriginationPrincipalAmount:  float64(GenerateRandomId(1000, 100000)),
		PastDueAmount:               float64(GenerateRandomId(1000, 100000)),
		YtdInterestPaid:             float64(GenerateRandomId(1000, 100000)),
		YtdPrincipalPaid:            float64(GenerateRandomId(1000, 100000)),
		PropertyAddressCity:         GenerateRandomString(10),
		PropertyAddressState:        GenerateRandomString(10),
		PropertyAddressStreet:       GenerateRandomString(10),
		PropertyAddressPostalCode:   GenerateRandomString(10),
		PropertyRegion:              GenerateRandomString(10),
		PropertyCountry:             GenerateRandomString(10),
		InterestRatePercentage:      float64(GenerateRandomId(1000, 100000)),
		InterestRateType:            GenerateRandomString(10),
	}
}

func GenerateMortgageAccounts(mortgageAccountCount int) []*schema.MortgageAccount {
	mortgageAccounts := make([]*schema.MortgageAccount, 0, mortgageAccountCount)
	for i := 0; i < mortgageAccountCount; i++ {
		mortgageAccounts = append(mortgageAccounts, GenerateSingleMortgage())
	}

	return mortgageAccounts
}

func GenerateInvestmentHolding() *schema.InvesmentHolding {
	return &schema.InvesmentHolding{
		Name:                     GenerateRandomString(10),
		PlaidAccountId:           GenerateRandomString(10),
		CostBasis:                float64(GenerateRandomId(1000, 100000)),
		InstitutionPrice:         float64(GenerateRandomId(1000, 100000)),
		InstitutionPriceAsOf:     GenerateRandomString(10),
		InstitutionPriceDatetime: GenerateRandomString(10),
		InstitutionValue:         float64(GenerateRandomId(1000, 100000)),
		IsoCurrencyCode:          GenerateRandomString(10),
		Quantity:                 float64(GenerateRandomId(1000, 100000)),
		SecurityId:               GenerateRandomString(10),
		UnofficialCurrencyCode:   GenerateRandomString(10),
	}
}

func GenerateInvestmentHoldings(holdingCount int) []*schema.InvesmentHolding {
	holdings := make([]*schema.InvesmentHolding, 0, holdingCount)
	for i := 0; i < holdingCount; i++ {
		holdings = append(holdings, GenerateInvestmentHolding())
	}

	return holdings
}

func GenerateSingleInvestmentSecurity() *schema.InvestmentSecurity {
	return &schema.InvestmentSecurity{
		ClosePrice:             float64(GenerateRandomId(1000, 100000)),
		ClosePriceAsOf:         GenerateRandomString(10),
		Cusip:                  GenerateRandomString(10),
		InstitutionId:          GenerateRandomString(10),
		InstitutionSecurityId:  GenerateRandomString(10),
		IsCashEquivalent:       false,
		Isin:                   GenerateRandomString(10),
		IsoCurrencyCode:        GenerateRandomString(10),
		Name:                   GenerateRandomString(10),
		ProxySecurityId:        GenerateRandomString(10),
		SecurityId:             GenerateRandomString(10),
		Sedol:                  GenerateRandomString(10),
		TickerSymbol:           GenerateRandomString(10),
		Type:                   GenerateRandomString(10),
		UnofficialCurrencyCode: GenerateRandomString(10),
		UpdateDatetime:         GenerateRandomString(10),
	}
}

func GenerateInvestmentSecurities(securityCount int) []*schema.InvestmentSecurity {
	securities := make([]*schema.InvestmentSecurity, 0, securityCount)
	for i := 0; i < securityCount; i++ {
		securities = append(securities, GenerateSingleInvestmentSecurity())
	}

	return securities
}

func GenereateRandomUserProfileForTest() *schema.UserProfile {
	return &schema.UserProfile{
		UserId:       uint64(GenerateRandomId(100000, 3000000)),
		BankAccounts: GenerateBankAccounts(3),
		InvestmentAccounts: []*schema.InvestmentAccount{
			{
				Id:             0,
				UserId:         GenerateRandomString(10),
				Name:           GenerateRandomString(10),
				Number:         GenerateRandomString(10),
				Type:           GenerateRandomString(10),
				Balance:        float32(GenerateRandomId(1000, 100000)),
				CurrentFunds:   float64(GenerateRandomId(1000, 100000)),
				BalanceLimit:   uint64(GenerateRandomId(1000, 100000)),
				PlaidAccountId: GenerateRandomString(10),
				Subtype:        GenerateRandomString(10),
				Holdings:       GenerateInvestmentHoldings(10),
				Securities:     GenerateInvestmentSecurities(10),
			},
		},
		CreditAccounts:   GenerateCreditAccounts(3),
		MortgageAccounts: GenerateMortgageAccounts(3),
		PlaidAccessToken: GenerateRandomString(10),
		StripeCustomerId: GenerateRandomString(10),
	}
}
