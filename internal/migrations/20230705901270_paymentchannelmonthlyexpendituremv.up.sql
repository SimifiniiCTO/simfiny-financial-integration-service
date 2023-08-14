CREATE MATERIALIZED VIEW IF NOT EXISTS PaymentChannelMonthlyExpenditureMV
TO PaymentChannelMonthlyExpenditure
AS
SELECT
    toYYYYMM(Time) as Month,
    PaymentChannel,
    UserId,
    sum(Sign * Amount) as TotalSpending
FROM
    TransactionInternal
WHERE
    Amount > 0 AND Sign = 1
GROUP BY
    Month, PaymentChannel, UserId;