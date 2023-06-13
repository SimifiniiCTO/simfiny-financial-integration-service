package apiv1

// GetDatabaseSchemas returns a suite of database schemas
func GetDatabaseSchemas() []interface{} {
	var models = []interface{}{
		LinkORM{},
		TokenORM{},
		PlaidLinkORM{},
		UserProfileORM{},
		StripeSubscriptionORM{},
		StudentLoanAccountORM{},
		CreditAccountORM{},
		MortgageAccountORM{},
		InvestmentAccountORM{},
		BankAccountORM{},
		PocketORM{},
		SmartGoalORM{},
		ForecastORM{},
		MilestoneORM{},
		BudgetORM{},
		CategoryORM{},
		InvesmentHoldingORM{},
		InvestmentSecurityORM{},
		AprORM{},
		PlaidSyncORM{},
	}

	return models
}

func GetClickhouseSchemas() []interface{} {
	var models = []interface{}{
		TransactionORM{},
		ReOccuringTransactionORM{},
	}
	return models
}
