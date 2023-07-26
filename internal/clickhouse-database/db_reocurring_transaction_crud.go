package clickhousedatabase

import (
	"context"
	"fmt"
	"strings"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/google/uuid"
)

// AddReOccuringTransaction adds a new reoccurring transaction to the database for a given user id
func (db *Db) AddReOccurringTransaction(ctx context.Context, userId *uint64, tx *schema.ReOccuringTransaction) (*string, error) {
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

	if tx.Id != "" {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// the link id must be provided at write time in order to properly associate the transaction
	// to a given plaid link. This facilitates the ability to delete all transactions associated
	// to a given plaid link
	if tx.LinkId == 0 {
		return nil, fmt.Errorf("transaction link ID must be greater than 0")
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	// associate the user id to the transaction of interest
	tx.UserId = *userId
	tx.Id = id.String()

	// validate the transaction object
	if err := tx.Validate(); err != nil {
		return nil, err
	}

	// convert the transaction object to orm type
	record, err := tx.ConvertToInternal()
	if err != nil {
		return nil, err
	}

	// create the transaction record
	_, err = db.queryEngine.NewInsert().Model(record).Exec(ctx)
	if err != nil {
		return nil, err
	}

	// get the newly created tx
	createdTx := new(schema.ReOccuringTransactionInternal)
	if err := db.queryEngine.NewSelect().Model(createdTx).Where("UserId = ?", record.UserId).Scan(ctx); err != nil {
		return nil, err
	}

	return &createdTx.ID, nil
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

	txRecords := make([]schema.ReOccuringTransactionInternal, 0, len(txs))
	// associate the user id to the transaction of interest
	for _, tx := range txs {
		if tx.Id != "" {
			db.Logger.Error("transaction ID must be 0 at creation time")
			return fmt.Errorf("transaction ID must be 0 at creation time")
		}

		if tx.LinkId == 0 {
			db.Logger.Error("transaction link ID must be greater than 0")
			return fmt.Errorf("transaction link ID must be greater than 0")
		}

		// associate the user id to the transaction of interest
		tx.UserId = *userId

		// validate transactions
		if err := tx.Validate(); err != nil {
			db.Logger.Error(err.Error())
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

// DeleteReOccuringTransaction delets a singular reoccurring transaction by its id
func (db *Db) DeleteReOccuringTransaction(ctx context.Context, txId *string) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-reocurring-transaction"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// raw query to delete the transaction
	query := fmt.Sprintf(`ALTER TABLE ReOccuringTransactionInternal DELETE WHERE ID = '%s'`, *txId)
	if err := db.queryEngine.
		NewRaw(query).
		Scan(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteReOccurringTransactionsByIds deletes a set of reoccurring transactions by their ids
func (db *Db) DeleteReOccurringTransactionsByIds(ctx context.Context, txIds []string) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-reocurring-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// Create a string with comma-separated values from txIds
	ids := "'" + strings.Join(txIds, "', '") + "'"

	sqlQuery := fmt.Sprintf("ALTER TABLE ReOccuringTransactionInternal DELETE WHERE ID IN (%s)", ids)

	if err := db.queryEngine.
		NewRaw(sqlQuery).
		Scan(ctx); err != nil {
		return err
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

	// raw query to delete the transaction
	query := fmt.Sprintf(`ALTER TABLE ReOccuringTransactionInternal DELETE WHERE LinkId = '%d'`, *linkId)
	if err := db.queryEngine.
		NewRaw(query).
		Scan(ctx); err != nil {
		return err
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

	// delete the transaction
	query := fmt.Sprintf(`ALTER TABLE ReOccuringTransactionInternal DELETE WHERE UserId = '%d'`, *userId)
	if err := db.queryEngine.
		NewRaw(query).
		Scan(ctx); err != nil {
		return err
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

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	query := fmt.Sprintf(`UserId = %d`, *userId)
	var transactions []schema.ReOccuringTransactionInternal
	if err := db.
		queryEngine.
		NewSelect().
		Model(&transactions).
		Where(query).
		Offset(offset).
		Limit(queryLimit).
		Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(transactions) == 0 {
		return nil, 0, fmt.Errorf("no transactions found")
	}

	results := make([]*schema.ReOccuringTransaction, 0, len(transactions))
	for _, tx := range transactions {
		txRecord, err := tx.ConvertToRecurringTransaction()
		if err != nil {
			return nil, 0, err
		}
		results = append(results, txRecord)
	}

	return results, nextPageNumber, nil
}

// UpdateReOccurringTransaction updates a singular reoccurring transaction
func (db *Db) UpdateReOccurringTransaction(ctx context.Context, userId *uint64, txId *string, tx *schema.ReOccuringTransaction) error {
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

	// update the transaction
	record, err := tx.ConvertToInternal()
	if err != nil {
		return err
	}

	query := `
		ALTER TABLE ReOccuringTransactionInternal UPDATE
		AccountId = ?,
		AverageAmount = ?,
		AverageAmountIsoCurrencyCode = ?,
		CategoryId = ?,
		Description = ?,
		FirstDate = ?,
		Flow = ?,
		Frequency = ?,
		IsActive = ?,
		LastAmount = ?,
		LastAmountIsoCurrencyCode = ?,
		LastDate = ?,
		LinkId = ?,
		MerchantName = ?,
		PersonalFinanceCategoryDetailed = ?,
		PersonalFinanceCategoryPrimary = ?,
		Status = ?,
		StreamId = ?,
		TransactionIds = ?,
		UpdatedTime = ?,
		UserId = ?
		WHERE ID = ?;
	`

	isActiveRef := func(b bool) int {
		if b {
			return 1
		} else {
			return 0
		}
	}(record.IsActive)

	if err := db.queryEngine.NewRaw(query,
		record.AccountId,
		record.AverageAmount,
		record.AverageAmountIsoCurrencyCode,
		record.CategoryId,
		record.Description,
		record.FirstDate,
		record.Flow,
		record.Frequency,
		isActiveRef,
		record.LastAmount,
		record.LastAmountIsoCurrencyCode,
		record.LastDate,
		record.LinkId,
		record.MerchantName,
		record.PersonalFinanceCategoryDetailed,
		record.PersonalFinanceCategoryPrimary,
		record.Status,
		record.StreamId,
		record.TransactionIds,
		record.UpdatedTime,
		record.UserId,
		record.ID).Scan(ctx); err != nil {
		return err
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

	for _, tx := range txs {
		if err := db.UpdateReOccurringTransaction(ctx, userId, &tx.Id, tx); err != nil {
			db.Logger.Error(err.Error())
		}
	}

	return nil
}

// GetReOcurringTransactionById gets a singular reoccurring transaction by its id
func (db *Db) GetReOcurringTransactionById(ctx context.Context, txId *string) (*schema.ReOccuringTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-reocurring-transaction-by-id"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	tx := new(schema.ReOccuringTransactionInternal)
	if err := db.queryEngine.NewSelect().Model(tx).Where("ID = ?", *txId).Scan(ctx); err != nil {
		return nil, err
	}

	res, err := tx.ConvertToRecurringTransaction()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *Db) GetUserReOccurringTransactions(ctx context.Context, userId *uint64) ([]*schema.ReOccuringTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-reocurring-transaction-for-user"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("userId must not be nil")
	}

	var result []schema.ReOccuringTransactionInternal
	sqlQuery := fmt.Sprintf("SELECT * FROM ReOccuringTransactionInternal WHERE UserId = %d", *userId)
	if err := db.queryEngine.
		NewRaw(sqlQuery).
		Scan(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []*schema.ReOccuringTransaction{}, nil
	}

	results := make([]*schema.ReOccuringTransaction, 0, len(result))
	for _, tx := range result {
		txRecord, err := tx.ConvertToRecurringTransaction()
		if err != nil {
			return nil, err
		}
		results = append(results, txRecord)
	}

	return results, nil
}
