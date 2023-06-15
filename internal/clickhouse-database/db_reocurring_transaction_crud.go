package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// AddReOccuringTransaction adds a new reoccurring transaction to the database for a given user id
func (db *Db) AddReOccurringTransaction(ctx context.Context, userId *uint64, tx *schema.ReOccuringTransaction) (*uint64, error) {
	// trace the operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-reocurring-transaction"); span != nil {
		defer span.End()
	}

	// validate input parameters
	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	if tx == nil {
		return nil, fmt.Errorf("transaction is nil")
	}

	if tx.Id != 0 {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// the link id must be provided at write time in order to properly associate the transaction
	// to a given plaid link. This facilitates the ability to delete all transactions associated
	// to a given plaid link
	if tx.LinkId == 0 {
		return nil, fmt.Errorf("transaction link ID must be greater than 0")
	}

	// associate the user id to the transaction of interest
	tx.UserId = *userId
	// convert the transaction object to orm type
	record, err := tx.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// validate the transaction object
	if err := tx.Validate(); err != nil {
		return nil, err
	}

	// create the transaction record
	t := db.QueryOperator.ReOccuringTransactionORM
	if err := t.WithContext(ctx).Create(&record); err != nil {
		return nil, err
	}

	return &record.Id, nil
}

// AddReOccurringTransactions adds a new set of reoccurring transactions to the database and associates them to a given user id
func (db *Db) AddReOccurringTransactions(ctx context.Context, userId *uint64, txs []*schema.ReOccuringTransaction) error {
	// trace the operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-reocurring-transactions"); span != nil {
		defer span.End()
	}

	// validate input parameters
	if userId == nil {
		return fmt.Errorf("user ID is nil")
	}

	if txs == nil {
		return fmt.Errorf("transactions is nil")
	}

	txRecords := make([]*schema.ReOccuringTransactionORM, 0, len(txs))
	// associate the user id to the transaction of interest
	for _, tx := range txs {
		if tx.Id != 0 {
			return fmt.Errorf("transaction ID must be 0 at creation time")
		}

		if tx.LinkId == 0 {
			return fmt.Errorf("transaction link ID must be greater than 0")
		}

		// associate the user id to the transaction of interest
		tx.UserId = *userId

		// validate transactions
		if err := tx.Validate(); err != nil {
			return err
		}

		record, err := tx.ToORM(ctx)
		if err != nil {
			return err
		}

		txRecords = append(txRecords, &record)
	}

	t := db.QueryOperator.ReOccuringTransactionORM
	if err := t.WithContext(ctx).Create(txRecords...); err != nil {
		return err
	}

	return nil
}

