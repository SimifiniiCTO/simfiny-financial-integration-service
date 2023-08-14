-- materialized view that will transform transactions into income metrics
CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeMetricsMV
TO IncomeMetrics
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    -- ref: https://clickhouse.com/docs/en/engines/table-engines/mergetree-family/collapsingmergetree
    -- we calculate the transaction count differently for collapsing merge tree engine tables
    sum(Sign) as TransactionCount,
    sum(Sign * Amount) as TotalIncome
FROM
    TransactionInternal
WHERE
    Amount > 0
GROUP BY
    Month,
    PersonalFinanceCategoryPrimary,
    UserId;