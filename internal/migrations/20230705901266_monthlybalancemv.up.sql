CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyBalanceMV
TO MonthlyBalance
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sum(Amount * Sign) as NetBalance
FROM
    TransactionInternal
WHERE
    Sign = 1
GROUP BY
    Month, UserId;
