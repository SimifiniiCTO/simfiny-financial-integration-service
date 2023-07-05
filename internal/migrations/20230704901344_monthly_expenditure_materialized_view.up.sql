CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyExpenditureMV
TO MonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    sum(Amount) as TotalSpending,
    UserId UInt64
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month, UserId;