CREATE MATERIALIZED VIEW IF NOT EXISTS MonthlyBalanceChangeByUserMV
TO MonthlyBalanceChangeByUser
AS
SELECT
    toYYYYMM(time) as Month,
    UserId,
    (maxIf(Balance, time == lastDayOfMonth(time)) - minIf(Balance, time == toDate(toStartOfMonth(time)))) as BalanceChange
FROM AccountBalanceHistoryInternal
GROUP BY Month, UserId;
