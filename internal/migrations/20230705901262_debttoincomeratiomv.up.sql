CREATE MATERIALIZED VIEW IF NOT EXISTS DebtToIncomeRatioMV
TO DebtToIncomeRatio
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    sum(if(Amount > 0, -Amount, 0) * Sign) / sum(if(Amount < 0, Amount, 0) * Sign) as Ratio
FROM
    TransactionInternal
WHERE
    Sign = 1
GROUP BY
    Month, UserId;