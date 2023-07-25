CREATE TABLE IF NOT EXISTS IncomeMetricsFinancialSubProfile
(
    Month UInt32,
    IncomeLastTwoWeeks Float64,
    IncomeLastMonth Float64,
    IncomeLastTwoMonths Float64,
    IncomeLastSixMonths Float64,
    IncomeLastYear Float64
)
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;