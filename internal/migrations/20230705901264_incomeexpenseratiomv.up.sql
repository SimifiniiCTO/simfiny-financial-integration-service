CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeExpenseRatioMV
TO IncomeExpenseRatio
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    (-sum(if(Amount < 0, Amount, 0))) / sum(if(Amount > 0, Amount, 0)) as IncomeExpenseRatio
FROM
    TransactionInternal
GROUP BY
    Month, UserId;