package database

import (
	"context"
	"errors"
	"fmt"

	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
)

func (db *Db) validateVirtualAcct(ctx context.Context, vAcctOrm *schema.VirtualAccountORM) error {
	if vAcctOrm == nil {
		err := errors.New("virtual account object cannot be nil")
		db.Logger.Error(err.Error())
		return err
	}

	if vAcctOrm.UserID == 0 {
		err := errors.New(fmt.Sprintf("userID cannot be %d", vAcctOrm.Id))
		db.Logger.Error(err.Error())
		return err
	}

	if vAcctOrm.AccessToken == "" {
		err := errors.New("accesstoken cannot be empty")
		db.Logger.Error(err.Error())
		return err
	}

	return nil
}
