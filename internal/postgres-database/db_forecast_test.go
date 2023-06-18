package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestDb_CreateForecast(t *testing.T) {
	type args struct {
		ctx      context.Context
		goalID   uint64
		forecast *schema.Forecast
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		shouldCreateGoal bool
	}{
		{
			name: "create forecast",
			args: args{
				ctx:      context.Background(),
				goalID:   uint64(helper.GenerateRandomId(10000, 300000)),
				forecast: helper.GenerateSingleForecast(),
			},
			shouldCreateGoal: true,
		},
		{
			name: "create forecast for non-existent goal",
			args: args{
				ctx:      context.Background(),
				goalID:   0,
				forecast: helper.GenerateSingleForecast(),
			},
			shouldCreateGoal: false,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateGoal {
				sampleUserID := uint64(helper.GenerateRandomId(10000, 3000000))
				_, err := conn.CreateUserProfile(tt.args.ctx, &schema.UserProfile{
					UserId: sampleUserID,
				})

				assert.Nil(t, err)
				// generate bank account
				bankAcct, err := conn.CreateBankAccount(tt.args.ctx, sampleUserID, helper.GenerateBankAccount())
				assert.Nil(t, err)

				// obtain first goal from the created bank account
				goal := bankAcct.Pockets[0].Goals[0] // guaranteed to be populated
				tt.args.goalID = goal.Id
			}

			got, err := conn.CreateForecast(tt.args.ctx, tt.args.goalID, tt.args.forecast)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateForecast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestDb_GetForecast(t *testing.T) {
	type args struct {
		ctx        context.Context
		goalID     uint64
		forecast   *schema.Forecast
		forecastID uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		shouldCreateGoal bool
	}{
		{
			name: "get a forecast that already exists",
			args: args{
				ctx:        context.Background(),
				goalID:     uint64(helper.GenerateRandomId(10000, 300000)),
				forecast:   helper.GenerateSingleForecast(),
				forecastID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateGoal: true,
		},
		{
			name: "get a forecast that does not exist",
			args: args{
				ctx:        context.Background(),
				goalID:     0,
				forecast:   helper.GenerateSingleForecast(),
				forecastID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateGoal: false,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateGoal {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				goal := bankAcct.Pockets[0].Goals[0]

				// create the forecast
				forecast, err := conn.CreateForecast(tt.args.ctx, goal.Id, tt.args.forecast)
				assert.Nil(t, err)

				// populate the args
				tt.args.forecast = forecast
				tt.args.goalID = goal.Id
				tt.args.forecastID = forecast.Id
			}

			got, err := conn.GetForecast(tt.args.ctx, tt.args.forecastID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetForecast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, tt.args.forecast, got)
			}
		})
	}
}

func TestDb_UpdateForecast(t *testing.T) {
	type args struct {
		ctx        context.Context
		goalID     uint64
		forecast   *schema.Forecast
		forecastID uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		shouldCreateGoal bool
	}{
		{
			name: "update a forecast that already exists",
			args: args{
				ctx:        context.Background(),
				goalID:     uint64(helper.GenerateRandomId(10000, 300000)),
				forecast:   helper.GenerateSingleForecast(),
				forecastID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateGoal: true,
		},
		{
			name: "update a forecast that does not exist",
			args: args{
				ctx:        context.Background(),
				goalID:     0,
				forecast:   helper.GenerateSingleForecast(),
				forecastID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateGoal: false,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const forecastedAmmount = "325782375872"
			if tt.shouldCreateGoal {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				goal := bankAcct.Pockets[0].Goals[0]

				// create the forecast
				forecast, err := conn.CreateForecast(tt.args.ctx, goal.Id, tt.args.forecast)
				assert.Nil(t, err)

				// populate the args
				forecast.ForecastedAmount = forecastedAmmount
				tt.args.forecast = forecast
				tt.args.goalID = goal.Id
				tt.args.forecastID = forecast.Id
			}

			if err := conn.UpdateForecast(tt.args.ctx, tt.args.forecast); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateForecast() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				forecast, err := conn.GetForecast(tt.args.ctx, tt.args.forecastID)
				assert.Nil(t, err)
				assert.Equal(t, forecastedAmmount, forecast.ForecastedAmount)
			}
		})
	}
}

func TestDb_DeleteForecast(t *testing.T) {
	type args struct {
		ctx        context.Context
		goalID     uint64
		forecast   *schema.Forecast
		forecastID uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		shouldCreateGoal bool
	}{
		{
			name: "delete a forecast that already exists",
			args: args{
				ctx:        context.Background(),
				goalID:     uint64(helper.GenerateRandomId(10000, 300000)),
				forecast:   helper.GenerateSingleForecast(),
				forecastID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateGoal: true,
		},
		{
			name: "delete a forecast that does not exist",
			args: args{
				ctx:        context.Background(),
				goalID:     0,
				forecast:   helper.GenerateSingleForecast(),
				forecastID: uint64(helper.GenerateRandomId(10000, 300000)),
			},
			shouldCreateGoal: false,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCreateGoal {
				_, bankAcct := createUserProfileAndBankAccountForTest(t, tt.args.ctx, conn)
				goal := bankAcct.Pockets[0].Goals[0]

				// create the forecast
				forecast, err := conn.CreateForecast(tt.args.ctx, goal.Id, tt.args.forecast)
				assert.Nil(t, err)

				// populate the args
				tt.args.forecast = forecast
				tt.args.goalID = goal.Id
				tt.args.forecastID = forecast.Id
			}

			if err := conn.DeleteForecast(tt.args.ctx, tt.args.forecastID); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateForecast() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, err := conn.GetForecast(tt.args.ctx, tt.args.forecastID)
				assert.NotNil(t, err)
			}
		})
	}
}