// DeleteReOccuringTransaction delets a singular reoccurring transaction by its id
func (db *Db) DeleteReOccuringTransaction(ctx context.Context, txId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-reocurring-transaction"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.ReOccuringTransactionORM
	if _, err := db.GetTransactionById(ctx, txId); err != nil {
		return err
	}

	// delete the transaction
	result, err := t.WithContext(ctx).Where(t.Id.Eq(*txId)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// DeleteReOccurringTransactionsByIds deletes a set of reoccurring transactions by their ids
func (db *Db) DeleteReOccurringTransactionsByIds(ctx context.Context, txIds []uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-reocurring-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.ReOccuringTransactionORM
	// delete the transaction
	result, err := t.WithContext(ctx).Where(t.Id.In(txIds...)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// DeleteReOccurringTransactionsByLinkId deletes a set of reoccurring transactions by their link id
func (db *Db) DeleteReOcurringTransactionsByLinkId(ctx context.Context, linkId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-reocurring-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if linkId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.ReOccuringTransactionORM
	// delete all transactions matching this link id
	result, err := t.WithContext(ctx).Where(t.LinkId.Eq(*linkId)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// DeleteUserReOcurringTransactons deletes all reoccurring transactions for a given user id
func (db *Db) DeleteUserReOcurringTransactons(ctx context.Context, userId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-user-reocurring-transaction"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.ReOccuringTransactionORM
	// delete the transaction
	result, err := t.WithContext(ctx).Where(t.UserId.Eq(*userId)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetReOcurringTransactions gets a set of reoccurring transactions for a given user id in a paginated format
func (db *Db) GetReOcurringTransactions(ctx context.Context, userId *uint64, pagenumber int64, limit int64) ([]*schema.ReOccuringTransaction, int64, error) {
	var (
		nextPageNumber int64
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-reocurring-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID must be 0 at creation time")
	}

	// sanitize the request params
	pageNumber, pageSize := db.sanitizePaginationParams(pagenumber, limit)
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	t := db.QueryOperator.ReOccuringTransactionORM
	txs, err := t.
		WithContext(ctx).
		Where(t.UserId.Eq(*userId)).
		Limit(int(pageSize)).Offset(int(pageSize * (pageNumber - 1))).
		Find()
	if err != nil {
		return nil, 0, err
	}

	if len(txs) == 0 {
		return nil, 0, fmt.Errorf("no transactions found")
	}

	results := make([]*schema.ReOccuringTransaction, 0, len(txs))
	for _, tx := range txs {
		txRecord, err := tx.ToPB(ctx)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, &txRecord)
	}

	return results, nextPageNumber, nil
}

// UpdateReOccurringTransaction updates a singular reoccurring transaction
func (db *Db) UpdateReOccurringTransaction(ctx context.Context, userId *uint64, txId *uint64, tx *schema.ReOccuringTransaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-reocurring-update-transaction"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	if tx == nil {
		return fmt.Errorf("transaction must not be nil")
	}

	// validate transactions
	if err := tx.Validate(); err != nil {
		return err
	}

	//	get the transacton by tx id
	t := db.QueryOperator.ReOccuringTransactionORM
	if _, err := db.GetTransactionById(ctx, txId); err != nil {
		return err
	}

	// update the transaction
	txOrm, err := tx.ToORM(ctx)
	if err != nil {
		return err
	}

	// perform the update operation
	result, err := t.WithContext(ctx).Updates(txOrm)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// UpdateReOccurringTransactions updates a set of reoccurring transactions
func (db *Db) UpdateReOccurringTransactions(ctx context.Context, userId *uint64, txs []*schema.ReOccuringTransaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-transaction"); span != nil {
		defer span.End()
	}

	if len(txs) == 0 {
		return fmt.Errorf("transactions length must be greater than 0")
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	txnsOrmRecords := make([]*schema.ReOccuringTransactionORM, 0, len(txs))
	for _, tx := range txs {
		// associate the user id with the transaction
		tx.UserId = *userId

		txnOrm, err := tx.ToORM(ctx)
		if err != nil {
			return err
		}

		txnsOrmRecords = append(txnsOrmRecords, &txnOrm)
	}

	t := db.QueryOperator.ReOccuringTransactionORM
	// perform the update operation
	result, err := t.WithContext(ctx).Where(t.UserId.Eq(*userId)).Updates(txnsOrmRecords)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// GetReOcurringTransactionById gets a singular reoccurring transaction by its id
func (db *Db) GetReOcurringTransactionById(ctx context.Context, txId *uint64) (*schema.ReOccuringTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-reocurring-transaction-by-id"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	t := db.QueryOperator.ReOccuringTransactionORM
	record, err := t.WithContext(ctx).Where(t.Id.Eq(*txId)).First()
	if err != nil {
		return nil, fmt.Errorf("transaction with id %d does not exist", txId)
	}

	tx, err := record.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (db *Db) GetUserReOccurringTransactions(ctx context.Context, userId *uint64) ([]*schema.ReOccuringTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-reocurring-transaction-for-user"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("userId must not be nil")
	}

	t := db.QueryOperator.ReOccuringTransactionORM
	txs, err := t.
		WithContext(ctx).
		Where(t.UserId.Eq(*userId)).
		Find()
	if err != nil {
		return nil, err
	}

	if len(txs) == 0 {
		return nil, fmt.Errorf("no transactions found")
	}

	results := make([]*schema.ReOccuringTransaction, 0, len(txs))
	for _, tx := range txs {
		txRecord, err := tx.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		results = append(results, &txRecord)
	}

	return results, nil
}
