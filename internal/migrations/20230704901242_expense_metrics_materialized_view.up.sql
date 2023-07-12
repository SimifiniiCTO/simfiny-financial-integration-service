-- materialized view for common expense metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS ExpenseMetricsMV
TO ExpenseMetrics
AS
SELECT
    toYYYYMM(Time) as Month,
    CategoryId,
    UserId,
    count() as TransactionCount,
    sum(Amount) as TotalExpenses
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month,
    CategoryId,
    UserId;
