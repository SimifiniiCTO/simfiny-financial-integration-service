CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyTransactionCountMV
TO MonthlyTransactionCount
AS
SELECT
    toYYYYMM(Time) AS Month,
    UserId,
    count() AS TransactionCount
FROM
    TransactionInternal
GROUP BY
    Month,
    UserId;