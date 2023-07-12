CREATE TABLE IF NOT EXISTS UserFinancialHealthMetricsTable (
    Time                            DateTime,
    UserId                          UInt64,
    MonthlyIncome                   Float64,
    MonthlyExpense                  Float64,
    TransactionDiversity            UInt64,
    DebtToIncomeRatio               Float64,
    OverdraftFrequency              UInt64
) ENGINE = MergeTree()
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, UserId);