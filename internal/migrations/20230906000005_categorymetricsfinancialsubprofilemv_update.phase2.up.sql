CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMetricsFinancialSubProfileMV
TO CategoryMetricsFinancialSubProfile
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    sum(Sign) as TransactionCount,
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
    Month, PersonalFinanceCategoryPrimary, UserId;
