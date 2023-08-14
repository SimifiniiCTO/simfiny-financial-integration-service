CREATE MATERIALIZED VIEW IF NOT EXISTS TotalInvestmentBySecurityMV
TO TotalInvestmentBySecurity
AS
SELECT
    UserId,
    SecurityId,
    sum(Amount * Sign) as TotalInvestment
FROM InvestmentTransactionInternal
WHERE
    Sign = 1
GROUP BY
    SecurityId, UserId;