-- Create Materialized View for Income Metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeMetricsFinancialSubProfileMV
TO IncomeMetricsFinancialSubProfile
AS
SELECT
    toYYYYMM(Time) as Month,
    -sumIf(Amount * Sign, Time >= now() - INTERVAL 2 WEEK) as IncomeLastTwoWeeks,
    -sumIf(Amount * Sign, Time >= now() - INTERVAL 1 MONTH) as IncomeLastMonth,
    -sumIf(Amount * Sign, Time >= now() - INTERVAL 2 MONTH) as IncomeLastTwoMonths,
    -sumIf(Amount * Sign, Time >= now() - INTERVAL 6 MONTH) as IncomeLastSixMonths,
    -sumIf(Amount * Sign, Time >= now() - INTERVAL 1 YEAR) as IncomeLastYear,
    UserId
FROM
    TransactionInternal
WHERE
    Amount < 0 AND Sign = 1
GROUP BY
    Month, UserId;