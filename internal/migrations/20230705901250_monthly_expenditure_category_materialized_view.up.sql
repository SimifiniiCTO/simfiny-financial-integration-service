CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyExpenditureMV
TO CategoryMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    CategoryId,
    sum(Amount) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month, CategoryId;