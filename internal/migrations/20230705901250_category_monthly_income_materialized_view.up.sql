CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyIncomeMV
TO CategoryMonthlyIncome
AS
SELECT
    toYYYYMM(Time) as Month,
    CategoryId,
    sum(Amount) as TotalIncome
FROM
    TransactionInternal
WHERE
    Amount > 0
GROUP BY
    Month, CategoryId;
