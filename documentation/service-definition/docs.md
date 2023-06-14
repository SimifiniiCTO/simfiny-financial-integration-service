# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/clickhouse.proto](#api_v1_clickhouse-proto)
    - [AverageTransactionAmountByCategoryMetric](#api-v1-AverageTransactionAmountByCategoryMetric)
    - [InvestmentTransaction](#api-v1-InvestmentTransaction)
    - [MonthlyTransactionCountByCategoryMetric](#api-v1-MonthlyTransactionCountByCategoryMetric)
    - [PersonalFinanceCategory](#api-v1-PersonalFinanceCategory)
    - [ReOccuringTransaction](#api-v1-ReOccuringTransaction)
    - [Transaction](#api-v1-Transaction)
    - [Transaction.Location](#api-v1-Transaction-Location)
    - [Transaction.PaymentMeta](#api-v1-Transaction-PaymentMeta)
    - [TransactionAmountByCountryMetric](#api-v1-TransactionAmountByCountryMetric)
    - [TransactionAmountDistributionByCategoryMetric](#api-v1-TransactionAmountDistributionByCategoryMetric)
    - [TransactionCountByMerchantPaymentChannelMetric](#api-v1-TransactionCountByMerchantPaymentChannelMetric)
  
    - [ReCurringFlow](#api-v1-ReCurringFlow)
    - [ReOccuringTransactionsFrequency](#api-v1-ReOccuringTransactionsFrequency)
    - [ReOccuringTransactionsStatus](#api-v1-ReOccuringTransactionsStatus)
  
- [api/v1/errors_ignore.proto](#api_v1_errors_ignore-proto)
    - [ErrorMessageRequest](#api-v1-ErrorMessageRequest)
    - [InternalErrorMessageResponse](#api-v1-InternalErrorMessageResponse)
    - [PathUnknownErrorMessageResponse](#api-v1-PathUnknownErrorMessageResponse)
    - [ValidationErrorMessageResponse](#api-v1-ValidationErrorMessageResponse)
  
    - [AuthErrorCode](#api-v1-AuthErrorCode)
    - [ErrorCode](#api-v1-ErrorCode)
    - [InternalErrorCode](#api-v1-InternalErrorCode)
    - [NotFoundErrorCode](#api-v1-NotFoundErrorCode)
  
- [api/v1/message.proto](#api_v1_message-proto)
    - [Apr](#api-v1-Apr)
    - [BankAccount](#api-v1-BankAccount)
    - [Budget](#api-v1-Budget)
    - [Category](#api-v1-Category)
    - [CreditAccount](#api-v1-CreditAccount)
    - [Forecast](#api-v1-Forecast)
    - [InvesmentHolding](#api-v1-InvesmentHolding)
    - [InvestmentAccount](#api-v1-InvestmentAccount)
    - [InvestmentSecurity](#api-v1-InvestmentSecurity)
    - [Link](#api-v1-Link)
    - [Milestone](#api-v1-Milestone)
    - [MortgageAccount](#api-v1-MortgageAccount)
    - [PlaidLink](#api-v1-PlaidLink)
    - [PlaidSync](#api-v1-PlaidSync)
    - [Pocket](#api-v1-Pocket)
    - [SmartGoal](#api-v1-SmartGoal)
    - [StripeSubscription](#api-v1-StripeSubscription)
    - [StudentLoanAccount](#api-v1-StudentLoanAccount)
    - [Token](#api-v1-Token)
    - [UserProfile](#api-v1-UserProfile)
  
    - [BankAccountStatus](#api-v1-BankAccountStatus)
    - [BankAccountType](#api-v1-BankAccountType)
    - [GoalStatus](#api-v1-GoalStatus)
    - [GoalType](#api-v1-GoalType)
    - [LinkStatus](#api-v1-LinkStatus)
    - [LinkType](#api-v1-LinkType)
    - [PocketType](#api-v1-PocketType)
    - [StripeSubscriptionStatus](#api-v1-StripeSubscriptionStatus)
  
- [api/v1/openapi.proto](#api_v1_openapi-proto)
- [api/v1/request_response.proto](#api_v1_request_response-proto)
    - [CreateBankAccountRequest](#api-v1-CreateBankAccountRequest)
    - [CreateBankAccountResponse](#api-v1-CreateBankAccountResponse)
    - [CreateBudgetRequest](#api-v1-CreateBudgetRequest)
    - [CreateBudgetResponse](#api-v1-CreateBudgetResponse)
    - [CreateManualLinkRequest](#api-v1-CreateManualLinkRequest)
    - [CreateManualLinkResponse](#api-v1-CreateManualLinkResponse)
    - [CreateMilestoneRequest](#api-v1-CreateMilestoneRequest)
    - [CreateMilestoneResponse](#api-v1-CreateMilestoneResponse)
    - [CreateSmartGoalRequest](#api-v1-CreateSmartGoalRequest)
    - [CreateSmartGoalResponse](#api-v1-CreateSmartGoalResponse)
    - [CreateSubscriptionRequest](#api-v1-CreateSubscriptionRequest)
    - [CreateSubscriptionResponse](#api-v1-CreateSubscriptionResponse)
    - [CreateUserProfileRequest](#api-v1-CreateUserProfileRequest)
    - [CreateUserProfileResponse](#api-v1-CreateUserProfileResponse)
    - [DeleteBankAccountRequest](#api-v1-DeleteBankAccountRequest)
    - [DeleteBankAccountResponse](#api-v1-DeleteBankAccountResponse)
    - [DeleteBudgetRequest](#api-v1-DeleteBudgetRequest)
    - [DeleteBudgetResponse](#api-v1-DeleteBudgetResponse)
    - [DeleteLinkRequest](#api-v1-DeleteLinkRequest)
    - [DeleteLinkResponse](#api-v1-DeleteLinkResponse)
    - [DeleteMilestoneRequest](#api-v1-DeleteMilestoneRequest)
    - [DeleteMilestoneResponse](#api-v1-DeleteMilestoneResponse)
    - [DeleteSmartGoalRequest](#api-v1-DeleteSmartGoalRequest)
    - [DeleteSmartGoalResponse](#api-v1-DeleteSmartGoalResponse)
    - [DeleteUserProfileRequest](#api-v1-DeleteUserProfileRequest)
    - [DeleteUserProfileResponse](#api-v1-DeleteUserProfileResponse)
    - [GetAllBudgetsRequest](#api-v1-GetAllBudgetsRequest)
    - [GetAllBudgetsResponse](#api-v1-GetAllBudgetsResponse)
    - [GetBankAccountRequest](#api-v1-GetBankAccountRequest)
    - [GetBankAccountResponse](#api-v1-GetBankAccountResponse)
    - [GetBudgetRequest](#api-v1-GetBudgetRequest)
    - [GetBudgetResponse](#api-v1-GetBudgetResponse)
    - [GetForecastRequest](#api-v1-GetForecastRequest)
    - [GetForecastResponse](#api-v1-GetForecastResponse)
    - [GetInvestmentAcccountRequest](#api-v1-GetInvestmentAcccountRequest)
    - [GetInvestmentAcccountResponse](#api-v1-GetInvestmentAcccountResponse)
    - [GetLiabilityAccountRequest](#api-v1-GetLiabilityAccountRequest)
    - [GetLiabilityAccountResponse](#api-v1-GetLiabilityAccountResponse)
    - [GetLinkRequest](#api-v1-GetLinkRequest)
    - [GetLinkResponse](#api-v1-GetLinkResponse)
    - [GetLinksRequest](#api-v1-GetLinksRequest)
    - [GetLinksResponse](#api-v1-GetLinksResponse)
    - [GetMilestoneRequest](#api-v1-GetMilestoneRequest)
    - [GetMilestoneResponse](#api-v1-GetMilestoneResponse)
    - [GetMilestonesBySmartGoalIdRequest](#api-v1-GetMilestonesBySmartGoalIdRequest)
    - [GetMilestonesBySmartGoalIdResponse](#api-v1-GetMilestonesBySmartGoalIdResponse)
    - [GetMortgageAccountRequest](#api-v1-GetMortgageAccountRequest)
    - [GetMortgageAccountResponse](#api-v1-GetMortgageAccountResponse)
    - [GetPocketRequest](#api-v1-GetPocketRequest)
    - [GetPocketResponse](#api-v1-GetPocketResponse)
    - [GetReCurringTransactionsRequest](#api-v1-GetReCurringTransactionsRequest)
    - [GetReCurringTransactionsResponse](#api-v1-GetReCurringTransactionsResponse)
    - [GetReCurringTransactionsResponse.ParticipantReCurringTransactions](#api-v1-GetReCurringTransactionsResponse-ParticipantReCurringTransactions)
    - [GetSmartGoalsByPocketIdRequest](#api-v1-GetSmartGoalsByPocketIdRequest)
    - [GetSmartGoalsByPocketIdResponse](#api-v1-GetSmartGoalsByPocketIdResponse)
    - [GetStudentLoanAccountRequest](#api-v1-GetStudentLoanAccountRequest)
    - [GetStudentLoanAccountResponse](#api-v1-GetStudentLoanAccountResponse)
    - [GetTransactionsRequest](#api-v1-GetTransactionsRequest)
    - [GetTransactionsResponse](#api-v1-GetTransactionsResponse)
    - [GetUserProfileRequest](#api-v1-GetUserProfileRequest)
    - [GetUserProfileResponse](#api-v1-GetUserProfileResponse)
    - [HealthCheckRequest](#api-v1-HealthCheckRequest)
    - [HealthCheckResponse](#api-v1-HealthCheckResponse)
    - [PlaidExchangeTokenRequest](#api-v1-PlaidExchangeTokenRequest)
    - [PlaidExchangeTokenResponse](#api-v1-PlaidExchangeTokenResponse)
    - [PlaidInitiateTokenExchangeRequest](#api-v1-PlaidInitiateTokenExchangeRequest)
    - [PlaidInitiateTokenExchangeResponse](#api-v1-PlaidInitiateTokenExchangeResponse)
    - [ProcessWebhookRequest](#api-v1-ProcessWebhookRequest)
    - [ProcessWebhookRequest.ErrorEntry](#api-v1-ProcessWebhookRequest-ErrorEntry)
    - [ProcessWebhookResponse](#api-v1-ProcessWebhookResponse)
    - [ReadynessCheckRequest](#api-v1-ReadynessCheckRequest)
    - [ReadynessCheckResponse](#api-v1-ReadynessCheckResponse)
    - [StripeWebhookRequest](#api-v1-StripeWebhookRequest)
    - [StripeWebhookResponse](#api-v1-StripeWebhookResponse)
    - [UpdateBankAccountRequest](#api-v1-UpdateBankAccountRequest)
    - [UpdateBankAccountResponse](#api-v1-UpdateBankAccountResponse)
    - [UpdateBudgetRequest](#api-v1-UpdateBudgetRequest)
    - [UpdateBudgetResponse](#api-v1-UpdateBudgetResponse)
    - [UpdateMilestoneRequest](#api-v1-UpdateMilestoneRequest)
    - [UpdateMilestoneResponse](#api-v1-UpdateMilestoneResponse)
    - [UpdateSmartGoalRequest](#api-v1-UpdateSmartGoalRequest)
    - [UpdateSmartGoalResponse](#api-v1-UpdateSmartGoalResponse)
    - [UpdateUserProfileRequest](#api-v1-UpdateUserProfileRequest)
    - [UpdateUserProfileResponse](#api-v1-UpdateUserProfileResponse)
  
- [api/v1/service.proto](#api_v1_service-proto)
    - [FinancialService](#api-v1-FinancialService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_clickhouse-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/clickhouse.proto



<a name="api-v1-AverageTransactionAmountByCategoryMetric"></a>

### AverageTransactionAmountByCategoryMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| amount | [double](#double) |  |  |






<a name="api-v1-InvestmentTransaction"></a>

### InvestmentTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_id | [string](#string) |  | @gotag: clickhouse:&#34;account_id&#34; |
| ammount | [string](#string) |  | @gotag: clickhouse:&#34;amount&#34; |
| investment_transaction_id | [string](#string) |  | @gotag: clickhouse:&#34;investment_transaction_id&#34; |
| security_id | [string](#string) |  | @gotag: clickhouse:&#34;security_id&#34; |
| date | [string](#string) |  | @gotag: clickhouse:&#34;date&#34; |
| name | [string](#string) |  | @gotag: clickhouse:&#34;name&#34; |
| quantity | [double](#double) |  | @gotag: clickhouse:&#34;quantity&#34; |
| amount | [double](#double) |  | @gotag: clickhouse:&#34;amount&#34; |
| price | [double](#double) |  | @gotag: clickhouse:&#34;price&#34; |
| fees | [double](#double) |  | @gotag: clickhouse:&#34;fees&#34; |
| type | [string](#string) |  | @gotag: clickhouse:&#34;type&#34; |
| subtype | [string](#string) |  | @gotag: clickhouse:&#34;subtype&#34; |
| iso_currency_code | [string](#string) |  | @gotag: clickhouse:&#34;iso_currency_code&#34; |
| unofficial_currency_code | [string](#string) |  | @gotag: clickhouse:&#34;unofficial_currency_code&#34; |
| link_id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;link_id&#34; |
| id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;id&#34; |
| user_id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;user_id&#34; |
| created_at | [string](#string) |  |  |






<a name="api-v1-MonthlyTransactionCountByCategoryMetric"></a>

### MonthlyTransactionCountByCategoryMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| count | [uint32](#uint32) |  |  |
| month | [string](#string) |  |  |






<a name="api-v1-PersonalFinanceCategory"></a>

### PersonalFinanceCategory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary | [string](#string) |  |  |
| detailed | [string](#string) |  |  |






<a name="api-v1-ReOccuringTransaction"></a>

### ReOccuringTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_id | [string](#string) |  | @gotag: clickhouse:&#34;account_id&#34; |
| stream_id | [string](#string) |  | @gotag: clickhouse:&#34;stream_id&#34; |
| category | [string](#string) | repeated | @gotag: clickhouse:&#34;category&#34; |
| category_id | [string](#string) |  | @gotag: clickhouse:&#34;category_id&#34; |
| description | [string](#string) |  | @gotag: clickhouse:&#34;description&#34; |
| merchant_name | [string](#string) |  | @gotag: clickhouse:&#34;merchant_name&#34; |
| personal_finance_category_primary | [string](#string) |  | @gotag: clickhouse:&#34;personal_finance_category_primary&#34; |
| personal_finance_category_detailed | [string](#string) |  | @gotag: clickhouse:&#34;personal_finance_category_detailed&#34; |
| first_date | [string](#string) |  | @gotag: clickhouse:&#34;first_date&#34; |
| last_date | [string](#string) |  | @gotag: clickhouse:&#34;last_date&#34; |
| frequency | [ReOccuringTransactionsFrequency](#api-v1-ReOccuringTransactionsFrequency) |  | @gotag: clickhouse:&#34;frequency&#34; |
| transaction_ids | [string](#string) | repeated | @gotag: clickhouse:&#34;transaction_ids&#34; |
| average_amount | [string](#string) |  | @gotag: clickhouse:&#34;average_amount&#34; |
| average_amount_iso_currency_code | [string](#string) |  | @gotag: clickhouse:&#34;average_amount_iso_currency_code&#34; |
| last_amount | [string](#string) |  | @gotag: clickhouse:&#34;last_amount&#34; |
| last_amount_iso_currency_code | [string](#string) |  | @gotag: clickhouse:&#34;last_amount_iso_currency_code&#34; |
| is_active | [bool](#bool) |  | @gotag: clickhouse:&#34;is_active&#34; |
| status | [ReOccuringTransactionsStatus](#api-v1-ReOccuringTransactionsStatus) |  | @gotag: clickhouse:&#34;status&#34; |
| updated_time | [string](#string) |  | @gotag: clickhouse:&#34;updated_time&#34; |
| user_id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;user_id&#34; |
| link_id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;link_id&#34; |
| id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;id&#34; |
| flow | [ReCurringFlow](#api-v1-ReCurringFlow) |  | @gotag: clickhouse:&#34;flow&#34; |






<a name="api-v1-Transaction"></a>

### Transaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_id | [string](#string) |  | @gotag: clickhouse:&#34;account_id&#34; |
| amount | [double](#double) |  | @gotag: clickhouse:&#34;amount&#34; |
| iso_currency_code | [string](#string) |  | @gotag: clickhouse:&#34;iso_currency_code&#34; |
| unofficial_currency_code | [string](#string) |  | @gotag: clickhouse:&#34;unofficial_currency_code&#34; |
| category | [string](#string) |  | @gotag: clickhouse:&#34;category&#34; |
| category_id | [string](#string) |  | @gotag: clickhouse:&#34;category_id&#34; |
| check_number | [string](#string) |  | @gotag: clickhouse:&#34;check_number&#34; |
| date | [string](#string) |  | @gotag: clickhouse:&#34;date&#34; |
| datetime | [string](#string) |  | @gotag: clickhouse:&#34;datetime&#34; |
| authorized_date | [string](#string) |  | @gotag: clickhouse:&#34;authorized_date&#34; |
| authorized_datetime | [string](#string) |  | @gotag: clickhouse:&#34;authorized_datetime&#34; |
| location | [Transaction.Location](#api-v1-Transaction-Location) |  | @gotag: clickhouse:&#34;location&#34; |
| name | [string](#string) |  | @gotag: clickhouse:&#34;name&#34; |
| merchant_name | [string](#string) |  | @gotag: clickhouse:&#34;merchant_name&#34; |
| payment_meta | [Transaction.PaymentMeta](#api-v1-Transaction-PaymentMeta) |  | @gotag: clickhouse:&#34;payment_meta&#34; |
| payment_channel | [string](#string) |  | @gotag: clickhouse:&#34;payment_channel&#34; |
| pending | [bool](#bool) |  | @gotag: clickhouse:&#34;pending&#34; |
| pending_transaction_id | [string](#string) |  | @gotag: clickhouse:&#34;pending_transaction_id&#34; |
| account_owner | [string](#string) |  | @gotag: clickhouse:&#34;account_owner&#34; |
| transaction_id | [string](#string) |  | @gotag: clickhouse:&#34;transaction_id&#34; |
| transaction_code | [string](#string) |  | @gotag: clickhouse:&#34;transaction_code&#34; |
| id | [uint64](#uint64) |  |  |
| user_id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;user_id&#34; |
| link_id | [uint64](#uint64) |  | @gotag: clickhouse:&#34;link_id&#34; |






<a name="api-v1-Transaction-Location"></a>

### Transaction.Location



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| city | [string](#string) |  |  |
| region | [string](#string) |  |  |
| postal_code | [string](#string) |  |  |
| country | [string](#string) |  |  |
| lat | [double](#double) |  |  |
| lon | [double](#double) |  |  |
| store_number | [string](#string) |  |  |






<a name="api-v1-Transaction-PaymentMeta"></a>

### Transaction.PaymentMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| by_order_of | [string](#string) |  |  |
| payee | [string](#string) |  |  |
| payer | [string](#string) |  |  |
| payment_method | [string](#string) |  |  |
| payment_processor | [string](#string) |  |  |
| ppd_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| reference_number | [string](#string) |  |  |






<a name="api-v1-TransactionAmountByCountryMetric"></a>

### TransactionAmountByCountryMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| country | [string](#string) |  |  |
| amount | [double](#double) |  |  |






<a name="api-v1-TransactionAmountDistributionByCategoryMetric"></a>

### TransactionAmountDistributionByCategoryMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| mean | [double](#double) |  |  |
| median | [double](#double) |  |  |
| standard_deviation | [double](#double) |  |  |






<a name="api-v1-TransactionCountByMerchantPaymentChannelMetric"></a>

### TransactionCountByMerchantPaymentChannelMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_name | [string](#string) |  |  |
| payment_channel | [string](#string) |  |  |
| transaction_count | [uint32](#uint32) |  |  |





 


<a name="api-v1-ReCurringFlow"></a>

### ReCurringFlow


| Name | Number | Description |
| ---- | ------ | ----------- |
| RE_CURRING_FLOW_UNSPECIFIED | 0 |  |
| RE_CURRING_FLOW_INFLOW | 1 |  |
| RE_CURRING_FLOW_OUTFLOW | 2 |  |



<a name="api-v1-ReOccuringTransactionsFrequency"></a>

### ReOccuringTransactionsFrequency


| Name | Number | Description |
| ---- | ------ | ----------- |
| RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED | 0 |  |
| RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY | 1 |  |
| RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY | 2 |  |
| RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY | 3 |  |
| RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY | 4 |  |
| RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY | 5 |  |



<a name="api-v1-ReOccuringTransactionsStatus"></a>

### ReOccuringTransactionsStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED | 0 |  |
| RE_OCCURING_TRANSACTIONS_STATUS_MATURE | 1 | A MATURE recurring stream should have at least 3 transactions and happen on a regular cadence (For Annual recurring stream, we will mark it MATURE after 2 instances). |
| RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION | 2 | When a recurring transaction first appears in the transaction history and before it fulfills the requirement of a mature stream, the status will be EARLY_DETECTION. |
| RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED | 3 | A stream that was previously in the EARLY_DETECTION status will move to the TOMBSTONED status when no further transactions were found at the next expected date. |


 

 

 



<a name="api_v1_errors_ignore-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/errors_ignore.proto



<a name="api-v1-ErrorMessageRequest"></a>

### ErrorMessageRequest







<a name="api-v1-InternalErrorMessageResponse"></a>

### InternalErrorMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [InternalErrorCode](#api-v1-InternalErrorCode) |  |  |
| message | [string](#string) |  |  |






<a name="api-v1-PathUnknownErrorMessageResponse"></a>

### PathUnknownErrorMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [NotFoundErrorCode](#api-v1-NotFoundErrorCode) |  |  |
| message | [string](#string) |  |  |






<a name="api-v1-ValidationErrorMessageResponse"></a>

### ValidationErrorMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ErrorCode](#api-v1-ErrorCode) |  |  |
| message | [string](#string) |  |  |





 


<a name="api-v1-AuthErrorCode"></a>

### AuthErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_auth_error | 0 |  |
| auth_failed_invalid_subject | 1001 |  |
| auth_failed_invalid_audience | 1002 |  |
| auth_failed_invalid_issuer | 1003 |  |
| invalid_claims | 1004 |  |
| auth_failed_invalid_bearer_token | 1005 |  |
| bearer_token_missing | 1010 |  |
| unauthenticated | 1500 |  |



<a name="api-v1-ErrorCode"></a>

### ErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_error | 0 |  |
| validation_error | 2000 |  |
| authorization_model_not_found | 2001 |  |
| authorization_model_resolution_too_complex | 2002 |  |
| invalid_write_input | 2003 |  |
| cannot_allow_duplicate_tuples_in_one_request | 2004 |  |
| cannot_allow_duplicate_types_in_one_request | 2005 |  |
| cannot_allow_multiple_references_to_one_relation | 2006 |  |
| invalid_continuation_token | 2007 |  |
| invalid_tuple_set | 2008 |  |
| invalid_check_input | 2009 |  |
| invalid_expand_input | 2010 |  |
| unsupported_user_set | 2011 |  |
| invalid_object_format | 2012 |  |
| write_failed_due_to_invalid_input | 2017 |  |
| authorization_model_assertions_not_found | 2018 |  |
| latest_authorization_model_not_found | 2020 |  |
| type_not_found | 2021 |  |
| relation_not_found | 2022 |  |
| empty_relation_definition | 2023 |  |
| invalid_user | 2025 |  |
| invalid_tuple | 2027 |  |
| unknown_relation | 2028 |  |
| store_id_invalid_length | 2030 |  |
| assertions_too_many_items | 2033 |  |
| id_too_long | 2034 |  |
| authorization_model_id_too_long | 2036 |  |
| tuple_key_value_not_specified | 2037 |  |
| tuple_keys_too_many_or_too_few_items | 2038 |  |
| page_size_invalid | 2039 |  |
| param_missing_value | 2040 |  |
| difference_base_missing_value | 2041 |  |
| subtract_base_missing_value | 2042 |  |
| object_too_long | 2043 |  |
| relation_too_long | 2044 |  |
| type_definitions_too_few_items | 2045 |  |
| type_invalid_length | 2046 |  |
| type_invalid_pattern | 2047 |  |
| relations_too_few_items | 2048 |  |
| relations_too_long | 2049 |  |
| relations_invalid_pattern | 2050 |  |
| object_invalid_pattern | 2051 |  |
| query_string_type_continuation_token_mismatch | 2052 |  |
| exceeded_entity_limit | 2053 |  |
| invalid_contextual_tuple | 2054 |  |
| duplicate_contextual_tuple | 2055 |  |
| invalid_authorization_model | 2056 |  |
| unsupported_schema_version | 2057 |  |



<a name="api-v1-InternalErrorCode"></a>

### InternalErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_internal_error | 0 |  |
| internal_error | 4000 |  |
| cancelled | 4003 |  |
| deadline_exceeded | 4004 |  |
| already_exists | 4005 |  |
| resource_exhausted | 4006 |  |
| failed_precondition | 4007 |  |
| aborted | 4008 |  |
| out_of_range | 4009 |  |
| unavailable | 4010 |  |
| data_loss | 4011 |  |



<a name="api-v1-NotFoundErrorCode"></a>

### NotFoundErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| no_not_found_error | 0 |  |
| undefined_endpoint | 5000 |  |
| store_id_not_found | 5002 |  |
| unimplemented | 5004 |  |


 

 

 



<a name="api_v1_message-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/message.proto



<a name="api-v1-Apr"></a>

### Apr



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| percentage | [double](#double) |  |  |
| type | [string](#string) |  |  |
| balance_subject_to_apr | [double](#double) |  |  |
| interest_charge_amount | [double](#double) |  |  |






<a name="api-v1-BankAccount"></a>

### BankAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| user_id | [uint64](#uint64) |  | the user id to which this bank account is tied to |
| name | [string](#string) |  | the bank account name |
| number | [string](#string) |  | the bank account number |
| type | [BankAccountType](#api-v1-BankAccountType) |  | the bank account type |
| balance | [float](#float) |  | the bank account balance |
| currency | [string](#string) |  | the bank account currency |
| current_funds | [double](#double) |  |  |
| balance_limit | [uint64](#uint64) |  |  |
| pockets | [Pocket](#api-v1-Pocket) | repeated | the set of &#34;virtualized accounts this user witholds&#34; NOTE: these pockets are automatically created by the system when a user connects a bank account |
| plaid_account_id | [string](#string) |  | plaid account id mapped to this bank account |
| subtype | [string](#string) |  | account subtype |
| status | [BankAccountStatus](#api-v1-BankAccountStatus) |  | the bank account status |






<a name="api-v1-Budget"></a>

### Budget
The Budgets table stores information about each budget created by the user,
including the name of the budget, the start and end dates, and the user ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| name | [string](#string) |  | The name of the budget |
| description | [string](#string) |  |  |
| start_date | [string](#string) |  | the time the goal was created |
| end_date | [string](#string) |  | the time the goal was updated |
| category | [Category](#api-v1-Category) |  | category associated with the goal |






<a name="api-v1-Category"></a>

### Category
The Categories table stores information about the different categories of expenses or income,
such as &#34;Housing&#34;, &#34;Food&#34;, &#34;Transportation&#34;, and &#34;Entertainment&#34;. Each category has one or more
subcategories, which are stored in the Subcategories table.

For example, the &#34;Housing&#34; category might have subcategories for &#34;Rent&#34;, &#34;Utilities&#34;, and &#34;Home Maintenance&#34;.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| name | [string](#string) |  | The name of the category |
| description | [string](#string) |  | The description of the category |
| subcategories | [string](#string) | repeated | the sub categories of the category |






<a name="api-v1-CreditAccount"></a>

### CreditAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| user_id | [uint64](#uint64) |  | the user id to which this bank account is tied to |
| name | [string](#string) |  | the account name |
| number | [string](#string) |  | the bank account number |
| type | [string](#string) |  | the bank account type |
| balance | [float](#float) |  | the bank account balance |
| current_funds | [double](#double) |  | current funds on the account |
| balance_limit | [uint64](#uint64) |  | balance limit |
| plaid_account_id | [string](#string) |  | plaid account id mapped to this bank account |
| subtype | [string](#string) |  | accoint subtype |
| is_overdue | [bool](#bool) |  | wether the account is overdue |
| last_payment_amount | [double](#double) |  | the last payment amount |
| last_payment_date | [string](#string) |  | the last payment date |
| last_statement_issue_date | [string](#string) |  | the last statement issue date |
| minimum_amount_due_date | [double](#double) |  | the minimum amount due date |
| next_payment_date | [string](#string) |  | the next payment date |
| aprs | [Apr](#api-v1-Apr) | repeated | the aprs |
| last_statement_balance | [double](#double) |  | the last statement balance |
| minimum_payment_amount | [double](#double) |  | the minimum payment amount |
| next_payment_due_date | [string](#string) |  | the next payment due date |






<a name="api-v1-Forecast"></a>

### Forecast
The Forecast table stores information about each forecast generated for a particular goal,
including the forecast date, the forecasted amount of money saved or invested for the
goal by the target date, and the variance between the forecasted and target amounts.
This allows the user to track how well they are progressing towards their goal and make adjustments as needed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| forecasted_amount | [string](#string) |  | the forecasted amount of the goal |
| forecasted_completion_date | [string](#string) |  | the forecasted completion date of the goal |
| variance_amount | [string](#string) |  | the forecasted variance of the goal between the forecasted and target amounts |






<a name="api-v1-InvesmentHolding"></a>

### InvesmentHolding



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| name | [string](#string) |  | The name of the investment holding |
| plaid_account_id | [string](#string) |  | plaid account id |
| cost_basis | [double](#double) |  |  |
| institution_price | [double](#double) |  |  |
| institution_price_as_of | [string](#string) |  |  |
| institution_price_datetime | [string](#string) |  |  |
| institution_value | [double](#double) |  |  |
| iso_currency_code | [string](#string) |  |  |
| quantity | [double](#double) |  |  |
| security_id | [string](#string) |  |  |
| unofficial_currency_code | [string](#string) |  |  |






<a name="api-v1-InvestmentAccount"></a>

### InvestmentAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| user_id | [uint64](#uint64) |  | the user id to which this bank account is tied to |
| name | [string](#string) |  | the account name |
| number | [string](#string) |  | the bank account number |
| type | [string](#string) |  | the bank account type |
| balance | [float](#float) |  | the bank account balance |
| current_funds | [double](#double) |  |  |
| balance_limit | [uint64](#uint64) |  |  |
| plaid_account_id | [string](#string) |  | plaid account id mapped to this bank account |
| subtype | [string](#string) |  | accoint subtype |
| holdings | [InvesmentHolding](#api-v1-InvesmentHolding) | repeated | invesment holding is the set of securities this account witholds |
| securities | [InvestmentSecurity](#api-v1-InvestmentSecurity) | repeated | the set of securities this account witholds |






<a name="api-v1-InvestmentSecurity"></a>

### InvestmentSecurity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| close_price | [double](#double) |  |  |
| close_price_as_of | [string](#string) |  |  |
| cusip | [string](#string) |  |  |
| institution_id | [string](#string) |  |  |
| institution_security_id | [string](#string) |  |  |
| is_cash_equivalent | [bool](#bool) |  |  |
| isin | [string](#string) |  |  |
| iso_currency_code | [string](#string) |  |  |
| name | [string](#string) |  |  |
| proxy_security_id | [string](#string) |  |  |
| security_id | [string](#string) |  |  |
| sedol | [string](#string) |  |  |
| ticker_symbol | [string](#string) |  |  |
| type | [string](#string) |  |  |
| unofficial_currency_code | [string](#string) |  |  |
| update_datetime | [string](#string) |  |  |






<a name="api-v1-Link"></a>

### Link
A Link represents a login at a financial institution. A single end-user of your application might have accounts at different financial
institutions, which means they would have multiple different Items. An Item is not the same as a financial institution account,
although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access
both their checking account and their savings account, a single Item would be associated with both of those accounts. Each Item 
linked within your application will have a corresponding access_token, which is a token that you can use to make API requests related
to that specific Item.
Two Items created for the same set of credentials at the same institution will be considered different and not share the same item_id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| link_status | [LinkStatus](#api-v1-LinkStatus) |  |  |
| plaid_link | [PlaidLink](#api-v1-PlaidLink) |  |  |
| plaid_new_accounts_available | [bool](#bool) |  |  |
| expiration_date | [string](#string) |  |  |
| institution_name | [string](#string) |  |  |
| custom_institution_name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| last_manual_sync | [string](#string) |  |  |
| last_successful_update | [string](#string) |  |  |
| token | [Token](#api-v1-Token) |  | token object witholds an access token which is a token used to make API requests related to a specific Item. You will typically obtain an access_token by calling /item/public_token/exchange. For more details, see the Token exchange flow. An access_token does not expire, although it may require updating, such as when a user changes their password, or when working with European institutions that comply with PSD2&#39;s 90-day consent window. For more information, see When to use update mode. Access tokens should always be stored securely, and associated with the user whose data they represent. If compromised, an access_token can be rotated via /item/access_token/invalidate. If no longer needed, it can be revoked via /item/remove.(gorm.field).has_one = {disable_association_autocreate: false disable_association_autoupdate: false preload: true}]; |
| bank_accounts | [BankAccount](#api-v1-BankAccount) | repeated | a link event - or client login event can have many connected bank accounts for example a log in link against one instition like chase can have many account (checking and savings) it is important though to ensure that if a link against an instition already exists, we dont fascilitate duplicated |
| investment_accounts | [InvestmentAccount](#api-v1-InvestmentAccount) | repeated | a link event - or client login event can have many connected investment accounts for example a log in link against one instition like fidelity can have many accounts (401k and investment account) it is important though to ensure that if a link against an instition already exists, we dont fascilitate duplicated |
| credit_accounts | [CreditAccount](#api-v1-CreditAccount) | repeated | credit accounts tied to a user |
| mortgage_accounts | [MortgageAccount](#api-v1-MortgageAccount) | repeated | mortgage accounts tied to a user |
| student_loan_accounts | [StudentLoanAccount](#api-v1-StudentLoanAccount) | repeated | student loan accounts tied to a link |
| plaid_institution_id | [string](#string) |  | the id of the institution this link is tied to and against |
| link_type | [LinkType](#api-v1-LinkType) |  | the type of link this is ... can be either a manual or plaid link type |
| error_code | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |
| new_accounts_available | [bool](#bool) |  |  |
| plaid_sync | [PlaidSync](#api-v1-PlaidSync) |  |  |






<a name="api-v1-Milestone"></a>

### Milestone
Milestone: represents a milestone in the context of simfinni. A financial milestone that is both smart
and achievable. A milestone is a sub goal of a goal and is tied to a goal by the goal id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| name | [string](#string) |  | The name of the milestone Validations: - must be at least 3 characters long |
| description | [string](#string) |  | The description of the miletone Validations: - must be at least 3 characters long |
| target_date | [string](#string) |  | the target date of the milestone Validations: - must be at least 3 characters long |
| target_amount | [string](#string) |  | the target amount of the milestone |
| is_completed | [bool](#bool) |  | wethe milestone is completed or not |
| budget | [Budget](#api-v1-Budget) |  | the budget associated with the milestone |






<a name="api-v1-MortgageAccount"></a>

### MortgageAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| plaid_account_id | [string](#string) |  |  |
| account_number | [string](#string) |  |  |
| current_late_fee | [double](#double) |  |  |
| escrow_balance | [double](#double) |  |  |
| has_pmi | [bool](#bool) |  |  |
| has_prepayment_penalty | [bool](#bool) |  |  |
| last_payment_amount | [double](#double) |  |  |
| last_payment_date | [string](#string) |  |  |
| loan_term | [string](#string) |  |  |
| loan_type_description | [string](#string) |  |  |
| maturity_date | [string](#string) |  |  |
| next_monthly_payment | [double](#double) |  |  |
| next_payment_due_date | [string](#string) |  |  |
| original_principal_balance | [double](#double) |  |  |
| original_property_value | [double](#double) |  |  |
| outstanding_principal_balance | [double](#double) |  |  |
| payment_amount | [double](#double) |  |  |
| payment_date | [string](#string) |  |  |
| origination_date | [string](#string) |  |  |
| origination_principal_amount | [double](#double) |  |  |
| past_due_amount | [double](#double) |  |  |
| ytd_interest_paid | [double](#double) |  |  |
| ytd_principal_paid | [double](#double) |  |  |
| property_address_city | [string](#string) |  |  |
| property_address_state | [string](#string) |  |  |
| property_address_street | [string](#string) |  |  |
| property_address_postal_code | [string](#string) |  |  |
| property_region | [string](#string) |  |  |
| property_country | [string](#string) |  |  |
| interest_rate_percentage | [double](#double) |  |  |
| interest_rate_type | [string](#string) |  |  |






<a name="api-v1-PlaidLink"></a>

### PlaidLink



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| products | [string](#string) | repeated |  |
| webhook_url | [string](#string) |  |  |
| institution_id | [string](#string) |  |  |
| institution_name | [string](#string) |  |  |
| use_plaid_sync | [bool](#bool) |  |  |
| item_id | [string](#string) |  |  |






<a name="api-v1-PlaidSync"></a>

### PlaidSync



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| time_stamp | [string](#string) |  |  |
| trigger | [string](#string) |  |  |
| next_cursor | [string](#string) |  |  |
| added | [int64](#int64) |  |  |
| removed | [int64](#int64) |  |  |
| modified | [int64](#int64) |  |  |






<a name="api-v1-Pocket"></a>

### Pocket
Pocket is an abstraction of a over a bank account. A user can has at most 4 pockets per connected account
NOTE: these pockets are automatically created by the system and should not be exposed for mutation
by any client. The only operations that can be performed against a pocket are:
1. Get the pocket
2. Get the pocket&#39;s smart goals
3. Adding a smart goal to the pocket


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| goals | [SmartGoal](#api-v1-SmartGoal) | repeated | the set of smart goals this user witholds |
| type | [PocketType](#api-v1-PocketType) |  | The type of the pocket |






<a name="api-v1-SmartGoal"></a>

### SmartGoal
SmartGoal: The Goals table stores information about each financial goal, including the name of the goal,
its description, the target amount of money the user wants to save or invest, and the expected date of completion.

The Goals table also includes columns for the start date of the goal, the current amount of money saved or
invested towards the goal, and a boolean flag indicating whether the goal has been achieved.
These additional columns allow the user to track their progress towards the goal and see how much
more they need to save or invest to reach their target amount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| user_id | [uint64](#uint64) |  | the user id to which this goal is tied to |
| name | [string](#string) |  | The name of the goal Validations: - must be at least 3 characters long |
| description | [string](#string) |  | The description of the goal Validations: - must be at least 3 characters long |
| is_completed | [bool](#bool) |  | wether the goal has been achieved or not |
| goal_type | [GoalType](#api-v1-GoalType) |  | The type of the goal |
| duration | [string](#string) |  | The duration of the goal |
| start_date | [string](#string) |  | the start date of the goal |
| end_date | [string](#string) |  | the end date of the goal |
| target_amount | [string](#string) |  | the target amount of the goal amount of money the user wants to save or invest |
| current_amount | [string](#string) |  | the current amount of the goal current amount of money saved or invested towards the goal |
| milestones | [Milestone](#api-v1-Milestone) | repeated | Milestones associated with the goal |
| forecasts | [Forecast](#api-v1-Forecast) |  | Forecasts associated with the goal |






<a name="api-v1-StripeSubscription"></a>

### StripeSubscription
StripeSubscription stores high level stripe subscription details of which the user profile has


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| stripe_subscription_id | [string](#string) |  | stripe subscription id tied to the customer |
| stripe_subscription_status | [StripeSubscriptionStatus](#api-v1-StripeSubscriptionStatus) |  | stripe subscription status |
| stripe_subscription_active_until | [string](#string) |  | stripe subscription active until |
| stripe_webhook_latest_timestamp | [string](#string) |  | stripe webhook latest timestamp |






<a name="api-v1-StudentLoanAccount"></a>

### StudentLoanAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| plaid_account_id | [string](#string) |  |  |
| disbursement_dates | [string](#string) | repeated |  |
| expected_payoff_date | [string](#string) |  |  |
| guarantor | [string](#string) |  |  |
| interest_rate_percentage | [double](#double) |  |  |
| is_overdue | [bool](#bool) |  |  |
| last_payment_amount | [double](#double) |  |  |
| last_payment_date | [string](#string) |  |  |
| last_statement_issue_date | [string](#string) |  |  |
| loan_name | [string](#string) |  |  |
| loan_end_date | [string](#string) |  |  |
| minimum_payment_amount | [double](#double) |  |  |
| next_payment_due_date | [string](#string) |  |  |
| origination_date | [string](#string) |  |  |
| origination_principal_amount | [double](#double) |  |  |
| outstanding_interest_amount | [double](#double) |  |  |
| payment_reference_number | [string](#string) |  |  |
| sequence_number | [string](#string) |  |  |
| ytd_interest_paid | [double](#double) |  |  |
| ytd_principal_paid | [double](#double) |  |  |
| loan_type | [string](#string) |  |  |
| pslf_status_estimated_eligibility_date | [string](#string) |  |  |
| pslf_status_payments_made | [int32](#int32) |  |  |
| pslf_status_payments_remaining | [int32](#int32) |  |  |
| repayment_plan_type | [string](#string) |  |  |
| repayment_plan_description | [string](#string) |  |  |
| servicer_address_city | [string](#string) |  |  |
| servicer_address_postal_code | [string](#string) |  |  |
| servicer_address_state | [string](#string) |  |  |
| servicer_address_street | [string](#string) |  |  |
| servicer_address_region | [string](#string) |  |  |
| servicer_address_country | [string](#string) |  |  |
| user_id | [uint64](#uint64) |  | the user id to which this bank account is tied to |
| name | [string](#string) |  | the account name |






<a name="api-v1-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| item_id | [string](#string) |  | the id of the item the token is tied to |
| key_id | [string](#string) |  |  |
| access_token | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="api-v1-UserProfile"></a>

### UserProfile
UserProfile stores high level user profile details
such as the id, user_id tied to the profile, and many more


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | id |
| user_id | [uint64](#uint64) |  | the user id tied to the profile |
| stripe_customer_id | [string](#string) |  |  |
| stripe_subscriptions | [StripeSubscription](#api-v1-StripeSubscription) |  | the stripe subscriptions the user profile actively maintains |
| link | [Link](#api-v1-Link) | repeated | a user profile can have many links (connected institutions) of which finanical accounts are tied to (checking, savings, etc) |





 


<a name="api-v1-BankAccountStatus"></a>

### BankAccountStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| BANK_ACCOUNT_STATUS_UNSPECIFIED | 0 |  |
| BANK_ACCOUNT_STATUS_ACTIVE | 1 |  |
| BANK_ACCOUNT_STATUS_INACTIVE | 2 |  |



<a name="api-v1-BankAccountType"></a>

### BankAccountType


| Name | Number | Description |
| ---- | ------ | ----------- |
| BANK_ACCOUNT_TYPE_UNSPECIFIED | 0 |  |
| BANK_ACCOUNT_TYPE_PLAID | 1 |  |
| BANK_ACCOUNT_TYPE_MANUAL | 2 |  |



<a name="api-v1-GoalStatus"></a>

### GoalStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| GOAL_STATUS_UNSPECIFIED | 0 |  |
| GOAL_STATUS_ACTIVE | 1 |  |
| GOAL_STATUS_INACTIVE | 2 |  |
| GOAL_STATUS_COMPLETED | 3 |  |
| GOAL_STATUS_DELETE | 4 |  |



<a name="api-v1-GoalType"></a>

### GoalType


| Name | Number | Description |
| ---- | ------ | ----------- |
| GOAL_TYPE_UNSPECIFIED | 0 |  |
| GOAL_TYPE_SAVINGS | 1 |  |
| GOAL_TYPE_INVESTMENT | 2 |  |
| GOAL_TYPE_DEBT | 3 |  |
| GOAL_TYPE_EXPENSE | 4 |  |



<a name="api-v1-LinkStatus"></a>

### LinkStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| LINK_STATUS_UNSPECIFIED | 0 |  |
| LINK_STATUS_SETUP | 1 |  |
| LINK_STATUS_PENDING | 2 |  |
| LINK_STATUS_ERROR | 3 |  |
| LINK_STATUS_SUCCESS | 4 |  |
| LINK_STATUS_PENDING_EXPIRATION | 5 |  |
| LINK_STATUS_REVOKED | 6 |  |



<a name="api-v1-LinkType"></a>

### LinkType


| Name | Number | Description |
| ---- | ------ | ----------- |
| LINK_TYPE_UNSPECIFIED | 0 |  |
| LINK_TYPE_PLAID | 1 |  |
| LINK_TYPE_MANUAL | 2 |  |



<a name="api-v1-PocketType"></a>

### PocketType


| Name | Number | Description |
| ---- | ------ | ----------- |
| POCKET_TYPE_UNSPECIFIED | 0 |  |
| POCKET_TYPE_DISCRETIONARY_SPENDING | 1 |  |
| POCKET_TYPE_FUN_MONEY | 2 |  |
| POCKET_TYPE_DEBT_REDUCTION | 3 |  |
| POCKET_TYPE_EMERGENCY_FUND | 4 |  |
| POCKET_TYPE_INVESTMENT | 5 |  |
| POCKET_TYPE_SHORT_TERM_SAVINGS | 6 |  |
| POCKET_TYPE_LONG_TERM_SAVINGS | 7 |  |



<a name="api-v1-StripeSubscriptionStatus"></a>

### StripeSubscriptionStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED | 0 |  |
| STRIPE_SUBSCRIPTION_STATUS_TRIALING | 1 |  |
| STRIPE_SUBSCRIPTION_STATUS_ACTIVE | 2 |  |
| STRIPE_SUBSCRIPTION_STATUS_PAST_DUE | 3 |  |
| STRIPE_SUBSCRIPTION_STATUS_CANCELED | 4 |  |
| STRIPE_SUBSCRIPTION_STATUS_UNPAID | 5 |  |
| STRIPE_SUBSCRIPTION_STATUS_COMPLETE | 6 |  |
| STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE | 7 |  |
| STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED | 8 |  |


 

 

 



<a name="api_v1_openapi-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/openapi.proto


 

 

 

 



<a name="api_v1_request_response-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/request_response.proto



<a name="api-v1-CreateBankAccountRequest"></a>

### CreateBankAccountRequest
CreateBankAccountRequest: Represents the request object invoked against the financial
service to create a bank account for a given user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The account ID associated with the user Validations: - user_id must be greater than 0 |
| bank_account | [BankAccount](#api-v1-BankAccount) |  | The bank account to create Validations: - cannot be nil hence required |






<a name="api-v1-CreateBankAccountResponse"></a>

### CreateBankAccountResponse
CreateBankAccountResponse: Represents the response object returned as a response to
the `create bank account` request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bank_account_id | [uint64](#uint64) |  | The bank account id |






<a name="api-v1-CreateBudgetRequest"></a>

### CreateBudgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestrone_id | [uint64](#uint64) |  | The milestone to associate this budget with |
| budget | [Budget](#api-v1-Budget) |  | The budget to create Validations: - cannot be nil hence required |






<a name="api-v1-CreateBudgetResponse"></a>

### CreateBudgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budget_id | [uint64](#uint64) |  | The budget id |






<a name="api-v1-CreateManualLinkRequest"></a>

### CreateManualLinkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| manual_account_link | [Link](#api-v1-Link) |  | The manual account link |






<a name="api-v1-CreateManualLinkResponse"></a>

### CreateManualLinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link_id | [uint64](#uint64) |  | The link&#39;s id |






<a name="api-v1-CreateMilestoneRequest"></a>

### CreateMilestoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id Validations: - smart_goal_id must be greater than 0 |
| milestone | [Milestone](#api-v1-Milestone) |  | The milestone to create Validations: - cannot be nil hence required |






<a name="api-v1-CreateMilestoneResponse"></a>

### CreateMilestoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestone_id | [uint64](#uint64) |  | The milestone id |






<a name="api-v1-CreateSmartGoalRequest"></a>

### CreateSmartGoalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pocket_id | [uint64](#uint64) |  | The pocket account id Validations: - pocket_account_id must be greater than 0 |
| smart_goal | [SmartGoal](#api-v1-SmartGoal) |  | The smart goal to create Validations: - cannot be nil hence required |






<a name="api-v1-CreateSmartGoalResponse"></a>

### CreateSmartGoalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id |






<a name="api-v1-CreateSubscriptionRequest"></a>

### CreateSubscriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  |  |
| price_id | [string](#string) |  |  |






<a name="api-v1-CreateSubscriptionResponse"></a>

### CreateSubscriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription_id | [string](#string) |  |  |
| payment_intent_client_secret | [string](#string) |  |  |






<a name="api-v1-CreateUserProfileRequest"></a>

### CreateUserProfileRequest
CreateUserProfileRequest: Represents the request object invoked against the user
service to create a user profile


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-v1-UserProfile) |  | User profile to create Validations: - cannot be nil hence required |
| email | [string](#string) |  | the email of the account to create |






<a name="api-v1-CreateUserProfileResponse"></a>

### CreateUserProfileResponse
CreateUserProfileResponse: Represents the response object returned as a response to
the `create user profile` request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  |  |






<a name="api-v1-DeleteBankAccountRequest"></a>

### DeleteBankAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The account ID associated with the user Validations: - user_id must be greater than 0 |
| bank_account_id | [uint64](#uint64) |  | The bank account id Validations: - bank_account_id must be greater than 0 |






<a name="api-v1-DeleteBankAccountResponse"></a>

### DeleteBankAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [bool](#bool) |  | The bank account id |






<a name="api-v1-DeleteBudgetRequest"></a>

### DeleteBudgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budget_id | [uint64](#uint64) |  | The budget id Validations: - budget_id must be greater than 0 |






<a name="api-v1-DeleteBudgetResponse"></a>

### DeleteBudgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [bool](#bool) |  | The budget id |






<a name="api-v1-DeleteLinkRequest"></a>

### DeleteLinkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| link_id | [uint64](#uint64) |  | The link id Validations: - link_id must be greater than 0 |






<a name="api-v1-DeleteLinkResponse"></a>

### DeleteLinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link_id | [uint64](#uint64) |  | The link&#39;s id |






<a name="api-v1-DeleteMilestoneRequest"></a>

### DeleteMilestoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestone_id | [uint64](#uint64) |  | The milestone id Validations: - milestone_id must be greater than 0 |






<a name="api-v1-DeleteMilestoneResponse"></a>

### DeleteMilestoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [bool](#bool) |  | The milestone id |






<a name="api-v1-DeleteSmartGoalRequest"></a>

### DeleteSmartGoalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id Validations: - smart_goal_id must be greater than 0 |






<a name="api-v1-DeleteSmartGoalResponse"></a>

### DeleteSmartGoalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [bool](#bool) |  | The smart goal id |






<a name="api-v1-DeleteUserProfileRequest"></a>

### DeleteUserProfileRequest
teUserProfileRequest: Represents the request object invoked against the user
service to delete a user profile


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The account ID associated with the user. NOTE: This user_id is the simfiny backend platform wide user id Validations: - user_id must be greater than 0 |






<a name="api-v1-DeleteUserProfileResponse"></a>

### DeleteUserProfileResponse
DeleteUserProfileResponse: Represents the response object returned as a response to
the `delete user profile` request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile_deleted | [bool](#bool) |  |  |






<a name="api-v1-GetAllBudgetsRequest"></a>

### GetAllBudgetsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pocket_id | [uint64](#uint64) |  | The pocket account id Validations: - pocket_account_id must be greater than 0 |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id Validations: - smart_goal_id must be greater than 0 |
| milestone_id | [uint64](#uint64) |  | The milestone id Validations: - milestone_id must be greater than 0 |






<a name="api-v1-GetAllBudgetsResponse"></a>

### GetAllBudgetsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budgets | [Budget](#api-v1-Budget) | repeated | The budgets |






<a name="api-v1-GetBankAccountRequest"></a>

### GetBankAccountRequest
GetBankAccountRequest: Represents the request object invoked against the financial
service to get a bank account for a given user and bank account id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bank_account_id | [uint64](#uint64) |  | The bank account id Validations: - bank_account_id must be greater than 0 |






<a name="api-v1-GetBankAccountResponse"></a>

### GetBankAccountResponse
GetBankAccountResponse: Represents the response object returned as a response to
the `get bank account` request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bank_account | [BankAccount](#api-v1-BankAccount) |  | The bank account |






<a name="api-v1-GetBudgetRequest"></a>

### GetBudgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budget_id | [uint64](#uint64) |  | The budget id Validations: - budget_id must be greater than 0 |






<a name="api-v1-GetBudgetResponse"></a>

### GetBudgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budget | [Budget](#api-v1-Budget) |  | The budget |






<a name="api-v1-GetForecastRequest"></a>

### GetForecastRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id Validations: - smart_goal_id must be greater than 0 |






<a name="api-v1-GetForecastResponse"></a>

### GetForecastResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forecast | [Forecast](#api-v1-Forecast) |  | The forecast |






<a name="api-v1-GetInvestmentAcccountRequest"></a>

### GetInvestmentAcccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| investment_account_id | [uint64](#uint64) |  | The investment account id Validations: - investment_account_id must be greater than 0 |






<a name="api-v1-GetInvestmentAcccountResponse"></a>

### GetInvestmentAcccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| investment_account | [InvestmentAccount](#api-v1-InvestmentAccount) |  | The investment account |






<a name="api-v1-GetLiabilityAccountRequest"></a>

### GetLiabilityAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| liability_account_id | [uint64](#uint64) |  | The liability account id Validations: - liability_account_id must be greater than 0 |






<a name="api-v1-GetLiabilityAccountResponse"></a>

### GetLiabilityAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| liability_account | [CreditAccount](#api-v1-CreditAccount) |  | The liability account |






<a name="api-v1-GetLinkRequest"></a>

### GetLinkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| link_id | [uint64](#uint64) |  | The link id Validations: - link_id must be greater than 0 |






<a name="api-v1-GetLinkResponse"></a>

### GetLinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link | [Link](#api-v1-Link) |  | The link |






<a name="api-v1-GetLinksRequest"></a>

### GetLinksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |






<a name="api-v1-GetLinksResponse"></a>

### GetLinksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| links | [Link](#api-v1-Link) | repeated | The links |






<a name="api-v1-GetMilestoneRequest"></a>

### GetMilestoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestone_id | [uint64](#uint64) |  | The milestone id Validations: - milestone_id must be greater than 0 |






<a name="api-v1-GetMilestoneResponse"></a>

### GetMilestoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestone | [Milestone](#api-v1-Milestone) |  | The milestone |






<a name="api-v1-GetMilestonesBySmartGoalIdRequest"></a>

### GetMilestonesBySmartGoalIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id Validations: - smart_goal_id must be greater than 0 |






<a name="api-v1-GetMilestonesBySmartGoalIdResponse"></a>

### GetMilestonesBySmartGoalIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestones | [Milestone](#api-v1-Milestone) | repeated | The milestones |






<a name="api-v1-GetMortgageAccountRequest"></a>

### GetMortgageAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| mortgage_account_id | [uint64](#uint64) |  | The mortage account id Validations: - mortage_account_id must be greater than 0 |






<a name="api-v1-GetMortgageAccountResponse"></a>

### GetMortgageAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mortage_account | [MortgageAccount](#api-v1-MortgageAccount) |  | The mortage account |






<a name="api-v1-GetPocketRequest"></a>

### GetPocketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pocket_id | [uint64](#uint64) |  | The pocket account id Validations: - pocket_account_id must be greater than 0 |






<a name="api-v1-GetPocketResponse"></a>

### GetPocketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pocket | [Pocket](#api-v1-Pocket) |  | The pocket account |






<a name="api-v1-GetReCurringTransactionsRequest"></a>

### GetReCurringTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |






<a name="api-v1-GetReCurringTransactionsResponse"></a>

### GetReCurringTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| re_ccuring_transactions | [ReOccuringTransaction](#api-v1-ReOccuringTransaction) | repeated | The re-occuring transactions |
| participant_re_ccuring_transactions | [GetReCurringTransactionsResponse.ParticipantReCurringTransactions](#api-v1-GetReCurringTransactionsResponse-ParticipantReCurringTransactions) | repeated |  |






<a name="api-v1-GetReCurringTransactionsResponse-ParticipantReCurringTransactions"></a>

### GetReCurringTransactionsResponse.ParticipantReCurringTransactions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reocurring_transaction_id | [uint64](#uint64) |  | The participant id |
| transactions | [Transaction](#api-v1-Transaction) | repeated | The transactions |






<a name="api-v1-GetSmartGoalsByPocketIdRequest"></a>

### GetSmartGoalsByPocketIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pocket_id | [uint64](#uint64) |  | The pocket account id Validations: - pocket_account_id must be greater than 0 |






<a name="api-v1-GetSmartGoalsByPocketIdResponse"></a>

### GetSmartGoalsByPocketIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goals | [SmartGoal](#api-v1-SmartGoal) | repeated | The smart goals |






<a name="api-v1-GetStudentLoanAccountRequest"></a>

### GetStudentLoanAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| student_loan_account_id | [uint64](#uint64) |  | The student loan account id Validations: - student_loan_account_id must be greater than 0 |






<a name="api-v1-GetStudentLoanAccountResponse"></a>

### GetStudentLoanAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| student_loan_account | [StudentLoanAccount](#api-v1-StudentLoanAccount) |  | The student loan account |






<a name="api-v1-GetTransactionsRequest"></a>

### GetTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| page_number | [uint64](#uint64) |  |  |
| page_size | [uint64](#uint64) |  |  |






<a name="api-v1-GetTransactionsResponse"></a>

### GetTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactions | [Transaction](#api-v1-Transaction) | repeated | The transactions |
| next_page_number | [uint64](#uint64) |  |  |






<a name="api-v1-GetUserProfileRequest"></a>

### GetUserProfileRequest
GetUserProfileRequest: Represents the request object invoked against the user
service to get a user profile


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The account ID associated with the user. NOTE: This user_id is the simfiny backend platform wide user id Validations: - user_id must be greater than 0 |






<a name="api-v1-GetUserProfileResponse"></a>

### GetUserProfileResponse
GetUserProfileResponse: Represents the response object returned as a response to
the `get user profile` request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-v1-UserProfile) |  |  |






<a name="api-v1-HealthCheckRequest"></a>

### HealthCheckRequest







<a name="api-v1-HealthCheckResponse"></a>

### HealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| healthy | [bool](#bool) |  |  |






<a name="api-v1-PlaidExchangeTokenRequest"></a>

### PlaidExchangeTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | The user id Validations: - user_id must be greater than 0 |
| public_token | [string](#string) |  | The public token Validations: - cannot be nil hence required |
| institution_id | [string](#string) |  | The institution id |
| institution_name | [string](#string) |  | The institution name |






<a name="api-v1-PlaidExchangeTokenResponse"></a>

### PlaidExchangeTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | wether the operation was successful |






<a name="api-v1-PlaidInitiateTokenExchangeRequest"></a>

### PlaidInitiateTokenExchangeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | A unique ID representing the end user. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`. It is currently used as a means of searching logs for the given user in the Plaid Dashboard. Validations: - user_id must be greater than 0 |
| full_name | [string](#string) |  | The user&#39;s full legal name. This is an optional field used in the [returning user experience](https://plaid.com/docs/link/returning-user) to associate Items to the user. |
| email | [string](#string) |  | The user&#39;s email address. This field is optional, but required to enable the [pre-authenticated returning user flow](https://plaid.com/docs/link/returning-user/#enabling-the-returning-user-experience). |
| phone_number | [string](#string) |  | The user&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. This field is optional, but required to enable the [returning user experience](https://plaid.com/docs/link/returning-user). |






<a name="api-v1-PlaidInitiateTokenExchangeResponse"></a>

### PlaidInitiateTokenExchangeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link_token | [string](#string) |  |  |
| expiration | [string](#string) |  |  |
| plaid_request_id | [string](#string) |  |  |






<a name="api-v1-ProcessWebhookRequest"></a>

### ProcessWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webhook_type | [string](#string) |  |  |
| webhook_code | [string](#string) |  |  |
| item_id | [string](#string) |  | The item_id of the Item associated with this webhook, warning, or error |
| initial_update_complete | [bool](#bool) |  | Indicates if initial pull information is available. |
| historical_update_complete | [string](#string) |  | Indicates if historical pull information is available. |
| environment | [string](#string) |  | The Plaid environment the webhook was sent from |
| new_transactions | [string](#string) | repeated | The number of new, unfetched transactions available |
| removed_transactions | [string](#string) | repeated | An array of transaction_ids corresponding to the removed transactions |
| error | [ProcessWebhookRequest.ErrorEntry](#api-v1-ProcessWebhookRequest-ErrorEntry) | repeated | We use standard HTTP response codes for success and failure notifications, and our errors are further classified by error_type. In general, 200 HTTP codes correspond to success, 40X codes are for developer- or user-related failures, and 50X codes are for Plaid-related issues. An Item with a non-null error object will only be part of an API response when calling /item/get to view Item status. Otherwise, error fields will be null if no error has occurred; if an error has occurred, an error code will be returned instead. |
| account_ids | [string](#string) | repeated | A list of account_ids for accounts that have new or updated recurring transactions data. |
| consent_expiration_time | [string](#string) |  | The time at which the user&#39;s access_token will expire. This field will only be present |
| account_ids_with_new_liabilities | [string](#string) | repeated | An array of account_id&#39;s for accounts that contain new liabilities.&#39; |
| account_ids_with_updated_liabilities | [string](#string) | repeated | An object with keys of account_id&#39;s that are mapped to their respective liabilities fields that changed. |
| new_holdings | [uint64](#uint64) |  | The number of new holdings reported since the last time this webhook was fired. |
| updated_holdings | [uint64](#uint64) |  | The number of updated holdings reported since the last time this webhook was fired. @gotag: json:&#34;updated_holdings&#34; |






<a name="api-v1-ProcessWebhookRequest-ErrorEntry"></a>

### ProcessWebhookRequest.ErrorEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="api-v1-ProcessWebhookResponse"></a>

### ProcessWebhookResponse







<a name="api-v1-ReadynessCheckRequest"></a>

### ReadynessCheckRequest







<a name="api-v1-ReadynessCheckResponse"></a>

### ReadynessCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| healthy | [bool](#bool) |  |  |






<a name="api-v1-StripeWebhookRequest"></a>

### StripeWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| body | [string](#string) |  |  |
| signature | [string](#string) |  |  |






<a name="api-v1-StripeWebhookResponse"></a>

### StripeWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="api-v1-UpdateBankAccountRequest"></a>

### UpdateBankAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bank_account | [BankAccount](#api-v1-BankAccount) |  | The bank account to update Validations: - cannot be nil hence required |






<a name="api-v1-UpdateBankAccountResponse"></a>

### UpdateBankAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [bool](#bool) |  | The bank account id |
| bank_account | [BankAccount](#api-v1-BankAccount) |  | The bank account |






<a name="api-v1-UpdateBudgetRequest"></a>

### UpdateBudgetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budget | [Budget](#api-v1-Budget) |  | The budget to update Validations: - cannot be nil hence required |






<a name="api-v1-UpdateBudgetResponse"></a>

### UpdateBudgetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| budget | [Budget](#api-v1-Budget) |  | The budget id |






<a name="api-v1-UpdateMilestoneRequest"></a>

### UpdateMilestoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestone | [Milestone](#api-v1-Milestone) |  | The milestone to update Validations: - cannot be nil hence required |






<a name="api-v1-UpdateMilestoneResponse"></a>

### UpdateMilestoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| milestone | [Milestone](#api-v1-Milestone) |  | The milestone id |






<a name="api-v1-UpdateSmartGoalRequest"></a>

### UpdateSmartGoalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal | [SmartGoal](#api-v1-SmartGoal) |  | The smart goal to update Validations: - cannot be nil hence required |






<a name="api-v1-UpdateSmartGoalResponse"></a>

### UpdateSmartGoalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| smart_goal_id | [uint64](#uint64) |  | The smart goal id |






<a name="api-v1-UpdateUserProfileRequest"></a>

### UpdateUserProfileRequest
UpdateUserProfileRequest: Represents the request object invoked against the user
service to update a user profile


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [UserProfile](#api-v1-UserProfile) |  | User profile to update Validation: - cannot nil hence required |






<a name="api-v1-UpdateUserProfileResponse"></a>

### UpdateUserProfileResponse
UpdateUserProfileResponse: Represents the response object returned as a response to
the `update user profile` request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile_updated | [bool](#bool) |  |  |
| profile | [UserProfile](#api-v1-UserProfile) |  |  |





 

 

 

 



<a name="api_v1_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/service.proto


 

 

 


<a name="api-v1-FinancialService"></a>

### FinancialService
FinancialService API.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PlaidInitiateTokenExchange | [PlaidInitiateTokenExchangeRequest](#api-v1-PlaidInitiateTokenExchangeRequest) | [PlaidInitiateTokenExchangeResponse](#api-v1-PlaidInitiateTokenExchangeResponse) |  |
| PlaidExchangeToken | [PlaidExchangeTokenRequest](#api-v1-PlaidExchangeTokenRequest) | [PlaidExchangeTokenResponse](#api-v1-PlaidExchangeTokenResponse) |  |
| CreateUserProfile | [CreateUserProfileRequest](#api-v1-CreateUserProfileRequest) | [CreateUserProfileResponse](#api-v1-CreateUserProfileResponse) |  |
| GetUserProfile | [GetUserProfileRequest](#api-v1-GetUserProfileRequest) | [GetUserProfileResponse](#api-v1-GetUserProfileResponse) |  |
| DeleteUserProfile | [DeleteUserProfileRequest](#api-v1-DeleteUserProfileRequest) | [DeleteUserProfileResponse](#api-v1-DeleteUserProfileResponse) |  |
| UpdateUserProfile | [UpdateUserProfileRequest](#api-v1-UpdateUserProfileRequest) | [UpdateUserProfileResponse](#api-v1-UpdateUserProfileResponse) |  |
| CreateBankAccount | [CreateBankAccountRequest](#api-v1-CreateBankAccountRequest) | [CreateBankAccountResponse](#api-v1-CreateBankAccountResponse) |  |
| GetBankAccount | [GetBankAccountRequest](#api-v1-GetBankAccountRequest) | [GetBankAccountResponse](#api-v1-GetBankAccountResponse) |  |
| UpdateBankAccount | [UpdateBankAccountRequest](#api-v1-UpdateBankAccountRequest) | [UpdateBankAccountResponse](#api-v1-UpdateBankAccountResponse) |  |
| DeleteBankAccount | [DeleteBankAccountRequest](#api-v1-DeleteBankAccountRequest) | [DeleteBankAccountResponse](#api-v1-DeleteBankAccountResponse) |  |
| GetPocket | [GetPocketRequest](#api-v1-GetPocketRequest) | [GetPocketResponse](#api-v1-GetPocketResponse) |  |
| GetSmartGoalsByPocketId | [GetSmartGoalsByPocketIdRequest](#api-v1-GetSmartGoalsByPocketIdRequest) | [GetSmartGoalsByPocketIdResponse](#api-v1-GetSmartGoalsByPocketIdResponse) |  |
| CreateSmartGoal | [CreateSmartGoalRequest](#api-v1-CreateSmartGoalRequest) | [CreateSmartGoalResponse](#api-v1-CreateSmartGoalResponse) |  |
| UpdateSmartGoal | [UpdateSmartGoalRequest](#api-v1-UpdateSmartGoalRequest) | [UpdateSmartGoalResponse](#api-v1-UpdateSmartGoalResponse) |  |
| DeleteSmartGoal | [DeleteSmartGoalRequest](#api-v1-DeleteSmartGoalRequest) | [DeleteSmartGoalResponse](#api-v1-DeleteSmartGoalResponse) |  |
| CreateMilestone | [CreateMilestoneRequest](#api-v1-CreateMilestoneRequest) | [CreateMilestoneResponse](#api-v1-CreateMilestoneResponse) |  |
| DeleteMilestone | [DeleteMilestoneRequest](#api-v1-DeleteMilestoneRequest) | [DeleteMilestoneResponse](#api-v1-DeleteMilestoneResponse) |  |
| UpdateMilestone | [UpdateMilestoneRequest](#api-v1-UpdateMilestoneRequest) | [UpdateMilestoneResponse](#api-v1-UpdateMilestoneResponse) |  |
| GetMilestone | [GetMilestoneRequest](#api-v1-GetMilestoneRequest) | [GetMilestoneResponse](#api-v1-GetMilestoneResponse) |  |
| GetMilestonesBySmartGoalId | [GetMilestonesBySmartGoalIdRequest](#api-v1-GetMilestonesBySmartGoalIdRequest) | [GetMilestonesBySmartGoalIdResponse](#api-v1-GetMilestonesBySmartGoalIdResponse) |  |
| GetForecast | [GetForecastRequest](#api-v1-GetForecastRequest) | [GetForecastResponse](#api-v1-GetForecastResponse) |  |
| CreateBudget | [CreateBudgetRequest](#api-v1-CreateBudgetRequest) | [CreateBudgetResponse](#api-v1-CreateBudgetResponse) |  |
| UpdateBudget | [UpdateBudgetRequest](#api-v1-UpdateBudgetRequest) | [UpdateBudgetResponse](#api-v1-UpdateBudgetResponse) |  |
| DeleteBudget | [DeleteBudgetRequest](#api-v1-DeleteBudgetRequest) | [DeleteBudgetResponse](#api-v1-DeleteBudgetResponse) |  |
| GetBudget | [GetBudgetRequest](#api-v1-GetBudgetRequest) | [GetBudgetResponse](#api-v1-GetBudgetResponse) |  |
| GetAllBudgets | [GetAllBudgetsRequest](#api-v1-GetAllBudgetsRequest) | [GetAllBudgetsResponse](#api-v1-GetAllBudgetsResponse) |  |
| HealthCheck | [HealthCheckRequest](#api-v1-HealthCheckRequest) | [HealthCheckResponse](#api-v1-HealthCheckResponse) |  |
| ReadynessCheck | [ReadynessCheckRequest](#api-v1-ReadynessCheckRequest) | [ReadynessCheckResponse](#api-v1-ReadynessCheckResponse) |  |
| GetInvestmentAcccount | [GetInvestmentAcccountRequest](#api-v1-GetInvestmentAcccountRequest) | [GetInvestmentAcccountResponse](#api-v1-GetInvestmentAcccountResponse) |  |
| GetMortgageAccount | [GetMortgageAccountRequest](#api-v1-GetMortgageAccountRequest) | [GetMortgageAccountResponse](#api-v1-GetMortgageAccountResponse) |  |
| GetLiabilityAccount | [GetLiabilityAccountRequest](#api-v1-GetLiabilityAccountRequest) | [GetLiabilityAccountResponse](#api-v1-GetLiabilityAccountResponse) |  |
| GetStudentLoanAccount | [GetStudentLoanAccountRequest](#api-v1-GetStudentLoanAccountRequest) | [GetStudentLoanAccountResponse](#api-v1-GetStudentLoanAccountResponse) |  |
| CreateManualLink | [CreateManualLinkRequest](#api-v1-CreateManualLinkRequest) | [CreateManualLinkResponse](#api-v1-CreateManualLinkResponse) |  |
| GetLink | [GetLinkRequest](#api-v1-GetLinkRequest) | [GetLinkResponse](#api-v1-GetLinkResponse) |  |
| GetLinks | [GetLinksRequest](#api-v1-GetLinksRequest) | [GetLinksResponse](#api-v1-GetLinksResponse) |  |
| DeleteLink | [DeleteLinkRequest](#api-v1-DeleteLinkRequest) | [DeleteLinkResponse](#api-v1-DeleteLinkResponse) |  |
| GetReCurringTransactions | [GetReCurringTransactionsRequest](#api-v1-GetReCurringTransactionsRequest) | [GetReCurringTransactionsResponse](#api-v1-GetReCurringTransactionsResponse) |  |
| GetTransactions | [GetTransactionsRequest](#api-v1-GetTransactionsRequest) | [GetTransactionsResponse](#api-v1-GetTransactionsResponse) |  |
| ProcessWebhook | [ProcessWebhookRequest](#api-v1-ProcessWebhookRequest) | [ProcessWebhookResponse](#api-v1-ProcessWebhookResponse) |  |
| StripeWebhook | [StripeWebhookRequest](#api-v1-StripeWebhookRequest) | [StripeWebhookResponse](#api-v1-StripeWebhookResponse) |  |
| CreateSubscription | [CreateSubscriptionRequest](#api-v1-CreateSubscriptionRequest) | [CreateSubscriptionResponse](#api-v1-CreateSubscriptionResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

