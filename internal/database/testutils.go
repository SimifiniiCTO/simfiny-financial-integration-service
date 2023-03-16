package database

import (
	"time"

	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
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

func genereateRandomUserProfileForTest() *schema.UserProfile {
	return &schema.UserProfile{
		UserId: uint64(helper.GenerateRandomId(100000, 3000000)),
		BankAccounts: []*schema.BankAccount{
			{
				UserId:       helper.GenerateRandomString(10),
				Name:         helper.GenerateRandomString(10),
				Number:       helper.GenerateRandomString(10),
				Type:         helper.GenerateRandomString(10),
				Balance:      float32(helper.GenerateRandomId(1000, 100000)),
				Currency:     helper.GenerateRandomString(10),
				CurrentFunds: float64(helper.GenerateRandomId(1000, 100000)),
				BalanceLimit: uint64(helper.GenerateRandomId(1000, 100000)),
				Pockets: []*schema.Pocket{
					{
						Goals: []*schema.SmartGoal{
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_DEBT,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
						},
						Type: schema.PocketType_POCKET_TYPE_DEBT_REDUCTION,
					},
					{
						Goals: []*schema.SmartGoal{
							{

								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_EXPENSE,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_SAVINGS,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_INVESTMENT,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
						},
						Type: schema.PocketType_POCKET_TYPE_EMERGENCY_FUND,
					},
					{
						Goals: []*schema.SmartGoal{
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_EXPENSE,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_SAVINGS,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_INVESTMENT,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
						},
						Type: schema.PocketType_POCKET_TYPE_FUN_MONEY,
					},
					{
						Goals: []*schema.SmartGoal{
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_EXPENSE,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_SAVINGS,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
							{
								UserId:        helper.GenerateRandomString(10),
								Name:          helper.GenerateRandomString(10),
								Description:   helper.GenerateRandomString(10),
								IsCompleted:   false,
								GoalType:      schema.GoalType_GOAL_TYPE_INVESTMENT,
								Duration:      helper.GenerateRandomString(10),
								StartDate:     helper.GenerateRandomString(10),
								EndDate:       helper.GenerateRandomString(10),
								TargetAmount:  helper.GenerateRandomString(10),
								CurrentAmount: helper.GenerateRandomString(10),
							},
						},
						Type: schema.PocketType_POCKET_TYPE_DISCRETIONARY_SPENDING,
					},
				},
				PlaidAccountId: helper.GenerateRandomString(10),
				Subtype:        helper.GenerateRandomString(10),
			},
		},
		InvestmentAccounts: []*schema.InvestmentAccount{
			{
				UserId:       helper.GenerateRandomString(10),
				Name:         helper.GenerateRandomString(10),
				Number:       helper.GenerateRandomString(10),
				Type:         helper.GenerateRandomString(10),
				Balance:      float32(helper.GenerateRandomId(1000, 100000)),
				CurrentFunds: float64(helper.GenerateRandomId(1000, 100000)),
				BalanceLimit: uint64(helper.GenerateRandomId(1000, 100000)),
			},
		},
		CreditAccounts: []*schema.CreditAccount{
			{
				UserId:                 helper.GenerateRandomString(10),
				Name:                   helper.GenerateRandomString(10),
				Number:                 helper.GenerateRandomString(10),
				Type:                   helper.GenerateRandomString(10),
				Balance:                float32(helper.GenerateRandomId(1000, 100000)),
				CurrentFunds:           float64(helper.GenerateRandomId(1000, 100000)),
				BalanceLimit:           uint64(helper.GenerateRandomId(1000, 100000)),
				PlaidAccountId:         helper.GenerateRandomString(10),
				Subtype:                helper.GenerateRandomString(10),
				IsOverdue:              false,
				LastPaymentAmount:      float64(helper.GenerateRandomId(1000, 100000)),
				LastPaymentDate:        helper.GenerateRandomString(10),
				LastStatementIssueDate: helper.GenerateRandomString(10),
				MinimumAmountDueDate:   float64(helper.GenerateRandomId(1000, 100000)),
				NextPaymentDate:        helper.GenerateRandomString(10),
				Aprs:                   []*schema.Apr{},
				LastStatementBalance:   float64(helper.GenerateRandomId(1000, 100000)),
				MinimumPaymentAmount:   float64(helper.GenerateRandomId(1000, 100000)),
				NextPaymentDueDate:     helper.GenerateRandomString(10),
			},
		},
		MortgageAccounts: []*schema.MortgageAccount{
			{
				PlaidAccountId:              helper.GenerateRandomString(10),
				AccountNumber:               helper.GenerateRandomString(10),
				CurrentLateFee:              float64(helper.GenerateRandomId(1000, 100000)),
				EscrowBalance:               float64(helper.GenerateRandomId(1000, 100000)),
				HasPmi:                      false,
				HasPrepaymentPenalty:        false,
				LastPaymentAmount:           float64(helper.GenerateRandomId(1000, 100000)),
				LastPaymentDate:             helper.GenerateRandomString(10),
				LoanTerm:                    helper.GenerateRandomString(10),
				LoanTypeDescription:         helper.GenerateRandomString(10),
				MaturityDate:                helper.GenerateRandomString(10),
				NextMonthlyPayment:          float64(helper.GenerateRandomId(1000, 100000)),
				NextPaymentDueDate:          helper.GenerateRandomString(10),
				OriginalPrincipalBalance:    float64(helper.GenerateRandomId(1000, 100000)),
				OriginalPropertyValue:       float64(helper.GenerateRandomId(1000, 100000)),
				OutstandingPrincipalBalance: float64(helper.GenerateRandomId(1000, 100000)),
				PaymentAmount:               float64(helper.GenerateRandomId(1000, 100000)),
				PaymentDate:                 helper.GenerateRandomString(10),
				OriginationDate:             helper.GenerateRandomString(10),
				OriginationPrincipalAmount:  float64(helper.GenerateRandomId(1000, 100000)),
				PastDueAmount:               float64(helper.GenerateRandomId(1000, 100000)),
				YtdInterestPaid:             float64(helper.GenerateRandomId(1000, 100000)),
				YtdPrincipalPaid:            float64(helper.GenerateRandomId(1000, 100000)),
				PropertyAddressCity:         helper.GenerateRandomString(10),
				PropertyAddressState:        helper.GenerateRandomString(10),
				PropertyAddressStreet:       helper.GenerateRandomString(10),
				PropertyAddressPostalCode:   helper.GenerateRandomString(10),
				PropertyRegion:              helper.GenerateRandomString(10),
				PropertyCountry:             helper.GenerateRandomString(10),
				InterestRatePercentage:      float64(helper.GenerateRandomId(1000, 100000)),
				InterestRateType:            helper.GenerateRandomString(10),
			},
			{
				PlaidAccountId:              helper.GenerateRandomString(10),
				AccountNumber:               helper.GenerateRandomString(10),
				CurrentLateFee:              float64(helper.GenerateRandomId(1000, 100000)),
				EscrowBalance:               float64(helper.GenerateRandomId(1000, 100000)),
				HasPmi:                      false,
				HasPrepaymentPenalty:        false,
				LastPaymentAmount:           float64(helper.GenerateRandomId(1000, 100000)),
				LastPaymentDate:             helper.GenerateRandomString(10),
				LoanTerm:                    helper.GenerateRandomString(10),
				LoanTypeDescription:         helper.GenerateRandomString(10),
				MaturityDate:                helper.GenerateRandomString(10),
				NextMonthlyPayment:          float64(helper.GenerateRandomId(1000, 100000)),
				NextPaymentDueDate:          helper.GenerateRandomString(10),
				OriginalPrincipalBalance:    float64(helper.GenerateRandomId(1000, 100000)),
				OriginalPropertyValue:       float64(helper.GenerateRandomId(1000, 100000)),
				OutstandingPrincipalBalance: float64(helper.GenerateRandomId(1000, 100000)),
				PaymentAmount:               float64(helper.GenerateRandomId(1000, 100000)),
				PaymentDate:                 helper.GenerateRandomString(10),
				OriginationDate:             helper.GenerateRandomString(10),
				OriginationPrincipalAmount:  float64(helper.GenerateRandomId(1000, 100000)),
				PastDueAmount:               float64(helper.GenerateRandomId(1000, 100000)),
				YtdInterestPaid:             float64(helper.GenerateRandomId(1000, 100000)),
				YtdPrincipalPaid:            float64(helper.GenerateRandomId(1000, 100000)),
				PropertyAddressCity:         helper.GenerateRandomString(10),
				PropertyAddressState:        helper.GenerateRandomString(10),
				PropertyAddressStreet:       helper.GenerateRandomString(10),
				PropertyAddressPostalCode:   helper.GenerateRandomString(10),
				PropertyRegion:              helper.GenerateRandomString(10),
				PropertyCountry:             helper.GenerateRandomString(10),
				InterestRatePercentage:      float64(helper.GenerateRandomId(1000, 100000)),
				InterestRateType:            helper.GenerateRandomString(10),
			},
		},
		PlaidAccessToken: helper.GenerateRandomString(10),
		StripeCustomerId: helper.GenerateRandomString(10),
	}
}
