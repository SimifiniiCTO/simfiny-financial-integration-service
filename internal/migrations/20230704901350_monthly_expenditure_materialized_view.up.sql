CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyExpenditureMV
TO MonthlyExpenditure
AS
SELECT
    UserId,
    toYYYYMM(Time) as Month,
    sum(Amount) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month, 
    UserId;