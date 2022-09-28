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
	FIND_VIRTUAL_ACCOUNT_USER_ID_OPERATION = "FIND_VIRTUAL_ACCOUNT_USER_ID_OPERATION"
	FIND_VIRTUAL_ACCOUNT_USER_ID_TX        = "FIND_VIRTUAL_ACCOUNT_USER_ID_TX"
)

// FindVirtualAcctByUserID performs a db tx to find a virtual account based on provided user id
func (db *Db) FindVirtualAcctByUserID(ctx context.Context, userID uint64) (*schema.VirtualAccount, error) {
	if userID == 0 {
		return nil, errors.New("invalid input argument. userID account cannot be 0")
	}

	txn := db.Telemetry.StartTransaction(FIND_VIRTUAL_ACCOUNT_USER_ID_OPERATION)
	defer txn.End()

	dbTxn := db.findVirtualAcctByUserIDTxn(ctx, txn, userID)
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

// findVirtualAcctByUserIDTxn encompasses the db tx used to find a virtual account based on provided user id
func (db *Db) findVirtualAcctByUserIDTxn(ctx context.Context, txn *newrelic.Transaction, userID uint64) core_database.CmplxTx {
	tx := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		table := schema.VirtualAccountORM{}.TableName()
		operation := "find_virtual_account_by_user_id"

		db.Logger.Info(fmt.Sprintf("started transaction - %s", operation))
		s := db.startDatastoreSegment(txn, operation, table)
		defer s.End()

		if userID == 0 {
			err := fmt.Errorf("invalid userID. userID cannot be 0")
			db.Logger.Error(err.Error())
			return nil, err
		}

		var account schema.VirtualAccountORM
		if err := tx.Where(&schema.VirtualAccountORM{UserID: userID}).First(&account).Error; err != nil {
			db.Logger.Error(err.Error())
			return nil, err
		}

		if !account.Active {
			errMsg := "virtual account does not exist. account was soft deleted"
			err := errors.New(errMsg)
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

	return tx
}
