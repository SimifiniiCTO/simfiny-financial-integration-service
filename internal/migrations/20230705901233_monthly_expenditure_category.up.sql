-- Category-wise Monthly Expenditure: This view provides the total spending per category each month.
CREATE TABLE IF NOT EXISTS CategoryMonthlyExpenditure 
(
    Month UInt32,
    CategoryId String,
    TotalSpending Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, CategoryId);