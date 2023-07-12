-- Debt-to-Income Ratio: This view calculates the ratio of expenditure to income for each user on a monthly basis,
-- which is often used to assess a person's financial health
CREATE TABLE IF NOT EXISTS DebtToIncomeRatio 
(
    Month UInt32,
    UserId UInt64,
    Ratio Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, UserId);