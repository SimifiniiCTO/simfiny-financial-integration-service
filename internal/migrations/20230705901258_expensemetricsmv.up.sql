-- materialized view for common expense metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS ExpenseMetricsMV
TO ExpenseMetrics
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    sum(Sign) as TransactionCount,
    sum(Amount * Sign) as TotalExpenses
FROM
    TransactionInternal
WHERE
    Amount > 0 AND Sign = 1
GROUP BY
    Month,
    PersonalFinanceCategoryPrimary,
    UserId;