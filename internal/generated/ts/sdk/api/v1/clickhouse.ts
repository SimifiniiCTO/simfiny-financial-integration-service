/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "api.v1";

export enum ReOccuringTransactionsFrequency {
  RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED = 0,
  RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY = 1,
  RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY = 2,
  RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY = 3,
  RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY = 4,
  RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY = 5,
  UNRECOGNIZED = -1,
}

export function reOccuringTransactionsFrequencyFromJSON(object: any): ReOccuringTransactionsFrequency {
  switch (object) {
    case 0:
    case "RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED":
      return ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED;
    case 1:
    case "RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY":
      return ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY;
    case 2:
    case "RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY":
      return ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY;
    case 3:
    case "RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY":
      return ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY;
    case 4:
    case "RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY":
      return ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY;
    case 5:
    case "RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY":
      return ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ReOccuringTransactionsFrequency.UNRECOGNIZED;
  }
}

export function reOccuringTransactionsFrequencyToJSON(object: ReOccuringTransactionsFrequency): string {
  switch (object) {
    case ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED:
      return "RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED";
    case ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY:
      return "RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY";
    case ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY:
      return "RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY";
    case ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY:
      return "RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY";
    case ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY:
      return "RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY";
    case ReOccuringTransactionsFrequency.RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY:
      return "RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY";
    case ReOccuringTransactionsFrequency.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum ReOccuringTransactionsStatus {
  RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED = 0,
  /**
   * RE_OCCURING_TRANSACTIONS_STATUS_MATURE - A MATURE recurring stream should have at least 3 transactions and happen
   *  on a regular cadence (For Annual recurring stream, we will mark it MATURE after 2 instances).
   */
  RE_OCCURING_TRANSACTIONS_STATUS_MATURE = 1,
  /**
   * RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION - When a recurring transaction first appears in the transaction history and before it fulfills
   * the requirement of a mature stream, the status will be EARLY_DETECTION.
   */
  RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION = 2,
  /**
   * RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED - A stream that was previously in the EARLY_DETECTION status will move to the TOMBSTONED
   * status when no further transactions were found at the next expected date.
   */
  RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED = 3,
  UNRECOGNIZED = -1,
}

export function reOccuringTransactionsStatusFromJSON(object: any): ReOccuringTransactionsStatus {
  switch (object) {
    case 0:
    case "RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED":
      return ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED;
    case 1:
    case "RE_OCCURING_TRANSACTIONS_STATUS_MATURE":
      return ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_MATURE;
    case 2:
    case "RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION":
      return ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION;
    case 3:
    case "RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED":
      return ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ReOccuringTransactionsStatus.UNRECOGNIZED;
  }
}

export function reOccuringTransactionsStatusToJSON(object: ReOccuringTransactionsStatus): string {
  switch (object) {
    case ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED:
      return "RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED";
    case ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_MATURE:
      return "RE_OCCURING_TRANSACTIONS_STATUS_MATURE";
    case ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION:
      return "RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION";
    case ReOccuringTransactionsStatus.RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED:
      return "RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED";
    case ReOccuringTransactionsStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum ReCurringFlow {
  RE_CURRING_FLOW_UNSPECIFIED = 0,
  RE_CURRING_FLOW_INFLOW = 1,
  RE_CURRING_FLOW_OUTFLOW = 2,
  UNRECOGNIZED = -1,
}

export function reCurringFlowFromJSON(object: any): ReCurringFlow {
  switch (object) {
    case 0:
    case "RE_CURRING_FLOW_UNSPECIFIED":
      return ReCurringFlow.RE_CURRING_FLOW_UNSPECIFIED;
    case 1:
    case "RE_CURRING_FLOW_INFLOW":
      return ReCurringFlow.RE_CURRING_FLOW_INFLOW;
    case 2:
    case "RE_CURRING_FLOW_OUTFLOW":
      return ReCurringFlow.RE_CURRING_FLOW_OUTFLOW;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ReCurringFlow.UNRECOGNIZED;
  }
}

export function reCurringFlowToJSON(object: ReCurringFlow): string {
  switch (object) {
    case ReCurringFlow.RE_CURRING_FLOW_UNSPECIFIED:
      return "RE_CURRING_FLOW_UNSPECIFIED";
    case ReCurringFlow.RE_CURRING_FLOW_INFLOW:
      return "RE_CURRING_FLOW_INFLOW";
    case ReCurringFlow.RE_CURRING_FLOW_OUTFLOW:
      return "RE_CURRING_FLOW_OUTFLOW";
    case ReCurringFlow.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface PersonalFinanceCategory {
  primary: string;
  detailed: string;
}

export interface ReOccuringTransaction {
  /** @gotag: clickhouse:"account_id" */
  accountId: string;
  /** @gotag: clickhouse:"stream_id" */
  streamId: string;
  /** @gotag: clickhouse:"category" */
  category: string[];
  /** @gotag: clickhouse:"category_id" */
  categoryId: string;
  /** @gotag: clickhouse:"description" */
  description: string;
  /** @gotag: clickhouse:"merchant_name" */
  merchantName: string;
  /** @gotag: clickhouse:"personal_finance_category_primary" */
  personalFinanceCategoryPrimary: string;
  /** @gotag: clickhouse:"personal_finance_category_detailed" */
  personalFinanceCategoryDetailed: string;
  /** @gotag: clickhouse:"first_date" */
  firstDate: string;
  /** @gotag: clickhouse:"last_date" */
  lastDate: string;
  /** @gotag: clickhouse:"frequency" */
  frequency: ReOccuringTransactionsFrequency;
  /** @gotag: clickhouse:"transaction_ids" */
  transactionIds: string[];
  /** @gotag: clickhouse:"average_amount" */
  averageAmount: string;
  /** @gotag: clickhouse:"average_amount_iso_currency_code" */
  averageAmountIsoCurrencyCode: string;
  /** @gotag: clickhouse:"last_amount" */
  lastAmount: string;
  /** @gotag: clickhouse:"last_amount_iso_currency_code" */
  lastAmountIsoCurrencyCode: string;
  /** @gotag: clickhouse:"is_active" */
  isActive: boolean;
  /** @gotag: clickhouse:"status" */
  status: ReOccuringTransactionsStatus;
  /** @gotag: clickhouse:"updated_time" */
  updatedTime: string;
  /** @gotag: clickhouse:"user_id" */
  userId: number;
  /** @gotag: clickhouse:"link_id" */
  linkId: number;
  /** @gotag: clickhouse:"id" */
  id: number;
  /** @gotag: clickhouse:"flow" */
  flow: ReCurringFlow;
}

export interface Transaction {
  /** @gotag: clickhouse:"account_id" */
  accountId: string;
  /** @gotag: clickhouse:"amount" */
  amount: number;
  /** @gotag: clickhouse:"iso_currency_code" */
  isoCurrencyCode: string;
  /** @gotag: clickhouse:"unofficial_currency_code" */
  unofficialCurrencyCode: string;
  /** @gotag: clickhouse:"category" */
  category: string;
  /** @gotag: clickhouse:"category_id" */
  categoryId: string;
  /** @gotag: clickhouse:"check_number" */
  checkNumber: string;
  /** @gotag: clickhouse:"date" */
  date: string;
  /** @gotag: clickhouse:"datetime" */
  datetime: string;
  /** @gotag: clickhouse:"authorized_date" */
  authorizedDate: string;
  /** @gotag: clickhouse:"authorized_datetime" */
  authorizedDatetime: string;
  /** @gotag: clickhouse:"location" */
  location:
    | Transaction_Location
    | undefined;
  /** @gotag: clickhouse:"name" */
  name: string;
  /** @gotag: clickhouse:"merchant_name" */
  merchantName: string;
  /** @gotag: clickhouse:"payment_meta" */
  paymentMeta:
    | Transaction_PaymentMeta
    | undefined;
  /** @gotag: clickhouse:"payment_channel" */
  paymentChannel: string;
  /** @gotag: clickhouse:"pending" */
  pending: boolean;
  /** @gotag: clickhouse:"pending_transaction_id" */
  pendingTransactionId: string;
  /** @gotag: clickhouse:"account_owner" */
  accountOwner: string;
  /** @gotag: clickhouse:"transaction_id" */
  transactionId: string;
  /** @gotag: clickhouse:"transaction_code" */
  transactionCode: string;
  id: number;
  /** @gotag: clickhouse:"user_id" */
  userId: number;
  /** @gotag: clickhouse:"link_id" */
  linkId: number;
}

export interface Transaction_Location {
  address: string;
  city: string;
  region: string;
  postalCode: string;
  country: string;
  lat: number;
  lon: number;
  storeNumber: string;
}

export interface Transaction_PaymentMeta {
  byOrderOf: string;
  payee: string;
  payer: string;
  paymentMethod: string;
  paymentProcessor: string;
  ppdId: string;
  reason: string;
  referenceNumber: string;
}

export interface TransactionAmountByCountryMetric {
  country: string;
  amount: number;
}

export interface AverageTransactionAmountByCategoryMetric {
  category: string;
  amount: number;
}

export interface MonthlyTransactionCountByCategoryMetric {
  category: string;
  count: number;
  month: string;
}

export interface TransactionCountByMerchantPaymentChannelMetric {
  merchantName: string;
  paymentChannel: string;
  transactionCount: number;
}

export interface TransactionAmountDistributionByCategoryMetric {
  category: string;
  mean: number;
  median: number;
  standardDeviation: number;
}

function createBasePersonalFinanceCategory(): PersonalFinanceCategory {
  return { primary: "", detailed: "" };
}

export const PersonalFinanceCategory = {
  encode(message: PersonalFinanceCategory, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.primary !== "") {
      writer.uint32(10).string(message.primary);
    }
    if (message.detailed !== "") {
      writer.uint32(18).string(message.detailed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PersonalFinanceCategory {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePersonalFinanceCategory();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.primary = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.detailed = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PersonalFinanceCategory {
    return {
      primary: isSet(object.primary) ? String(object.primary) : "",
      detailed: isSet(object.detailed) ? String(object.detailed) : "",
    };
  },

  toJSON(message: PersonalFinanceCategory): unknown {
    const obj: any = {};
    message.primary !== undefined && (obj.primary = message.primary);
    message.detailed !== undefined && (obj.detailed = message.detailed);
    return obj;
  },

  create<I extends Exact<DeepPartial<PersonalFinanceCategory>, I>>(base?: I): PersonalFinanceCategory {
    return PersonalFinanceCategory.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PersonalFinanceCategory>, I>>(object: I): PersonalFinanceCategory {
    const message = createBasePersonalFinanceCategory();
    message.primary = object.primary ?? "";
    message.detailed = object.detailed ?? "";
    return message;
  },
};

function createBaseReOccuringTransaction(): ReOccuringTransaction {
  return {
    accountId: "",
    streamId: "",
    category: [],
    categoryId: "",
    description: "",
    merchantName: "",
    personalFinanceCategoryPrimary: "",
    personalFinanceCategoryDetailed: "",
    firstDate: "",
    lastDate: "",
    frequency: 0,
    transactionIds: [],
    averageAmount: "",
    averageAmountIsoCurrencyCode: "",
    lastAmount: "",
    lastAmountIsoCurrencyCode: "",
    isActive: false,
    status: 0,
    updatedTime: "",
    userId: 0,
    linkId: 0,
    id: 0,
    flow: 0,
  };
}

export const ReOccuringTransaction = {
  encode(message: ReOccuringTransaction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    if (message.streamId !== "") {
      writer.uint32(18).string(message.streamId);
    }
    for (const v of message.category) {
      writer.uint32(26).string(v!);
    }
    if (message.categoryId !== "") {
      writer.uint32(34).string(message.categoryId);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.merchantName !== "") {
      writer.uint32(50).string(message.merchantName);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(58).string(message.personalFinanceCategoryPrimary);
    }
    if (message.personalFinanceCategoryDetailed !== "") {
      writer.uint32(66).string(message.personalFinanceCategoryDetailed);
    }
    if (message.firstDate !== "") {
      writer.uint32(74).string(message.firstDate);
    }
    if (message.lastDate !== "") {
      writer.uint32(82).string(message.lastDate);
    }
    if (message.frequency !== 0) {
      writer.uint32(88).int32(message.frequency);
    }
    for (const v of message.transactionIds) {
      writer.uint32(98).string(v!);
    }
    if (message.averageAmount !== "") {
      writer.uint32(106).string(message.averageAmount);
    }
    if (message.averageAmountIsoCurrencyCode !== "") {
      writer.uint32(114).string(message.averageAmountIsoCurrencyCode);
    }
    if (message.lastAmount !== "") {
      writer.uint32(122).string(message.lastAmount);
    }
    if (message.lastAmountIsoCurrencyCode !== "") {
      writer.uint32(130).string(message.lastAmountIsoCurrencyCode);
    }
    if (message.isActive === true) {
      writer.uint32(136).bool(message.isActive);
    }
    if (message.status !== 0) {
      writer.uint32(144).int32(message.status);
    }
    if (message.updatedTime !== "") {
      writer.uint32(154).string(message.updatedTime);
    }
    if (message.userId !== 0) {
      writer.uint32(160).uint64(message.userId);
    }
    if (message.linkId !== 0) {
      writer.uint32(168).uint64(message.linkId);
    }
    if (message.id !== 0) {
      writer.uint32(176).uint64(message.id);
    }
    if (message.flow !== 0) {
      writer.uint32(184).int32(message.flow);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReOccuringTransaction {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReOccuringTransaction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accountId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.streamId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.category.push(reader.string());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.categoryId = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.description = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.merchantName = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.personalFinanceCategoryDetailed = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.firstDate = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.lastDate = reader.string();
          continue;
        case 11:
          if (tag !== 88) {
            break;
          }

          message.frequency = reader.int32() as any;
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.transactionIds.push(reader.string());
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.averageAmount = reader.string();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.averageAmountIsoCurrencyCode = reader.string();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.lastAmount = reader.string();
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.lastAmountIsoCurrencyCode = reader.string();
          continue;
        case 17:
          if (tag !== 136) {
            break;
          }

          message.isActive = reader.bool();
          continue;
        case 18:
          if (tag !== 144) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 19:
          if (tag !== 154) {
            break;
          }

          message.updatedTime = reader.string();
          continue;
        case 20:
          if (tag !== 160) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 21:
          if (tag !== 168) {
            break;
          }

          message.linkId = longToNumber(reader.uint64() as Long);
          continue;
        case 22:
          if (tag !== 176) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 23:
          if (tag !== 184) {
            break;
          }

          message.flow = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReOccuringTransaction {
    return {
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      streamId: isSet(object.streamId) ? String(object.streamId) : "",
      category: Array.isArray(object?.category) ? object.category.map((e: any) => String(e)) : [],
      categoryId: isSet(object.categoryId) ? String(object.categoryId) : "",
      description: isSet(object.description) ? String(object.description) : "",
      merchantName: isSet(object.merchantName) ? String(object.merchantName) : "",
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      personalFinanceCategoryDetailed: isSet(object.personalFinanceCategoryDetailed)
        ? String(object.personalFinanceCategoryDetailed)
        : "",
      firstDate: isSet(object.firstDate) ? String(object.firstDate) : "",
      lastDate: isSet(object.lastDate) ? String(object.lastDate) : "",
      frequency: isSet(object.frequency) ? reOccuringTransactionsFrequencyFromJSON(object.frequency) : 0,
      transactionIds: Array.isArray(object?.transactionIds) ? object.transactionIds.map((e: any) => String(e)) : [],
      averageAmount: isSet(object.averageAmount) ? String(object.averageAmount) : "",
      averageAmountIsoCurrencyCode: isSet(object.averageAmountIsoCurrencyCode)
        ? String(object.averageAmountIsoCurrencyCode)
        : "",
      lastAmount: isSet(object.lastAmount) ? String(object.lastAmount) : "",
      lastAmountIsoCurrencyCode: isSet(object.lastAmountIsoCurrencyCode)
        ? String(object.lastAmountIsoCurrencyCode)
        : "",
      isActive: isSet(object.isActive) ? Boolean(object.isActive) : false,
      status: isSet(object.status) ? reOccuringTransactionsStatusFromJSON(object.status) : 0,
      updatedTime: isSet(object.updatedTime) ? String(object.updatedTime) : "",
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      linkId: isSet(object.linkId) ? Number(object.linkId) : 0,
      id: isSet(object.id) ? Number(object.id) : 0,
      flow: isSet(object.flow) ? reCurringFlowFromJSON(object.flow) : 0,
    };
  },

  toJSON(message: ReOccuringTransaction): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.streamId !== undefined && (obj.streamId = message.streamId);
    if (message.category) {
      obj.category = message.category.map((e) => e);
    } else {
      obj.category = [];
    }
    message.categoryId !== undefined && (obj.categoryId = message.categoryId);
    message.description !== undefined && (obj.description = message.description);
    message.merchantName !== undefined && (obj.merchantName = message.merchantName);
    message.personalFinanceCategoryPrimary !== undefined &&
      (obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary);
    message.personalFinanceCategoryDetailed !== undefined &&
      (obj.personalFinanceCategoryDetailed = message.personalFinanceCategoryDetailed);
    message.firstDate !== undefined && (obj.firstDate = message.firstDate);
    message.lastDate !== undefined && (obj.lastDate = message.lastDate);
    message.frequency !== undefined && (obj.frequency = reOccuringTransactionsFrequencyToJSON(message.frequency));
    if (message.transactionIds) {
      obj.transactionIds = message.transactionIds.map((e) => e);
    } else {
      obj.transactionIds = [];
    }
    message.averageAmount !== undefined && (obj.averageAmount = message.averageAmount);
    message.averageAmountIsoCurrencyCode !== undefined &&
      (obj.averageAmountIsoCurrencyCode = message.averageAmountIsoCurrencyCode);
    message.lastAmount !== undefined && (obj.lastAmount = message.lastAmount);
    message.lastAmountIsoCurrencyCode !== undefined &&
      (obj.lastAmountIsoCurrencyCode = message.lastAmountIsoCurrencyCode);
    message.isActive !== undefined && (obj.isActive = message.isActive);
    message.status !== undefined && (obj.status = reOccuringTransactionsStatusToJSON(message.status));
    message.updatedTime !== undefined && (obj.updatedTime = message.updatedTime);
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.flow !== undefined && (obj.flow = reCurringFlowToJSON(message.flow));
    return obj;
  },

  create<I extends Exact<DeepPartial<ReOccuringTransaction>, I>>(base?: I): ReOccuringTransaction {
    return ReOccuringTransaction.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReOccuringTransaction>, I>>(object: I): ReOccuringTransaction {
    const message = createBaseReOccuringTransaction();
    message.accountId = object.accountId ?? "";
    message.streamId = object.streamId ?? "";
    message.category = object.category?.map((e) => e) || [];
    message.categoryId = object.categoryId ?? "";
    message.description = object.description ?? "";
    message.merchantName = object.merchantName ?? "";
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.personalFinanceCategoryDetailed = object.personalFinanceCategoryDetailed ?? "";
    message.firstDate = object.firstDate ?? "";
    message.lastDate = object.lastDate ?? "";
    message.frequency = object.frequency ?? 0;
    message.transactionIds = object.transactionIds?.map((e) => e) || [];
    message.averageAmount = object.averageAmount ?? "";
    message.averageAmountIsoCurrencyCode = object.averageAmountIsoCurrencyCode ?? "";
    message.lastAmount = object.lastAmount ?? "";
    message.lastAmountIsoCurrencyCode = object.lastAmountIsoCurrencyCode ?? "";
    message.isActive = object.isActive ?? false;
    message.status = object.status ?? 0;
    message.updatedTime = object.updatedTime ?? "";
    message.userId = object.userId ?? 0;
    message.linkId = object.linkId ?? 0;
    message.id = object.id ?? 0;
    message.flow = object.flow ?? 0;
    return message;
  },
};

function createBaseTransaction(): Transaction {
  return {
    accountId: "",
    amount: 0,
    isoCurrencyCode: "",
    unofficialCurrencyCode: "",
    category: "",
    categoryId: "",
    checkNumber: "",
    date: "",
    datetime: "",
    authorizedDate: "",
    authorizedDatetime: "",
    location: undefined,
    name: "",
    merchantName: "",
    paymentMeta: undefined,
    paymentChannel: "",
    pending: false,
    pendingTransactionId: "",
    accountOwner: "",
    transactionId: "",
    transactionCode: "",
    id: 0,
    userId: 0,
    linkId: 0,
  };
}

export const Transaction = {
  encode(message: Transaction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    if (message.amount !== 0) {
      writer.uint32(17).double(message.amount);
    }
    if (message.isoCurrencyCode !== "") {
      writer.uint32(26).string(message.isoCurrencyCode);
    }
    if (message.unofficialCurrencyCode !== "") {
      writer.uint32(34).string(message.unofficialCurrencyCode);
    }
    if (message.category !== "") {
      writer.uint32(42).string(message.category);
    }
    if (message.categoryId !== "") {
      writer.uint32(50).string(message.categoryId);
    }
    if (message.checkNumber !== "") {
      writer.uint32(58).string(message.checkNumber);
    }
    if (message.date !== "") {
      writer.uint32(66).string(message.date);
    }
    if (message.datetime !== "") {
      writer.uint32(74).string(message.datetime);
    }
    if (message.authorizedDate !== "") {
      writer.uint32(82).string(message.authorizedDate);
    }
    if (message.authorizedDatetime !== "") {
      writer.uint32(90).string(message.authorizedDatetime);
    }
    if (message.location !== undefined) {
      Transaction_Location.encode(message.location, writer.uint32(98).fork()).ldelim();
    }
    if (message.name !== "") {
      writer.uint32(106).string(message.name);
    }
    if (message.merchantName !== "") {
      writer.uint32(114).string(message.merchantName);
    }
    if (message.paymentMeta !== undefined) {
      Transaction_PaymentMeta.encode(message.paymentMeta, writer.uint32(122).fork()).ldelim();
    }
    if (message.paymentChannel !== "") {
      writer.uint32(130).string(message.paymentChannel);
    }
    if (message.pending === true) {
      writer.uint32(136).bool(message.pending);
    }
    if (message.pendingTransactionId !== "") {
      writer.uint32(146).string(message.pendingTransactionId);
    }
    if (message.accountOwner !== "") {
      writer.uint32(154).string(message.accountOwner);
    }
    if (message.transactionId !== "") {
      writer.uint32(162).string(message.transactionId);
    }
    if (message.transactionCode !== "") {
      writer.uint32(170).string(message.transactionCode);
    }
    if (message.id !== 0) {
      writer.uint32(176).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(184).uint64(message.userId);
    }
    if (message.linkId !== 0) {
      writer.uint32(192).uint64(message.linkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Transaction {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransaction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.accountId = reader.string();
          continue;
        case 2:
          if (tag !== 17) {
            break;
          }

          message.amount = reader.double();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.isoCurrencyCode = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.unofficialCurrencyCode = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.category = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.categoryId = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.checkNumber = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.date = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.datetime = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.authorizedDate = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.authorizedDatetime = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.location = Transaction_Location.decode(reader, reader.uint32());
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.name = reader.string();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.merchantName = reader.string();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.paymentMeta = Transaction_PaymentMeta.decode(reader, reader.uint32());
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.paymentChannel = reader.string();
          continue;
        case 17:
          if (tag !== 136) {
            break;
          }

          message.pending = reader.bool();
          continue;
        case 18:
          if (tag !== 146) {
            break;
          }

          message.pendingTransactionId = reader.string();
          continue;
        case 19:
          if (tag !== 154) {
            break;
          }

          message.accountOwner = reader.string();
          continue;
        case 20:
          if (tag !== 162) {
            break;
          }

          message.transactionId = reader.string();
          continue;
        case 21:
          if (tag !== 170) {
            break;
          }

          message.transactionCode = reader.string();
          continue;
        case 22:
          if (tag !== 176) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 23:
          if (tag !== 184) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 24:
          if (tag !== 192) {
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

  fromJSON(object: any): Transaction {
    return {
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      isoCurrencyCode: isSet(object.isoCurrencyCode) ? String(object.isoCurrencyCode) : "",
      unofficialCurrencyCode: isSet(object.unofficialCurrencyCode) ? String(object.unofficialCurrencyCode) : "",
      category: isSet(object.category) ? String(object.category) : "",
      categoryId: isSet(object.categoryId) ? String(object.categoryId) : "",
      checkNumber: isSet(object.checkNumber) ? String(object.checkNumber) : "",
      date: isSet(object.date) ? String(object.date) : "",
      datetime: isSet(object.datetime) ? String(object.datetime) : "",
      authorizedDate: isSet(object.authorizedDate) ? String(object.authorizedDate) : "",
      authorizedDatetime: isSet(object.authorizedDatetime) ? String(object.authorizedDatetime) : "",
      location: isSet(object.location) ? Transaction_Location.fromJSON(object.location) : undefined,
      name: isSet(object.name) ? String(object.name) : "",
      merchantName: isSet(object.merchantName) ? String(object.merchantName) : "",
      paymentMeta: isSet(object.paymentMeta) ? Transaction_PaymentMeta.fromJSON(object.paymentMeta) : undefined,
      paymentChannel: isSet(object.paymentChannel) ? String(object.paymentChannel) : "",
      pending: isSet(object.pending) ? Boolean(object.pending) : false,
      pendingTransactionId: isSet(object.pendingTransactionId) ? String(object.pendingTransactionId) : "",
      accountOwner: isSet(object.accountOwner) ? String(object.accountOwner) : "",
      transactionId: isSet(object.transactionId) ? String(object.transactionId) : "",
      transactionCode: isSet(object.transactionCode) ? String(object.transactionCode) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      linkId: isSet(object.linkId) ? Number(object.linkId) : 0,
    };
  },

  toJSON(message: Transaction): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.amount !== undefined && (obj.amount = message.amount);
    message.isoCurrencyCode !== undefined && (obj.isoCurrencyCode = message.isoCurrencyCode);
    message.unofficialCurrencyCode !== undefined && (obj.unofficialCurrencyCode = message.unofficialCurrencyCode);
    message.category !== undefined && (obj.category = message.category);
    message.categoryId !== undefined && (obj.categoryId = message.categoryId);
    message.checkNumber !== undefined && (obj.checkNumber = message.checkNumber);
    message.date !== undefined && (obj.date = message.date);
    message.datetime !== undefined && (obj.datetime = message.datetime);
    message.authorizedDate !== undefined && (obj.authorizedDate = message.authorizedDate);
    message.authorizedDatetime !== undefined && (obj.authorizedDatetime = message.authorizedDatetime);
    message.location !== undefined &&
      (obj.location = message.location ? Transaction_Location.toJSON(message.location) : undefined);
    message.name !== undefined && (obj.name = message.name);
    message.merchantName !== undefined && (obj.merchantName = message.merchantName);
    message.paymentMeta !== undefined &&
      (obj.paymentMeta = message.paymentMeta ? Transaction_PaymentMeta.toJSON(message.paymentMeta) : undefined);
    message.paymentChannel !== undefined && (obj.paymentChannel = message.paymentChannel);
    message.pending !== undefined && (obj.pending = message.pending);
    message.pendingTransactionId !== undefined && (obj.pendingTransactionId = message.pendingTransactionId);
    message.accountOwner !== undefined && (obj.accountOwner = message.accountOwner);
    message.transactionId !== undefined && (obj.transactionId = message.transactionId);
    message.transactionCode !== undefined && (obj.transactionCode = message.transactionCode);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    return obj;
  },

  create<I extends Exact<DeepPartial<Transaction>, I>>(base?: I): Transaction {
    return Transaction.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Transaction>, I>>(object: I): Transaction {
    const message = createBaseTransaction();
    message.accountId = object.accountId ?? "";
    message.amount = object.amount ?? 0;
    message.isoCurrencyCode = object.isoCurrencyCode ?? "";
    message.unofficialCurrencyCode = object.unofficialCurrencyCode ?? "";
    message.category = object.category ?? "";
    message.categoryId = object.categoryId ?? "";
    message.checkNumber = object.checkNumber ?? "";
    message.date = object.date ?? "";
    message.datetime = object.datetime ?? "";
    message.authorizedDate = object.authorizedDate ?? "";
    message.authorizedDatetime = object.authorizedDatetime ?? "";
    message.location = (object.location !== undefined && object.location !== null)
      ? Transaction_Location.fromPartial(object.location)
      : undefined;
    message.name = object.name ?? "";
    message.merchantName = object.merchantName ?? "";
    message.paymentMeta = (object.paymentMeta !== undefined && object.paymentMeta !== null)
      ? Transaction_PaymentMeta.fromPartial(object.paymentMeta)
      : undefined;
    message.paymentChannel = object.paymentChannel ?? "";
    message.pending = object.pending ?? false;
    message.pendingTransactionId = object.pendingTransactionId ?? "";
    message.accountOwner = object.accountOwner ?? "";
    message.transactionId = object.transactionId ?? "";
    message.transactionCode = object.transactionCode ?? "";
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.linkId = object.linkId ?? 0;
    return message;
  },
};

function createBaseTransaction_Location(): Transaction_Location {
  return { address: "", city: "", region: "", postalCode: "", country: "", lat: 0, lon: 0, storeNumber: "" };
}

export const Transaction_Location = {
  encode(message: Transaction_Location, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.city !== "") {
      writer.uint32(18).string(message.city);
    }
    if (message.region !== "") {
      writer.uint32(26).string(message.region);
    }
    if (message.postalCode !== "") {
      writer.uint32(34).string(message.postalCode);
    }
    if (message.country !== "") {
      writer.uint32(42).string(message.country);
    }
    if (message.lat !== 0) {
      writer.uint32(49).double(message.lat);
    }
    if (message.lon !== 0) {
      writer.uint32(57).double(message.lon);
    }
    if (message.storeNumber !== "") {
      writer.uint32(66).string(message.storeNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Transaction_Location {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransaction_Location();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.address = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.city = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.region = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.postalCode = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.country = reader.string();
          continue;
        case 6:
          if (tag !== 49) {
            break;
          }

          message.lat = reader.double();
          continue;
        case 7:
          if (tag !== 57) {
            break;
          }

          message.lon = reader.double();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.storeNumber = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Transaction_Location {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      city: isSet(object.city) ? String(object.city) : "",
      region: isSet(object.region) ? String(object.region) : "",
      postalCode: isSet(object.postalCode) ? String(object.postalCode) : "",
      country: isSet(object.country) ? String(object.country) : "",
      lat: isSet(object.lat) ? Number(object.lat) : 0,
      lon: isSet(object.lon) ? Number(object.lon) : 0,
      storeNumber: isSet(object.storeNumber) ? String(object.storeNumber) : "",
    };
  },

  toJSON(message: Transaction_Location): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.city !== undefined && (obj.city = message.city);
    message.region !== undefined && (obj.region = message.region);
    message.postalCode !== undefined && (obj.postalCode = message.postalCode);
    message.country !== undefined && (obj.country = message.country);
    message.lat !== undefined && (obj.lat = message.lat);
    message.lon !== undefined && (obj.lon = message.lon);
    message.storeNumber !== undefined && (obj.storeNumber = message.storeNumber);
    return obj;
  },

  create<I extends Exact<DeepPartial<Transaction_Location>, I>>(base?: I): Transaction_Location {
    return Transaction_Location.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Transaction_Location>, I>>(object: I): Transaction_Location {
    const message = createBaseTransaction_Location();
    message.address = object.address ?? "";
    message.city = object.city ?? "";
    message.region = object.region ?? "";
    message.postalCode = object.postalCode ?? "";
    message.country = object.country ?? "";
    message.lat = object.lat ?? 0;
    message.lon = object.lon ?? 0;
    message.storeNumber = object.storeNumber ?? "";
    return message;
  },
};

function createBaseTransaction_PaymentMeta(): Transaction_PaymentMeta {
  return {
    byOrderOf: "",
    payee: "",
    payer: "",
    paymentMethod: "",
    paymentProcessor: "",
    ppdId: "",
    reason: "",
    referenceNumber: "",
  };
}

export const Transaction_PaymentMeta = {
  encode(message: Transaction_PaymentMeta, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.byOrderOf !== "") {
      writer.uint32(10).string(message.byOrderOf);
    }
    if (message.payee !== "") {
      writer.uint32(18).string(message.payee);
    }
    if (message.payer !== "") {
      writer.uint32(26).string(message.payer);
    }
    if (message.paymentMethod !== "") {
      writer.uint32(34).string(message.paymentMethod);
    }
    if (message.paymentProcessor !== "") {
      writer.uint32(42).string(message.paymentProcessor);
    }
    if (message.ppdId !== "") {
      writer.uint32(50).string(message.ppdId);
    }
    if (message.reason !== "") {
      writer.uint32(58).string(message.reason);
    }
    if (message.referenceNumber !== "") {
      writer.uint32(66).string(message.referenceNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Transaction_PaymentMeta {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransaction_PaymentMeta();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.byOrderOf = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.payee = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.payer = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.paymentMethod = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.paymentProcessor = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.ppdId = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.reason = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.referenceNumber = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Transaction_PaymentMeta {
    return {
      byOrderOf: isSet(object.byOrderOf) ? String(object.byOrderOf) : "",
      payee: isSet(object.payee) ? String(object.payee) : "",
      payer: isSet(object.payer) ? String(object.payer) : "",
      paymentMethod: isSet(object.paymentMethod) ? String(object.paymentMethod) : "",
      paymentProcessor: isSet(object.paymentProcessor) ? String(object.paymentProcessor) : "",
      ppdId: isSet(object.ppdId) ? String(object.ppdId) : "",
      reason: isSet(object.reason) ? String(object.reason) : "",
      referenceNumber: isSet(object.referenceNumber) ? String(object.referenceNumber) : "",
    };
  },

  toJSON(message: Transaction_PaymentMeta): unknown {
    const obj: any = {};
    message.byOrderOf !== undefined && (obj.byOrderOf = message.byOrderOf);
    message.payee !== undefined && (obj.payee = message.payee);
    message.payer !== undefined && (obj.payer = message.payer);
    message.paymentMethod !== undefined && (obj.paymentMethod = message.paymentMethod);
    message.paymentProcessor !== undefined && (obj.paymentProcessor = message.paymentProcessor);
    message.ppdId !== undefined && (obj.ppdId = message.ppdId);
    message.reason !== undefined && (obj.reason = message.reason);
    message.referenceNumber !== undefined && (obj.referenceNumber = message.referenceNumber);
    return obj;
  },

  create<I extends Exact<DeepPartial<Transaction_PaymentMeta>, I>>(base?: I): Transaction_PaymentMeta {
    return Transaction_PaymentMeta.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Transaction_PaymentMeta>, I>>(object: I): Transaction_PaymentMeta {
    const message = createBaseTransaction_PaymentMeta();
    message.byOrderOf = object.byOrderOf ?? "";
    message.payee = object.payee ?? "";
    message.payer = object.payer ?? "";
    message.paymentMethod = object.paymentMethod ?? "";
    message.paymentProcessor = object.paymentProcessor ?? "";
    message.ppdId = object.ppdId ?? "";
    message.reason = object.reason ?? "";
    message.referenceNumber = object.referenceNumber ?? "";
    return message;
  },
};

function createBaseTransactionAmountByCountryMetric(): TransactionAmountByCountryMetric {
  return { country: "", amount: 0 };
}

export const TransactionAmountByCountryMetric = {
  encode(message: TransactionAmountByCountryMetric, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.country !== "") {
      writer.uint32(10).string(message.country);
    }
    if (message.amount !== 0) {
      writer.uint32(17).double(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransactionAmountByCountryMetric {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransactionAmountByCountryMetric();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.country = reader.string();
          continue;
        case 2:
          if (tag !== 17) {
            break;
          }

          message.amount = reader.double();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TransactionAmountByCountryMetric {
    return {
      country: isSet(object.country) ? String(object.country) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
    };
  },

  toJSON(message: TransactionAmountByCountryMetric): unknown {
    const obj: any = {};
    message.country !== undefined && (obj.country = message.country);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  create<I extends Exact<DeepPartial<TransactionAmountByCountryMetric>, I>>(
    base?: I,
  ): TransactionAmountByCountryMetric {
    return TransactionAmountByCountryMetric.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TransactionAmountByCountryMetric>, I>>(
    object: I,
  ): TransactionAmountByCountryMetric {
    const message = createBaseTransactionAmountByCountryMetric();
    message.country = object.country ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseAverageTransactionAmountByCategoryMetric(): AverageTransactionAmountByCategoryMetric {
  return { category: "", amount: 0 };
}

export const AverageTransactionAmountByCategoryMetric = {
  encode(message: AverageTransactionAmountByCategoryMetric, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.category !== "") {
      writer.uint32(10).string(message.category);
    }
    if (message.amount !== 0) {
      writer.uint32(17).double(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AverageTransactionAmountByCategoryMetric {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAverageTransactionAmountByCategoryMetric();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.category = reader.string();
          continue;
        case 2:
          if (tag !== 17) {
            break;
          }

          message.amount = reader.double();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AverageTransactionAmountByCategoryMetric {
    return {
      category: isSet(object.category) ? String(object.category) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
    };
  },

  toJSON(message: AverageTransactionAmountByCategoryMetric): unknown {
    const obj: any = {};
    message.category !== undefined && (obj.category = message.category);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  create<I extends Exact<DeepPartial<AverageTransactionAmountByCategoryMetric>, I>>(
    base?: I,
  ): AverageTransactionAmountByCategoryMetric {
    return AverageTransactionAmountByCategoryMetric.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AverageTransactionAmountByCategoryMetric>, I>>(
    object: I,
  ): AverageTransactionAmountByCategoryMetric {
    const message = createBaseAverageTransactionAmountByCategoryMetric();
    message.category = object.category ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseMonthlyTransactionCountByCategoryMetric(): MonthlyTransactionCountByCategoryMetric {
  return { category: "", count: 0, month: "" };
}

export const MonthlyTransactionCountByCategoryMetric = {
  encode(message: MonthlyTransactionCountByCategoryMetric, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.category !== "") {
      writer.uint32(10).string(message.category);
    }
    if (message.count !== 0) {
      writer.uint32(16).uint32(message.count);
    }
    if (message.month !== "") {
      writer.uint32(26).string(message.month);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MonthlyTransactionCountByCategoryMetric {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMonthlyTransactionCountByCategoryMetric();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.category = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.count = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.month = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MonthlyTransactionCountByCategoryMetric {
    return {
      category: isSet(object.category) ? String(object.category) : "",
      count: isSet(object.count) ? Number(object.count) : 0,
      month: isSet(object.month) ? String(object.month) : "",
    };
  },

  toJSON(message: MonthlyTransactionCountByCategoryMetric): unknown {
    const obj: any = {};
    message.category !== undefined && (obj.category = message.category);
    message.count !== undefined && (obj.count = Math.round(message.count));
    message.month !== undefined && (obj.month = message.month);
    return obj;
  },

  create<I extends Exact<DeepPartial<MonthlyTransactionCountByCategoryMetric>, I>>(
    base?: I,
  ): MonthlyTransactionCountByCategoryMetric {
    return MonthlyTransactionCountByCategoryMetric.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<MonthlyTransactionCountByCategoryMetric>, I>>(
    object: I,
  ): MonthlyTransactionCountByCategoryMetric {
    const message = createBaseMonthlyTransactionCountByCategoryMetric();
    message.category = object.category ?? "";
    message.count = object.count ?? 0;
    message.month = object.month ?? "";
    return message;
  },
};

function createBaseTransactionCountByMerchantPaymentChannelMetric(): TransactionCountByMerchantPaymentChannelMetric {
  return { merchantName: "", paymentChannel: "", transactionCount: 0 };
}

export const TransactionCountByMerchantPaymentChannelMetric = {
  encode(
    message: TransactionCountByMerchantPaymentChannelMetric,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.merchantName !== "") {
      writer.uint32(10).string(message.merchantName);
    }
    if (message.paymentChannel !== "") {
      writer.uint32(18).string(message.paymentChannel);
    }
    if (message.transactionCount !== 0) {
      writer.uint32(24).uint32(message.transactionCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransactionCountByMerchantPaymentChannelMetric {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransactionCountByMerchantPaymentChannelMetric();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.merchantName = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.paymentChannel = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.transactionCount = reader.uint32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TransactionCountByMerchantPaymentChannelMetric {
    return {
      merchantName: isSet(object.merchantName) ? String(object.merchantName) : "",
      paymentChannel: isSet(object.paymentChannel) ? String(object.paymentChannel) : "",
      transactionCount: isSet(object.transactionCount) ? Number(object.transactionCount) : 0,
    };
  },

  toJSON(message: TransactionCountByMerchantPaymentChannelMetric): unknown {
    const obj: any = {};
    message.merchantName !== undefined && (obj.merchantName = message.merchantName);
    message.paymentChannel !== undefined && (obj.paymentChannel = message.paymentChannel);
    message.transactionCount !== undefined && (obj.transactionCount = Math.round(message.transactionCount));
    return obj;
  },

  create<I extends Exact<DeepPartial<TransactionCountByMerchantPaymentChannelMetric>, I>>(
    base?: I,
  ): TransactionCountByMerchantPaymentChannelMetric {
    return TransactionCountByMerchantPaymentChannelMetric.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TransactionCountByMerchantPaymentChannelMetric>, I>>(
    object: I,
  ): TransactionCountByMerchantPaymentChannelMetric {
    const message = createBaseTransactionCountByMerchantPaymentChannelMetric();
    message.merchantName = object.merchantName ?? "";
    message.paymentChannel = object.paymentChannel ?? "";
    message.transactionCount = object.transactionCount ?? 0;
    return message;
  },
};

function createBaseTransactionAmountDistributionByCategoryMetric(): TransactionAmountDistributionByCategoryMetric {
  return { category: "", mean: 0, median: 0, standardDeviation: 0 };
}

export const TransactionAmountDistributionByCategoryMetric = {
  encode(message: TransactionAmountDistributionByCategoryMetric, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.category !== "") {
      writer.uint32(10).string(message.category);
    }
    if (message.mean !== 0) {
      writer.uint32(17).double(message.mean);
    }
    if (message.median !== 0) {
      writer.uint32(25).double(message.median);
    }
    if (message.standardDeviation !== 0) {
      writer.uint32(33).double(message.standardDeviation);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransactionAmountDistributionByCategoryMetric {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransactionAmountDistributionByCategoryMetric();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.category = reader.string();
          continue;
        case 2:
          if (tag !== 17) {
            break;
          }

          message.mean = reader.double();
          continue;
        case 3:
          if (tag !== 25) {
            break;
          }

          message.median = reader.double();
          continue;
        case 4:
          if (tag !== 33) {
            break;
          }

          message.standardDeviation = reader.double();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TransactionAmountDistributionByCategoryMetric {
    return {
      category: isSet(object.category) ? String(object.category) : "",
      mean: isSet(object.mean) ? Number(object.mean) : 0,
      median: isSet(object.median) ? Number(object.median) : 0,
      standardDeviation: isSet(object.standardDeviation) ? Number(object.standardDeviation) : 0,
    };
  },

  toJSON(message: TransactionAmountDistributionByCategoryMetric): unknown {
    const obj: any = {};
    message.category !== undefined && (obj.category = message.category);
    message.mean !== undefined && (obj.mean = message.mean);
    message.median !== undefined && (obj.median = message.median);
    message.standardDeviation !== undefined && (obj.standardDeviation = message.standardDeviation);
    return obj;
  },

  create<I extends Exact<DeepPartial<TransactionAmountDistributionByCategoryMetric>, I>>(
    base?: I,
  ): TransactionAmountDistributionByCategoryMetric {
    return TransactionAmountDistributionByCategoryMetric.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TransactionAmountDistributionByCategoryMetric>, I>>(
    object: I,
  ): TransactionAmountDistributionByCategoryMetric {
    const message = createBaseTransactionAmountDistributionByCategoryMetric();
    message.category = object.category ?? "";
    message.mean = object.mean ?? 0;
    message.median = object.median ?? 0;
    message.standardDeviation = object.standardDeviation ?? 0;
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
