-- Monthly Total Quantity by Security and User
CREATE TABLE IF NOT EXISTS MonthlyTotalQuantityBySecurityAndUser 
(
    Month UInt32,
    UserId UInt64,
    SecurityId String,
    TotalQuantity AggregateFunction(sum, Float64)
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, UserId, SecurityId);
