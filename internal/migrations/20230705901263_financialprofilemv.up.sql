CREATE MATERIALIZED VIEW IF NOT EXISTS FinancialProfileMV
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sumIf(Amount, Amount > 0) as TotalIncome,
    -sumIf(Amount, Amount < 0) as TotalExpenses,
    count() as NumberOfTransactions
FROM TransactionInternal
GROUP BY Month, UserId;