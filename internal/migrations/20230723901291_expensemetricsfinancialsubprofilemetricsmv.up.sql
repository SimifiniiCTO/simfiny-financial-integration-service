-- Create Materialized View for Expense Metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS ExpenseMetricsFinancialSubProfileMetricsMV
TO ExpenseMetricsFinancialSubProfileMetrics
AS
SELECT
    toYYYYMM(Time) as Month,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 WEEK) as SpentLastWeek,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 MONTH) as SpentLastMonth,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 6 MONTH) as SpentLastSixMonths,
    avgIf(Amount * Sign, Time >= now() - INTERVAL 1 MONTH) as AverageMonthlyDiscretionarySpending,
    avgIf(Amount * Sign, Time >= now() - INTERVAL 1 MONTH AND PaymentMetaPaymentMethod = 'Recurring') as AverageMonthlyRecurringSpending,
    UserId
FROM
    TransactionInternal
WHERE
    Amount < 0 AND Sign = 1
GROUP BY
    Month, UserId;