-- Net Savings per Month: This view provides the net savings (income - expenditure) per month.
CREATE TABLE IF NOT EXISTS MonthlySavings 
(
    Month UInt32,
    UserId UInt64,
    NetSavings Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, UserId);