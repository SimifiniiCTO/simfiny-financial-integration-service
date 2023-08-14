CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyExpenditureMV
TO CategoryMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    sum(Amount * Sign) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount > 0 AND Sign = 1
GROUP BY
    Month, PersonalFinanceCategoryPrimary, UserId;
