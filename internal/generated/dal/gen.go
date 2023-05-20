// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                     = new(Query)
	AprORM                *aprORM
	BankAccountORM        *bankAccountORM
	BudgetORM             *budgetORM
	CategoryORM           *categoryORM
	CreditAccountORM      *creditAccountORM
	ForecastORM           *forecastORM
	InvesmentHoldingORM   *invesmentHoldingORM
	InvestmentAccountORM  *investmentAccountORM
	InvestmentSecurityORM *investmentSecurityORM
	LinkORM               *linkORM
	MilestoneORM          *milestoneORM
	MortgageAccountORM    *mortgageAccountORM
	PlaidLinkORM          *plaidLinkORM
	PlaidSyncORM          *plaidSyncORM
	PocketORM             *pocketORM
	SmartGoalORM          *smartGoalORM
	StripeSubscriptionORM *stripeSubscriptionORM
	StudentLoanAccountORM *studentLoanAccountORM
	TokenORM              *tokenORM
	TransactionORM        *transactionORM
	UserProfileORM        *userProfileORM
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AprORM = &Q.AprORM
	BankAccountORM = &Q.BankAccountORM
	BudgetORM = &Q.BudgetORM
	CategoryORM = &Q.CategoryORM
	CreditAccountORM = &Q.CreditAccountORM
	ForecastORM = &Q.ForecastORM
	InvesmentHoldingORM = &Q.InvesmentHoldingORM
	InvestmentAccountORM = &Q.InvestmentAccountORM
	InvestmentSecurityORM = &Q.InvestmentSecurityORM
	LinkORM = &Q.LinkORM
	MilestoneORM = &Q.MilestoneORM
	MortgageAccountORM = &Q.MortgageAccountORM
	PlaidLinkORM = &Q.PlaidLinkORM
	PlaidSyncORM = &Q.PlaidSyncORM
	PocketORM = &Q.PocketORM
	SmartGoalORM = &Q.SmartGoalORM
	StripeSubscriptionORM = &Q.StripeSubscriptionORM
	StudentLoanAccountORM = &Q.StudentLoanAccountORM
	TokenORM = &Q.TokenORM
	TransactionORM = &Q.TransactionORM
	UserProfileORM = &Q.UserProfileORM
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                    db,
		AprORM:                newAprORM(db, opts...),
		BankAccountORM:        newBankAccountORM(db, opts...),
		BudgetORM:             newBudgetORM(db, opts...),
		CategoryORM:           newCategoryORM(db, opts...),
		CreditAccountORM:      newCreditAccountORM(db, opts...),
		ForecastORM:           newForecastORM(db, opts...),
		InvesmentHoldingORM:   newInvesmentHoldingORM(db, opts...),
		InvestmentAccountORM:  newInvestmentAccountORM(db, opts...),
		InvestmentSecurityORM: newInvestmentSecurityORM(db, opts...),
		LinkORM:               newLinkORM(db, opts...),
		MilestoneORM:          newMilestoneORM(db, opts...),
		MortgageAccountORM:    newMortgageAccountORM(db, opts...),
		PlaidLinkORM:          newPlaidLinkORM(db, opts...),
		PlaidSyncORM:          newPlaidSyncORM(db, opts...),
		PocketORM:             newPocketORM(db, opts...),
		SmartGoalORM:          newSmartGoalORM(db, opts...),
		StripeSubscriptionORM: newStripeSubscriptionORM(db, opts...),
		StudentLoanAccountORM: newStudentLoanAccountORM(db, opts...),
		TokenORM:              newTokenORM(db, opts...),
		TransactionORM:        newTransactionORM(db, opts...),
		UserProfileORM:        newUserProfileORM(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AprORM                aprORM
	BankAccountORM        bankAccountORM
	BudgetORM             budgetORM
	CategoryORM           categoryORM
	CreditAccountORM      creditAccountORM
	ForecastORM           forecastORM
	InvesmentHoldingORM   invesmentHoldingORM
	InvestmentAccountORM  investmentAccountORM
	InvestmentSecurityORM investmentSecurityORM
	LinkORM               linkORM
	MilestoneORM          milestoneORM
	MortgageAccountORM    mortgageAccountORM
	PlaidLinkORM          plaidLinkORM
	PlaidSyncORM          plaidSyncORM
	PocketORM             pocketORM
	SmartGoalORM          smartGoalORM
	StripeSubscriptionORM stripeSubscriptionORM
	StudentLoanAccountORM studentLoanAccountORM
	TokenORM              tokenORM
	TransactionORM        transactionORM
	UserProfileORM        userProfileORM
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		AprORM:                q.AprORM.clone(db),
		BankAccountORM:        q.BankAccountORM.clone(db),
		BudgetORM:             q.BudgetORM.clone(db),
		CategoryORM:           q.CategoryORM.clone(db),
		CreditAccountORM:      q.CreditAccountORM.clone(db),
		ForecastORM:           q.ForecastORM.clone(db),
		InvesmentHoldingORM:   q.InvesmentHoldingORM.clone(db),
		InvestmentAccountORM:  q.InvestmentAccountORM.clone(db),
		InvestmentSecurityORM: q.InvestmentSecurityORM.clone(db),
		LinkORM:               q.LinkORM.clone(db),
		MilestoneORM:          q.MilestoneORM.clone(db),
		MortgageAccountORM:    q.MortgageAccountORM.clone(db),
		PlaidLinkORM:          q.PlaidLinkORM.clone(db),
		PlaidSyncORM:          q.PlaidSyncORM.clone(db),
		PocketORM:             q.PocketORM.clone(db),
		SmartGoalORM:          q.SmartGoalORM.clone(db),
		StripeSubscriptionORM: q.StripeSubscriptionORM.clone(db),
		StudentLoanAccountORM: q.StudentLoanAccountORM.clone(db),
		TokenORM:              q.TokenORM.clone(db),
		TransactionORM:        q.TransactionORM.clone(db),
		UserProfileORM:        q.UserProfileORM.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		AprORM:                q.AprORM.replaceDB(db),
		BankAccountORM:        q.BankAccountORM.replaceDB(db),
		BudgetORM:             q.BudgetORM.replaceDB(db),
		CategoryORM:           q.CategoryORM.replaceDB(db),
		CreditAccountORM:      q.CreditAccountORM.replaceDB(db),
		ForecastORM:           q.ForecastORM.replaceDB(db),
		InvesmentHoldingORM:   q.InvesmentHoldingORM.replaceDB(db),
		InvestmentAccountORM:  q.InvestmentAccountORM.replaceDB(db),
		InvestmentSecurityORM: q.InvestmentSecurityORM.replaceDB(db),
		LinkORM:               q.LinkORM.replaceDB(db),
		MilestoneORM:          q.MilestoneORM.replaceDB(db),
		MortgageAccountORM:    q.MortgageAccountORM.replaceDB(db),
		PlaidLinkORM:          q.PlaidLinkORM.replaceDB(db),
		PlaidSyncORM:          q.PlaidSyncORM.replaceDB(db),
		PocketORM:             q.PocketORM.replaceDB(db),
		SmartGoalORM:          q.SmartGoalORM.replaceDB(db),
		StripeSubscriptionORM: q.StripeSubscriptionORM.replaceDB(db),
		StudentLoanAccountORM: q.StudentLoanAccountORM.replaceDB(db),
		TokenORM:              q.TokenORM.replaceDB(db),
		TransactionORM:        q.TransactionORM.replaceDB(db),
		UserProfileORM:        q.UserProfileORM.replaceDB(db),
	}
}

