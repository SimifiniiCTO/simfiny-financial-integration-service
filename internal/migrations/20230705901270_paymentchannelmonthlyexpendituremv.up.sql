CREATE MATERIALIZED VIEW IF NOT EXISTS PaymentChannelMonthlyExpenditureMV
TO PaymentChannelMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    PaymentChannel,
    sum(Amount) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount > 0
GROUP BY
    Month, PaymentChannel;