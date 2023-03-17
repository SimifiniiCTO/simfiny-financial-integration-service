package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreatePocket implements DatabaseOperations
func (db *Db) CreatePocket(ctx context.Context, bankAccountID uint64, pocket *schema.Pocket) (*schema.Pocket, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-pocket"); span != nil {
		defer span.End()
	}

	if pocket == nil {
		return nil, fmt.Errorf("pocket must be provided")
	}

	if bankAccountID == 0 {
		return nil, fmt.Errorf("bank account id must be provided. got: %d", bankAccountID)
	}

	// ensure the bank account exists
	b := db.QueryOperator.BankAccountORM
	bankAcct, err := b.WithContext(ctx).Where(b.Id.Eq(bankAccountID)).First()
	if err != nil {
		return nil, fmt.Errorf("bank account with id %d does not exist", bankAccountID)
	}

	pocketOrm, err := pocket.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// associate the bank account with the pocket
	if err := b.Pockets.Model(bankAcct).Append(&pocketOrm); err != nil {
		return nil, err
	}

	// update the bank account with the new pocket
	result, err := b.WithContext(ctx).Updates(bankAcct)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// convert orm to schema
	// Note the update operation populate the record id in the pocketOrm
	record, err := pocketOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// DeletePocket implements DatabaseOperations
func (db *Db) DeletePocket(ctx context.Context, pocketID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-pocket"); span != nil {
		defer span.End()
	}

	if pocketID == 0 {
		return fmt.Errorf("pocket id must be provided. got: %d", pocketID)
	}

	// ensure the pocket exists
	p := db.QueryOperator.PocketORM
	if _, err := p.WithContext(ctx).Where(p.Id.Eq(pocketID)).First(); err != nil {
		return fmt.Errorf("pocket with id %d does not exist", pocketID)
	}

	// perform the delete operation
	// TODO: ensure all associations are deleted
	result, err := p.WithContext(ctx).Where(p.Id.Eq(pocketID)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetPocket implements DatabaseOperations
func (db *Db) GetPocket(ctx context.Context, pocketID uint64) (*schema.Pocket, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-pocket"); span != nil {
		defer span.End()
	}

	if pocketID == 0 {
		return nil, fmt.Errorf("pocket id must be provided. got: %d", pocketID)
	}

	// ensure the pocket exists
	p := db.QueryOperator.PocketORM
	pocketOrm, err := p.
		WithContext(ctx).
		Where(p.Id.Eq(pocketID)).
		Preload(p.Goals.Forecasts).
		Preload(p.Goals.Milestones.Budget.Category).
		First()
	if err != nil {
		return nil, fmt.Errorf("pocket with id %d does not exist", pocketID)
	}

	// convert orm to schema
	record, err := pocketOrm.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// UpdatePocket implements DatabaseOperations
func (db *Db) UpdatePocket(ctx context.Context, pocket *schema.Pocket) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-pocket"); span != nil {
		defer span.End()
	}

	if pocket == nil {
		return fmt.Errorf("pocket must be provided")
	}

	if pocket.Id == 0 {
		return fmt.Errorf("pocket id must be provided. got: %d", pocket.Id)
	}

	// ensure the pocket exists
	p := db.QueryOperator.PocketORM
	if _, err := p.WithContext(ctx).Where(p.Id.Eq(pocket.Id)).First(); err != nil {
		return fmt.Errorf("pocket with id %d does not exist", pocket.Id)
	}

	// update the pocket
	pocketOrm, err := pocket.ToORM(ctx)
	if err != nil {
		return err
	}

	// perform the update operation
	result, err := p.WithContext(ctx).Updates(pocketOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
