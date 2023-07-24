-- Create Materialized View for Expense Metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS ExpenseMetricsFinancialSubProfileMetricsMV
TO ExpenseMetricsFinancialSubProfileMetrics
AS
SELECT
    toYYYYMM(Time) as Month,
    sumIf(Amount, Time >= now() - INTERVAL 1 WEEK) as SpentLastWeek,
    sumIf(Amount, Time >= now() - INTERVAL 1 MONTH) as SpentLastMonth,
    sumIf(Amount, Time >= now() - INTERVAL 6 MONTH) as SpentLastSixMonths,
    avgIf(Amount, Time >= now() - INTERVAL 1 MONTH) as AverageMonthlyDiscretionarySpending,
    avgIf(Amount, Time >= now() - INTERVAL 1 MONTH AND PaymentMetaPaymentMethod = 'Recurring') as AverageMonthlyRecurringSpending
FROM
    TransactionInternal
WHERE
    Amount > 0
GROUP BY
    Month;