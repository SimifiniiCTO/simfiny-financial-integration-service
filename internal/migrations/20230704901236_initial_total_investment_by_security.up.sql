-- Total investmnet by security
CREATE TABLE IF NOT EXISTS TotalInvestmentBySecurity 
(
    UserId UInt64,
    SecurityId String,
    TotalInvestment Float64
) 
ENGINE = AggregatingMergeTree()
ORDER BY SecurityId;