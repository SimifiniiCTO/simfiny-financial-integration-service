CREATE MATERIALIZED VIEW IF NOT EXISTS UserFinancialHealthMetrics
ENGINE = AggregatingMergeTree()
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, UserId)
AS
SELECT
    Time,
    UserId,
    sumIf(Amount, Sign = 1) AS MonthlyIncome,
    sumIf(Amount, Sign = -1) AS MonthlyExpense,
    countDistinctIf(CategoryId, Sign = -1) AS TransactionDiversity,
    if(sumIf(Amount, Sign = -1) = 0, 0, sumIf(Amount, Sign = 1) / abs(sumIf(Amount, Sign = -1))) AS DebtToIncomeRatio,
    countIf(Amount < 0) AS OverdraftFrequency
FROM TransactionInternal
GROUP BY Time, UserId