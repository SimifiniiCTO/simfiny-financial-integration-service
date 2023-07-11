package clickhousedatabase

import (
	"context"
	"fmt"
	"strings"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/google/uuid"
)

// AddInvestmentTransactions is adding investment transactions to the Clickhouse database. It takes in a context, a
// user ID, and a slice of InvestmentTransaction objects. It first checks that the user ID is not nil
// and that the length of the transactions slice is greater than 0. It then creates a slice of
// InvestmentTransactionORM objects from the InvestmentTransaction objects, associates the user ID with
// each transaction, validates each transaction, and checks that the transaction ID is 0 at creation
// time. Finally, it uses the QueryOperator to create the transactions in the database and returns an
// error if there is one.
func (db *Db) AddInvestmentTransactions(ctx context.Context, userId *uint64, txs []*schema.InvestmentTransaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-investment-add-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return fmt.Errorf("user ID cannot be nil")
	}

	if len(txs) == 0 {
		return fmt.Errorf("transactions length must be greater than 0")
	}

	transactions := make([]schema.InvestmentTransactionInternal, 0, len(txs))
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

		transactions = append(transactions, *record)
	}

	if _, err := db.queryEngine.NewInsert().Model(&transactions).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteInvestmentTransactions is deleting investment transactions from the Clickhouse database. It takes in a context
// and a variable number of transaction IDs as uint64 values. It first checks that the length of the
// transaction IDs slice is greater than 0. It then uses the QueryOperator to delete the transactions
// from the database and returns an error if there is one.
func (db *Db) DeleteInvestmentTransactions(ctx context.Context, txIds ...string) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-investment-transaction"); span != nil {
		defer span.End()
	}

	if len(txIds) == 0 {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// Create a string with comma-separated values from txIds
	ids := "'" + strings.Join(txIds, "', '") + "'"

	sqlQuery := fmt.Sprintf("ALTER TABLE InvestmentTransactionInternal DELETE WHERE ID IN (%s)", ids)

	if err := db.queryEngine.
		NewRaw(sqlQuery).
		Scan(ctx); err != nil {
		return err
	}

	return nil
}

// GetInvestmentTransactions is retrieving investment transactions from the Clickhouse database in a paginated manner.
func (db *Db) GetInvestmentTransactions(ctx context.Context, userId *uint64, pageSize int64, pageNumber int64) ([]*schema.InvestmentTransaction, int64, error) {
	var (
		nextPageNumber int64
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-investment-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID is nil")
	}

	pageNumber, pageSize = db.sanitizePaginationParams(pageNumber, pageSize)
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	query := fmt.Sprintf(`UserId = %d`, *userId)
	var transactions []schema.InvestmentTransactionInternal
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
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.InvestmentTransaction, 0, len(transactions))
	for _, record := range transactions {
		txRecord, err := record.ConvertToInvestmentTransaction()
		if err != nil {
			return nil, 0, err
		}
		txs = append(txs, txRecord)
	}

	return txs, nextPageNumber, nil
}

// GetAllInvestmentTransactions is retrieving all investment transactions from the Clickhouse database.
func (db *Db) GetAllInvestmentTransactions(ctx context.Context, userId *uint64) ([]*schema.InvestmentTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-investment-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	query := fmt.Sprintf(`UserId = %d`, *userId)
	var transactions []schema.InvestmentTransactionInternal
	if err := db.
		queryEngine.
		NewSelect().
		Model(&transactions).
		Where(query).
		Scan(ctx); err != nil {
		return nil, err
	}

	txs := make([]*schema.InvestmentTransaction, 0, len(transactions))
	for _, tx := range transactions {
		txRecord, err := tx.ConvertToInvestmentTransaction()
		if err != nil {
			return nil, err
		}
		txs = append(txs, txRecord)
	}

	return txs, nil
}

// UpdateInvestmentTransactions is updating investment transactions in the Clickhouse database.
func (db *Db) UpdateInvestmentTransactions(ctx context.Context, userId *uint64, txs []*schema.InvestmentTransaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-transaction"); span != nil {
		defer span.End()
	}

	if len(txs) == 0 {
		return fmt.Errorf("transactions length must be greater than 0")
	}

	if userId == nil {
		return fmt.Errorf("user ID must be 0 at creation time")
	}

	// txnsOrmRecords := make([]*schema.InvestmentTransactionORM, 0, len(txs))
	// for _, tx := range txs {
	// 	tx.UserId = *userId

	// 	txnOrm, err := tx.ToORM(ctx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	txnsOrmRecords = append(txnsOrmRecords, &txnOrm)
	// }

	// t := db.QueryOperator.InvestmentTransactionORM
	// for _, tx := range txnsOrmRecords {
	// 	if tx.Id == "" {
	// 		return fmt.Errorf("transaction ID must be 0 at creation time")
	// 	}

	// 	// update the transaction
	// 	if _, err := t.WithContext(ctx).Updates(tx); err != nil {
	// 		return err
	// 	}
	// }

	transactions := make([]schema.InvestmentTransactionInternal, 0, len(txs))
	for _, tx := range txs {
		if tx.Id != "" {
			return fmt.Errorf("transaction ID must be 0 at creation time")
		}

		// associate the user id to the transaction of interest
		tx.UserId = *userId

		// validate transactions
		if err := tx.Validate(); err != nil {
			return err
		}

		record, err := tx.ConvertToInternal()
		if err != nil {
			return err
		}

		transactions = append(transactions, *record)
	}

	return nil
}

// GetInvestmentTransactionById is retrieving investment transactions from the Clickhouse database by ID.
func (db *Db) GetInvestmentTransactionById(ctx context.Context, txId *string) (*schema.InvestmentTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-InvestmentTransaction-by-id"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return nil, fmt.Errorf("InvestmentTransaction ID must be 0 at creation time")
	}

	tx := new(schema.InvestmentTransactionInternal)
	if err := db.queryEngine.NewSelect().Model(tx).Where("ID = ?", *txId).Scan(ctx); err != nil {
		return nil, err
	}

	res, err := tx.ConvertToInvestmentTransaction()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetInvestmentTransactionsByPlaidTransactionIds is retrieving investment transactions from the Clickhouse database by Plaid transaction IDs.
func (db *Db) GetInvestmentTransactionsByPlaidTransactionIds(ctx context.Context, txIds []string) ([]*schema.InvestmentTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	var result []schema.InvestmentTransactionInternal
	ids := "'" + strings.Join(txIds, "', '") + "'"
	sqlQuery := fmt.Sprintf("SELECT * FROM InvestmentTransactionInternal WHERE InvestmentTransactionId IN (%s)", ids)
	if err := db.queryEngine.
		NewRaw(sqlQuery).
		Scan(ctx, &result); err != nil {
		return nil, err
	}

	resultSet := make([]*schema.InvestmentTransaction, 0, len(result))
	if len(result) > 0 {
		for _, tx := range result {
			txRecord, err := tx.ConvertToInvestmentTransaction()
			if err != nil {
				return nil, err
			}

			resultSet = append(resultSet, txRecord)
		}
	}

	return resultSet, nil
}
