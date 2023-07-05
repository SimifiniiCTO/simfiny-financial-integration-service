CREATE TABLE IF NOT EXISTS ReOccuringTransactionInternal (
    AccountId                       String,
    AverageAmount                   String,
    AverageAmountIsoCurrencyCode    String,
    CategoryId                      String,
    Description                     String,
    FirstDate                       String,
    Flow                            String,
    Frequency                       String,
    Id                              String DEFAULT generateUUIDv4(),
    IsActive                        UInt8,
    LastAmount                      String,
    LastAmountIsoCurrencyCode       String,
    LastDate                        String,
    LinkId                          UInt64,
    MerchantName                    String,
    PersonalFinanceCategoryDetailed String,
    PersonalFinanceCategoryPrimary  String,
    Sign                            Int8,
    Status                          String,
    StreamId                        String,
    Time                            DateTime,
    TransactionIds                  String,
    UpdatedTime                     String,
    UserId                          UInt64
) ENGINE = CollapsingMergeTree(Sign)
PARTITION BY toYYYYMM(Time)
ORDER BY (Time, Id)
PRIMARY KEY (Time, Id);