package clickhousedatabase

import (
	"context"
	"fmt"
	"strings"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// AddTransaction adds a single transaction to the Clickhouse database.
func (db *Db) AddTransaction(ctx context.Context, userId *uint64, tx *schema.Transaction) (*string, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-add-transaction"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	if tx == nil {
		return nil, fmt.Errorf("transaction is nil")
	}

	if tx.Id != "" {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// associate the user id to the transaction of interest
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	tx.UserId = *userId
	tx.Id = id.String()

	// validate the transaction object
	if err := tx.Validate(); err != nil {
		return nil, err
	}

	record, err := tx.ConvertToInternal()
	if err != nil {
		return nil, err
	}

	_, err = db.queryEngine.NewInsert().Model(record).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// get the newly created tx
	createdTx := new(schema.TransactionInternal)
	if err := db.queryEngine.NewSelect().Model(createdTx).Where("UserId = ?", record.UserId).Scan(ctx); err != nil {
		return nil, err
	}

	return &record.ID, nil
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

	txRecords := make([]schema.TransactionInternal, 0, len(txs))
	for _, tx := range txs {
		if tx.Id != "" {
			return fmt.Errorf("transaction ID must be 0 at creation time")
		}

		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		// associate the user id to the transaction of interest
		tx.UserId = *userId
		tx.Id = id.String()

		// validate transactions
		if err := tx.Validate(); err != nil {
			return err
		}

		record, err := tx.ConvertToInternal()
		if err != nil {
			return err
		}

		txRecords = append(txRecords, *record)
	}

	if _, err := db.queryEngine.NewInsert().Model(&txRecords).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteTransaction deletes a single transaction.
func (db *Db) DeleteTransaction(ctx context.Context, txId *string) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transaction"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// raw query to delete the transaction
	query := fmt.Sprintf(`ALTER TABLE TransactionInternal DELETE WHERE ID = '%s'`, *txId)
	if err := db.queryEngine.
		NewRaw(query).
		Scan(ctx); err != nil {
		return err
	}

	return nil
}

func (db *Db) DeleteTransactionsByIds(ctx context.Context, txIds []string) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// Create a string with comma-separated values from txIds
	ids := "'" + strings.Join(txIds, "', '") + "'"

	sqlQuery := fmt.Sprintf("ALTER TABLE TransactionInternal DELETE WHERE ID IN (%s)", ids)

	if err := db.queryEngine.
		NewRaw(sqlQuery).
		Scan(ctx); err != nil {
		return err
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

	if err := db.queryEngine.
		NewRaw(
			fmt.
				Sprintf("ALTER TABLE TransactionInternal DELETE WHERE LinkId = %d", *linkId)).
		Scan(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteUserTransactons deletes all transactions for a given user.
func (db *Db) DeleteUserTransactions(ctx context.Context, userId *uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transaction"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	// raw query to delete the transaction
	if err := db.queryEngine.
		NewRaw(
			fmt.
				Sprintf("ALTER TABLE TransactionInternal DELETE WHERE UserId = %d", *userId)).
		Scan(ctx); err != nil {
		return err
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

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	var transactions []schema.TransactionInternal
	if err := db.
		queryEngine.
		NewSelect().
		Model(&transactions).
		Where("UserId = ?", *userId).
		Where("Sign = 1").
		Offset(offset).
		Limit(queryLimit).
		Order("Time DESC").
		Scan(ctx); err != nil {
		return nil, 0, err
	}

	results := make([]*schema.Transaction, 0, len(transactions))
	if len(transactions) > 0 {
		for _, tx := range transactions {
			txRecord, err := tx.ConvertToTransaction()
			if err != nil {
				return nil, 0, err
			}
			results = append(results, txRecord)
		}
	}

	return results, nextPageNumber, nil
}

func (db *Db) GetTransactionsForAccount(ctx context.Context, userId *uint64, pagenumber int64, limit int64, accountId string) ([]*schema.Transaction, int64, error) {
	var (
		nextPageNumber int64
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID must be 0 at creation time")
	}

	if accountId == "" {
		return nil, 0, fmt.Errorf("account ID was not provided")
	}

	// sanitize the request params
	pageNumber, pageSize := db.sanitizePaginationParams(pagenumber, limit)
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	var transactions []schema.TransactionInternal
	if err := db.
		queryEngine.
		NewSelect().
		Model(&transactions).
		Where("UserId = ?", *userId).
		Where("AccountId = ?", accountId).
		Where("Sign = 1").
		Offset(offset).
		Limit(queryLimit).
		Order("Time DESC").
		Scan(ctx); err != nil {
		return nil, 0, err
	}

	results := make([]*schema.Transaction, 0, len(transactions))
	if len(transactions) > 0 {
		for _, tx := range transactions {
			txRecord, err := tx.ConvertToTransaction()
			if err != nil {
				return nil, 0, err
			}
			results = append(results, txRecord)
		}
	}

	return results, nextPageNumber, nil
}

func Stringify(strs []string) string {
	quotedStrs := make([]string, len(strs))
	for i, str := range strs {
		quotedStrs[i] = fmt.Sprintf("'%s'", str)
	}
	return fmt.Sprintf("[%s]", strings.Join(quotedStrs, ", "))
}

// UpdateTransaction updates a single transaction.
func (db *Db) UpdateTransaction(ctx context.Context, userId *uint64, txId *string, tx *schema.Transaction) error {
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
	storedTx, err := db.GetTransactionById(ctx, txId)
	if err != nil {
		return err
	}

	// since we use collapsing merge tree, we need to delete the old row and insert a new one
	// we first delete the old row
	storedTx.Sign = -1
	if _, err := db.AddTransaction(ctx, userId, storedTx); err != nil {
		return err
	}

	// now we add the updated transaction
	tx.Sign = 1
	tx.Id = *txId
	if _, err := db.AddTransaction(ctx, userId, tx); err != nil {
		return err
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

	// GetTransactionsByPlaidTransactionIds
	transactionIds := make([]string, 0, len(txs))
	transactionIdToOldTransaction := make(map[string]*schema.Transaction, 0)

	// we first iterate over the transactions and get the transaction ids
	for _, tx := range txs {
		// validate transactions
		if err := tx.Validate(); err != nil {
			return err
		}

		transactionIds = append(transactionIds, tx.TransactionId)
	}

	// get all the old transactions based on the transaction ids
	oldTransactions, err := db.GetTransactionsByPlaidTransactionIds(ctx, transactionIds)
	if err != nil {
		return err
	}

	transactionSetForUpdate := make([]*schema.Transaction, 0, len(oldTransactions))
	// now we iterate over the old transactions and ensure they have a negative sign
	for _, tx := range oldTransactions {
		currentTx := tx
		currentTx.Sign = -1
		transactionIdToOldTransaction[currentTx.TransactionId] = currentTx

		db.Logger.Info("old transactions", zap.Any("transactions", currentTx))
		transactionSetForUpdate = append(transactionSetForUpdate, currentTx)
	}

	transactionsSetForWrite := make([]*schema.Transaction, 0, len(txs))
	// iterate over the new transactions and ensure they have a positive sign
	for _, tx := range txs {
		currentTx := tx
		currentTx.Sign = 1

		// We should have a strategy to determine what a "duplicate" is.
		// One common way is to use a unique identifier like Id
		if oldTx, ok := transactionIdToOldTransaction[currentTx.TransactionId]; ok {
			currentTx.Id = oldTx.Id
		}

		transactionsSetForWrite = append(transactionsSetForWrite, currentTx)
	}

	// save this to the database
	if err := db.AddTransactions(ctx, userId, transactionSetForUpdate); err != nil {
		return err
	}

	// now updated transactions are added
	if err := db.AddTransactions(ctx, userId, transactionsSetForWrite); err != nil {
		return err
	}

	return nil
}

// GetTransactionById Gets a transaction by ID.
func (db *Db) GetTransactionById(ctx context.Context, txId *string) (*schema.Transaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-transaction-by-id"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	tx := new(schema.TransactionInternal)
	if err := db.queryEngine.NewSelect().Model(tx).Where("ID = ?", *txId).Where("Sign = 1").Scan(ctx); err != nil {
		return nil, err
	}

	res, err := tx.ConvertToTransaction()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *Db) GetTransactionByUserId(ctx context.Context, userId *uint64) (*schema.Transaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-transaction-by-id"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID cannot be nil")
	}

	tx := new(schema.TransactionInternal)
	if err := db.queryEngine.NewSelect().Model(tx).Where("UserId = ?", *userId).Where("Sign = 1").Scan(ctx); err != nil {
		return nil, err
	}

	res, err := tx.ConvertToTransaction()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTransactionsByPlaidTransactionIds gets transactions by plaid transaction ids.
func (db *Db) GetTransactionsByPlaidTransactionIds(ctx context.Context, txIds []string) ([]*schema.Transaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	var result []schema.TransactionInternal
	ids := "'" + strings.Join(txIds, "', '") + "'"
	sqlQuery := fmt.Sprintf("SELECT * FROM TransactionInternal WHERE TransactionId IN (%s)", ids)
	if err := db.queryEngine.
		NewRaw(sqlQuery).
		Scan(ctx, &result); err != nil {
		return nil, err
	}

	resultSet := make([]*schema.Transaction, 0, len(result))
	if len(result) > 0 {
		for _, tx := range result {
			txRecord, err := tx.ConvertToTransaction()
			if err != nil {
				return nil, err
			}

			resultSet = append(resultSet, txRecord)
		}
	}

	return resultSet, nil
}

// `sanitizePaginationParams` is a method of the `Db` struct in the `postgres` package.
// It takes in a page number and a page size as parameters. It returns a sanitized
// page number and page size.
func (*Db) sanitizePaginationParams(pageNumber int64, pageSize int64) (int64, int64) {
	if pageNumber <= 0 {
		pageNumber = 1
	}

	if pageSize > 300 {
		pageSize = 200
	} else if pageSize <= 0 {
		pageSize = 100
	}
	return pageNumber, pageSize
}
