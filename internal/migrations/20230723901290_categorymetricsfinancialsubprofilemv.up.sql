-- 1. Backup is an operational task, so it's not represented as an SQL command here.
-- Ensure you've backed up your data!

-- 2. Create a temporary materialized view with the updated logic.
CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMetricsFinancialSubProfileMV_temp
TO CategoryMetricsFinancialSubProfile_temp
AS
SELECT
    toYYYYMM(Time) as Month,
    PersonalFinanceCategoryPrimary,
    UserId,
    sum(Sign) as TransactionCount,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 WEEK) as SpentLastWeek,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 2 WEEK) as SpentLastTwoWeeks,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 MONTH) as SpentLastMonth,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 6 MONTH) as SpentLastSixMonths,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 1 YEAR) as SpentLastYear,
    sumIf(Amount * Sign, Time >= now() - INTERVAL 2 YEAR) as SpentLastTwoYears
FROM
    TransactionInternal
WHERE
    Amount > 0 AND Sign = 1
GROUP BY
    Month, PersonalFinanceCategoryPrimary, UserId;

-- Populate the view (if necessary). Depending on how your DB and materialized views are set up, you might need to insert or refresh the data.

-- 3. Once you have tested the data in the temporary materialized view and are satisfied, you can rename the old view and the new one.
RENAME TABLE CategoryMetricsFinancialSubProfileMV TO CategoryMetricsFinancialSubProfileMV_backup;
RENAME TABLE CategoryMetricsFinancialSubProfileMV_temp TO CategoryMetricsFinancialSubProfileMV;

-- Additional steps like repointing to the original table and cleanup are operational and may vary based on your specific setup and needs. They aren't represented here.
