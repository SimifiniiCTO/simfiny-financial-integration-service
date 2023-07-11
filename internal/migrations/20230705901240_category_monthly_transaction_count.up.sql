-- Monthly Number of Transactions by Category: This view provides the number
-- of transactions per category each month.
CREATE TABLE IF NOT EXISTS CategoryMonthlyTransactionCount 
(
    Month UInt32,
    CategoryId String,
    TransactionCount AggregateFunction(count)
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, CategoryId);