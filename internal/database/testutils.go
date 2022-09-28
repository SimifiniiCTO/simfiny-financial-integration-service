package database

import (
	"time"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"go.uber.org/zap"
)

const (
	DefaultMaxConnectionAttempts  int           = 3
	DefaultMaxRetriesPerOperation int           = 3
	DefaultRetryTimeout           time.Duration = 50 * time.Millisecond
	DefaultRetrySleepInterval     time.Duration = 25 * time.Millisecond
)

const (
	EMPTY = ""
)

var (
	Port                int    = 6000
	Host                string = "localhost"
	User                string = "financial_integration_svc"
	Password            string = "financial_integration_svc"
	Dbname              string = "financial_integration_svc"
	DefaultDbConnParams        = ConnectionParameters{
		Host:         Host,
		User:         User,
		Password:     Password,
		DatabaseName: Dbname,
		Port:         Port,
	}

	DefaultConnInitializationParams = ConnectionInitializationParams{
		ConnectionParams:       &DefaultDbConnParams,
		Logger:                 zap.L(),
		MaxConnectionAttempts:  DefaultMaxConnectionAttempts,
		MaxRetriesPerOperation: DefaultMaxRetriesPerOperation,
		RetryTimeOut:           DefaultRetryTimeout,
		RetrySleepInterval:     DefaultRetrySleepInterval,
	}
)

func generateRandomizedAccountWithRandomId() *schema.VirtualAccount {
	acct := generateRandomizedAccount()
	acct.Id = uint64(helper.GenerateRandomId(10000, 3000000))
	return acct
}

// generateRandomizedAccount generates a random account
func generateRandomizedAccount() *schema.VirtualAccount {
	return &schema.VirtualAccount{
		Id:          0,
		AccessToken: helper.GenerateRandomString(10),
		UserID:      uint64(helper.GenerateRandomId(10000, 3000000)),
		DepositAccountID: []*schema.DepositAccount{
			{
				Id:             0,
				PlaidAccountID: helper.GenerateRandomString(10),
				AccountSubtype: helper.GenerateRandomString(10),
				AccountType:    helper.GenerateRandomString(10),
				AccountName:    helper.GenerateRandomString(10),
				BalanceID: &schema.Balance{
					Id:             uint64(helper.GenerateRandomId(10000, 3000000)),
					AvailableFunds: 0,
					CurrentFunds:   320.76,
					CurrencyCode:   "USD",
					BalanceLimit:   0,
				},
			},
		},
		CreditAccountID: []*schema.CreditAccount{
			{
				Id:                     0,
				PlaidAccountID:         helper.GenerateRandomString(10),
				AccountSubtype:         helper.GenerateRandomString(10),
				AccountType:            helper.GenerateRandomString(10),
				IsOverdue:              false,
				LastPaymentAmount:      3000,
				LastPaymentDate:        helper.GenerateRandomString(10),
				LastStatementIssueDate: helper.GenerateRandomString(10),
				MinimumPaymentAmount:   2000,
				NextPaymentDueDate:     helper.GenerateRandomString(10),
				Aprs: []*schema.APR{
					{
						Id:                   0,
						APRPercentage:        20.5,
						APRType:              helper.GenerateRandomString(10),
						BalanceSubjectToAPR:  2000.5,
						InterestChargeAmount: 100.5,
					},
				},
			},
		},
		MortgageLoanAccountID: []*schema.MortgageLoanAccount{
			{
				Id:                         0,
				PlaidAccountID:             helper.GenerateRandomString(10),
				AccountSubtype:             helper.GenerateRandomString(10),
				AccountType:                helper.GenerateRandomString(10),
				AccountNumber:              helper.GenerateRandomString(10),
				CurrentLateFee:             300.5,
				EscrowBalance:              200,
				HasPmi:                     false,
				HasPrepaymentPenalty:       false,
				LastPaymentAmount:          100,
				LastPaymentDate:            helper.GenerateRandomString(10),
				LoanTerm:                   helper.GenerateRandomString(10),
				LoanTypeDescription:        helper.GenerateRandomString(10),
				MaturityDate:               helper.GenerateRandomString(10),
				NexMonthlyPayment:          300,
				NextPaymentDueDate:         helper.GenerateRandomString(10),
				OriginationDate:            helper.GenerateRandomString(10),
				OriginationPrincipalAmount: 40,
				PastDueAmount:              5000,
				InterestID:                 nil,
				AddressID:                  nil,
				InterestPaidYTD:            3000,
				PrincipalPaidYTD:           1800,
			},
		},
		StudentLoanAccountID: []*schema.StudentLoanAccount{
			{
				Id:                         0,
				PlaidAccountID:             helper.GenerateRandomString(10),
				AccountSubtype:             helper.GenerateRandomString(10),
				AccountType:                helper.GenerateRandomString(10),
				AccountNumber:              helper.GenerateRandomString(10),
				DisbursementDates:          nil,
				ExpectedPayoffDate:         helper.GenerateRandomString(10),
				Guarantor:                  helper.GenerateRandomString(10),
				InterestRatePercentage:     500.7,
				IsOverdue:                  false,
				LastPaymentAmount:          400,
				LastPaymentDate:            helper.GenerateRandomString(10),
				LastStatementIssueDate:     helper.GenerateRandomString(10),
				LoanStatementIssueDate:     helper.GenerateRandomString(10),
				LoanName:                   helper.GenerateRandomString(10),
				LoanRepaymentEndDate:       helper.GenerateRandomString(10),
				MinimumPaymentAmount:       70,
				NextPaymentDueDate:         helper.GenerateRandomString(10),
				OriginationDate:            helper.GenerateRandomString(10),
				OriginationPrincipalAmount: 90,
				OutstandingInterestAmount:  960,
				PaymentReferenceNumber:     helper.GenerateRandomString(10),
				SequenceNumber:             helper.GenerateRandomString(10),
				PslfID:                     nil,
				RepaymentPlan:              helper.GenerateRandomString(10),
				ServisorAddressID:          nil,
				InterestPaidYTD:            840,
				PrincipalPaidYTD:           120,
			},
		},
		InvestmentAccountID: []*schema.InvestmentAccount{
			{
				Id:             0,
				PlaidAccountID: helper.GenerateRandomString(10),
				AccountSubtype: helper.GenerateRandomString(10),
				AccountType:    helper.GenerateRandomString(10),
				AccountName:    helper.GenerateRandomString(10),
				BalanceID: &schema.Balance{
					Id:             0,
					AvailableFunds: 1090,
					CurrentFunds:   900,
					CurrencyCode:   "USD",
					BalanceLimit:   9990,
				},
				SecurityID: []*schema.Security{
					{
						Id:           0,
						CurrencyCode: "USD",
						SecurityID:   "GOOG",
						TickerSymbol: "GOOG",
						SecurityType: "Equity",
						SecurityName: "Google",
					},
				},
			},
		},
		Active: false,
	}
}
