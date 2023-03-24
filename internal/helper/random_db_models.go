package helper

import (
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GenerateBugetCategory generates a random budget category
func GenerateBugetCategory() *schema.Category {
	return &schema.Category{
		Id:          0,
		Name:        GenerateRandomString(10),
		Description: GenerateRandomString(10),
		Subcategories: []string{
			GenerateRandomString(10),
			GenerateRandomString(10),
			GenerateRandomString(10),
		},
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

// GenerateBudgets generates a slice of random budgets
func GenerateBudgets(budgetCount int) []*schema.Budget {
	budgets := make([]*schema.Budget, 0, budgetCount)
	for i := 0; i < budgetCount; i++ {
		budgets = append(budgets, GenerateSingleBudget())
	}
	return budgets
}

// GenerateSingleMilestone generates a single milestone
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

// GenerateMilestone generates a slice of milestones
func GenerateMilestones(milestoneCount int) []*schema.Milestone {
	milestones := make([]*schema.Milestone, 0, milestoneCount)
	for i := 0; i < milestoneCount; i++ {
		milestones = append(milestones, GenerateSingleMilestone(3))
	}

	return milestones
}

// GenerateSingleForecast generates a single forecast
func GenerateSingleForecast() *schema.Forecast {
	return &schema.Forecast{
		Id:                       0,
		ForecastedAmount:         GenerateRandomString(10),
		ForecastedCompletionDate: GenerateRandomString(10),
		VarianceAmount:           GenerateRandomString(10),
	}
}

// GenerateSingleGoal generates a single goal
func GenerateSingleGoal(numMilestones int) *schema.SmartGoal {
	return &schema.SmartGoal{
		Id:           0,
		UserId:       uint64(GenerateRandomId(10000, 30000)),
		Name:         GenerateRandomString(10),
		Description:  GenerateRandomString(10),
		IsCompleted:  false,
		GoalType:     schema.GoalType_GOAL_TYPE_DEBT,
		Duration:     GenerateRandomString(10),
		StartDate:    GenerateRandomString(10),
		EndDate:      GenerateRandomString(10),
		TargetAmount: GenerateRandomString(10),
		Milestones:   GenerateMilestones(numMilestones),
		Forecasts:    GenerateSingleForecast(),
	}
}

// GenerateGoals generates a slice of goals
func GenerateGoals(goalCount int) []*schema.SmartGoal {
	goals := make([]*schema.SmartGoal, 0, goalCount)
	for i := 0; i < goalCount; i++ {
		goals = append(goals, GenerateSingleGoal(10))
	}

	return goals
}

// GenerateSinglePocket generates a single pocket
func GenerateSinglePocket() *schema.Pocket {
	return &schema.Pocket{
		Id:    0,
		Type:  schema.PocketType_POCKET_TYPE_DEBT_REDUCTION,
		Goals: GenerateGoals(3),
	}
}

// GenerateSinglePocketWithGoals generates a single pocket with a specified number of goals
func GenerateSinglePocketWithGoals(goalCount int) *schema.Pocket {
	return &schema.Pocket{
		Id:    0,
		Type:  schema.PocketType_POCKET_TYPE_DEBT_REDUCTION,
		Goals: GenerateGoals(goalCount),
	}
}

// GeneratePockets generates a slice of pockets
func GeneratePockets(pocketCount int) []*schema.Pocket {
	pockets := make([]*schema.Pocket, 0, pocketCount)
	for i := 0; i < pocketCount; i++ {
		pockets = append(pockets, GenerateSinglePocket())
	}

	return pockets
}

// GenerateSingleBankAccount generates a single bank account
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

// GenerateBankAccounts generates a slice of bank accounts
func GenerateBankAccounts(bankAccountCount int) []*schema.BankAccount {
	bankAccounts := make([]*schema.BankAccount, 0, bankAccountCount)
	for i := 0; i < bankAccountCount; i++ {
		bankAccounts = append(bankAccounts, GenerateBankAccount())
	}

	return bankAccounts
}

// GenerateSingleCreditAccount generates a single credit account
func GenerateSingleCreditAccount() *schema.CreditAccount {
	return &schema.CreditAccount{
		UserId:                 uint64(GenerateRandomId(10000, 30000)),
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

// GenerateCreditAccounts generates a slice of credit accounts
func GenerateCreditAccounts(creditAccountCount int) []*schema.CreditAccount {
	creditAccounts := make([]*schema.CreditAccount, 0, creditAccountCount)
	for i := 0; i < creditAccountCount; i++ {
		creditAccounts = append(creditAccounts, GenerateSingleCreditAccount())
	}

	return creditAccounts
}

// GenerateSingleMortgage generates a single mortgage
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

// GenerateMortgageAccounts generates a slice of mortgage accounts
func GenerateMortgageAccounts(mortgageAccountCount int) []*schema.MortgageAccount {
	mortgageAccounts := make([]*schema.MortgageAccount, 0, mortgageAccountCount)
	for i := 0; i < mortgageAccountCount; i++ {
		mortgageAccounts = append(mortgageAccounts, GenerateSingleMortgage())
	}

	return mortgageAccounts
}

// GenerateSingleInvestmentAccount generates a single investment account
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

// GenerateInvestmentHoldings generates a slice of investment holdings
func GenerateInvestmentHoldings(holdingCount int) []*schema.InvesmentHolding {
	holdings := make([]*schema.InvesmentHolding, 0, holdingCount)
	for i := 0; i < holdingCount; i++ {
		holdings = append(holdings, GenerateInvestmentHolding())
	}

	return holdings
}

// GenerateSingleInvestmentAccount generates a single investment account
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

// GenerateInvestmentHoldings generates a slice of investment holdings
func GenerateInvestmentSecurities(securityCount int) []*schema.InvestmentSecurity {
	securities := make([]*schema.InvestmentSecurity, 0, securityCount)
	for i := 0; i < securityCount; i++ {
		securities = append(securities, GenerateSingleInvestmentSecurity())
	}

	return securities
}

// GenerateSingleInvestmentAccount generates a single investment account
func GenerateSingleInvestmentAccount() *schema.InvestmentAccount {
	return &schema.InvestmentAccount{
		UserId:         0,
		Name:           GenerateRandomString(10),
		Number:         GenerateRandomString(10),
		Type:           GenerateRandomString(10),
		Balance:        float32(GenerateRandomId(1000, 10000)),
		CurrentFunds:   float64(GenerateRandomId(1000, 10000)),
		BalanceLimit:   uint64(GenerateRandomId(1000, 10000)),
		PlaidAccountId: GenerateRandomString(10),
		Subtype:        GenerateRandomString(10),
		Holdings:       GenerateInvestmentHoldings(10),
		Securities:     GenerateInvestmentSecurities(10),
	}
}

// GenerateInvestmentAccounts generates a slice of investment accounts
func GenerateInvestmentAccounts(numInvestmentAcct int) []*schema.InvestmentAccount {
	investmentAccounts := make([]*schema.InvestmentAccount, 0, numInvestmentAcct)
	for i := 0; i < numInvestmentAcct; i++ {
		investmentAccounts = append(investmentAccounts, GenerateSingleInvestmentAccount())
	}

	return investmentAccounts
}

// GenerateToken generates a token
func GenerateToken() *schema.Token {
	return &schema.Token{
		AccessToken: GenerateRandomString(10),
		ItemId:      GenerateRandomString(10),
		KeyId:       GenerateRandomString(10),
		Version:     GenerateRandomString(10),
	}
}

// GeneratePlaidLink generates a plaid link
func GeneratePlaidLink() *schema.PlaidLink {
	return &schema.PlaidLink{
		Products: []string{
			"transactions",
			"auth",
			"identity",
			"balance",
			"assets",
		},
		WebhookUrl:      "https://webhook.site/4f3b1b1b-1b1b-4f3b-1b1b-4f3b1b1b4f3b",
		InstitutionId:   GenerateRandomString(10),
		InstitutionName: GenerateRandomString(10),
		UsePlaidSync:    false,
		ItemId:          GenerateRandomString(10),
	}
}

// GenerateLink generates a link
func GenerateLink(linkType schema.LinkType) *schema.Link {
	return &schema.Link{
		Id:                        0,
		LinkStatus:                schema.LinkStatus_LINK_STATUS_SUCCESS,
		PlaidLink:                 GeneratePlaidLink(),
		PlaidNewAccountsAvailable: false,
		ExpirationDate:            GenerateRandomString(10),
		InstitutionName:           GenerateRandomString(10),
		CustomInstitutionName:     GenerateRandomString(10),
		Description:               GenerateRandomString(10),
		LastManualSync:            GenerateRandomString(10),
		LastSuccessfulUpdate:      GenerateRandomString(10),
		Token:                     GenerateToken(),
		BankAccounts:              GenerateBankAccounts(4),
		InvestmentAccounts:        GenerateInvestmentAccounts(4),
		CreditAccounts:            GenerateCreditAccounts(10),
		MortgageAccounts:          GenerateMortgageAccounts(10),
		PlaidInstitutionId:        GenerateRandomString(10),
		LinkType:                  linkType,
	}
}

// GenerateSingleSubscription generates a single subscription
func GenerateSingleSubscription() *schema.StripeSubscription {
	return &schema.StripeSubscription{
		StripeSubscriptionId:          GenerateRandomString(20),
		StripeSubscriptionStatus:      schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_PAST_DUE,
		StripeSubscriptionActiveUntil: GenerateRandomString(20),
		StripeWebhookLatestTimestamp:  GenerateRandomString(20),
	}
}

// GenerateStripeSubscriptions generates a slice of subscriptions
func GenerateStripeSubscriptions(subscriptionCount int) []*schema.StripeSubscription {
	subscriptions := make([]*schema.StripeSubscription, 0, subscriptionCount)
	for i := 0; i < subscriptionCount; i++ {
		subscriptions = append(subscriptions, GenerateSingleSubscription())
	}

	return subscriptions
}

// GenereateRandomUserProfileForTest generates a random user profile for testing
func GenereateRandomUserProfileForTest() *schema.UserProfile {
	return &schema.UserProfile{
		UserId:              uint64(GenerateRandomId(100000, 3000000)),
		StripeSubscriptions: GenerateSingleSubscription(),
		Link: []*schema.Link{
			GenerateLink(schema.LinkType_LINK_TYPE_MANUAL),
			GenerateLink(schema.LinkType_LINK_TYPE_PLAID),
		},
	}
}
