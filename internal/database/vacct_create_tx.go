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
	CREATE_VIRTUAL_ACCOUNT_OPERATION    = "CREATE_VIRTUAL_ACCOUNT_OPERATION"
	CREATE_VIRTUAL_ACCOUNT_OPERATION_TX = "CREATE_VIRTUAL_ACCOUNT_OPERATION_TX"
)

// CreateVirtualAccount executes a db tx to create a virtual account and store locally in db
func (db *Db) CreateVirtualAccount(ctx context.Context, vAcct *schema.VirtualAccount, accessToken string) (*schema.VirtualAccount, error) {
	if vAcct == nil {
		return nil, errors.New("invalid input argument. vAcct cannot be nil")
	}

	if accessToken == EMPTY {
		return nil, errors.New("invalid input argument. accessToken cannot be empty")
	}

	txn := db.Telemetry.StartTransaction(CREATE_VIRTUAL_ACCOUNT_OPERATION)
	defer txn.End()

	dbTxn := db.createVirtualAccountTxn(ctx, txn, vAcct, accessToken)
	res, err := db.Conn.PerformComplexTransaction(ctx, dbTxn)
	if err != nil {
		return nil, err
	}

	acc, ok := res.(schema.VirtualAccount)
	if !ok {
		return nil, errors.New("failed to convert interface to type")
	}

	return &acc, nil
}

// createVirtualAccountTxn with holds actual creat account tx logic
func (db *Db) createVirtualAccountTxn(ctx context.Context, txn *newrelic.Transaction, vAcct *schema.VirtualAccount,
	accessToken string) core_database.CmplxTx {
	return func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		table := schema.VirtualAccountORM{}.TableName()
		operation := "create_virtual_account"

		db.Logger.Info(fmt.Sprintf("starting transaction - %s", operation))
		s := db.startDatastoreSegment(txn, operation, table)
		defer s.End()

		if vAcct == nil {
			err := errors.New("invalid virtual account parameter. param is nil")
			db.Logger.Error(err.Error())
			return nil, err
		}

		if accessToken == EMPTY {
			return nil, errors.New("invalid input argument. accessToken cannot be empty")
		}

		vAcctOrm, err := vAcct.ToORM(ctx)
		if err != nil {
			db.Logger.Error(err.Error())
			return nil, err
		}

		// validate account
		if err = db.validateVirtualAcct(ctx, &vAcctOrm); err != nil {
			db.Logger.Error(err.Error())
			return nil, err
		}

		// check if record already exists based on user ID
		if _, err = db.FindVirtualAcctByUserID(ctx, vAcctOrm.UserID); err == nil {
			return nil, errors.New("account already exists")
		}

		// set acct as active & populate access token
		vAcctOrm.Active = true
		vAcctOrm.AccessToken = accessToken

		// save account
		if err = db.Conn.Engine.Create(&vAcctOrm).Error; err != nil {
			db.Logger.Error(err.Error())
			return nil, err
		}

		db.Logger.Info(fmt.Sprintf("successfully completed transaction - %s", operation))

		return vAcctOrm.ToPB(ctx)
	}
}
