/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "api.v1";

export enum BankAccountStatus {
  BANK_ACCOUNT_STATUS_UNSPECIFIED = 0,
  BANK_ACCOUNT_STATUS_ACTIVE = 1,
  BANK_ACCOUNT_STATUS_INACTIVE = 2,
  UNRECOGNIZED = -1,
}

export function bankAccountStatusFromJSON(object: any): BankAccountStatus {
  switch (object) {
    case 0:
    case "BANK_ACCOUNT_STATUS_UNSPECIFIED":
      return BankAccountStatus.BANK_ACCOUNT_STATUS_UNSPECIFIED;
    case 1:
    case "BANK_ACCOUNT_STATUS_ACTIVE":
      return BankAccountStatus.BANK_ACCOUNT_STATUS_ACTIVE;
    case 2:
    case "BANK_ACCOUNT_STATUS_INACTIVE":
      return BankAccountStatus.BANK_ACCOUNT_STATUS_INACTIVE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return BankAccountStatus.UNRECOGNIZED;
  }
}

export function bankAccountStatusToJSON(object: BankAccountStatus): string {
  switch (object) {
    case BankAccountStatus.BANK_ACCOUNT_STATUS_UNSPECIFIED:
      return "BANK_ACCOUNT_STATUS_UNSPECIFIED";
    case BankAccountStatus.BANK_ACCOUNT_STATUS_ACTIVE:
      return "BANK_ACCOUNT_STATUS_ACTIVE";
    case BankAccountStatus.BANK_ACCOUNT_STATUS_INACTIVE:
      return "BANK_ACCOUNT_STATUS_INACTIVE";
    case BankAccountStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum GoalStatus {
  GOAL_STATUS_UNSPECIFIED = 0,
  GOAL_STATUS_ACTIVE = 1,
  GOAL_STATUS_INACTIVE = 2,
  GOAL_STATUS_COMPLETED = 3,
  GOAL_STATUS_DELETE = 4,
  UNRECOGNIZED = -1,
}

export function goalStatusFromJSON(object: any): GoalStatus {
  switch (object) {
    case 0:
    case "GOAL_STATUS_UNSPECIFIED":
      return GoalStatus.GOAL_STATUS_UNSPECIFIED;
    case 1:
    case "GOAL_STATUS_ACTIVE":
      return GoalStatus.GOAL_STATUS_ACTIVE;
    case 2:
    case "GOAL_STATUS_INACTIVE":
      return GoalStatus.GOAL_STATUS_INACTIVE;
    case 3:
    case "GOAL_STATUS_COMPLETED":
      return GoalStatus.GOAL_STATUS_COMPLETED;
    case 4:
    case "GOAL_STATUS_DELETE":
      return GoalStatus.GOAL_STATUS_DELETE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return GoalStatus.UNRECOGNIZED;
  }
}

export function goalStatusToJSON(object: GoalStatus): string {
  switch (object) {
    case GoalStatus.GOAL_STATUS_UNSPECIFIED:
      return "GOAL_STATUS_UNSPECIFIED";
    case GoalStatus.GOAL_STATUS_ACTIVE:
      return "GOAL_STATUS_ACTIVE";
    case GoalStatus.GOAL_STATUS_INACTIVE:
      return "GOAL_STATUS_INACTIVE";
    case GoalStatus.GOAL_STATUS_COMPLETED:
      return "GOAL_STATUS_COMPLETED";
    case GoalStatus.GOAL_STATUS_DELETE:
      return "GOAL_STATUS_DELETE";
    case GoalStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum GoalType {
  GOAL_TYPE_UNSPECIFIED = 0,
  GOAL_TYPE_SAVINGS = 1,
  GOAL_TYPE_INVESTMENT = 2,
  GOAL_TYPE_DEBT = 3,
  GOAL_TYPE_EXPENSE = 4,
  UNRECOGNIZED = -1,
}

export function goalTypeFromJSON(object: any): GoalType {
  switch (object) {
    case 0:
    case "GOAL_TYPE_UNSPECIFIED":
      return GoalType.GOAL_TYPE_UNSPECIFIED;
    case 1:
    case "GOAL_TYPE_SAVINGS":
      return GoalType.GOAL_TYPE_SAVINGS;
    case 2:
    case "GOAL_TYPE_INVESTMENT":
      return GoalType.GOAL_TYPE_INVESTMENT;
    case 3:
    case "GOAL_TYPE_DEBT":
      return GoalType.GOAL_TYPE_DEBT;
    case 4:
    case "GOAL_TYPE_EXPENSE":
      return GoalType.GOAL_TYPE_EXPENSE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return GoalType.UNRECOGNIZED;
  }
}

export function goalTypeToJSON(object: GoalType): string {
  switch (object) {
    case GoalType.GOAL_TYPE_UNSPECIFIED:
      return "GOAL_TYPE_UNSPECIFIED";
    case GoalType.GOAL_TYPE_SAVINGS:
      return "GOAL_TYPE_SAVINGS";
    case GoalType.GOAL_TYPE_INVESTMENT:
      return "GOAL_TYPE_INVESTMENT";
    case GoalType.GOAL_TYPE_DEBT:
      return "GOAL_TYPE_DEBT";
    case GoalType.GOAL_TYPE_EXPENSE:
      return "GOAL_TYPE_EXPENSE";
    case GoalType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum PocketType {
  POCKET_TYPE_UNSPECIFIED = 0,
  POCKET_TYPE_DISCRETIONARY_SPENDING = 1,
  POCKET_TYPE_FUN_MONEY = 2,
  POCKET_TYPE_DEBT_REDUCTION = 3,
  POCKET_TYPE_EMERGENCY_FUND = 4,
  POCKET_TYPE_INVESTMENT = 5,
  POCKET_TYPE_SHORT_TERM_SAVINGS = 6,
  POCKET_TYPE_LONG_TERM_SAVINGS = 7,
  UNRECOGNIZED = -1,
}

export function pocketTypeFromJSON(object: any): PocketType {
  switch (object) {
    case 0:
    case "POCKET_TYPE_UNSPECIFIED":
      return PocketType.POCKET_TYPE_UNSPECIFIED;
    case 1:
    case "POCKET_TYPE_DISCRETIONARY_SPENDING":
      return PocketType.POCKET_TYPE_DISCRETIONARY_SPENDING;
    case 2:
    case "POCKET_TYPE_FUN_MONEY":
      return PocketType.POCKET_TYPE_FUN_MONEY;
    case 3:
    case "POCKET_TYPE_DEBT_REDUCTION":
      return PocketType.POCKET_TYPE_DEBT_REDUCTION;
    case 4:
    case "POCKET_TYPE_EMERGENCY_FUND":
      return PocketType.POCKET_TYPE_EMERGENCY_FUND;
    case 5:
    case "POCKET_TYPE_INVESTMENT":
      return PocketType.POCKET_TYPE_INVESTMENT;
    case 6:
    case "POCKET_TYPE_SHORT_TERM_SAVINGS":
      return PocketType.POCKET_TYPE_SHORT_TERM_SAVINGS;
    case 7:
    case "POCKET_TYPE_LONG_TERM_SAVINGS":
      return PocketType.POCKET_TYPE_LONG_TERM_SAVINGS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PocketType.UNRECOGNIZED;
  }
}

export function pocketTypeToJSON(object: PocketType): string {
  switch (object) {
    case PocketType.POCKET_TYPE_UNSPECIFIED:
      return "POCKET_TYPE_UNSPECIFIED";
    case PocketType.POCKET_TYPE_DISCRETIONARY_SPENDING:
      return "POCKET_TYPE_DISCRETIONARY_SPENDING";
    case PocketType.POCKET_TYPE_FUN_MONEY:
      return "POCKET_TYPE_FUN_MONEY";
    case PocketType.POCKET_TYPE_DEBT_REDUCTION:
      return "POCKET_TYPE_DEBT_REDUCTION";
    case PocketType.POCKET_TYPE_EMERGENCY_FUND:
      return "POCKET_TYPE_EMERGENCY_FUND";
    case PocketType.POCKET_TYPE_INVESTMENT:
      return "POCKET_TYPE_INVESTMENT";
    case PocketType.POCKET_TYPE_SHORT_TERM_SAVINGS:
      return "POCKET_TYPE_SHORT_TERM_SAVINGS";
    case PocketType.POCKET_TYPE_LONG_TERM_SAVINGS:
      return "POCKET_TYPE_LONG_TERM_SAVINGS";
    case PocketType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum BankAccountType {
  BANK_ACCOUNT_TYPE_UNSPECIFIED = 0,
  BANK_ACCOUNT_TYPE_PLAID = 1,
  BANK_ACCOUNT_TYPE_MANUAL = 2,
  UNRECOGNIZED = -1,
}

export function bankAccountTypeFromJSON(object: any): BankAccountType {
  switch (object) {
    case 0:
    case "BANK_ACCOUNT_TYPE_UNSPECIFIED":
      return BankAccountType.BANK_ACCOUNT_TYPE_UNSPECIFIED;
    case 1:
    case "BANK_ACCOUNT_TYPE_PLAID":
      return BankAccountType.BANK_ACCOUNT_TYPE_PLAID;
    case 2:
    case "BANK_ACCOUNT_TYPE_MANUAL":
      return BankAccountType.BANK_ACCOUNT_TYPE_MANUAL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return BankAccountType.UNRECOGNIZED;
  }
}

export function bankAccountTypeToJSON(object: BankAccountType): string {
  switch (object) {
    case BankAccountType.BANK_ACCOUNT_TYPE_UNSPECIFIED:
      return "BANK_ACCOUNT_TYPE_UNSPECIFIED";
    case BankAccountType.BANK_ACCOUNT_TYPE_PLAID:
      return "BANK_ACCOUNT_TYPE_PLAID";
    case BankAccountType.BANK_ACCOUNT_TYPE_MANUAL:
      return "BANK_ACCOUNT_TYPE_MANUAL";
    case BankAccountType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum StripeSubscriptionStatus {
  STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED = 0,
  STRIPE_SUBSCRIPTION_STATUS_TRIALING = 1,
  STRIPE_SUBSCRIPTION_STATUS_ACTIVE = 2,
  STRIPE_SUBSCRIPTION_STATUS_PAST_DUE = 3,
  STRIPE_SUBSCRIPTION_STATUS_CANCELED = 4,
  STRIPE_SUBSCRIPTION_STATUS_UNPAID = 5,
  STRIPE_SUBSCRIPTION_STATUS_COMPLETE = 6,
  STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE = 7,
  STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED = 8,
  STRIPE_SUBSCRIPTION_STATUS_CREATED = 9,
  STRIPE_SUBSCRIPTION_STATUS_PAUSED = 10,
  UNRECOGNIZED = -1,
}

export function stripeSubscriptionStatusFromJSON(object: any): StripeSubscriptionStatus {
  switch (object) {
    case 0:
    case "STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED;
    case 1:
    case "STRIPE_SUBSCRIPTION_STATUS_TRIALING":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_TRIALING;
    case 2:
    case "STRIPE_SUBSCRIPTION_STATUS_ACTIVE":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_ACTIVE;
    case 3:
    case "STRIPE_SUBSCRIPTION_STATUS_PAST_DUE":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_PAST_DUE;
    case 4:
    case "STRIPE_SUBSCRIPTION_STATUS_CANCELED":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_CANCELED;
    case 5:
    case "STRIPE_SUBSCRIPTION_STATUS_UNPAID":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_UNPAID;
    case 6:
    case "STRIPE_SUBSCRIPTION_STATUS_COMPLETE":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_COMPLETE;
    case 7:
    case "STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE;
    case 8:
    case "STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED;
    case 9:
    case "STRIPE_SUBSCRIPTION_STATUS_CREATED":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_CREATED;
    case 10:
    case "STRIPE_SUBSCRIPTION_STATUS_PAUSED":
      return StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_PAUSED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StripeSubscriptionStatus.UNRECOGNIZED;
  }
}

export function stripeSubscriptionStatusToJSON(object: StripeSubscriptionStatus): string {
  switch (object) {
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED:
      return "STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_TRIALING:
      return "STRIPE_SUBSCRIPTION_STATUS_TRIALING";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_ACTIVE:
      return "STRIPE_SUBSCRIPTION_STATUS_ACTIVE";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_PAST_DUE:
      return "STRIPE_SUBSCRIPTION_STATUS_PAST_DUE";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_CANCELED:
      return "STRIPE_SUBSCRIPTION_STATUS_CANCELED";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_UNPAID:
      return "STRIPE_SUBSCRIPTION_STATUS_UNPAID";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_COMPLETE:
      return "STRIPE_SUBSCRIPTION_STATUS_COMPLETE";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE:
      return "STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED:
      return "STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_CREATED:
      return "STRIPE_SUBSCRIPTION_STATUS_CREATED";
    case StripeSubscriptionStatus.STRIPE_SUBSCRIPTION_STATUS_PAUSED:
      return "STRIPE_SUBSCRIPTION_STATUS_PAUSED";
    case StripeSubscriptionStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum LinkStatus {
  LINK_STATUS_UNSPECIFIED = 0,
  LINK_STATUS_SETUP = 1,
  LINK_STATUS_PENDING = 2,
  LINK_STATUS_ERROR = 3,
  LINK_STATUS_SUCCESS = 4,
  LINK_STATUS_PENDING_EXPIRATION = 5,
  LINK_STATUS_REVOKED = 6,
  UNRECOGNIZED = -1,
}

export function linkStatusFromJSON(object: any): LinkStatus {
  switch (object) {
    case 0:
    case "LINK_STATUS_UNSPECIFIED":
      return LinkStatus.LINK_STATUS_UNSPECIFIED;
    case 1:
    case "LINK_STATUS_SETUP":
      return LinkStatus.LINK_STATUS_SETUP;
    case 2:
    case "LINK_STATUS_PENDING":
      return LinkStatus.LINK_STATUS_PENDING;
    case 3:
    case "LINK_STATUS_ERROR":
      return LinkStatus.LINK_STATUS_ERROR;
    case 4:
    case "LINK_STATUS_SUCCESS":
      return LinkStatus.LINK_STATUS_SUCCESS;
    case 5:
    case "LINK_STATUS_PENDING_EXPIRATION":
      return LinkStatus.LINK_STATUS_PENDING_EXPIRATION;
    case 6:
    case "LINK_STATUS_REVOKED":
      return LinkStatus.LINK_STATUS_REVOKED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return LinkStatus.UNRECOGNIZED;
  }
}

export function linkStatusToJSON(object: LinkStatus): string {
  switch (object) {
    case LinkStatus.LINK_STATUS_UNSPECIFIED:
      return "LINK_STATUS_UNSPECIFIED";
    case LinkStatus.LINK_STATUS_SETUP:
      return "LINK_STATUS_SETUP";
    case LinkStatus.LINK_STATUS_PENDING:
      return "LINK_STATUS_PENDING";
    case LinkStatus.LINK_STATUS_ERROR:
      return "LINK_STATUS_ERROR";
    case LinkStatus.LINK_STATUS_SUCCESS:
      return "LINK_STATUS_SUCCESS";
    case LinkStatus.LINK_STATUS_PENDING_EXPIRATION:
      return "LINK_STATUS_PENDING_EXPIRATION";
    case LinkStatus.LINK_STATUS_REVOKED:
      return "LINK_STATUS_REVOKED";
    case LinkStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum LinkType {
  LINK_TYPE_UNSPECIFIED = 0,
  LINK_TYPE_PLAID = 1,
  LINK_TYPE_MANUAL = 2,
  UNRECOGNIZED = -1,
}

export function linkTypeFromJSON(object: any): LinkType {
  switch (object) {
    case 0:
    case "LINK_TYPE_UNSPECIFIED":
      return LinkType.LINK_TYPE_UNSPECIFIED;
    case 1:
    case "LINK_TYPE_PLAID":
      return LinkType.LINK_TYPE_PLAID;
    case 2:
    case "LINK_TYPE_MANUAL":
      return LinkType.LINK_TYPE_MANUAL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return LinkType.UNRECOGNIZED;
  }
}

export function linkTypeToJSON(object: LinkType): string {
  switch (object) {
    case LinkType.LINK_TYPE_UNSPECIFIED:
      return "LINK_TYPE_UNSPECIFIED";
    case LinkType.LINK_TYPE_PLAID:
      return "LINK_TYPE_PLAID";
    case LinkType.LINK_TYPE_MANUAL:
      return "LINK_TYPE_MANUAL";
    case LinkType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** StripeSubscription stores high level stripe subscription details of which the user profile has */
export interface StripeSubscription {
  id: number;
  /** stripe subscription id tied to the customer */
  stripeSubscriptionId: string;
  /** stripe subscription status */
  stripeSubscriptionStatus: StripeSubscriptionStatus;
  /** stripe subscription active until */
  stripeSubscriptionActiveUntil: string;
  /** stripe webhook latest timestamp */
  stripeWebhookLatestTimestamp: string;
  /** wether the subscription is trialing */
  isTrialing: boolean;
}

/**
 * UserProfile stores high level user profile details
 * such as the id, user_id tied to the profile, and many more
 */
export interface UserProfile {
  /** id */
  id: number;
  /** the user id tied to the profile */
  userId: number;
  stripeCustomerId: string;
  /** the stripe subscriptions the user profile actively maintains */
  stripeSubscriptions:
    | StripeSubscription
    | undefined;
  /** a user profile can have many links (connected institutions) of which finanical accounts are tied to (checking, savings, etc) */
  link: Link[];
}

/**
 * A Link represents a login at a financial institution. A single end-user of your application might have accounts at different financial
 * institutions, which means they would have multiple different Items. An Item is not the same as a financial institution account,
 * although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access
 * both their checking account and their savings account, a single Item would be associated with both of those accounts. Each Item
 * linked within your application will have a corresponding access_token, which is a token that you can use to make API requests related
 * to that specific Item.
 * Two Items created for the same set of credentials at the same institution will be considered different and not share the same item_id.
 */
export interface Link {
  /** id */
  id: number;
  linkStatus: LinkStatus;
  plaidLink: PlaidLink | undefined;
  plaidNewAccountsAvailable: boolean;
  expirationDate: string;
  institutionName: string;
  customInstitutionName: string;
  description: string;
  lastManualSync: string;
  lastSuccessfulUpdate: string;
  /**
   * token object witholds an access token which is a token used to make API requests related to a specific Item. You will typically obtain an access_token
   * by calling /item/public_token/exchange. For more details, see the Token exchange flow. An access_token does not expire,
   * although it may require updating, such as when a user changes their password, or when working with European institutions
   * that comply with PSD2's 90-day consent window. For more information, see When to use update mode.
   * Access tokens should always be stored securely, and associated with the user whose data they represent.
   * If compromised, an access_token can be rotated via /item/access_token/invalidate. If no longer needed,
   * it can be revoked via /item/remove.(gorm.field).has_one = {disable_association_autocreate: false disable_association_autoupdate: false preload: true}];
   */
  token:
    | Token
    | undefined;
  /**
   * a link event - or client login event can have many connected bank accounts
   * for example a log in link against one instition like chase can have many account (checking and savings)
   * it is important though to ensure that if a link against an instition already exists, we dont fascilitate duplicated
   */
  bankAccounts: BankAccount[];
  /**
   * a link event - or client login event can have many connected investment accounts
   * for example a log in link against one instition like fidelity can have many accounts (401k and investment account)
   * it is important though to ensure that if a link against an instition already exists, we dont fascilitate duplicated
   */
  investmentAccounts: InvestmentAccount[];
  /** credit accounts tied to a user */
  creditAccounts: CreditAccount[];
  /** mortgage accounts tied to a user */
  mortgageAccounts: MortgageAccount[];
  /** student loan accounts tied to a link */
  studentLoanAccounts: StudentLoanAccount[];
  /** the id of the institution this link is tied to and against */
  plaidInstitutionId: string;
  /** the type of link this is ... can be either a manual or plaid link type */
  linkType: LinkType;
  errorCode: string;
  updatedAt: string;
  newAccountsAvailable: boolean;
  plaidSync: PlaidSync | undefined;
}

export interface PlaidSync {
  /** id */
  id: number;
  timeStamp: string;
  trigger: string;
  nextCursor: string;
  added: number;
  removed: number;
  modified: number;
}

export interface Token {
  /** id */
  id: number;
  /** the id of the item the token is tied to */
  itemId: string;
  keyId: string;
  accessToken: string;
  version: string;
}

export interface PlaidLink {
  /** id */
  id: number;
  products: string[];
  webhookUrl: string;
  institutionId: string;
  institutionName: string;
  usePlaidSync: boolean;
  itemId: string;
}

export interface StudentLoanAccount {
  /** id */
  id: number;
  plaidAccountId: string;
  disbursementDates: string[];
  expectedPayoffDate: string;
  guarantor: string;
  interestRatePercentage: number;
  isOverdue: boolean;
  lastPaymentAmount: number;
  lastPaymentDate: string;
  lastStatementIssueDate: string;
  loanName: string;
  loanEndDate: string;
  minimumPaymentAmount: number;
  nextPaymentDueDate: string;
  originationDate: string;
  originationPrincipalAmount: number;
  outstandingInterestAmount: number;
  paymentReferenceNumber: string;
  sequenceNumber: string;
  ytdInterestPaid: number;
  ytdPrincipalPaid: number;
  loanType: string;
  pslfStatusEstimatedEligibilityDate: string;
  pslfStatusPaymentsMade: number;
  pslfStatusPaymentsRemaining: number;
  repaymentPlanType: string;
  repaymentPlanDescription: string;
  servicerAddressCity: string;
  servicerAddressPostalCode: string;
  servicerAddressState: string;
  servicerAddressStreet: string;
  servicerAddressRegion: string;
  servicerAddressCountry: string;
  /** the user id to which this bank account is tied to */
  userId: number;
  /** the account name */
  name: string;
}

export interface CreditAccount {
  /** id */
  id: number;
  /** the user id to which this bank account is tied to */
  userId: number;
  /** the account name */
  name: string;
  /** the bank account number */
  number: string;
  /** the bank account type */
  type: string;
  /** the bank account balance */
  balance: number;
  /** current funds on the account */
  currentFunds: number;
  /** balance limit */
  balanceLimit: number;
  /** plaid account id mapped to this bank account */
  plaidAccountId: string;
  /** accoint subtype */
  subtype: string;
  /** wether the account is overdue */
  isOverdue: boolean;
  /** the last payment amount */
  lastPaymentAmount: number;
  /** the last payment date */
  lastPaymentDate: string;
  /** the last statement issue date */
  lastStatementIssueDate: string;
  /** the minimum amount due date */
  minimumAmountDueDate: number;
  /** the next payment date */
  nextPaymentDate: string;
  /** the aprs */
  aprs: Apr[];
  /** the last statement balance */
  lastStatementBalance: number;
  /** the minimum payment amount */
  minimumPaymentAmount: number;
  /** the next payment due date */
  nextPaymentDueDate: string;
}

export interface MortgageAccount {
  id: number;
  plaidAccountId: string;
  accountNumber: string;
  currentLateFee: number;
  escrowBalance: number;
  hasPmi: boolean;
  hasPrepaymentPenalty: boolean;
  lastPaymentAmount: number;
  lastPaymentDate: string;
  loanTerm: string;
  loanTypeDescription: string;
  maturityDate: string;
  nextMonthlyPayment: number;
  nextPaymentDueDate: string;
  originalPrincipalBalance: number;
  originalPropertyValue: number;
  outstandingPrincipalBalance: number;
  paymentAmount: number;
  paymentDate: string;
  originationDate: string;
  originationPrincipalAmount: number;
  pastDueAmount: number;
  ytdInterestPaid: number;
  ytdPrincipalPaid: number;
  propertyAddressCity: string;
  propertyAddressState: string;
  propertyAddressStreet: string;
  propertyAddressPostalCode: string;
  propertyRegion: string;
  propertyCountry: string;
  interestRatePercentage: number;
  interestRateType: string;
}

export interface InvestmentAccount {
  /** id */
  id: number;
  /** the user id to which this bank account is tied to */
  userId: number;
  /** the account name */
  name: string;
  /** the bank account number */
  number: string;
  /** the bank account type */
  type: string;
  /** the bank account balance */
  balance: number;
  currentFunds: number;
  balanceLimit: number;
  /** plaid account id mapped to this bank account */
  plaidAccountId: string;
  /** accoint subtype */
  subtype: string;
  /** invesment holding is the set of securities this account witholds */
  holdings: InvesmentHolding[];
  /** the set of securities this account witholds */
  securities: InvestmentSecurity[];
}

export interface BankAccount {
  /** id */
  id: number;
  /** the user id to which this bank account is tied to */
  userId: number;
  /** the bank account name */
  name: string;
  /** the bank account number */
  number: string;
  /** the bank account type */
  type: BankAccountType;
  /** the bank account balance */
  balance: number;
  /** the bank account currency */
  currency: string;
  currentFunds: number;
  balanceLimit: number;
  /**
   * the set of "virtualized accounts this user witholds"
   * NOTE: these pockets are automatically created by the system
   * when a user connects a bank account
   */
  pockets: Pocket[];
  /** plaid account id mapped to this bank account */
  plaidAccountId: string;
  /** account subtype */
  subtype: string;
  /** the bank account status */
  status: BankAccountStatus;
}

/**
 * Pocket is an abstraction of a over a bank account. A user can has at most 4 pockets per connected account
 * NOTE: these pockets are automatically created by the system and should not be exposed for mutation
 * by any client. The only operations that can be performed against a pocket are:
 * 1. Get the pocket
 * 2. Get the pocket's smart goals
 * 3. Adding a smart goal to the pocket
 */
export interface Pocket {
  /** id */
  id: number;
  /** the set of smart goals this user witholds */
  goals: SmartGoal[];
  /** The type of the pocket */
  type: PocketType;
}

/**
 * SmartGoal: The Goals table stores information about each financial goal, including the name of the goal,
 * its description, the target amount of money the user wants to save or invest, and the expected date of completion.
 *
 * The Goals table also includes columns for the start date of the goal, the current amount of money saved or
 * invested towards the goal, and a boolean flag indicating whether the goal has been achieved.
 * These additional columns allow the user to track their progress towards the goal and see how much
 * more they need to save or invest to reach their target amount.
 */
export interface SmartGoal {
  /** id */
  id: number;
  /** the user id to which this goal is tied to */
  userId: number;
  /**
   * The name of the goal
   * Validations:
   * - must be at least 3 characters long
   */
  name: string;
  /**
   * The description of the goal
   * Validations:
   * - must be at least 3 characters long
   */
  description: string;
  /** wether the goal has been achieved or not */
  isCompleted: boolean;
  /** The type of the goal */
  goalType: GoalType;
  /** The duration of the goal */
  duration: string;
  /** the start date of the goal */
  startDate: string;
  /** the end date of the goal */
  endDate: string;
  /**
   * the target amount of the goal
   * amount of money the user wants to save or invest
   */
  targetAmount: string;
  /**
   * the current amount of the goal
   * current amount of money saved or invested towards the goal
   */
  currentAmount: string;
  /** Milestones associated with the goal */
  milestones: Milestone[];
  /** Forecasts associated with the goal */
  forecasts: Forecast | undefined;
}

/**
 * The Forecast table stores information about each forecast generated for a particular goal,
 * including the forecast date, the forecasted amount of money saved or invested for the
 * goal by the target date, and the variance between the forecasted and target amounts.
 * This allows the user to track how well they are progressing towards their goal and make adjustments as needed.
 */
export interface Forecast {
  /** id */
  id: number;
  /** the forecasted amount of the goal */
  forecastedAmount: string;
  /** the forecasted completion date of the goal */
  forecastedCompletionDate: string;
  /** the forecasted variance of the goal between the forecasted and target amounts */
  varianceAmount: string;
}

/**
 * Milestone: represents a milestone in the context of simfinni. A financial milestone that is both smart
 * and achievable. A milestone is a sub goal of a goal and is tied to a goal by the goal id
 */
export interface Milestone {
  /** id */
  id: number;
  /**
   * The name of the milestone
   * Validations:
   * - must be at least 3 characters long
   */
  name: string;
  /**
   * The description of the miletone
   * Validations:
   * - must be at least 3 characters long
   */
  description: string;
  /**
   * the target date of the milestone
   * Validations:
   * - must be at least 3 characters long
   */
  targetDate: string;
  /** the target amount of the milestone */
  targetAmount: string;
  /** wethe milestone is completed or not */
  isCompleted: boolean;
  /** the budget associated with the milestone */
  budget: Budget | undefined;
}

/**
 * The Budgets table stores information about each budget created by the user,
 * including the name of the budget, the start and end dates, and the user ID.
 */
export interface Budget {
  /** id */
  id: number;
  /** The name of the budget */
  name: string;
  description: string;
  /** the time the goal was created */
  startDate: string;
  /** the time the goal was updated */
  endDate: string;
  /** category associated with the goal */
  category: Category | undefined;
}

/**
 * The Categories table stores information about the different categories of expenses or income,
 * such as "Housing", "Food", "Transportation", and "Entertainment". Each category has one or more
 * subcategories, which are stored in the Subcategories table.
 *
 * For example, the "Housing" category might have subcategories for "Rent", "Utilities", and "Home Maintenance".
 */
export interface Category {
  /** id */
  id: number;
  /** The name of the category */
  name: string;
  /** The description of the category */
  description: string;
  /** the sub categories of the category */
  subcategories: string[];
}

export interface InvesmentHolding {
  /** id */
  id: number;
  /** The name of the investment holding */
  name: string;
  /** plaid account id */
  plaidAccountId: string;
  costBasis: number;
  institutionPrice: number;
  institutionPriceAsOf: string;
  institutionPriceDatetime: string;
  institutionValue: number;
  isoCurrencyCode: string;
  quantity: number;
  securityId: string;
  unofficialCurrencyCode: string;
}

export interface InvestmentSecurity {
  /** id */
  id: number;
  closePrice: number;
  closePriceAsOf: string;
  cusip: string;
  institutionId: string;
  institutionSecurityId: string;
  isCashEquivalent: boolean;
  isin: string;
  isoCurrencyCode: string;
  name: string;
  proxySecurityId: string;
  securityId: string;
  sedol: string;
  tickerSymbol: string;
  type: string;
  unofficialCurrencyCode: string;
  updateDatetime: string;
}

export interface Apr {
  id: number;
  percentage: number;
  type: string;
  balanceSubjectToApr: number;
  interestChargeAmount: number;
}

function createBaseStripeSubscription(): StripeSubscription {
  return {
    id: 0,
    stripeSubscriptionId: "",
    stripeSubscriptionStatus: 0,
    stripeSubscriptionActiveUntil: "",
    stripeWebhookLatestTimestamp: "",
    isTrialing: false,
  };
}

export const StripeSubscription = {
  encode(message: StripeSubscription, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.stripeSubscriptionId !== "") {
      writer.uint32(18).string(message.stripeSubscriptionId);
    }
    if (message.stripeSubscriptionStatus !== 0) {
      writer.uint32(24).int32(message.stripeSubscriptionStatus);
    }
    if (message.stripeSubscriptionActiveUntil !== "") {
      writer.uint32(34).string(message.stripeSubscriptionActiveUntil);
    }
    if (message.stripeWebhookLatestTimestamp !== "") {
      writer.uint32(42).string(message.stripeWebhookLatestTimestamp);
    }
    if (message.isTrialing === true) {
      writer.uint32(48).bool(message.isTrialing);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StripeSubscription {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStripeSubscription();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.stripeSubscriptionId = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.stripeSubscriptionStatus = reader.int32() as any;
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.stripeSubscriptionActiveUntil = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.stripeWebhookLatestTimestamp = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.isTrialing = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): StripeSubscription {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      stripeSubscriptionId: isSet(object.stripeSubscriptionId) ? String(object.stripeSubscriptionId) : "",
      stripeSubscriptionStatus: isSet(object.stripeSubscriptionStatus)
        ? stripeSubscriptionStatusFromJSON(object.stripeSubscriptionStatus)
        : 0,
      stripeSubscriptionActiveUntil: isSet(object.stripeSubscriptionActiveUntil)
        ? String(object.stripeSubscriptionActiveUntil)
        : "",
      stripeWebhookLatestTimestamp: isSet(object.stripeWebhookLatestTimestamp)
        ? String(object.stripeWebhookLatestTimestamp)
        : "",
      isTrialing: isSet(object.isTrialing) ? Boolean(object.isTrialing) : false,
    };
  },

  toJSON(message: StripeSubscription): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.stripeSubscriptionId !== undefined && (obj.stripeSubscriptionId = message.stripeSubscriptionId);
    message.stripeSubscriptionStatus !== undefined &&
      (obj.stripeSubscriptionStatus = stripeSubscriptionStatusToJSON(message.stripeSubscriptionStatus));
    message.stripeSubscriptionActiveUntil !== undefined &&
      (obj.stripeSubscriptionActiveUntil = message.stripeSubscriptionActiveUntil);
    message.stripeWebhookLatestTimestamp !== undefined &&
      (obj.stripeWebhookLatestTimestamp = message.stripeWebhookLatestTimestamp);
    message.isTrialing !== undefined && (obj.isTrialing = message.isTrialing);
    return obj;
  },

  create<I extends Exact<DeepPartial<StripeSubscription>, I>>(base?: I): StripeSubscription {
    return StripeSubscription.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StripeSubscription>, I>>(object: I): StripeSubscription {
    const message = createBaseStripeSubscription();
    message.id = object.id ?? 0;
    message.stripeSubscriptionId = object.stripeSubscriptionId ?? "";
    message.stripeSubscriptionStatus = object.stripeSubscriptionStatus ?? 0;
    message.stripeSubscriptionActiveUntil = object.stripeSubscriptionActiveUntil ?? "";
    message.stripeWebhookLatestTimestamp = object.stripeWebhookLatestTimestamp ?? "";
    message.isTrialing = object.isTrialing ?? false;
    return message;
  },
};

function createBaseUserProfile(): UserProfile {
  return { id: 0, userId: 0, stripeCustomerId: "", stripeSubscriptions: undefined, link: [] };
}

export const UserProfile = {
  encode(message: UserProfile, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.stripeCustomerId !== "") {
      writer.uint32(26).string(message.stripeCustomerId);
    }
    if (message.stripeSubscriptions !== undefined) {
      StripeSubscription.encode(message.stripeSubscriptions, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.link) {
      Link.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserProfile {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserProfile();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.stripeCustomerId = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.stripeSubscriptions = StripeSubscription.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.link.push(Link.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserProfile {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      stripeCustomerId: isSet(object.stripeCustomerId) ? String(object.stripeCustomerId) : "",
      stripeSubscriptions: isSet(object.stripeSubscriptions)
        ? StripeSubscription.fromJSON(object.stripeSubscriptions)
        : undefined,
      link: Array.isArray(object?.link) ? object.link.map((e: any) => Link.fromJSON(e)) : [],
    };
  },

  toJSON(message: UserProfile): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.stripeCustomerId !== undefined && (obj.stripeCustomerId = message.stripeCustomerId);
    message.stripeSubscriptions !== undefined && (obj.stripeSubscriptions = message.stripeSubscriptions
      ? StripeSubscription.toJSON(message.stripeSubscriptions)
      : undefined);
    if (message.link) {
      obj.link = message.link.map((e) => e ? Link.toJSON(e) : undefined);
    } else {
      obj.link = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UserProfile>, I>>(base?: I): UserProfile {
    return UserProfile.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserProfile>, I>>(object: I): UserProfile {
    const message = createBaseUserProfile();
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.stripeCustomerId = object.stripeCustomerId ?? "";
    message.stripeSubscriptions = (object.stripeSubscriptions !== undefined && object.stripeSubscriptions !== null)
      ? StripeSubscription.fromPartial(object.stripeSubscriptions)
      : undefined;
    message.link = object.link?.map((e) => Link.fromPartial(e)) || [];
    return message;
  },
};

function createBaseLink(): Link {
  return {
    id: 0,
    linkStatus: 0,
    plaidLink: undefined,
    plaidNewAccountsAvailable: false,
    expirationDate: "",
    institutionName: "",
    customInstitutionName: "",
    description: "",
    lastManualSync: "",
    lastSuccessfulUpdate: "",
    token: undefined,
    bankAccounts: [],
    investmentAccounts: [],
    creditAccounts: [],
    mortgageAccounts: [],
    studentLoanAccounts: [],
    plaidInstitutionId: "",
    linkType: 0,
    errorCode: "",
    updatedAt: "",
    newAccountsAvailable: false,
    plaidSync: undefined,
  };
}

export const Link = {
  encode(message: Link, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.linkStatus !== 0) {
      writer.uint32(24).int32(message.linkStatus);
    }
    if (message.plaidLink !== undefined) {
      PlaidLink.encode(message.plaidLink, writer.uint32(34).fork()).ldelim();
    }
    if (message.plaidNewAccountsAvailable === true) {
      writer.uint32(40).bool(message.plaidNewAccountsAvailable);
    }
    if (message.expirationDate !== "") {
      writer.uint32(50).string(message.expirationDate);
    }
    if (message.institutionName !== "") {
      writer.uint32(58).string(message.institutionName);
    }
    if (message.customInstitutionName !== "") {
      writer.uint32(66).string(message.customInstitutionName);
    }
    if (message.description !== "") {
      writer.uint32(74).string(message.description);
    }
    if (message.lastManualSync !== "") {
      writer.uint32(82).string(message.lastManualSync);
    }
    if (message.lastSuccessfulUpdate !== "") {
      writer.uint32(90).string(message.lastSuccessfulUpdate);
    }
    if (message.token !== undefined) {
      Token.encode(message.token, writer.uint32(98).fork()).ldelim();
    }
    for (const v of message.bankAccounts) {
      BankAccount.encode(v!, writer.uint32(106).fork()).ldelim();
    }
    for (const v of message.investmentAccounts) {
      InvestmentAccount.encode(v!, writer.uint32(114).fork()).ldelim();
    }
    for (const v of message.creditAccounts) {
      CreditAccount.encode(v!, writer.uint32(122).fork()).ldelim();
    }
    for (const v of message.mortgageAccounts) {
      MortgageAccount.encode(v!, writer.uint32(130).fork()).ldelim();
    }
    for (const v of message.studentLoanAccounts) {
      StudentLoanAccount.encode(v!, writer.uint32(138).fork()).ldelim();
    }
    if (message.plaidInstitutionId !== "") {
      writer.uint32(146).string(message.plaidInstitutionId);
    }
    if (message.linkType !== 0) {
      writer.uint32(152).int32(message.linkType);
    }
    if (message.errorCode !== "") {
      writer.uint32(162).string(message.errorCode);
    }
    if (message.updatedAt !== "") {
      writer.uint32(170).string(message.updatedAt);
    }
    if (message.newAccountsAvailable === true) {
      writer.uint32(176).bool(message.newAccountsAvailable);
    }
    if (message.plaidSync !== undefined) {
      PlaidSync.encode(message.plaidSync, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Link {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLink();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.linkStatus = reader.int32() as any;
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.plaidLink = PlaidLink.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.plaidNewAccountsAvailable = reader.bool();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.expirationDate = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.institutionName = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.customInstitutionName = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.description = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.lastManualSync = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.lastSuccessfulUpdate = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.token = Token.decode(reader, reader.uint32());
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.bankAccounts.push(BankAccount.decode(reader, reader.uint32()));
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.investmentAccounts.push(InvestmentAccount.decode(reader, reader.uint32()));
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.creditAccounts.push(CreditAccount.decode(reader, reader.uint32()));
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.mortgageAccounts.push(MortgageAccount.decode(reader, reader.uint32()));
          continue;
        case 17:
          if (tag !== 138) {
            break;
          }

          message.studentLoanAccounts.push(StudentLoanAccount.decode(reader, reader.uint32()));
          continue;
        case 18:
          if (tag !== 146) {
            break;
          }

          message.plaidInstitutionId = reader.string();
          continue;
        case 19:
          if (tag !== 152) {
            break;
          }

          message.linkType = reader.int32() as any;
          continue;
        case 20:
          if (tag !== 162) {
            break;
          }

          message.errorCode = reader.string();
          continue;
        case 21:
          if (tag !== 170) {
            break;
          }

          message.updatedAt = reader.string();
          continue;
        case 22:
          if (tag !== 176) {
            break;
          }

          message.newAccountsAvailable = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.plaidSync = PlaidSync.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Link {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      linkStatus: isSet(object.linkStatus) ? linkStatusFromJSON(object.linkStatus) : 0,
      plaidLink: isSet(object.plaidLink) ? PlaidLink.fromJSON(object.plaidLink) : undefined,
      plaidNewAccountsAvailable: isSet(object.plaidNewAccountsAvailable)
        ? Boolean(object.plaidNewAccountsAvailable)
        : false,
      expirationDate: isSet(object.expirationDate) ? String(object.expirationDate) : "",
      institutionName: isSet(object.institutionName) ? String(object.institutionName) : "",
      customInstitutionName: isSet(object.customInstitutionName) ? String(object.customInstitutionName) : "",
      description: isSet(object.description) ? String(object.description) : "",
      lastManualSync: isSet(object.lastManualSync) ? String(object.lastManualSync) : "",
      lastSuccessfulUpdate: isSet(object.lastSuccessfulUpdate) ? String(object.lastSuccessfulUpdate) : "",
      token: isSet(object.token) ? Token.fromJSON(object.token) : undefined,
      bankAccounts: Array.isArray(object?.bankAccounts)
        ? object.bankAccounts.map((e: any) => BankAccount.fromJSON(e))
        : [],
      investmentAccounts: Array.isArray(object?.investmentAccounts)
        ? object.investmentAccounts.map((e: any) => InvestmentAccount.fromJSON(e))
        : [],
      creditAccounts: Array.isArray(object?.creditAccounts)
        ? object.creditAccounts.map((e: any) => CreditAccount.fromJSON(e))
        : [],
      mortgageAccounts: Array.isArray(object?.mortgageAccounts)
        ? object.mortgageAccounts.map((e: any) => MortgageAccount.fromJSON(e))
        : [],
      studentLoanAccounts: Array.isArray(object?.studentLoanAccounts)
        ? object.studentLoanAccounts.map((e: any) => StudentLoanAccount.fromJSON(e))
        : [],
      plaidInstitutionId: isSet(object.plaidInstitutionId) ? String(object.plaidInstitutionId) : "",
      linkType: isSet(object.linkType) ? linkTypeFromJSON(object.linkType) : 0,
      errorCode: isSet(object.errorCode) ? String(object.errorCode) : "",
      updatedAt: isSet(object.updatedAt) ? String(object.updatedAt) : "",
      newAccountsAvailable: isSet(object.newAccountsAvailable) ? Boolean(object.newAccountsAvailable) : false,
      plaidSync: isSet(object.plaidSync) ? PlaidSync.fromJSON(object.plaidSync) : undefined,
    };
  },

  toJSON(message: Link): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.linkStatus !== undefined && (obj.linkStatus = linkStatusToJSON(message.linkStatus));
    message.plaidLink !== undefined &&
      (obj.plaidLink = message.plaidLink ? PlaidLink.toJSON(message.plaidLink) : undefined);
    message.plaidNewAccountsAvailable !== undefined &&
      (obj.plaidNewAccountsAvailable = message.plaidNewAccountsAvailable);
    message.expirationDate !== undefined && (obj.expirationDate = message.expirationDate);
    message.institutionName !== undefined && (obj.institutionName = message.institutionName);
    message.customInstitutionName !== undefined && (obj.customInstitutionName = message.customInstitutionName);
    message.description !== undefined && (obj.description = message.description);
    message.lastManualSync !== undefined && (obj.lastManualSync = message.lastManualSync);
    message.lastSuccessfulUpdate !== undefined && (obj.lastSuccessfulUpdate = message.lastSuccessfulUpdate);
    message.token !== undefined && (obj.token = message.token ? Token.toJSON(message.token) : undefined);
    if (message.bankAccounts) {
      obj.bankAccounts = message.bankAccounts.map((e) => e ? BankAccount.toJSON(e) : undefined);
    } else {
      obj.bankAccounts = [];
    }
    if (message.investmentAccounts) {
      obj.investmentAccounts = message.investmentAccounts.map((e) => e ? InvestmentAccount.toJSON(e) : undefined);
    } else {
      obj.investmentAccounts = [];
    }
    if (message.creditAccounts) {
      obj.creditAccounts = message.creditAccounts.map((e) => e ? CreditAccount.toJSON(e) : undefined);
    } else {
      obj.creditAccounts = [];
    }
    if (message.mortgageAccounts) {
      obj.mortgageAccounts = message.mortgageAccounts.map((e) => e ? MortgageAccount.toJSON(e) : undefined);
    } else {
      obj.mortgageAccounts = [];
    }
    if (message.studentLoanAccounts) {
      obj.studentLoanAccounts = message.studentLoanAccounts.map((e) => e ? StudentLoanAccount.toJSON(e) : undefined);
    } else {
      obj.studentLoanAccounts = [];
    }
    message.plaidInstitutionId !== undefined && (obj.plaidInstitutionId = message.plaidInstitutionId);
    message.linkType !== undefined && (obj.linkType = linkTypeToJSON(message.linkType));
    message.errorCode !== undefined && (obj.errorCode = message.errorCode);
    message.updatedAt !== undefined && (obj.updatedAt = message.updatedAt);
    message.newAccountsAvailable !== undefined && (obj.newAccountsAvailable = message.newAccountsAvailable);
    message.plaidSync !== undefined &&
      (obj.plaidSync = message.plaidSync ? PlaidSync.toJSON(message.plaidSync) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<Link>, I>>(base?: I): Link {
    return Link.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Link>, I>>(object: I): Link {
    const message = createBaseLink();
    message.id = object.id ?? 0;
    message.linkStatus = object.linkStatus ?? 0;
    message.plaidLink = (object.plaidLink !== undefined && object.plaidLink !== null)
      ? PlaidLink.fromPartial(object.plaidLink)
      : undefined;
    message.plaidNewAccountsAvailable = object.plaidNewAccountsAvailable ?? false;
    message.expirationDate = object.expirationDate ?? "";
    message.institutionName = object.institutionName ?? "";
    message.customInstitutionName = object.customInstitutionName ?? "";
    message.description = object.description ?? "";
    message.lastManualSync = object.lastManualSync ?? "";
    message.lastSuccessfulUpdate = object.lastSuccessfulUpdate ?? "";
    message.token = (object.token !== undefined && object.token !== null) ? Token.fromPartial(object.token) : undefined;
    message.bankAccounts = object.bankAccounts?.map((e) => BankAccount.fromPartial(e)) || [];
    message.investmentAccounts = object.investmentAccounts?.map((e) => InvestmentAccount.fromPartial(e)) || [];
    message.creditAccounts = object.creditAccounts?.map((e) => CreditAccount.fromPartial(e)) || [];
    message.mortgageAccounts = object.mortgageAccounts?.map((e) => MortgageAccount.fromPartial(e)) || [];
    message.studentLoanAccounts = object.studentLoanAccounts?.map((e) => StudentLoanAccount.fromPartial(e)) || [];
    message.plaidInstitutionId = object.plaidInstitutionId ?? "";
    message.linkType = object.linkType ?? 0;
    message.errorCode = object.errorCode ?? "";
    message.updatedAt = object.updatedAt ?? "";
    message.newAccountsAvailable = object.newAccountsAvailable ?? false;
    message.plaidSync = (object.plaidSync !== undefined && object.plaidSync !== null)
      ? PlaidSync.fromPartial(object.plaidSync)
      : undefined;
    return message;
  },
};

function createBasePlaidSync(): PlaidSync {
  return { id: 0, timeStamp: "", trigger: "", nextCursor: "", added: 0, removed: 0, modified: 0 };
}

export const PlaidSync = {
  encode(message: PlaidSync, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.timeStamp !== "") {
      writer.uint32(26).string(message.timeStamp);
    }
    if (message.trigger !== "") {
      writer.uint32(34).string(message.trigger);
    }
    if (message.nextCursor !== "") {
      writer.uint32(42).string(message.nextCursor);
    }
    if (message.added !== 0) {
      writer.uint32(48).int64(message.added);
    }
    if (message.removed !== 0) {
      writer.uint32(56).int64(message.removed);
    }
    if (message.modified !== 0) {
      writer.uint32(64).int64(message.modified);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PlaidSync {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaidSync();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.timeStamp = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.trigger = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.nextCursor = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.added = longToNumber(reader.int64() as Long);
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.removed = longToNumber(reader.int64() as Long);
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.modified = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PlaidSync {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      timeStamp: isSet(object.timeStamp) ? String(object.timeStamp) : "",
      trigger: isSet(object.trigger) ? String(object.trigger) : "",
      nextCursor: isSet(object.nextCursor) ? String(object.nextCursor) : "",
      added: isSet(object.added) ? Number(object.added) : 0,
      removed: isSet(object.removed) ? Number(object.removed) : 0,
      modified: isSet(object.modified) ? Number(object.modified) : 0,
    };
  },

  toJSON(message: PlaidSync): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.timeStamp !== undefined && (obj.timeStamp = message.timeStamp);
    message.trigger !== undefined && (obj.trigger = message.trigger);
    message.nextCursor !== undefined && (obj.nextCursor = message.nextCursor);
    message.added !== undefined && (obj.added = Math.round(message.added));
    message.removed !== undefined && (obj.removed = Math.round(message.removed));
    message.modified !== undefined && (obj.modified = Math.round(message.modified));
    return obj;
  },

  create<I extends Exact<DeepPartial<PlaidSync>, I>>(base?: I): PlaidSync {
    return PlaidSync.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PlaidSync>, I>>(object: I): PlaidSync {
    const message = createBasePlaidSync();
    message.id = object.id ?? 0;
    message.timeStamp = object.timeStamp ?? "";
    message.trigger = object.trigger ?? "";
    message.nextCursor = object.nextCursor ?? "";
    message.added = object.added ?? 0;
    message.removed = object.removed ?? 0;
    message.modified = object.modified ?? 0;
    return message;
  },
};

function createBaseToken(): Token {
  return { id: 0, itemId: "", keyId: "", accessToken: "", version: "" };
}

export const Token = {
  encode(message: Token, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.itemId !== "") {
      writer.uint32(18).string(message.itemId);
    }
    if (message.keyId !== "") {
      writer.uint32(26).string(message.keyId);
    }
    if (message.accessToken !== "") {
      writer.uint32(34).string(message.accessToken);
    }
    if (message.version !== "") {
      writer.uint32(42).string(message.version);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Token {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.itemId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.keyId = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.accessToken = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.version = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Token {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      itemId: isSet(object.itemId) ? String(object.itemId) : "",
      keyId: isSet(object.keyId) ? String(object.keyId) : "",
      accessToken: isSet(object.accessToken) ? String(object.accessToken) : "",
      version: isSet(object.version) ? String(object.version) : "",
    };
  },

  toJSON(message: Token): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.itemId !== undefined && (obj.itemId = message.itemId);
    message.keyId !== undefined && (obj.keyId = message.keyId);
    message.accessToken !== undefined && (obj.accessToken = message.accessToken);
    message.version !== undefined && (obj.version = message.version);
    return obj;
  },

  create<I extends Exact<DeepPartial<Token>, I>>(base?: I): Token {
    return Token.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Token>, I>>(object: I): Token {
    const message = createBaseToken();
    message.id = object.id ?? 0;
    message.itemId = object.itemId ?? "";
    message.keyId = object.keyId ?? "";
    message.accessToken = object.accessToken ?? "";
    message.version = object.version ?? "";
    return message;
  },
};

function createBasePlaidLink(): PlaidLink {
  return {
    id: 0,
    products: [],
    webhookUrl: "",
    institutionId: "",
    institutionName: "",
    usePlaidSync: false,
    itemId: "",
  };
}

export const PlaidLink = {
  encode(message: PlaidLink, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    for (const v of message.products) {
      writer.uint32(18).string(v!);
    }
    if (message.webhookUrl !== "") {
      writer.uint32(26).string(message.webhookUrl);
    }
    if (message.institutionId !== "") {
      writer.uint32(34).string(message.institutionId);
    }
    if (message.institutionName !== "") {
      writer.uint32(42).string(message.institutionName);
    }
    if (message.usePlaidSync === true) {
      writer.uint32(48).bool(message.usePlaidSync);
    }
    if (message.itemId !== "") {
      writer.uint32(58).string(message.itemId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PlaidLink {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaidLink();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.products.push(reader.string());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.webhookUrl = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.institutionId = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.institutionName = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.usePlaidSync = reader.bool();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.itemId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PlaidLink {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      products: Array.isArray(object?.products) ? object.products.map((e: any) => String(e)) : [],
      webhookUrl: isSet(object.webhookUrl) ? String(object.webhookUrl) : "",
      institutionId: isSet(object.institutionId) ? String(object.institutionId) : "",
      institutionName: isSet(object.institutionName) ? String(object.institutionName) : "",
      usePlaidSync: isSet(object.usePlaidSync) ? Boolean(object.usePlaidSync) : false,
      itemId: isSet(object.itemId) ? String(object.itemId) : "",
    };
  },

  toJSON(message: PlaidLink): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    if (message.products) {
      obj.products = message.products.map((e) => e);
    } else {
      obj.products = [];
    }
    message.webhookUrl !== undefined && (obj.webhookUrl = message.webhookUrl);
    message.institutionId !== undefined && (obj.institutionId = message.institutionId);
    message.institutionName !== undefined && (obj.institutionName = message.institutionName);
    message.usePlaidSync !== undefined && (obj.usePlaidSync = message.usePlaidSync);
    message.itemId !== undefined && (obj.itemId = message.itemId);
    return obj;
  },

  create<I extends Exact<DeepPartial<PlaidLink>, I>>(base?: I): PlaidLink {
    return PlaidLink.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PlaidLink>, I>>(object: I): PlaidLink {
    const message = createBasePlaidLink();
    message.id = object.id ?? 0;
    message.products = object.products?.map((e) => e) || [];
    message.webhookUrl = object.webhookUrl ?? "";
    message.institutionId = object.institutionId ?? "";
    message.institutionName = object.institutionName ?? "";
    message.usePlaidSync = object.usePlaidSync ?? false;
    message.itemId = object.itemId ?? "";
    return message;
  },
};

function createBaseStudentLoanAccount(): StudentLoanAccount {
  return {
    id: 0,
    plaidAccountId: "",
    disbursementDates: [],
    expectedPayoffDate: "",
    guarantor: "",
    interestRatePercentage: 0,
    isOverdue: false,
    lastPaymentAmount: 0,
    lastPaymentDate: "",
    lastStatementIssueDate: "",
    loanName: "",
    loanEndDate: "",
    minimumPaymentAmount: 0,
    nextPaymentDueDate: "",
    originationDate: "",
    originationPrincipalAmount: 0,
    outstandingInterestAmount: 0,
    paymentReferenceNumber: "",
    sequenceNumber: "",
    ytdInterestPaid: 0,
    ytdPrincipalPaid: 0,
    loanType: "",
    pslfStatusEstimatedEligibilityDate: "",
    pslfStatusPaymentsMade: 0,
    pslfStatusPaymentsRemaining: 0,
    repaymentPlanType: "",
    repaymentPlanDescription: "",
    servicerAddressCity: "",
    servicerAddressPostalCode: "",
    servicerAddressState: "",
    servicerAddressStreet: "",
    servicerAddressRegion: "",
    servicerAddressCountry: "",
    userId: 0,
    name: "",
  };
}

export const StudentLoanAccount = {
  encode(message: StudentLoanAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.plaidAccountId !== "") {
      writer.uint32(18).string(message.plaidAccountId);
    }
    for (const v of message.disbursementDates) {
      writer.uint32(26).string(v!);
    }
    if (message.expectedPayoffDate !== "") {
      writer.uint32(34).string(message.expectedPayoffDate);
    }
    if (message.guarantor !== "") {
      writer.uint32(42).string(message.guarantor);
    }
    if (message.interestRatePercentage !== 0) {
      writer.uint32(49).double(message.interestRatePercentage);
    }
    if (message.isOverdue === true) {
      writer.uint32(56).bool(message.isOverdue);
    }
    if (message.lastPaymentAmount !== 0) {
      writer.uint32(65).double(message.lastPaymentAmount);
    }
    if (message.lastPaymentDate !== "") {
      writer.uint32(74).string(message.lastPaymentDate);
    }
    if (message.lastStatementIssueDate !== "") {
      writer.uint32(82).string(message.lastStatementIssueDate);
    }
    if (message.loanName !== "") {
      writer.uint32(90).string(message.loanName);
    }
    if (message.loanEndDate !== "") {
      writer.uint32(98).string(message.loanEndDate);
    }
    if (message.minimumPaymentAmount !== 0) {
      writer.uint32(105).double(message.minimumPaymentAmount);
    }
    if (message.nextPaymentDueDate !== "") {
      writer.uint32(114).string(message.nextPaymentDueDate);
    }
    if (message.originationDate !== "") {
      writer.uint32(122).string(message.originationDate);
    }
    if (message.originationPrincipalAmount !== 0) {
      writer.uint32(129).double(message.originationPrincipalAmount);
    }
    if (message.outstandingInterestAmount !== 0) {
      writer.uint32(137).double(message.outstandingInterestAmount);
    }
    if (message.paymentReferenceNumber !== "") {
      writer.uint32(146).string(message.paymentReferenceNumber);
    }
    if (message.sequenceNumber !== "") {
      writer.uint32(170).string(message.sequenceNumber);
    }
    if (message.ytdInterestPaid !== 0) {
      writer.uint32(185).double(message.ytdInterestPaid);
    }
    if (message.ytdPrincipalPaid !== 0) {
      writer.uint32(193).double(message.ytdPrincipalPaid);
    }
    if (message.loanType !== "") {
      writer.uint32(202).string(message.loanType);
    }
    if (message.pslfStatusEstimatedEligibilityDate !== "") {
      writer.uint32(210).string(message.pslfStatusEstimatedEligibilityDate);
    }
    if (message.pslfStatusPaymentsMade !== 0) {
      writer.uint32(216).int32(message.pslfStatusPaymentsMade);
    }
    if (message.pslfStatusPaymentsRemaining !== 0) {
      writer.uint32(224).int32(message.pslfStatusPaymentsRemaining);
    }
    if (message.repaymentPlanType !== "") {
      writer.uint32(234).string(message.repaymentPlanType);
    }
    if (message.repaymentPlanDescription !== "") {
      writer.uint32(242).string(message.repaymentPlanDescription);
    }
    if (message.servicerAddressCity !== "") {
      writer.uint32(250).string(message.servicerAddressCity);
    }
    if (message.servicerAddressPostalCode !== "") {
      writer.uint32(258).string(message.servicerAddressPostalCode);
    }
    if (message.servicerAddressState !== "") {
      writer.uint32(266).string(message.servicerAddressState);
    }
    if (message.servicerAddressStreet !== "") {
      writer.uint32(274).string(message.servicerAddressStreet);
    }
    if (message.servicerAddressRegion !== "") {
      writer.uint32(282).string(message.servicerAddressRegion);
    }
    if (message.servicerAddressCountry !== "") {
      writer.uint32(290).string(message.servicerAddressCountry);
    }
    if (message.userId !== 0) {
      writer.uint32(296).uint64(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(306).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StudentLoanAccount {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStudentLoanAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.disbursementDates.push(reader.string());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.expectedPayoffDate = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.guarantor = reader.string();
          continue;
        case 6:
          if (tag !== 49) {
            break;
          }

          message.interestRatePercentage = reader.double();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.isOverdue = reader.bool();
          continue;
        case 8:
          if (tag !== 65) {
            break;
          }

          message.lastPaymentAmount = reader.double();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.lastPaymentDate = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.lastStatementIssueDate = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.loanName = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.loanEndDate = reader.string();
          continue;
        case 13:
          if (tag !== 105) {
            break;
          }

          message.minimumPaymentAmount = reader.double();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.nextPaymentDueDate = reader.string();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.originationDate = reader.string();
          continue;
        case 16:
          if (tag !== 129) {
            break;
          }

          message.originationPrincipalAmount = reader.double();
          continue;
        case 17:
          if (tag !== 137) {
            break;
          }

          message.outstandingInterestAmount = reader.double();
          continue;
        case 18:
          if (tag !== 146) {
            break;
          }

          message.paymentReferenceNumber = reader.string();
          continue;
        case 21:
          if (tag !== 170) {
            break;
          }

          message.sequenceNumber = reader.string();
          continue;
        case 23:
          if (tag !== 185) {
            break;
          }

          message.ytdInterestPaid = reader.double();
          continue;
        case 24:
          if (tag !== 193) {
            break;
          }

          message.ytdPrincipalPaid = reader.double();
          continue;
        case 25:
          if (tag !== 202) {
            break;
          }

          message.loanType = reader.string();
          continue;
        case 26:
          if (tag !== 210) {
            break;
          }

          message.pslfStatusEstimatedEligibilityDate = reader.string();
          continue;
        case 27:
          if (tag !== 216) {
            break;
          }

          message.pslfStatusPaymentsMade = reader.int32();
          continue;
        case 28:
          if (tag !== 224) {
            break;
          }

          message.pslfStatusPaymentsRemaining = reader.int32();
          continue;
        case 29:
          if (tag !== 234) {
            break;
          }

          message.repaymentPlanType = reader.string();
          continue;
        case 30:
          if (tag !== 242) {
            break;
          }

          message.repaymentPlanDescription = reader.string();
          continue;
        case 31:
          if (tag !== 250) {
            break;
          }

          message.servicerAddressCity = reader.string();
          continue;
        case 32:
          if (tag !== 258) {
            break;
          }

          message.servicerAddressPostalCode = reader.string();
          continue;
        case 33:
          if (tag !== 266) {
            break;
          }

          message.servicerAddressState = reader.string();
          continue;
        case 34:
          if (tag !== 274) {
            break;
          }

          message.servicerAddressStreet = reader.string();
          continue;
        case 35:
          if (tag !== 282) {
            break;
          }

          message.servicerAddressRegion = reader.string();
          continue;
        case 36:
          if (tag !== 290) {
            break;
          }

          message.servicerAddressCountry = reader.string();
          continue;
        case 37:
          if (tag !== 296) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 38:
          if (tag !== 306) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): StudentLoanAccount {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      disbursementDates: Array.isArray(object?.disbursementDates)
        ? object.disbursementDates.map((e: any) => String(e))
        : [],
      expectedPayoffDate: isSet(object.expectedPayoffDate) ? String(object.expectedPayoffDate) : "",
      guarantor: isSet(object.guarantor) ? String(object.guarantor) : "",
      interestRatePercentage: isSet(object.interestRatePercentage) ? Number(object.interestRatePercentage) : 0,
      isOverdue: isSet(object.isOverdue) ? Boolean(object.isOverdue) : false,
      lastPaymentAmount: isSet(object.lastPaymentAmount) ? Number(object.lastPaymentAmount) : 0,
      lastPaymentDate: isSet(object.lastPaymentDate) ? String(object.lastPaymentDate) : "",
      lastStatementIssueDate: isSet(object.lastStatementIssueDate) ? String(object.lastStatementIssueDate) : "",
      loanName: isSet(object.loanName) ? String(object.loanName) : "",
      loanEndDate: isSet(object.loanEndDate) ? String(object.loanEndDate) : "",
      minimumPaymentAmount: isSet(object.minimumPaymentAmount) ? Number(object.minimumPaymentAmount) : 0,
      nextPaymentDueDate: isSet(object.nextPaymentDueDate) ? String(object.nextPaymentDueDate) : "",
      originationDate: isSet(object.originationDate) ? String(object.originationDate) : "",
      originationPrincipalAmount: isSet(object.originationPrincipalAmount)
        ? Number(object.originationPrincipalAmount)
        : 0,
      outstandingInterestAmount: isSet(object.outstandingInterestAmount) ? Number(object.outstandingInterestAmount) : 0,
      paymentReferenceNumber: isSet(object.paymentReferenceNumber) ? String(object.paymentReferenceNumber) : "",
      sequenceNumber: isSet(object.sequenceNumber) ? String(object.sequenceNumber) : "",
      ytdInterestPaid: isSet(object.ytdInterestPaid) ? Number(object.ytdInterestPaid) : 0,
      ytdPrincipalPaid: isSet(object.ytdPrincipalPaid) ? Number(object.ytdPrincipalPaid) : 0,
      loanType: isSet(object.loanType) ? String(object.loanType) : "",
      pslfStatusEstimatedEligibilityDate: isSet(object.pslfStatusEstimatedEligibilityDate)
        ? String(object.pslfStatusEstimatedEligibilityDate)
        : "",
      pslfStatusPaymentsMade: isSet(object.pslfStatusPaymentsMade) ? Number(object.pslfStatusPaymentsMade) : 0,
      pslfStatusPaymentsRemaining: isSet(object.pslfStatusPaymentsRemaining)
        ? Number(object.pslfStatusPaymentsRemaining)
        : 0,
      repaymentPlanType: isSet(object.repaymentPlanType) ? String(object.repaymentPlanType) : "",
      repaymentPlanDescription: isSet(object.repaymentPlanDescription) ? String(object.repaymentPlanDescription) : "",
      servicerAddressCity: isSet(object.servicerAddressCity) ? String(object.servicerAddressCity) : "",
      servicerAddressPostalCode: isSet(object.servicerAddressPostalCode)
        ? String(object.servicerAddressPostalCode)
        : "",
      servicerAddressState: isSet(object.servicerAddressState) ? String(object.servicerAddressState) : "",
      servicerAddressStreet: isSet(object.servicerAddressStreet) ? String(object.servicerAddressStreet) : "",
      servicerAddressRegion: isSet(object.servicerAddressRegion) ? String(object.servicerAddressRegion) : "",
      servicerAddressCountry: isSet(object.servicerAddressCountry) ? String(object.servicerAddressCountry) : "",
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: StudentLoanAccount): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.plaidAccountId !== undefined && (obj.plaidAccountId = message.plaidAccountId);
    if (message.disbursementDates) {
      obj.disbursementDates = message.disbursementDates.map((e) => e);
    } else {
      obj.disbursementDates = [];
    }
    message.expectedPayoffDate !== undefined && (obj.expectedPayoffDate = message.expectedPayoffDate);
    message.guarantor !== undefined && (obj.guarantor = message.guarantor);
    message.interestRatePercentage !== undefined && (obj.interestRatePercentage = message.interestRatePercentage);
    message.isOverdue !== undefined && (obj.isOverdue = message.isOverdue);
    message.lastPaymentAmount !== undefined && (obj.lastPaymentAmount = message.lastPaymentAmount);
    message.lastPaymentDate !== undefined && (obj.lastPaymentDate = message.lastPaymentDate);
    message.lastStatementIssueDate !== undefined && (obj.lastStatementIssueDate = message.lastStatementIssueDate);
    message.loanName !== undefined && (obj.loanName = message.loanName);
    message.loanEndDate !== undefined && (obj.loanEndDate = message.loanEndDate);
    message.minimumPaymentAmount !== undefined && (obj.minimumPaymentAmount = message.minimumPaymentAmount);
    message.nextPaymentDueDate !== undefined && (obj.nextPaymentDueDate = message.nextPaymentDueDate);
    message.originationDate !== undefined && (obj.originationDate = message.originationDate);
    message.originationPrincipalAmount !== undefined &&
      (obj.originationPrincipalAmount = message.originationPrincipalAmount);
    message.outstandingInterestAmount !== undefined &&
      (obj.outstandingInterestAmount = message.outstandingInterestAmount);
    message.paymentReferenceNumber !== undefined && (obj.paymentReferenceNumber = message.paymentReferenceNumber);
    message.sequenceNumber !== undefined && (obj.sequenceNumber = message.sequenceNumber);
    message.ytdInterestPaid !== undefined && (obj.ytdInterestPaid = message.ytdInterestPaid);
    message.ytdPrincipalPaid !== undefined && (obj.ytdPrincipalPaid = message.ytdPrincipalPaid);
    message.loanType !== undefined && (obj.loanType = message.loanType);
    message.pslfStatusEstimatedEligibilityDate !== undefined &&
      (obj.pslfStatusEstimatedEligibilityDate = message.pslfStatusEstimatedEligibilityDate);
    message.pslfStatusPaymentsMade !== undefined &&
      (obj.pslfStatusPaymentsMade = Math.round(message.pslfStatusPaymentsMade));
    message.pslfStatusPaymentsRemaining !== undefined &&
      (obj.pslfStatusPaymentsRemaining = Math.round(message.pslfStatusPaymentsRemaining));
    message.repaymentPlanType !== undefined && (obj.repaymentPlanType = message.repaymentPlanType);
    message.repaymentPlanDescription !== undefined && (obj.repaymentPlanDescription = message.repaymentPlanDescription);
    message.servicerAddressCity !== undefined && (obj.servicerAddressCity = message.servicerAddressCity);
    message.servicerAddressPostalCode !== undefined &&
      (obj.servicerAddressPostalCode = message.servicerAddressPostalCode);
    message.servicerAddressState !== undefined && (obj.servicerAddressState = message.servicerAddressState);
    message.servicerAddressStreet !== undefined && (obj.servicerAddressStreet = message.servicerAddressStreet);
    message.servicerAddressRegion !== undefined && (obj.servicerAddressRegion = message.servicerAddressRegion);
    message.servicerAddressCountry !== undefined && (obj.servicerAddressCountry = message.servicerAddressCountry);
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<StudentLoanAccount>, I>>(base?: I): StudentLoanAccount {
    return StudentLoanAccount.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StudentLoanAccount>, I>>(object: I): StudentLoanAccount {
    const message = createBaseStudentLoanAccount();
    message.id = object.id ?? 0;
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.disbursementDates = object.disbursementDates?.map((e) => e) || [];
    message.expectedPayoffDate = object.expectedPayoffDate ?? "";
    message.guarantor = object.guarantor ?? "";
    message.interestRatePercentage = object.interestRatePercentage ?? 0;
    message.isOverdue = object.isOverdue ?? false;
    message.lastPaymentAmount = object.lastPaymentAmount ?? 0;
    message.lastPaymentDate = object.lastPaymentDate ?? "";
    message.lastStatementIssueDate = object.lastStatementIssueDate ?? "";
    message.loanName = object.loanName ?? "";
    message.loanEndDate = object.loanEndDate ?? "";
    message.minimumPaymentAmount = object.minimumPaymentAmount ?? 0;
    message.nextPaymentDueDate = object.nextPaymentDueDate ?? "";
    message.originationDate = object.originationDate ?? "";
    message.originationPrincipalAmount = object.originationPrincipalAmount ?? 0;
    message.outstandingInterestAmount = object.outstandingInterestAmount ?? 0;
    message.paymentReferenceNumber = object.paymentReferenceNumber ?? "";
    message.sequenceNumber = object.sequenceNumber ?? "";
    message.ytdInterestPaid = object.ytdInterestPaid ?? 0;
    message.ytdPrincipalPaid = object.ytdPrincipalPaid ?? 0;
    message.loanType = object.loanType ?? "";
    message.pslfStatusEstimatedEligibilityDate = object.pslfStatusEstimatedEligibilityDate ?? "";
    message.pslfStatusPaymentsMade = object.pslfStatusPaymentsMade ?? 0;
    message.pslfStatusPaymentsRemaining = object.pslfStatusPaymentsRemaining ?? 0;
    message.repaymentPlanType = object.repaymentPlanType ?? "";
    message.repaymentPlanDescription = object.repaymentPlanDescription ?? "";
    message.servicerAddressCity = object.servicerAddressCity ?? "";
    message.servicerAddressPostalCode = object.servicerAddressPostalCode ?? "";
    message.servicerAddressState = object.servicerAddressState ?? "";
    message.servicerAddressStreet = object.servicerAddressStreet ?? "";
    message.servicerAddressRegion = object.servicerAddressRegion ?? "";
    message.servicerAddressCountry = object.servicerAddressCountry ?? "";
    message.userId = object.userId ?? 0;
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseCreditAccount(): CreditAccount {
  return {
    id: 0,
    userId: 0,
    name: "",
    number: "",
    type: "",
    balance: 0,
    currentFunds: 0,
    balanceLimit: 0,
    plaidAccountId: "",
    subtype: "",
    isOverdue: false,
    lastPaymentAmount: 0,
    lastPaymentDate: "",
    lastStatementIssueDate: "",
    minimumAmountDueDate: 0,
    nextPaymentDate: "",
    aprs: [],
    lastStatementBalance: 0,
    minimumPaymentAmount: 0,
    nextPaymentDueDate: "",
  };
}

export const CreditAccount = {
  encode(message: CreditAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.number !== "") {
      writer.uint32(34).string(message.number);
    }
    if (message.type !== "") {
      writer.uint32(42).string(message.type);
    }
    if (message.balance !== 0) {
      writer.uint32(53).float(message.balance);
    }
    if (message.currentFunds !== 0) {
      writer.uint32(73).double(message.currentFunds);
    }
    if (message.balanceLimit !== 0) {
      writer.uint32(80).uint64(message.balanceLimit);
    }
    if (message.plaidAccountId !== "") {
      writer.uint32(98).string(message.plaidAccountId);
    }
    if (message.subtype !== "") {
      writer.uint32(106).string(message.subtype);
    }
    if (message.isOverdue === true) {
      writer.uint32(112).bool(message.isOverdue);
    }
    if (message.lastPaymentAmount !== 0) {
      writer.uint32(121).double(message.lastPaymentAmount);
    }
    if (message.lastPaymentDate !== "") {
      writer.uint32(130).string(message.lastPaymentDate);
    }
    if (message.lastStatementIssueDate !== "") {
      writer.uint32(138).string(message.lastStatementIssueDate);
    }
    if (message.minimumAmountDueDate !== 0) {
      writer.uint32(145).double(message.minimumAmountDueDate);
    }
    if (message.nextPaymentDate !== "") {
      writer.uint32(154).string(message.nextPaymentDate);
    }
    for (const v of message.aprs) {
      Apr.encode(v!, writer.uint32(162).fork()).ldelim();
    }
    if (message.lastStatementBalance !== 0) {
      writer.uint32(169).double(message.lastStatementBalance);
    }
    if (message.minimumPaymentAmount !== 0) {
      writer.uint32(177).double(message.minimumPaymentAmount);
    }
    if (message.nextPaymentDueDate !== "") {
      writer.uint32(186).string(message.nextPaymentDueDate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreditAccount {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreditAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.number = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.type = reader.string();
          continue;
        case 6:
          if (tag !== 53) {
            break;
          }

          message.balance = reader.float();
          continue;
        case 9:
          if (tag !== 73) {
            break;
          }

          message.currentFunds = reader.double();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.balanceLimit = longToNumber(reader.uint64() as Long);
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.subtype = reader.string();
          continue;
        case 14:
          if (tag !== 112) {
            break;
          }

          message.isOverdue = reader.bool();
          continue;
        case 15:
          if (tag !== 121) {
            break;
          }

          message.lastPaymentAmount = reader.double();
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.lastPaymentDate = reader.string();
          continue;
        case 17:
          if (tag !== 138) {
            break;
          }

          message.lastStatementIssueDate = reader.string();
          continue;
        case 18:
          if (tag !== 145) {
            break;
          }

          message.minimumAmountDueDate = reader.double();
          continue;
        case 19:
          if (tag !== 154) {
            break;
          }

          message.nextPaymentDate = reader.string();
          continue;
        case 20:
          if (tag !== 162) {
            break;
          }

          message.aprs.push(Apr.decode(reader, reader.uint32()));
          continue;
        case 21:
          if (tag !== 169) {
            break;
          }

          message.lastStatementBalance = reader.double();
          continue;
        case 22:
          if (tag !== 177) {
            break;
          }

          message.minimumPaymentAmount = reader.double();
          continue;
        case 23:
          if (tag !== 186) {
            break;
          }

          message.nextPaymentDueDate = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreditAccount {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      number: isSet(object.number) ? String(object.number) : "",
      type: isSet(object.type) ? String(object.type) : "",
      balance: isSet(object.balance) ? Number(object.balance) : 0,
      currentFunds: isSet(object.currentFunds) ? Number(object.currentFunds) : 0,
      balanceLimit: isSet(object.balanceLimit) ? Number(object.balanceLimit) : 0,
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      subtype: isSet(object.subtype) ? String(object.subtype) : "",
      isOverdue: isSet(object.isOverdue) ? Boolean(object.isOverdue) : false,
      lastPaymentAmount: isSet(object.lastPaymentAmount) ? Number(object.lastPaymentAmount) : 0,
      lastPaymentDate: isSet(object.lastPaymentDate) ? String(object.lastPaymentDate) : "",
      lastStatementIssueDate: isSet(object.lastStatementIssueDate) ? String(object.lastStatementIssueDate) : "",
      minimumAmountDueDate: isSet(object.minimumAmountDueDate) ? Number(object.minimumAmountDueDate) : 0,
      nextPaymentDate: isSet(object.nextPaymentDate) ? String(object.nextPaymentDate) : "",
      aprs: Array.isArray(object?.aprs) ? object.aprs.map((e: any) => Apr.fromJSON(e)) : [],
      lastStatementBalance: isSet(object.lastStatementBalance) ? Number(object.lastStatementBalance) : 0,
      minimumPaymentAmount: isSet(object.minimumPaymentAmount) ? Number(object.minimumPaymentAmount) : 0,
      nextPaymentDueDate: isSet(object.nextPaymentDueDate) ? String(object.nextPaymentDueDate) : "",
    };
  },

  toJSON(message: CreditAccount): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.name !== undefined && (obj.name = message.name);
    message.number !== undefined && (obj.number = message.number);
    message.type !== undefined && (obj.type = message.type);
    message.balance !== undefined && (obj.balance = message.balance);
    message.currentFunds !== undefined && (obj.currentFunds = message.currentFunds);
    message.balanceLimit !== undefined && (obj.balanceLimit = Math.round(message.balanceLimit));
    message.plaidAccountId !== undefined && (obj.plaidAccountId = message.plaidAccountId);
    message.subtype !== undefined && (obj.subtype = message.subtype);
    message.isOverdue !== undefined && (obj.isOverdue = message.isOverdue);
    message.lastPaymentAmount !== undefined && (obj.lastPaymentAmount = message.lastPaymentAmount);
    message.lastPaymentDate !== undefined && (obj.lastPaymentDate = message.lastPaymentDate);
    message.lastStatementIssueDate !== undefined && (obj.lastStatementIssueDate = message.lastStatementIssueDate);
    message.minimumAmountDueDate !== undefined && (obj.minimumAmountDueDate = message.minimumAmountDueDate);
    message.nextPaymentDate !== undefined && (obj.nextPaymentDate = message.nextPaymentDate);
    if (message.aprs) {
      obj.aprs = message.aprs.map((e) => e ? Apr.toJSON(e) : undefined);
    } else {
      obj.aprs = [];
    }
    message.lastStatementBalance !== undefined && (obj.lastStatementBalance = message.lastStatementBalance);
    message.minimumPaymentAmount !== undefined && (obj.minimumPaymentAmount = message.minimumPaymentAmount);
    message.nextPaymentDueDate !== undefined && (obj.nextPaymentDueDate = message.nextPaymentDueDate);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreditAccount>, I>>(base?: I): CreditAccount {
    return CreditAccount.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreditAccount>, I>>(object: I): CreditAccount {
    const message = createBaseCreditAccount();
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.name = object.name ?? "";
    message.number = object.number ?? "";
    message.type = object.type ?? "";
    message.balance = object.balance ?? 0;
    message.currentFunds = object.currentFunds ?? 0;
    message.balanceLimit = object.balanceLimit ?? 0;
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.subtype = object.subtype ?? "";
    message.isOverdue = object.isOverdue ?? false;
    message.lastPaymentAmount = object.lastPaymentAmount ?? 0;
    message.lastPaymentDate = object.lastPaymentDate ?? "";
    message.lastStatementIssueDate = object.lastStatementIssueDate ?? "";
    message.minimumAmountDueDate = object.minimumAmountDueDate ?? 0;
    message.nextPaymentDate = object.nextPaymentDate ?? "";
    message.aprs = object.aprs?.map((e) => Apr.fromPartial(e)) || [];
    message.lastStatementBalance = object.lastStatementBalance ?? 0;
    message.minimumPaymentAmount = object.minimumPaymentAmount ?? 0;
    message.nextPaymentDueDate = object.nextPaymentDueDate ?? "";
    return message;
  },
};

function createBaseMortgageAccount(): MortgageAccount {
  return {
    id: 0,
    plaidAccountId: "",
    accountNumber: "",
    currentLateFee: 0,
    escrowBalance: 0,
    hasPmi: false,
    hasPrepaymentPenalty: false,
    lastPaymentAmount: 0,
    lastPaymentDate: "",
    loanTerm: "",
    loanTypeDescription: "",
    maturityDate: "",
    nextMonthlyPayment: 0,
    nextPaymentDueDate: "",
    originalPrincipalBalance: 0,
    originalPropertyValue: 0,
    outstandingPrincipalBalance: 0,
    paymentAmount: 0,
    paymentDate: "",
    originationDate: "",
    originationPrincipalAmount: 0,
    pastDueAmount: 0,
    ytdInterestPaid: 0,
    ytdPrincipalPaid: 0,
    propertyAddressCity: "",
    propertyAddressState: "",
    propertyAddressStreet: "",
    propertyAddressPostalCode: "",
    propertyRegion: "",
    propertyCountry: "",
    interestRatePercentage: 0,
    interestRateType: "",
  };
}

export const MortgageAccount = {
  encode(message: MortgageAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.plaidAccountId !== "") {
      writer.uint32(18).string(message.plaidAccountId);
    }
    if (message.accountNumber !== "") {
      writer.uint32(26).string(message.accountNumber);
    }
    if (message.currentLateFee !== 0) {
      writer.uint32(33).double(message.currentLateFee);
    }
    if (message.escrowBalance !== 0) {
      writer.uint32(41).double(message.escrowBalance);
    }
    if (message.hasPmi === true) {
      writer.uint32(48).bool(message.hasPmi);
    }
    if (message.hasPrepaymentPenalty === true) {
      writer.uint32(56).bool(message.hasPrepaymentPenalty);
    }
    if (message.lastPaymentAmount !== 0) {
      writer.uint32(73).double(message.lastPaymentAmount);
    }
    if (message.lastPaymentDate !== "") {
      writer.uint32(82).string(message.lastPaymentDate);
    }
    if (message.loanTerm !== "") {
      writer.uint32(90).string(message.loanTerm);
    }
    if (message.loanTypeDescription !== "") {
      writer.uint32(98).string(message.loanTypeDescription);
    }
    if (message.maturityDate !== "") {
      writer.uint32(106).string(message.maturityDate);
    }
    if (message.nextMonthlyPayment !== 0) {
      writer.uint32(113).double(message.nextMonthlyPayment);
    }
    if (message.nextPaymentDueDate !== "") {
      writer.uint32(122).string(message.nextPaymentDueDate);
    }
    if (message.originalPrincipalBalance !== 0) {
      writer.uint32(129).double(message.originalPrincipalBalance);
    }
    if (message.originalPropertyValue !== 0) {
      writer.uint32(137).double(message.originalPropertyValue);
    }
    if (message.outstandingPrincipalBalance !== 0) {
      writer.uint32(145).double(message.outstandingPrincipalBalance);
    }
    if (message.paymentAmount !== 0) {
      writer.uint32(153).double(message.paymentAmount);
    }
    if (message.paymentDate !== "") {
      writer.uint32(162).string(message.paymentDate);
    }
    if (message.originationDate !== "") {
      writer.uint32(202).string(message.originationDate);
    }
    if (message.originationPrincipalAmount !== 0) {
      writer.uint32(209).double(message.originationPrincipalAmount);
    }
    if (message.pastDueAmount !== 0) {
      writer.uint32(225).double(message.pastDueAmount);
    }
    if (message.ytdInterestPaid !== 0) {
      writer.uint32(233).double(message.ytdInterestPaid);
    }
    if (message.ytdPrincipalPaid !== 0) {
      writer.uint32(241).double(message.ytdPrincipalPaid);
    }
    if (message.propertyAddressCity !== "") {
      writer.uint32(250).string(message.propertyAddressCity);
    }
    if (message.propertyAddressState !== "") {
      writer.uint32(258).string(message.propertyAddressState);
    }
    if (message.propertyAddressStreet !== "") {
      writer.uint32(266).string(message.propertyAddressStreet);
    }
    if (message.propertyAddressPostalCode !== "") {
      writer.uint32(274).string(message.propertyAddressPostalCode);
    }
    if (message.propertyRegion !== "") {
      writer.uint32(282).string(message.propertyRegion);
    }
    if (message.propertyCountry !== "") {
      writer.uint32(290).string(message.propertyCountry);
    }
    if (message.interestRatePercentage !== 0) {
      writer.uint32(297).double(message.interestRatePercentage);
    }
    if (message.interestRateType !== "") {
      writer.uint32(306).string(message.interestRateType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MortgageAccount {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMortgageAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.accountNumber = reader.string();
          continue;
        case 4:
          if (tag !== 33) {
            break;
          }

          message.currentLateFee = reader.double();
          continue;
        case 5:
          if (tag !== 41) {
            break;
          }

          message.escrowBalance = reader.double();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.hasPmi = reader.bool();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.hasPrepaymentPenalty = reader.bool();
          continue;
        case 9:
          if (tag !== 73) {
            break;
          }

          message.lastPaymentAmount = reader.double();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.lastPaymentDate = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.loanTerm = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.loanTypeDescription = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.maturityDate = reader.string();
          continue;
        case 14:
          if (tag !== 113) {
            break;
          }

          message.nextMonthlyPayment = reader.double();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.nextPaymentDueDate = reader.string();
          continue;
        case 16:
          if (tag !== 129) {
            break;
          }

          message.originalPrincipalBalance = reader.double();
          continue;
        case 17:
          if (tag !== 137) {
            break;
          }

          message.originalPropertyValue = reader.double();
          continue;
        case 18:
          if (tag !== 145) {
            break;
          }

          message.outstandingPrincipalBalance = reader.double();
          continue;
        case 19:
          if (tag !== 153) {
            break;
          }

          message.paymentAmount = reader.double();
          continue;
        case 20:
          if (tag !== 162) {
            break;
          }

          message.paymentDate = reader.string();
          continue;
        case 25:
          if (tag !== 202) {
            break;
          }

          message.originationDate = reader.string();
          continue;
        case 26:
          if (tag !== 209) {
            break;
          }

          message.originationPrincipalAmount = reader.double();
          continue;
        case 28:
          if (tag !== 225) {
            break;
          }

          message.pastDueAmount = reader.double();
          continue;
        case 29:
          if (tag !== 233) {
            break;
          }

          message.ytdInterestPaid = reader.double();
          continue;
        case 30:
          if (tag !== 241) {
            break;
          }

          message.ytdPrincipalPaid = reader.double();
          continue;
        case 31:
          if (tag !== 250) {
            break;
          }

          message.propertyAddressCity = reader.string();
          continue;
        case 32:
          if (tag !== 258) {
            break;
          }

          message.propertyAddressState = reader.string();
          continue;
        case 33:
          if (tag !== 266) {
            break;
          }

          message.propertyAddressStreet = reader.string();
          continue;
        case 34:
          if (tag !== 274) {
            break;
          }

          message.propertyAddressPostalCode = reader.string();
          continue;
        case 35:
          if (tag !== 282) {
            break;
          }

          message.propertyRegion = reader.string();
          continue;
        case 36:
          if (tag !== 290) {
            break;
          }

          message.propertyCountry = reader.string();
          continue;
        case 37:
          if (tag !== 297) {
            break;
          }

          message.interestRatePercentage = reader.double();
          continue;
        case 38:
          if (tag !== 306) {
            break;
          }

          message.interestRateType = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MortgageAccount {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      accountNumber: isSet(object.accountNumber) ? String(object.accountNumber) : "",
      currentLateFee: isSet(object.currentLateFee) ? Number(object.currentLateFee) : 0,
      escrowBalance: isSet(object.escrowBalance) ? Number(object.escrowBalance) : 0,
      hasPmi: isSet(object.hasPmi) ? Boolean(object.hasPmi) : false,
      hasPrepaymentPenalty: isSet(object.hasPrepaymentPenalty) ? Boolean(object.hasPrepaymentPenalty) : false,
      lastPaymentAmount: isSet(object.lastPaymentAmount) ? Number(object.lastPaymentAmount) : 0,
      lastPaymentDate: isSet(object.lastPaymentDate) ? String(object.lastPaymentDate) : "",
      loanTerm: isSet(object.loanTerm) ? String(object.loanTerm) : "",
      loanTypeDescription: isSet(object.loanTypeDescription) ? String(object.loanTypeDescription) : "",
      maturityDate: isSet(object.maturityDate) ? String(object.maturityDate) : "",
      nextMonthlyPayment: isSet(object.nextMonthlyPayment) ? Number(object.nextMonthlyPayment) : 0,
      nextPaymentDueDate: isSet(object.nextPaymentDueDate) ? String(object.nextPaymentDueDate) : "",
      originalPrincipalBalance: isSet(object.originalPrincipalBalance) ? Number(object.originalPrincipalBalance) : 0,
      originalPropertyValue: isSet(object.originalPropertyValue) ? Number(object.originalPropertyValue) : 0,
      outstandingPrincipalBalance: isSet(object.outstandingPrincipalBalance)
        ? Number(object.outstandingPrincipalBalance)
        : 0,
      paymentAmount: isSet(object.paymentAmount) ? Number(object.paymentAmount) : 0,
      paymentDate: isSet(object.paymentDate) ? String(object.paymentDate) : "",
      originationDate: isSet(object.originationDate) ? String(object.originationDate) : "",
      originationPrincipalAmount: isSet(object.originationPrincipalAmount)
        ? Number(object.originationPrincipalAmount)
        : 0,
      pastDueAmount: isSet(object.pastDueAmount) ? Number(object.pastDueAmount) : 0,
      ytdInterestPaid: isSet(object.ytdInterestPaid) ? Number(object.ytdInterestPaid) : 0,
      ytdPrincipalPaid: isSet(object.ytdPrincipalPaid) ? Number(object.ytdPrincipalPaid) : 0,
      propertyAddressCity: isSet(object.propertyAddressCity) ? String(object.propertyAddressCity) : "",
      propertyAddressState: isSet(object.propertyAddressState) ? String(object.propertyAddressState) : "",
      propertyAddressStreet: isSet(object.propertyAddressStreet) ? String(object.propertyAddressStreet) : "",
      propertyAddressPostalCode: isSet(object.propertyAddressPostalCode)
        ? String(object.propertyAddressPostalCode)
        : "",
      propertyRegion: isSet(object.propertyRegion) ? String(object.propertyRegion) : "",
      propertyCountry: isSet(object.propertyCountry) ? String(object.propertyCountry) : "",
      interestRatePercentage: isSet(object.interestRatePercentage) ? Number(object.interestRatePercentage) : 0,
      interestRateType: isSet(object.interestRateType) ? String(object.interestRateType) : "",
    };
  },

  toJSON(message: MortgageAccount): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.plaidAccountId !== undefined && (obj.plaidAccountId = message.plaidAccountId);
    message.accountNumber !== undefined && (obj.accountNumber = message.accountNumber);
    message.currentLateFee !== undefined && (obj.currentLateFee = message.currentLateFee);
    message.escrowBalance !== undefined && (obj.escrowBalance = message.escrowBalance);
    message.hasPmi !== undefined && (obj.hasPmi = message.hasPmi);
    message.hasPrepaymentPenalty !== undefined && (obj.hasPrepaymentPenalty = message.hasPrepaymentPenalty);
    message.lastPaymentAmount !== undefined && (obj.lastPaymentAmount = message.lastPaymentAmount);
    message.lastPaymentDate !== undefined && (obj.lastPaymentDate = message.lastPaymentDate);
    message.loanTerm !== undefined && (obj.loanTerm = message.loanTerm);
    message.loanTypeDescription !== undefined && (obj.loanTypeDescription = message.loanTypeDescription);
    message.maturityDate !== undefined && (obj.maturityDate = message.maturityDate);
    message.nextMonthlyPayment !== undefined && (obj.nextMonthlyPayment = message.nextMonthlyPayment);
    message.nextPaymentDueDate !== undefined && (obj.nextPaymentDueDate = message.nextPaymentDueDate);
    message.originalPrincipalBalance !== undefined && (obj.originalPrincipalBalance = message.originalPrincipalBalance);
    message.originalPropertyValue !== undefined && (obj.originalPropertyValue = message.originalPropertyValue);
    message.outstandingPrincipalBalance !== undefined &&
      (obj.outstandingPrincipalBalance = message.outstandingPrincipalBalance);
    message.paymentAmount !== undefined && (obj.paymentAmount = message.paymentAmount);
    message.paymentDate !== undefined && (obj.paymentDate = message.paymentDate);
    message.originationDate !== undefined && (obj.originationDate = message.originationDate);
    message.originationPrincipalAmount !== undefined &&
      (obj.originationPrincipalAmount = message.originationPrincipalAmount);
    message.pastDueAmount !== undefined && (obj.pastDueAmount = message.pastDueAmount);
    message.ytdInterestPaid !== undefined && (obj.ytdInterestPaid = message.ytdInterestPaid);
    message.ytdPrincipalPaid !== undefined && (obj.ytdPrincipalPaid = message.ytdPrincipalPaid);
    message.propertyAddressCity !== undefined && (obj.propertyAddressCity = message.propertyAddressCity);
    message.propertyAddressState !== undefined && (obj.propertyAddressState = message.propertyAddressState);
    message.propertyAddressStreet !== undefined && (obj.propertyAddressStreet = message.propertyAddressStreet);
    message.propertyAddressPostalCode !== undefined &&
      (obj.propertyAddressPostalCode = message.propertyAddressPostalCode);
    message.propertyRegion !== undefined && (obj.propertyRegion = message.propertyRegion);
    message.propertyCountry !== undefined && (obj.propertyCountry = message.propertyCountry);
    message.interestRatePercentage !== undefined && (obj.interestRatePercentage = message.interestRatePercentage);
    message.interestRateType !== undefined && (obj.interestRateType = message.interestRateType);
    return obj;
  },

  create<I extends Exact<DeepPartial<MortgageAccount>, I>>(base?: I): MortgageAccount {
    return MortgageAccount.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<MortgageAccount>, I>>(object: I): MortgageAccount {
    const message = createBaseMortgageAccount();
    message.id = object.id ?? 0;
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.accountNumber = object.accountNumber ?? "";
    message.currentLateFee = object.currentLateFee ?? 0;
    message.escrowBalance = object.escrowBalance ?? 0;
    message.hasPmi = object.hasPmi ?? false;
    message.hasPrepaymentPenalty = object.hasPrepaymentPenalty ?? false;
    message.lastPaymentAmount = object.lastPaymentAmount ?? 0;
    message.lastPaymentDate = object.lastPaymentDate ?? "";
    message.loanTerm = object.loanTerm ?? "";
    message.loanTypeDescription = object.loanTypeDescription ?? "";
    message.maturityDate = object.maturityDate ?? "";
    message.nextMonthlyPayment = object.nextMonthlyPayment ?? 0;
    message.nextPaymentDueDate = object.nextPaymentDueDate ?? "";
    message.originalPrincipalBalance = object.originalPrincipalBalance ?? 0;
    message.originalPropertyValue = object.originalPropertyValue ?? 0;
    message.outstandingPrincipalBalance = object.outstandingPrincipalBalance ?? 0;
    message.paymentAmount = object.paymentAmount ?? 0;
    message.paymentDate = object.paymentDate ?? "";
    message.originationDate = object.originationDate ?? "";
    message.originationPrincipalAmount = object.originationPrincipalAmount ?? 0;
    message.pastDueAmount = object.pastDueAmount ?? 0;
    message.ytdInterestPaid = object.ytdInterestPaid ?? 0;
    message.ytdPrincipalPaid = object.ytdPrincipalPaid ?? 0;
    message.propertyAddressCity = object.propertyAddressCity ?? "";
    message.propertyAddressState = object.propertyAddressState ?? "";
    message.propertyAddressStreet = object.propertyAddressStreet ?? "";
    message.propertyAddressPostalCode = object.propertyAddressPostalCode ?? "";
    message.propertyRegion = object.propertyRegion ?? "";
    message.propertyCountry = object.propertyCountry ?? "";
    message.interestRatePercentage = object.interestRatePercentage ?? 0;
    message.interestRateType = object.interestRateType ?? "";
    return message;
  },
};

function createBaseInvestmentAccount(): InvestmentAccount {
  return {
    id: 0,
    userId: 0,
    name: "",
    number: "",
    type: "",
    balance: 0,
    currentFunds: 0,
    balanceLimit: 0,
    plaidAccountId: "",
    subtype: "",
    holdings: [],
    securities: [],
  };
}

export const InvestmentAccount = {
  encode(message: InvestmentAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.number !== "") {
      writer.uint32(34).string(message.number);
    }
    if (message.type !== "") {
      writer.uint32(42).string(message.type);
    }
    if (message.balance !== 0) {
      writer.uint32(53).float(message.balance);
    }
    if (message.currentFunds !== 0) {
      writer.uint32(73).double(message.currentFunds);
    }
    if (message.balanceLimit !== 0) {
      writer.uint32(80).uint64(message.balanceLimit);
    }
    if (message.plaidAccountId !== "") {
      writer.uint32(98).string(message.plaidAccountId);
    }
    if (message.subtype !== "") {
      writer.uint32(106).string(message.subtype);
    }
    for (const v of message.holdings) {
      InvesmentHolding.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.securities) {
      InvestmentSecurity.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InvestmentAccount {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInvestmentAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.number = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.type = reader.string();
          continue;
        case 6:
          if (tag !== 53) {
            break;
          }

          message.balance = reader.float();
          continue;
        case 9:
          if (tag !== 73) {
            break;
          }

          message.currentFunds = reader.double();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.balanceLimit = longToNumber(reader.uint64() as Long);
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.subtype = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.holdings.push(InvesmentHolding.decode(reader, reader.uint32()));
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.securities.push(InvestmentSecurity.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InvestmentAccount {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      number: isSet(object.number) ? String(object.number) : "",
      type: isSet(object.type) ? String(object.type) : "",
      balance: isSet(object.balance) ? Number(object.balance) : 0,
      currentFunds: isSet(object.currentFunds) ? Number(object.currentFunds) : 0,
      balanceLimit: isSet(object.balanceLimit) ? Number(object.balanceLimit) : 0,
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      subtype: isSet(object.subtype) ? String(object.subtype) : "",
      holdings: Array.isArray(object?.holdings) ? object.holdings.map((e: any) => InvesmentHolding.fromJSON(e)) : [],
      securities: Array.isArray(object?.securities)
        ? object.securities.map((e: any) => InvestmentSecurity.fromJSON(e))
        : [],
    };
  },

  toJSON(message: InvestmentAccount): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.name !== undefined && (obj.name = message.name);
    message.number !== undefined && (obj.number = message.number);
    message.type !== undefined && (obj.type = message.type);
    message.balance !== undefined && (obj.balance = message.balance);
    message.currentFunds !== undefined && (obj.currentFunds = message.currentFunds);
    message.balanceLimit !== undefined && (obj.balanceLimit = Math.round(message.balanceLimit));
    message.plaidAccountId !== undefined && (obj.plaidAccountId = message.plaidAccountId);
    message.subtype !== undefined && (obj.subtype = message.subtype);
    if (message.holdings) {
      obj.holdings = message.holdings.map((e) => e ? InvesmentHolding.toJSON(e) : undefined);
    } else {
      obj.holdings = [];
    }
    if (message.securities) {
      obj.securities = message.securities.map((e) => e ? InvestmentSecurity.toJSON(e) : undefined);
    } else {
      obj.securities = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<InvestmentAccount>, I>>(base?: I): InvestmentAccount {
    return InvestmentAccount.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<InvestmentAccount>, I>>(object: I): InvestmentAccount {
    const message = createBaseInvestmentAccount();
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.name = object.name ?? "";
    message.number = object.number ?? "";
    message.type = object.type ?? "";
    message.balance = object.balance ?? 0;
    message.currentFunds = object.currentFunds ?? 0;
    message.balanceLimit = object.balanceLimit ?? 0;
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.subtype = object.subtype ?? "";
    message.holdings = object.holdings?.map((e) => InvesmentHolding.fromPartial(e)) || [];
    message.securities = object.securities?.map((e) => InvestmentSecurity.fromPartial(e)) || [];
    return message;
  },
};

function createBaseBankAccount(): BankAccount {
  return {
    id: 0,
    userId: 0,
    name: "",
    number: "",
    type: 0,
    balance: 0,
    currency: "",
    currentFunds: 0,
    balanceLimit: 0,
    pockets: [],
    plaidAccountId: "",
    subtype: "",
    status: 0,
  };
}

export const BankAccount = {
  encode(message: BankAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.number !== "") {
      writer.uint32(34).string(message.number);
    }
    if (message.type !== 0) {
      writer.uint32(40).int32(message.type);
    }
    if (message.balance !== 0) {
      writer.uint32(53).float(message.balance);
    }
    if (message.currency !== "") {
      writer.uint32(58).string(message.currency);
    }
    if (message.currentFunds !== 0) {
      writer.uint32(73).double(message.currentFunds);
    }
    if (message.balanceLimit !== 0) {
      writer.uint32(80).uint64(message.balanceLimit);
    }
    for (const v of message.pockets) {
      Pocket.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    if (message.plaidAccountId !== "") {
      writer.uint32(106).string(message.plaidAccountId);
    }
    if (message.subtype !== "") {
      writer.uint32(114).string(message.subtype);
    }
    if (message.status !== 0) {
      writer.uint32(120).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BankAccount {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBankAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.number = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 6:
          if (tag !== 53) {
            break;
          }

          message.balance = reader.float();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.currency = reader.string();
          continue;
        case 9:
          if (tag !== 73) {
            break;
          }

          message.currentFunds = reader.double();
          continue;
        case 10:
          if (tag !== 80) {
            break;
          }

          message.balanceLimit = longToNumber(reader.uint64() as Long);
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.pockets.push(Pocket.decode(reader, reader.uint32()));
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.subtype = reader.string();
          continue;
        case 15:
          if (tag !== 120) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BankAccount {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      number: isSet(object.number) ? String(object.number) : "",
      type: isSet(object.type) ? bankAccountTypeFromJSON(object.type) : 0,
      balance: isSet(object.balance) ? Number(object.balance) : 0,
      currency: isSet(object.currency) ? String(object.currency) : "",
      currentFunds: isSet(object.currentFunds) ? Number(object.currentFunds) : 0,
      balanceLimit: isSet(object.balanceLimit) ? Number(object.balanceLimit) : 0,
      pockets: Array.isArray(object?.pockets) ? object.pockets.map((e: any) => Pocket.fromJSON(e)) : [],
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      subtype: isSet(object.subtype) ? String(object.subtype) : "",
      status: isSet(object.status) ? bankAccountStatusFromJSON(object.status) : 0,
    };
  },

  toJSON(message: BankAccount): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.name !== undefined && (obj.name = message.name);
    message.number !== undefined && (obj.number = message.number);
    message.type !== undefined && (obj.type = bankAccountTypeToJSON(message.type));
    message.balance !== undefined && (obj.balance = message.balance);
    message.currency !== undefined && (obj.currency = message.currency);
    message.currentFunds !== undefined && (obj.currentFunds = message.currentFunds);
    message.balanceLimit !== undefined && (obj.balanceLimit = Math.round(message.balanceLimit));
    if (message.pockets) {
      obj.pockets = message.pockets.map((e) => e ? Pocket.toJSON(e) : undefined);
    } else {
      obj.pockets = [];
    }
    message.plaidAccountId !== undefined && (obj.plaidAccountId = message.plaidAccountId);
    message.subtype !== undefined && (obj.subtype = message.subtype);
    message.status !== undefined && (obj.status = bankAccountStatusToJSON(message.status));
    return obj;
  },

  create<I extends Exact<DeepPartial<BankAccount>, I>>(base?: I): BankAccount {
    return BankAccount.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BankAccount>, I>>(object: I): BankAccount {
    const message = createBaseBankAccount();
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.name = object.name ?? "";
    message.number = object.number ?? "";
    message.type = object.type ?? 0;
    message.balance = object.balance ?? 0;
    message.currency = object.currency ?? "";
    message.currentFunds = object.currentFunds ?? 0;
    message.balanceLimit = object.balanceLimit ?? 0;
    message.pockets = object.pockets?.map((e) => Pocket.fromPartial(e)) || [];
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.subtype = object.subtype ?? "";
    message.status = object.status ?? 0;
    return message;
  },
};

function createBasePocket(): Pocket {
  return { id: 0, goals: [], type: 0 };
}

export const Pocket = {
  encode(message: Pocket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    for (const v of message.goals) {
      SmartGoal.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.type !== 0) {
      writer.uint32(32).int32(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Pocket {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePocket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.goals.push(SmartGoal.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Pocket {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      goals: Array.isArray(object?.goals) ? object.goals.map((e: any) => SmartGoal.fromJSON(e)) : [],
      type: isSet(object.type) ? pocketTypeFromJSON(object.type) : 0,
    };
  },

  toJSON(message: Pocket): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    if (message.goals) {
      obj.goals = message.goals.map((e) => e ? SmartGoal.toJSON(e) : undefined);
    } else {
      obj.goals = [];
    }
    message.type !== undefined && (obj.type = pocketTypeToJSON(message.type));
    return obj;
  },

  create<I extends Exact<DeepPartial<Pocket>, I>>(base?: I): Pocket {
    return Pocket.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Pocket>, I>>(object: I): Pocket {
    const message = createBasePocket();
    message.id = object.id ?? 0;
    message.goals = object.goals?.map((e) => SmartGoal.fromPartial(e)) || [];
    message.type = object.type ?? 0;
    return message;
  },
};

function createBaseSmartGoal(): SmartGoal {
  return {
    id: 0,
    userId: 0,
    name: "",
    description: "",
    isCompleted: false,
    goalType: 0,
    duration: "",
    startDate: "",
    endDate: "",
    targetAmount: "",
    currentAmount: "",
    milestones: [],
    forecasts: undefined,
  };
}

export const SmartGoal = {
  encode(message: SmartGoal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.isCompleted === true) {
      writer.uint32(40).bool(message.isCompleted);
    }
    if (message.goalType !== 0) {
      writer.uint32(72).int32(message.goalType);
    }
    if (message.duration !== "") {
      writer.uint32(82).string(message.duration);
    }
    if (message.startDate !== "") {
      writer.uint32(90).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(98).string(message.endDate);
    }
    if (message.targetAmount !== "") {
      writer.uint32(106).string(message.targetAmount);
    }
    if (message.currentAmount !== "") {
      writer.uint32(114).string(message.currentAmount);
    }
    for (const v of message.milestones) {
      Milestone.encode(v!, writer.uint32(122).fork()).ldelim();
    }
    if (message.forecasts !== undefined) {
      Forecast.encode(message.forecasts, writer.uint32(130).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SmartGoal {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSmartGoal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.userId = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.description = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.isCompleted = reader.bool();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.goalType = reader.int32() as any;
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.duration = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.startDate = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.endDate = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.targetAmount = reader.string();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.currentAmount = reader.string();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.milestones.push(Milestone.decode(reader, reader.uint32()));
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.forecasts = Forecast.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SmartGoal {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      userId: isSet(object.userId) ? Number(object.userId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      isCompleted: isSet(object.isCompleted) ? Boolean(object.isCompleted) : false,
      goalType: isSet(object.goalType) ? goalTypeFromJSON(object.goalType) : 0,
      duration: isSet(object.duration) ? String(object.duration) : "",
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      targetAmount: isSet(object.targetAmount) ? String(object.targetAmount) : "",
      currentAmount: isSet(object.currentAmount) ? String(object.currentAmount) : "",
      milestones: Array.isArray(object?.milestones) ? object.milestones.map((e: any) => Milestone.fromJSON(e)) : [],
      forecasts: isSet(object.forecasts) ? Forecast.fromJSON(object.forecasts) : undefined,
    };
  },

  toJSON(message: SmartGoal): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.userId !== undefined && (obj.userId = Math.round(message.userId));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.isCompleted !== undefined && (obj.isCompleted = message.isCompleted);
    message.goalType !== undefined && (obj.goalType = goalTypeToJSON(message.goalType));
    message.duration !== undefined && (obj.duration = message.duration);
    message.startDate !== undefined && (obj.startDate = message.startDate);
    message.endDate !== undefined && (obj.endDate = message.endDate);
    message.targetAmount !== undefined && (obj.targetAmount = message.targetAmount);
    message.currentAmount !== undefined && (obj.currentAmount = message.currentAmount);
    if (message.milestones) {
      obj.milestones = message.milestones.map((e) => e ? Milestone.toJSON(e) : undefined);
    } else {
      obj.milestones = [];
    }
    message.forecasts !== undefined &&
      (obj.forecasts = message.forecasts ? Forecast.toJSON(message.forecasts) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<SmartGoal>, I>>(base?: I): SmartGoal {
    return SmartGoal.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SmartGoal>, I>>(object: I): SmartGoal {
    const message = createBaseSmartGoal();
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.isCompleted = object.isCompleted ?? false;
    message.goalType = object.goalType ?? 0;
    message.duration = object.duration ?? "";
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.targetAmount = object.targetAmount ?? "";
    message.currentAmount = object.currentAmount ?? "";
    message.milestones = object.milestones?.map((e) => Milestone.fromPartial(e)) || [];
    message.forecasts = (object.forecasts !== undefined && object.forecasts !== null)
      ? Forecast.fromPartial(object.forecasts)
      : undefined;
    return message;
  },
};

function createBaseForecast(): Forecast {
  return { id: 0, forecastedAmount: "", forecastedCompletionDate: "", varianceAmount: "" };
}

export const Forecast = {
  encode(message: Forecast, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.forecastedAmount !== "") {
      writer.uint32(18).string(message.forecastedAmount);
    }
    if (message.forecastedCompletionDate !== "") {
      writer.uint32(26).string(message.forecastedCompletionDate);
    }
    if (message.varianceAmount !== "") {
      writer.uint32(34).string(message.varianceAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Forecast {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForecast();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.forecastedAmount = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.forecastedCompletionDate = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.varianceAmount = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Forecast {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      forecastedAmount: isSet(object.forecastedAmount) ? String(object.forecastedAmount) : "",
      forecastedCompletionDate: isSet(object.forecastedCompletionDate) ? String(object.forecastedCompletionDate) : "",
      varianceAmount: isSet(object.varianceAmount) ? String(object.varianceAmount) : "",
    };
  },

  toJSON(message: Forecast): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.forecastedAmount !== undefined && (obj.forecastedAmount = message.forecastedAmount);
    message.forecastedCompletionDate !== undefined && (obj.forecastedCompletionDate = message.forecastedCompletionDate);
    message.varianceAmount !== undefined && (obj.varianceAmount = message.varianceAmount);
    return obj;
  },

  create<I extends Exact<DeepPartial<Forecast>, I>>(base?: I): Forecast {
    return Forecast.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Forecast>, I>>(object: I): Forecast {
    const message = createBaseForecast();
    message.id = object.id ?? 0;
    message.forecastedAmount = object.forecastedAmount ?? "";
    message.forecastedCompletionDate = object.forecastedCompletionDate ?? "";
    message.varianceAmount = object.varianceAmount ?? "";
    return message;
  },
};

function createBaseMilestone(): Milestone {
  return { id: 0, name: "", description: "", targetDate: "", targetAmount: "", isCompleted: false, budget: undefined };
}

export const Milestone = {
  encode(message: Milestone, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.targetDate !== "") {
      writer.uint32(34).string(message.targetDate);
    }
    if (message.targetAmount !== "") {
      writer.uint32(42).string(message.targetAmount);
    }
    if (message.isCompleted === true) {
      writer.uint32(48).bool(message.isCompleted);
    }
    if (message.budget !== undefined) {
      Budget.encode(message.budget, writer.uint32(106).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Milestone {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMilestone();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.targetDate = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.targetAmount = reader.string();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.isCompleted = reader.bool();
          continue;
        case 13:
          if (tag !== 106) {
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

  fromJSON(object: any): Milestone {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      targetDate: isSet(object.targetDate) ? String(object.targetDate) : "",
      targetAmount: isSet(object.targetAmount) ? String(object.targetAmount) : "",
      isCompleted: isSet(object.isCompleted) ? Boolean(object.isCompleted) : false,
      budget: isSet(object.budget) ? Budget.fromJSON(object.budget) : undefined,
    };
  },

  toJSON(message: Milestone): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.targetDate !== undefined && (obj.targetDate = message.targetDate);
    message.targetAmount !== undefined && (obj.targetAmount = message.targetAmount);
    message.isCompleted !== undefined && (obj.isCompleted = message.isCompleted);
    message.budget !== undefined && (obj.budget = message.budget ? Budget.toJSON(message.budget) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<Milestone>, I>>(base?: I): Milestone {
    return Milestone.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Milestone>, I>>(object: I): Milestone {
    const message = createBaseMilestone();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.targetDate = object.targetDate ?? "";
    message.targetAmount = object.targetAmount ?? "";
    message.isCompleted = object.isCompleted ?? false;
    message.budget = (object.budget !== undefined && object.budget !== null)
      ? Budget.fromPartial(object.budget)
      : undefined;
    return message;
  },
};

function createBaseBudget(): Budget {
  return { id: 0, name: "", description: "", startDate: "", endDate: "", category: undefined };
}

export const Budget = {
  encode(message: Budget, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.startDate !== "") {
      writer.uint32(34).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(42).string(message.endDate);
    }
    if (message.category !== undefined) {
      Category.encode(message.category, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Budget {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBudget();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.startDate = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.endDate = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.category = Category.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Budget {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      category: isSet(object.category) ? Category.fromJSON(object.category) : undefined,
    };
  },

  toJSON(message: Budget): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.startDate !== undefined && (obj.startDate = message.startDate);
    message.endDate !== undefined && (obj.endDate = message.endDate);
    message.category !== undefined && (obj.category = message.category ? Category.toJSON(message.category) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<Budget>, I>>(base?: I): Budget {
    return Budget.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Budget>, I>>(object: I): Budget {
    const message = createBaseBudget();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.category = (object.category !== undefined && object.category !== null)
      ? Category.fromPartial(object.category)
      : undefined;
    return message;
  },
};

function createBaseCategory(): Category {
  return { id: 0, name: "", description: "", subcategories: [] };
}

export const Category = {
  encode(message: Category, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    for (const v of message.subcategories) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Category {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCategory();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.subcategories.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Category {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      subcategories: Array.isArray(object?.subcategories) ? object.subcategories.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: Category): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    if (message.subcategories) {
      obj.subcategories = message.subcategories.map((e) => e);
    } else {
      obj.subcategories = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Category>, I>>(base?: I): Category {
    return Category.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Category>, I>>(object: I): Category {
    const message = createBaseCategory();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.subcategories = object.subcategories?.map((e) => e) || [];
    return message;
  },
};

function createBaseInvesmentHolding(): InvesmentHolding {
  return {
    id: 0,
    name: "",
    plaidAccountId: "",
    costBasis: 0,
    institutionPrice: 0,
    institutionPriceAsOf: "",
    institutionPriceDatetime: "",
    institutionValue: 0,
    isoCurrencyCode: "",
    quantity: 0,
    securityId: "",
    unofficialCurrencyCode: "",
  };
}

export const InvesmentHolding = {
  encode(message: InvesmentHolding, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.plaidAccountId !== "") {
      writer.uint32(26).string(message.plaidAccountId);
    }
    if (message.costBasis !== 0) {
      writer.uint32(33).double(message.costBasis);
    }
    if (message.institutionPrice !== 0) {
      writer.uint32(41).double(message.institutionPrice);
    }
    if (message.institutionPriceAsOf !== "") {
      writer.uint32(50).string(message.institutionPriceAsOf);
    }
    if (message.institutionPriceDatetime !== "") {
      writer.uint32(58).string(message.institutionPriceDatetime);
    }
    if (message.institutionValue !== 0) {
      writer.uint32(65).double(message.institutionValue);
    }
    if (message.isoCurrencyCode !== "") {
      writer.uint32(74).string(message.isoCurrencyCode);
    }
    if (message.quantity !== 0) {
      writer.uint32(81).double(message.quantity);
    }
    if (message.securityId !== "") {
      writer.uint32(90).string(message.securityId);
    }
    if (message.unofficialCurrencyCode !== "") {
      writer.uint32(98).string(message.unofficialCurrencyCode);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InvesmentHolding {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInvesmentHolding();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.plaidAccountId = reader.string();
          continue;
        case 4:
          if (tag !== 33) {
            break;
          }

          message.costBasis = reader.double();
          continue;
        case 5:
          if (tag !== 41) {
            break;
          }

          message.institutionPrice = reader.double();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.institutionPriceAsOf = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.institutionPriceDatetime = reader.string();
          continue;
        case 8:
          if (tag !== 65) {
            break;
          }

          message.institutionValue = reader.double();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.isoCurrencyCode = reader.string();
          continue;
        case 10:
          if (tag !== 81) {
            break;
          }

          message.quantity = reader.double();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.securityId = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.unofficialCurrencyCode = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InvesmentHolding {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      plaidAccountId: isSet(object.plaidAccountId) ? String(object.plaidAccountId) : "",
      costBasis: isSet(object.costBasis) ? Number(object.costBasis) : 0,
      institutionPrice: isSet(object.institutionPrice) ? Number(object.institutionPrice) : 0,
      institutionPriceAsOf: isSet(object.institutionPriceAsOf) ? String(object.institutionPriceAsOf) : "",
      institutionPriceDatetime: isSet(object.institutionPriceDatetime) ? String(object.institutionPriceDatetime) : "",
      institutionValue: isSet(object.institutionValue) ? Number(object.institutionValue) : 0,
      isoCurrencyCode: isSet(object.isoCurrencyCode) ? String(object.isoCurrencyCode) : "",
      quantity: isSet(object.quantity) ? Number(object.quantity) : 0,
      securityId: isSet(object.securityId) ? String(object.securityId) : "",
      unofficialCurrencyCode: isSet(object.unofficialCurrencyCode) ? String(object.unofficialCurrencyCode) : "",
    };
  },

  toJSON(message: InvesmentHolding): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.plaidAccountId !== undefined && (obj.plaidAccountId = message.plaidAccountId);
    message.costBasis !== undefined && (obj.costBasis = message.costBasis);
    message.institutionPrice !== undefined && (obj.institutionPrice = message.institutionPrice);
    message.institutionPriceAsOf !== undefined && (obj.institutionPriceAsOf = message.institutionPriceAsOf);
    message.institutionPriceDatetime !== undefined && (obj.institutionPriceDatetime = message.institutionPriceDatetime);
    message.institutionValue !== undefined && (obj.institutionValue = message.institutionValue);
    message.isoCurrencyCode !== undefined && (obj.isoCurrencyCode = message.isoCurrencyCode);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    message.securityId !== undefined && (obj.securityId = message.securityId);
    message.unofficialCurrencyCode !== undefined && (obj.unofficialCurrencyCode = message.unofficialCurrencyCode);
    return obj;
  },

  create<I extends Exact<DeepPartial<InvesmentHolding>, I>>(base?: I): InvesmentHolding {
    return InvesmentHolding.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<InvesmentHolding>, I>>(object: I): InvesmentHolding {
    const message = createBaseInvesmentHolding();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.plaidAccountId = object.plaidAccountId ?? "";
    message.costBasis = object.costBasis ?? 0;
    message.institutionPrice = object.institutionPrice ?? 0;
    message.institutionPriceAsOf = object.institutionPriceAsOf ?? "";
    message.institutionPriceDatetime = object.institutionPriceDatetime ?? "";
    message.institutionValue = object.institutionValue ?? 0;
    message.isoCurrencyCode = object.isoCurrencyCode ?? "";
    message.quantity = object.quantity ?? 0;
    message.securityId = object.securityId ?? "";
    message.unofficialCurrencyCode = object.unofficialCurrencyCode ?? "";
    return message;
  },
};

function createBaseInvestmentSecurity(): InvestmentSecurity {
  return {
    id: 0,
    closePrice: 0,
    closePriceAsOf: "",
    cusip: "",
    institutionId: "",
    institutionSecurityId: "",
    isCashEquivalent: false,
    isin: "",
    isoCurrencyCode: "",
    name: "",
    proxySecurityId: "",
    securityId: "",
    sedol: "",
    tickerSymbol: "",
    type: "",
    unofficialCurrencyCode: "",
    updateDatetime: "",
  };
}

export const InvestmentSecurity = {
  encode(message: InvestmentSecurity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.closePrice !== 0) {
      writer.uint32(17).double(message.closePrice);
    }
    if (message.closePriceAsOf !== "") {
      writer.uint32(26).string(message.closePriceAsOf);
    }
    if (message.cusip !== "") {
      writer.uint32(34).string(message.cusip);
    }
    if (message.institutionId !== "") {
      writer.uint32(42).string(message.institutionId);
    }
    if (message.institutionSecurityId !== "") {
      writer.uint32(50).string(message.institutionSecurityId);
    }
    if (message.isCashEquivalent === true) {
      writer.uint32(56).bool(message.isCashEquivalent);
    }
    if (message.isin !== "") {
      writer.uint32(66).string(message.isin);
    }
    if (message.isoCurrencyCode !== "") {
      writer.uint32(74).string(message.isoCurrencyCode);
    }
    if (message.name !== "") {
      writer.uint32(82).string(message.name);
    }
    if (message.proxySecurityId !== "") {
      writer.uint32(90).string(message.proxySecurityId);
    }
    if (message.securityId !== "") {
      writer.uint32(98).string(message.securityId);
    }
    if (message.sedol !== "") {
      writer.uint32(106).string(message.sedol);
    }
    if (message.tickerSymbol !== "") {
      writer.uint32(114).string(message.tickerSymbol);
    }
    if (message.type !== "") {
      writer.uint32(122).string(message.type);
    }
    if (message.unofficialCurrencyCode !== "") {
      writer.uint32(130).string(message.unofficialCurrencyCode);
    }
    if (message.updateDatetime !== "") {
      writer.uint32(138).string(message.updateDatetime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InvestmentSecurity {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInvestmentSecurity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 17) {
            break;
          }

          message.closePrice = reader.double();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.closePriceAsOf = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.cusip = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.institutionId = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.institutionSecurityId = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.isCashEquivalent = reader.bool();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.isin = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.isoCurrencyCode = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.name = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.proxySecurityId = reader.string();
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.securityId = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.sedol = reader.string();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.tickerSymbol = reader.string();
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.type = reader.string();
          continue;
        case 16:
          if (tag !== 130) {
            break;
          }

          message.unofficialCurrencyCode = reader.string();
          continue;
        case 17:
          if (tag !== 138) {
            break;
          }

          message.updateDatetime = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InvestmentSecurity {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      closePrice: isSet(object.closePrice) ? Number(object.closePrice) : 0,
      closePriceAsOf: isSet(object.closePriceAsOf) ? String(object.closePriceAsOf) : "",
      cusip: isSet(object.cusip) ? String(object.cusip) : "",
      institutionId: isSet(object.institutionId) ? String(object.institutionId) : "",
      institutionSecurityId: isSet(object.institutionSecurityId) ? String(object.institutionSecurityId) : "",
      isCashEquivalent: isSet(object.isCashEquivalent) ? Boolean(object.isCashEquivalent) : false,
      isin: isSet(object.isin) ? String(object.isin) : "",
      isoCurrencyCode: isSet(object.isoCurrencyCode) ? String(object.isoCurrencyCode) : "",
      name: isSet(object.name) ? String(object.name) : "",
      proxySecurityId: isSet(object.proxySecurityId) ? String(object.proxySecurityId) : "",
      securityId: isSet(object.securityId) ? String(object.securityId) : "",
      sedol: isSet(object.sedol) ? String(object.sedol) : "",
      tickerSymbol: isSet(object.tickerSymbol) ? String(object.tickerSymbol) : "",
      type: isSet(object.type) ? String(object.type) : "",
      unofficialCurrencyCode: isSet(object.unofficialCurrencyCode) ? String(object.unofficialCurrencyCode) : "",
      updateDatetime: isSet(object.updateDatetime) ? String(object.updateDatetime) : "",
    };
  },

  toJSON(message: InvestmentSecurity): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.closePrice !== undefined && (obj.closePrice = message.closePrice);
    message.closePriceAsOf !== undefined && (obj.closePriceAsOf = message.closePriceAsOf);
    message.cusip !== undefined && (obj.cusip = message.cusip);
    message.institutionId !== undefined && (obj.institutionId = message.institutionId);
    message.institutionSecurityId !== undefined && (obj.institutionSecurityId = message.institutionSecurityId);
    message.isCashEquivalent !== undefined && (obj.isCashEquivalent = message.isCashEquivalent);
    message.isin !== undefined && (obj.isin = message.isin);
    message.isoCurrencyCode !== undefined && (obj.isoCurrencyCode = message.isoCurrencyCode);
    message.name !== undefined && (obj.name = message.name);
    message.proxySecurityId !== undefined && (obj.proxySecurityId = message.proxySecurityId);
    message.securityId !== undefined && (obj.securityId = message.securityId);
    message.sedol !== undefined && (obj.sedol = message.sedol);
    message.tickerSymbol !== undefined && (obj.tickerSymbol = message.tickerSymbol);
    message.type !== undefined && (obj.type = message.type);
    message.unofficialCurrencyCode !== undefined && (obj.unofficialCurrencyCode = message.unofficialCurrencyCode);
    message.updateDatetime !== undefined && (obj.updateDatetime = message.updateDatetime);
    return obj;
  },

  create<I extends Exact<DeepPartial<InvestmentSecurity>, I>>(base?: I): InvestmentSecurity {
    return InvestmentSecurity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<InvestmentSecurity>, I>>(object: I): InvestmentSecurity {
    const message = createBaseInvestmentSecurity();
    message.id = object.id ?? 0;
    message.closePrice = object.closePrice ?? 0;
    message.closePriceAsOf = object.closePriceAsOf ?? "";
    message.cusip = object.cusip ?? "";
    message.institutionId = object.institutionId ?? "";
    message.institutionSecurityId = object.institutionSecurityId ?? "";
    message.isCashEquivalent = object.isCashEquivalent ?? false;
    message.isin = object.isin ?? "";
    message.isoCurrencyCode = object.isoCurrencyCode ?? "";
    message.name = object.name ?? "";
    message.proxySecurityId = object.proxySecurityId ?? "";
    message.securityId = object.securityId ?? "";
    message.sedol = object.sedol ?? "";
    message.tickerSymbol = object.tickerSymbol ?? "";
    message.type = object.type ?? "";
    message.unofficialCurrencyCode = object.unofficialCurrencyCode ?? "";
    message.updateDatetime = object.updateDatetime ?? "";
    return message;
  },
};

function createBaseApr(): Apr {
  return { id: 0, percentage: 0, type: "", balanceSubjectToApr: 0, interestChargeAmount: 0 };
}

export const Apr = {
  encode(message: Apr, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.percentage !== 0) {
      writer.uint32(17).double(message.percentage);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.balanceSubjectToApr !== 0) {
      writer.uint32(33).double(message.balanceSubjectToApr);
    }
    if (message.interestChargeAmount !== 0) {
      writer.uint32(41).double(message.interestChargeAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Apr {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseApr();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 17) {
            break;
          }

          message.percentage = reader.double();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.type = reader.string();
          continue;
        case 4:
          if (tag !== 33) {
            break;
          }

          message.balanceSubjectToApr = reader.double();
          continue;
        case 5:
          if (tag !== 41) {
            break;
          }

          message.interestChargeAmount = reader.double();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Apr {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      percentage: isSet(object.percentage) ? Number(object.percentage) : 0,
      type: isSet(object.type) ? String(object.type) : "",
      balanceSubjectToApr: isSet(object.balanceSubjectToApr) ? Number(object.balanceSubjectToApr) : 0,
      interestChargeAmount: isSet(object.interestChargeAmount) ? Number(object.interestChargeAmount) : 0,
    };
  },

  toJSON(message: Apr): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.percentage !== undefined && (obj.percentage = message.percentage);
    message.type !== undefined && (obj.type = message.type);
    message.balanceSubjectToApr !== undefined && (obj.balanceSubjectToApr = message.balanceSubjectToApr);
    message.interestChargeAmount !== undefined && (obj.interestChargeAmount = message.interestChargeAmount);
    return obj;
  },

  create<I extends Exact<DeepPartial<Apr>, I>>(base?: I): Apr {
    return Apr.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Apr>, I>>(object: I): Apr {
    const message = createBaseApr();
    message.id = object.id ?? 0;
    message.percentage = object.percentage ?? 0;
    message.type = object.type ?? "";
    message.balanceSubjectToApr = object.balanceSubjectToApr ?? 0;
    message.interestChargeAmount = object.interestChargeAmount ?? 0;
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
