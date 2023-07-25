package clickhousedatabase

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/uptrace/go-clickhouse/ch"
)

type TransactionAggregatedParams struct {
	Month                   uint32
	PersonalFinanceCategory string
	LocationCity            string
	PaymentChannel          string
	MerchantName            string
	PageSize                uint64
	PageNumber              uint64
}

func buildTransactionAggregatesClickHouseQuery(params *TransactionAggregatedParams, query *ch.SelectQuery) {
	if params.Month != 0 {
		query = query.Where("Month = ?", params.Month)
	}

	if params.PersonalFinanceCategory != "" {
		query = query.Where("PersonalFinanceCategoryPrimary = ?", params.PersonalFinanceCategory)
	}

	if params.LocationCity != "" {
		query = query.Where("LocationCity = ?", params.LocationCity)
	}

	if params.PaymentChannel != "" {
		query = query.Where("PaymentChannel = ?", params.PaymentChannel)
	}

	if params.MerchantName != "" {
		query = query.Where("MerchantName = ?", params.MerchantName)
	}
}

func (db *Db) GetTransactionAggregates(ctx context.Context, userId *uint64, params *TransactionAggregatedParams) ([]*schema.TransactionAggregatesByMonth, int64, error) {
	var (
		nextPageNumber        int64
		transactionAggregates []schema.TransactionAggregatesByMonthInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-transaction-aggregates"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&transactionAggregates).Offset(offset).Limit(queryLimit)
	buildTransactionAggregatesClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(transactionAggregates) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.TransactionAggregatesByMonth, 0, len(transactionAggregates))
	for _, record := range transactionAggregates {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}
