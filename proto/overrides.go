package proto

func Schemas() []interface{} {
	schemas := make([]interface{}, 0)
	schemas = append(schemas,
		VirtualAccountORM{},
		SecurityORM{},
		AddressORM{},
		BalanceORM{},
		CreditAccountORM{},
		DepositAccountORM{},
		InvestmentAccountORM{},
		MortgageLoanAccountORM{},
		InterestORM{},
		PslfORM{},
		StudentLoanAccountORM{},
		APRORM{})
	return schemas
}
