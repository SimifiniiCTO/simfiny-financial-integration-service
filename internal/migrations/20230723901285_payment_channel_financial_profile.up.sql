CREATE TABLE IF NOT EXISTS PaymentChannelMetricsFinancialSubProfile
(
    Month UInt32,
    PaymentChannel String,
    TransactionCount UInt64,
    SpentLastWeek Float64,
    SpentLastTwoWeeks Float64,
    SpentLastMonth Float64,
    SpentLastSixMonths Float64,
    SpentLastYear Float64,
    SpentLastTwoYears Float64,
    UserId UInt64
)
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, PaymentChannel);
