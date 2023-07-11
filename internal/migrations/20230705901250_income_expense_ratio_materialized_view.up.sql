CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeExpenseRatioMV
TO IncomeExpenseRatio
AS
SELECT
    toYYYYMM(Time) as Month,
    sum(if(Amount > 0, Amount, 0)) / sum(if(Amount < 0, -Amount, 0)) as Ratio
FROM
    TransactionInternal
GROUP BY
    Month;