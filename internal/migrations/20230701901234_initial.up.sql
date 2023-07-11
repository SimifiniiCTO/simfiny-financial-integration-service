CREATE TABLE IF NOT EXISTS TransactionInternal (
    AccountId                       LowCardinality(String),
    AccountOwner                    LowCardinality(String),
    Amount                          Float64,
    AuthorizedDate                  LowCardinality(String),
    AuthorizedDatetime              LowCardinality(String),
    CategoryId                      LowCardinality(String),
    CheckNumber                     LowCardinality(String),
    CurrentDate                     LowCardinality(String),
    CurrentDatetime                 LowCardinality(String),
    ID                              String DEFAULT generateUUIDv4(),
    IsoCurrencyCode                 LowCardinality(String),
    LinkId                          UInt64,
    LocationAddress                 LowCardinality(String),
    LocationCity                    LowCardinality(String),
    LocationCountry                 LowCardinality(String),
    LocationLat                     Float64,
    LocationLon                     Float64,
    LocationPostalCode              LowCardinality(String),
    LocationRegion                  LowCardinality(String),
    LocationStoreNumber             LowCardinality(String),
    MerchantName                    LowCardinality(String),
    Name                            LowCardinality(String),
    PaymentChannel                  LowCardinality(String),
    PaymentMetaByOrderOf            LowCardinality(String),
    PaymentMetaPayee                LowCardinality(String),
    PaymentMetaPayer                LowCardinality(String),
    PaymentMetaPaymentMethod        LowCardinality(String),
    PaymentMetaPaymentProcessor     LowCardinality(String),
    PaymentMetaPpdId                LowCardinality(String),
    PaymentMetaReason               LowCardinality(String),
    PaymentMetaReferenceNumber      LowCardinality(String),
    Pending                         Bool,
    PendingTransactionId            LowCardinality(String),
    PersonalFinanceCategoryDetailed LowCardinality(String),
    PersonalFinanceCategoryPrimary  LowCardinality(String),
    Sign                            Int8,
    Time                            DateTime,
    TransactionCode                 LowCardinality(String),
    TransactionId                   LowCardinality(String),
    UnofficialCurrencyCode          LowCardinality(String),
    UserId                          UInt64,
    Categories                      Array(String)
) ENGINE = CollapsingMergeTree(Sign)
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, ID)
PRIMARY KEY (Time, ID);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS FinancialProfileMV
-- TO FinancialProfile
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     UserId,
--     sumIf(Amount, Amount > 0) as TotalIncome,
--     -sumIf(Amount, Amount < 0) as TotalExpenses,
--     count() as NumberOfTransactions,
--     (
--         SELECT CategoryId FROM 
--         (
--             SELECT 
--                 CategoryId, 
--                 -sum(Amount) as CategoryExpense 
--             FROM TransactionInternal 
--             WHERE Amount < 0 AND UserId = u.UserId AND toYYYYMM(Time) = Month
--             GROUP BY CategoryId 
--             ORDER BY CategoryExpense DESC 
--             LIMIT 1
--         )
--     ) as MostExpensiveCategory
-- FROM TransactionInternal as u
-- GROUP BY Month, UserId;

-- In this query, we use a subquery to find the category with the maximum expenditure for each user.
-- This information could be valuable for identifying areas where budgeting could help.
-- CREATE TABLE IF NOT EXISTS FinancialProfile 
-- (
--     Month UInt32,
--     UserId UInt64,
--     TotalIncome AggregateFunction(sum, Float64),
--     TotalExpenses AggregateFunction(sum, Float64),
--     NumberOfTransactions AggregateFunction(count),
--     MostExpensiveCategory String
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, UserId);