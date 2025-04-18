CREATE TABLE IF NOT EXISTS MonthlyExpenditure 
(
    Month UInt32,
    UserId UInt64,
    TotalSpending Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;