-- Monthly Balance: This view provides the net balance per month (income - expenditure).
CREATE TABLE IF NOT EXISTS MonthlyBalance 
(
    Month UInt32,
    NetBalance Float64,
    UserId UInt64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;
