package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/uptrace/go-clickhouse/ch"
)

type DebtToIncomeParams struct {
	Month      uint32
	PageSize   uint64
	PageNumber uint64
}

type ExpenseMetricsParams struct {
	Month                          uint32
	PersonalFinanceCategoryPrimary string
	PageSize                       uint64
	PageNumber                     uint64
}

type FinancialProfileParams struct {
	Month      uint32
	PageSize   uint64
	PageNumber uint64
}

func (db *Db) GetDebtToIncomeRatio(ctx context.Context, userId *uint64, params *DebtToIncomeParams) ([]*schema.DebtToIncomeRatio, int64, error) {
	var (
		nextPageNumber       int64
		debtToIncomeRatios   []schema.DebtToIncomeRatioInternal
		buildClickHouseQuery = func(params *DebtToIncomeParams, query *ch.SelectQuery) {
			if params.Month != 0 {
				query = query.Where("Month = ?", params.Month)
			}
		}
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-debt-to-income-ratio"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&debtToIncomeRatios).Offset(offset).Limit(queryLimit)
	buildClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(debtToIncomeRatios) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.DebtToIncomeRatio, 0, len(debtToIncomeRatios))
	for _, record := range debtToIncomeRatios {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetExpenseMetrics(ctx context.Context, userId *uint64, params *ExpenseMetricsParams) ([]*schema.ExpenseMetrics, int64, error) {
	var (
		nextPageNumber       int64
		expenseMetrics       []schema.ExpenseMetricsInternal
		buildClickHouseQuery = func(params *ExpenseMetricsParams, query *ch.SelectQuery) {
			if params.Month != 0 {
				query = query.Where("Month = ?", params.Month)
			}

			if params.PersonalFinanceCategoryPrimary != "" {
				query = query.Where("PersonalFinanceCategoryPrimary = ?", params.PersonalFinanceCategoryPrimary)
			}
		}
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-debt-to-income-ratio"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&expenseMetrics).Offset(offset).Limit(queryLimit)
	buildClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(expenseMetrics) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.ExpenseMetrics, 0, len(expenseMetrics))
	for _, record := range expenseMetrics {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetFinancialProfile(ctx context.Context, userId *uint64, params *FinancialProfileParams) ([]*schema.FinancialProfile, int64, error) {
	var (
		nextPageNumber       int64
		financialProfile     []schema.FinancialProfileInternal
		buildClickHouseQuery = func(params *FinancialProfileParams, query *ch.SelectQuery) {
			if params.Month != 0 {
				query = query.Where("Month = ?", params.Month)
			}
		}
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-financial-profile"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&financialProfile).Offset(offset).Limit(queryLimit)
	buildClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(financialProfile) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.FinancialProfile, 0, len(financialProfile))
	for _, record := range financialProfile {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}
