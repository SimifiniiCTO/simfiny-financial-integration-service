-- Create Materialized View for Merchant Metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS MerchantMetricsFinancialSubProfileMV
TO MerchantMetricsFinancialSubProfile
AS
SELECT
    toYYYYMM(Time) as Month,
    MerchantName,
    UserId,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 WEEK) as SpentLastWeek,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 2 WEEK) as SpentLastTwoWeeks,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 MONTH) as SpentLastMonth,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 6 MONTH) as SpentLastSixMonths,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 YEAR) as SpentLastYear,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 2 YEAR) as SpentLastTwoYears
FROM
    TransactionInternal
WHERE
    Amount > 0 AND Sign = 1
GROUP BY
    Month, MerchantName, UserId;