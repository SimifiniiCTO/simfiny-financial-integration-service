-- Create Materialized View for Payment Channel Metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS PaymentChannelFinancialSubProfileMetricsMV
TO PaymentChannelMetricsFinancialSubProfile
AS
SELECT
    toYYYYMM(Time) as Month,
    PaymentChannel,
    UserId,
    count() as TransactionCount,
    sumIf(Amount, Time >= now() - INTERVAL 1 WEEK) as SpentLastWeek,
    sumIf(Amount, Time >= now() - INTERVAL 2 WEEK) as SpentLastTwoWeeks,
    sumIf(Amount, Time >= now() - INTERVAL 1 MONTH) as SpentLastMonth,
    sumIf(Amount, Time >= now() - INTERVAL 6 MONTH) as SpentLastSixMonths,
    sumIf(Amount, Time >= now() - INTERVAL 1 YEAR) as SpentLastYear,
    sumIf(Amount, Time >= now() - INTERVAL 2 YEAR) as SpentLastTwoYears
FROM
    TransactionInternal
WHERE
    Amount > 0
GROUP BY
    Month, PaymentChannel, UserId;