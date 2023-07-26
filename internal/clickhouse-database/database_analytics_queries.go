package clickhousedatabase

import (
	"context"
	"fmt"
	"time"

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
	selectQuery := db.queryEngine.NewSelect().Model(&incomeExpenseRatios).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&incomeMetricss).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&merchantMonthlyExpenditures).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyExpenditures).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyIncomes).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&monthlySavings).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyBalances).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&monthlyTransactionCount).Offset(offset).Limit(queryLimit)
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
	selectQuery := db.queryEngine.NewSelect().Model(&paymentChannelMonthlyExpenditures).Offset(offset).Limit(queryLimit)
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

	// query for all the profile based on the user ID
	if err := db.queryEngine.NewSelect().Model(&categoryProfiles).Where("UserId = ? AND Month = ?", *userId, getCurrentMonthRepresentation()).Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&expenseProfiles).Where("UserId = ? AND Month = ?", *userId, getCurrentMonthRepresentation()).Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&incomeProfiles).Where("UserId = ? AND Month = ?", *userId, getCurrentMonthRepresentation()).Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&locationProfiles).Where("UserId = ? AND Month = ?", *userId, getCurrentMonthRepresentation()).Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&merchantProfiles).Where("UserId = ? AND Month = ?", *userId, getCurrentMonthRepresentation()).Scan(ctx); err != nil {
		return nil, err
	}

	if err := db.queryEngine.NewSelect().Model(&paymentChannelsProfiles).Where("UserId = ? AND Month = ?", *userId, getCurrentMonthRepresentation()).Scan(ctx); err != nil {
		return nil, err
	}

	// convert all profiles
	categoryProfilesProto := make([]*schema.CategoryMetricsFinancialSubProfile, 0, len(categoryProfiles))
	for idx, record := range categoryProfiles {
		record := record.ConvertToProto()
		categoryProfilesProto = append(categoryProfilesProto, record)

		if idx > limit {
			break
		}
	}

	expenseProfilesProto := make([]*schema.ExpenseMetricsFinancialSubProfileMetrics, 0, len(expenseProfiles))
	for idx, record := range expenseProfiles {
		record := record.ConvertToProto()
		expenseProfilesProto = append(expenseProfilesProto, record)

		if idx > limit {
			break
		}
	}

	incomeProfilesProto := make([]*schema.IncomeMetricsFinancialSubProfile, 0, len(incomeProfiles))
	for idx, record := range incomeProfiles {
		record := record.ConvertToProto()
		incomeProfilesProto = append(incomeProfilesProto, record)

		if idx > limit {
			break
		}
	}

	locationProfilesProto := make([]*schema.LocationFinancialSubProfile, 0, len(locationProfiles))
	for idx, record := range locationProfiles {
		record := record.ConvertToProto()
		locationProfilesProto = append(locationProfilesProto, record)

		if idx > limit {
			break
		}
	}

	merchantProfilesProto := make([]*schema.MerchantMetricsFinancialSubProfile, 0, len(merchantProfiles))
	for idx, record := range merchantProfiles {
		record := record.ConvertToProto()
		merchantProfilesProto = append(merchantProfilesProto, record)
		if idx > limit {
			break
		}
	}

	paymentChannelsProfilesProto := make([]*schema.PaymentChannelMetricsFinancialSubProfile, 0, len(paymentChannelsProfiles))
	for idx, record := range paymentChannelsProfiles {
		record := record.ConvertToProto()
		paymentChannelsProfilesProto = append(paymentChannelsProfilesProto, record)
		if idx > limit {
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

func getCurrentMonthRepresentation() string {
	now := time.Now()

	year, month, _ := now.Date()

	return fmt.Sprintf("%d%02d", year, int(month))
}
