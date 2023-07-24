CREATE TABLE IF NOT EXISTS ExpenseMetricsFinancialSubProfileMetrics
(
    Month UInt32,
    SpentLastWeek Float64,
    SpentLastMonth Float64,
    SpentLastSixMonths Float64,
    AverageMonthlyDiscretionarySpending Float64,
    AverageMonthlyRecurringSpending Float64
)
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;
