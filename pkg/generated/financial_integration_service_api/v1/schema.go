package financial_integration_service_apiv1

// GetDatabaseSchemas returns a suite of database schemas
func GetDatabaseSchemas() []interface{} {
	var models = []interface{}{
		LinkORM{},
		ActionableInsightORM{},
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
		InvestmentTransactionORM{},
	}
	return models
}

func GetClickhouseTableNameToSchemaMap() map[string]interface{} {
	schemaMap := make(map[string]interface{}, 0)
	schemaMap[TransactionORM{}.TableName()] = TransactionORM{}
	schemaMap[ReOccuringTransactionORM{}.TableName()] = ReOccuringTransactionORM{}
	schemaMap[InvestmentTransactionORM{}.TableName()] = InvestmentTransactionORM{}
	return schemaMap
}
