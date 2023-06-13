/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Any } from "../../google/protobuf/any";
import { ReOccuringTransaction, Transaction } from "./clickhouse";
import {
  BankAccount,
  Budget,
  CreditAccount,
  Forecast,
  InvestmentAccount,
  Link,
  Milestone,
  MortgageAccount,
  Pocket,
  SmartGoal,
  StudentLoanAccount,
  UserProfile,
} from "./message";

export const protobufPackage = "api.v1";

/**
 * CreateUserProfileRequest: Represents the request object invoked against the user
 * service to create a user profile
 */
export interface CreateUserProfileRequest {
  /**
   * User profile to create
   * Validations:
   * - cannot be nil hence required
   */
  profile:
    | UserProfile
    | undefined;
  /** the email of the account to create */
  email: string;
}

/**
 * CreateUserProfileResponse: Represents the response object returned as a response to
 * the `create user profile` request
 */
export interface CreateUserProfileResponse {
  userId: number;
}

/**
 * GetUserProfileRequest: Represents the request object invoked against the user
 * service to get a user profile
 */
export interface GetUserProfileRequest {
  /**
   * The account ID associated with the user.
   * NOTE: This user_id is the simfiny backend platform wide user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
}

/**
 * GetUserProfileResponse: Represents the response object returned as a response to
 * the `get user profile` request
 */
export interface GetUserProfileResponse {
  profile: UserProfile | undefined;
}

/**
 * teUserProfileRequest: Represents the request object invoked against the user
 * service to delete a user profile
 */
export interface DeleteUserProfileRequest {
  /**
   * The account ID associated with the user.
   * NOTE: This user_id is the simfiny backend platform wide user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
}

/**
 * DeleteUserProfileResponse: Represents the response object returned as a response to
 * the `delete user profile` request
 */
export interface DeleteUserProfileResponse {
  profileDeleted: boolean;
}

/**
 * UpdateUserProfileRequest: Represents the request object invoked against the user
 * service to update a user profile
 */
export interface UpdateUserProfileRequest {
  /**
   * User profile to update
   * Validation:
   * - cannot nil hence required
   */
  profile: UserProfile | undefined;
}

/**
 * UpdateUserProfileResponse: Represents the response object returned as a response to
 * the `update user profile` request
 */
export interface UpdateUserProfileResponse {
  profileUpdated: boolean;
  profile: UserProfile | undefined;
}

/**
 * CreateBankAccountRequest: Represents the request object invoked against the financial
 * service to create a bank account for a given user
 */
export interface CreateBankAccountRequest {
  /**
   * The account ID associated with the user
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The bank account to create
   * Validations:
   * - cannot be nil hence required
   */
  bankAccount: BankAccount | undefined;
}

/**
 * CreateBankAccountResponse: Represents the response object returned as a response to
 * the `create bank account` request
 */
export interface CreateBankAccountResponse {
  /** The bank account id */
  bankAccountId: number;
}

/**
 * GetBankAccountRequest: Represents the request object invoked against the financial
 * service to get a bank account for a given user and bank account id
 */
export interface GetBankAccountRequest {
  /**
   * The bank account id
   * Validations:
   * - bank_account_id must be greater than 0
   */
  bankAccountId: number;
}

/**
 * GetBankAccountResponse: Represents the response object returned as a response to
 * the `get bank account` request
 */
export interface GetBankAccountResponse {
  /** The bank account */
  bankAccount: BankAccount | undefined;
}

export interface DeleteBankAccountRequest {
  /**
   * The account ID associated with the user
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The bank account id
   * Validations:
   * - bank_account_id must be greater than 0
   */
  bankAccountId: number;
}

export interface DeleteBankAccountResponse {
  /** The bank account id */
  deleted: boolean;
}

export interface UpdateBankAccountRequest {
  /**
   * The bank account to update
   * Validations:
   * - cannot be nil hence required
   */
  bankAccount: BankAccount | undefined;
}

export interface UpdateBankAccountResponse {
  /** The bank account id */
  updated: boolean;
  /** The bank account */
  bankAccount: BankAccount | undefined;
}

export interface GetPocketRequest {
  /**
   * The pocket account id
   * Validations:
   * - pocket_account_id must be greater than 0
   */
  pocketId: number;
}

export interface GetPocketResponse {
  /** The pocket account */
  pocket: Pocket | undefined;
}

export interface GetSmartGoalsByPocketIdRequest {
  /**
   * The pocket account id
   * Validations:
   * - pocket_account_id must be greater than 0
   */
  pocketId: number;
}

export interface GetSmartGoalsByPocketIdResponse {
  /** The smart goals */
  smartGoals: SmartGoal[];
}

export interface CreateSmartGoalRequest {
  /**
   * The pocket account id
   * Validations:
   * - pocket_account_id must be greater than 0
   */
  pocketId: number;
  /**
   * The smart goal to create
   * Validations:
   * - cannot be nil hence required
   */
  smartGoal: SmartGoal | undefined;
}

export interface CreateSmartGoalResponse {
  /** The smart goal id */
  smartGoalId: number;
}

export interface UpdateSmartGoalRequest {
  /**
   * The smart goal to update
   * Validations:
   * - cannot be nil hence required
   */
  smartGoal: SmartGoal | undefined;
}

export interface UpdateSmartGoalResponse {
  /** The smart goal id */
  smartGoalId: number;
}

export interface DeleteSmartGoalRequest {
  /**
   * The smart goal id
   * Validations:
   * - smart_goal_id must be greater than 0
   */
  smartGoalId: number;
}

export interface DeleteSmartGoalResponse {
  /** The smart goal id */
  deleted: boolean;
}

export interface CreateMilestoneRequest {
  /**
   * The smart goal id
   * Validations:
   * - smart_goal_id must be greater than 0
   */
  smartGoalId: number;
  /**
   * The milestone to create
   * Validations:
   * - cannot be nil hence required
   */
  milestone: Milestone | undefined;
}

export interface CreateMilestoneResponse {
  /** The milestone id */
  milestoneId: number;
}

export interface DeleteMilestoneRequest {
  /**
   * The milestone id
   * Validations:
   * - milestone_id must be greater than 0
   */
  milestoneId: number;
}

export interface DeleteMilestoneResponse {
  /** The milestone id */
  deleted: boolean;
}

export interface UpdateMilestoneRequest {
  /**
   * The milestone to update
   * Validations:
   * - cannot be nil hence required
   */
  milestone: Milestone | undefined;
}

export interface UpdateMilestoneResponse {
  /** The milestone id */
  milestone: Milestone | undefined;
}

export interface GetMilestonesBySmartGoalIdRequest {
  /**
   * The smart goal id
   * Validations:
   * - smart_goal_id must be greater than 0
   */
  smartGoalId: number;
}

export interface GetMilestonesBySmartGoalIdResponse {
  /** The milestones */
  milestones: Milestone[];
}

export interface GetMilestoneRequest {
  /**
   * The milestone id
   * Validations:
   * - milestone_id must be greater than 0
   */
  milestoneId: number;
}

export interface GetMilestoneResponse {
  /** The milestone */
  milestone: Milestone | undefined;
}

export interface GetForecastRequest {
  /**
   * The smart goal id
   * Validations:
   * - smart_goal_id must be greater than 0
   */
  smartGoalId: number;
}

export interface GetForecastResponse {
  /** The forecast */
  forecast: Forecast | undefined;
}

export interface CreateBudgetRequest {
  /** The milestone to associate this budget with */
  milestroneId: number;
  /**
   * The budget to create
   * Validations:
   * - cannot be nil hence required
   */
  budget: Budget | undefined;
}

export interface CreateBudgetResponse {
  /** The budget id */
  budgetId: number;
}

export interface UpdateBudgetRequest {
  /**
   * The budget to update
   * Validations:
   * - cannot be nil hence required
   */
  budget: Budget | undefined;
}

export interface UpdateBudgetResponse {
  /** The budget id */
  budget: Budget | undefined;
}

export interface DeleteBudgetRequest {
  /**
   * The budget id
   * Validations:
   * - budget_id must be greater than 0
   */
  budgetId: number;
}

export interface DeleteBudgetResponse {
  /** The budget id */
  deleted: boolean;
}

export interface GetBudgetRequest {
  /**
   * The budget id
   * Validations:
   * - budget_id must be greater than 0
   */
  budgetId: number;
}

export interface GetBudgetResponse {
  /** The budget */
  budget: Budget | undefined;
}

export interface GetAllBudgetsRequest {
  /**
   * The pocket account id
   * Validations:
   * - pocket_account_id must be greater than 0
   */
  pocketId: number;
  /**
   * The smart goal id
   * Validations:
   * - smart_goal_id must be greater than 0
   */
  smartGoalId: number;
  /**
   * The milestone id
   * Validations:
   * - milestone_id must be greater than 0
   */
  milestoneId: number;
}

export interface GetAllBudgetsResponse {
  /** The budgets */
  budgets: Budget[];
}

export interface HealthCheckRequest {
}

export interface HealthCheckResponse {
  healthy: boolean;
}

export interface ReadynessCheckRequest {
}

export interface ReadynessCheckResponse {
  healthy: boolean;
}

export interface PlaidInitiateTokenExchangeRequest {
  /**
   * A unique ID representing the end user. Typically this will be a user ID number from your application.
   * Personally identifiable information, such as an email address or phone number,
   * should not be used in the `client_user_id`. It is currently used as a means of searching logs
   * for the given user in the Plaid Dashboard.
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The user's full legal name. This is an optional field used in
   * the [returning user experience](https://plaid.com/docs/link/returning-user) to associate Items to the user.
   */
  fullName: string;
  /**
   * The user's email address. This field is optional, but required to enable the
   * [pre-authenticated returning user flow](https://plaid.com/docs/link/returning-user/#enabling-the-returning-user-experience).
   */
  email: string;
  /**
   * The user's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
   * This field is optional, but required to enable the [returning user experience](https://plaid.com/docs/link/returning-user).
   */
  phoneNumber: string;
}

export interface PlaidInitiateTokenExchangeResponse {
  linkToken: string;
  expiration: string;
  plaidRequestId: string;
}

export interface PlaidExchangeTokenRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The public token
   * Validations:
   * - cannot be nil hence required
   */
  publicToken: string;
  /** The institution id */
  institutionId: string;
  /** The institution name */
  institutionName: string;
}

export interface PlaidExchangeTokenResponse {
  /** wether the operation was successful */
  success: boolean;
}

export interface GetInvestmentAcccountRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The investment account id
   * Validations:
   * - investment_account_id must be greater than 0
   */
  investmentAccountId: number;
}

export interface GetInvestmentAcccountResponse {
  /** The investment account */
  investmentAccount: InvestmentAccount | undefined;
}

export interface GetMortgageAccountRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The mortage account id
   * Validations:
   * - mortage_account_id must be greater than 0
   */
  mortgageAccountId: number;
}

export interface GetMortgageAccountResponse {
  /** The mortage account */
  mortageAccount: MortgageAccount | undefined;
}

export interface GetLiabilityAccountRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The liability account id
   * Validations:
   * - liability_account_id must be greater than 0
   */
  liabilityAccountId: number;
}

export interface GetLiabilityAccountResponse {
  /** The liability account */
  liabilityAccount: CreditAccount | undefined;
}

export interface GetStudentLoanAccountRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The student loan account id
   * Validations:
   * - student_loan_account_id must be greater than 0
   */
  studentLoanAccountId: number;
}

export interface GetStudentLoanAccountResponse {
  /** The student loan account */
  studentLoanAccount: StudentLoanAccount | undefined;
}

export interface CreateManualLinkRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /** The manual account link */
  manualAccountLink: Link | undefined;
}

export interface CreateManualLinkResponse {
  /** The link's id */
  linkId: number;
}

export interface GetLinkRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The link id
   * Validations:
   * - link_id must be greater than 0
   */
  linkId: number;
}

export interface GetLinkResponse {
  /** The link */
  link: Link | undefined;
}

export interface GetLinksRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
}

export interface GetLinksResponse {
  /** The links */
  links: Link[];
}

export interface DeleteLinkRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  /**
   * The link id
   * Validations:
   * - link_id must be greater than 0
   */
  linkId: number;
}

export interface DeleteLinkResponse {
  /** The link's id */
  linkId: number;
}

export interface GetReCurringTransactionsRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
}

export interface GetReCurringTransactionsResponse {
  /** The re-occuring transactions */
  reCcuringTransactions: ReOccuringTransaction[];
  participantReCcuringTransactions: GetReCurringTransactionsResponse_ParticipantReCurringTransactions[];
}

export interface GetReCurringTransactionsResponse_ParticipantReCurringTransactions {
  /** The participant id */
  reocurringTransactionId: number;
  /** The transactions */
  transactions: Transaction[];
}

export interface GetTransactionsRequest {
  /**
   * The user id
   * Validations:
   * - user_id must be greater than 0
   */
  userId: number;
  pageNumber: number;
  pageSize: number;
}

export interface GetTransactionsResponse {
  /** The transactions */
  transactions: Transaction[];
  nextPageNumber: number;
}