type queryCtx struct {
	AprORM                IAprORMDo
	BankAccountORM        IBankAccountORMDo
	BudgetORM             IBudgetORMDo
	CategoryORM           ICategoryORMDo
	CreditAccountORM      ICreditAccountORMDo
	ForecastORM           IForecastORMDo
	InvesmentHoldingORM   IInvesmentHoldingORMDo
	InvestmentAccountORM  IInvestmentAccountORMDo
	InvestmentSecurityORM IInvestmentSecurityORMDo
	LinkORM               ILinkORMDo
	MilestoneORM          IMilestoneORMDo
	MortgageAccountORM    IMortgageAccountORMDo
	PlaidLinkORM          IPlaidLinkORMDo
	PlaidSyncORM          IPlaidSyncORMDo
	PocketORM             IPocketORMDo
	SmartGoalORM          ISmartGoalORMDo
	StripeSubscriptionORM IStripeSubscriptionORMDo
	StudentLoanAccountORM IStudentLoanAccountORMDo
	TokenORM              ITokenORMDo
	TransactionORM        ITransactionORMDo
	UserProfileORM        IUserProfileORMDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AprORM:                q.AprORM.WithContext(ctx),
		BankAccountORM:        q.BankAccountORM.WithContext(ctx),
		BudgetORM:             q.BudgetORM.WithContext(ctx),
		CategoryORM:           q.CategoryORM.WithContext(ctx),
		CreditAccountORM:      q.CreditAccountORM.WithContext(ctx),
		ForecastORM:           q.ForecastORM.WithContext(ctx),
		InvesmentHoldingORM:   q.InvesmentHoldingORM.WithContext(ctx),
		InvestmentAccountORM:  q.InvestmentAccountORM.WithContext(ctx),
		InvestmentSecurityORM: q.InvestmentSecurityORM.WithContext(ctx),
		LinkORM:               q.LinkORM.WithContext(ctx),
		MilestoneORM:          q.MilestoneORM.WithContext(ctx),
		MortgageAccountORM:    q.MortgageAccountORM.WithContext(ctx),
		PlaidLinkORM:          q.PlaidLinkORM.WithContext(ctx),
		PlaidSyncORM:          q.PlaidSyncORM.WithContext(ctx),
		PocketORM:             q.PocketORM.WithContext(ctx),
		SmartGoalORM:          q.SmartGoalORM.WithContext(ctx),
		StripeSubscriptionORM: q.StripeSubscriptionORM.WithContext(ctx),
		StudentLoanAccountORM: q.StudentLoanAccountORM.WithContext(ctx),
		TokenORM:              q.TokenORM.WithContext(ctx),
		TransactionORM:        q.TransactionORM.WithContext(ctx),
		UserProfileORM:        q.UserProfileORM.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
