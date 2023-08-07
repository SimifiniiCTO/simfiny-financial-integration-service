CREATE MATERIALIZED VIEW IF NOT EXISTS DebtToIncomeRatioMV
TO DebtToIncomeRatio
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sum(if(Amount > 0, -Amount, 0)) / sum(if(Amount < 0, Amount, 0)) as Ratio
FROM
    TransactionInternal
GROUP BY
    Month, UserId;