export interface ProcessWebhookRequest {
  webhookType: string;
  webhookCode: string;
  /** The item_id of the Item associated with this webhook, warning, or error */
  itemId: string;
  /** Indicates if initial pull information is available. */
  initialUpdateComplete: boolean;
  /** Indicates if historical pull information is available. */
  historicalUpdateComplete: string;
  /** The Plaid environment the webhook was sent from */
  environment: string;
  /** The number of new, unfetched transactions available */
  newTransactions: string[];
  /** An array of transaction_ids corresponding to the removed transactions */
  removedTransactions: string[];
  /**
   * We use standard HTTP response codes for success and failure notifications,
   * and our errors are further classified by error_type. In general, 200 HTTP codes
   * correspond to success, 40X codes are for developer- or user-related failures, and
   * 50X codes are for Plaid-related issues. An Item with a non-null error object will
   * only be part of an API response when calling /item/get to view Item status. Otherwise,
   * error fields will be null if no error has occurred; if an error has occurred, an error
   * code will be returned instead.
   */
  error: { [key: string]: Any };
  /** A list of account_ids for accounts that have new or updated recurring transactions data. */
  accountIds: string[];
  /** The time at which the user's access_token will expire. This field will only be present */
  consentExpirationTime: string;
  /** An array of account_id's for accounts that contain new liabilities.' */
  accountIdsWithNewLiabilities: string[];
  /** An object with keys of account_id's that are mapped to their respective liabilities fields that changed. */
  accountIdsWithUpdatedLiabilities: string[];
  /** The number of new holdings reported since the last time this webhook was fired. */
  newHoldings: number;
  /**
   * The number of updated holdings reported since the last time this webhook was fired.
   * @gotag: json:"updated_holdings"
   */
  updatedHoldings: number;
}

export interface ProcessWebhookRequest_ErrorEntry {
  key: string;
  value: Any | undefined;
}

export interface ProcessWebhookResponse {
}

export interface StripeWebhookRequest {
  body: string;
  signature: string;
}

export interface StripeWebhookResponse {
  message: string;
}

function createBaseCreateUserProfileRequest(): CreateUserProfileRequest {
  return { profile: undefined, email: "" };
}

