CREATE TABLE IF NOT EXISTS AccountBalanceHistoryInternal(
    Time             DateTime,
    AccountId        String,
    IsoCurrencyCode String,
    Balance           Float64,
    UserId           UInt64,
    Sign              Int8,
    Id               String
) ENGINE = CollapsingMergeTree(Sign)
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, Id)