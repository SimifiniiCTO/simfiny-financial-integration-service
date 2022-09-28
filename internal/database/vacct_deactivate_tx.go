package database

import (
	"context"
	"errors"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/newrelic/go-agent/v3/newrelic"
	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"gorm.io/gorm"
)

const (
	DEACTIVATE_VIRTUAL_ACCOUNT_OPERATION    = "DEACTIVATE_VIRTUAL_ACCOUNT_OPERATION"
	DEACTIVATE_VIRTUAL_ACCOUNT_OPERATION_TX = "DEACTIVATE_VIRTUAL_ACCOUNT_OPERATION_TX"
)

// DeactivateVirtualAccount executes a db tx that deactivates a virtual account
func (db *Db) DeactivateVirtualAccount(ctx context.Context, vAcctID uint64) error {
	if vAcctID == 0 {
		return errors.New("invalid input argument. vAcctID cannot be 0")
	}

	txn := db.Telemetry.StartTransaction(DEACTIVATE_VIRTUAL_ACCOUNT_OPERATION)
	defer txn.End()

	dbTxn := db.deactivateVirtualAccountTxn(ctx, txn, vAcctID)
	if err := db.Conn.PerformTransaction(ctx, dbTxn); err != nil {
		return err
	}

	return nil
}

// deactivateVirtualAccountTxn encompasses the db tx that deactivates a virtual account
func (db *Db) deactivateVirtualAccountTxn(ctx context.Context, txn *newrelic.Transaction, vAcctID uint64) core_database.Tx {
	return func(ctx context.Context, tx *gorm.DB) error {
		table := schema.VirtualAccountORM{}.TableName()
		operation := "deactivate_virtual_account"

		db.Logger.Info(fmt.Sprintf("starting transaction - %s", operation))
		s := db.startDatastoreSegment(txn, operation, table)
		defer s.End()

		if vAcctID == 0 {
			err := errors.New(fmt.Sprintf("invalid input parameter. virtualAcctID: %d", vAcctID))
			db.Logger.Error(err.Error())
			return err
		}

		vAcct, err := db.FindVirtualAcctID(ctx, vAcctID)
		if err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		if err := db.Conn.Engine.Model(&schema.VirtualAccountORM{}).Where("id", vAcct.Id).Update("active", "false").Error; err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		db.Logger.Info(fmt.Sprintf("successfully completed transaction - %s", operation))

		return nil
	}
}
