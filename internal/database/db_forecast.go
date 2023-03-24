package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"gorm.io/gen/field"
)

// CreateForecast implements DatabaseOperations
func (db *Db) CreateForecast(ctx context.Context, goalID uint64, forecast *schema.Forecast) (*schema.Forecast, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-forecast"); span != nil {
		defer span.End()
	}

	if forecast == nil {
		return nil, fmt.Errorf("forecast must be provided")
	}

	if goalID == 0 {
		return nil, fmt.Errorf("goal id must be provided. got: %d", goalID)
	}

	// ensure the goal exists
	g := db.QueryOperator.SmartGoalORM
	goal, err := g.WithContext(ctx).Where(g.Id.Eq(goalID)).First()
	if err != nil {
		return nil, fmt.Errorf("goal with id %d does not exist", goalID)
	}

	// convert schema to orm
	forecastOrm, err := forecast.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// save the forecast to the goal
	if err := g.Forecasts.Model(goal).Append(&forecastOrm); err != nil {
		return nil, err
	}

	result, err := g.WithContext(ctx).Where(g.Id.Eq(goalID)).Updates(goal)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// convert orm to schema
	record, err := forecastOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// GetForecast implements DatabaseOperations
func (db *Db) GetForecast(ctx context.Context, forecastID uint64) (*schema.Forecast, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-forecast"); span != nil {
		defer span.End()
	}

	if forecastID == 0 {
		return nil, fmt.Errorf("forecast id must be provided. got: %d", forecastID)
	}

	// ensure the forecast exists
	f := db.QueryOperator.ForecastORM
	forecastOrm, err := f.
		WithContext(ctx).
		Where(f.Id.Eq(forecastID)).
		First()
	if err != nil {
		return nil, fmt.Errorf("forecast with id %d does not exist", forecastID)
	}

	// convert orm to schema
	record, err := forecastOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// UpdateForecast implements DatabaseOperations
func (db *Db) UpdateForecast(ctx context.Context, forecast *schema.Forecast) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-forecast"); span != nil {
		defer span.End()
	}

	if forecast == nil {
		return fmt.Errorf("forecast must be provided")
	}

	if forecast.Id == 0 {
		return fmt.Errorf("forecast id must be provided. got: %d", forecast.Id)
	}

	// ensure the forecast exists
	f := db.QueryOperator.ForecastORM
	if _, err := f.WithContext(ctx).Where(f.Id.Eq(forecast.Id)).First(); err != nil {
		return fmt.Errorf("forecast with id %d does not exist", forecast.Id)
	}

	// convert schema to orm
	forecastOrm, err := forecast.ToORM(ctx)
	if err != nil {
		return err
	}

	result, err := f.WithContext(ctx).Where(f.Id.Eq(forecast.Id)).Updates(&forecastOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// DeleteForecast implements DatabaseOperations
func (db *Db) DeleteForecast(ctx context.Context, forecastID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-forecast"); span != nil {
		defer span.End()
	}

	if forecastID == 0 {
		return fmt.Errorf("forecast id must be provided. got: %d", forecastID)
	}

	// ensure the forecast exists
	f := db.QueryOperator.ForecastORM
	if _, err := f.WithContext(ctx).Where(f.Id.Eq(forecastID)).First(); err != nil {
		return fmt.Errorf("forecast with id %d does not exist", forecastID)
	}

	result, err := f.WithContext(ctx).Where(f.Id.Eq(forecastID)).Select(field.AssociationFields).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
