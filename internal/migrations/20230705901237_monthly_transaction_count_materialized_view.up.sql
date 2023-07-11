CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyTransactionCountMV
TO MonthlyTransactionCount
AS
SELECT
    toYYYYMM(Time) as Month,
    count() as TransactionCount
FROM
    TransactionInternal
GROUP BY
    Month;