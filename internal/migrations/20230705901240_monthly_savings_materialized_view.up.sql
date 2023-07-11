CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlySavingsMV
TO MonthlySavings
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sum(Amount) as NetSavings
FROM
    TransactionInternal
GROUP BY
    Month, UserId;
