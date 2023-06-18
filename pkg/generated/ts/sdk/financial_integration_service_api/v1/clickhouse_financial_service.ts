/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Any } from "../../google/protobuf/any";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "financial_integration_service_api.v1";

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

export interface InvestmentTransaction {
  /** @gotag: ch:"account_id" */
  accountId: string;
  /** @gotag: ch:"amount" */
  ammount: string;
  /** @gotag: ch:"investment_transaction_id" */
  investmentTransactionId: string;
  /** @gotag: ch:"security_id" */
  securityId: string;
  /** @gotag: ch:"date" */
  date: string;
  /** @gotag: ch:"name" */
  name: string;
  /** @gotag: ch:"quantity" */
  quantity: number;
  /** @gotag: ch:"amount" */
  amount: number;
  /** @gotag: ch:"price" */
  price: number;
  /** @gotag: ch:"fees" */
  fees: number;
  /** @gotag: ch:"type" */
  type: string;
  /** @gotag: ch:"subtype" */
  subtype: string;
  /** @gotag: ch:"iso_currency_code" */
  isoCurrencyCode: string;
  /** @gotag: ch:"unofficial_currency_code" */
  unofficialCurrencyCode: string;
  /** @gotag: ch:"link_id" */
  linkId: number;
  /** @gotag: ch:"id" */
  id: number;
  /** @gotag: ch:"user_id" */
  userId: number;
  createdAt: string;
  sign: number;
  time: Date | undefined;
  additionalProperties: Any | undefined;
}

export interface ReOccuringTransaction {
  /** @gotag: ch:"account_id" */
  accountId: string;
  /** @gotag: ch:"stream_id" */
  streamId: string;
  /** @gotag: ch:"category_id" */
  categoryId: string;
  /** @gotag: ch:"description" */
  description: string;
  /** @gotag: ch:"merchant_name" */
  merchantName: string;
  /** @gotag: ch:"personal_finance_category_primary" */
  personalFinanceCategoryPrimary: string;
  /** @gotag: ch:"personal_finance_category_detailed" */
  personalFinanceCategoryDetailed: string;
  /** @gotag: ch:"first_date" */
  firstDate: string;
  /** @gotag: ch:"last_date" */
  lastDate: string;
  /** @gotag: ch:"frequency" */
  frequency: ReOccuringTransactionsFrequency;
  /** @gotag: ch:"transaction_ids,array" */
  transactionIds: string;
  /** @gotag: ch:"average_amount" */
  averageAmount: string;
  /** @gotag: ch:"average_amount_iso_currency_code" */
  averageAmountIsoCurrencyCode: string;
  /** @gotag: ch:"last_amount" */
  lastAmount: string;
  /** @gotag: ch:"last_amount_iso_currency_code" */
  lastAmountIsoCurrencyCode: string;
  /** @gotag: ch:"is_active" */
  isActive: boolean;
  /** @gotag: ch:"status" */
  status: ReOccuringTransactionsStatus;
  /** @gotag: ch:"updated_time" */
  updatedTime: string;
  /** @gotag: ch:"user_id" */
  userId: number;
  /** @gotag: ch:"link_id" */
  linkId: number;
  /** @gotag: ch:"id" */
  id: number;
  /** @gotag: ch:"flow" */
  flow: ReCurringFlow;
  sign: number;
  time: Date | undefined;
  additionalProperties: Any | undefined;
}

