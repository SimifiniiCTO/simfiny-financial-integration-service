CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyIncomeMV
TO MonthlyIncome
AS
SELECT
    UserId,
    toYYYYMM(Time) as Month,
    sum(Amount) as TotalIncome
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month, 
    UserId;
