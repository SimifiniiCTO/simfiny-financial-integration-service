CREATE TABLE IF NOT EXISTS FinancialProfile 
(
    Month UInt32,
    UserId UInt64,
    TotalIncome Float64,
    TotalExpenses Float64,
    NumberOfTransactions UInt64,
    MostExpensiveCategory String
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, UserId);