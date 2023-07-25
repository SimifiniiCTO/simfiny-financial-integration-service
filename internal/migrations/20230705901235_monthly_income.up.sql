-- Monthly Income: This view provides the total income per month.
CREATE TABLE IF NOT EXISTS MonthlyIncome 
(
    Month UInt32,
    UserId UInt64,
    TotalIncome Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;
