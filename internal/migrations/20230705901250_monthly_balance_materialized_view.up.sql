CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyBalanceMV
TO MonthlyBalance
AS
SELECT
    toYYYYMM(Time) as Month,
    sum(Amount) as NetBalance
FROM
    TransactionInternal
GROUP BY
    Month;