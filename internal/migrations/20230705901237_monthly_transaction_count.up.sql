-- Monthly Transaction Count: This view provides the total number of transactions per month.
CREATE TABLE IF NOT EXISTS MonthlyTransactionCount 
(
    Month Date,
    UserId UInt64,
    TransactionCount AggregateFunction(count)
) ENGINE = AggregatingMergeTree()
PARTITION BY toYYYYMM(Month)
ORDER BY (Month, UserId);