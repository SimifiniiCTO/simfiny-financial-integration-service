package database

import (
	"context"
	"fmt"
	"time"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func (db *Db) GetLastPlaidSync(ctx context.Context, userId, linkId uint64) (*schema.PlaidSync, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-last-plaid-sync"); span != nil {
		defer span.End()
	}

	link, err := db.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return nil, err
	}

	return link.PlaidSync, nil
}

func (db *Db) RecordPlaidSync(ctx context.Context, userId, plaidLinkId uint64, trigger, nextCursor string, added, modified, removed int64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-record-plaid-sync"); span != nil {
		defer span.End()
	}

	// ensure the link exists
	link, err := db.GetLink(ctx, userId, plaidLinkId, false)
	if err != nil {
		return err
	}

	if link.PlaidLink == nil {
		return fmt.Errorf("link with id %d does not have a plaid link", plaidLinkId)
	}

	// convert the link to orm
	linkOrm, err := link.ToORM(ctx)
	if err != nil {
		return fmt.Errorf("failed to convert link to orm: %v", err)
	}

	plaidSync := &schema.PlaidSyncORM{
		Added:      added,
		Id:         userId,
		Modified:   modified,
		NextCursor: nextCursor,
		Removed:    removed,
		TimeStamp:  time.Now().String(),
		Trigger:    trigger,
	}

	// update the link
	l := db.QueryOperator.LinkORM
	if err := l.PlaidSync.Model(&linkOrm).Replace(plaidSync); err != nil {
		return fmt.Errorf("failed to append plaid sync to link: %v", err)
	}

	result, err := l.WithContext(ctx).Updates(&linkOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
