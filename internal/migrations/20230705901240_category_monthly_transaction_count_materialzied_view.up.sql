CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyTransactionCountMV
TO CategoryMonthlyTransactionCount
AS
SELECT
    toYYYYMM(Time) as Month,
    CategoryId,
    count() as TransactionCount
FROM
    TransactionInternal
GROUP BY
    Month;