export const CreateUserProfileRequest = {
  encode(message: CreateUserProfileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.profile !== undefined) {
      UserProfile.encode(message.profile, writer.uint32(10).fork()).ldelim();
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserProfileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateUserProfileRequest {
    return {
      profile: isSet(object.profile) ? UserProfile.fromJSON(object.profile) : undefined,
      email: isSet(object.email) ? String(object.email) : "",
    };
  },

  toJSON(message: CreateUserProfileRequest): unknown {
    const obj: any = {};
    message.profile !== undefined && (obj.profile = message.profile ? UserProfile.toJSON(message.profile) : undefined);
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateUserProfileRequest>, I>>(base?: I): CreateUserProfileRequest {
    return CreateUserProfileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserProfileRequest>, I>>(object: I): CreateUserProfileRequest {
    const message = createBaseCreateUserProfileRequest();
    message.profile = (object.profile !== undefined && object.profile !== null)
      ? UserProfile.fromPartial(object.profile)
      : undefined;
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseCreateUserProfileResponse(): CreateUserProfileResponse {
  return { userId: 0 };
}

export const CreateUserProfileResponse = {
  encode(message: CreateUserProfileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserProfileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserProfileResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateUserProfileResponse {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: CreateUserProfileResponse): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateUserProfileResponse>, I>>(base?: I): CreateUserProfileResponse {
    return CreateUserProfileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserProfileResponse>, I>>(object: I): CreateUserProfileResponse {
    const message = createBaseCreateUserProfileResponse();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseGetUserProfileRequest(): GetUserProfileRequest {
  return { userId: 0 };
}

export const GetUserProfileRequest = {
  encode(message: GetUserProfileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserProfileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserProfileRequest {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: GetUserProfileRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserProfileRequest>, I>>(base?: I): GetUserProfileRequest {
    return GetUserProfileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserProfileRequest>, I>>(object: I): GetUserProfileRequest {
    const message = createBaseGetUserProfileRequest();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseGetUserProfileResponse(): GetUserProfileResponse {
  return { profile: undefined };
}

export const GetUserProfileResponse = {
  encode(message: GetUserProfileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.profile !== undefined) {
      UserProfile.encode(message.profile, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserProfileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserProfileResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserProfileResponse {
    return { profile: isSet(object.profile) ? UserProfile.fromJSON(object.profile) : undefined };
  },

  toJSON(message: GetUserProfileResponse): unknown {
    const obj: any = {};
    message.profile !== undefined && (obj.profile = message.profile ? UserProfile.toJSON(message.profile) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserProfileResponse>, I>>(base?: I): GetUserProfileResponse {
    return GetUserProfileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserProfileResponse>, I>>(object: I): GetUserProfileResponse {
    const message = createBaseGetUserProfileResponse();
    message.profile = (object.profile !== undefined && object.profile !== null)
      ? UserProfile.fromPartial(object.profile)
      : undefined;
    return message;
  },
};

function createBaseDeleteUserProfileRequest(): DeleteUserProfileRequest {
  return { userId: 0 };
}

export const DeleteUserProfileRequest = {
  encode(message: DeleteUserProfileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteUserProfileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteUserProfileRequest {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: DeleteUserProfileRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteUserProfileRequest>, I>>(base?: I): DeleteUserProfileRequest {
    return DeleteUserProfileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteUserProfileRequest>, I>>(object: I): DeleteUserProfileRequest {
    const message = createBaseDeleteUserProfileRequest();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseDeleteUserProfileResponse(): DeleteUserProfileResponse {
  return { profileDeleted: false };
}

export const DeleteUserProfileResponse = {
  encode(message: DeleteUserProfileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.profileDeleted === true) {
      writer.uint32(8).bool(message.profileDeleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteUserProfileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserProfileResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.profileDeleted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteUserProfileResponse {
    return { profileDeleted: isSet(object.profileDeleted) ? Boolean(object.profileDeleted) : false };
  },

  toJSON(message: DeleteUserProfileResponse): unknown {
    const obj: any = {};
    message.profileDeleted !== undefined && (obj.profileDeleted = message.profileDeleted);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteUserProfileResponse>, I>>(base?: I): DeleteUserProfileResponse {
    return DeleteUserProfileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteUserProfileResponse>, I>>(object: I): DeleteUserProfileResponse {
    const message = createBaseDeleteUserProfileResponse();
    message.profileDeleted = object.profileDeleted ?? false;
    return message;
  },
};

function createBaseUpdateUserProfileRequest(): UpdateUserProfileRequest {
  return { profile: undefined };
}

export const UpdateUserProfileRequest = {
  encode(message: UpdateUserProfileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.profile !== undefined) {
      UserProfile.encode(message.profile, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserProfileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateUserProfileRequest {
    return { profile: isSet(object.profile) ? UserProfile.fromJSON(object.profile) : undefined };
  },

  toJSON(message: UpdateUserProfileRequest): unknown {
    const obj: any = {};
    message.profile !== undefined && (obj.profile = message.profile ? UserProfile.toJSON(message.profile) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateUserProfileRequest>, I>>(base?: I): UpdateUserProfileRequest {
    return UpdateUserProfileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserProfileRequest>, I>>(object: I): UpdateUserProfileRequest {
    const message = createBaseUpdateUserProfileRequest();
    message.profile = (object.profile !== undefined && object.profile !== null)
      ? UserProfile.fromPartial(object.profile)
      : undefined;
    return message;
  },
};

function createBaseUpdateUserProfileResponse(): UpdateUserProfileResponse {
  return { profileUpdated: false, profile: undefined };
}

export const UpdateUserProfileResponse = {
  encode(message: UpdateUserProfileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.profileUpdated === true) {
      writer.uint32(8).bool(message.profileUpdated);
    }
    if (message.profile !== undefined) {
      UserProfile.encode(message.profile, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserProfileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserProfileResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.profileUpdated = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateUserProfileResponse {
    return {
      profileUpdated: isSet(object.profileUpdated) ? Boolean(object.profileUpdated) : false,
      profile: isSet(object.profile) ? UserProfile.fromJSON(object.profile) : undefined,
    };
  },

  toJSON(message: UpdateUserProfileResponse): unknown {
    const obj: any = {};
    message.profileUpdated !== undefined && (obj.profileUpdated = message.profileUpdated);
    message.profile !== undefined && (obj.profile = message.profile ? UserProfile.toJSON(message.profile) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateUserProfileResponse>, I>>(base?: I): UpdateUserProfileResponse {
    return UpdateUserProfileResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserProfileResponse>, I>>(object: I): UpdateUserProfileResponse {
    const message = createBaseUpdateUserProfileResponse();
    message.profileUpdated = object.profileUpdated ?? false;
    message.profile = (object.profile !== undefined && object.profile !== null)
      ? UserProfile.fromPartial(object.profile)
      : undefined;
    return message;
  },
};

function createBaseCreateBankAccountRequest(): CreateBankAccountRequest {
  return { userId: 0, bankAccount: undefined };
}

export const CreateBankAccountRequest = {
  encode(message: CreateBankAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.bankAccount !== undefined) {
      BankAccount.encode(message.bankAccount, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBankAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBankAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateBankAccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      bankAccount: isSet(object.bankAccount) ? BankAccount.fromJSON(object.bankAccount) : undefined,
    };
  },

  toJSON(message: CreateBankAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.bankAccount !== undefined &&
      (obj.bankAccount = message.bankAccount ? BankAccount.toJSON(message.bankAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBankAccountRequest>, I>>(base?: I): CreateBankAccountRequest {
    return CreateBankAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateBankAccountRequest>, I>>(object: I): CreateBankAccountRequest {
    const message = createBaseCreateBankAccountRequest();
    message.userId = object.userId ?? 0;
    message.bankAccount = (object.bankAccount !== undefined && object.bankAccount !== null)
      ? BankAccount.fromPartial(object.bankAccount)
      : undefined;
    return message;
  },
};

function createBaseCreateBankAccountResponse(): CreateBankAccountResponse {
  return { bankAccountId: 0 };
}

export const CreateBankAccountResponse = {
  encode(message: CreateBankAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bankAccountId !== 0) {
      writer.uint32(8).uint64(message.bankAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBankAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBankAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.bankAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateBankAccountResponse {
    return { bankAccountId: isSet(object.bankAccountId) ? Number(object.bankAccountId) : 0 };
  },

  toJSON(message: CreateBankAccountResponse): unknown {
    const obj: any = {};
    message.bankAccountId !== undefined && (obj.bankAccountId = Math.round(message.bankAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBankAccountResponse>, I>>(base?: I): CreateBankAccountResponse {
    return CreateBankAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateBankAccountResponse>, I>>(object: I): CreateBankAccountResponse {
    const message = createBaseCreateBankAccountResponse();
    message.bankAccountId = object.bankAccountId ?? 0;
    return message;
  },
};

function createBaseGetBankAccountRequest(): GetBankAccountRequest {
  return { bankAccountId: 0 };
}

export const GetBankAccountRequest = {
  encode(message: GetBankAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bankAccountId !== 0) {
      writer.uint32(8).uint64(message.bankAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBankAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBankAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.bankAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetBankAccountRequest {
    return { bankAccountId: isSet(object.bankAccountId) ? Number(object.bankAccountId) : 0 };
  },

  toJSON(message: GetBankAccountRequest): unknown {
    const obj: any = {};
    message.bankAccountId !== undefined && (obj.bankAccountId = Math.round(message.bankAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetBankAccountRequest>, I>>(base?: I): GetBankAccountRequest {
    return GetBankAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetBankAccountRequest>, I>>(object: I): GetBankAccountRequest {
    const message = createBaseGetBankAccountRequest();
    message.bankAccountId = object.bankAccountId ?? 0;
    return message;
  },
};

function createBaseGetBankAccountResponse(): GetBankAccountResponse {
  return { bankAccount: undefined };
}

export const GetBankAccountResponse = {
  encode(message: GetBankAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bankAccount !== undefined) {
      BankAccount.encode(message.bankAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBankAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBankAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetBankAccountResponse {
    return { bankAccount: isSet(object.bankAccount) ? BankAccount.fromJSON(object.bankAccount) : undefined };
  },

  toJSON(message: GetBankAccountResponse): unknown {
    const obj: any = {};
    message.bankAccount !== undefined &&
      (obj.bankAccount = message.bankAccount ? BankAccount.toJSON(message.bankAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetBankAccountResponse>, I>>(base?: I): GetBankAccountResponse {
    return GetBankAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetBankAccountResponse>, I>>(object: I): GetBankAccountResponse {
    const message = createBaseGetBankAccountResponse();
    message.bankAccount = (object.bankAccount !== undefined && object.bankAccount !== null)
      ? BankAccount.fromPartial(object.bankAccount)
      : undefined;
    return message;
  },
};

function createBaseDeleteBankAccountRequest(): DeleteBankAccountRequest {
  return { userId: 0, bankAccountId: 0 };
}

export const DeleteBankAccountRequest = {
  encode(message: DeleteBankAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.bankAccountId !== 0) {
      writer.uint32(16).uint64(message.bankAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBankAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBankAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.bankAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteBankAccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      bankAccountId: isSet(object.bankAccountId) ? Number(object.bankAccountId) : 0,
    };
  },

  toJSON(message: DeleteBankAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.bankAccountId !== undefined && (obj.bankAccountId = Math.round(message.bankAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteBankAccountRequest>, I>>(base?: I): DeleteBankAccountRequest {
    return DeleteBankAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBankAccountRequest>, I>>(object: I): DeleteBankAccountRequest {
    const message = createBaseDeleteBankAccountRequest();
    message.userId = object.userId ?? 0;
    message.bankAccountId = object.bankAccountId ?? 0;
    return message;
  },
};

function createBaseDeleteBankAccountResponse(): DeleteBankAccountResponse {
  return { deleted: false };
}

export const DeleteBankAccountResponse = {
  encode(message: DeleteBankAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deleted === true) {
      writer.uint32(8).bool(message.deleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBankAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBankAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteBankAccountResponse {
    return { deleted: isSet(object.deleted) ? Boolean(object.deleted) : false };
  },

  toJSON(message: DeleteBankAccountResponse): unknown {
    const obj: any = {};
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteBankAccountResponse>, I>>(base?: I): DeleteBankAccountResponse {
    return DeleteBankAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBankAccountResponse>, I>>(object: I): DeleteBankAccountResponse {
    const message = createBaseDeleteBankAccountResponse();
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseUpdateBankAccountRequest(): UpdateBankAccountRequest {
  return { bankAccount: undefined };
}

export const UpdateBankAccountRequest = {
  encode(message: UpdateBankAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bankAccount !== undefined) {
      BankAccount.encode(message.bankAccount, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBankAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBankAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3:
          if (tag !== 26) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateBankAccountRequest {
    return { bankAccount: isSet(object.bankAccount) ? BankAccount.fromJSON(object.bankAccount) : undefined };
  },

  toJSON(message: UpdateBankAccountRequest): unknown {
    const obj: any = {};
    message.bankAccount !== undefined &&
      (obj.bankAccount = message.bankAccount ? BankAccount.toJSON(message.bankAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateBankAccountRequest>, I>>(base?: I): UpdateBankAccountRequest {
    return UpdateBankAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateBankAccountRequest>, I>>(object: I): UpdateBankAccountRequest {
    const message = createBaseUpdateBankAccountRequest();
    message.bankAccount = (object.bankAccount !== undefined && object.bankAccount !== null)
      ? BankAccount.fromPartial(object.bankAccount)
      : undefined;
    return message;
  },
};

function createBaseUpdateBankAccountResponse(): UpdateBankAccountResponse {
  return { updated: false, bankAccount: undefined };
}

export const UpdateBankAccountResponse = {
  encode(message: UpdateBankAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.updated === true) {
      writer.uint32(8).bool(message.updated);
    }
    if (message.bankAccount !== undefined) {
      BankAccount.encode(message.bankAccount, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBankAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBankAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.updated = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateBankAccountResponse {
    return {
      updated: isSet(object.updated) ? Boolean(object.updated) : false,
      bankAccount: isSet(object.bankAccount) ? BankAccount.fromJSON(object.bankAccount) : undefined,
    };
  },

  toJSON(message: UpdateBankAccountResponse): unknown {
    const obj: any = {};
    message.updated !== undefined && (obj.updated = message.updated);
    message.bankAccount !== undefined &&
      (obj.bankAccount = message.bankAccount ? BankAccount.toJSON(message.bankAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateBankAccountResponse>, I>>(base?: I): UpdateBankAccountResponse {
    return UpdateBankAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateBankAccountResponse>, I>>(object: I): UpdateBankAccountResponse {
    const message = createBaseUpdateBankAccountResponse();
    message.updated = object.updated ?? false;
    message.bankAccount = (object.bankAccount !== undefined && object.bankAccount !== null)
      ? BankAccount.fromPartial(object.bankAccount)
      : undefined;
    return message;
  },
};

function createBaseGetPocketRequest(): GetPocketRequest {
  return { pocketId: 0 };
}

export const GetPocketRequest = {
  encode(message: GetPocketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pocketId !== 0) {
      writer.uint32(16).uint64(message.pocketId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPocketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPocketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPocketRequest {
    return { pocketId: isSet(object.pocketId) ? Number(object.pocketId) : 0 };
  },

  toJSON(message: GetPocketRequest): unknown {
    const obj: any = {};
    message.pocketId !== undefined && (obj.pocketId = Math.round(message.pocketId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPocketRequest>, I>>(base?: I): GetPocketRequest {
    return GetPocketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPocketRequest>, I>>(object: I): GetPocketRequest {
    const message = createBaseGetPocketRequest();
    message.pocketId = object.pocketId ?? 0;
    return message;
  },
};

function createBaseGetPocketResponse(): GetPocketResponse {
  return { pocket: undefined };
}

export const GetPocketResponse = {
  encode(message: GetPocketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pocket !== undefined) {
      Pocket.encode(message.pocket, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPocketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPocketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pocket = Pocket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPocketResponse {
    return { pocket: isSet(object.pocket) ? Pocket.fromJSON(object.pocket) : undefined };
  },

  toJSON(message: GetPocketResponse): unknown {
    const obj: any = {};
    message.pocket !== undefined && (obj.pocket = message.pocket ? Pocket.toJSON(message.pocket) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPocketResponse>, I>>(base?: I): GetPocketResponse {
    return GetPocketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPocketResponse>, I>>(object: I): GetPocketResponse {
    const message = createBaseGetPocketResponse();
    message.pocket = (object.pocket !== undefined && object.pocket !== null)
      ? Pocket.fromPartial(object.pocket)
      : undefined;
    return message;
  },
};

function createBaseGetSmartGoalsByPocketIdRequest(): GetSmartGoalsByPocketIdRequest {
  return { pocketId: 0 };
}

export const GetSmartGoalsByPocketIdRequest = {
  encode(message: GetSmartGoalsByPocketIdRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pocketId !== 0) {
      writer.uint32(16).uint64(message.pocketId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSmartGoalsByPocketIdRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSmartGoalsByPocketIdRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetSmartGoalsByPocketIdRequest {
    return { pocketId: isSet(object.pocketId) ? Number(object.pocketId) : 0 };
  },

  toJSON(message: GetSmartGoalsByPocketIdRequest): unknown {
    const obj: any = {};
    message.pocketId !== undefined && (obj.pocketId = Math.round(message.pocketId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetSmartGoalsByPocketIdRequest>, I>>(base?: I): GetSmartGoalsByPocketIdRequest {
    return GetSmartGoalsByPocketIdRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetSmartGoalsByPocketIdRequest>, I>>(
    object: I,
  ): GetSmartGoalsByPocketIdRequest {
    const message = createBaseGetSmartGoalsByPocketIdRequest();
    message.pocketId = object.pocketId ?? 0;
    return message;
  },
};

function createBaseGetSmartGoalsByPocketIdResponse(): GetSmartGoalsByPocketIdResponse {
  return { smartGoals: [] };
}

export const GetSmartGoalsByPocketIdResponse = {
  encode(message: GetSmartGoalsByPocketIdResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.smartGoals) {
      SmartGoal.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSmartGoalsByPocketIdResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSmartGoalsByPocketIdResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.smartGoals.push(SmartGoal.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetSmartGoalsByPocketIdResponse {
    return {
      smartGoals: Array.isArray(object?.smartGoals) ? object.smartGoals.map((e: any) => SmartGoal.fromJSON(e)) : [],
    };
  },

  toJSON(message: GetSmartGoalsByPocketIdResponse): unknown {
    const obj: any = {};
    if (message.smartGoals) {
      obj.smartGoals = message.smartGoals.map((e) => e ? SmartGoal.toJSON(e) : undefined);
    } else {
      obj.smartGoals = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetSmartGoalsByPocketIdResponse>, I>>(base?: I): GetSmartGoalsByPocketIdResponse {
    return GetSmartGoalsByPocketIdResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetSmartGoalsByPocketIdResponse>, I>>(
    object: I,
  ): GetSmartGoalsByPocketIdResponse {
    const message = createBaseGetSmartGoalsByPocketIdResponse();
    message.smartGoals = object.smartGoals?.map((e) => SmartGoal.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateSmartGoalRequest(): CreateSmartGoalRequest {
  return { pocketId: 0, smartGoal: undefined };
}

export const CreateSmartGoalRequest = {
  encode(message: CreateSmartGoalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pocketId !== 0) {
      writer.uint32(16).uint64(message.pocketId);
    }
    if (message.smartGoal !== undefined) {
      SmartGoal.encode(message.smartGoal, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateSmartGoalRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateSmartGoalRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.smartGoal = SmartGoal.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateSmartGoalRequest {
    return {
      pocketId: isSet(object.pocketId) ? Number(object.pocketId) : 0,
      smartGoal: isSet(object.smartGoal) ? SmartGoal.fromJSON(object.smartGoal) : undefined,
    };
  },

  toJSON(message: CreateSmartGoalRequest): unknown {
    const obj: any = {};
    message.pocketId !== undefined && (obj.pocketId = Math.round(message.pocketId));
    message.smartGoal !== undefined &&
      (obj.smartGoal = message.smartGoal ? SmartGoal.toJSON(message.smartGoal) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateSmartGoalRequest>, I>>(base?: I): CreateSmartGoalRequest {
    return CreateSmartGoalRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateSmartGoalRequest>, I>>(object: I): CreateSmartGoalRequest {
    const message = createBaseCreateSmartGoalRequest();
    message.pocketId = object.pocketId ?? 0;
    message.smartGoal = (object.smartGoal !== undefined && object.smartGoal !== null)
      ? SmartGoal.fromPartial(object.smartGoal)
      : undefined;
    return message;
  },
};

function createBaseCreateSmartGoalResponse(): CreateSmartGoalResponse {
  return { smartGoalId: 0 };
}

export const CreateSmartGoalResponse = {
  encode(message: CreateSmartGoalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoalId !== 0) {
      writer.uint32(8).uint64(message.smartGoalId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateSmartGoalResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateSmartGoalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateSmartGoalResponse {
    return { smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0 };
  },

  toJSON(message: CreateSmartGoalResponse): unknown {
    const obj: any = {};
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateSmartGoalResponse>, I>>(base?: I): CreateSmartGoalResponse {
    return CreateSmartGoalResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateSmartGoalResponse>, I>>(object: I): CreateSmartGoalResponse {
    const message = createBaseCreateSmartGoalResponse();
    message.smartGoalId = object.smartGoalId ?? 0;
    return message;
  },
};

function createBaseUpdateSmartGoalRequest(): UpdateSmartGoalRequest {
  return { smartGoal: undefined };
}

export const UpdateSmartGoalRequest = {
  encode(message: UpdateSmartGoalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoal !== undefined) {
      SmartGoal.encode(message.smartGoal, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateSmartGoalRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateSmartGoalRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3:
          if (tag !== 26) {
            break;
          }

          message.smartGoal = SmartGoal.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateSmartGoalRequest {
    return { smartGoal: isSet(object.smartGoal) ? SmartGoal.fromJSON(object.smartGoal) : undefined };
  },

  toJSON(message: UpdateSmartGoalRequest): unknown {
    const obj: any = {};
    message.smartGoal !== undefined &&
      (obj.smartGoal = message.smartGoal ? SmartGoal.toJSON(message.smartGoal) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateSmartGoalRequest>, I>>(base?: I): UpdateSmartGoalRequest {
    return UpdateSmartGoalRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateSmartGoalRequest>, I>>(object: I): UpdateSmartGoalRequest {
    const message = createBaseUpdateSmartGoalRequest();
    message.smartGoal = (object.smartGoal !== undefined && object.smartGoal !== null)
      ? SmartGoal.fromPartial(object.smartGoal)
      : undefined;
    return message;
  },
};

function createBaseUpdateSmartGoalResponse(): UpdateSmartGoalResponse {
  return { smartGoalId: 0 };
}

export const UpdateSmartGoalResponse = {
  encode(message: UpdateSmartGoalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoalId !== 0) {
      writer.uint32(8).uint64(message.smartGoalId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateSmartGoalResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateSmartGoalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateSmartGoalResponse {
    return { smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0 };
  },

  toJSON(message: UpdateSmartGoalResponse): unknown {
    const obj: any = {};
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateSmartGoalResponse>, I>>(base?: I): UpdateSmartGoalResponse {
    return UpdateSmartGoalResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateSmartGoalResponse>, I>>(object: I): UpdateSmartGoalResponse {
    const message = createBaseUpdateSmartGoalResponse();
    message.smartGoalId = object.smartGoalId ?? 0;
    return message;
  },
};

function createBaseDeleteSmartGoalRequest(): DeleteSmartGoalRequest {
  return { smartGoalId: 0 };
}

export const DeleteSmartGoalRequest = {
  encode(message: DeleteSmartGoalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoalId !== 0) {
      writer.uint32(16).uint64(message.smartGoalId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteSmartGoalRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteSmartGoalRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteSmartGoalRequest {
    return { smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0 };
  },

  toJSON(message: DeleteSmartGoalRequest): unknown {
    const obj: any = {};
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteSmartGoalRequest>, I>>(base?: I): DeleteSmartGoalRequest {
    return DeleteSmartGoalRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteSmartGoalRequest>, I>>(object: I): DeleteSmartGoalRequest {
    const message = createBaseDeleteSmartGoalRequest();
    message.smartGoalId = object.smartGoalId ?? 0;
    return message;
  },
};

function createBaseDeleteSmartGoalResponse(): DeleteSmartGoalResponse {
  return { deleted: false };
}

export const DeleteSmartGoalResponse = {
  encode(message: DeleteSmartGoalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deleted === true) {
      writer.uint32(8).bool(message.deleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteSmartGoalResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteSmartGoalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteSmartGoalResponse {
    return { deleted: isSet(object.deleted) ? Boolean(object.deleted) : false };
  },

  toJSON(message: DeleteSmartGoalResponse): unknown {
    const obj: any = {};
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteSmartGoalResponse>, I>>(base?: I): DeleteSmartGoalResponse {
    return DeleteSmartGoalResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteSmartGoalResponse>, I>>(object: I): DeleteSmartGoalResponse {
    const message = createBaseDeleteSmartGoalResponse();
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseCreateMilestoneRequest(): CreateMilestoneRequest {
  return { smartGoalId: 0, milestone: undefined };
}

export const CreateMilestoneRequest = {
  encode(message: CreateMilestoneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoalId !== 0) {
      writer.uint32(8).uint64(message.smartGoalId);
    }
    if (message.milestone !== undefined) {
      Milestone.encode(message.milestone, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMilestoneRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMilestoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMilestoneRequest {
    return {
      smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0,
      milestone: isSet(object.milestone) ? Milestone.fromJSON(object.milestone) : undefined,
    };
  },

  toJSON(message: CreateMilestoneRequest): unknown {
    const obj: any = {};
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    message.milestone !== undefined &&
      (obj.milestone = message.milestone ? Milestone.toJSON(message.milestone) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateMilestoneRequest>, I>>(base?: I): CreateMilestoneRequest {
    return CreateMilestoneRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateMilestoneRequest>, I>>(object: I): CreateMilestoneRequest {
    const message = createBaseCreateMilestoneRequest();
    message.smartGoalId = object.smartGoalId ?? 0;
    message.milestone = (object.milestone !== undefined && object.milestone !== null)
      ? Milestone.fromPartial(object.milestone)
      : undefined;
    return message;
  },
};

function createBaseCreateMilestoneResponse(): CreateMilestoneResponse {
  return { milestoneId: 0 };
}

export const CreateMilestoneResponse = {
  encode(message: CreateMilestoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestoneId !== 0) {
      writer.uint32(8).uint64(message.milestoneId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMilestoneResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMilestoneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMilestoneResponse {
    return { milestoneId: isSet(object.milestoneId) ? Number(object.milestoneId) : 0 };
  },

  toJSON(message: CreateMilestoneResponse): unknown {
    const obj: any = {};
    message.milestoneId !== undefined && (obj.milestoneId = Math.round(message.milestoneId));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateMilestoneResponse>, I>>(base?: I): CreateMilestoneResponse {
    return CreateMilestoneResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateMilestoneResponse>, I>>(object: I): CreateMilestoneResponse {
    const message = createBaseCreateMilestoneResponse();
    message.milestoneId = object.milestoneId ?? 0;
    return message;
  },
};

function createBaseDeleteMilestoneRequest(): DeleteMilestoneRequest {
  return { milestoneId: 0 };
}

export const DeleteMilestoneRequest = {
  encode(message: DeleteMilestoneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestoneId !== 0) {
      writer.uint32(16).uint64(message.milestoneId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteMilestoneRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteMilestoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteMilestoneRequest {
    return { milestoneId: isSet(object.milestoneId) ? Number(object.milestoneId) : 0 };
  },

  toJSON(message: DeleteMilestoneRequest): unknown {
    const obj: any = {};
    message.milestoneId !== undefined && (obj.milestoneId = Math.round(message.milestoneId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteMilestoneRequest>, I>>(base?: I): DeleteMilestoneRequest {
    return DeleteMilestoneRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteMilestoneRequest>, I>>(object: I): DeleteMilestoneRequest {
    const message = createBaseDeleteMilestoneRequest();
    message.milestoneId = object.milestoneId ?? 0;
    return message;
  },
};

function createBaseDeleteMilestoneResponse(): DeleteMilestoneResponse {
  return { deleted: false };
}

export const DeleteMilestoneResponse = {
  encode(message: DeleteMilestoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deleted === true) {
      writer.uint32(8).bool(message.deleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteMilestoneResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteMilestoneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteMilestoneResponse {
    return { deleted: isSet(object.deleted) ? Boolean(object.deleted) : false };
  },

  toJSON(message: DeleteMilestoneResponse): unknown {
    const obj: any = {};
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteMilestoneResponse>, I>>(base?: I): DeleteMilestoneResponse {
    return DeleteMilestoneResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteMilestoneResponse>, I>>(object: I): DeleteMilestoneResponse {
    const message = createBaseDeleteMilestoneResponse();
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseUpdateMilestoneRequest(): UpdateMilestoneRequest {
  return { milestone: undefined };
}

export const UpdateMilestoneRequest = {
  encode(message: UpdateMilestoneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestone !== undefined) {
      Milestone.encode(message.milestone, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateMilestoneRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateMilestoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateMilestoneRequest {
    return { milestone: isSet(object.milestone) ? Milestone.fromJSON(object.milestone) : undefined };
  },

  toJSON(message: UpdateMilestoneRequest): unknown {
    const obj: any = {};
    message.milestone !== undefined &&
      (obj.milestone = message.milestone ? Milestone.toJSON(message.milestone) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateMilestoneRequest>, I>>(base?: I): UpdateMilestoneRequest {
    return UpdateMilestoneRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateMilestoneRequest>, I>>(object: I): UpdateMilestoneRequest {
    const message = createBaseUpdateMilestoneRequest();
    message.milestone = (object.milestone !== undefined && object.milestone !== null)
      ? Milestone.fromPartial(object.milestone)
      : undefined;
    return message;
  },
};

function createBaseUpdateMilestoneResponse(): UpdateMilestoneResponse {
  return { milestone: undefined };
}

export const UpdateMilestoneResponse = {
  encode(message: UpdateMilestoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestone !== undefined) {
      Milestone.encode(message.milestone, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateMilestoneResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateMilestoneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateMilestoneResponse {
    return { milestone: isSet(object.milestone) ? Milestone.fromJSON(object.milestone) : undefined };
  },

  toJSON(message: UpdateMilestoneResponse): unknown {
    const obj: any = {};
    message.milestone !== undefined &&
      (obj.milestone = message.milestone ? Milestone.toJSON(message.milestone) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateMilestoneResponse>, I>>(base?: I): UpdateMilestoneResponse {
    return UpdateMilestoneResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateMilestoneResponse>, I>>(object: I): UpdateMilestoneResponse {
    const message = createBaseUpdateMilestoneResponse();
    message.milestone = (object.milestone !== undefined && object.milestone !== null)
      ? Milestone.fromPartial(object.milestone)
      : undefined;
    return message;
  },
};

function createBaseGetMilestonesBySmartGoalIdRequest(): GetMilestonesBySmartGoalIdRequest {
  return { smartGoalId: 0 };
}

export const GetMilestonesBySmartGoalIdRequest = {
  encode(message: GetMilestonesBySmartGoalIdRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoalId !== 0) {
      writer.uint32(8).uint64(message.smartGoalId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMilestonesBySmartGoalIdRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMilestonesBySmartGoalIdRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMilestonesBySmartGoalIdRequest {
    return { smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0 };
  },

  toJSON(message: GetMilestonesBySmartGoalIdRequest): unknown {
    const obj: any = {};
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMilestonesBySmartGoalIdRequest>, I>>(
    base?: I,
  ): GetMilestonesBySmartGoalIdRequest {
    return GetMilestonesBySmartGoalIdRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMilestonesBySmartGoalIdRequest>, I>>(
    object: I,
  ): GetMilestonesBySmartGoalIdRequest {
    const message = createBaseGetMilestonesBySmartGoalIdRequest();
    message.smartGoalId = object.smartGoalId ?? 0;
    return message;
  },
};

function createBaseGetMilestonesBySmartGoalIdResponse(): GetMilestonesBySmartGoalIdResponse {
  return { milestones: [] };
}

export const GetMilestonesBySmartGoalIdResponse = {
  encode(message: GetMilestonesBySmartGoalIdResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.milestones) {
      Milestone.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMilestonesBySmartGoalIdResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMilestonesBySmartGoalIdResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.milestones.push(Milestone.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMilestonesBySmartGoalIdResponse {
    return {
      milestones: Array.isArray(object?.milestones) ? object.milestones.map((e: any) => Milestone.fromJSON(e)) : [],
    };
  },

  toJSON(message: GetMilestonesBySmartGoalIdResponse): unknown {
    const obj: any = {};
    if (message.milestones) {
      obj.milestones = message.milestones.map((e) => e ? Milestone.toJSON(e) : undefined);
    } else {
      obj.milestones = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMilestonesBySmartGoalIdResponse>, I>>(
    base?: I,
  ): GetMilestonesBySmartGoalIdResponse {
    return GetMilestonesBySmartGoalIdResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMilestonesBySmartGoalIdResponse>, I>>(
    object: I,
  ): GetMilestonesBySmartGoalIdResponse {
    const message = createBaseGetMilestonesBySmartGoalIdResponse();
    message.milestones = object.milestones?.map((e) => Milestone.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetMilestoneRequest(): GetMilestoneRequest {
  return { milestoneId: 0 };
}

export const GetMilestoneRequest = {
  encode(message: GetMilestoneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestoneId !== 0) {
      writer.uint32(16).uint64(message.milestoneId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMilestoneRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMilestoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMilestoneRequest {
    return { milestoneId: isSet(object.milestoneId) ? Number(object.milestoneId) : 0 };
  },

  toJSON(message: GetMilestoneRequest): unknown {
    const obj: any = {};
    message.milestoneId !== undefined && (obj.milestoneId = Math.round(message.milestoneId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMilestoneRequest>, I>>(base?: I): GetMilestoneRequest {
    return GetMilestoneRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMilestoneRequest>, I>>(object: I): GetMilestoneRequest {
    const message = createBaseGetMilestoneRequest();
    message.milestoneId = object.milestoneId ?? 0;
    return message;
  },
};

function createBaseGetMilestoneResponse(): GetMilestoneResponse {
  return { milestone: undefined };
}

export const GetMilestoneResponse = {
  encode(message: GetMilestoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestone !== undefined) {
      Milestone.encode(message.milestone, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMilestoneResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMilestoneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMilestoneResponse {
    return { milestone: isSet(object.milestone) ? Milestone.fromJSON(object.milestone) : undefined };
  },

  toJSON(message: GetMilestoneResponse): unknown {
    const obj: any = {};
    message.milestone !== undefined &&
      (obj.milestone = message.milestone ? Milestone.toJSON(message.milestone) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMilestoneResponse>, I>>(base?: I): GetMilestoneResponse {
    return GetMilestoneResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMilestoneResponse>, I>>(object: I): GetMilestoneResponse {
    const message = createBaseGetMilestoneResponse();
    message.milestone = (object.milestone !== undefined && object.milestone !== null)
      ? Milestone.fromPartial(object.milestone)
      : undefined;
    return message;
  },
};

function createBaseGetForecastRequest(): GetForecastRequest {
  return { smartGoalId: 0 };
}

export const GetForecastRequest = {
  encode(message: GetForecastRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.smartGoalId !== 0) {
      writer.uint32(8).uint64(message.smartGoalId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetForecastRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetForecastRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetForecastRequest {
    return { smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0 };
  },

  toJSON(message: GetForecastRequest): unknown {
    const obj: any = {};
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetForecastRequest>, I>>(base?: I): GetForecastRequest {
    return GetForecastRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetForecastRequest>, I>>(object: I): GetForecastRequest {
    const message = createBaseGetForecastRequest();
    message.smartGoalId = object.smartGoalId ?? 0;
    return message;
  },
};

function createBaseGetForecastResponse(): GetForecastResponse {
  return { forecast: undefined };
}

export const GetForecastResponse = {
  encode(message: GetForecastResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.forecast !== undefined) {
      Forecast.encode(message.forecast, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetForecastResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetForecastResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.forecast = Forecast.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetForecastResponse {
    return { forecast: isSet(object.forecast) ? Forecast.fromJSON(object.forecast) : undefined };
  },

  toJSON(message: GetForecastResponse): unknown {
    const obj: any = {};
    message.forecast !== undefined && (obj.forecast = message.forecast ? Forecast.toJSON(message.forecast) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetForecastResponse>, I>>(base?: I): GetForecastResponse {
    return GetForecastResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetForecastResponse>, I>>(object: I): GetForecastResponse {
    const message = createBaseGetForecastResponse();
    message.forecast = (object.forecast !== undefined && object.forecast !== null)
      ? Forecast.fromPartial(object.forecast)
      : undefined;
    return message;
  },
};

function createBaseCreateBudgetRequest(): CreateBudgetRequest {
  return { milestroneId: 0, budget: undefined };
}

export const CreateBudgetRequest = {
  encode(message: CreateBudgetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.milestroneId !== 0) {
      writer.uint32(16).uint64(message.milestroneId);
    }
    if (message.budget !== undefined) {
      Budget.encode(message.budget, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBudgetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBudgetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.milestroneId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateBudgetRequest {
    return {
      milestroneId: isSet(object.milestroneId) ? Number(object.milestroneId) : 0,
      budget: isSet(object.budget) ? Budget.fromJSON(object.budget) : undefined,
    };
  },

  toJSON(message: CreateBudgetRequest): unknown {
    const obj: any = {};
    message.milestroneId !== undefined && (obj.milestroneId = Math.round(message.milestroneId));
    message.budget !== undefined && (obj.budget = message.budget ? Budget.toJSON(message.budget) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBudgetRequest>, I>>(base?: I): CreateBudgetRequest {
    return CreateBudgetRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateBudgetRequest>, I>>(object: I): CreateBudgetRequest {
    const message = createBaseCreateBudgetRequest();
    message.milestroneId = object.milestroneId ?? 0;
    message.budget = (object.budget !== undefined && object.budget !== null)
      ? Budget.fromPartial(object.budget)
      : undefined;
    return message;
  },
};

function createBaseCreateBudgetResponse(): CreateBudgetResponse {
  return { budgetId: 0 };
}

export const CreateBudgetResponse = {
  encode(message: CreateBudgetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.budgetId !== 0) {
      writer.uint32(8).uint64(message.budgetId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBudgetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBudgetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.budgetId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateBudgetResponse {
    return { budgetId: isSet(object.budgetId) ? Number(object.budgetId) : 0 };
  },

  toJSON(message: CreateBudgetResponse): unknown {
    const obj: any = {};
    message.budgetId !== undefined && (obj.budgetId = Math.round(message.budgetId));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBudgetResponse>, I>>(base?: I): CreateBudgetResponse {
    return CreateBudgetResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateBudgetResponse>, I>>(object: I): CreateBudgetResponse {
    const message = createBaseCreateBudgetResponse();
    message.budgetId = object.budgetId ?? 0;
    return message;
  },
};

function createBaseUpdateBudgetRequest(): UpdateBudgetRequest {
  return { budget: undefined };
}

export const UpdateBudgetRequest = {
  encode(message: UpdateBudgetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.budget !== undefined) {
      Budget.encode(message.budget, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBudgetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBudgetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateBudgetRequest {
    return { budget: isSet(object.budget) ? Budget.fromJSON(object.budget) : undefined };
  },

  toJSON(message: UpdateBudgetRequest): unknown {
    const obj: any = {};
    message.budget !== undefined && (obj.budget = message.budget ? Budget.toJSON(message.budget) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateBudgetRequest>, I>>(base?: I): UpdateBudgetRequest {
    return UpdateBudgetRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateBudgetRequest>, I>>(object: I): UpdateBudgetRequest {
    const message = createBaseUpdateBudgetRequest();
    message.budget = (object.budget !== undefined && object.budget !== null)
      ? Budget.fromPartial(object.budget)
      : undefined;
    return message;
  },
};

function createBaseUpdateBudgetResponse(): UpdateBudgetResponse {
  return { budget: undefined };
}

export const UpdateBudgetResponse = {
  encode(message: UpdateBudgetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.budget !== undefined) {
      Budget.encode(message.budget, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBudgetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBudgetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateBudgetResponse {
    return { budget: isSet(object.budget) ? Budget.fromJSON(object.budget) : undefined };
  },

  toJSON(message: UpdateBudgetResponse): unknown {
    const obj: any = {};
    message.budget !== undefined && (obj.budget = message.budget ? Budget.toJSON(message.budget) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateBudgetResponse>, I>>(base?: I): UpdateBudgetResponse {
    return UpdateBudgetResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateBudgetResponse>, I>>(object: I): UpdateBudgetResponse {
    const message = createBaseUpdateBudgetResponse();
    message.budget = (object.budget !== undefined && object.budget !== null)
      ? Budget.fromPartial(object.budget)
      : undefined;
    return message;
  },
};

function createBaseDeleteBudgetRequest(): DeleteBudgetRequest {
  return { budgetId: 0 };
}

export const DeleteBudgetRequest = {
  encode(message: DeleteBudgetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.budgetId !== 0) {
      writer.uint32(16).uint64(message.budgetId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBudgetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBudgetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.budgetId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteBudgetRequest {
    return { budgetId: isSet(object.budgetId) ? Number(object.budgetId) : 0 };
  },

  toJSON(message: DeleteBudgetRequest): unknown {
    const obj: any = {};
    message.budgetId !== undefined && (obj.budgetId = Math.round(message.budgetId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteBudgetRequest>, I>>(base?: I): DeleteBudgetRequest {
    return DeleteBudgetRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBudgetRequest>, I>>(object: I): DeleteBudgetRequest {
    const message = createBaseDeleteBudgetRequest();
    message.budgetId = object.budgetId ?? 0;
    return message;
  },
};

function createBaseDeleteBudgetResponse(): DeleteBudgetResponse {
  return { deleted: false };
}

export const DeleteBudgetResponse = {
  encode(message: DeleteBudgetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deleted === true) {
      writer.uint32(8).bool(message.deleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBudgetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBudgetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteBudgetResponse {
    return { deleted: isSet(object.deleted) ? Boolean(object.deleted) : false };
  },

  toJSON(message: DeleteBudgetResponse): unknown {
    const obj: any = {};
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteBudgetResponse>, I>>(base?: I): DeleteBudgetResponse {
    return DeleteBudgetResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBudgetResponse>, I>>(object: I): DeleteBudgetResponse {
    const message = createBaseDeleteBudgetResponse();
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseGetBudgetRequest(): GetBudgetRequest {
  return { budgetId: 0 };
}

export const GetBudgetRequest = {
  encode(message: GetBudgetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.budgetId !== 0) {
      writer.uint32(8).uint64(message.budgetId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBudgetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBudgetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.budgetId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetBudgetRequest {
    return { budgetId: isSet(object.budgetId) ? Number(object.budgetId) : 0 };
  },

  toJSON(message: GetBudgetRequest): unknown {
    const obj: any = {};
    message.budgetId !== undefined && (obj.budgetId = Math.round(message.budgetId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetBudgetRequest>, I>>(base?: I): GetBudgetRequest {
    return GetBudgetRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetBudgetRequest>, I>>(object: I): GetBudgetRequest {
    const message = createBaseGetBudgetRequest();
    message.budgetId = object.budgetId ?? 0;
    return message;
  },
};

function createBaseGetBudgetResponse(): GetBudgetResponse {
  return { budget: undefined };
}

export const GetBudgetResponse = {
  encode(message: GetBudgetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.budget !== undefined) {
      Budget.encode(message.budget, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBudgetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBudgetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetBudgetResponse {
    return { budget: isSet(object.budget) ? Budget.fromJSON(object.budget) : undefined };
  },

  toJSON(message: GetBudgetResponse): unknown {
    const obj: any = {};
    message.budget !== undefined && (obj.budget = message.budget ? Budget.toJSON(message.budget) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetBudgetResponse>, I>>(base?: I): GetBudgetResponse {
    return GetBudgetResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetBudgetResponse>, I>>(object: I): GetBudgetResponse {
    const message = createBaseGetBudgetResponse();
    message.budget = (object.budget !== undefined && object.budget !== null)
      ? Budget.fromPartial(object.budget)
      : undefined;
    return message;
  },
};

function createBaseGetAllBudgetsRequest(): GetAllBudgetsRequest {
  return { pocketId: 0, smartGoalId: 0, milestoneId: 0 };
}

export const GetAllBudgetsRequest = {
  encode(message: GetAllBudgetsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pocketId !== 0) {
      writer.uint32(8).uint64(message.pocketId);
    }
    if (message.smartGoalId !== 0) {
      writer.uint32(16).uint64(message.smartGoalId);
    }
    if (message.milestoneId !== 0) {
      writer.uint32(24).uint64(message.milestoneId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAllBudgetsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAllBudgetsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetAllBudgetsRequest {
    return {
      pocketId: isSet(object.pocketId) ? Number(object.pocketId) : 0,
      smartGoalId: isSet(object.smartGoalId) ? Number(object.smartGoalId) : 0,
      milestoneId: isSet(object.milestoneId) ? Number(object.milestoneId) : 0,
    };
  },

  toJSON(message: GetAllBudgetsRequest): unknown {
    const obj: any = {};
    message.pocketId !== undefined && (obj.pocketId = Math.round(message.pocketId));
    message.smartGoalId !== undefined && (obj.smartGoalId = Math.round(message.smartGoalId));
    message.milestoneId !== undefined && (obj.milestoneId = Math.round(message.milestoneId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetAllBudgetsRequest>, I>>(base?: I): GetAllBudgetsRequest {
    return GetAllBudgetsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetAllBudgetsRequest>, I>>(object: I): GetAllBudgetsRequest {
    const message = createBaseGetAllBudgetsRequest();
    message.pocketId = object.pocketId ?? 0;
    message.smartGoalId = object.smartGoalId ?? 0;
    message.milestoneId = object.milestoneId ?? 0;
    return message;
  },
};

function createBaseGetAllBudgetsResponse(): GetAllBudgetsResponse {
  return { budgets: [] };
}

export const GetAllBudgetsResponse = {
  encode(message: GetAllBudgetsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.budgets) {
      Budget.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAllBudgetsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAllBudgetsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.budgets.push(Budget.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetAllBudgetsResponse {
    return { budgets: Array.isArray(object?.budgets) ? object.budgets.map((e: any) => Budget.fromJSON(e)) : [] };
  },

  toJSON(message: GetAllBudgetsResponse): unknown {
    const obj: any = {};
    if (message.budgets) {
      obj.budgets = message.budgets.map((e) => e ? Budget.toJSON(e) : undefined);
    } else {
      obj.budgets = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetAllBudgetsResponse>, I>>(base?: I): GetAllBudgetsResponse {
    return GetAllBudgetsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetAllBudgetsResponse>, I>>(object: I): GetAllBudgetsResponse {
    const message = createBaseGetAllBudgetsResponse();
    message.budgets = object.budgets?.map((e) => Budget.fromPartial(e)) || [];
    return message;
  },
};

function createBaseHealthCheckRequest(): HealthCheckRequest {
  return {};
}

export const HealthCheckRequest = {
  encode(_: HealthCheckRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HealthCheckRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHealthCheckRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): HealthCheckRequest {
    return {};
  },

  toJSON(_: HealthCheckRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<HealthCheckRequest>, I>>(base?: I): HealthCheckRequest {
    return HealthCheckRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<HealthCheckRequest>, I>>(_: I): HealthCheckRequest {
    const message = createBaseHealthCheckRequest();
    return message;
  },
};

function createBaseHealthCheckResponse(): HealthCheckResponse {
  return { healthy: false };
}

export const HealthCheckResponse = {
  encode(message: HealthCheckResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.healthy === true) {
      writer.uint32(8).bool(message.healthy);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HealthCheckResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHealthCheckResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.healthy = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): HealthCheckResponse {
    return { healthy: isSet(object.healthy) ? Boolean(object.healthy) : false };
  },

  toJSON(message: HealthCheckResponse): unknown {
    const obj: any = {};
    message.healthy !== undefined && (obj.healthy = message.healthy);
    return obj;
  },

  create<I extends Exact<DeepPartial<HealthCheckResponse>, I>>(base?: I): HealthCheckResponse {
    return HealthCheckResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<HealthCheckResponse>, I>>(object: I): HealthCheckResponse {
    const message = createBaseHealthCheckResponse();
    message.healthy = object.healthy ?? false;
    return message;
  },
};

function createBaseReadynessCheckRequest(): ReadynessCheckRequest {
  return {};
}

export const ReadynessCheckRequest = {
  encode(_: ReadynessCheckRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadynessCheckRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadynessCheckRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): ReadynessCheckRequest {
    return {};
  },

  toJSON(_: ReadynessCheckRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadynessCheckRequest>, I>>(base?: I): ReadynessCheckRequest {
    return ReadynessCheckRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadynessCheckRequest>, I>>(_: I): ReadynessCheckRequest {
    const message = createBaseReadynessCheckRequest();
    return message;
  },
};

function createBaseReadynessCheckResponse(): ReadynessCheckResponse {
  return { healthy: false };
}

export const ReadynessCheckResponse = {
  encode(message: ReadynessCheckResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.healthy === true) {
      writer.uint32(8).bool(message.healthy);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadynessCheckResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadynessCheckResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.healthy = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadynessCheckResponse {
    return { healthy: isSet(object.healthy) ? Boolean(object.healthy) : false };
  },

  toJSON(message: ReadynessCheckResponse): unknown {
    const obj: any = {};
    message.healthy !== undefined && (obj.healthy = message.healthy);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadynessCheckResponse>, I>>(base?: I): ReadynessCheckResponse {
    return ReadynessCheckResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadynessCheckResponse>, I>>(object: I): ReadynessCheckResponse {
    const message = createBaseReadynessCheckResponse();
    message.healthy = object.healthy ?? false;
    return message;
  },
};

function createBasePlaidInitiateTokenExchangeRequest(): PlaidInitiateTokenExchangeRequest {
  return { userId: 0, fullName: "", email: "", phoneNumber: "" };
}

export const PlaidInitiateTokenExchangeRequest = {
  encode(message: PlaidInitiateTokenExchangeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.fullName !== "") {
      writer.uint32(18).string(message.fullName);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    if (message.phoneNumber !== "") {
      writer.uint32(34).string(message.phoneNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PlaidInitiateTokenExchangeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaidInitiateTokenExchangeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.fullName = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.email = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.phoneNumber = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PlaidInitiateTokenExchangeRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      fullName: isSet(object.fullName) ? String(object.fullName) : "",
      email: isSet(object.email) ? String(object.email) : "",
      phoneNumber: isSet(object.phoneNumber) ? String(object.phoneNumber) : "",
    };
  },

  toJSON(message: PlaidInitiateTokenExchangeRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.fullName !== undefined && (obj.fullName = message.fullName);
    message.email !== undefined && (obj.email = message.email);
    message.phoneNumber !== undefined && (obj.phoneNumber = message.phoneNumber);
    return obj;
  },

  create<I extends Exact<DeepPartial<PlaidInitiateTokenExchangeRequest>, I>>(
    base?: I,
  ): PlaidInitiateTokenExchangeRequest {
    return PlaidInitiateTokenExchangeRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PlaidInitiateTokenExchangeRequest>, I>>(
    object: I,
  ): PlaidInitiateTokenExchangeRequest {
    const message = createBasePlaidInitiateTokenExchangeRequest();
    message.userId = object.userId ?? 0;
    message.fullName = object.fullName ?? "";
    message.email = object.email ?? "";
    message.phoneNumber = object.phoneNumber ?? "";
    return message;
  },
};

function createBasePlaidInitiateTokenExchangeResponse(): PlaidInitiateTokenExchangeResponse {
  return { linkToken: "", expiration: "", plaidRequestId: "" };
}

export const PlaidInitiateTokenExchangeResponse = {
  encode(message: PlaidInitiateTokenExchangeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.linkToken !== "") {
      writer.uint32(10).string(message.linkToken);
    }
    if (message.expiration !== "") {
      writer.uint32(18).string(message.expiration);
    }
    if (message.plaidRequestId !== "") {
      writer.uint32(26).string(message.plaidRequestId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PlaidInitiateTokenExchangeResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaidInitiateTokenExchangeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.linkToken = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.expiration = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.plaidRequestId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PlaidInitiateTokenExchangeResponse {
    return {
      linkToken: isSet(object.linkToken) ? String(object.linkToken) : "",
      expiration: isSet(object.expiration) ? String(object.expiration) : "",
      plaidRequestId: isSet(object.plaidRequestId) ? String(object.plaidRequestId) : "",
    };
  },

  toJSON(message: PlaidInitiateTokenExchangeResponse): unknown {
    const obj: any = {};
    message.linkToken !== undefined && (obj.linkToken = message.linkToken);
    message.expiration !== undefined && (obj.expiration = message.expiration);
    message.plaidRequestId !== undefined && (obj.plaidRequestId = message.plaidRequestId);
    return obj;
  },

  create<I extends Exact<DeepPartial<PlaidInitiateTokenExchangeResponse>, I>>(
    base?: I,
  ): PlaidInitiateTokenExchangeResponse {
    return PlaidInitiateTokenExchangeResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PlaidInitiateTokenExchangeResponse>, I>>(
    object: I,
  ): PlaidInitiateTokenExchangeResponse {
    const message = createBasePlaidInitiateTokenExchangeResponse();
    message.linkToken = object.linkToken ?? "";
    message.expiration = object.expiration ?? "";
    message.plaidRequestId = object.plaidRequestId ?? "";
    return message;
  },
};

function createBasePlaidExchangeTokenRequest(): PlaidExchangeTokenRequest {
  return { userId: 0, publicToken: "", institutionId: "", institutionName: "" };
}

export const PlaidExchangeTokenRequest = {
  encode(message: PlaidExchangeTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.publicToken !== "") {
      writer.uint32(18).string(message.publicToken);
    }
    if (message.institutionId !== "") {
      writer.uint32(26).string(message.institutionId);
    }
    if (message.institutionName !== "") {
      writer.uint32(34).string(message.institutionName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PlaidExchangeTokenRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaidExchangeTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.publicToken = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.institutionId = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.institutionName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PlaidExchangeTokenRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      publicToken: isSet(object.publicToken) ? String(object.publicToken) : "",
      institutionId: isSet(object.institutionId) ? String(object.institutionId) : "",
      institutionName: isSet(object.institutionName) ? String(object.institutionName) : "",
    };
  },

  toJSON(message: PlaidExchangeTokenRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.publicToken !== undefined && (obj.publicToken = message.publicToken);
    message.institutionId !== undefined && (obj.institutionId = message.institutionId);
    message.institutionName !== undefined && (obj.institutionName = message.institutionName);
    return obj;
  },

  create<I extends Exact<DeepPartial<PlaidExchangeTokenRequest>, I>>(base?: I): PlaidExchangeTokenRequest {
    return PlaidExchangeTokenRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PlaidExchangeTokenRequest>, I>>(object: I): PlaidExchangeTokenRequest {
    const message = createBasePlaidExchangeTokenRequest();
    message.userId = object.userId ?? 0;
    message.publicToken = object.publicToken ?? "";
    message.institutionId = object.institutionId ?? "";
    message.institutionName = object.institutionName ?? "";
    return message;
  },
};

function createBasePlaidExchangeTokenResponse(): PlaidExchangeTokenResponse {
  return { success: false };
}

export const PlaidExchangeTokenResponse = {
  encode(message: PlaidExchangeTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PlaidExchangeTokenResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaidExchangeTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PlaidExchangeTokenResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: PlaidExchangeTokenResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  create<I extends Exact<DeepPartial<PlaidExchangeTokenResponse>, I>>(base?: I): PlaidExchangeTokenResponse {
    return PlaidExchangeTokenResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PlaidExchangeTokenResponse>, I>>(object: I): PlaidExchangeTokenResponse {
    const message = createBasePlaidExchangeTokenResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseGetInvestmentAcccountRequest(): GetInvestmentAcccountRequest {
  return { userId: 0, investmentAccountId: 0 };
}

export const GetInvestmentAcccountRequest = {
  encode(message: GetInvestmentAcccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.investmentAccountId !== 0) {
      writer.uint32(16).uint64(message.investmentAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetInvestmentAcccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetInvestmentAcccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.investmentAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetInvestmentAcccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      investmentAccountId: isSet(object.investmentAccountId) ? Number(object.investmentAccountId) : 0,
    };
  },

  toJSON(message: GetInvestmentAcccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.investmentAccountId !== undefined && (obj.investmentAccountId = Math.round(message.investmentAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetInvestmentAcccountRequest>, I>>(base?: I): GetInvestmentAcccountRequest {
    return GetInvestmentAcccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetInvestmentAcccountRequest>, I>>(object: I): GetInvestmentAcccountRequest {
    const message = createBaseGetInvestmentAcccountRequest();
    message.userId = object.userId ?? 0;
    message.investmentAccountId = object.investmentAccountId ?? 0;
    return message;
  },
};

function createBaseGetInvestmentAcccountResponse(): GetInvestmentAcccountResponse {
  return { investmentAccount: undefined };
}

export const GetInvestmentAcccountResponse = {
  encode(message: GetInvestmentAcccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investmentAccount !== undefined) {
      InvestmentAccount.encode(message.investmentAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetInvestmentAcccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetInvestmentAcccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.investmentAccount = InvestmentAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetInvestmentAcccountResponse {
    return {
      investmentAccount: isSet(object.investmentAccount)
        ? InvestmentAccount.fromJSON(object.investmentAccount)
        : undefined,
    };
  },

  toJSON(message: GetInvestmentAcccountResponse): unknown {
    const obj: any = {};
    message.investmentAccount !== undefined && (obj.investmentAccount = message.investmentAccount
      ? InvestmentAccount.toJSON(message.investmentAccount)
      : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetInvestmentAcccountResponse>, I>>(base?: I): GetInvestmentAcccountResponse {
    return GetInvestmentAcccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetInvestmentAcccountResponse>, I>>(
    object: I,
  ): GetInvestmentAcccountResponse {
    const message = createBaseGetInvestmentAcccountResponse();
    message.investmentAccount = (object.investmentAccount !== undefined && object.investmentAccount !== null)
      ? InvestmentAccount.fromPartial(object.investmentAccount)
      : undefined;
    return message;
  },
};

function createBaseGetMortgageAccountRequest(): GetMortgageAccountRequest {
  return { userId: 0, mortgageAccountId: 0 };
}

export const GetMortgageAccountRequest = {
  encode(message: GetMortgageAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.mortgageAccountId !== 0) {
      writer.uint32(16).uint64(message.mortgageAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMortgageAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMortgageAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.mortgageAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMortgageAccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      mortgageAccountId: isSet(object.mortgageAccountId) ? Number(object.mortgageAccountId) : 0,
    };
  },

  toJSON(message: GetMortgageAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.mortgageAccountId !== undefined && (obj.mortgageAccountId = Math.round(message.mortgageAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMortgageAccountRequest>, I>>(base?: I): GetMortgageAccountRequest {
    return GetMortgageAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMortgageAccountRequest>, I>>(object: I): GetMortgageAccountRequest {
    const message = createBaseGetMortgageAccountRequest();
    message.userId = object.userId ?? 0;
    message.mortgageAccountId = object.mortgageAccountId ?? 0;
    return message;
  },
};

function createBaseGetMortgageAccountResponse(): GetMortgageAccountResponse {
  return { mortageAccount: undefined };
}

export const GetMortgageAccountResponse = {
  encode(message: GetMortgageAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mortageAccount !== undefined) {
      MortgageAccount.encode(message.mortageAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMortgageAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMortgageAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.mortageAccount = MortgageAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMortgageAccountResponse {
    return {
      mortageAccount: isSet(object.mortageAccount) ? MortgageAccount.fromJSON(object.mortageAccount) : undefined,
    };
  },

  toJSON(message: GetMortgageAccountResponse): unknown {
    const obj: any = {};
    message.mortageAccount !== undefined &&
      (obj.mortageAccount = message.mortageAccount ? MortgageAccount.toJSON(message.mortageAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMortgageAccountResponse>, I>>(base?: I): GetMortgageAccountResponse {
    return GetMortgageAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMortgageAccountResponse>, I>>(object: I): GetMortgageAccountResponse {
    const message = createBaseGetMortgageAccountResponse();
    message.mortageAccount = (object.mortageAccount !== undefined && object.mortageAccount !== null)
      ? MortgageAccount.fromPartial(object.mortageAccount)
      : undefined;
    return message;
  },
};

function createBaseGetLiabilityAccountRequest(): GetLiabilityAccountRequest {
  return { userId: 0, liabilityAccountId: 0 };
}

export const GetLiabilityAccountRequest = {
  encode(message: GetLiabilityAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.liabilityAccountId !== 0) {
      writer.uint32(16).uint64(message.liabilityAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLiabilityAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLiabilityAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.liabilityAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLiabilityAccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      liabilityAccountId: isSet(object.liabilityAccountId) ? Number(object.liabilityAccountId) : 0,
    };
  },

  toJSON(message: GetLiabilityAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.liabilityAccountId !== undefined && (obj.liabilityAccountId = Math.round(message.liabilityAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLiabilityAccountRequest>, I>>(base?: I): GetLiabilityAccountRequest {
    return GetLiabilityAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLiabilityAccountRequest>, I>>(object: I): GetLiabilityAccountRequest {
    const message = createBaseGetLiabilityAccountRequest();
    message.userId = object.userId ?? 0;
    message.liabilityAccountId = object.liabilityAccountId ?? 0;
    return message;
  },
};

function createBaseGetLiabilityAccountResponse(): GetLiabilityAccountResponse {
  return { liabilityAccount: undefined };
}

export const GetLiabilityAccountResponse = {
  encode(message: GetLiabilityAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.liabilityAccount !== undefined) {
      CreditAccount.encode(message.liabilityAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLiabilityAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLiabilityAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.liabilityAccount = CreditAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLiabilityAccountResponse {
    return {
      liabilityAccount: isSet(object.liabilityAccount) ? CreditAccount.fromJSON(object.liabilityAccount) : undefined,
    };
  },

  toJSON(message: GetLiabilityAccountResponse): unknown {
    const obj: any = {};
    message.liabilityAccount !== undefined &&
      (obj.liabilityAccount = message.liabilityAccount ? CreditAccount.toJSON(message.liabilityAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLiabilityAccountResponse>, I>>(base?: I): GetLiabilityAccountResponse {
    return GetLiabilityAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLiabilityAccountResponse>, I>>(object: I): GetLiabilityAccountResponse {
    const message = createBaseGetLiabilityAccountResponse();
    message.liabilityAccount = (object.liabilityAccount !== undefined && object.liabilityAccount !== null)
      ? CreditAccount.fromPartial(object.liabilityAccount)
      : undefined;
    return message;
  },
};

function createBaseGetStudentLoanAccountRequest(): GetStudentLoanAccountRequest {
  return { userId: 0, studentLoanAccountId: 0 };
}

export const GetStudentLoanAccountRequest = {
  encode(message: GetStudentLoanAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.studentLoanAccountId !== 0) {
      writer.uint32(16).uint64(message.studentLoanAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetStudentLoanAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetStudentLoanAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.studentLoanAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetStudentLoanAccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      studentLoanAccountId: isSet(object.studentLoanAccountId) ? Number(object.studentLoanAccountId) : 0,
    };
  },

  toJSON(message: GetStudentLoanAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.studentLoanAccountId !== undefined && (obj.studentLoanAccountId = Math.round(message.studentLoanAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetStudentLoanAccountRequest>, I>>(base?: I): GetStudentLoanAccountRequest {
    return GetStudentLoanAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetStudentLoanAccountRequest>, I>>(object: I): GetStudentLoanAccountRequest {
    const message = createBaseGetStudentLoanAccountRequest();
    message.userId = object.userId ?? 0;
    message.studentLoanAccountId = object.studentLoanAccountId ?? 0;
    return message;
  },
};

function createBaseGetStudentLoanAccountResponse(): GetStudentLoanAccountResponse {
  return { studentLoanAccount: undefined };
}

export const GetStudentLoanAccountResponse = {
  encode(message: GetStudentLoanAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.studentLoanAccount !== undefined) {
      StudentLoanAccount.encode(message.studentLoanAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetStudentLoanAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetStudentLoanAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.studentLoanAccount = StudentLoanAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetStudentLoanAccountResponse {
    return {
      studentLoanAccount: isSet(object.studentLoanAccount)
        ? StudentLoanAccount.fromJSON(object.studentLoanAccount)
        : undefined,
    };
  },

  toJSON(message: GetStudentLoanAccountResponse): unknown {
    const obj: any = {};
    message.studentLoanAccount !== undefined && (obj.studentLoanAccount = message.studentLoanAccount
      ? StudentLoanAccount.toJSON(message.studentLoanAccount)
      : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetStudentLoanAccountResponse>, I>>(base?: I): GetStudentLoanAccountResponse {
    return GetStudentLoanAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetStudentLoanAccountResponse>, I>>(
    object: I,
  ): GetStudentLoanAccountResponse {
    const message = createBaseGetStudentLoanAccountResponse();
    message.studentLoanAccount = (object.studentLoanAccount !== undefined && object.studentLoanAccount !== null)
      ? StudentLoanAccount.fromPartial(object.studentLoanAccount)
      : undefined;
    return message;
  },
};

function createBaseCreateManualLinkRequest(): CreateManualLinkRequest {
  return { userId: 0, manualAccountLink: undefined };
}

export const CreateManualLinkRequest = {
  encode(message: CreateManualLinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.manualAccountLink !== undefined) {
      Link.encode(message.manualAccountLink, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateManualLinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateManualLinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.manualAccountLink = Link.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateManualLinkRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      manualAccountLink: isSet(object.manualAccountLink) ? Link.fromJSON(object.manualAccountLink) : undefined,
    };
  },

  toJSON(message: CreateManualLinkRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.manualAccountLink !== undefined &&
      (obj.manualAccountLink = message.manualAccountLink ? Link.toJSON(message.manualAccountLink) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateManualLinkRequest>, I>>(base?: I): CreateManualLinkRequest {
    return CreateManualLinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateManualLinkRequest>, I>>(object: I): CreateManualLinkRequest {
    const message = createBaseCreateManualLinkRequest();
    message.userId = object.userId ?? 0;
    message.manualAccountLink = (object.manualAccountLink !== undefined && object.manualAccountLink !== null)
      ? Link.fromPartial(object.manualAccountLink)
      : undefined;
    return message;
  },
};

function createBaseCreateManualLinkResponse(): CreateManualLinkResponse {
  return { linkId: 0 };
}

export const CreateManualLinkResponse = {
  encode(message: CreateManualLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.linkId !== 0) {
      writer.uint32(8).uint64(message.linkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateManualLinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateManualLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.linkId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateManualLinkResponse {
    return { linkId: isSet(object.linkId) ? Number(object.linkId) : 0 };
  },

  toJSON(message: CreateManualLinkResponse): unknown {
    const obj: any = {};
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateManualLinkResponse>, I>>(base?: I): CreateManualLinkResponse {
    return CreateManualLinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateManualLinkResponse>, I>>(object: I): CreateManualLinkResponse {
    const message = createBaseCreateManualLinkResponse();
    message.linkId = object.linkId ?? 0;
    return message;
  },
};

function createBaseGetLinkRequest(): GetLinkRequest {
  return { userId: 0, linkId: 0 };
}

export const GetLinkRequest = {
  encode(message: GetLinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.linkId !== 0) {
      writer.uint32(16).uint64(message.linkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.linkId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLinkRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      linkId: isSet(object.linkId) ? Number(object.linkId) : 0,
    };
  },

  toJSON(message: GetLinkRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLinkRequest>, I>>(base?: I): GetLinkRequest {
    return GetLinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLinkRequest>, I>>(object: I): GetLinkRequest {
    const message = createBaseGetLinkRequest();
    message.userId = object.userId ?? 0;
    message.linkId = object.linkId ?? 0;
    return message;
  },
};

function createBaseGetLinkResponse(): GetLinkResponse {
  return { link: undefined };
}

export const GetLinkResponse = {
  encode(message: GetLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.link !== undefined) {
      Link.encode(message.link, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.link = Link.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLinkResponse {
    return { link: isSet(object.link) ? Link.fromJSON(object.link) : undefined };
  },

  toJSON(message: GetLinkResponse): unknown {
    const obj: any = {};
    message.link !== undefined && (obj.link = message.link ? Link.toJSON(message.link) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLinkResponse>, I>>(base?: I): GetLinkResponse {
    return GetLinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLinkResponse>, I>>(object: I): GetLinkResponse {
    const message = createBaseGetLinkResponse();
    message.link = (object.link !== undefined && object.link !== null) ? Link.fromPartial(object.link) : undefined;
    return message;
  },
};

function createBaseGetLinksRequest(): GetLinksRequest {
  return { userId: 0 };
}

export const GetLinksRequest = {
  encode(message: GetLinksRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLinksRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLinksRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLinksRequest {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: GetLinksRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLinksRequest>, I>>(base?: I): GetLinksRequest {
    return GetLinksRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLinksRequest>, I>>(object: I): GetLinksRequest {
    const message = createBaseGetLinksRequest();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseGetLinksResponse(): GetLinksResponse {
  return { links: [] };
}

export const GetLinksResponse = {
  encode(message: GetLinksResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.links) {
      Link.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetLinksResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetLinksResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.links.push(Link.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetLinksResponse {
    return { links: Array.isArray(object?.links) ? object.links.map((e: any) => Link.fromJSON(e)) : [] };
  },

  toJSON(message: GetLinksResponse): unknown {
    const obj: any = {};
    if (message.links) {
      obj.links = message.links.map((e) => e ? Link.toJSON(e) : undefined);
    } else {
      obj.links = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetLinksResponse>, I>>(base?: I): GetLinksResponse {
    return GetLinksResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetLinksResponse>, I>>(object: I): GetLinksResponse {
    const message = createBaseGetLinksResponse();
    message.links = object.links?.map((e) => Link.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDeleteLinkRequest(): DeleteLinkRequest {
  return { userId: 0, linkId: 0 };
}

export const DeleteLinkRequest = {
  encode(message: DeleteLinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.linkId !== 0) {
      writer.uint32(16).uint64(message.linkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteLinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteLinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.linkId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteLinkRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      linkId: isSet(object.linkId) ? Number(object.linkId) : 0,
    };
  },

  toJSON(message: DeleteLinkRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteLinkRequest>, I>>(base?: I): DeleteLinkRequest {
    return DeleteLinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteLinkRequest>, I>>(object: I): DeleteLinkRequest {
    const message = createBaseDeleteLinkRequest();
    message.userId = object.userId ?? 0;
    message.linkId = object.linkId ?? 0;
    return message;
  },
};

function createBaseDeleteLinkResponse(): DeleteLinkResponse {
  return { linkId: 0 };
}

export const DeleteLinkResponse = {
  encode(message: DeleteLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.linkId !== 0) {
      writer.uint32(8).uint64(message.linkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteLinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.linkId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteLinkResponse {
    return { linkId: isSet(object.linkId) ? Number(object.linkId) : 0 };
  },

  toJSON(message: DeleteLinkResponse): unknown {
    const obj: any = {};
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteLinkResponse>, I>>(base?: I): DeleteLinkResponse {
    return DeleteLinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteLinkResponse>, I>>(object: I): DeleteLinkResponse {
    const message = createBaseDeleteLinkResponse();
    message.linkId = object.linkId ?? 0;
    return message;
  },
};

function createBaseGetReCurringTransactionsRequest(): GetReCurringTransactionsRequest {
  return { userId: 0 };
}

export const GetReCurringTransactionsRequest = {
  encode(message: GetReCurringTransactionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetReCurringTransactionsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetReCurringTransactionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetReCurringTransactionsRequest {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: GetReCurringTransactionsRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetReCurringTransactionsRequest>, I>>(base?: I): GetReCurringTransactionsRequest {
    return GetReCurringTransactionsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetReCurringTransactionsRequest>, I>>(
    object: I,
  ): GetReCurringTransactionsRequest {
    const message = createBaseGetReCurringTransactionsRequest();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseGetReCurringTransactionsResponse(): GetReCurringTransactionsResponse {
  return { reCcuringTransactions: [], participantReCcuringTransactions: [] };
}

export const GetReCurringTransactionsResponse = {
  encode(message: GetReCurringTransactionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.reCcuringTransactions) {
      ReOccuringTransaction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.participantReCcuringTransactions) {
      GetReCurringTransactionsResponse_ParticipantReCurringTransactions.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetReCurringTransactionsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetReCurringTransactionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.reCcuringTransactions.push(ReOccuringTransaction.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.participantReCcuringTransactions.push(
            GetReCurringTransactionsResponse_ParticipantReCurringTransactions.decode(reader, reader.uint32()),
          );
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetReCurringTransactionsResponse {
    return {
      reCcuringTransactions: Array.isArray(object?.reCcuringTransactions)
        ? object.reCcuringTransactions.map((e: any) => ReOccuringTransaction.fromJSON(e))
        : [],
      participantReCcuringTransactions: Array.isArray(object?.participantReCcuringTransactions)
        ? object.participantReCcuringTransactions.map((e: any) =>
          GetReCurringTransactionsResponse_ParticipantReCurringTransactions.fromJSON(e)
        )
        : [],
    };
  },

  toJSON(message: GetReCurringTransactionsResponse): unknown {
    const obj: any = {};
    if (message.reCcuringTransactions) {
      obj.reCcuringTransactions = message.reCcuringTransactions.map((e) =>
        e ? ReOccuringTransaction.toJSON(e) : undefined
      );
    } else {
      obj.reCcuringTransactions = [];
    }
    if (message.participantReCcuringTransactions) {
      obj.participantReCcuringTransactions = message.participantReCcuringTransactions.map((e) =>
        e ? GetReCurringTransactionsResponse_ParticipantReCurringTransactions.toJSON(e) : undefined
      );
    } else {
      obj.participantReCcuringTransactions = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetReCurringTransactionsResponse>, I>>(
    base?: I,
  ): GetReCurringTransactionsResponse {
    return GetReCurringTransactionsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetReCurringTransactionsResponse>, I>>(
    object: I,
  ): GetReCurringTransactionsResponse {
    const message = createBaseGetReCurringTransactionsResponse();
    message.reCcuringTransactions = object.reCcuringTransactions?.map((e) => ReOccuringTransaction.fromPartial(e)) ||
      [];
    message.participantReCcuringTransactions =
      object.participantReCcuringTransactions?.map((e) =>
        GetReCurringTransactionsResponse_ParticipantReCurringTransactions.fromPartial(e)
      ) || [];
    return message;
  },
};

function createBaseGetReCurringTransactionsResponse_ParticipantReCurringTransactions(): GetReCurringTransactionsResponse_ParticipantReCurringTransactions {
  return { reocurringTransactionId: 0, transactions: [] };
}

export const GetReCurringTransactionsResponse_ParticipantReCurringTransactions = {
  encode(
    message: GetReCurringTransactionsResponse_ParticipantReCurringTransactions,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.reocurringTransactionId !== 0) {
      writer.uint32(8).uint64(message.reocurringTransactionId);
    }
    for (const v of message.transactions) {
      Transaction.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): GetReCurringTransactionsResponse_ParticipantReCurringTransactions {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetReCurringTransactionsResponse_ParticipantReCurringTransactions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.reocurringTransactionId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.transactions.push(Transaction.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetReCurringTransactionsResponse_ParticipantReCurringTransactions {
    return {
      reocurringTransactionId: isSet(object.reocurringTransactionId) ? Number(object.reocurringTransactionId) : 0,
      transactions: Array.isArray(object?.transactions)
        ? object.transactions.map((e: any) => Transaction.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetReCurringTransactionsResponse_ParticipantReCurringTransactions): unknown {
    const obj: any = {};
    message.reocurringTransactionId !== undefined &&
      (obj.reocurringTransactionId = Math.round(message.reocurringTransactionId));
    if (message.transactions) {
      obj.transactions = message.transactions.map((e) => e ? Transaction.toJSON(e) : undefined);
    } else {
      obj.transactions = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetReCurringTransactionsResponse_ParticipantReCurringTransactions>, I>>(
    base?: I,
  ): GetReCurringTransactionsResponse_ParticipantReCurringTransactions {
    return GetReCurringTransactionsResponse_ParticipantReCurringTransactions.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetReCurringTransactionsResponse_ParticipantReCurringTransactions>, I>>(
    object: I,
  ): GetReCurringTransactionsResponse_ParticipantReCurringTransactions {
    const message = createBaseGetReCurringTransactionsResponse_ParticipantReCurringTransactions();
    message.reocurringTransactionId = object.reocurringTransactionId ?? 0;
    message.transactions = object.transactions?.map((e) => Transaction.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetTransactionsRequest(): GetTransactionsRequest {
  return { userId: 0, pageNumber: 0, pageSize: 0 };
}

export const GetTransactionsRequest = {
  encode(message: GetTransactionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).uint64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).uint64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTransactionsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTransactionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.uint64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTransactionsRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetTransactionsRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.pageNumber !== undefined && (obj.pageNumber = Math.round(message.pageNumber));
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTransactionsRequest>, I>>(base?: I): GetTransactionsRequest {
    return GetTransactionsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetTransactionsRequest>, I>>(object: I): GetTransactionsRequest {
    const message = createBaseGetTransactionsRequest();
    message.userId = object.userId ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetTransactionsResponse(): GetTransactionsResponse {
  return { transactions: [], nextPageNumber: 0 };
}

export const GetTransactionsResponse = {
  encode(message: GetTransactionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.transactions) {
      Transaction.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).uint64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTransactionsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTransactionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.transactions.push(Transaction.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTransactionsResponse {
    return {
      transactions: Array.isArray(object?.transactions)
        ? object.transactions.map((e: any) => Transaction.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetTransactionsResponse): unknown {
    const obj: any = {};
    if (message.transactions) {
      obj.transactions = message.transactions.map((e) => e ? Transaction.toJSON(e) : undefined);
    } else {
      obj.transactions = [];
    }
    message.nextPageNumber !== undefined && (obj.nextPageNumber = Math.round(message.nextPageNumber));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTransactionsResponse>, I>>(base?: I): GetTransactionsResponse {
    return GetTransactionsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetTransactionsResponse>, I>>(object: I): GetTransactionsResponse {
    const message = createBaseGetTransactionsResponse();
    message.transactions = object.transactions?.map((e) => Transaction.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseProcessWebhookRequest(): ProcessWebhookRequest {
  return {
    webhookType: "",
    webhookCode: "",
    itemId: "",
    initialUpdateComplete: false,
    historicalUpdateComplete: "",
    environment: "",
    newTransactions: [],
    removedTransactions: [],
    error: {},
    accountIds: [],
    consentExpirationTime: "",
    accountIdsWithNewLiabilities: [],
    accountIdsWithUpdatedLiabilities: [],
    newHoldings: 0,
    updatedHoldings: 0,
  };
}

export const ProcessWebhookRequest = {
  encode(message: ProcessWebhookRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.webhookType !== "") {
      writer.uint32(10).string(message.webhookType);
    }
    if (message.webhookCode !== "") {
      writer.uint32(18).string(message.webhookCode);
    }
    if (message.itemId !== "") {
      writer.uint32(26).string(message.itemId);
    }
    if (message.initialUpdateComplete === true) {
      writer.uint32(32).bool(message.initialUpdateComplete);
    }
    if (message.historicalUpdateComplete !== "") {
      writer.uint32(42).string(message.historicalUpdateComplete);
    }
    if (message.environment !== "") {
      writer.uint32(50).string(message.environment);
    }
    for (const v of message.newTransactions) {
      writer.uint32(58).string(v!);
    }
    for (const v of message.removedTransactions) {
      writer.uint32(66).string(v!);
    }
    Object.entries(message.error).forEach(([key, value]) => {
      ProcessWebhookRequest_ErrorEntry.encode({ key: key as any, value }, writer.uint32(74).fork()).ldelim();
    });
    for (const v of message.accountIds) {
      writer.uint32(82).string(v!);
    }
    if (message.consentExpirationTime !== "") {
      writer.uint32(90).string(message.consentExpirationTime);
    }
    for (const v of message.accountIdsWithNewLiabilities) {
      writer.uint32(98).string(v!);
    }
    for (const v of message.accountIdsWithUpdatedLiabilities) {
      writer.uint32(106).string(v!);
    }
    if (message.newHoldings !== 0) {
      writer.uint32(112).uint64(message.newHoldings);
    }
    if (message.updatedHoldings !== 0) {
      writer.uint32(120).uint64(message.updatedHoldings);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProcessWebhookRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProcessWebhookRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.webhookType = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.webhookCode = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.itemId = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.initialUpdateComplete = reader.bool();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.historicalUpdateComplete = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.environment = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.newTransactions.push(reader.string());
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.removedTransactions.push(reader.string());
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          const entry9 = ProcessWebhookRequest_ErrorEntry.decode(reader, reader.uint32());
          if (entry9.value !== undefined) {
            message.error[entry9.key] = entry9.value;
          }
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.accountIds.push(reader.string());
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.consentExpirationTime = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.accountIdsWithNewLiabilities.push(reader.string());
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.accountIdsWithUpdatedLiabilities.push(reader.string());
          continue;
        case 14:
          if (tag !== 112) {
            break;
          }

          message.newHoldings = longToNumber(reader.uint64() as Long);
          continue;
        case 15:
          if (tag !== 120) {
            break;
          }

          message.updatedHoldings = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ProcessWebhookRequest {
    return {
      webhookType: isSet(object.webhookType) ? String(object.webhookType) : "",
      webhookCode: isSet(object.webhookCode) ? String(object.webhookCode) : "",
      itemId: isSet(object.itemId) ? String(object.itemId) : "",
      initialUpdateComplete: isSet(object.initialUpdateComplete) ? Boolean(object.initialUpdateComplete) : false,
      historicalUpdateComplete: isSet(object.historicalUpdateComplete) ? String(object.historicalUpdateComplete) : "",
      environment: isSet(object.environment) ? String(object.environment) : "",
      newTransactions: Array.isArray(object?.newTransactions) ? object.newTransactions.map((e: any) => String(e)) : [],
      removedTransactions: Array.isArray(object?.removedTransactions)
        ? object.removedTransactions.map((e: any) => String(e))
        : [],
      error: isObject(object.error)
        ? Object.entries(object.error).reduce<{ [key: string]: Any }>((acc, [key, value]) => {
          acc[key] = Any.fromJSON(value);
          return acc;
        }, {})
        : {},
      accountIds: Array.isArray(object?.accountIds) ? object.accountIds.map((e: any) => String(e)) : [],
      consentExpirationTime: isSet(object.consentExpirationTime) ? String(object.consentExpirationTime) : "",
      accountIdsWithNewLiabilities: Array.isArray(object?.accountIdsWithNewLiabilities)
        ? object.accountIdsWithNewLiabilities.map((e: any) => String(e))
        : [],
      accountIdsWithUpdatedLiabilities: Array.isArray(object?.accountIdsWithUpdatedLiabilities)
        ? object.accountIdsWithUpdatedLiabilities.map((e: any) => String(e))
        : [],
      newHoldings: isSet(object.newHoldings) ? Number(object.newHoldings) : 0,
      updatedHoldings: isSet(object.updatedHoldings) ? Number(object.updatedHoldings) : 0,
    };
  },

  toJSON(message: ProcessWebhookRequest): unknown {
    const obj: any = {};
    message.webhookType !== undefined && (obj.webhookType = message.webhookType);
    message.webhookCode !== undefined && (obj.webhookCode = message.webhookCode);
    message.itemId !== undefined && (obj.itemId = message.itemId);
    message.initialUpdateComplete !== undefined && (obj.initialUpdateComplete = message.initialUpdateComplete);
    message.historicalUpdateComplete !== undefined && (obj.historicalUpdateComplete = message.historicalUpdateComplete);
    message.environment !== undefined && (obj.environment = message.environment);
    if (message.newTransactions) {
      obj.newTransactions = message.newTransactions.map((e) => e);
    } else {
      obj.newTransactions = [];
    }
    if (message.removedTransactions) {
      obj.removedTransactions = message.removedTransactions.map((e) => e);
    } else {
      obj.removedTransactions = [];
    }
    obj.error = {};
    if (message.error) {
      Object.entries(message.error).forEach(([k, v]) => {
        obj.error[k] = Any.toJSON(v);
      });
    }
    if (message.accountIds) {
      obj.accountIds = message.accountIds.map((e) => e);
    } else {
      obj.accountIds = [];
    }
    message.consentExpirationTime !== undefined && (obj.consentExpirationTime = message.consentExpirationTime);
    if (message.accountIdsWithNewLiabilities) {
      obj.accountIdsWithNewLiabilities = message.accountIdsWithNewLiabilities.map((e) => e);
    } else {
      obj.accountIdsWithNewLiabilities = [];
    }
    if (message.accountIdsWithUpdatedLiabilities) {
      obj.accountIdsWithUpdatedLiabilities = message.accountIdsWithUpdatedLiabilities.map((e) => e);
    } else {
      obj.accountIdsWithUpdatedLiabilities = [];
    }
    message.newHoldings !== undefined && (obj.newHoldings = Math.round(message.newHoldings));
    message.updatedHoldings !== undefined && (obj.updatedHoldings = Math.round(message.updatedHoldings));
    return obj;
  },

  create<I extends Exact<DeepPartial<ProcessWebhookRequest>, I>>(base?: I): ProcessWebhookRequest {
    return ProcessWebhookRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProcessWebhookRequest>, I>>(object: I): ProcessWebhookRequest {
    const message = createBaseProcessWebhookRequest();
    message.webhookType = object.webhookType ?? "";
    message.webhookCode = object.webhookCode ?? "";
    message.itemId = object.itemId ?? "";
    message.initialUpdateComplete = object.initialUpdateComplete ?? false;
    message.historicalUpdateComplete = object.historicalUpdateComplete ?? "";
    message.environment = object.environment ?? "";
    message.newTransactions = object.newTransactions?.map((e) => e) || [];
    message.removedTransactions = object.removedTransactions?.map((e) => e) || [];
    message.error = Object.entries(object.error ?? {}).reduce<{ [key: string]: Any }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Any.fromPartial(value);
      }
      return acc;
    }, {});
    message.accountIds = object.accountIds?.map((e) => e) || [];
    message.consentExpirationTime = object.consentExpirationTime ?? "";
    message.accountIdsWithNewLiabilities = object.accountIdsWithNewLiabilities?.map((e) => e) || [];
    message.accountIdsWithUpdatedLiabilities = object.accountIdsWithUpdatedLiabilities?.map((e) => e) || [];
    message.newHoldings = object.newHoldings ?? 0;
    message.updatedHoldings = object.updatedHoldings ?? 0;
    return message;
  },
};

function createBaseProcessWebhookRequest_ErrorEntry(): ProcessWebhookRequest_ErrorEntry {
  return { key: "", value: undefined };
}

export const ProcessWebhookRequest_ErrorEntry = {
  encode(message: ProcessWebhookRequest_ErrorEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Any.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProcessWebhookRequest_ErrorEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProcessWebhookRequest_ErrorEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = Any.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ProcessWebhookRequest_ErrorEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? Any.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: ProcessWebhookRequest_ErrorEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? Any.toJSON(message.value) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ProcessWebhookRequest_ErrorEntry>, I>>(
    base?: I,
  ): ProcessWebhookRequest_ErrorEntry {
    return ProcessWebhookRequest_ErrorEntry.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProcessWebhookRequest_ErrorEntry>, I>>(
    object: I,
  ): ProcessWebhookRequest_ErrorEntry {
    const message = createBaseProcessWebhookRequest_ErrorEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null) ? Any.fromPartial(object.value) : undefined;
    return message;
  },
};

function createBaseProcessWebhookResponse(): ProcessWebhookResponse {
  return {};
}

export const ProcessWebhookResponse = {
  encode(_: ProcessWebhookResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProcessWebhookResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProcessWebhookResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): ProcessWebhookResponse {
    return {};
  },

  toJSON(_: ProcessWebhookResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ProcessWebhookResponse>, I>>(base?: I): ProcessWebhookResponse {
    return ProcessWebhookResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ProcessWebhookResponse>, I>>(_: I): ProcessWebhookResponse {
    const message = createBaseProcessWebhookResponse();
    return message;
  },
};

function createBaseStripeWebhookRequest(): StripeWebhookRequest {
  return { body: "", signature: "" };
}

export const StripeWebhookRequest = {
  encode(message: StripeWebhookRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.body !== "") {
      writer.uint32(10).string(message.body);
    }
    if (message.signature !== "") {
      writer.uint32(18).string(message.signature);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StripeWebhookRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStripeWebhookRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.body = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.signature = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): StripeWebhookRequest {
    return {
      body: isSet(object.body) ? String(object.body) : "",
      signature: isSet(object.signature) ? String(object.signature) : "",
    };
  },

  toJSON(message: StripeWebhookRequest): unknown {
    const obj: any = {};
    message.body !== undefined && (obj.body = message.body);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  create<I extends Exact<DeepPartial<StripeWebhookRequest>, I>>(base?: I): StripeWebhookRequest {
    return StripeWebhookRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StripeWebhookRequest>, I>>(object: I): StripeWebhookRequest {
    const message = createBaseStripeWebhookRequest();
    message.body = object.body ?? "";
    message.signature = object.signature ?? "";
    return message;
  },
};

function createBaseStripeWebhookResponse(): StripeWebhookResponse {
  return { message: "" };
}

export const StripeWebhookResponse = {
  encode(message: StripeWebhookResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StripeWebhookResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStripeWebhookResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): StripeWebhookResponse {
    return { message: isSet(object.message) ? String(object.message) : "" };
  },

  toJSON(message: StripeWebhookResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  create<I extends Exact<DeepPartial<StripeWebhookResponse>, I>>(base?: I): StripeWebhookResponse {
    return StripeWebhookResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StripeWebhookResponse>, I>>(object: I): StripeWebhookResponse {
    const message = createBaseStripeWebhookResponse();
    message.message = object.message ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
