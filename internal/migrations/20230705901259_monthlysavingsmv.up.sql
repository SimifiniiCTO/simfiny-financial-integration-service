CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlySavingsMV
TO MonthlySavings
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sum(Amount * Sign) as NetSavings
FROM
    TransactionInternal
WHERE
    Sign = 1
GROUP BY
    Month, UserId;
