-- Monthly Number of Transactions by Category: This view provides the number
-- of transactions per category each month.
CREATE TABLE IF NOT EXISTS CategoryMonthlyTransactionCount 
(
    Month UInt32,
    PersonalFinanceCategoryPrimary String,
    TransactionCount UInt32,
    UserId UInt64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, PersonalFinanceCategoryPrimary);