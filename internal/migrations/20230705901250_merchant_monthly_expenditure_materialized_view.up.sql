CREATE MATERIALIZED VIEW IF NOT EXISTS MerchantMonthlyExpenditureMV
TO MerchantMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    MerchantName,
    sum(Amount) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount < 0
GROUP BY
    Month, MerchantName;