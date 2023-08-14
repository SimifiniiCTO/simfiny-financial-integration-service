CREATE MATERIALIZED VIEW IF NOT EXISTS FinancialProfileMV
TO FinancialProfile
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sumIf(Amount * Sign, Amount > 0) as TotalIncome,
    -sumIf(Amount * Sign, Amount < 0) as TotalExpenses,
    sum(Sign) as NumberOfTransactions
FROM TransactionInternal
WHERE
    Sign = 1
GROUP BY Month, UserId;
