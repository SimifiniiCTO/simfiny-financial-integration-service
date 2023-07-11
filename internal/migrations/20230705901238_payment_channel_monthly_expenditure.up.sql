
-- PaymentChannel-wise Monthly Expenditure: This view shows how much a user spends 
-- through different payment channels each month.
CREATE TABLE IF NOT EXISTS PaymentChannelMonthlyExpenditure 
(
    Month UInt32,
    PaymentChannel String,
    TotalSpending Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, PaymentChannel);