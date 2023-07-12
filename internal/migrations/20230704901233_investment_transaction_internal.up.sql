CREATE TABLE IF NOT EXISTS InvestmentTransactionInternal (
    AccountId               String,
    Amount                  Float64,
    CreatedAt               String,
    CurrentDate             String,
    Fees                    Float64,
    ID                      String DEFAULT generateUUIDv4(),
    InvestmentTransactionId String,
    IsoCurrencyCode         String,
    LinkId                  UInt64,
    Name                    String,
    Price                   Float64,
    Quantity                Float64,
    SecurityId              String,
    Sign                    Int8,
    Subtype                 String,
    Time                    DateTime,
    Type                    String,
    UnofficialCurrencyCode  String,
    UserId                  UInt64
) ENGINE = CollapsingMergeTree(Sign)
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, UserId, ID)
PRIMARY KEY (Time, UserId, ID);
