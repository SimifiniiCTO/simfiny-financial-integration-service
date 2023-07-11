-- Monthly Transaction Count: This view provides the total number of transactions per month.
CREATE TABLE IF NOT EXISTS MonthlyTransactionCount 
(
    Month UInt32,
    TransactionCount AggregateFunction(count)
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;