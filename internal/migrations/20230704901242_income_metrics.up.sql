-- target table where income metrics will be stored
CREATE TABLE IF NOT EXISTS IncomeMetrics 
(
    Month UInt32,
    CategoryId String,
    TransactionCount UInt64,
    TotalIncome Float32,
    UserId UInt64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, CategoryId, UserId);