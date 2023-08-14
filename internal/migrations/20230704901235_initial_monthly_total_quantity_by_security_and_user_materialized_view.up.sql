CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyTotalQuantityBySecurityAndUserMV
TO MonthlyTotalQuantityBySecurityAndUser
AS
SELECT
    toYYYYMM(Time) as Month,
    UserId,
    SecurityId,
    sum(Quantity * Sign) as TotalQuantity
FROM InvestmentTransactionInternal
WHERE
    Sign = 1
GROUP BY
    Month, UserId, SecurityId;