export interface Transaction {
  /** @gotag: ch:"account_id" */
  accountId: string;
  /** @gotag: ch:"amount" */
  amount: number;
  /** @gotag: ch:"iso_currency_code" */
  isoCurrencyCode: string;
  /** @gotag: ch:"unofficial_currency_code" */
  unofficialCurrencyCode: string;
  /** @gotag: ch:"category_id" */
  categoryId: string;
  /** @gotag: ch:"check_number" */
  checkNumber: string;
  /** @gotag: ch:"date" */
  date: string;
  /** @gotag: ch:"datetime" */
  datetime: string;
  /** @gotag: ch:"authorized_date" */
  authorizedDate: string;
  /** @gotag: ch:"authorized_datetime" */
  authorizedDatetime: string;
  /** @gotag: ch:"name" */
  name: string;
  /** @gotag: ch:"merchant_name" */
  merchantName: string;
  /** @gotag: ch:"payment_channel" */
  paymentChannel: string;
  /** @gotag: ch:"pending" */
  pending: boolean;
  /** @gotag: ch:"pending_transaction_id" */
  pendingTransactionId: string;
  /** @gotag: ch:"account_owner" */
  accountOwner: string;
  /** @gotag: ch:"transaction_id" */
  transactionId: string;
  /** @gotag: ch:"transaction_code" */
  transactionCode: string;
  id: number;
  /** @gotag: ch:"user_id" */
  userId: number;
  /** @gotag: ch:"link_id" */
  linkId: number;
  sign: number;
  /** @gotag: ch:"personal_finance_category_primary" */
  personalFinanceCategoryPrimary: string;
  /** @gotag: ch:"personal_finance_category_detailed" */
  personalFinanceCategoryDetailed: string;
  locationAddress: string;
  locationCity: string;
  locationRegion: string;
  locationPostalCode: string;
  locationCountry: string;
  locationLat: number;
  locationLon: number;
  locationStoreNumber: string;
  paymentMetaByOrderOf: string;
  paymentMetaPayee: string;
  paymentMetaPayer: string;
  paymentMetaPaymentMethod: string;
  paymentMetaPaymentProcessor: string;
  paymentMetaPpdId: string;
  paymentMetaReason: string;
  paymentMetaReferenceNumber: string;
  time: Date | undefined;
  additionalProperties: Any | undefined;
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

function createBaseInvestmentTransaction(): InvestmentTransaction {
  return {
    accountId: "",
    ammount: "",
    investmentTransactionId: "",
    securityId: "",
    date: "",
    name: "",
    quantity: 0,
    amount: 0,
    price: 0,
    fees: 0,
    type: "",
    subtype: "",
    isoCurrencyCode: "",
    unofficialCurrencyCode: "",
    linkId: 0,
    id: 0,
    userId: 0,
    createdAt: "",
    sign: 0,
    time: undefined,
    additionalProperties: undefined,
  };
}

export const InvestmentTransaction = {
  encode(message: InvestmentTransaction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    if (message.ammount !== "") {
      writer.uint32(18).string(message.ammount);
    }
    if (message.investmentTransactionId !== "") {
      writer.uint32(26).string(message.investmentTransactionId);
    }
    if (message.securityId !== "") {
      writer.uint32(34).string(message.securityId);
    }
    if (message.date !== "") {
      writer.uint32(42).string(message.date);
    }
    if (message.name !== "") {
      writer.uint32(50).string(message.name);
    }
    if (message.quantity !== 0) {
      writer.uint32(57).double(message.quantity);
    }
    if (message.amount !== 0) {
      writer.uint32(65).double(message.amount);
    }
    if (message.price !== 0) {
      writer.uint32(73).double(message.price);
    }
    if (message.fees !== 0) {
      writer.uint32(81).double(message.fees);
    }
    if (message.type !== "") {
      writer.uint32(90).string(message.type);
    }
    if (message.subtype !== "") {
      writer.uint32(98).string(message.subtype);
    }
    if (message.isoCurrencyCode !== "") {
      writer.uint32(106).string(message.isoCurrencyCode);
    }
    if (message.unofficialCurrencyCode !== "") {
      writer.uint32(114).string(message.unofficialCurrencyCode);
    }
    if (message.linkId !== 0) {
      writer.uint32(120).uint64(message.linkId);
    }
    if (message.id !== 0) {
      writer.uint32(128).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(136).uint64(message.userId);
    }
    if (message.createdAt !== "") {
      writer.uint32(146).string(message.createdAt);
    }
    if (message.sign !== 0) {
      writer.uint32(152).int32(message.sign);
    }
    if (message.time !== undefined) {
      Timestamp.encode(toTimestamp(message.time), writer.uint32(162).fork()).ldelim();
    }
    if (message.additionalProperties !== undefined) {
      Any.encode(message.additionalProperties, writer.uint32(170).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InvestmentTransaction {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInvestmentTransaction();
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

          message.ammount = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.investmentTransactionId = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.securityId = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.date = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.name = reader.string();
          continue;
        case 7:
          if (tag !== 57) {
            break;
          }

          message.quantity = reader.double();
          continue;
        case 8:
          if (tag !== 65) {
            break;
          }

          message.amount = reader.double();
          continue;
        case 9:
          if (tag !== 73) {
            break;
          }

          message.price = reader.double();
          continue;
        case 10:
          if (tag !== 81) {
            break;
          }

          message.fees = reader.double();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.type = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.subtype = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.isoCurrencyCode = reader.string();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.unofficialCurrencyCode = reader.string();
          continue;
        case 15:
          if (tag !== 120) {
            break;
          }

          message.linkId = longToNumber(reader.uint64() as Long);
          continue;
        case 16:
          if (tag !== 128) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 17:
          if (tag !== 136) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 18:
          if (tag !== 146) {
            break;
          }

          message.createdAt = reader.string();
          continue;
        case 19:
          if (tag !== 152) {
            break;
          }

          message.sign = reader.int32();
          continue;
        case 20:
          if (tag !== 162) {
            break;
          }

          message.time = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 21:
          if (tag !== 170) {
            break;
          }

          message.additionalProperties = Any.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InvestmentTransaction {
    return {
      accountId: isSet(object.accountId) ? String(object.accountId) : "",
      ammount: isSet(object.ammount) ? String(object.ammount) : "",
      investmentTransactionId: isSet(object.investmentTransactionId) ? String(object.investmentTransactionId) : "",
      securityId: isSet(object.securityId) ? String(object.securityId) : "",
      date: isSet(object.date) ? String(object.date) : "",
      name: isSet(object.name) ? String(object.name) : "",
      quantity: isSet(object.quantity) ? Number(object.quantity) : 0,
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      price: isSet(object.price) ? Number(object.price) : 0,
      fees: isSet(object.fees) ? Number(object.fees) : 0,
      type: isSet(object.type) ? String(object.type) : "",
      subtype: isSet(object.subtype) ? String(object.subtype) : "",
      isoCurrencyCode: isSet(object.isoCurrencyCode) ? String(object.isoCurrencyCode) : "",
      unofficialCurrencyCode: isSet(object.unofficialCurrencyCode) ? String(object.unofficialCurrencyCode) : "",
      linkId: isSet(object.linkId) ? Number(object.linkId) : 0,
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      createdAt: isSet(object.createdAt) ? String(object.createdAt) : "",
      sign: isSet(object.sign) ? Number(object.sign) : 0,
      time: isSet(object.time) ? fromJsonTimestamp(object.time) : undefined,
      additionalProperties: isSet(object.additionalProperties) ? Any.fromJSON(object.additionalProperties) : undefined,
    };
  },

  toJSON(message: InvestmentTransaction): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.ammount !== undefined && (obj.ammount = message.ammount);
    message.investmentTransactionId !== undefined && (obj.investmentTransactionId = message.investmentTransactionId);
    message.securityId !== undefined && (obj.securityId = message.securityId);
    message.date !== undefined && (obj.date = message.date);
    message.name !== undefined && (obj.name = message.name);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    message.fees !== undefined && (obj.fees = message.fees);
    message.type !== undefined && (obj.type = message.type);
    message.subtype !== undefined && (obj.subtype = message.subtype);
    message.isoCurrencyCode !== undefined && (obj.isoCurrencyCode = message.isoCurrencyCode);
    message.unofficialCurrencyCode !== undefined && (obj.unofficialCurrencyCode = message.unofficialCurrencyCode);
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    message.sign !== undefined && (obj.sign = Math.round(message.sign));
    message.time !== undefined && (obj.time = message.time.toISOString());
    message.additionalProperties !== undefined &&
      (obj.additionalProperties = message.additionalProperties ? Any.toJSON(message.additionalProperties) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<InvestmentTransaction>, I>>(base?: I): InvestmentTransaction {
    return InvestmentTransaction.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<InvestmentTransaction>, I>>(object: I): InvestmentTransaction {
    const message = createBaseInvestmentTransaction();
    message.accountId = object.accountId ?? "";
    message.ammount = object.ammount ?? "";
    message.investmentTransactionId = object.investmentTransactionId ?? "";
    message.securityId = object.securityId ?? "";
    message.date = object.date ?? "";
    message.name = object.name ?? "";
    message.quantity = object.quantity ?? 0;
    message.amount = object.amount ?? 0;
    message.price = object.price ?? 0;
    message.fees = object.fees ?? 0;
    message.type = object.type ?? "";
    message.subtype = object.subtype ?? "";
    message.isoCurrencyCode = object.isoCurrencyCode ?? "";
    message.unofficialCurrencyCode = object.unofficialCurrencyCode ?? "";
    message.linkId = object.linkId ?? 0;
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.createdAt = object.createdAt ?? "";
    message.sign = object.sign ?? 0;
    message.time = object.time ?? undefined;
    message.additionalProperties = (object.additionalProperties !== undefined && object.additionalProperties !== null)
      ? Any.fromPartial(object.additionalProperties)
      : undefined;
    return message;
  },
};

function createBaseReOccuringTransaction(): ReOccuringTransaction {
  return {
    accountId: "",
    streamId: "",
    categoryId: "",
    description: "",
    merchantName: "",
    personalFinanceCategoryPrimary: "",
    personalFinanceCategoryDetailed: "",
    firstDate: "",
    lastDate: "",
    frequency: 0,
    transactionIds: "",
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
    sign: 0,
    time: undefined,
    additionalProperties: undefined,
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
    if (message.transactionIds !== "") {
      writer.uint32(98).string(message.transactionIds);
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
    if (message.sign !== 0) {
      writer.uint32(192).int32(message.sign);
    }
    if (message.time !== undefined) {
      Timestamp.encode(toTimestamp(message.time), writer.uint32(202).fork()).ldelim();
    }
    if (message.additionalProperties !== undefined) {
      Any.encode(message.additionalProperties, writer.uint32(210).fork()).ldelim();
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

          message.transactionIds = reader.string();
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
        case 24:
          if (tag !== 192) {
            break;
          }

          message.sign = reader.int32();
          continue;
        case 25:
          if (tag !== 202) {
            break;
          }

          message.time = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 26:
          if (tag !== 210) {
            break;
          }

          message.additionalProperties = Any.decode(reader, reader.uint32());
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
      transactionIds: isSet(object.transactionIds) ? String(object.transactionIds) : "",
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
      sign: isSet(object.sign) ? Number(object.sign) : 0,
      time: isSet(object.time) ? fromJsonTimestamp(object.time) : undefined,
      additionalProperties: isSet(object.additionalProperties) ? Any.fromJSON(object.additionalProperties) : undefined,
    };
  },

  toJSON(message: ReOccuringTransaction): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.streamId !== undefined && (obj.streamId = message.streamId);
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
    message.transactionIds !== undefined && (obj.transactionIds = message.transactionIds);
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
    message.sign !== undefined && (obj.sign = Math.round(message.sign));
    message.time !== undefined && (obj.time = message.time.toISOString());
    message.additionalProperties !== undefined &&
      (obj.additionalProperties = message.additionalProperties ? Any.toJSON(message.additionalProperties) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReOccuringTransaction>, I>>(base?: I): ReOccuringTransaction {
    return ReOccuringTransaction.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReOccuringTransaction>, I>>(object: I): ReOccuringTransaction {
    const message = createBaseReOccuringTransaction();
    message.accountId = object.accountId ?? "";
    message.streamId = object.streamId ?? "";
    message.categoryId = object.categoryId ?? "";
    message.description = object.description ?? "";
    message.merchantName = object.merchantName ?? "";
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.personalFinanceCategoryDetailed = object.personalFinanceCategoryDetailed ?? "";
    message.firstDate = object.firstDate ?? "";
    message.lastDate = object.lastDate ?? "";
    message.frequency = object.frequency ?? 0;
    message.transactionIds = object.transactionIds ?? "";
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
    message.sign = object.sign ?? 0;
    message.time = object.time ?? undefined;
    message.additionalProperties = (object.additionalProperties !== undefined && object.additionalProperties !== null)
      ? Any.fromPartial(object.additionalProperties)
      : undefined;
    return message;
  },
};

function createBaseTransaction(): Transaction {
  return {
    accountId: "",
    amount: 0,
    isoCurrencyCode: "",
    unofficialCurrencyCode: "",
    categoryId: "",
    checkNumber: "",
    date: "",
    datetime: "",
    authorizedDate: "",
    authorizedDatetime: "",
    name: "",
    merchantName: "",
    paymentChannel: "",
    pending: false,
    pendingTransactionId: "",
    accountOwner: "",
    transactionId: "",
    transactionCode: "",
    id: 0,
    userId: 0,
    linkId: 0,
    sign: 0,
    personalFinanceCategoryPrimary: "",
    personalFinanceCategoryDetailed: "",
    locationAddress: "",
    locationCity: "",
    locationRegion: "",
    locationPostalCode: "",
    locationCountry: "",
    locationLat: 0,
    locationLon: 0,
    locationStoreNumber: "",
    paymentMetaByOrderOf: "",
    paymentMetaPayee: "",
    paymentMetaPayer: "",
    paymentMetaPaymentMethod: "",
    paymentMetaPaymentProcessor: "",
    paymentMetaPpdId: "",
    paymentMetaReason: "",
    paymentMetaReferenceNumber: "",
    time: undefined,
    additionalProperties: undefined,
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
    if (message.name !== "") {
      writer.uint32(106).string(message.name);
    }
    if (message.merchantName !== "") {
      writer.uint32(114).string(message.merchantName);
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
    if (message.sign !== 0) {
      writer.uint32(200).int32(message.sign);
    }
    if (message.personalFinanceCategoryPrimary !== "") {
      writer.uint32(210).string(message.personalFinanceCategoryPrimary);
    }
    if (message.personalFinanceCategoryDetailed !== "") {
      writer.uint32(218).string(message.personalFinanceCategoryDetailed);
    }
    if (message.locationAddress !== "") {
      writer.uint32(226).string(message.locationAddress);
    }
    if (message.locationCity !== "") {
      writer.uint32(234).string(message.locationCity);
    }
    if (message.locationRegion !== "") {
      writer.uint32(242).string(message.locationRegion);
    }
    if (message.locationPostalCode !== "") {
      writer.uint32(250).string(message.locationPostalCode);
    }
    if (message.locationCountry !== "") {
      writer.uint32(258).string(message.locationCountry);
    }
    if (message.locationLat !== 0) {
      writer.uint32(265).double(message.locationLat);
    }
    if (message.locationLon !== 0) {
      writer.uint32(273).double(message.locationLon);
    }
    if (message.locationStoreNumber !== "") {
      writer.uint32(282).string(message.locationStoreNumber);
    }
    if (message.paymentMetaByOrderOf !== "") {
      writer.uint32(290).string(message.paymentMetaByOrderOf);
    }
    if (message.paymentMetaPayee !== "") {
      writer.uint32(298).string(message.paymentMetaPayee);
    }
    if (message.paymentMetaPayer !== "") {
      writer.uint32(306).string(message.paymentMetaPayer);
    }
    if (message.paymentMetaPaymentMethod !== "") {
      writer.uint32(314).string(message.paymentMetaPaymentMethod);
    }
    if (message.paymentMetaPaymentProcessor !== "") {
      writer.uint32(322).string(message.paymentMetaPaymentProcessor);
    }
    if (message.paymentMetaPpdId !== "") {
      writer.uint32(330).string(message.paymentMetaPpdId);
    }
    if (message.paymentMetaReason !== "") {
      writer.uint32(338).string(message.paymentMetaReason);
    }
    if (message.paymentMetaReferenceNumber !== "") {
      writer.uint32(346).string(message.paymentMetaReferenceNumber);
    }
    if (message.time !== undefined) {
      Timestamp.encode(toTimestamp(message.time), writer.uint32(354).fork()).ldelim();
    }
    if (message.additionalProperties !== undefined) {
      Any.encode(message.additionalProperties, writer.uint32(362).fork()).ldelim();
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
        case 25:
          if (tag !== 200) {
            break;
          }

          message.sign = reader.int32();
          continue;
        case 26:
          if (tag !== 210) {
            break;
          }

          message.personalFinanceCategoryPrimary = reader.string();
          continue;
        case 27:
          if (tag !== 218) {
            break;
          }

          message.personalFinanceCategoryDetailed = reader.string();
          continue;
        case 28:
          if (tag !== 226) {
            break;
          }

          message.locationAddress = reader.string();
          continue;
        case 29:
          if (tag !== 234) {
            break;
          }

          message.locationCity = reader.string();
          continue;
        case 30:
          if (tag !== 242) {
            break;
          }

          message.locationRegion = reader.string();
          continue;
        case 31:
          if (tag !== 250) {
            break;
          }

          message.locationPostalCode = reader.string();
          continue;
        case 32:
          if (tag !== 258) {
            break;
          }

          message.locationCountry = reader.string();
          continue;
        case 33:
          if (tag !== 265) {
            break;
          }

          message.locationLat = reader.double();
          continue;
        case 34:
          if (tag !== 273) {
            break;
          }

          message.locationLon = reader.double();
          continue;
        case 35:
          if (tag !== 282) {
            break;
          }

          message.locationStoreNumber = reader.string();
          continue;
        case 36:
          if (tag !== 290) {
            break;
          }

          message.paymentMetaByOrderOf = reader.string();
          continue;
        case 37:
          if (tag !== 298) {
            break;
          }

          message.paymentMetaPayee = reader.string();
          continue;
        case 38:
          if (tag !== 306) {
            break;
          }

          message.paymentMetaPayer = reader.string();
          continue;
        case 39:
          if (tag !== 314) {
            break;
          }

          message.paymentMetaPaymentMethod = reader.string();
          continue;
        case 40:
          if (tag !== 322) {
            break;
          }

          message.paymentMetaPaymentProcessor = reader.string();
          continue;
        case 41:
          if (tag !== 330) {
            break;
          }

          message.paymentMetaPpdId = reader.string();
          continue;
        case 42:
          if (tag !== 338) {
            break;
          }

          message.paymentMetaReason = reader.string();
          continue;
        case 43:
          if (tag !== 346) {
            break;
          }

          message.paymentMetaReferenceNumber = reader.string();
          continue;
        case 44:
          if (tag !== 354) {
            break;
          }

          message.time = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 45:
          if (tag !== 362) {
            break;
          }

          message.additionalProperties = Any.decode(reader, reader.uint32());
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
      categoryId: isSet(object.categoryId) ? String(object.categoryId) : "",
      checkNumber: isSet(object.checkNumber) ? String(object.checkNumber) : "",
      date: isSet(object.date) ? String(object.date) : "",
      datetime: isSet(object.datetime) ? String(object.datetime) : "",
      authorizedDate: isSet(object.authorizedDate) ? String(object.authorizedDate) : "",
      authorizedDatetime: isSet(object.authorizedDatetime) ? String(object.authorizedDatetime) : "",
      name: isSet(object.name) ? String(object.name) : "",
      merchantName: isSet(object.merchantName) ? String(object.merchantName) : "",
      paymentChannel: isSet(object.paymentChannel) ? String(object.paymentChannel) : "",
      pending: isSet(object.pending) ? Boolean(object.pending) : false,
      pendingTransactionId: isSet(object.pendingTransactionId) ? String(object.pendingTransactionId) : "",
      accountOwner: isSet(object.accountOwner) ? String(object.accountOwner) : "",
      transactionId: isSet(object.transactionId) ? String(object.transactionId) : "",
      transactionCode: isSet(object.transactionCode) ? String(object.transactionCode) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      linkId: isSet(object.linkId) ? Number(object.linkId) : 0,
      sign: isSet(object.sign) ? Number(object.sign) : 0,
      personalFinanceCategoryPrimary: isSet(object.personalFinanceCategoryPrimary)
        ? String(object.personalFinanceCategoryPrimary)
        : "",
      personalFinanceCategoryDetailed: isSet(object.personalFinanceCategoryDetailed)
        ? String(object.personalFinanceCategoryDetailed)
        : "",
      locationAddress: isSet(object.locationAddress) ? String(object.locationAddress) : "",
      locationCity: isSet(object.locationCity) ? String(object.locationCity) : "",
      locationRegion: isSet(object.locationRegion) ? String(object.locationRegion) : "",
      locationPostalCode: isSet(object.locationPostalCode) ? String(object.locationPostalCode) : "",
      locationCountry: isSet(object.locationCountry) ? String(object.locationCountry) : "",
      locationLat: isSet(object.locationLat) ? Number(object.locationLat) : 0,
      locationLon: isSet(object.locationLon) ? Number(object.locationLon) : 0,
      locationStoreNumber: isSet(object.locationStoreNumber) ? String(object.locationStoreNumber) : "",
      paymentMetaByOrderOf: isSet(object.paymentMetaByOrderOf) ? String(object.paymentMetaByOrderOf) : "",
      paymentMetaPayee: isSet(object.paymentMetaPayee) ? String(object.paymentMetaPayee) : "",
      paymentMetaPayer: isSet(object.paymentMetaPayer) ? String(object.paymentMetaPayer) : "",
      paymentMetaPaymentMethod: isSet(object.paymentMetaPaymentMethod) ? String(object.paymentMetaPaymentMethod) : "",
      paymentMetaPaymentProcessor: isSet(object.paymentMetaPaymentProcessor)
        ? String(object.paymentMetaPaymentProcessor)
        : "",
      paymentMetaPpdId: isSet(object.paymentMetaPpdId) ? String(object.paymentMetaPpdId) : "",
      paymentMetaReason: isSet(object.paymentMetaReason) ? String(object.paymentMetaReason) : "",
      paymentMetaReferenceNumber: isSet(object.paymentMetaReferenceNumber)
        ? String(object.paymentMetaReferenceNumber)
        : "",
      time: isSet(object.time) ? fromJsonTimestamp(object.time) : undefined,
      additionalProperties: isSet(object.additionalProperties) ? Any.fromJSON(object.additionalProperties) : undefined,
    };
  },

  toJSON(message: Transaction): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    message.amount !== undefined && (obj.amount = message.amount);
    message.isoCurrencyCode !== undefined && (obj.isoCurrencyCode = message.isoCurrencyCode);
    message.unofficialCurrencyCode !== undefined && (obj.unofficialCurrencyCode = message.unofficialCurrencyCode);
    message.categoryId !== undefined && (obj.categoryId = message.categoryId);
    message.checkNumber !== undefined && (obj.checkNumber = message.checkNumber);
    message.date !== undefined && (obj.date = message.date);
    message.datetime !== undefined && (obj.datetime = message.datetime);
    message.authorizedDate !== undefined && (obj.authorizedDate = message.authorizedDate);
    message.authorizedDatetime !== undefined && (obj.authorizedDatetime = message.authorizedDatetime);
    message.name !== undefined && (obj.name = message.name);
    message.merchantName !== undefined && (obj.merchantName = message.merchantName);
    message.paymentChannel !== undefined && (obj.paymentChannel = message.paymentChannel);
    message.pending !== undefined && (obj.pending = message.pending);
    message.pendingTransactionId !== undefined && (obj.pendingTransactionId = message.pendingTransactionId);
    message.accountOwner !== undefined && (obj.accountOwner = message.accountOwner);
    message.transactionId !== undefined && (obj.transactionId = message.transactionId);
    message.transactionCode !== undefined && (obj.transactionCode = message.transactionCode);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.linkId !== undefined && (obj.linkId = Math.round(message.linkId));
    message.sign !== undefined && (obj.sign = Math.round(message.sign));
    message.personalFinanceCategoryPrimary !== undefined &&
      (obj.personalFinanceCategoryPrimary = message.personalFinanceCategoryPrimary);
    message.personalFinanceCategoryDetailed !== undefined &&
      (obj.personalFinanceCategoryDetailed = message.personalFinanceCategoryDetailed);
    message.locationAddress !== undefined && (obj.locationAddress = message.locationAddress);
    message.locationCity !== undefined && (obj.locationCity = message.locationCity);
    message.locationRegion !== undefined && (obj.locationRegion = message.locationRegion);
    message.locationPostalCode !== undefined && (obj.locationPostalCode = message.locationPostalCode);
    message.locationCountry !== undefined && (obj.locationCountry = message.locationCountry);
    message.locationLat !== undefined && (obj.locationLat = message.locationLat);
    message.locationLon !== undefined && (obj.locationLon = message.locationLon);
    message.locationStoreNumber !== undefined && (obj.locationStoreNumber = message.locationStoreNumber);
    message.paymentMetaByOrderOf !== undefined && (obj.paymentMetaByOrderOf = message.paymentMetaByOrderOf);
    message.paymentMetaPayee !== undefined && (obj.paymentMetaPayee = message.paymentMetaPayee);
    message.paymentMetaPayer !== undefined && (obj.paymentMetaPayer = message.paymentMetaPayer);
    message.paymentMetaPaymentMethod !== undefined && (obj.paymentMetaPaymentMethod = message.paymentMetaPaymentMethod);
    message.paymentMetaPaymentProcessor !== undefined &&
      (obj.paymentMetaPaymentProcessor = message.paymentMetaPaymentProcessor);
    message.paymentMetaPpdId !== undefined && (obj.paymentMetaPpdId = message.paymentMetaPpdId);
    message.paymentMetaReason !== undefined && (obj.paymentMetaReason = message.paymentMetaReason);
    message.paymentMetaReferenceNumber !== undefined &&
      (obj.paymentMetaReferenceNumber = message.paymentMetaReferenceNumber);
    message.time !== undefined && (obj.time = message.time.toISOString());
    message.additionalProperties !== undefined &&
      (obj.additionalProperties = message.additionalProperties ? Any.toJSON(message.additionalProperties) : undefined);
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
    message.categoryId = object.categoryId ?? "";
    message.checkNumber = object.checkNumber ?? "";
    message.date = object.date ?? "";
    message.datetime = object.datetime ?? "";
    message.authorizedDate = object.authorizedDate ?? "";
    message.authorizedDatetime = object.authorizedDatetime ?? "";
    message.name = object.name ?? "";
    message.merchantName = object.merchantName ?? "";
    message.paymentChannel = object.paymentChannel ?? "";
    message.pending = object.pending ?? false;
    message.pendingTransactionId = object.pendingTransactionId ?? "";
    message.accountOwner = object.accountOwner ?? "";
    message.transactionId = object.transactionId ?? "";
    message.transactionCode = object.transactionCode ?? "";
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.linkId = object.linkId ?? 0;
    message.sign = object.sign ?? 0;
    message.personalFinanceCategoryPrimary = object.personalFinanceCategoryPrimary ?? "";
    message.personalFinanceCategoryDetailed = object.personalFinanceCategoryDetailed ?? "";
    message.locationAddress = object.locationAddress ?? "";
    message.locationCity = object.locationCity ?? "";
    message.locationRegion = object.locationRegion ?? "";
    message.locationPostalCode = object.locationPostalCode ?? "";
    message.locationCountry = object.locationCountry ?? "";
    message.locationLat = object.locationLat ?? 0;
    message.locationLon = object.locationLon ?? 0;
    message.locationStoreNumber = object.locationStoreNumber ?? "";
    message.paymentMetaByOrderOf = object.paymentMetaByOrderOf ?? "";
    message.paymentMetaPayee = object.paymentMetaPayee ?? "";
    message.paymentMetaPayer = object.paymentMetaPayer ?? "";
    message.paymentMetaPaymentMethod = object.paymentMetaPaymentMethod ?? "";
    message.paymentMetaPaymentProcessor = object.paymentMetaPaymentProcessor ?? "";
    message.paymentMetaPpdId = object.paymentMetaPpdId ?? "";
    message.paymentMetaReason = object.paymentMetaReason ?? "";
    message.paymentMetaReferenceNumber = object.paymentMetaReferenceNumber ?? "";
    message.time = object.time ?? undefined;
    message.additionalProperties = (object.additionalProperties !== undefined && object.additionalProperties !== null)
      ? Any.fromPartial(object.additionalProperties)
      : undefined;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
