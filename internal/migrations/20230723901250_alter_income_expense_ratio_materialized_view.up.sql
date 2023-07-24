-- Step 1: Drop the existing materialized view
DROP MATERIALIZED VIEW IF EXISTS IncomeExpenseRatioMV;

-- Step 2: Create the materialized view with updated column
-- INCOME according to plaid is when the ammount is actually negative. 
-- 'Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.'
-- ref: https://plaid.com/docs/api/products/transactions/
CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeExpenseRatioMV
TO IncomeExpenseRatio
AS
SELECT
    toYYYYMM(Time) as Month,
    (-sum(if(Amount < 0, Amount, 0))) / sum(if(Amount > 0, Amount, 0)) as IncomeExpenseRatio
FROM
    TransactionInternal
GROUP BY
    Month;