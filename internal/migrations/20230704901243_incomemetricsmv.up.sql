-- materialized view that will transform transactions into income metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeMetricsMV
TO IncomeMetrics
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    count() as TransactionCount,
    sum(Amount) as TotalIncome
FROM
    TransactionInternal
WHERE
    Amount > 0
GROUP BY
    Month,
    PersonalFinanceCategoryPrimary,
    UserId;