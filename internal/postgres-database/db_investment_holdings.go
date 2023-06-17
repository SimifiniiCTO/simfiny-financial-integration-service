package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// AddInvestmentHoldings adds a set of investment holdings into a target investment account
func (db *Db) AddInvestmentHoldings(ctx context.Context, userId *uint64, investmentAccountId *uint64, investmentHoldings []*schema.InvesmentHolding) (*schema.InvestmentAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-investment-holding"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	if investmentHoldings == nil {
		return nil, fmt.Errorf("investment holding is nil")
	}

	// validate all present holdings
	for _, holding := range investmentHoldings {
		if holding.Id != 0 {
			return nil, fmt.Errorf("investment holding ID must be 0 at creation time")
		}

		// validate the investment holding object
		if err := holding.Validate(); err != nil {
			return nil, err
		}
	}

	// ensure that the investment account exists
	ia := db.QueryOperator.InvestmentAccountORM
	investmentAcct, err := ia.WithContext(ctx).Where(ia.Id.Eq(*investmentAccountId)).First()
	if err != nil {
		return nil, fmt.Errorf("investment account with id %d does not exist", *investmentAccountId)
	}

	// convert all the investment holdings to orm
	records := make([]*schema.InvesmentHoldingORM, 0, len(investmentHoldings))
	for _, holding := range investmentHoldings {
		record, err := holding.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)
	}

	// perform the bulk insert operation
	if err := ia.Holdings.Model(investmentAcct).Append(records...); err != nil {
		return nil, err
	}

	result, err := investmentAcct.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *Db) UpdateInvestmentHoldings(ctx context.Context, userId *uint64, investmentAccountId *uint64, investmentHoldings []*schema.InvesmentHolding) (*schema.InvestmentAccount, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-investment-holding"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	if investmentHoldings == nil {
		return nil, fmt.Errorf("investment holding is nil")
	}

	// validate all present holdings
	for _, holding := range investmentHoldings {
		// validate the investment holding object
		if err := holding.Validate(); err != nil {
			return nil, err
		}
	}

	// ensure that the investment account exists
	ia := db.QueryOperator.InvestmentAccountORM
	investmentAcct, err := ia.WithContext(ctx).Where(ia.Id.Eq(*investmentAccountId)).First()
	if err != nil {
		return nil, fmt.Errorf("investment account with id %d does not exist", *investmentAccountId)
	}

	// convert all the investment holdings to orm
	records := make([]*schema.InvesmentHoldingORM, 0, len(investmentHoldings))
	for _, holding := range investmentHoldings {
		record, err := holding.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)
	}

	// perform the bulk update operation
	if err := ia.Holdings.Model(investmentAcct).Replace(records...); err != nil {
		return nil, err
	}

	result, err := investmentAcct.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *Db) GetInvestmentHoldings(ctx context.Context, investmentHoldingIds ...uint64) ([]*schema.InvesmentHolding, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-investment-holding"); span != nil {
		defer span.End()
	}

	if investmentHoldingIds == nil {
		return nil, fmt.Errorf("investment holding id must be provided. got: %d", investmentHoldingIds)
	}

	// ensure the investment holdings exist
	ih := db.QueryOperator.InvesmentHoldingORM
	records, err := ih.WithContext(ctx).Where(ih.Id.In(investmentHoldingIds...)).Find()
	if err != nil {
		return nil, err
	}

	// convert the orm to pb
	holdings := make([]*schema.InvesmentHolding, 0, len(records))
	for _, record := range records {
		holding, err := record.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		holdings = append(holdings, &holding)
	}

	return holdings, nil
}

func (db *Db) DeleteInvestmentHoldings(ctx context.Context, investmentHoldingIds ...uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-investment-holding"); span != nil {
		defer span.End()
	}

	if investmentHoldingIds == nil {
		return fmt.Errorf("investment holding id must be provided. got: %d", investmentHoldingIds)
	}

	// delete the investment holdings by id
	ih := db.QueryOperator.InvesmentHoldingORM
	result, err := ih.WithContext(ctx).Where(ih.Id.In(investmentHoldingIds...)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
