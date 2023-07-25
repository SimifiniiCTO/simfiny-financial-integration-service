CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyExpenditureMV
TO CategoryMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    sum(Amount) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month, PersonalFinanceCategoryPrimary;