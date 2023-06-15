package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func (db *Db) AddInvestmentTransactions(ctx context.Context, userId *uint64, txs []*schema.InvestmentTransaction) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-investment-add-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil
	}

	if txs == nil {
		return nil
	}

	transactions := make([]*schema.InvestmentTransactionORM, 0, len(txs))
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

		transactions = append(transactions, &record)
	}

	t := db.queryOperator.InvestmentTransactionORM
	if err := t.WithContext(ctx).Create(transactions...); err != nil {
		return err
	}

	return nil
}

// DeleteTransaction implements ClickhouseDatabaseOperations.
func (db *Db) DeleteInvestmentTransactions(ctx context.Context, txIds ...uint64) error {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-investment-transaction"); span != nil {
		defer span.End()
	}

	if len(txIds) == 0 {
		return fmt.Errorf("transaction ID must be 0 at creation time")
	}

	// delete the investment transactions by id
	t := db.queryOperator.InvestmentTransactionORM
	result, err := t.WithContext(ctx).Where(t.Id.In(txIds...)).Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

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

	t := db.queryOperator.InvestmentTransactionORM
	var records []*schema.InvestmentTransactionORM
	records, err := t.WithContext(ctx).
		Where(t.UserId.Eq(*userId)).
		Limit(int(pageSize)).Offset(int(pageSize * (pageNumber - 1))).
		Find()
	if err != nil {
		return nil, 0, err
	}

	txs := make([]*schema.InvestmentTransaction, 0, len(records))
	for _, record := range records {
		tx, err := record.ToPB(ctx)
		if err != nil {
			return nil, 0, err
		}

		txs = append(txs, &tx)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetAllInvestmentTransactions(ctx context.Context, userId *uint64) ([]*schema.InvestmentTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-investment-transactions"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID is nil")
	}

	t := db.queryOperator.InvestmentTransactionORM
	var records []*schema.InvestmentTransactionORM
	records, err := t.WithContext(ctx).
		Where(t.UserId.Eq(*userId)).
		Find()
	if err != nil {
		return nil, err
	}

	txs := make([]*schema.InvestmentTransaction, 0, len(records))
	for _, record := range records {
		tx, err := record.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		txs = append(txs, &tx)
	}

	return txs, nil
}

// UpdateTransaction implements ClickhouseDatabaseOperations.
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

	txnsOrmRecords := make([]*schema.InvestmentTransactionORM, 0, len(txs))
	for _, tx := range txs {
		tx.UserId = *userId

		txnOrm, err := tx.ToORM(ctx)
		if err != nil {
			return err
		}

		txnsOrmRecords = append(txnsOrmRecords, &txnOrm)
	}

	t := db.queryOperator.InvestmentTransactionORM
	for _, tx := range txnsOrmRecords {
		if tx.Id == 0 {
			return fmt.Errorf("transaction ID must be 0 at creation time")
		}

		// update the transaction
		if _, err := t.WithContext(ctx).Updates(tx); err != nil {
			return err
		}
	}

	return nil
}

func (db *Db) GetInvestmentTransactionById(ctx context.Context, txId *uint64) (*schema.InvestmentTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-InvestmentTransaction-by-id"); span != nil {
		defer span.End()
	}

	if txId == nil {
		return nil, fmt.Errorf("InvestmentTransaction ID must be 0 at creation time")
	}

	t := db.queryOperator.InvestmentTransactionORM
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

func (db *Db) GetInvestmentTransactionsByPlaidTransactionIds(ctx context.Context, txIds []string) ([]*schema.InvestmentTransaction, error) {
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-transactions-by-ids"); span != nil {
		defer span.End()
	}

	if txIds == nil {
		return nil, fmt.Errorf("transaction ID must be 0 at creation time")
	}

	//	get the transacton by tx id
	t := db.queryOperator.InvestmentTransactionORM
	// delete the transaction
	result, err := t.WithContext(ctx).Where(t.InvestmentTransactionId.In(txIds...)).Find()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("no txn found")
	}

	resultSet := make([]*schema.InvestmentTransaction, 0, len(result))
	for _, tx := range result {
		txRecord, err := tx.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		resultSet = append(resultSet, &txRecord)
	}

	return resultSet, nil
}
