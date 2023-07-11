CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyTotalQuantityBySecurityAndUserMV
TO MonthlyTotalQuantityBySecurityAndUser
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    SecurityId,
    sum(Quantity) as TotalQuantity
FROM InvestmentTransactionInternal
GROUP BY Month, UserId, SecurityId;