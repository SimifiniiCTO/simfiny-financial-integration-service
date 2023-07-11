-- Category-wise Monthly Income: This view provides the total income per category each month.
CREATE TABLE IF NOT EXISTS CategoryMonthlyIncome 
(
    Month UInt32,
    CategoryId String,
    TotalIncome Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, CategoryId);