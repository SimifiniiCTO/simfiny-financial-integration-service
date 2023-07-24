CREATE TABLE IF NOT EXISTS MerchantMetricsFinancialSubProfile
(
    Month UInt32,
    MerchantName String,
    TransactionsLastWeek UInt64,
    TransactionsLastTwoWeeks UInt64,
    TransactionsLastMonth UInt64,
    TransactionsLastSixMonths UInt64,
    TransactionsLastYear UInt64,
    TransactionsLastTwoYears UInt64,
    SpentLastWeek Float64,
    SpentLastTwoWeeks Float64,
    SpentLastMonth Float64,
    SpentLastSixMonths Float64,
    SpentLastYear Float64,
    SpentLastTwoYears Float64
)
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, MerchantName);