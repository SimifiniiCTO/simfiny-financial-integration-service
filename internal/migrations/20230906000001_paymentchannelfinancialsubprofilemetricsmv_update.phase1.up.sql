-- To perform migrations of materialized views, we first have to drop the materialized view and then recreate it.
-- Because we cannot perform both operations in one go in one file, we opt to split the migration into two files.
DROP TABLE IF EXISTS PaymentChannelFinancialSubProfileMetricsMV;
