-- Create Materialized View for Income Metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeMetricsFinancialSubProfileMV
TO IncomeMetricsFinancialSubProfile
AS
SELECT
    toYYYYMM(Time) as Month,
    -sumIf(Amount, Time >= now() - INTERVAL 2 WEEK) as IncomeLastTwoWeeks,
    -sumIf(Amount, Time >= now() - INTERVAL 1 MONTH) as IncomeLastMonth,
    -sumIf(Amount, Time >= now() - INTERVAL 2 MONTH) as IncomeLastTwoMonths,
    -sumIf(Amount, Time >= now() - INTERVAL 6 MONTH) as IncomeLastSixMonths,
    -sumIf(Amount, Time >= now() - INTERVAL 1 YEAR) as IncomeLastYear
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month;