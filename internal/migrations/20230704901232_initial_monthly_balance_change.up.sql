CREATE TABLE IF NOT EXISTS MonthlyBalanceChangeByUser 
(
    Month UInt32,
    UserId UInt64,
    BalanceChange AggregateFunction(sum, Float64)
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, UserId);
