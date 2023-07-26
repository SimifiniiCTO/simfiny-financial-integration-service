/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import {
  GetAccountBalanceHistoryRequest,
  GetAccountBalanceHistoryResponse,
  GetCategoryMonthlyTransactionCountRequest,
  GetCategoryMonthlyTransactionCountResponse,
  GetDebtToIncomeRatioRequest,
  GetDebtToIncomeRatioResponse,
  GetExpenseMetricsRequest,
  GetExpenseMetricsResponse,
  GetFinancialProfileRequest,
  GetFinancialProfileResponse,
  GetIncomeExpenseRatioRequest,
  GetIncomeExpenseRatioResponse,
  GetIncomeMetricsRequest,
  GetIncomeMetricsResponse,
  GetMelodyFinancialContextRequest,
  GetMelodyFinancialContextResponse,
  GetMerchantMonthlyExpenditureRequest,
  GetMerchantMonthlyExpenditureResponse,
  GetMonthlyBalanceRequest,
  GetMonthlyBalanceResponse,
  GetMonthlyExpenditureRequest,
  GetMonthlyExpenditureResponse,
  GetMonthlyIncomeRequest,
  GetMonthlyIncomeResponse,
  GetMonthlySavingsRequest,
  GetMonthlySavingsResponse,
  GetMonthlyTotalQuantityBySecurityAndUserRequest,
  GetMonthlyTotalQuantityBySecurityAndUserResponse,
  GetMonthlyTransactionCountRequest,
  GetMonthlyTransactionCountResponse,
  GetPaymentChannelMonthlyExpenditureRequest,
  GetPaymentChannelMonthlyExpenditureResponse,
  GetTotalInvestmentBySecurityRequest,
  GetTotalInvestmentBySecurityResponse,
  GetTransactionAggregatesRequest,
  GetTransactionAggregatesResponse,
  GetUserAccountBalanceHistoryRequest,
  GetUserAccountBalanceHistoryResponse,
  GetUserCategoryMonthlyExpenditureRequest,
  GetUserCategoryMonthlyExpenditureResponse,
  GetUserCategoryMonthlyIncomeRequest,
  GetUserCategoryMonthlyIncomeResponse,
} from "./request_response_financial_analytics_service";
import {
  CreateBankAccountRequest,
  CreateBankAccountResponse,
  CreateBudgetRequest,
  CreateBudgetResponse,
  CreateManualLinkRequest,
  CreateManualLinkResponse,
  CreateMilestoneRequest,
  CreateMilestoneResponse,
  CreateSmartGoalRequest,
  CreateSmartGoalResponse,
  CreateSubscriptionRequest,
  CreateSubscriptionResponse,
  CreateUserProfileRequest,
  CreateUserProfileResponse,
  DeleteBankAccountRequest,
  DeleteBankAccountResponse,
  DeleteBudgetRequest,
  DeleteBudgetResponse,
  DeleteLinkRequest,
  DeleteLinkResponse,
  DeleteMilestoneRequest,
  DeleteMilestoneResponse,
  DeleteSmartGoalRequest,
  DeleteSmartGoalResponse,
  DeleteUserProfileRequest,
  DeleteUserProfileResponse,
  GetAllBudgetsRequest,
  GetAllBudgetsResponse,
  GetBankAccountRequest,
  GetBankAccountResponse,
  GetBudgetRequest,
  GetBudgetResponse,
  GetForecastRequest,
  GetForecastResponse,
  GetInvestmentAcccountRequest,
  GetInvestmentAcccountResponse,
  GetLiabilityAccountRequest,
  GetLiabilityAccountResponse,
  GetLinkRequest,
  GetLinkResponse,
  GetLinksRequest,
  GetLinksResponse,
  GetMilestoneRequest,
  GetMilestoneResponse,
  GetMilestonesBySmartGoalIdRequest,
  GetMilestonesBySmartGoalIdResponse,
  GetMortgageAccountRequest,
  GetMortgageAccountResponse,
  GetPocketRequest,
  GetPocketResponse,
  GetReCurringTransactionsRequest,
  GetReCurringTransactionsResponse,
  GetSmartGoalsByPocketIdRequest,
  GetSmartGoalsByPocketIdResponse,
  GetStudentLoanAccountRequest,
  GetStudentLoanAccountResponse,
  GetTransactionsForBankAccountRequest,
  GetTransactionsForBankAccountResponse,
  GetTransactionsRequest,
  GetTransactionsResponse,
  GetUserProfileRequest,
  GetUserProfileResponse,
  HealthCheckRequest,
  HealthCheckResponse,
  PlaidExchangeTokenRequest,
  PlaidExchangeTokenResponse,
  PlaidInitiateTokenExchangeRequest,
  PlaidInitiateTokenExchangeResponse,
  PlaidInitiateTokenUpdateRequest,
  PlaidInitiateTokenUpdateResponse,
  ProcessWebhookRequest,
  ProcessWebhookResponse,
  ReadynessCheckRequest,
  ReadynessCheckResponse,
  StripeWebhookRequest,
  StripeWebhookResponse,
  UpdateBankAccountRequest,
  UpdateBankAccountResponse,
  UpdateBudgetRequest,
  UpdateBudgetResponse,
  UpdateMilestoneRequest,
  UpdateMilestoneResponse,
  UpdateSmartGoalRequest,
  UpdateSmartGoalResponse,
  UpdateUserProfileRequest,
  UpdateUserProfileResponse,
} from "./request_response_financial_service";

