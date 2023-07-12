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
type TransactionAggregatesByMonth struct {
	ch.CHModel       `ch:"partition:toYYYYMM(time)"`
	Month            time.Time `ch:"type:Date"`
	CategoryId       string    `ch:"type:String"`
	LocationCity     string    `ch:"type:String"`
	PaymentChannel   string    `ch:"type:String"`
	MerchantName     string    `ch:"type:String"`
	TransactionCount uint64    `ch:"type:UInt64"`
	TotalAmount      float64   `ch:"type:Float64"`
}

/*
* This view will aggregate all expenses (transactions where the Amount is less than zero) by category and month.
* If your definition of "expense metrics" is different, you might need to adjust the query accordingly.*
 */
type ExpenseMetrics struct {
	ch.CHModel       `ch:"partition:toYYYYMM(time)"`
	Month            time.Time `ch:"type:Date"`
	CategoryId       string    `ch:"type:String"`
	TransactionCount uint64    `ch:"type:UInt64"`
	TotalExpenses    float64   `ch:"type:Float64"`
}
