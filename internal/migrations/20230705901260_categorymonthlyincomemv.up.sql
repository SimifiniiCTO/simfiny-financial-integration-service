-- INCOME according to plaid is when the ammount is actually negative. 
-- 'Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.'
-- ref: https://plaid.com/docs/api/products/transactions/
CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyIncomeMV
TO CategoryMonthlyIncome
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    -sum(Amount * Sign) as TotalIncome  -- Make income positive for reporting
FROM
    TransactionInternal
WHERE
    Amount < 0 AND Sign = 1
GROUP BY
    Month, PersonalFinanceCategoryPrimary, UserId;
