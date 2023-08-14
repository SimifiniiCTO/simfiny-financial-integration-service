CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyTransactionCountMV
TO CategoryMonthlyTransactionCount
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    sum(Sign) as TransactionCount
FROM
    TransactionInternal
WHERE
    Sign = 1
GROUP BY
    Month, PersonalFinanceCategoryPrimary, UserId;
