package clickhousedatabase

import (
	"context"
	"fmt"
	"time"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// AddTransaction adds a single transaction to the Clickhouse database.
func (db *Db) AddTransaction(ctx context.Context, userId *uint64, tx *schema.Transaction) (*uint64, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-transaction"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	if tx == nil {
		return nil, fmt.Errorf("transaction is nil")
	}

	if tx.Id != 0 {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// associate the user id to the transaction of interest
	tx.UserId = *userId

	record, err := tx.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// validate the transaction object
	if err := tx.Validate(); err != nil {
		return nil, err
	}

	t := db.QueryOperator.TransactionORM
	if err := t.WithContext(ctx).Create(&record); err != nil {
		return nil, err
	}

	return &record.Id, nil
}

// AddTransactions adds transactions to the Clickhouse database.
func (db *Db) AddTransactions(ctx context.Context, userId *uint64, txs []*schema.Transaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return fmt.Errorf("user ID is nil")
	}

	if txs == nil {
		return fmt.Errorf("transactions is nil")
	}

	txRecords := make([]*schema.TransactionORM, 0, len(txs))
	for _, tx := range txs {
		if tx.Id != 0 {
			return fmt.Errorf("transaction ID must be 0 at creation time")
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

	t := db.QueryOperator.TransactionORM
	if err := t.WithContext(ctx).Create(txRecords...); err != nil {
		return err
	}

	return nil
}

// DeleteTransaction deletes a single transaction.
func (db *Db) DeleteTransaction(ctx context.Context, txId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transaction"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.TransactionORM
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

// DeleteTransactionsByIds deletes transactions by ids.
func (db *Db) DeleteTransactionsByIds(ctx context.Context, txIds []uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.TransactionORM
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

// DeleteTransactionsByLinkId deletes all transactions for a given link id.
func (db *Db) DeleteTransactionsByLinkId(ctx context.Context, linkId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if linkId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.TransactionORM
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

// DeleteUserTransactons deletes all transactions for a given user.
func (db *Db) DeleteUserTransactons(ctx context.Context, userId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transaction"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.TransactionORM
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

// GetTransactions gets a set of transactions in a paginated fashion.
func (db *Db) GetTransactions(ctx context.Context, userId *uint64, pagenumber int64, limit int64) ([]*schema.Transaction, int64, error) {
	var (
		nextPageNumber int64
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-transactions"); span != nil {
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

	t := db.QueryOperator.TransactionORM
	txs, err := t.
		WithContext(ctx).
		Limit(int(pageSize)).Offset(int(pageSize * (pageNumber - 1))).
		Find()
	if err != nil {
		return nil, 0, err
	}

	results := make([]*schema.Transaction, 0, len(txs))
	for _, tx := range txs {
		txRecord, err := tx.ToPB(ctx)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, &txRecord)
	}

	return results, nextPageNumber, nil
}

// UpdateTransaction updates a single transaction.
func (db *Db) UpdateTransaction(ctx context.Context, userId *uint64, txId *uint64, tx *schema.Transaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-transaction"); span != nil {
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
	t := db.QueryOperator.TransactionORM
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

// UpdateTransactions updates transactions.
func (db *Db) UpdateTransactions(ctx context.Context, userId *uint64, txs []*schema.Transaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-transaction"); span != nil {
		defer span.End()
	}

	if len(txs) == 0 {
		return fmt.Errorf("transactions length must be greater than 0")
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	txnsOrmRecords := make([]*schema.TransactionORM, 0, len(txs))
	for _, tx := range txs {
		// associate the user id with the transaction
		tx.UserId = *userId

		txnOrm, err := tx.ToORM(ctx)
		if err != nil {
			return err
		}

		txnsOrmRecords = append(txnsOrmRecords, &txnOrm)
	}

	t := db.QueryOperator.TransactionORM
	// perform the update operation

	for _, tx := range txnsOrmRecords {
		result, err := t.WithContext(ctx).Where(t.UserId.Eq(*userId)).Updates(tx)
		if err != nil {
			return err
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}
	}

	return nil
}

// GetTransactionById Gets a transaction by ID.
func (db *Db) GetTransactionById(ctx context.Context, txId *uint64) (*schema.Transaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-transaction-by-id"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	t := db.QueryOperator.TransactionORM
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

// GetTransactionsByPlaidTransactionIds gets transactions by plaid transaction ids.
func (db *Db) GetTransactionsByPlaidTransactionIds(ctx context.Context, txIds []string) ([]*schema.Transaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.QueryOperator.TransactionORM
	// delete the transaction
	result, err := t.WithContext(ctx).Where(t.TransactionId.In(txIds...)).Find()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("no txn found")
	}

	resultSet := make([]*schema.Transaction, 0, len(result))
	for _, tx := range result {
		txRecord, err := tx.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		resultSet = append(resultSet, &txRecord)
	}

	return resultSet, nil
}

type MonthlyCategoryExpense struct {
	Month       time.Time
	Category    string
	TotalAmount float64
}

// `sanitizePaginationParams` is a method of the `Db` struct in the `postgres` package.
// It takes in a page number and a page size as parameters. It returns a sanitized
// page number and page size.
func (*Db) sanitizePaginationParams(pageNumber int64, pageSize int64) (int64, int64) {
	if pageNumber <= 0 {
		pageNumber = 1
	}

	if pageSize > 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = 10
	}
	return pageNumber, pageSize
}
