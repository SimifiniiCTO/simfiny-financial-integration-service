CREATE MATERIALIZED VIEW IF NOT EXISTS TotalInvestmentBySecurityMV
TO TotalInvestmentBySecurity
AS
SELECT
    UserId,
    SecurityId,
    sum(Amount) as TotalInvestment
FROM InvestmentTransactionInternal
GROUP BY SecurityId, UserId;