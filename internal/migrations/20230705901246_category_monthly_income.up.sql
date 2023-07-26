-- Category-wise Monthly Income: This view provides the total income per category each month.
CREATE TABLE IF NOT EXISTS CategoryMonthlyIncome 
(
    Month UInt32,
    PersonalFinanceCategoryPrimary String,
    TotalIncome Float64,
    UserId UInt64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, PersonalFinanceCategoryPrimary);