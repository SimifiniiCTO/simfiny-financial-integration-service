-- table where expense metric information will be aggregated and stored
CREATE TABLE IF NOT EXISTS ExpenseMetrics 
(
    Month UInt32,
    PersonalFinanceCategoryPrimary String,
    TransactionCount UInt64,
    TotalExpenses Float32,
    UserId UInt64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, PersonalFinanceCategoryPrimary, UserId);