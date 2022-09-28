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
	SAVE_VIRTUAL_ACCOUNT_OPERATION    = "SAVE_VIRTUAL_ACCOUNT_OPERATION"
	SAVE_VIRTUAL_ACCOUNT_OPERATION_TX = "SAVE_VIRTUAL_ACCOUNT_OPERATION_TX"
)

// SaveVirtualAccountRecord saves a record in the database
func (db *Db) SaveVirtualAccountRecord(ctx context.Context, vAcct *schema.VirtualAccount) error {
	if vAcct == nil {
		return errors.New("invalid virtual account parameter. param is nil")
	}

	txn := db.Telemetry.StartTransaction(SAVE_VIRTUAL_ACCOUNT_OPERATION)
	defer txn.End()

	tx := db.saveVirtualAccountTxFunc(ctx, txn, vAcct)
	if err := db.Conn.PerformTransaction(ctx, tx); err != nil {
		return err
	}

	return nil
}

// saveVirtualAccountTxFunc wraps the logic in a db tx and returns it
func (db *Db) saveVirtualAccountTxFunc(ctx context.Context, txn *newrelic.Transaction, vAcct *schema.VirtualAccount) core_database.Tx {
	return func(ctx context.Context, tx *gorm.DB) error {
		table := schema.VirtualAccountORM{}.TableName()
		operation := "save_virtual_account"

		s := db.startDatastoreSegment(txn, operation, table)
		defer s.End()

		db.Logger.Info(fmt.Sprintf("starting transaction - %s", operation))
		if vAcct == nil {
			err := errors.New("invalid virtual account parameter. param is nil")
			db.Logger.Error(err.Error())
			return err
		}

		vAcctOrm, err := vAcct.ToORM(ctx)
		if err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		if err = db.validateVirtualAcct(ctx, &vAcctOrm); err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&vAcctOrm).Error; err != nil {
			db.Logger.Error(err.Error())
			return err
		}

		db.Logger.Info(fmt.Sprintf("successfully completed transaction - %s", operation))

		return nil
	}
}
