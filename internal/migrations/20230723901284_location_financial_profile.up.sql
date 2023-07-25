CREATE TABLE IF NOT EXISTS LocationFinancialSubProfile
(
    Month UInt32,
    LocationCity String,
    TransactionCount UInt64,
    SpentLastWeek Float64,
    SpentLastTwoWeeks Float64,
    SpentLastMonth Float64,
    SpentLastSixMonths Float64,
    SpentLastYear Float64,
    SpentLastTwoYears Float64
)
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, LocationCity);
