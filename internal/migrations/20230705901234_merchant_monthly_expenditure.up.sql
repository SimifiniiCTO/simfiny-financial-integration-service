-- Merchant-wise Monthly Expenditure: This view provides the total spending per merchant each month.
CREATE TABLE IF NOT EXISTS MerchantMonthlyExpenditure 
(
    Month UInt32,
    MerchantName String,
    TotalSpending Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, MerchantName);
