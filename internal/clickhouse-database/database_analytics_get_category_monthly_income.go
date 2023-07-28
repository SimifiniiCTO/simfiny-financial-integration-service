package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/uptrace/go-clickhouse/ch"
)

type CategoryMonthlyIncomeParams struct {
	Month                          uint32
	PersonalFinanceCategoryPrimary string
	PageSize                       uint64
	PageNumber                     uint64
}

func buildCategoryMonthlyIncomeClickHouseQuery(params *CategoryMonthlyIncomeParams, query *ch.SelectQuery) {
	if params.Month != 0 {
		query = query.Where("Month = ?", params.Month)
	}

	if params.PersonalFinanceCategoryPrimary != "" {
		query = query.Where("PersonalFinanceCategoryPrimary = ?", params.PersonalFinanceCategoryPrimary)
	}
}

func (db *Db) GetCategoryMonthlyIncome(ctx context.Context, userId *uint64, params *CategoryMonthlyIncomeParams) ([]*schema.CategoryMonthlyIncome, int64, error) {
	var (
		nextPageNumber           int64
		categorizedMonthlyIncome []schema.CategoryMonthlyIncomeInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-category-monthly-expenditures"); span != nil {
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
	selectQuery := db.queryEngine.
		NewSelect().
		Model(&categorizedMonthlyIncome).
		Where("UserId = ?", *userId).
		Offset(offset).
		Limit(queryLimit)
	buildCategoryMonthlyIncomeClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(categorizedMonthlyIncome) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.CategoryMonthlyIncome, 0, len(categorizedMonthlyIncome))
	for _, record := range categorizedMonthlyIncome {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}
