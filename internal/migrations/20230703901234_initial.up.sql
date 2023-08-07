CREATE TABLE IF NOT EXISTS AccountBalanceHistoryInternal(
    Time             DateTime,
    AccountId        LowCardinality(String),
    IsoCurrencyCode  LowCardinality(String),
    Balance          Float64,
    UserId           UInt64,
    Sign             Int8,
    Id               String DEFAULT generateUUIDv4(),
) ENGINE = CollapsingMergeTree(Sign)
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, Id)