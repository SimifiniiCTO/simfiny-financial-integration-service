/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
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
  GetSmartGoalsByPocketIdRequest,
  GetSmartGoalsByPocketIdResponse,
  GetStudentLoanAccountRequest,
  GetStudentLoanAccountResponse,
  GetUserProfileRequest,
  GetUserProfileResponse,
  HealthCheckRequest,
  HealthCheckResponse,
  PlaidExchangeTokenRequest,
  PlaidExchangeTokenResponse,
  PlaidInitiateTokenExchangeRequest,
  PlaidInitiateTokenExchangeResponse,
  ProcessWebhookRequest,
  ProcessWebhookResponse,
  ReadynessCheckRequest,
  ReadynessCheckResponse,
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
} from "./request_response";

export const protobufPackage = "api.v1";

/** FinancialService API. */
export interface FinancialService {
  PlaidInitiateTokenExchange(request: PlaidInitiateTokenExchangeRequest): Promise<PlaidInitiateTokenExchangeResponse>;
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
  ProcessWebhook(request: ProcessWebhookRequest): Promise<ProcessWebhookResponse>;
}

export class FinancialServiceClientImpl implements FinancialService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "api.v1.FinancialService";
    this.rpc = rpc;
    this.PlaidInitiateTokenExchange = this.PlaidInitiateTokenExchange.bind(this);
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
    this.ProcessWebhook = this.ProcessWebhook.bind(this);
  }
  PlaidInitiateTokenExchange(request: PlaidInitiateTokenExchangeRequest): Promise<PlaidInitiateTokenExchangeResponse> {
    const data = PlaidInitiateTokenExchangeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PlaidInitiateTokenExchange", data);
    return promise.then((data) => PlaidInitiateTokenExchangeResponse.decode(_m0.Reader.create(data)));
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

  ProcessWebhook(request: ProcessWebhookRequest): Promise<ProcessWebhookResponse> {
    const data = ProcessWebhookRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ProcessWebhook", data);
    return promise.then((data) => ProcessWebhookResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
