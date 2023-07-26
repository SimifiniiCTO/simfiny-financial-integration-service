-- Monthly Transaction Count: This view provides the total number of transactions per month.
CREATE TABLE IF NOT EXISTS MonthlyTransactionCount 
(
    Month UInt32,
    UserId UInt64,
    TransactionCount UInt64
) ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, UserId);