package database

import (
	"context"
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/newrelic/go-agent/v3/newrelic"
	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"gorm.io/gorm"
)

const (
	UPDATE_VIRTUAL_ACCOUNT_OPERATION    = "UPDATE_VIRTUAL_ACCOUNT_OPERATION"
	UPDATE_VIRTUAL_ACCOUNT_OPERATION_TX = "UPDATE_VIRTUAL_ACCOUNT_OPERATION_TX"
)

// UpdateVirtualAccount executes a tx to update a virtual account
func (db *Db) UpdateVirtualAccount(ctx context.Context, vAcct *schema.VirtualAccount, vAcctID uint64) error {
	if vAcctID == 0 {
		return errors.New("invalid input argument. virtual account id cannot be 0")
	}

	txn := db.Telemetry.StartTransaction(UPDATE_VIRTUAL_ACCOUNT_OPERATION)
	defer txn.End()

	dbTxn := db.updateVirtualAccountTxn(ctx, txn, vAcct, vAcctID)
	return db.Conn.PerformTransaction(ctx, dbTxn)
}

// updateVirtualAccountTxn encompasses the tx logic used to update a virtual account
func (db *Db) updateVirtualAccountTxn(ctx context.Context, txn *newrelic.Transaction, vAcct *schema.VirtualAccount, vAcctID uint64) core_database.Tx {
	return func(ctx context.Context, tx *gorm.DB) error {
		table := schema.VirtualAccountORM{}.TableName()
		operation := "update_virtual_account"
		s := db.startDatastoreSegment(txn, operation, table)
		defer s.End()

		if vAcct == nil {
			err := errors.New("invalid input argument. virtual account object cannot be nil")
			db.Logger.Error(err.Error())
			return err
		}

		if vAcctID == 0 {
			err := errors.New("invalid input argument. virtual account id cannot be 0")
			db.Logger.Error(err.Error())
			return err
		}

		vAcctOrm, err := vAcct.ToORM(ctx)
		if err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		if err := db.validateVirtualAcct(ctx, &vAcctOrm); err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		if _, err := db.FindVirtualAcctID(ctx, vAcctID); err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		if err := db.SaveVirtualAccountRecord(ctx, vAcct); err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		return nil
	}
}
