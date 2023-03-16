package apiv1

// GetDatabaseSchemas returns a suite of database schemas
func GetDatabaseSchemas() []interface{} {
	var models = []interface{}{
		UserProfileORM{},
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
	}

	return models
}
