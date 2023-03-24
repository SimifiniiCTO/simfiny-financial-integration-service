package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

// createUserProfileAndBankAccountForTest creates a user profile and bank account for testing purposes
func createUserProfileAndBankAccountForTest(t *testing.T, ctx context.Context, conn *Db) (*schema.UserProfile, *schema.BankAccount) {
	sampleUserID := uint64(helper.GenerateRandomId(10000, 3000000))
	userProfile, err := conn.CreateUserProfile(ctx, &schema.UserProfile{
		UserId: sampleUserID,
	})
	assert.Nil(t, err)

	// generate bank account
	bankAccount, err := conn.CreateBankAccount(ctx, sampleUserID, helper.GenerateBankAccount())
	assert.Nil(t, err)

	return userProfile, bankAccount
}

func generateBankAccount() *schema.BankAccount {
	return &schema.BankAccount{
		UserId:       uint64(helper.GenerateRandomId(100000, 3000000)),
		Name:         helper.GenerateRandomString(10),
		Number:       helper.GenerateRandomString(10),
		Type:         schema.BankAccountType_BANK_ACCOUNT_TYPE_MANUAL,
		Balance:      float32(helper.GenerateRandomId(1000, 100000)),
		Currency:     helper.GenerateRandomString(10),
		CurrentFunds: float64(helper.GenerateRandomId(1000, 100000)),
		BalanceLimit: uint64(helper.GenerateRandomId(1000, 100000)),
		Pockets: []*schema.Pocket{
			{
				Goals: []*schema.SmartGoal{
					{
						UserId:        uint64(helper.GenerateRandomId(100000, 3000000)),
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

						UserId:        uint64(helper.GenerateRandomId(10000, 3000000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
						UserId:        uint64(helper.GenerateRandomId(10000, 300000)),
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
	}
}
