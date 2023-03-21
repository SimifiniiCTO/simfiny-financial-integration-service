/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import {
  BankAccount,
  Budget,
  CreditAccount,
  Forecast,
  InvestmentAccount,
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

export interface GetMortageAccountRequest {
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
  mortageAccountId: number;
}

export interface GetMortageAccountResponse {
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
          if (tag != 10) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.profileDeleted = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.profileUpdated = reader.bool();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.profile = UserProfile.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.bankAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.bankAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.bankAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 26) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.updated = reader.bool();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.bankAccount = BankAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.pocket = Pocket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.smartGoals.push(SmartGoal.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.smartGoal = SmartGoal.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 26) {
            break;
          }

          message.smartGoal = SmartGoal.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 18) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.milestones.push(Milestone.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.milestone = Milestone.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.forecast = Forecast.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.milestroneId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.budgetId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 18) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 16) {
            break;
          }

          message.budgetId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.deleted = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.budgetId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.budget = Budget.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.pocketId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.smartGoalId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.milestoneId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.budgets.push(Budget.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.healthy = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.healthy = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.fullName = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.email = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.phoneNumber = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.linkToken = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.expiration = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.plaidRequestId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.publicToken = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.institutionId = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.institutionName = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.investmentAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.investmentAccount = InvestmentAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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

function createBaseGetMortageAccountRequest(): GetMortageAccountRequest {
  return { userId: 0, mortageAccountId: 0 };
}

export const GetMortageAccountRequest = {
  encode(message: GetMortageAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.mortageAccountId !== 0) {
      writer.uint32(16).uint64(message.mortageAccountId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMortageAccountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMortageAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.mortageAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMortageAccountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      mortageAccountId: isSet(object.mortageAccountId) ? Number(object.mortageAccountId) : 0,
    };
  },

  toJSON(message: GetMortageAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.mortageAccountId !== undefined && (obj.mortageAccountId = Math.round(message.mortageAccountId));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMortageAccountRequest>, I>>(base?: I): GetMortageAccountRequest {
    return GetMortageAccountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMortageAccountRequest>, I>>(object: I): GetMortageAccountRequest {
    const message = createBaseGetMortageAccountRequest();
    message.userId = object.userId ?? 0;
    message.mortageAccountId = object.mortageAccountId ?? 0;
    return message;
  },
};

function createBaseGetMortageAccountResponse(): GetMortageAccountResponse {
  return { mortageAccount: undefined };
}

export const GetMortageAccountResponse = {
  encode(message: GetMortageAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mortageAccount !== undefined) {
      MortgageAccount.encode(message.mortageAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMortageAccountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMortageAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.mortageAccount = MortgageAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMortageAccountResponse {
    return {
      mortageAccount: isSet(object.mortageAccount) ? MortgageAccount.fromJSON(object.mortageAccount) : undefined,
    };
  },

  toJSON(message: GetMortageAccountResponse): unknown {
    const obj: any = {};
    message.mortageAccount !== undefined &&
      (obj.mortageAccount = message.mortageAccount ? MortgageAccount.toJSON(message.mortageAccount) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMortageAccountResponse>, I>>(base?: I): GetMortageAccountResponse {
    return GetMortageAccountResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMortageAccountResponse>, I>>(object: I): GetMortageAccountResponse {
    const message = createBaseGetMortageAccountResponse();
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.liabilityAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.liabilityAccount = CreditAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 8) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.studentLoanAccountId = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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
          if (tag != 10) {
            break;
          }

          message.studentLoanAccount = StudentLoanAccount.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
