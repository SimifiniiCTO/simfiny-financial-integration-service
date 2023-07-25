package financial_integration_service_apiv1

import (
	time "time"

	"github.com/uptrace/go-clickhouse/ch"
)

/*
* In this view, we're counting the number of transactions and summing the amount for each category,
* city, and payment channel for each month. To get the aggregates for the past week, two weeks, month,
* six months, one year, or two years, you would filter this table in your queries. However, remember
* that weeks do not align perfectly with months, so the results may not be exactly what you expect.
 */
type TransactionAggregatesByMonthInternal struct {
	ch.CHModel                     `ch:"TransactionAggregatesByMonth,partition:toYYYYMM(time)"`
	Month                          uint32  `ch:"Month,type:UInt32"`
	PersonalFinanceCategoryPrimary string  `ch:"PersonalFinanceCategoryPrimary,type:String"`
	LocationCity                   string  `ch:"LocationCity,type:String"`
	PaymentChannel                 string  `ch:"PaymentChannel,type:String"`
	MerchantName                   string  `ch:"MerchantName,type:String"`
	TransactionCount               uint64  `ch:"TransactionCount,type:UInt64"`
	TotalAmount                    float64 `ch:"TotalAmount,type:Float64"`
	UserId                         uint64  `ch:"UserId,type:UInt64"`
}

func (source *TransactionAggregatesByMonthInternal) ConvertToORM() *TransactionAggregatesByMonthORM {
	return &TransactionAggregatesByMonthORM{
		LocationCity:                   source.LocationCity,
		MerchantName:                   source.MerchantName,
		Month:                          source.Month,
		PaymentChannel:                 source.PaymentChannel,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalAmount:                    source.TotalAmount,
		TransactionCount:               source.TransactionCount,
		UserId:                         source.UserId,
	}
}

func (source *TransactionAggregatesByMonthORM) ConvertToInternal() *TransactionAggregatesByMonthInternal {
	return &TransactionAggregatesByMonthInternal{
		LocationCity:                   source.LocationCity,
		MerchantName:                   source.MerchantName,
		Month:                          source.Month,
		PaymentChannel:                 source.PaymentChannel,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalAmount:                    source.TotalAmount,
		TransactionCount:               source.TransactionCount,
		UserId:                         source.UserId,
	}
}

func (source *TransactionAggregatesByMonthInternal) ConvertToProto() *TransactionAggregatesByMonth {
	return &TransactionAggregatesByMonth{
		LocationCity:                   source.LocationCity,
		MerchantName:                   source.MerchantName,
		Month:                          source.Month,
		PaymentChannel:                 source.PaymentChannel,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalAmount:                    source.TotalAmount,
		TransactionCount:               source.TransactionCount,
		UserId:                         source.UserId,
	}
}

type CategoryMonthlyExpenditureInternal struct {
	ch.CHModel                     `ch:"CategoryMonthlyExpenditure,partition:toYYYYMM(time)"`
	Month                          uint32  `ch:"Month,type:UInt32"`
	PersonalFinanceCategoryPrimary string  `ch:"PersonalFinanceCategoryPrimary,type:String"`
	TotalSpending                  float64 `ch:"TotalSpending,type:Float64"`
	UserId                         uint64  `ch:"UserId,type:UInt64"`
}

func (source *CategoryMonthlyExpenditureInternal) ConvertToORM() *CategoryMonthlyExpenditureORM {
	return &CategoryMonthlyExpenditureORM{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalSpending:                  source.TotalSpending,
		UserId:                         source.UserId,
	}
}

func (source *CategoryMonthlyExpenditureORM) ConvertToInternal() *CategoryMonthlyExpenditureInternal {
	return &CategoryMonthlyExpenditureInternal{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalSpending:                  source.TotalSpending,
		UserId:                         source.UserId,
	}
}

func (source *CategoryMonthlyExpenditureInternal) ConvertToProto() *CategoryMonthlyExpenditure {
	return &CategoryMonthlyExpenditure{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalSpending:                  source.TotalSpending,
		UserId:                         source.UserId,
	}
}

