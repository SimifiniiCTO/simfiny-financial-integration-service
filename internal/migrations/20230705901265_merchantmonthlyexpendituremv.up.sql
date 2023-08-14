CREATE MATERIALIZED VIEW IF NOT EXISTS MerchantMonthlyExpenditureMV
TO MerchantMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    MerchantName,
    UserId,
    sum(Amount * Sign) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount > 0 AND Sign = 1
GROUP BY
    Month, MerchantName, UserId;
