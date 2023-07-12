-- MonthlyIncome captures the sum of all incoming transactions per month per user.
-- MonthlyExpense captures the sum of all outgoing transactions (expenses) per month per user.
-- TransactionDiversity counts the distinct categories of the user's expenses, providing a measure of how diversified their expenses are.
-- DebtToIncomeRatio calculates the ratio of income to debt. Note that it uses an absolute value for the denominator to prevent division by a negative number (assuming that debt is represented by negative values).
-- OverdraftFrequency counts the number of times a user's account went negative, indicating an overdraft. It's assuming that the Amount is negative in such cases.
-- These metrics are still calculated on a monthly basis. The time partitioning and ordering in the AggregatingMergeTree will allow efficient querying for these metrics over time periods.
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