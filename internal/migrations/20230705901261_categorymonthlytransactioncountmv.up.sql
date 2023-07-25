CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyTransactionCountMV
TO CategoryMonthlyTransactionCount
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    count() as TransactionCount
FROM
    TransactionInternal
GROUP BY
    Month,
    PersonalFinanceCategoryPrimary;