export const protobufPackage = "financial_integration_service_api.v1";

/** FinancialService API. */
export interface FinancialService {
  PlaidInitiateTokenExchange(request: PlaidInitiateTokenExchangeRequest): Promise<PlaidInitiateTokenExchangeResponse>;
  PlaidInitiateTokenUpdate(request: PlaidInitiateTokenUpdateRequest): Promise<PlaidInitiateTokenUpdateResponse>;
  PlaidExchangeToken(request: PlaidExchangeTokenRequest): Promise<PlaidExchangeTokenResponse>;
  CreateUserProfile(request: CreateUserProfileRequest): Promise<CreateUserProfileResponse>;
  GetUserProfile(request: GetUserProfileRequest): Promise<GetUserProfileResponse>;
  DeleteUserProfile(request: DeleteUserProfileRequest): Promise<DeleteUserProfileResponse>;
  UpdateUserProfile(request: UpdateUserProfileRequest): Promise<UpdateUserProfileResponse>;
  CreateBankAccount(request: CreateBankAccountRequest): Promise<CreateBankAccountResponse>;
  GetBankAccount(request: GetBankAccountRequest): Promise<GetBankAccountResponse>;
  UpdateBankAccount(request: UpdateBankAccountRequest): Promise<UpdateBankAccountResponse>;
  DeleteBankAccount(request: DeleteBankAccountRequest): Promise<DeleteBankAccountResponse>;
  GetPocket(request: GetPocketRequest): Promise<GetPocketResponse>;
  GetSmartGoalsByPocketId(request: GetSmartGoalsByPocketIdRequest): Promise<GetSmartGoalsByPocketIdResponse>;
  CreateSmartGoal(request: CreateSmartGoalRequest): Promise<CreateSmartGoalResponse>;
  UpdateSmartGoal(request: UpdateSmartGoalRequest): Promise<UpdateSmartGoalResponse>;
  DeleteSmartGoal(request: DeleteSmartGoalRequest): Promise<DeleteSmartGoalResponse>;
  CreateMilestone(request: CreateMilestoneRequest): Promise<CreateMilestoneResponse>;
  DeleteMilestone(request: DeleteMilestoneRequest): Promise<DeleteMilestoneResponse>;
  UpdateMilestone(request: UpdateMilestoneRequest): Promise<UpdateMilestoneResponse>;
  GetMilestone(request: GetMilestoneRequest): Promise<GetMilestoneResponse>;
  GetMilestonesBySmartGoalId(request: GetMilestonesBySmartGoalIdRequest): Promise<GetMilestonesBySmartGoalIdResponse>;
  GetForecast(request: GetForecastRequest): Promise<GetForecastResponse>;
  CreateBudget(request: CreateBudgetRequest): Promise<CreateBudgetResponse>;
  UpdateBudget(request: UpdateBudgetRequest): Promise<UpdateBudgetResponse>;
  DeleteBudget(request: DeleteBudgetRequest): Promise<DeleteBudgetResponse>;
  GetBudget(request: GetBudgetRequest): Promise<GetBudgetResponse>;
  GetAllBudgets(request: GetAllBudgetsRequest): Promise<GetAllBudgetsResponse>;
  HealthCheck(request: HealthCheckRequest): Promise<HealthCheckResponse>;
  ReadynessCheck(request: ReadynessCheckRequest): Promise<ReadynessCheckResponse>;
  GetInvestmentAcccount(request: GetInvestmentAcccountRequest): Promise<GetInvestmentAcccountResponse>;
  GetMortgageAccount(request: GetMortgageAccountRequest): Promise<GetMortgageAccountResponse>;
  GetLiabilityAccount(request: GetLiabilityAccountRequest): Promise<GetLiabilityAccountResponse>;
  GetStudentLoanAccount(request: GetStudentLoanAccountRequest): Promise<GetStudentLoanAccountResponse>;
  CreateManualLink(request: CreateManualLinkRequest): Promise<CreateManualLinkResponse>;
  GetLink(request: GetLinkRequest): Promise<GetLinkResponse>;
  GetLinks(request: GetLinksRequest): Promise<GetLinksResponse>;
  DeleteLink(request: DeleteLinkRequest): Promise<DeleteLinkResponse>;
  GetReCurringTransactions(request: GetReCurringTransactionsRequest): Promise<GetReCurringTransactionsResponse>;
  GetTransactions(request: GetTransactionsRequest): Promise<GetTransactionsResponse>;
  ProcessWebhook(request: ProcessWebhookRequest): Promise<ProcessWebhookResponse>;
  StripeWebhook(request: StripeWebhookRequest): Promise<StripeWebhookResponse>;
  CreateSubscription(request: CreateSubscriptionRequest): Promise<CreateSubscriptionResponse>;
  /** Transaction Aggregates By Queries */
  GetTransactionAggregates(request: GetTransactionAggregatesRequest): Promise<GetTransactionAggregatesResponse>;
  GetUserAccountBalanceHistory(
    request: GetUserAccountBalanceHistoryRequest,
  ): Promise<GetUserAccountBalanceHistoryResponse>;
  GetAccountBalanceHistory(request: GetAccountBalanceHistoryRequest): Promise<GetAccountBalanceHistoryResponse>;
  /** Get user category monthly expenditure */
  GetUserCategoryMonthlyExpenditure(
    request: GetUserCategoryMonthlyExpenditureRequest,
  ): Promise<GetUserCategoryMonthlyExpenditureResponse>;
  /** Get CategoryMonthlyIncome by Category and User - This would return all CategoryMonthlyIncome records for a specific user for a specific personal finance category */
  GetUserCategoryMonthlyIncome(
    request: GetUserCategoryMonthlyIncomeRequest,
  ): Promise<GetUserCategoryMonthlyIncomeResponse>;
  /** Get CategoryMonthlyTransactionCount by User - This would return all CategoryMonthlyTransactionCount records for a specific user */
  GetCategoryMonthlyTransactionCount(
    request: GetCategoryMonthlyTransactionCountRequest,
  ): Promise<GetCategoryMonthlyTransactionCountResponse>;
  GetDebtToIncomeRatio(request: GetDebtToIncomeRatioRequest): Promise<GetDebtToIncomeRatioResponse>;
  GetExpenseMetrics(request: GetExpenseMetricsRequest): Promise<GetExpenseMetricsResponse>;
  GetFinancialProfile(request: GetFinancialProfileRequest): Promise<GetFinancialProfileResponse>;
  GetIncomeExpenseRatio(request: GetIncomeExpenseRatioRequest): Promise<GetIncomeExpenseRatioResponse>;
  GetIncomeMetrics(request: GetIncomeMetricsRequest): Promise<GetIncomeMetricsResponse>;
  GetMerchantMonthlyExpenditure(
    request: GetMerchantMonthlyExpenditureRequest,
  ): Promise<GetMerchantMonthlyExpenditureResponse>;
  /** @here */
  GetMonthlyBalance(request: GetMonthlyBalanceRequest): Promise<GetMonthlyBalanceResponse>;
  GetMonthlyExpenditure(request: GetMonthlyExpenditureRequest): Promise<GetMonthlyExpenditureResponse>;
  GetMonthlyIncome(request: GetMonthlyIncomeRequest): Promise<GetMonthlyIncomeResponse>;
  GetMonthlySavings(request: GetMonthlySavingsRequest): Promise<GetMonthlySavingsResponse>;
  GetMonthlyTotalQuantityBySecurityAndUser(
    request: GetMonthlyTotalQuantityBySecurityAndUserRequest,
  ): Promise<GetMonthlyTotalQuantityBySecurityAndUserResponse>;
  GetMonthlyTransactionCount(request: GetMonthlyTransactionCountRequest): Promise<GetMonthlyTransactionCountResponse>;
  GetPaymentChannelMonthlyExpenditure(
    request: GetPaymentChannelMonthlyExpenditureRequest,
  ): Promise<GetPaymentChannelMonthlyExpenditureResponse>;
  GetTotalInvestmentBySecurity(
    request: GetTotalInvestmentBySecurityRequest,
  ): Promise<GetTotalInvestmentBySecurityResponse>;
  GetMelodyFinancialContext(request: GetMelodyFinancialContextRequest): Promise<GetMelodyFinancialContextResponse>;
  GetTransactionsForBankAccount(
    request: GetTransactionsForBankAccountRequest,
  ): Promise<GetTransactionsForBankAccountResponse>;
}