type CategoryMonthlyIncomeInternal struct {
	ch.CHModel                     `ch:"CategoryMonthlyIncome,partition:toYYYYMM(time)"`
	Month                          uint32  `ch:"Month,type:UInt32"`
	PersonalFinanceCategoryPrimary string  `ch:"PersonalFinanceCategoryPrimary,type:String"`
	TotalIncome                    float64 `ch:"TotalIncome,type:Float64"`
	UserId                         uint64  `ch:"UserId,type:UInt64"`
}

func (source *CategoryMonthlyIncomeInternal) ConvertToORM() *CategoryMonthlyIncomeORM {
	return &CategoryMonthlyIncomeORM{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalIncome:                    source.TotalIncome,
		UserId:                         source.UserId,
	}
}

func (source *CategoryMonthlyIncomeORM) ConvertToInternal() *CategoryMonthlyIncomeInternal {
	return &CategoryMonthlyIncomeInternal{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalIncome:                    source.TotalIncome,
		UserId:                         source.UserId,
	}
}

func (source *CategoryMonthlyIncomeInternal) ConvertToProto() *CategoryMonthlyIncome {
	return &CategoryMonthlyIncome{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		TotalIncome:                    source.TotalIncome,
		UserId:                         source.UserId,
	}
}

/*
* This view will aggregate all expenses (transactions where the Amount is less than zero) by category and month.
* If your definition of "expense metrics" is different, you might need to adjust the query accordingly.*
 */
type ExpenseMetricsInternal struct {
	ch.CHModel       `ch:"partition:toYYYYMM(time)"`
	Month            time.Time `ch:"type:Date"`
	CategoryId       string    `ch:"type:String"`
	TransactionCount uint64    `ch:"type:UInt64"`
	TotalExpenses    float64   `ch:"type:Float64"`
}

type CategoryMetricsFinancialSubProfileInternal struct {
	ch.CHModel                     `ch:"partition:toYYYYMM(time)"`
	Month                          uint32  `ch:"type:UInt32"`
	PersonalFinanceCategoryPrimary string  `ch:"type:String"`
	TransactionCount               uint64  `ch:"type:UInt64"`
	SpentLastWeek                  float64 `ch:"type:Float64"`
	SpentLastTwoWeeks              float64 `ch:"type:Float64"`
	SpentLastMonth                 float64 `ch:"type:Float64"`
	SpentLastSixMonths             float64 `ch:"type:Float64"`
	SpentLastYear                  float64 `ch:"type:Float64"`
	SpentLastTwoYears              float64 `ch:"type:Float64"`
	UserId                         uint64  `ch:"type:UInt64"`
}

func (source *CategoryMetricsFinancialSubProfileInternal) ConvertToORM() *CategoryMetricsFinancialSubProfileORM {
	return &CategoryMetricsFinancialSubProfileORM{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		SpentLastMonth:                 source.SpentLastMonth,
		SpentLastSixMonths:             source.SpentLastSixMonths,
		SpentLastTwoWeeks:              source.SpentLastTwoWeeks,
		SpentLastTwoYears:              source.SpentLastTwoYears,
		SpentLastWeek:                  source.SpentLastWeek,
		SpentLastYear:                  source.SpentLastYear,
		TransactionCount:               source.TransactionCount,
		UserId:                         source.UserId,
	}
}

func (source *CategoryMetricsFinancialSubProfileORM) ConvertToInternal() *CategoryMetricsFinancialSubProfileInternal {
	return &CategoryMetricsFinancialSubProfileInternal{
		Month:                          source.Month,
		PersonalFinanceCategoryPrimary: source.PersonalFinanceCategoryPrimary,
		SpentLastMonth:                 source.SpentLastMonth,
		SpentLastSixMonths:             source.SpentLastSixMonths,
		SpentLastTwoWeeks:              source.SpentLastTwoWeeks,
		SpentLastTwoYears:              source.SpentLastTwoYears,
		SpentLastWeek:                  source.SpentLastWeek,
		SpentLastYear:                  source.SpentLastYear,
		TransactionCount:               source.TransactionCount,
		UserId:                         source.UserId,
	}
}
