/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import {
  AccountBalanceHistory,
  CategoryMonthlyExpenditure,
  CategoryMonthlyIncome,
  CategoryMonthlyTransactionCount,
  DebtToIncomeRatio,
  ExpenseMetrics,
  FinancialProfile,
  IncomeExpenseRatio,
  IncomeMetrics,
  MelodyFinancialContext,
  MerchantMonthlyExpenditure,
  MonthlyBalance,
  MonthlyExpenditure,
  MonthlyIncome,
  MonthlySavings,
  MonthlyTotalQuantityBySecurityAndUser,
  MonthlyTransactionCount,
  PaymentChannelMonthlyExpenditure,
  TotalInvestmentBySecurity,
  TransactionAggregatesByMonth,
} from "./clickhouse_financial_service";
import Long = require("long");

export const protobufPackage = "financial_integration_service_api.v1";

/** transaction aggregates by month */
export interface GetTransactionAggregatesRequest {
  userId: number;
  month: number;
  personalFinanceCategoryPrimary: string;
  locationCity: string;
  paymentChannel: string;
  merchantName: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetTransactionAggregatesResponse {
  transactionAggregates: TransactionAggregatesByMonth[];
  nextPageNumber: number;
}

/** Account Balance */
export interface GetUserAccountBalanceHistoryRequest {
  /** User ID */
  userId: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetUserAccountBalanceHistoryResponse {
  /** List of account balance history records */
  accountBalanceHistory: AccountBalanceHistory[];
}

export interface GetAccountBalanceHistoryRequest {
  /** Account ID */
  plaidAccountId: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetAccountBalanceHistoryResponse {
  /** List of account balance history records for specific account */
  accountBalanceHistory: AccountBalanceHistory[];
}

/** CategoryMonthlyExpenditure */
export interface GetUserCategoryMonthlyExpenditureRequest {
  /** User ID */
  userId: number;
  /** Personal finance category */
  personalFinanceCategoryPrimary: string;
  /** Month in the format of YYYYMM */
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetUserCategoryMonthlyExpenditureResponse {
  /** List of CategoryMonthlyExpenditure records for the user */
  categoryMonthlyExpenditure: CategoryMonthlyExpenditure[];
  nextPageNumber: number;
}

export interface GetUserCategoryMonthlyIncomeRequest {
  userId: number;
  personalFinanceCategoryPrimary: string;
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetUserCategoryMonthlyIncomeResponse {
  categoryMonthlyIncome: CategoryMonthlyIncome[];
  nextPageNumber: number;
}

export interface GetCategoryMonthlyTransactionCountRequest {
  /** has to be present and defined */
  userId: number;
  /** optional */
  month: number;
  /** optional */
  personalFinanceCategoryPrimary: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetCategoryMonthlyTransactionCountResponse {
  categoryMonthlyTransactionCount: CategoryMonthlyTransactionCount[];
  nextPageNumber: number;
}

export interface GetDebtToIncomeRatioRequest {
  userId: number;
  /** optional */
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetDebtToIncomeRatioResponse {
  debtToIncomeRatios: DebtToIncomeRatio[];
  nextPageNumber: number;
}

export interface GetExpenseMetricsRequest {
  userId: number;
  /** optonal */
  month: number;
  /** optional */
  personalFinanceCategoryPrimary: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetExpenseMetricsResponse {
  expenseMetrics: ExpenseMetrics[];
  nextPageNumber: number;
}

/** GetFinancialProfile RPC */
export interface GetFinancialProfileRequest {
  userId: number;
  /** optional */
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetFinancialProfileResponse {
  financialProfiles: FinancialProfile[];
  nextPageNumber: number;
}

/** GetIncomeExpenseRatio RPC */
export interface GetIncomeExpenseRatioRequest {
  userId: number;
  /** optional */
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetIncomeExpenseRatioResponse {
  incomeExpenseRatios: IncomeExpenseRatio[];
  nextPageNumber: number;
}

/** GetIncomeMetrics RPC */
export interface GetIncomeMetricsRequest {
  userId: number;
  /** optional */
  month: number;
  /** optional */
  personalFinanceCategoryPrimary: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetIncomeMetricsResponse {
  incomeMetrics: IncomeMetrics[];
  nextPageNumber: number;
}

/** GetMerchantMonthlyExpenditure RPC */
export interface GetMerchantMonthlyExpenditureRequest {
  userId: number;
  /** optional */
  month: number;
  /** optional */
  merchantName: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMerchantMonthlyExpenditureResponse {
  merchantMonthlyExpenditures: MerchantMonthlyExpenditure[];
  nextPageNumber: number;
}

/** For example, for MonthlyBalance: */
export interface GetMonthlyBalanceRequest {
  userId: number;
  /** optional */
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMonthlyBalanceResponse {
  monthlyBalances: MonthlyBalance[];
  nextPageNumber: number;
}

/** GetMonthlyExpenditure RPC */
export interface GetMonthlyExpenditureRequest {
  userId: number;
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMonthlyExpenditureResponse {
  monthlyExpenditures: MonthlyExpenditure[];
  nextPageNumber: number;
}

/** GetMonthlyIncome RPC */
export interface GetMonthlyIncomeRequest {
  userId: number;
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMonthlyIncomeResponse {
  monthlyIncomes: MonthlyIncome[];
  nextPageNumber: number;
}

/** GetMonthlySavings RPC */
export interface GetMonthlySavingsRequest {
  userId: number;
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMonthlySavingsResponse {
  monthlySavings: MonthlySavings[];
  nextPageNumber: number;
}

/** GetMonthlyTotalQuantityBySecurityAndUser RPC */
export interface GetMonthlyTotalQuantityBySecurityAndUserRequest {
  userId: number;
  month: number;
  securityId: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMonthlyTotalQuantityBySecurityAndUserResponse {
  monthlyTotalQuantityBySecurityAndUser: MonthlyTotalQuantityBySecurityAndUser[];
  nextPageNumber: number;
}

/** GetMonthlyTransactionCount RPC */
export interface GetMonthlyTransactionCountRequest {
  userId: number;
  month: number;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetMonthlyTransactionCountResponse {
  monthlyTransactionCounts: MonthlyTransactionCount[];
  nextPageNumber: number;
}

/** GetPaymentChannelMonthlyExpenditure RPC */
export interface GetPaymentChannelMonthlyExpenditureRequest {
  userId: number;
  month: number;
  paymentChannel: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetPaymentChannelMonthlyExpenditureResponse {
  paymentChannelMonthlyExpenditure: PaymentChannelMonthlyExpenditure[];
  nextPageNumber: number;
}

/** GetTotalInvestmentBySecurity RPC */
export interface GetTotalInvestmentBySecurityRequest {
  userId: number;
  securityId: string;
  pageNumber: number;
  /** Number of items to return per page. */
  pageSize: number;
}

export interface GetTotalInvestmentBySecurityResponse {
  totalInvestmentBySecurity: TotalInvestmentBySecurity[];
  nextPageNumber: number;
}

export interface GetMelodyFinancialContextRequest {
  userId: number;
}

export interface GetMelodyFinancialContextResponse {
  melodyFinancialContext: MelodyFinancialContext | undefined;
}

function createBaseGetTransactionAggregatesRequest(): GetTransactionAggregatesRequest {
  return {
    userId: 0,
    month: 0,
    personalFinanceCategoryPrimary: "",
    locationCity: "",
    paymentChannel: "",
    merchantName: "",
    pageNumber: 0,
    pageSize: 0,
  };
}

export const GetTransactionAggregatesRequest = {
  encode(message: GetTransactionAggregatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(26).string(message.personalFinanceCategoryPrimary);
    }
    if (message.locationCity !== "") {
      writer.uint32(34).string(message.locationCity);
    }
    if (message.paymentChannel !== "") {
      writer.uint32(42).string(message.paymentChannel);
    }
    if (message.merchantName !== "") {
      writer.uint32(50).string(message.merchantName);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(56).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(64).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTransactionAggregatesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTransactionAggregatesRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.locationCity = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.paymentChannel = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.merchantName = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTransactionAggregatesRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      locationCity: isSet(object.locationCity) ? String(object.locationCity) : "",
      paymentChannel: isSet(object.paymentChannel) ? String(object.paymentChannel) : "",
      merchantName: isSet(object.merchantName) ? String(object.merchantName) : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetTransactionAggregatesRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary;
    }
    if (message.locationCity !== "") {
      obj.locationCity = message.locationCity;
    }
    if (message.paymentChannel !== "") {
      obj.paymentChannel = message.paymentChannel;
    }
    if (message.merchantName !== "") {
      obj.merchantName = message.merchantName;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTransactionAggregatesRequest>, I>>(base?: I): GetTransactionAggregatesRequest {
    return GetTransactionAggregatesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetTransactionAggregatesRequest>, I>>(
    object: I,
  ): GetTransactionAggregatesRequest {
    const message = createBaseGetTransactionAggregatesRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.locationCity = object.locationCity ?? "";
    message.paymentChannel = object.paymentChannel ?? "";
    message.merchantName = object.merchantName ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetTransactionAggregatesResponse(): GetTransactionAggregatesResponse {
  return { transactionAggregates: [], nextPageNumber: 0 };
}

export const GetTransactionAggregatesResponse = {
  encode(message: GetTransactionAggregatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.transactionAggregates) {
      TransactionAggregatesByMonth.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTransactionAggregatesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTransactionAggregatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.transactionAggregates.push(TransactionAggregatesByMonth.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTransactionAggregatesResponse {
    return {
      transactionAggregates: Array.isArray(object?.transactionAggregates)
        ? object.transactionAggregates.map((e: any) => TransactionAggregatesByMonth.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetTransactionAggregatesResponse): unknown {
    const obj: any = {};
    if (message.transactionAggregates?.length) {
      obj.transactionAggregates = message.transactionAggregates.map((e) => TransactionAggregatesByMonth.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTransactionAggregatesResponse>, I>>(
    base?: I,
  ): GetTransactionAggregatesResponse {
    return GetTransactionAggregatesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetTransactionAggregatesResponse>, I>>(
    object: I,
  ): GetTransactionAggregatesResponse {
    const message = createBaseGetTransactionAggregatesResponse();
    message.transactionAggregates =
      object.transactionAggregates?.map((e) => TransactionAggregatesByMonth.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetUserAccountBalanceHistoryRequest(): GetUserAccountBalanceHistoryRequest {
  return { userId: 0, pageNumber: 0, pageSize: 0 };
}

export const GetUserAccountBalanceHistoryRequest = {
  encode(message: GetUserAccountBalanceHistoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(16).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(24).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserAccountBalanceHistoryRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserAccountBalanceHistoryRequest();
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

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserAccountBalanceHistoryRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetUserAccountBalanceHistoryRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserAccountBalanceHistoryRequest>, I>>(
    base?: I,
  ): GetUserAccountBalanceHistoryRequest {
    return GetUserAccountBalanceHistoryRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserAccountBalanceHistoryRequest>, I>>(
    object: I,
  ): GetUserAccountBalanceHistoryRequest {
    const message = createBaseGetUserAccountBalanceHistoryRequest();
    message.userId = object.userId ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetUserAccountBalanceHistoryResponse(): GetUserAccountBalanceHistoryResponse {
  return { accountBalanceHistory: [] };
}

export const GetUserAccountBalanceHistoryResponse = {
  encode(message: GetUserAccountBalanceHistoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.accountBalanceHistory) {
      AccountBalanceHistory.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserAccountBalanceHistoryResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserAccountBalanceHistoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accountBalanceHistory.push(AccountBalanceHistory.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserAccountBalanceHistoryResponse {
    return {
      accountBalanceHistory: Array.isArray(object?.accountBalanceHistory)
        ? object.accountBalanceHistory.map((e: any) => AccountBalanceHistory.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetUserAccountBalanceHistoryResponse): unknown {
    const obj: any = {};
    if (message.accountBalanceHistory?.length) {
      obj.accountBalanceHistory = message.accountBalanceHistory.map((e) => AccountBalanceHistory.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserAccountBalanceHistoryResponse>, I>>(
    base?: I,
  ): GetUserAccountBalanceHistoryResponse {
    return GetUserAccountBalanceHistoryResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserAccountBalanceHistoryResponse>, I>>(
    object: I,
  ): GetUserAccountBalanceHistoryResponse {
    const message = createBaseGetUserAccountBalanceHistoryResponse();
    message.accountBalanceHistory = object.accountBalanceHistory?.map((e) => AccountBalanceHistory.fromPartial(e)) ||
      [];
    return message;
  },
};

function createBaseGetAccountBalanceHistoryRequest(): GetAccountBalanceHistoryRequest {
  return { plaidAccountId: "", pageNumber: 0, pageSize: 0 };
}

export const GetAccountBalanceHistoryRequest = {
  encode(message: GetAccountBalanceHistoryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.plaidAccountId !== "") {
      writer.uint32(10).string(message.plaidAccountId);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(16).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(24).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAccountBalanceHistoryRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAccountBalanceHistoryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetAccountBalanceHistoryRequest {
    return {
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetAccountBalanceHistoryRequest): unknown {
    const obj: any = {};
    if (message.plaidAccountId !== "") {
      obj.plaidAccountId = message.plaidAccountId;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetAccountBalanceHistoryRequest>, I>>(base?: I): GetAccountBalanceHistoryRequest {
    return GetAccountBalanceHistoryRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetAccountBalanceHistoryRequest>, I>>(
    object: I,
  ): GetAccountBalanceHistoryRequest {
    const message = createBaseGetAccountBalanceHistoryRequest();
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetAccountBalanceHistoryResponse(): GetAccountBalanceHistoryResponse {
  return { accountBalanceHistory: [] };
}

export const GetAccountBalanceHistoryResponse = {
  encode(message: GetAccountBalanceHistoryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.accountBalanceHistory) {
      AccountBalanceHistory.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAccountBalanceHistoryResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAccountBalanceHistoryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accountBalanceHistory.push(AccountBalanceHistory.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetAccountBalanceHistoryResponse {
    return {
      accountBalanceHistory: Array.isArray(object?.accountBalanceHistory)
        ? object.accountBalanceHistory.map((e: any) => AccountBalanceHistory.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetAccountBalanceHistoryResponse): unknown {
    const obj: any = {};
    if (message.accountBalanceHistory?.length) {
      obj.accountBalanceHistory = message.accountBalanceHistory.map((e) => AccountBalanceHistory.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetAccountBalanceHistoryResponse>, I>>(
    base?: I,
  ): GetAccountBalanceHistoryResponse {
    return GetAccountBalanceHistoryResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetAccountBalanceHistoryResponse>, I>>(
    object: I,
  ): GetAccountBalanceHistoryResponse {
    const message = createBaseGetAccountBalanceHistoryResponse();
    message.accountBalanceHistory = object.accountBalanceHistory?.map((e) => AccountBalanceHistory.fromPartial(e)) ||
      [];
    return message;
  },
};

function createBaseGetUserCategoryMonthlyExpenditureRequest(): GetUserCategoryMonthlyExpenditureRequest {
  return { userId: 0, personalFinanceCategoryPrimary: "", month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetUserCategoryMonthlyExpenditureRequest = {
  encode(message: GetUserCategoryMonthlyExpenditureRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(18).string(message.personalFinanceCategoryPrimary);
    }
    if (message.month !== 0) {
      writer.uint32(24).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserCategoryMonthlyExpenditureRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserCategoryMonthlyExpenditureRequest();
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

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.month = reader.uint32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserCategoryMonthlyExpenditureRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetUserCategoryMonthlyExpenditureRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary;
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserCategoryMonthlyExpenditureRequest>, I>>(
    base?: I,
  ): GetUserCategoryMonthlyExpenditureRequest {
    return GetUserCategoryMonthlyExpenditureRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserCategoryMonthlyExpenditureRequest>, I>>(
    object: I,
  ): GetUserCategoryMonthlyExpenditureRequest {
    const message = createBaseGetUserCategoryMonthlyExpenditureRequest();
    message.userId = object.userId ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetUserCategoryMonthlyExpenditureResponse(): GetUserCategoryMonthlyExpenditureResponse {
  return { categoryMonthlyExpenditure: [], nextPageNumber: 0 };
}

export const GetUserCategoryMonthlyExpenditureResponse = {
  encode(message: GetUserCategoryMonthlyExpenditureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.categoryMonthlyExpenditure) {
      CategoryMonthlyExpenditure.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserCategoryMonthlyExpenditureResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserCategoryMonthlyExpenditureResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.categoryMonthlyExpenditure.push(CategoryMonthlyExpenditure.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserCategoryMonthlyExpenditureResponse {
    return {
      categoryMonthlyExpenditure: Array.isArray(object?.categoryMonthlyExpenditure)
        ? object.categoryMonthlyExpenditure.map((e: any) => CategoryMonthlyExpenditure.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetUserCategoryMonthlyExpenditureResponse): unknown {
    const obj: any = {};
    if (message.categoryMonthlyExpenditure?.length) {
      obj.categoryMonthlyExpenditure = message.categoryMonthlyExpenditure.map((e) =>
        CategoryMonthlyExpenditure.toJSON(e)
      );
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserCategoryMonthlyExpenditureResponse>, I>>(
    base?: I,
  ): GetUserCategoryMonthlyExpenditureResponse {
    return GetUserCategoryMonthlyExpenditureResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserCategoryMonthlyExpenditureResponse>, I>>(
    object: I,
  ): GetUserCategoryMonthlyExpenditureResponse {
    const message = createBaseGetUserCategoryMonthlyExpenditureResponse();
    message.categoryMonthlyExpenditure =
      object.categoryMonthlyExpenditure?.map((e) => CategoryMonthlyExpenditure.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetUserCategoryMonthlyIncomeRequest(): GetUserCategoryMonthlyIncomeRequest {
  return { userId: 0, personalFinanceCategoryPrimary: "", month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetUserCategoryMonthlyIncomeRequest = {
  encode(message: GetUserCategoryMonthlyIncomeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(18).string(message.personalFinanceCategoryPrimary);
    }
    if (message.month !== 0) {
      writer.uint32(24).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserCategoryMonthlyIncomeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserCategoryMonthlyIncomeRequest();
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

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.month = reader.uint32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserCategoryMonthlyIncomeRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetUserCategoryMonthlyIncomeRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary;
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserCategoryMonthlyIncomeRequest>, I>>(
    base?: I,
  ): GetUserCategoryMonthlyIncomeRequest {
    return GetUserCategoryMonthlyIncomeRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserCategoryMonthlyIncomeRequest>, I>>(
    object: I,
  ): GetUserCategoryMonthlyIncomeRequest {
    const message = createBaseGetUserCategoryMonthlyIncomeRequest();
    message.userId = object.userId ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetUserCategoryMonthlyIncomeResponse(): GetUserCategoryMonthlyIncomeResponse {
  return { categoryMonthlyIncome: [], nextPageNumber: 0 };
}

export const GetUserCategoryMonthlyIncomeResponse = {
  encode(message: GetUserCategoryMonthlyIncomeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.categoryMonthlyIncome) {
      CategoryMonthlyIncome.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserCategoryMonthlyIncomeResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserCategoryMonthlyIncomeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.categoryMonthlyIncome.push(CategoryMonthlyIncome.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserCategoryMonthlyIncomeResponse {
    return {
      categoryMonthlyIncome: Array.isArray(object?.categoryMonthlyIncome)
        ? object.categoryMonthlyIncome.map((e: any) => CategoryMonthlyIncome.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetUserCategoryMonthlyIncomeResponse): unknown {
    const obj: any = {};
    if (message.categoryMonthlyIncome?.length) {
      obj.categoryMonthlyIncome = message.categoryMonthlyIncome.map((e) => CategoryMonthlyIncome.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserCategoryMonthlyIncomeResponse>, I>>(
    base?: I,
  ): GetUserCategoryMonthlyIncomeResponse {
    return GetUserCategoryMonthlyIncomeResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserCategoryMonthlyIncomeResponse>, I>>(
    object: I,
  ): GetUserCategoryMonthlyIncomeResponse {
    const message = createBaseGetUserCategoryMonthlyIncomeResponse();
    message.categoryMonthlyIncome = object.categoryMonthlyIncome?.map((e) => CategoryMonthlyIncome.fromPartial(e)) ||
      [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetCategoryMonthlyTransactionCountRequest(): GetCategoryMonthlyTransactionCountRequest {
  return { userId: 0, month: 0, personalFinanceCategoryPrimary: "", pageNumber: 0, pageSize: 0 };
}

export const GetCategoryMonthlyTransactionCountRequest = {
  encode(message: GetCategoryMonthlyTransactionCountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(26).string(message.personalFinanceCategoryPrimary);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetCategoryMonthlyTransactionCountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetCategoryMonthlyTransactionCountRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetCategoryMonthlyTransactionCountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetCategoryMonthlyTransactionCountRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetCategoryMonthlyTransactionCountRequest>, I>>(
    base?: I,
  ): GetCategoryMonthlyTransactionCountRequest {
    return GetCategoryMonthlyTransactionCountRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetCategoryMonthlyTransactionCountRequest>, I>>(
    object: I,
  ): GetCategoryMonthlyTransactionCountRequest {
    const message = createBaseGetCategoryMonthlyTransactionCountRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetCategoryMonthlyTransactionCountResponse(): GetCategoryMonthlyTransactionCountResponse {
  return { categoryMonthlyTransactionCount: [], nextPageNumber: 0 };
}

export const GetCategoryMonthlyTransactionCountResponse = {
  encode(message: GetCategoryMonthlyTransactionCountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.categoryMonthlyTransactionCount) {
      CategoryMonthlyTransactionCount.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetCategoryMonthlyTransactionCountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetCategoryMonthlyTransactionCountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.categoryMonthlyTransactionCount.push(CategoryMonthlyTransactionCount.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetCategoryMonthlyTransactionCountResponse {
    return {
      categoryMonthlyTransactionCount: Array.isArray(object?.categoryMonthlyTransactionCount)
        ? object.categoryMonthlyTransactionCount.map((e: any) => CategoryMonthlyTransactionCount.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetCategoryMonthlyTransactionCountResponse): unknown {
    const obj: any = {};
    if (message.categoryMonthlyTransactionCount?.length) {
      obj.categoryMonthlyTransactionCount = message.categoryMonthlyTransactionCount.map((e) =>
        CategoryMonthlyTransactionCount.toJSON(e)
      );
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetCategoryMonthlyTransactionCountResponse>, I>>(
    base?: I,
  ): GetCategoryMonthlyTransactionCountResponse {
    return GetCategoryMonthlyTransactionCountResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetCategoryMonthlyTransactionCountResponse>, I>>(
    object: I,
  ): GetCategoryMonthlyTransactionCountResponse {
    const message = createBaseGetCategoryMonthlyTransactionCountResponse();
    message.categoryMonthlyTransactionCount =
      object.categoryMonthlyTransactionCount?.map((e) => CategoryMonthlyTransactionCount.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetDebtToIncomeRatioRequest(): GetDebtToIncomeRatioRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetDebtToIncomeRatioRequest = {
  encode(message: GetDebtToIncomeRatioRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDebtToIncomeRatioRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetDebtToIncomeRatioRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetDebtToIncomeRatioRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetDebtToIncomeRatioRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetDebtToIncomeRatioRequest>, I>>(base?: I): GetDebtToIncomeRatioRequest {
    return GetDebtToIncomeRatioRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetDebtToIncomeRatioRequest>, I>>(object: I): GetDebtToIncomeRatioRequest {
    const message = createBaseGetDebtToIncomeRatioRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetDebtToIncomeRatioResponse(): GetDebtToIncomeRatioResponse {
  return { debtToIncomeRatios: [], nextPageNumber: 0 };
}

export const GetDebtToIncomeRatioResponse = {
  encode(message: GetDebtToIncomeRatioResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.debtToIncomeRatios) {
      DebtToIncomeRatio.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDebtToIncomeRatioResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetDebtToIncomeRatioResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.debtToIncomeRatios.push(DebtToIncomeRatio.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetDebtToIncomeRatioResponse {
    return {
      debtToIncomeRatios: Array.isArray(object?.debtToIncomeRatios)
        ? object.debtToIncomeRatios.map((e: any) => DebtToIncomeRatio.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetDebtToIncomeRatioResponse): unknown {
    const obj: any = {};
    if (message.debtToIncomeRatios?.length) {
      obj.debtToIncomeRatios = message.debtToIncomeRatios.map((e) => DebtToIncomeRatio.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetDebtToIncomeRatioResponse>, I>>(base?: I): GetDebtToIncomeRatioResponse {
    return GetDebtToIncomeRatioResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetDebtToIncomeRatioResponse>, I>>(object: I): GetDebtToIncomeRatioResponse {
    const message = createBaseGetDebtToIncomeRatioResponse();
    message.debtToIncomeRatios = object.debtToIncomeRatios?.map((e) => DebtToIncomeRatio.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetExpenseMetricsRequest(): GetExpenseMetricsRequest {
  return { userId: 0, month: 0, personalFinanceCategoryPrimary: "", pageNumber: 0, pageSize: 0 };
}

export const GetExpenseMetricsRequest = {
  encode(message: GetExpenseMetricsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(26).string(message.personalFinanceCategoryPrimary);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetExpenseMetricsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetExpenseMetricsRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetExpenseMetricsRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetExpenseMetricsRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetExpenseMetricsRequest>, I>>(base?: I): GetExpenseMetricsRequest {
    return GetExpenseMetricsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetExpenseMetricsRequest>, I>>(object: I): GetExpenseMetricsRequest {
    const message = createBaseGetExpenseMetricsRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetExpenseMetricsResponse(): GetExpenseMetricsResponse {
  return { expenseMetrics: [], nextPageNumber: 0 };
}

export const GetExpenseMetricsResponse = {
  encode(message: GetExpenseMetricsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.expenseMetrics) {
      ExpenseMetrics.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetExpenseMetricsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetExpenseMetricsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.expenseMetrics.push(ExpenseMetrics.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetExpenseMetricsResponse {
    return {
      expenseMetrics: Array.isArray(object?.expenseMetrics)
        ? object.expenseMetrics.map((e: any) => ExpenseMetrics.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetExpenseMetricsResponse): unknown {
    const obj: any = {};
    if (message.expenseMetrics?.length) {
      obj.expenseMetrics = message.expenseMetrics.map((e) => ExpenseMetrics.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetExpenseMetricsResponse>, I>>(base?: I): GetExpenseMetricsResponse {
    return GetExpenseMetricsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetExpenseMetricsResponse>, I>>(object: I): GetExpenseMetricsResponse {
    const message = createBaseGetExpenseMetricsResponse();
    message.expenseMetrics = object.expenseMetrics?.map((e) => ExpenseMetrics.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetFinancialProfileRequest(): GetFinancialProfileRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetFinancialProfileRequest = {
  encode(message: GetFinancialProfileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetFinancialProfileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetFinancialProfileRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetFinancialProfileRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetFinancialProfileRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetFinancialProfileRequest>, I>>(base?: I): GetFinancialProfileRequest {
    return GetFinancialProfileRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetFinancialProfileRequest>, I>>(object: I): GetFinancialProfileRequest {
    const message = createBaseGetFinancialProfileRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetFinancialProfileResponse(): GetFinancialProfileResponse {
  return { financialProfiles: [], nextPageNumber: 0 };
}

export const GetFinancialProfileResponse = {
  encode(message: GetFinancialProfileResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.financialProfiles) {
      FinancialProfile.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetFinancialProfileResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetFinancialProfileResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.financialProfiles.push(FinancialProfile.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetFinancialProfileResponse {
    return {
      financialProfiles: Array.isArray(object?.financialProfiles)
        ? object.financialProfiles.map((e: any) => FinancialProfile.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetFinancialProfileResponse): unknown {
    const obj: any = {};
    if (message.financialProfiles?.length) {
      obj.financialProfiles = message.financialProfiles.map((e) => FinancialProfile.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetFinancialProfileResponse>, I>>(base?: I): GetFinancialProfileResponse {
    return GetFinancialProfileResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetFinancialProfileResponse>, I>>(object: I): GetFinancialProfileResponse {
    const message = createBaseGetFinancialProfileResponse();
    message.financialProfiles = object.financialProfiles?.map((e) => FinancialProfile.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetIncomeExpenseRatioRequest(): GetIncomeExpenseRatioRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetIncomeExpenseRatioRequest = {
  encode(message: GetIncomeExpenseRatioRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetIncomeExpenseRatioRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetIncomeExpenseRatioRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetIncomeExpenseRatioRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetIncomeExpenseRatioRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetIncomeExpenseRatioRequest>, I>>(base?: I): GetIncomeExpenseRatioRequest {
    return GetIncomeExpenseRatioRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetIncomeExpenseRatioRequest>, I>>(object: I): GetIncomeExpenseRatioRequest {
    const message = createBaseGetIncomeExpenseRatioRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetIncomeExpenseRatioResponse(): GetIncomeExpenseRatioResponse {
  return { incomeExpenseRatios: [], nextPageNumber: 0 };
}

export const GetIncomeExpenseRatioResponse = {
  encode(message: GetIncomeExpenseRatioResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.incomeExpenseRatios) {
      IncomeExpenseRatio.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetIncomeExpenseRatioResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetIncomeExpenseRatioResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.incomeExpenseRatios.push(IncomeExpenseRatio.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetIncomeExpenseRatioResponse {
    return {
      incomeExpenseRatios: Array.isArray(object?.incomeExpenseRatios)
        ? object.incomeExpenseRatios.map((e: any) => IncomeExpenseRatio.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetIncomeExpenseRatioResponse): unknown {
    const obj: any = {};
    if (message.incomeExpenseRatios?.length) {
      obj.incomeExpenseRatios = message.incomeExpenseRatios.map((e) => IncomeExpenseRatio.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetIncomeExpenseRatioResponse>, I>>(base?: I): GetIncomeExpenseRatioResponse {
    return GetIncomeExpenseRatioResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetIncomeExpenseRatioResponse>, I>>(
    object: I,
  ): GetIncomeExpenseRatioResponse {
    const message = createBaseGetIncomeExpenseRatioResponse();
    message.incomeExpenseRatios = object.incomeExpenseRatios?.map((e) => IncomeExpenseRatio.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetIncomeMetricsRequest(): GetIncomeMetricsRequest {
  return { userId: 0, month: 0, personalFinanceCategoryPrimary: "", pageNumber: 0, pageSize: 0 };
}

export const GetIncomeMetricsRequest = {
  encode(message: GetIncomeMetricsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(26).string(message.personalFinanceCategoryPrimary);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetIncomeMetricsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetIncomeMetricsRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetIncomeMetricsRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetIncomeMetricsRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetIncomeMetricsRequest>, I>>(base?: I): GetIncomeMetricsRequest {
    return GetIncomeMetricsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetIncomeMetricsRequest>, I>>(object: I): GetIncomeMetricsRequest {
    const message = createBaseGetIncomeMetricsRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetIncomeMetricsResponse(): GetIncomeMetricsResponse {
  return { incomeMetrics: [], nextPageNumber: 0 };
}

export const GetIncomeMetricsResponse = {
  encode(message: GetIncomeMetricsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.incomeMetrics) {
      IncomeMetrics.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetIncomeMetricsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetIncomeMetricsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.incomeMetrics.push(IncomeMetrics.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetIncomeMetricsResponse {
    return {
      incomeMetrics: Array.isArray(object?.incomeMetrics)
        ? object.incomeMetrics.map((e: any) => IncomeMetrics.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetIncomeMetricsResponse): unknown {
    const obj: any = {};
    if (message.incomeMetrics?.length) {
      obj.incomeMetrics = message.incomeMetrics.map((e) => IncomeMetrics.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetIncomeMetricsResponse>, I>>(base?: I): GetIncomeMetricsResponse {
    return GetIncomeMetricsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetIncomeMetricsResponse>, I>>(object: I): GetIncomeMetricsResponse {
    const message = createBaseGetIncomeMetricsResponse();
    message.incomeMetrics = object.incomeMetrics?.map((e) => IncomeMetrics.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMerchantMonthlyExpenditureRequest(): GetMerchantMonthlyExpenditureRequest {
  return { userId: 0, month: 0, merchantName: "", pageNumber: 0, pageSize: 0 };
}

export const GetMerchantMonthlyExpenditureRequest = {
  encode(message: GetMerchantMonthlyExpenditureRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.merchantName !== "") {
      writer.uint32(26).string(message.merchantName);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMerchantMonthlyExpenditureRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMerchantMonthlyExpenditureRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.merchantName = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMerchantMonthlyExpenditureRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      merchantName: isSet(object.merchantName) ? String(object.merchantName) : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMerchantMonthlyExpenditureRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.merchantName !== "") {
      obj.merchantName = message.merchantName;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMerchantMonthlyExpenditureRequest>, I>>(
    base?: I,
  ): GetMerchantMonthlyExpenditureRequest {
    return GetMerchantMonthlyExpenditureRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMerchantMonthlyExpenditureRequest>, I>>(
    object: I,
  ): GetMerchantMonthlyExpenditureRequest {
    const message = createBaseGetMerchantMonthlyExpenditureRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.merchantName = object.merchantName ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMerchantMonthlyExpenditureResponse(): GetMerchantMonthlyExpenditureResponse {
  return { merchantMonthlyExpenditures: [], nextPageNumber: 0 };
}

export const GetMerchantMonthlyExpenditureResponse = {
  encode(message: GetMerchantMonthlyExpenditureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.merchantMonthlyExpenditures) {
      MerchantMonthlyExpenditure.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMerchantMonthlyExpenditureResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMerchantMonthlyExpenditureResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.merchantMonthlyExpenditures.push(MerchantMonthlyExpenditure.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMerchantMonthlyExpenditureResponse {
    return {
      merchantMonthlyExpenditures: Array.isArray(object?.merchantMonthlyExpenditures)
        ? object.merchantMonthlyExpenditures.map((e: any) => MerchantMonthlyExpenditure.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMerchantMonthlyExpenditureResponse): unknown {
    const obj: any = {};
    if (message.merchantMonthlyExpenditures?.length) {
      obj.merchantMonthlyExpenditures = message.merchantMonthlyExpenditures.map((e) =>
        MerchantMonthlyExpenditure.toJSON(e)
      );
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMerchantMonthlyExpenditureResponse>, I>>(
    base?: I,
  ): GetMerchantMonthlyExpenditureResponse {
    return GetMerchantMonthlyExpenditureResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMerchantMonthlyExpenditureResponse>, I>>(
    object: I,
  ): GetMerchantMonthlyExpenditureResponse {
    const message = createBaseGetMerchantMonthlyExpenditureResponse();
    message.merchantMonthlyExpenditures =
      object.merchantMonthlyExpenditures?.map((e) => MerchantMonthlyExpenditure.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMonthlyBalanceRequest(): GetMonthlyBalanceRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetMonthlyBalanceRequest = {
  encode(message: GetMonthlyBalanceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyBalanceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyBalanceRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyBalanceRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMonthlyBalanceRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyBalanceRequest>, I>>(base?: I): GetMonthlyBalanceRequest {
    return GetMonthlyBalanceRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyBalanceRequest>, I>>(object: I): GetMonthlyBalanceRequest {
    const message = createBaseGetMonthlyBalanceRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMonthlyBalanceResponse(): GetMonthlyBalanceResponse {
  return { monthlyBalances: [], nextPageNumber: 0 };
}

export const GetMonthlyBalanceResponse = {
  encode(message: GetMonthlyBalanceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.monthlyBalances) {
      MonthlyBalance.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyBalanceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyBalanceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.monthlyBalances.push(MonthlyBalance.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyBalanceResponse {
    return {
      monthlyBalances: Array.isArray(object?.monthlyBalances)
        ? object.monthlyBalances.map((e: any) => MonthlyBalance.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMonthlyBalanceResponse): unknown {
    const obj: any = {};
    if (message.monthlyBalances?.length) {
      obj.monthlyBalances = message.monthlyBalances.map((e) => MonthlyBalance.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyBalanceResponse>, I>>(base?: I): GetMonthlyBalanceResponse {
    return GetMonthlyBalanceResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyBalanceResponse>, I>>(object: I): GetMonthlyBalanceResponse {
    const message = createBaseGetMonthlyBalanceResponse();
    message.monthlyBalances = object.monthlyBalances?.map((e) => MonthlyBalance.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMonthlyExpenditureRequest(): GetMonthlyExpenditureRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetMonthlyExpenditureRequest = {
  encode(message: GetMonthlyExpenditureRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyExpenditureRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyExpenditureRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyExpenditureRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMonthlyExpenditureRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyExpenditureRequest>, I>>(base?: I): GetMonthlyExpenditureRequest {
    return GetMonthlyExpenditureRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyExpenditureRequest>, I>>(object: I): GetMonthlyExpenditureRequest {
    const message = createBaseGetMonthlyExpenditureRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMonthlyExpenditureResponse(): GetMonthlyExpenditureResponse {
  return { monthlyExpenditures: [], nextPageNumber: 0 };
}

export const GetMonthlyExpenditureResponse = {
  encode(message: GetMonthlyExpenditureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.monthlyExpenditures) {
      MonthlyExpenditure.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyExpenditureResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyExpenditureResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.monthlyExpenditures.push(MonthlyExpenditure.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyExpenditureResponse {
    return {
      monthlyExpenditures: Array.isArray(object?.monthlyExpenditures)
        ? object.monthlyExpenditures.map((e: any) => MonthlyExpenditure.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMonthlyExpenditureResponse): unknown {
    const obj: any = {};
    if (message.monthlyExpenditures?.length) {
      obj.monthlyExpenditures = message.monthlyExpenditures.map((e) => MonthlyExpenditure.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyExpenditureResponse>, I>>(base?: I): GetMonthlyExpenditureResponse {
    return GetMonthlyExpenditureResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyExpenditureResponse>, I>>(
    object: I,
  ): GetMonthlyExpenditureResponse {
    const message = createBaseGetMonthlyExpenditureResponse();
    message.monthlyExpenditures = object.monthlyExpenditures?.map((e) => MonthlyExpenditure.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMonthlyIncomeRequest(): GetMonthlyIncomeRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetMonthlyIncomeRequest = {
  encode(message: GetMonthlyIncomeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyIncomeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyIncomeRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyIncomeRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMonthlyIncomeRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyIncomeRequest>, I>>(base?: I): GetMonthlyIncomeRequest {
    return GetMonthlyIncomeRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyIncomeRequest>, I>>(object: I): GetMonthlyIncomeRequest {
    const message = createBaseGetMonthlyIncomeRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMonthlyIncomeResponse(): GetMonthlyIncomeResponse {
  return { monthlyIncomes: [], nextPageNumber: 0 };
}

export const GetMonthlyIncomeResponse = {
  encode(message: GetMonthlyIncomeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.monthlyIncomes) {
      MonthlyIncome.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyIncomeResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyIncomeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.monthlyIncomes.push(MonthlyIncome.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyIncomeResponse {
    return {
      monthlyIncomes: Array.isArray(object?.monthlyIncomes)
        ? object.monthlyIncomes.map((e: any) => MonthlyIncome.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMonthlyIncomeResponse): unknown {
    const obj: any = {};
    if (message.monthlyIncomes?.length) {
      obj.monthlyIncomes = message.monthlyIncomes.map((e) => MonthlyIncome.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyIncomeResponse>, I>>(base?: I): GetMonthlyIncomeResponse {
    return GetMonthlyIncomeResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyIncomeResponse>, I>>(object: I): GetMonthlyIncomeResponse {
    const message = createBaseGetMonthlyIncomeResponse();
    message.monthlyIncomes = object.monthlyIncomes?.map((e) => MonthlyIncome.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMonthlySavingsRequest(): GetMonthlySavingsRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetMonthlySavingsRequest = {
  encode(message: GetMonthlySavingsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(24).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlySavingsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlySavingsRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlySavingsRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMonthlySavingsRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlySavingsRequest>, I>>(base?: I): GetMonthlySavingsRequest {
    return GetMonthlySavingsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlySavingsRequest>, I>>(object: I): GetMonthlySavingsRequest {
    const message = createBaseGetMonthlySavingsRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMonthlySavingsResponse(): GetMonthlySavingsResponse {
  return { monthlySavings: [], nextPageNumber: 0 };
}

export const GetMonthlySavingsResponse = {
  encode(message: GetMonthlySavingsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.monthlySavings) {
      MonthlySavings.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlySavingsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlySavingsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.monthlySavings.push(MonthlySavings.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlySavingsResponse {
    return {
      monthlySavings: Array.isArray(object?.monthlySavings)
        ? object.monthlySavings.map((e: any) => MonthlySavings.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMonthlySavingsResponse): unknown {
    const obj: any = {};
    if (message.monthlySavings?.length) {
      obj.monthlySavings = message.monthlySavings.map((e) => MonthlySavings.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlySavingsResponse>, I>>(base?: I): GetMonthlySavingsResponse {
    return GetMonthlySavingsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlySavingsResponse>, I>>(object: I): GetMonthlySavingsResponse {
    const message = createBaseGetMonthlySavingsResponse();
    message.monthlySavings = object.monthlySavings?.map((e) => MonthlySavings.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMonthlyTotalQuantityBySecurityAndUserRequest(): GetMonthlyTotalQuantityBySecurityAndUserRequest {
  return { userId: 0, month: 0, securityId: "", pageNumber: 0, pageSize: 0 };
}

export const GetMonthlyTotalQuantityBySecurityAndUserRequest = {
  encode(
    message: GetMonthlyTotalQuantityBySecurityAndUserRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.securityId !== "") {
      writer.uint32(26).string(message.securityId);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyTotalQuantityBySecurityAndUserRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyTotalQuantityBySecurityAndUserRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.securityId = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyTotalQuantityBySecurityAndUserRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      securityId: isSet(object.securityId) ? String(object.securityId) : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMonthlyTotalQuantityBySecurityAndUserRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.securityId !== "") {
      obj.securityId = message.securityId;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyTotalQuantityBySecurityAndUserRequest>, I>>(
    base?: I,
  ): GetMonthlyTotalQuantityBySecurityAndUserRequest {
    return GetMonthlyTotalQuantityBySecurityAndUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyTotalQuantityBySecurityAndUserRequest>, I>>(
    object: I,
  ): GetMonthlyTotalQuantityBySecurityAndUserRequest {
    const message = createBaseGetMonthlyTotalQuantityBySecurityAndUserRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.securityId = object.securityId ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMonthlyTotalQuantityBySecurityAndUserResponse(): GetMonthlyTotalQuantityBySecurityAndUserResponse {
  return { monthlyTotalQuantityBySecurityAndUser: [], nextPageNumber: 0 };
}

export const GetMonthlyTotalQuantityBySecurityAndUserResponse = {
  encode(
    message: GetMonthlyTotalQuantityBySecurityAndUserResponse,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    for (const v of message.monthlyTotalQuantityBySecurityAndUser) {
      MonthlyTotalQuantityBySecurityAndUser.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyTotalQuantityBySecurityAndUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyTotalQuantityBySecurityAndUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.monthlyTotalQuantityBySecurityAndUser.push(
            MonthlyTotalQuantityBySecurityAndUser.decode(reader, reader.uint32()),
          );
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyTotalQuantityBySecurityAndUserResponse {
    return {
      monthlyTotalQuantityBySecurityAndUser: Array.isArray(object?.monthlyTotalQuantityBySecurityAndUser)
        ? object.monthlyTotalQuantityBySecurityAndUser.map((e: any) =>
          MonthlyTotalQuantityBySecurityAndUser.fromJSON(e)
        )
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMonthlyTotalQuantityBySecurityAndUserResponse): unknown {
    const obj: any = {};
    if (message.monthlyTotalQuantityBySecurityAndUser?.length) {
      obj.monthlyTotalQuantityBySecurityAndUser = message.monthlyTotalQuantityBySecurityAndUser.map((e) =>
        MonthlyTotalQuantityBySecurityAndUser.toJSON(e)
      );
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyTotalQuantityBySecurityAndUserResponse>, I>>(
    base?: I,
  ): GetMonthlyTotalQuantityBySecurityAndUserResponse {
    return GetMonthlyTotalQuantityBySecurityAndUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyTotalQuantityBySecurityAndUserResponse>, I>>(
    object: I,
  ): GetMonthlyTotalQuantityBySecurityAndUserResponse {
    const message = createBaseGetMonthlyTotalQuantityBySecurityAndUserResponse();
    message.monthlyTotalQuantityBySecurityAndUser =
      object.monthlyTotalQuantityBySecurityAndUser?.map((e) => MonthlyTotalQuantityBySecurityAndUser.fromPartial(e)) ||
      [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMonthlyTransactionCountRequest(): GetMonthlyTransactionCountRequest {
  return { userId: 0, month: 0, pageNumber: 0, pageSize: 0 };
}

export const GetMonthlyTransactionCountRequest = {
  encode(message: GetMonthlyTransactionCountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyTransactionCountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyTransactionCountRequest();
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

          message.month = reader.uint32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyTransactionCountRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetMonthlyTransactionCountRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyTransactionCountRequest>, I>>(
    base?: I,
  ): GetMonthlyTransactionCountRequest {
    return GetMonthlyTransactionCountRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyTransactionCountRequest>, I>>(
    object: I,
  ): GetMonthlyTransactionCountRequest {
    const message = createBaseGetMonthlyTransactionCountRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetMonthlyTransactionCountResponse(): GetMonthlyTransactionCountResponse {
  return { monthlyTransactionCounts: [], nextPageNumber: 0 };
}

export const GetMonthlyTransactionCountResponse = {
  encode(message: GetMonthlyTransactionCountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.monthlyTransactionCounts) {
      MonthlyTransactionCount.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMonthlyTransactionCountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMonthlyTransactionCountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.monthlyTransactionCounts.push(MonthlyTransactionCount.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMonthlyTransactionCountResponse {
    return {
      monthlyTransactionCounts: Array.isArray(object?.monthlyTransactionCounts)
        ? object.monthlyTransactionCounts.map((e: any) => MonthlyTransactionCount.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetMonthlyTransactionCountResponse): unknown {
    const obj: any = {};
    if (message.monthlyTransactionCounts?.length) {
      obj.monthlyTransactionCounts = message.monthlyTransactionCounts.map((e) => MonthlyTransactionCount.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMonthlyTransactionCountResponse>, I>>(
    base?: I,
  ): GetMonthlyTransactionCountResponse {
    return GetMonthlyTransactionCountResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMonthlyTransactionCountResponse>, I>>(
    object: I,
  ): GetMonthlyTransactionCountResponse {
    const message = createBaseGetMonthlyTransactionCountResponse();
    message.monthlyTransactionCounts =
      object.monthlyTransactionCounts?.map((e) => MonthlyTransactionCount.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetPaymentChannelMonthlyExpenditureRequest(): GetPaymentChannelMonthlyExpenditureRequest {
  return { userId: 0, month: 0, paymentChannel: "", pageNumber: 0, pageSize: 0 };
}

export const GetPaymentChannelMonthlyExpenditureRequest = {
  encode(message: GetPaymentChannelMonthlyExpenditureRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.month !== 0) {
      writer.uint32(16).uint32(message.month);
    }
    if (message.paymentChannel !== "") {
      writer.uint32(26).string(message.paymentChannel);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPaymentChannelMonthlyExpenditureRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPaymentChannelMonthlyExpenditureRequest();
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

          message.month = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.paymentChannel = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPaymentChannelMonthlyExpenditureRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      month: isSet(object.month) ? Number(object.month) : 0,
      paymentChannel: isSet(object.paymentChannel) ? String(object.paymentChannel) : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetPaymentChannelMonthlyExpenditureRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.month !== 0) {
      obj.month = Math.round(message.month);
    }
    if (message.paymentChannel !== "") {
      obj.paymentChannel = message.paymentChannel;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPaymentChannelMonthlyExpenditureRequest>, I>>(
    base?: I,
  ): GetPaymentChannelMonthlyExpenditureRequest {
    return GetPaymentChannelMonthlyExpenditureRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetPaymentChannelMonthlyExpenditureRequest>, I>>(
    object: I,
  ): GetPaymentChannelMonthlyExpenditureRequest {
    const message = createBaseGetPaymentChannelMonthlyExpenditureRequest();
    message.userId = object.userId ?? 0;
    message.month = object.month ?? 0;
    message.paymentChannel = object.paymentChannel ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetPaymentChannelMonthlyExpenditureResponse(): GetPaymentChannelMonthlyExpenditureResponse {
  return { paymentChannelMonthlyExpenditure: [], nextPageNumber: 0 };
}

export const GetPaymentChannelMonthlyExpenditureResponse = {
  encode(message: GetPaymentChannelMonthlyExpenditureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.paymentChannelMonthlyExpenditure) {
      PaymentChannelMonthlyExpenditure.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPaymentChannelMonthlyExpenditureResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPaymentChannelMonthlyExpenditureResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.paymentChannelMonthlyExpenditure.push(
            PaymentChannelMonthlyExpenditure.decode(reader, reader.uint32()),
          );
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPaymentChannelMonthlyExpenditureResponse {
    return {
      paymentChannelMonthlyExpenditure: Array.isArray(object?.paymentChannelMonthlyExpenditure)
        ? object.paymentChannelMonthlyExpenditure.map((e: any) => PaymentChannelMonthlyExpenditure.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetPaymentChannelMonthlyExpenditureResponse): unknown {
    const obj: any = {};
    if (message.paymentChannelMonthlyExpenditure?.length) {
      obj.paymentChannelMonthlyExpenditure = message.paymentChannelMonthlyExpenditure.map((e) =>
        PaymentChannelMonthlyExpenditure.toJSON(e)
      );
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPaymentChannelMonthlyExpenditureResponse>, I>>(
    base?: I,
  ): GetPaymentChannelMonthlyExpenditureResponse {
    return GetPaymentChannelMonthlyExpenditureResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetPaymentChannelMonthlyExpenditureResponse>, I>>(
    object: I,
  ): GetPaymentChannelMonthlyExpenditureResponse {
    const message = createBaseGetPaymentChannelMonthlyExpenditureResponse();
    message.paymentChannelMonthlyExpenditure =
      object.paymentChannelMonthlyExpenditure?.map((e) => PaymentChannelMonthlyExpenditure.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetTotalInvestmentBySecurityRequest(): GetTotalInvestmentBySecurityRequest {
  return { userId: 0, securityId: "", pageNumber: 0, pageSize: 0 };
}

export const GetTotalInvestmentBySecurityRequest = {
  encode(message: GetTotalInvestmentBySecurityRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    if (message.securityId !== "") {
      writer.uint32(18).string(message.securityId);
    }
    if (message.pageNumber !== 0) {
      writer.uint32(32).int64(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      writer.uint32(40).int64(message.pageSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTotalInvestmentBySecurityRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTotalInvestmentBySecurityRequest();
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

          message.securityId = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageNumber = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.pageSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTotalInvestmentBySecurityRequest {
    return {
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      securityId: isSet(object.securityId) ? String(object.securityId) : "",
      pageNumber: isSet(object.pageNumber) ? Number(object.pageNumber) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
    };
  },

  toJSON(message: GetTotalInvestmentBySecurityRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    if (message.securityId !== "") {
      obj.securityId = message.securityId;
    }
    if (message.pageNumber !== 0) {
      obj.pageNumber = Math.round(message.pageNumber);
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTotalInvestmentBySecurityRequest>, I>>(
    base?: I,
  ): GetTotalInvestmentBySecurityRequest {
    return GetTotalInvestmentBySecurityRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetTotalInvestmentBySecurityRequest>, I>>(
    object: I,
  ): GetTotalInvestmentBySecurityRequest {
    const message = createBaseGetTotalInvestmentBySecurityRequest();
    message.userId = object.userId ?? 0;
    message.securityId = object.securityId ?? "";
    message.pageNumber = object.pageNumber ?? 0;
    message.pageSize = object.pageSize ?? 0;
    return message;
  },
};

function createBaseGetTotalInvestmentBySecurityResponse(): GetTotalInvestmentBySecurityResponse {
  return { totalInvestmentBySecurity: [], nextPageNumber: 0 };
}

export const GetTotalInvestmentBySecurityResponse = {
  encode(message: GetTotalInvestmentBySecurityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.totalInvestmentBySecurity) {
      TotalInvestmentBySecurity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageNumber !== 0) {
      writer.uint32(16).int64(message.nextPageNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTotalInvestmentBySecurityResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTotalInvestmentBySecurityResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.totalInvestmentBySecurity.push(TotalInvestmentBySecurity.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.nextPageNumber = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTotalInvestmentBySecurityResponse {
    return {
      totalInvestmentBySecurity: Array.isArray(object?.totalInvestmentBySecurity)
        ? object.totalInvestmentBySecurity.map((e: any) => TotalInvestmentBySecurity.fromJSON(e))
        : [],
      nextPageNumber: isSet(object.nextPageNumber) ? Number(object.nextPageNumber) : 0,
    };
  },

  toJSON(message: GetTotalInvestmentBySecurityResponse): unknown {
    const obj: any = {};
    if (message.totalInvestmentBySecurity?.length) {
      obj.totalInvestmentBySecurity = message.totalInvestmentBySecurity.map((e) => TotalInvestmentBySecurity.toJSON(e));
    }
    if (message.nextPageNumber !== 0) {
      obj.nextPageNumber = Math.round(message.nextPageNumber);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTotalInvestmentBySecurityResponse>, I>>(
    base?: I,
  ): GetTotalInvestmentBySecurityResponse {
    return GetTotalInvestmentBySecurityResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetTotalInvestmentBySecurityResponse>, I>>(
    object: I,
  ): GetTotalInvestmentBySecurityResponse {
    const message = createBaseGetTotalInvestmentBySecurityResponse();
    message.totalInvestmentBySecurity =
      object.totalInvestmentBySecurity?.map((e) => TotalInvestmentBySecurity.fromPartial(e)) || [];
    message.nextPageNumber = object.nextPageNumber ?? 0;
    return message;
  },
};

function createBaseGetMelodyFinancialContextRequest(): GetMelodyFinancialContextRequest {
  return { userId: 0 };
}

export const GetMelodyFinancialContextRequest = {
  encode(message: GetMelodyFinancialContextRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== 0) {
      writer.uint32(8).uint64(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMelodyFinancialContextRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMelodyFinancialContextRequest();
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

  fromJSON(object: any): GetMelodyFinancialContextRequest {
    return { userId: isSet(object.userId) ? Number(object.userId) : 0 };
  },

  toJSON(message: GetMelodyFinancialContextRequest): unknown {
    const obj: any = {};
    if (message.userId !== 0) {
      obj.userId = Math.round(message.userId);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMelodyFinancialContextRequest>, I>>(
    base?: I,
  ): GetMelodyFinancialContextRequest {
    return GetMelodyFinancialContextRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMelodyFinancialContextRequest>, I>>(
    object: I,
  ): GetMelodyFinancialContextRequest {
    const message = createBaseGetMelodyFinancialContextRequest();
    message.userId = object.userId ?? 0;
    return message;
  },
};

function createBaseGetMelodyFinancialContextResponse(): GetMelodyFinancialContextResponse {
  return { melodyFinancialContext: undefined };
}

export const GetMelodyFinancialContextResponse = {
  encode(message: GetMelodyFinancialContextResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.melodyFinancialContext !== undefined) {
      MelodyFinancialContext.encode(message.melodyFinancialContext, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMelodyFinancialContextResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMelodyFinancialContextResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.melodyFinancialContext = MelodyFinancialContext.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMelodyFinancialContextResponse {
    return {
      melodyFinancialContext: isSet(object.melodyFinancialContext)
        ? MelodyFinancialContext.fromJSON(object.melodyFinancialContext)
        : undefined,
    };
  },

  toJSON(message: GetMelodyFinancialContextResponse): unknown {
    const obj: any = {};
    if (message.melodyFinancialContext !== undefined) {
      obj.melodyFinancialContext = MelodyFinancialContext.toJSON(message.melodyFinancialContext);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMelodyFinancialContextResponse>, I>>(
    base?: I,
  ): GetMelodyFinancialContextResponse {
    return GetMelodyFinancialContextResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetMelodyFinancialContextResponse>, I>>(
    object: I,
  ): GetMelodyFinancialContextResponse {
    const message = createBaseGetMelodyFinancialContextResponse();
    message.melodyFinancialContext =
      (object.melodyFinancialContext !== undefined && object.melodyFinancialContext !== null)
        ? MelodyFinancialContext.fromPartial(object.melodyFinancialContext)
        : undefined;
    return message;
  },
};

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
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

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
