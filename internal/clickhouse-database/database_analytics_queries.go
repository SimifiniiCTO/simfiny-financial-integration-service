package clickhousedatabase

import (
	"context"
	"fmt"
	"time"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/hashicorp/go-set"
	"github.com/uptrace/go-clickhouse/ch"
	"go.uber.org/zap"
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

type IncomeExpenseMetricsParams struct {
	Month      uint32
	PageSize   uint64
	PageNumber uint64
}

type IncomeMetricsParams struct {
	Month                          uint32
	PersonalFinanceCategoryPrimary string
	PageSize                       uint64
	PageNumber                     uint64
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
	selectQuery := db.queryEngine.NewSelect().Model(&debtToIncomeRatios).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&expenseMetrics).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&financialProfile).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
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

func (db *Db) GetIncomeExpenseRatio(ctx context.Context, userId *uint64, params *IncomeExpenseMetricsParams) ([]*schema.IncomeExpenseRatio, int64, error) {
	var (
		nextPageNumber       int64
		incomeExpenseRatios  []schema.IncomeExpenseRatioInternal
		buildClickHouseQuery = func(params *IncomeExpenseMetricsParams, query *ch.SelectQuery) {
			if params.Month != 0 {
				query = query.Where("Month = ?", params.Month)
			}
		}
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-income-expense-ratio"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&incomeExpenseRatios).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	buildClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(incomeExpenseRatios) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.IncomeExpenseRatio, 0, len(incomeExpenseRatios))
	for _, record := range incomeExpenseRatios {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetIncomeMetrics(ctx context.Context, userId *uint64, params *IncomeMetricsParams) ([]*schema.IncomeMetrics, int64, error) {
	var (
		nextPageNumber       int64
		incomeMetricss       []schema.IncomeMetricsInternal
		buildClickHouseQuery = func(params *IncomeMetricsParams, query *ch.SelectQuery) {
			if params.Month != 0 {
				query = query.Where("Month = ?", params.Month)
			}

			if params.PersonalFinanceCategoryPrimary != "" {
				query = query.Where("PersonalFinanceCategoryPrimary = ?", params.PersonalFinanceCategoryPrimary)
			}
		}
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-income-metrics"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&incomeMetricss).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	buildClickHouseQuery(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(incomeMetricss) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.IncomeMetrics, 0, len(incomeMetricss))
	for _, record := range incomeMetricss {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

type BaseParams struct {
	Month          uint32
	MerchantName   string
	PaymentChannel string
	PageSize       uint64
	PageNumber     uint64
}

func clickhouseQueryBuilder(params *BaseParams, query *ch.SelectQuery) {
	if params.Month != 0 {
		query = query.Where("Month = ?", params.Month)
	}

	if params.MerchantName != "" {
		query = query.Where("MerchantName = ?", params.MerchantName)
	}

	if params.PaymentChannel != "" {
		query = query.Where("PaymentChannel = ?", params.PaymentChannel)
	}
}

func (db *Db) GetMerchantMonthlyExpenditure(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.MerchantMonthlyExpenditure, int64, error) {
	var (
		nextPageNumber              int64
		merchantMonthlyExpenditures []schema.MerchantMonthlyExpenditureInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-merchant-monthly-expenditure"); span != nil {
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
	selectQuery := db.queryEngine.NewSelect().Model(&merchantMonthlyExpenditures).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(merchantMonthlyExpenditures) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.MerchantMonthlyExpenditure, 0, len(merchantMonthlyExpenditures))
	for _, record := range merchantMonthlyExpenditures {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetMonthlyExpenditure(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.MonthlyExpenditure, int64, error) {
	var (
		nextPageNumber      int64
		monthlyExpenditures []schema.MonthlyExpenditureInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-expenditure"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	// ensure only the month is used for the query
	params.MerchantName = ""

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyExpenditures).Offset(offset).Where("UserId = ?", *userId).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(monthlyExpenditures) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.MonthlyExpenditure, 0, len(monthlyExpenditures))
	for _, record := range monthlyExpenditures {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetMonthlyIncome(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.MonthlyIncome, int64, error) {
	var (
		nextPageNumber int64
		monthlyIncomes []schema.MonthlyIncomeInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-income"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	// ensure only the month is used for the query
	params.MerchantName = ""

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyIncomes).Offset(offset).Where("UserId = ?", *userId).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(monthlyIncomes) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.MonthlyIncome, 0, len(monthlyIncomes))
	for _, record := range monthlyIncomes {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetMonthlySavings(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.MonthlySavings, int64, error) {
	var (
		nextPageNumber int64
		monthlySavings []schema.MonthlySavingsInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-savings"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	// ensure only the month is used for the query
	params.MerchantName = ""

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&monthlySavings).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(monthlySavings) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.MonthlySavings, 0, len(monthlySavings))
	for _, record := range monthlySavings {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetMonthlyBalance(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.MonthlyBalance, int64, error) {
	var (
		nextPageNumber  int64
		monthlyBalances []schema.MonthlyBalanceInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-balance"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	// ensure only the month is used for the query
	params.MerchantName = ""

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyBalances).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(monthlyBalances) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.MonthlyBalance, 0, len(monthlyBalances))
	for _, record := range monthlyBalances {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetMonthlyTransactionCount(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.MonthlyTransactionCount, int64, error) {
	var (
		nextPageNumber          int64
		monthlyTransactionCount []schema.MonthlyTransactionCountInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-transaction-count"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	// ensure only the month is used for the query
	params.MerchantName = ""

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyTransactionCount).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(monthlyTransactionCount) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.MonthlyTransactionCount, 0, len(monthlyTransactionCount))
	for _, record := range monthlyTransactionCount {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

func (db *Db) GetMonthlyPaymentChannelExpenditure(ctx context.Context, userId *uint64, params *BaseParams) ([]*schema.PaymentChannelMonthlyExpenditure, int64, error) {
	var (
		nextPageNumber                    int64
		paymentChannelMonthlyExpenditures []schema.PaymentChannelMonthlyExpenditureInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-payment-channel-expenditures"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, 0, fmt.Errorf("user ID cannot be nil")
	}

	if params == nil {
		return nil, 0, fmt.Errorf("params cannot be nil")
	}

	// ensure only the month is used for the query
	params.MerchantName = ""

	pageNumber, pageSize := db.sanitizePaginationParams(int64(params.PageNumber), int64(params.PageSize))
	if pageNumber == 0 {
		nextPageNumber = 2
	} else {
		nextPageNumber = pageNumber + 1
	}

	offset := int(pageSize * (pageNumber - 1))
	queryLimit := int(pageSize)
	selectQuery := db.queryEngine.NewSelect().Model(&paymentChannelMonthlyExpenditures).Where("UserId = ?", *userId).Offset(offset).Limit(queryLimit)
	clickhouseQueryBuilder(params, selectQuery)
	// sort by month in descending order
	if err := selectQuery.Order("Month DESC").Scan(ctx); err != nil {
		return nil, 0, err
	}

	if len(paymentChannelMonthlyExpenditures) == 0 {
		return nil, 0, fmt.Errorf("no records found")
	}

	txs := make([]*schema.PaymentChannelMonthlyExpenditure, 0, len(paymentChannelMonthlyExpenditures))
	for _, record := range paymentChannelMonthlyExpenditures {
		record := record.ConvertToProto()
		txs = append(txs, record)
	}

	return txs, nextPageNumber, nil
}

// TODO: refactor this with order by caluse
func (db *Db) GetFinancialContextForCurrentMonth(ctx context.Context, userId *uint64, limit int) (*schema.MelodyFinancialContext, error) {
	var (
		financialContext        *schema.MelodyFinancialContext
		categoryProfiles        []*schema.CategoryMetricsFinancialSubProfileInternal
		expenseProfiles         []*schema.ExpenseMetricsFinancialSubProfileMetricsInternal
		incomeProfiles          []*schema.IncomeMetricsFinancialSubProfileInternal
		locationProfiles        []*schema.LocationFinancialSubProfileInternal
		merchantProfiles        []*schema.MerchantMetricsFinancialSubProfileInternal
		paymentChannelsProfiles []*schema.PaymentChannelMetricsFinancialSubProfileInternal
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-payment-channel-expenditures"); span != nil {
		defer span.End()
	}

	if userId == nil {
		return nil, fmt.Errorf("user ID cannot be nil")
	}

	if limit == 0 {
		limit = 4
	} else if limit > 5 {
		limit = 4
	}

	db.Logger.Info("limit", zap.Int("limit", limit))
	db.Logger.Info("month", zap.String("month", getCurrentMonthRepresentation()))

	// we first get the current months representation
	// TODO: refactor this
	month := getPreviousMonthRepresentation()

	// we use the current months numerical representation to the query the database. If no results are provided
	// it is safe to assume that the sourced context for the current month has not yet been generated
	// hence we use the previous months representation
	if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	// if no results are found for the current month, we use the previous months representation
	if len(categoryProfiles) == 0 {
		month = getPreviousMonthRepresentation()
		if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("TransactionCount DESC").Scan(ctx); err != nil {
			return nil, err
		}
	}

	// query for all the profile based on the user ID
	if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&expenseProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("SpentLastMonth DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&incomeProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("IncomeLastMonth DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&locationProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&merchantProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("SpentLastMonth DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&paymentChannelsProfiles).Where("UserId = ? AND Month = ?", *userId, month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	// convert all profiles
	categoryProfilesProto := make([]*schema.CategoryMetricsFinancialSubProfile, 0, len(categoryProfiles))
	for idx, record := range categoryProfiles {
		record := record.ConvertToProto()
		categoryProfilesProto = append(categoryProfilesProto, record)

		if idx >= limit {
			break
		}
	}

	expenseProfilesProto := make([]*schema.ExpenseMetricsFinancialSubProfileMetrics, 0, len(expenseProfiles))
	for idx, record := range expenseProfiles {
		record := record.ConvertToProto()
		expenseProfilesProto = append(expenseProfilesProto, record)

		if idx >= limit {
			break
		}
	}

	incomeProfilesProto := make([]*schema.IncomeMetricsFinancialSubProfile, 0, len(incomeProfiles))
	for idx, record := range incomeProfiles {
		record := record.ConvertToProto()
		incomeProfilesProto = append(incomeProfilesProto, record)

		if idx >= limit {
			break
		}
	}

	locationProfilesProto := make([]*schema.LocationFinancialSubProfile, 0, len(locationProfiles))
	for idx, record := range locationProfiles {
		record := record.ConvertToProto()
		locationProfilesProto = append(locationProfilesProto, record)

		if idx >= limit {
			break
		}
	}

	merchantProfilesProto := make([]*schema.MerchantMetricsFinancialSubProfile, 0, len(merchantProfiles))
	for idx, record := range merchantProfiles {
		record := record.ConvertToProto()
		merchantProfilesProto = append(merchantProfilesProto, record)
		if idx >= limit {
			break
		}
	}

	paymentChannelsProfilesProto := make([]*schema.PaymentChannelMetricsFinancialSubProfile, 0, len(paymentChannelsProfiles))
	for idx, record := range paymentChannelsProfiles {
		record := record.ConvertToProto()
		paymentChannelsProfilesProto = append(paymentChannelsProfilesProto, record)
		if idx >= limit {
			break
		}
	}

	financialContext = &schema.MelodyFinancialContext{
		Categories:      categoryProfilesProto,
		Expenses:        expenseProfilesProto,
		Income:          incomeProfilesProto,
		Locations:       locationProfilesProto,
		Merchants:       merchantProfilesProto,
		PaymentChannels: paymentChannelsProfilesProto,
	}

	return financialContext, nil
}

// TODO: refactor this with order by caluse
func (db *Db) GetAllFinancialContextsForCurrentMonth(ctx context.Context, limit int) (map[uint64]*schema.MelodyFinancialContext, error) {
	var (
		financialContexts       = make(map[uint64]*schema.MelodyFinancialContext, 0)
		categoryProfiles        []*schema.CategoryMetricsFinancialSubProfileInternal
		expenseProfiles         []*schema.ExpenseMetricsFinancialSubProfileMetricsInternal
		incomeProfiles          []*schema.IncomeMetricsFinancialSubProfileInternal
		locationProfiles        []*schema.LocationFinancialSubProfileInternal
		merchantProfiles        []*schema.MerchantMetricsFinancialSubProfileInternal
		paymentChannelsProfiles []*schema.PaymentChannelMetricsFinancialSubProfileInternal
		// hash maps
		userIdToCategorySubProfile       = make(map[uint64][]*schema.CategoryMetricsFinancialSubProfile, 0)
		userIdToExpensesSubProfile       = make(map[uint64][]*schema.ExpenseMetricsFinancialSubProfileMetrics, 0)
		userIdToIncomesSubProfile        = make(map[uint64][]*schema.IncomeMetricsFinancialSubProfile, 0)
		userIdToLocationSubProfile       = make(map[uint64][]*schema.LocationFinancialSubProfile, 0)
		userIdToMerchantSubProfile       = make(map[uint64][]*schema.MerchantMetricsFinancialSubProfile, 0)
		userIdToPaymentChannelSubProfile = make(map[uint64][]*schema.PaymentChannelMetricsFinancialSubProfile, 0)
	)

	if span := db.startDatastoreSpan(ctx, "dbtxn-get-monthly-payment-channel-expenditures"); span != nil {
		defer span.End()
	}

	if limit == 0 {
		limit = 4
	} else if limit > 5 {
		limit = 4
	}

	// we first get the current months representation
	month := getCurrentMonthRepresentation()

	// we use the current months numerical representation to the query the database. If no results are provided
	// it is safe to assume that the sourced context for the current month has not yet been generated
	// hence we use the previous months representation
	if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("Month = ?", month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	// if no results are found for the current month, we use the previous months representation
	if len(categoryProfiles) == 0 {
		month = getPreviousMonthRepresentation()
		if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("Month = ?", month).Order("TransactionCount DESC").Scan(ctx); err != nil {
			return nil, err
		}
	}

	// query for all the profile based on the user ID
	if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("Month = ?", month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&expenseProfiles).Where("Month = ?", month).Order("SpentLastMonth DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&incomeProfiles).Where("Month = ?", month).Order("IncomeLastMonth DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&locationProfiles).Where("Month = ?", month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&merchantProfiles).Where("Month = ?", month).Order("SpentLastMonth DESC").Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&paymentChannelsProfiles).Where("Month = ?", month).Order("TransactionCount DESC").Scan(ctx); err != nil {
		return nil, err
	}

	uniqueUserIds := set.New[uint64](0)
	// convert all profiles
	for idx, record := range categoryProfiles {
		uniqueUserIds.Insert(record.UserId)
		data := record.ConvertToProto()
		if _, ok := userIdToCategorySubProfile[record.UserId]; !ok {
			set := make([]*schema.CategoryMetricsFinancialSubProfile, 0)
			set = append(set, data)
			userIdToCategorySubProfile[record.UserId] = set
		} else {
			set := userIdToCategorySubProfile[record.UserId]
			set = append(set, data)
			userIdToCategorySubProfile[record.UserId] = set
		}

		if idx >= limit {
			break
		}
	}

	for idx, record := range expenseProfiles {
		uniqueUserIds.Insert(record.UserId)
		data := record.ConvertToProto()
		if _, ok := userIdToExpensesSubProfile[record.UserId]; !ok {
			set := make([]*schema.ExpenseMetricsFinancialSubProfileMetrics, 0)
			set = append(set, data)
			userIdToExpensesSubProfile[record.UserId] = set
		} else {
			set := userIdToExpensesSubProfile[record.UserId]
			set = append(set, data)
			userIdToExpensesSubProfile[record.UserId] = set
		}

		if idx >= limit {
			break
		}
	}

	for idx, record := range incomeProfiles {
		uniqueUserIds.Insert(record.UserId)
		data := record.ConvertToProto()
		if _, ok := userIdToIncomesSubProfile[record.UserId]; !ok {
			set := make([]*schema.IncomeMetricsFinancialSubProfile, 0)
			set = append(set, data)
			userIdToIncomesSubProfile[record.UserId] = set
		} else {
			set := userIdToIncomesSubProfile[record.UserId]
			set = append(set, data)
			userIdToIncomesSubProfile[record.UserId] = set
		}

		if idx >= limit {
			break
		}
	}

	for idx, record := range locationProfiles {
		uniqueUserIds.Insert(record.UserId)
		data := record.ConvertToProto()
		if _, ok := userIdToLocationSubProfile[record.UserId]; !ok {
			set := make([]*schema.LocationFinancialSubProfile, 0)
			set = append(set, data)
			userIdToLocationSubProfile[record.UserId] = set
		} else {
			set := userIdToLocationSubProfile[record.UserId]
			set = append(set, data)
			userIdToLocationSubProfile[record.UserId] = set
		}

		if idx >= limit {
			break
		}
	}

	for idx, record := range merchantProfiles {
		uniqueUserIds.Insert(record.UserId)
		data := record.ConvertToProto()
		if _, ok := userIdToMerchantSubProfile[record.UserId]; !ok {
			set := make([]*schema.MerchantMetricsFinancialSubProfile, 0)
			set = append(set, data)
			userIdToMerchantSubProfile[record.UserId] = set
		} else {
			set := userIdToMerchantSubProfile[record.UserId]
			set = append(set, data)
			userIdToMerchantSubProfile[record.UserId] = set
		}

		if idx >= limit {
			break
		}
	}

	for idx, record := range paymentChannelsProfiles {
		uniqueUserIds.Insert(record.UserId)
		data := record.ConvertToProto()
		if _, ok := userIdToPaymentChannelSubProfile[record.UserId]; !ok {
			set := make([]*schema.PaymentChannelMetricsFinancialSubProfile, 0)
			set = append(set, data)
			userIdToPaymentChannelSubProfile[record.UserId] = set
		} else {
			set := userIdToPaymentChannelSubProfile[record.UserId]
			set = append(set, data)
			userIdToPaymentChannelSubProfile[record.UserId] = set
		}

		if idx >= limit {
			break
		}
	}

	uniqueUserIds.ForEach(func(userId uint64) bool {
		// construct financial context
		financialContext := &schema.MelodyFinancialContext{}

		// extract the category profiles form the userIdToCategoryProfilemap
		if recs, ok := userIdToCategorySubProfile[userId]; ok {
			financialContext.Categories = recs
		}

		if recs, ok := userIdToExpensesSubProfile[userId]; ok {
			financialContext.Expenses = recs
		}

		if recs, ok := userIdToIncomesSubProfile[userId]; ok {
			financialContext.Income = recs
		}

		if recs, ok := userIdToLocationSubProfile[userId]; ok {
			financialContext.Locations = recs
		}

		if recs, ok := userIdToMerchantSubProfile[userId]; ok {
			financialContext.Merchants = recs
		}

		if recs, ok := userIdToPaymentChannelSubProfile[userId]; ok {
			financialContext.PaymentChannels = recs
		}

		financialContexts[userId] = financialContext
		return true
	})

	return financialContexts, nil
}

func getCurrentMonthRepresentation() string {
	now := time.Now()

	year, month, _ := now.Date()

	return fmt.Sprintf("%d%02d", year, int(month))
}

func getPreviousMonthRepresentation() string {
	now := time.Now().AddDate(0, -1, 0)

	year, month, _ := now.Date()

	return fmt.Sprintf("%d%02d", year, int(month))
}