export const FinancialServiceServiceName = "financial_integration_service_api.v1.FinancialService";
export class FinancialServiceClientImpl implements FinancialService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || FinancialServiceServiceName;
    this.rpc = rpc;
    this.PlaidInitiateTokenExchange = this.PlaidInitiateTokenExchange.bind(this);
    this.PlaidInitiateTokenUpdate = this.PlaidInitiateTokenUpdate.bind(this);
    this.PlaidExchangeToken = this.PlaidExchangeToken.bind(this);
    this.CreateUserProfile = this.CreateUserProfile.bind(this);
    this.GetUserProfile = this.GetUserProfile.bind(this);
    this.DeleteUserProfile = this.DeleteUserProfile.bind(this);
    this.UpdateUserProfile = this.UpdateUserProfile.bind(this);
    this.CreateBankAccount = this.CreateBankAccount.bind(this);
    this.GetBankAccount = this.GetBankAccount.bind(this);
    this.UpdateBankAccount = this.UpdateBankAccount.bind(this);
    this.DeleteBankAccount = this.DeleteBankAccount.bind(this);
    this.GetPocket = this.GetPocket.bind(this);
    this.GetSmartGoalsByPocketId = this.GetSmartGoalsByPocketId.bind(this);
    this.CreateSmartGoal = this.CreateSmartGoal.bind(this);
    this.UpdateSmartGoal = this.UpdateSmartGoal.bind(this);
    this.DeleteSmartGoal = this.DeleteSmartGoal.bind(this);
    this.CreateMilestone = this.CreateMilestone.bind(this);
    this.DeleteMilestone = this.DeleteMilestone.bind(this);
    this.UpdateMilestone = this.UpdateMilestone.bind(this);
    this.GetMilestone = this.GetMilestone.bind(this);
    this.GetMilestonesBySmartGoalId = this.GetMilestonesBySmartGoalId.bind(this);
    this.GetForecast = this.GetForecast.bind(this);
    this.CreateBudget = this.CreateBudget.bind(this);
    this.UpdateBudget = this.UpdateBudget.bind(this);
    this.DeleteBudget = this.DeleteBudget.bind(this);
    this.GetBudget = this.GetBudget.bind(this);
    this.GetAllBudgets = this.GetAllBudgets.bind(this);
    this.HealthCheck = this.HealthCheck.bind(this);
    this.ReadynessCheck = this.ReadynessCheck.bind(this);
    this.GetInvestmentAcccount = this.GetInvestmentAcccount.bind(this);
    this.GetMortgageAccount = this.GetMortgageAccount.bind(this);
    this.GetLiabilityAccount = this.GetLiabilityAccount.bind(this);
    this.GetStudentLoanAccount = this.GetStudentLoanAccount.bind(this);
    this.CreateManualLink = this.CreateManualLink.bind(this);
    this.GetLink = this.GetLink.bind(this);
    this.GetLinks = this.GetLinks.bind(this);
    this.DeleteLink = this.DeleteLink.bind(this);
    this.GetReCurringTransactions = this.GetReCurringTransactions.bind(this);
    this.GetTransactions = this.GetTransactions.bind(this);
    this.ProcessWebhook = this.ProcessWebhook.bind(this);
    this.StripeWebhook = this.StripeWebhook.bind(this);
    this.CreateSubscription = this.CreateSubscription.bind(this);
    this.GetTransactionAggregates = this.GetTransactionAggregates.bind(this);
    this.GetUserAccountBalanceHistory = this.GetUserAccountBalanceHistory.bind(this);
    this.GetAccountBalanceHistory = this.GetAccountBalanceHistory.bind(this);
    this.GetUserCategoryMonthlyExpenditure = this.GetUserCategoryMonthlyExpenditure.bind(this);
    this.GetUserCategoryMonthlyIncome = this.GetUserCategoryMonthlyIncome.bind(this);
    this.GetCategoryMonthlyTransactionCount = this.GetCategoryMonthlyTransactionCount.bind(this);
    this.GetDebtToIncomeRatio = this.GetDebtToIncomeRatio.bind(this);
    this.GetExpenseMetrics = this.GetExpenseMetrics.bind(this);
    this.GetFinancialProfile = this.GetFinancialProfile.bind(this);
    this.GetIncomeExpenseRatio = this.GetIncomeExpenseRatio.bind(this);
    this.GetIncomeMetrics = this.GetIncomeMetrics.bind(this);
    this.GetMerchantMonthlyExpenditure = this.GetMerchantMonthlyExpenditure.bind(this);
    this.GetMonthlyBalance = this.GetMonthlyBalance.bind(this);
    this.GetMonthlyExpenditure = this.GetMonthlyExpenditure.bind(this);
    this.GetMonthlyIncome = this.GetMonthlyIncome.bind(this);
    this.GetMonthlySavings = this.GetMonthlySavings.bind(this);
    this.GetMonthlyTotalQuantityBySecurityAndUser = this.GetMonthlyTotalQuantityBySecurityAndUser.bind(this);
    this.GetMonthlyTransactionCount = this.GetMonthlyTransactionCount.bind(this);
    this.GetPaymentChannelMonthlyExpenditure = this.GetPaymentChannelMonthlyExpenditure.bind(this);
    this.GetTotalInvestmentBySecurity = this.GetTotalInvestmentBySecurity.bind(this);
    this.GetMelodyFinancialContext = this.GetMelodyFinancialContext.bind(this);
    this.GetTransactionsForBankAccount = this.GetTransactionsForBankAccount.bind(this);
  }
  PlaidInitiateTokenExchange(request: PlaidInitiateTokenExchangeRequest): Promise<PlaidInitiateTokenExchangeResponse> {
    const data = PlaidInitiateTokenExchangeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PlaidInitiateTokenExchange", data);
    return promise.then((data) => PlaidInitiateTokenExchangeResponse.decode(_m0.Reader.create(data)));
  }

  PlaidInitiateTokenUpdate(request: PlaidInitiateTokenUpdateRequest): Promise<PlaidInitiateTokenUpdateResponse> {
    const data = PlaidInitiateTokenUpdateRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PlaidInitiateTokenUpdate", data);
    return promise.then((data) => PlaidInitiateTokenUpdateResponse.decode(_m0.Reader.create(data)));
  }

  PlaidExchangeToken(request: PlaidExchangeTokenRequest): Promise<PlaidExchangeTokenResponse> {
    const data = PlaidExchangeTokenRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PlaidExchangeToken", data);
    return promise.then((data) => PlaidExchangeTokenResponse.decode(_m0.Reader.create(data)));
  }

  CreateUserProfile(request: CreateUserProfileRequest): Promise<CreateUserProfileResponse> {
    const data = CreateUserProfileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateUserProfile", data);
    return promise.then((data) => CreateUserProfileResponse.decode(_m0.Reader.create(data)));
  }

  GetUserProfile(request: GetUserProfileRequest): Promise<GetUserProfileResponse> {
    const data = GetUserProfileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserProfile", data);
    return promise.then((data) => GetUserProfileResponse.decode(_m0.Reader.create(data)));
  }

  DeleteUserProfile(request: DeleteUserProfileRequest): Promise<DeleteUserProfileResponse> {
    const data = DeleteUserProfileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteUserProfile", data);
    return promise.then((data) => DeleteUserProfileResponse.decode(_m0.Reader.create(data)));
  }

  UpdateUserProfile(request: UpdateUserProfileRequest): Promise<UpdateUserProfileResponse> {
    const data = UpdateUserProfileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateUserProfile", data);
    return promise.then((data) => UpdateUserProfileResponse.decode(_m0.Reader.create(data)));
  }

  CreateBankAccount(request: CreateBankAccountRequest): Promise<CreateBankAccountResponse> {
    const data = CreateBankAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateBankAccount", data);
    return promise.then((data) => CreateBankAccountResponse.decode(_m0.Reader.create(data)));
  }

  GetBankAccount(request: GetBankAccountRequest): Promise<GetBankAccountResponse> {
    const data = GetBankAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetBankAccount", data);
    return promise.then((data) => GetBankAccountResponse.decode(_m0.Reader.create(data)));
  }

  UpdateBankAccount(request: UpdateBankAccountRequest): Promise<UpdateBankAccountResponse> {
    const data = UpdateBankAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateBankAccount", data);
    return promise.then((data) => UpdateBankAccountResponse.decode(_m0.Reader.create(data)));
  }

  DeleteBankAccount(request: DeleteBankAccountRequest): Promise<DeleteBankAccountResponse> {
    const data = DeleteBankAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteBankAccount", data);
    return promise.then((data) => DeleteBankAccountResponse.decode(_m0.Reader.create(data)));
  }

  GetPocket(request: GetPocketRequest): Promise<GetPocketResponse> {
    const data = GetPocketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPocket", data);
    return promise.then((data) => GetPocketResponse.decode(_m0.Reader.create(data)));
  }

  GetSmartGoalsByPocketId(request: GetSmartGoalsByPocketIdRequest): Promise<GetSmartGoalsByPocketIdResponse> {
    const data = GetSmartGoalsByPocketIdRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetSmartGoalsByPocketId", data);
    return promise.then((data) => GetSmartGoalsByPocketIdResponse.decode(_m0.Reader.create(data)));
  }

  CreateSmartGoal(request: CreateSmartGoalRequest): Promise<CreateSmartGoalResponse> {
    const data = CreateSmartGoalRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateSmartGoal", data);
    return promise.then((data) => CreateSmartGoalResponse.decode(_m0.Reader.create(data)));
  }

  UpdateSmartGoal(request: UpdateSmartGoalRequest): Promise<UpdateSmartGoalResponse> {
    const data = UpdateSmartGoalRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateSmartGoal", data);
    return promise.then((data) => UpdateSmartGoalResponse.decode(_m0.Reader.create(data)));
  }

  DeleteSmartGoal(request: DeleteSmartGoalRequest): Promise<DeleteSmartGoalResponse> {
    const data = DeleteSmartGoalRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteSmartGoal", data);
    return promise.then((data) => DeleteSmartGoalResponse.decode(_m0.Reader.create(data)));
  }

  CreateMilestone(request: CreateMilestoneRequest): Promise<CreateMilestoneResponse> {
    const data = CreateMilestoneRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateMilestone", data);
    return promise.then((data) => CreateMilestoneResponse.decode(_m0.Reader.create(data)));
  }

  DeleteMilestone(request: DeleteMilestoneRequest): Promise<DeleteMilestoneResponse> {
    const data = DeleteMilestoneRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteMilestone", data);
    return promise.then((data) => DeleteMilestoneResponse.decode(_m0.Reader.create(data)));
  }

  UpdateMilestone(request: UpdateMilestoneRequest): Promise<UpdateMilestoneResponse> {
    const data = UpdateMilestoneRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateMilestone", data);
    return promise.then((data) => UpdateMilestoneResponse.decode(_m0.Reader.create(data)));
  }

  GetMilestone(request: GetMilestoneRequest): Promise<GetMilestoneResponse> {
    const data = GetMilestoneRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMilestone", data);
    return promise.then((data) => GetMilestoneResponse.decode(_m0.Reader.create(data)));
  }

  GetMilestonesBySmartGoalId(request: GetMilestonesBySmartGoalIdRequest): Promise<GetMilestonesBySmartGoalIdResponse> {
    const data = GetMilestonesBySmartGoalIdRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMilestonesBySmartGoalId", data);
    return promise.then((data) => GetMilestonesBySmartGoalIdResponse.decode(_m0.Reader.create(data)));
  }

  GetForecast(request: GetForecastRequest): Promise<GetForecastResponse> {
    const data = GetForecastRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetForecast", data);
    return promise.then((data) => GetForecastResponse.decode(_m0.Reader.create(data)));
  }

  CreateBudget(request: CreateBudgetRequest): Promise<CreateBudgetResponse> {
    const data = CreateBudgetRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateBudget", data);
    return promise.then((data) => CreateBudgetResponse.decode(_m0.Reader.create(data)));
  }

  UpdateBudget(request: UpdateBudgetRequest): Promise<UpdateBudgetResponse> {
    const data = UpdateBudgetRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateBudget", data);
    return promise.then((data) => UpdateBudgetResponse.decode(_m0.Reader.create(data)));
  }

  DeleteBudget(request: DeleteBudgetRequest): Promise<DeleteBudgetResponse> {
    const data = DeleteBudgetRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteBudget", data);
    return promise.then((data) => DeleteBudgetResponse.decode(_m0.Reader.create(data)));
  }

  GetBudget(request: GetBudgetRequest): Promise<GetBudgetResponse> {
    const data = GetBudgetRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetBudget", data);
    return promise.then((data) => GetBudgetResponse.decode(_m0.Reader.create(data)));
  }

  GetAllBudgets(request: GetAllBudgetsRequest): Promise<GetAllBudgetsResponse> {
    const data = GetAllBudgetsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetAllBudgets", data);
    return promise.then((data) => GetAllBudgetsResponse.decode(_m0.Reader.create(data)));
  }

  HealthCheck(request: HealthCheckRequest): Promise<HealthCheckResponse> {
    const data = HealthCheckRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "HealthCheck", data);
    return promise.then((data) => HealthCheckResponse.decode(_m0.Reader.create(data)));
  }

  ReadynessCheck(request: ReadynessCheckRequest): Promise<ReadynessCheckResponse> {
    const data = ReadynessCheckRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ReadynessCheck", data);
    return promise.then((data) => ReadynessCheckResponse.decode(_m0.Reader.create(data)));
  }

  GetInvestmentAcccount(request: GetInvestmentAcccountRequest): Promise<GetInvestmentAcccountResponse> {
    const data = GetInvestmentAcccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetInvestmentAcccount", data);
    return promise.then((data) => GetInvestmentAcccountResponse.decode(_m0.Reader.create(data)));
  }

  GetMortgageAccount(request: GetMortgageAccountRequest): Promise<GetMortgageAccountResponse> {
    const data = GetMortgageAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMortgageAccount", data);
    return promise.then((data) => GetMortgageAccountResponse.decode(_m0.Reader.create(data)));
  }

  GetLiabilityAccount(request: GetLiabilityAccountRequest): Promise<GetLiabilityAccountResponse> {
    const data = GetLiabilityAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetLiabilityAccount", data);
    return promise.then((data) => GetLiabilityAccountResponse.decode(_m0.Reader.create(data)));
  }

  GetStudentLoanAccount(request: GetStudentLoanAccountRequest): Promise<GetStudentLoanAccountResponse> {
    const data = GetStudentLoanAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetStudentLoanAccount", data);
    return promise.then((data) => GetStudentLoanAccountResponse.decode(_m0.Reader.create(data)));
  }

  CreateManualLink(request: CreateManualLinkRequest): Promise<CreateManualLinkResponse> {
    const data = CreateManualLinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateManualLink", data);
    return promise.then((data) => CreateManualLinkResponse.decode(_m0.Reader.create(data)));
  }

  GetLink(request: GetLinkRequest): Promise<GetLinkResponse> {
    const data = GetLinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetLink", data);
    return promise.then((data) => GetLinkResponse.decode(_m0.Reader.create(data)));
  }

  GetLinks(request: GetLinksRequest): Promise<GetLinksResponse> {
    const data = GetLinksRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetLinks", data);
    return promise.then((data) => GetLinksResponse.decode(_m0.Reader.create(data)));
  }

  DeleteLink(request: DeleteLinkRequest): Promise<DeleteLinkResponse> {
    const data = DeleteLinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteLink", data);
    return promise.then((data) => DeleteLinkResponse.decode(_m0.Reader.create(data)));
  }

  GetReCurringTransactions(request: GetReCurringTransactionsRequest): Promise<GetReCurringTransactionsResponse> {
    const data = GetReCurringTransactionsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetReCurringTransactions", data);
    return promise.then((data) => GetReCurringTransactionsResponse.decode(_m0.Reader.create(data)));
  }

  GetTransactions(request: GetTransactionsRequest): Promise<GetTransactionsResponse> {
    const data = GetTransactionsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetTransactions", data);
    return promise.then((data) => GetTransactionsResponse.decode(_m0.Reader.create(data)));
  }

  ProcessWebhook(request: ProcessWebhookRequest): Promise<ProcessWebhookResponse> {
    const data = ProcessWebhookRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ProcessWebhook", data);
    return promise.then((data) => ProcessWebhookResponse.decode(_m0.Reader.create(data)));
  }

  StripeWebhook(request: StripeWebhookRequest): Promise<StripeWebhookResponse> {
    const data = StripeWebhookRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "StripeWebhook", data);
    return promise.then((data) => StripeWebhookResponse.decode(_m0.Reader.create(data)));
  }

  CreateSubscription(request: CreateSubscriptionRequest): Promise<CreateSubscriptionResponse> {
    const data = CreateSubscriptionRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateSubscription", data);
    return promise.then((data) => CreateSubscriptionResponse.decode(_m0.Reader.create(data)));
  }

  GetTransactionAggregates(request: GetTransactionAggregatesRequest): Promise<GetTransactionAggregatesResponse> {
    const data = GetTransactionAggregatesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetTransactionAggregates", data);
    return promise.then((data) => GetTransactionAggregatesResponse.decode(_m0.Reader.create(data)));
  }

  GetUserAccountBalanceHistory(
    request: GetUserAccountBalanceHistoryRequest,
  ): Promise<GetUserAccountBalanceHistoryResponse> {
    const data = GetUserAccountBalanceHistoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserAccountBalanceHistory", data);
    return promise.then((data) => GetUserAccountBalanceHistoryResponse.decode(_m0.Reader.create(data)));
  }

  GetAccountBalanceHistory(request: GetAccountBalanceHistoryRequest): Promise<GetAccountBalanceHistoryResponse> {
    const data = GetAccountBalanceHistoryRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetAccountBalanceHistory", data);
    return promise.then((data) => GetAccountBalanceHistoryResponse.decode(_m0.Reader.create(data)));
  }

  GetUserCategoryMonthlyExpenditure(
    request: GetUserCategoryMonthlyExpenditureRequest,
  ): Promise<GetUserCategoryMonthlyExpenditureResponse> {
    const data = GetUserCategoryMonthlyExpenditureRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserCategoryMonthlyExpenditure", data);
    return promise.then((data) => GetUserCategoryMonthlyExpenditureResponse.decode(_m0.Reader.create(data)));
  }

  GetUserCategoryMonthlyIncome(
    request: GetUserCategoryMonthlyIncomeRequest,
  ): Promise<GetUserCategoryMonthlyIncomeResponse> {
    const data = GetUserCategoryMonthlyIncomeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserCategoryMonthlyIncome", data);
    return promise.then((data) => GetUserCategoryMonthlyIncomeResponse.decode(_m0.Reader.create(data)));
  }

  GetCategoryMonthlyTransactionCount(
    request: GetCategoryMonthlyTransactionCountRequest,
  ): Promise<GetCategoryMonthlyTransactionCountResponse> {
    const data = GetCategoryMonthlyTransactionCountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetCategoryMonthlyTransactionCount", data);
    return promise.then((data) => GetCategoryMonthlyTransactionCountResponse.decode(_m0.Reader.create(data)));
  }

  GetDebtToIncomeRatio(request: GetDebtToIncomeRatioRequest): Promise<GetDebtToIncomeRatioResponse> {
    const data = GetDebtToIncomeRatioRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetDebtToIncomeRatio", data);
    return promise.then((data) => GetDebtToIncomeRatioResponse.decode(_m0.Reader.create(data)));
  }

  GetExpenseMetrics(request: GetExpenseMetricsRequest): Promise<GetExpenseMetricsResponse> {
    const data = GetExpenseMetricsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetExpenseMetrics", data);
    return promise.then((data) => GetExpenseMetricsResponse.decode(_m0.Reader.create(data)));
  }

  GetFinancialProfile(request: GetFinancialProfileRequest): Promise<GetFinancialProfileResponse> {
    const data = GetFinancialProfileRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetFinancialProfile", data);
    return promise.then((data) => GetFinancialProfileResponse.decode(_m0.Reader.create(data)));
  }

  GetIncomeExpenseRatio(request: GetIncomeExpenseRatioRequest): Promise<GetIncomeExpenseRatioResponse> {
    const data = GetIncomeExpenseRatioRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetIncomeExpenseRatio", data);
    return promise.then((data) => GetIncomeExpenseRatioResponse.decode(_m0.Reader.create(data)));
  }

  GetIncomeMetrics(request: GetIncomeMetricsRequest): Promise<GetIncomeMetricsResponse> {
    const data = GetIncomeMetricsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetIncomeMetrics", data);
    return promise.then((data) => GetIncomeMetricsResponse.decode(_m0.Reader.create(data)));
  }

  GetMerchantMonthlyExpenditure(
    request: GetMerchantMonthlyExpenditureRequest,
  ): Promise<GetMerchantMonthlyExpenditureResponse> {
    const data = GetMerchantMonthlyExpenditureRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMerchantMonthlyExpenditure", data);
    return promise.then((data) => GetMerchantMonthlyExpenditureResponse.decode(_m0.Reader.create(data)));
  }

  GetMonthlyBalance(request: GetMonthlyBalanceRequest): Promise<GetMonthlyBalanceResponse> {
    const data = GetMonthlyBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMonthlyBalance", data);
    return promise.then((data) => GetMonthlyBalanceResponse.decode(_m0.Reader.create(data)));
  }

  GetMonthlyExpenditure(request: GetMonthlyExpenditureRequest): Promise<GetMonthlyExpenditureResponse> {
    const data = GetMonthlyExpenditureRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMonthlyExpenditure", data);
    return promise.then((data) => GetMonthlyExpenditureResponse.decode(_m0.Reader.create(data)));
  }

  GetMonthlyIncome(request: GetMonthlyIncomeRequest): Promise<GetMonthlyIncomeResponse> {
    const data = GetMonthlyIncomeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMonthlyIncome", data);
    return promise.then((data) => GetMonthlyIncomeResponse.decode(_m0.Reader.create(data)));
  }

  GetMonthlySavings(request: GetMonthlySavingsRequest): Promise<GetMonthlySavingsResponse> {
    const data = GetMonthlySavingsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMonthlySavings", data);
    return promise.then((data) => GetMonthlySavingsResponse.decode(_m0.Reader.create(data)));
  }

  GetMonthlyTotalQuantityBySecurityAndUser(
    request: GetMonthlyTotalQuantityBySecurityAndUserRequest,
  ): Promise<GetMonthlyTotalQuantityBySecurityAndUserResponse> {
    const data = GetMonthlyTotalQuantityBySecurityAndUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMonthlyTotalQuantityBySecurityAndUser", data);
    return promise.then((data) => GetMonthlyTotalQuantityBySecurityAndUserResponse.decode(_m0.Reader.create(data)));
  }

  GetMonthlyTransactionCount(request: GetMonthlyTransactionCountRequest): Promise<GetMonthlyTransactionCountResponse> {
    const data = GetMonthlyTransactionCountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMonthlyTransactionCount", data);
    return promise.then((data) => GetMonthlyTransactionCountResponse.decode(_m0.Reader.create(data)));
  }

  GetPaymentChannelMonthlyExpenditure(
    request: GetPaymentChannelMonthlyExpenditureRequest,
  ): Promise<GetPaymentChannelMonthlyExpenditureResponse> {
    const data = GetPaymentChannelMonthlyExpenditureRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPaymentChannelMonthlyExpenditure", data);
    return promise.then((data) => GetPaymentChannelMonthlyExpenditureResponse.decode(_m0.Reader.create(data)));
  }

  GetTotalInvestmentBySecurity(
    request: GetTotalInvestmentBySecurityRequest,
  ): Promise<GetTotalInvestmentBySecurityResponse> {
    const data = GetTotalInvestmentBySecurityRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetTotalInvestmentBySecurity", data);
    return promise.then((data) => GetTotalInvestmentBySecurityResponse.decode(_m0.Reader.create(data)));
  }

  GetMelodyFinancialContext(request: GetMelodyFinancialContextRequest): Promise<GetMelodyFinancialContextResponse> {
    const data = GetMelodyFinancialContextRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMelodyFinancialContext", data);
    return promise.then((data) => GetMelodyFinancialContextResponse.decode(_m0.Reader.create(data)));
  }

  GetTransactionsForBankAccount(
    request: GetTransactionsForBankAccountRequest,
  ): Promise<GetTransactionsForBankAccountResponse> {
    const data = GetTransactionsForBankAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetTransactionsForBankAccount", data);
    return promise.then((data) => GetTransactionsForBankAccountResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
