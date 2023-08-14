CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyIncomeMV
TO MonthlyIncome
AS
SELECT
    UserId,
    toYYYYMM(Time) as Month,
    -sum(Amount * Sign) as TotalIncome  -- Convert the negative income values into positive
FROM
    TransactionInternal
WHERE
    Amount < 0 AND Sign = 1
GROUP BY
    Month, 
    UserId;
