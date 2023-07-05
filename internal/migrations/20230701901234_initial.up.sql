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





-- -- Category-wise Monthly Expenditure: This view provides the total spending per category each month.
-- CREATE TABLE IF NOT EXISTS CategoryMonthlyExpenditure 
-- (
--     Month UInt32,
--     CategoryId String,
--     TotalSpending AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, CategoryId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyExpenditureMV
-- TO CategoryMonthlyExpenditure
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     CategoryId,
--     sum(Amount) as TotalSpending
-- FROM
--     TransactionInternal
-- WHERE
--     Amount < 0
-- GROUP BY
--     Month, CategoryId;

-- -- Category-wise Monthly Income: This view provides the total income per category each month.
-- CREATE TABLE IF NOT EXISTS CategoryMonthlyIncome 
-- (
--     Month UInt32,
--     CategoryId String,
--     TotalIncome AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, CategoryId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyIncomeMV
-- TO CategoryMonthlyIncome
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     CategoryId,
--     sum(Amount) as TotalIncome
-- FROM
--     TransactionInternal
-- WHERE
--     Amount > 0
-- GROUP BY
--     Month, CategoryId;


-- -- Merchant-wise Monthly Expenditure: This view provides the total spending per merchant each month.
-- CREATE TABLE IF NOT EXISTS MerchantMonthlyExpenditure 
-- (
--     Month UInt32,
--     MerchantName String,
--     TotalSpending AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, MerchantName);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MerchantMonthlyExpenditureMV
-- TO MerchantMonthlyExpenditure
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     MerchantName,
--     sum(Amount) as TotalSpending
-- FROM
--     TransactionInternal
-- WHERE
--     Amount < 0
-- GROUP BY
--     Month, MerchantName;

-- -- Monthly Income: This view provides the total income per month.
-- CREATE TABLE IF NOT EXISTS MonthlyIncome 
-- (
--     Month UInt32,
--     TotalIncome AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY Month;

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyIncomeMV
-- TO MonthlyIncome
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     sum(Amount) as TotalIncome
-- FROM
--     TransactionInternal
-- WHERE
--     Amount > 0
-- GROUP BY
--     Month;

-- -- Monthly Balance: This view provides the net balance per month (income - expenditure).
-- CREATE TABLE IF NOT EXISTS MonthlyBalance 
-- (
--     Month UInt32,
--     NetBalance AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY Month;

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyBalanceMV
-- TO MonthlyBalance
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     sum(Amount) as NetBalance
-- FROM
--     TransactionInternal
-- GROUP BY
--     Month;

-- -- Monthly Transaction Count: This view provides the total number of transactions per month.
-- CREATE TABLE IF NOT EXISTS MonthlyTransactionCount 
-- (
--     Month UInt32,
--     TransactionCount AggregateFunction(count)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY Month;

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyTransactionCountMV
-- TO MonthlyTransactionCount
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     count() as TransactionCount
-- FROM
--     TransactionInternal
-- GROUP BY
--     Month;

-- -- PaymentChannel-wise Monthly Expenditure: This view shows how much a user spends 
-- -- through different payment channels each month.
-- CREATE TABLE IF NOT EXISTS PaymentChannelMonthlyExpenditure 
-- (
--     Month UInt32,
--     PaymentChannel String,
--     TotalSpending AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, PaymentChannel);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS PaymentChannelMonthlyExpenditureMV
-- TO PaymentChannelMonthlyExpenditure
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     PaymentChannel,
--     sum(Amount) as TotalSpending
-- FROM
--     TransactionInternal
-- WHERE
--     Amount < 0
-- GROUP BY
--     Month, PaymentChannel;

-- -- Income vs Expenditure Ratio: This view provides the ratio of total income to total expenditure per month.
-- CREATE TABLE IF NOT EXISTS IncomeExpenseRatio 
-- (
--     Month UInt32,
--     Ratio AggregateFunction(avg, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY Month;

-- CREATE MATERIALIZED VIEW IF NOT EXISTS IncomeExpenseRatioMV
-- TO IncomeExpenseRatio
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     sum(if(Amount > 0, Amount, 0)) / sum(if(Amount < 0, -Amount, 0)) as Ratio
-- FROM
--     TransactionInternal
-- GROUP BY
--     Month;

-- -- Monthly Number of Transactions by Category: This view provides the number
-- -- of transactions per category each month.
-- CREATE TABLE IF NOT EXISTS CategoryMonthlyTransactionCount 
-- (
--     Month UInt32,
--     CategoryId String,
--     TransactionCount AggregateFunction(count)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, CategoryId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS CategoryMonthlyTransactionCountMV
-- TO CategoryMonthlyTransactionCount
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     CategoryId,
--     count() as TransactionCount
-- FROM
--     TransactionInternal
-- GROUP


-- -- Net Savings per Month: This view provides the net savings (income - expenditure) per month.
-- CREATE TABLE IF NOT EXISTS MonthlySavings 
-- (
--     Month UInt32,
--     UserId UInt64,
--     NetSavings AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, UserId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlySavingsMV
-- TO MonthlySavings
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     UserId,
--     sum(Amount) as NetSavings
-- FROM
--     TransactionInternal
-- GROUP BY
--     Month, UserId;


-- -- Debt-to-Income Ratio: This view calculates the ratio of expenditure to income for each user on a monthly basis,
-- -- which is often used to assess a person's financial health
-- CREATE TABLE IF NOT EXISTS DebtToIncomeRatio 
-- (
--     Month UInt32,
--     UserId UInt64,
--     Ratio AggregateFunction(avg, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, UserId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS DebtToIncomeRatioMV
-- TO DebtToIncomeRatio
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     UserId,
--     sum(if(Amount < 0, -Amount, 0)) / sum(if(Amount > 0, Amount, 0)) as Ratio
-- FROM
--     TransactionInternal
-- GROUP BY
--     Month, UserId;


-- -- Monthly Income Growth Rate: This view calculates the month-over-month growth rate of a user's income, which can indicate whether the user's income 
-- -- is increasing or decreasing over time.
-- -- Note: The LAG function is not supported in ClickHouse as of the knowledge cutoff in September 2021. 
-- -- You may need to calculate the growth rate outside
-- --  of ClickHouse, for example, in your application code or in a BI tool.
-- CREATE TABLE IF NOT EXISTS MonthlyIncomeGrowthRate 
-- (
--     Month UInt32,
--     UserId UInt64,
--     GrowthRate AggregateFunction(avg, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, UserId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyIncomeGrowthRateMV
-- TO MonthlyIncomeGrowthRate
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     UserId,
--     (sum(Amount) - lag(sum(Amount), 1, 0) over (partition by UserId order by Month)) / lag(sum(Amount), 1, 0) as GrowthRate
-- FROM
--     TransactionInternal
-- WHERE
--     Amount > 0
-- GROUP BY
--     Month, UserId;

-- -- Monthly Expense Category Distribution: This view provides the distribution of expenses across different 
-- -- categories per month, which can reveal a user's spending habits.
-- CREATE TABLE IF NOT EXISTS MonthlyExpenseCategoryDistribution 
-- (
--     Month UInt32,
--     UserId UInt64,
--     CategoryId String,
--     TotalSpending AggregateFunction(sum, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, UserId, CategoryId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyExpenseCategoryDistributionMV
-- TO MonthlyExpenseCategoryDistribution
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     UserId,
--     CategoryId,
--     sum(Amount) as TotalSpending
-- FROM
--     TransactionInternal
-- WHERE
--     Amount < 0
-- GROUP BY
--     Month, UserId, CategoryId;

-- -- In this query, we use a subquery to find the category with the maximum expenditure for each user.
-- -- This information could be valuable for identifying areas where budgeting could help.
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

-- --  Monthly Average Transaction Value by User
-- CREATE TABLE IF NOT EXISTS MonthlyAvgTransactionValueByUser 
-- (
--     Month UInt32,
--     UserId UInt64,
--     AvgTransactionValue AggregateFunction(avg, Float64)
-- ) 
-- ENGINE = AggregatingMergeTree()
-- PARTITION BY Month
-- ORDER BY (Month, UserId);

-- CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyAvgTransactionValueByUserMV
-- TO MonthlyAvgTransactionValueByUser
-- AS
-- SELECT
--     toYYYYMM(Time) as Month,
--     UserId,
--     avg(Amount) as AvgTransactionValue
-- FROM TransactionInternal
-- GROUP BY Month, UserId;
