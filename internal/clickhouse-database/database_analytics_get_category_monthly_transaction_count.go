package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/uptrace/go-clickhouse/ch"
)

type CategoryMonthlyTransactionCountParams struct {
	Month                          uint32
	PersonalFinanceCategoryPrimary string
	PageSize                       uint64
	PageNumber                     uint64
}

func buildCategoryMonthlyTransactionCountClickHouseQuery(params *CategoryMonthlyTransactionCountParams, query *ch.SelectQuery) {
	if params.Month != 0 {
		query = query.Where("Month = ?", params.Month)
	}

	if params.PersonalFinanceCategoryPrimary != "" {
		query = query.Where("PersonalFinanceCategoryPrimary = ?", params.PersonalFinanceCategoryPrimary)
	}
}

func (db *Db) GetCategoryMonthlyTransactionCount(ctx context.Context, userId *uint64, params *CategoryMonthlyTransactionCountParams) ([]*schema.CategoryMonthlyTransactionCount, int64, error) {
	var (
		nextPageNumber                     int64
		categorizedMonthlyTransactionCount []schema.CategoryMonthlyTransactionCountInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-category-monthly-transaction-count"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&categorizedMonthlyTransactionCount).Offset(offset).Limit(queryLimit)
	buildCategoryMonthlyTransactionCountClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(categorizedMonthlyTransactionCount) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.CategoryMonthlyTransactionCount, 0, len(categorizedMonthlyTransactionCount))
	for _, record := range categorizedMonthlyTransactionCount {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}
