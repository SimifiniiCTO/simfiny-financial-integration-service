-- In this schema, the countState() and sumState() functions are used instead of count() and sum().
-- These functions return the intermediate state of the count and sum calculations, which is compatible with
-- the AggregateFunction(count) and AggregateFunction(sum, Float32) data types in the TransactionAggregatesByMonth table.
--
-- Remember, when you query the TransactionAggregatesByMonth table, you'll need to
-- use the -Merge function combinator to finalize the calculation from the aggregate state. For example:
--
-- SELECT 
--     Month,
--     PersonalFinanceCategoryPrimary,
--     LocationCity,
--     PaymentChannel,
--     MerchantName,
--     countMerge(TransactionCount) as TransactionCount,
--     sumMerge(TotalAmount) as TotalAmount
-- FROM
--     TransactionAggregatesByMonth
-- GROUP BY
--     Month,
--     PersonalFinanceCategoryPrimary,
--     LocationCity,
--     PaymentChannel,
--     MerchantName;

CREATE MATERIALIZED VIEW IF NOT EXISTS TransactionAggregatesByMonthMV
TO TransactionAggregatesByMonth
AS
SELECT
    toYYYYMM(Time) AS Month,
    PersonalFinanceCategoryPrimary,
    LocationCity,
    PaymentChannel,
    MerchantName,
    UserId,
    count() AS TransactionCount,
    sum(Amount) AS TotalAmount
FROM
    TransactionInternal
GROUP BY
    Month,
    PersonalFinanceCategoryPrimary,
    LocationCity,
    PaymentChannel,
    UserId,
    MerchantName;