package database

import (
	"context"
	"errors"
	"fmt"

	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/newrelic/go-agent/v3/newrelic"
	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"gorm.io/gorm"
)

const (
	FIND_VIRTUAL_ACCOUNT_ID_OPERATION = "FIND_VIRTUAL_ACCOUNT_ID_OPERATION"
	FIND_VIRTUAL_ACCOUNT_ID_TX        = "FIND_VIRTUAL_ACCOUNT_ID_TX"
)

// FindVirtualAcctID performs a db tx to obtain a virtual account based on provided ID
func (db *Db) FindVirtualAcctID(ctx context.Context, vAcctID uint64) (*schema.VirtualAccount, error) {
	if vAcctID == 0 {
		return nil, errors.New("invalid input argument. virtual account cannot be 0")
	}

	txn := db.Telemetry.StartTransaction(FIND_VIRTUAL_ACCOUNT_ID_OPERATION)
	defer txn.End()

	dbTxn := db.findVirtualAcctByIDTxn(ctx, txn, vAcctID)
	result, err := db.Conn.PerformComplexTransaction(ctx, dbTxn)
	if err != nil {
		return nil, err
	}

	acc, ok := result.(*schema.VirtualAccount)
	if !ok {
		return nil, errors.New("failed to interface to type")
	}

	return acc, nil
}

// findVirtualAcctByIDTxn defines the business logic for getting a virtual account from the db
func (db *Db) findVirtualAcctByIDTxn(ctx context.Context, txn *newrelic.Transaction, vAcctID uint64) core_database.CmplxTx {
	return func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		table := schema.VirtualAccountORM{}.TableName()
		operation := "find_virtual_account_by_id"

		db.Logger.Info(fmt.Sprintf("started transaction - %s", operation))
		s := db.startDatastoreSegment(txn, operation, table)
		defer s.End()

		// perform validations
		if vAcctID == 0 {
			err := errors.New("invalid input argument. vAcctID cannot be 0")
			db.Logger.Error(err.Error())
			return nil, err
		}

		// find the account of interest from the database
		var account schema.VirtualAccountORM
		if err := tx.Where(&schema.VirtualAccountORM{Id: vAcctID}).First(&account).Error; err != nil {
			db.Logger.Error(err.Error())
			return nil, err
		}

		// ensure account is active
		if !account.Active {
			err := errors.New("virtual account does not exist. account was soft deleted")
			db.Logger.Error(err.Error())
			return nil, err
		}

		acct, err := account.ToPB(ctx)
		if err != nil {
			db.Logger.Error(err.Error())
			return nil, err
		}

		db.Logger.Info(fmt.Sprintf("successfully completed transaction - %s", operation))

		return &acct, nil
	}
}
