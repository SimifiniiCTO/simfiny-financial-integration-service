-- Income vs Expenditure Ratio: This view provides the ratio of total income to total expenditure per month.
CREATE TABLE IF NOT EXISTS IncomeExpenseRatio 
(
    Month UInt32,
    Ratio Float64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY Month;
