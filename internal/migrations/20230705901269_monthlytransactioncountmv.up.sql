CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyTransactionCountMV
TO MonthlyTransactionCount
AS
SELECT
    toYYYYMM(Time) AS Month,
    UserId,
    sum(Sign) AS TransactionCount
FROM
    TransactionInternal
WHERE
    Sign = 1
GROUP BY
    Month,
    UserId;