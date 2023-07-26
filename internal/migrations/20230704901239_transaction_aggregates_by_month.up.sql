-- target table for the below materialized view
CREATE TABLE IF NOT EXISTS TransactionAggregatesByMonth 
(
    Month UInt32,
    PersonalFinanceCategoryPrimary String,
    LocationCity String,
    PaymentChannel String,
    MerchantName String,
    TransactionCount UInt64,
    TotalAmount Float64,
    UserId UInt64
) 
ENGINE = AggregatingMergeTree()
PARTITION BY Month
ORDER BY (Month, PersonalFinanceCategoryPrimary, LocationCity, PaymentChannel, MerchantName);
