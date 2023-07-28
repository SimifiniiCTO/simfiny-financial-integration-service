package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/uptrace/go-clickhouse/ch"
)

type CategoryMonthlyExpenditureParams struct {
	Month                          uint32
	PersonalFinanceCategoryPrimary string
	PageSize                       uint64
	PageNumber                     uint64
}

func buildCategoryMonthlyExpenditureClickHouseQuery(params *CategoryMonthlyExpenditureParams, query *ch.SelectQuery) {
	if params.Month != 0 {
		query = query.Where("Month = ?", params.Month)
	}

	if params.PersonalFinanceCategoryPrimary != "" {
		query = query.Where("PersonalFinanceCategoryPrimary = ?", params.PersonalFinanceCategoryPrimary)
	}
}

func (db *Db) GetCategoryMonthlyExpenditures(ctx context.Context, userId *uint64, params *CategoryMonthlyExpenditureParams) ([]*schema.CategoryMonthlyExpenditure, int64, error) {
	var (
		nextPageNumber                 int64
		categorizedMonthlyExpenditures []schema.CategoryMonthlyExpenditureInternal
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
		Model(&categorizedMonthlyExpenditures).
		Offset(offset).
		Where("UserId = ?", *userId).
		Limit(queryLimit)
	buildCategoryMonthlyExpenditureClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(categorizedMonthlyExpenditures) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.CategoryMonthlyExpenditure, 0, len(categorizedMonthlyExpenditures))
	for _, record := range categorizedMonthlyExpenditures {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}
