package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	taskhandler "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/task-handler"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"

	encryptdecrypt "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/encrypt_decrypt"
	"github.com/gogo/status"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

// The PlaidWebhook type is a struct that contains various fields related to Plaid webhook events.
// @property {string} WebhookType - a string that indicates the type of webhook event being sent by
// Plaid, such as "TRANSACTIONS" or "ITEM".
// @property {string} WebhookCode - WebhookCode is a string property of the PlaidWebhook struct that
// represents the specific code associated with the webhook event. This code provides additional
// information about the type of event that occurred.
// @property {string} ItemId - The unique identifier for the Plaid item associated with the webhook
// event.
// @property Error - A map containing details about any errors that occurred during the webhook event.
// The keys in the map correspond to specific error fields, such as "error_type" and "error_code". The
// values are interface{} types, which can hold any data type.
// @property {string} NewWebhookURL - The new URL that Plaid will send future webhook events to. This
// property is only present if the webhook_type is "WEBHOOK_UPDATE_ACKNOWLEDGED".
// @property {int64} NewTransactions - The number of new transactions that have been added since the
// last webhook was received.
// @property {[]string} RemovedTransactions - RemovedTransactions is a property of the PlaidWebhook
// struct that represents an array of transaction IDs that were removed from the user's account. This
// property is included in the webhook payload when Plaid detects that a transaction has been deleted
// or voided.
// @property ConsentExpirationTime - ConsentExpirationTime is a property of the PlaidWebhook struct
// that represents the time when the user's consent for the Plaid API access will expire. It is a
// pointer to a time.Time object, which means it can be nil if there is no expiration time set.
// @property {bool} InitialUpdateComplete - A boolean value indicating whether the initial data
// synchronization for the Plaid item has been completed or not. If it is true, it means that the
// initial data synchronization has been completed.
// @property {string} HistoricalUpdateComplete - The value of this property indicates whether the
// historical transaction data for the item has been fully retrieved and updated. It can have one of
// the following values:
// @property {string} Environment - The environment in which the webhook was triggered. This could be
// "sandbox" or "production" depending on the Plaid environment being used.
// @property {[]string} AccountIds - A slice of strings representing the Plaid account IDs associated
// with the webhook event.
// @property {[]string} AccountIdsWithNewLiabilities - This property is a slice of strings that
// contains the Plaid account IDs for accounts that have new liabilities. Liabilities refer to debts or
// obligations that a person or organization owes, such as loans or credit card balances. This property
// is included in the Plaid webhook response to notify the client application of any
// @property {[]string} AccountIdsWithUpdatedLiabilities - This property is a slice of strings that
// contains the account IDs for accounts that have updated liabilities. Liabilities refer to any debts
// or obligations that a person or organization owes, such as loans or credit card balances. This
// property is used in Plaid's webhook system to notify the client application when there are
// @property {uint64} NewHoldings - The number of new holdings that have been added to the Plaid
// account associated with the webhook event. Holdings refer to the assets held by the account, such as
// stocks or bonds.
// @property {uint64} UpdatedHoldings - The number of holdings that have been updated in the Plaid
// account.
type PlaidWebhook struct {
	WebhookType                      string                 `json:"webhook_type"`
	WebhookCode                      string                 `json:"webhook_code"`
	ItemId                           string                 `json:"item_id"`
	Error                            map[string]interface{} `json:"error"`
	NewWebhookURL                    string                 `json:"new_webhook_url"`
	NewTransactions                  int64                  `json:"new_transactions"`
	RemovedTransactions              []string               `json:"removed_transactions"`
	ConsentExpirationTime            *time.Time             `json:"consent_expiration_time"`
	InitialUpdateComplete            bool                   `json:"initial_update_complete"`
	HistoricalUpdateComplete         string                 `json:"historical_update_complete"`
	Environment                      string                 `json:"environment"`
	AccountIds                       []string               `json:"account_ids"`
	AccountIdsWithNewLiabilities     []string               `json:"account_ids_with_new_liabilities"`
	AccountIdsWithUpdatedLiabilities []string               `json:"account_ids_with_updated_liabilities"`
	NewHoldings                      uint64                 `json:"new_holdings"`
	UpdatedHoldings                  uint64                 `json:"updated_holdings"`
}

// The PlaidClaims type is a struct that extends the jwt.StandardClaims type in Go.
// @property  - The `PlaidClaims` struct is a custom struct that extends the `jwt.StandardClaims`
// struct from the `github.com/dgrijalva/jwt-go` package. It is used to define additional claims that
// can be included in a JSON Web Token (JWT) generated by the Plaid
type PlaidClaims struct {
	jwt.StandardClaims
}

// Valid is used to validate the JWT token claims. It checks if the token has expired, if it was
// issued in the future, and if it is not yet valid. If any of these checks fail, it returns a
// `jwt.ValidationError` with the appropriate error code. The method also adds a buffer of 5 seconds to
// the issued at timestamp to account for any clock drift between the server and Plaid's side.
func (c PlaidClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := jwt.TimeFunc().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if !c.VerifyExpiresAt(now, false) {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	// I'm adding 5 seconds onto the now timestamp because I was running into an issue periodically that the clock on
	// the server would be slightly different than the clock on Plaid's side. And the issued at was causing a conflict
	// where it was just a few seconds (sometimes just one) out of bounds for this to be handled. So adding a buffer of
	// 5 seconds to account for any clock drift between our servers and Plaid's.
	if !c.VerifyIssuedAt(now+5, false) {
		vErr.Inner = fmt.Errorf("token used before issued, %d | %d", now, c.IssuedAt)
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if !c.VerifyNotBefore(now, false) {
		vErr.Inner = fmt.Errorf("token is not valid yet")
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

// handlerPlaidWebhook handles incoming webhook requests
// from Plaid. It reads the request body, verifies the Plaid-Verification header, parses the JWT token,
// and exchanges the kid for a public key to verify the token. If the token is valid, it unmarshals the
// payload into a `PlaidWebhook` struct and processes the webhook request using the
// `processWebhookRequest` function. If there are any errors during this process, it returns an
// appropriate HTTP response code and error message.
func (s *Server) handlerPlaidWebhook(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer req.Body.Close()

	ctx := req.Context()

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-process-plaid-webhook-request")
		defer span.End()
	}

	verification := req.Header.Get("Plaid-Verification")
	if strings.TrimSpace(verification) == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var kid string
	var claims PlaidClaims

	result, err := jwt.ParseWithClaims(verification, &claims, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method for the JWT token is ES256 per Plaid's documentation. Anything else should be
		// rejected.
		method, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok || method.Name != "ES256" {
			return nil, errors.Errorf("invalid signing method")
		}

		// Look for a kid field, we are going to use this to exchange for a public key that we can use to verify the
		// JWT token.
		value, ok := token.Header["kid"]
		if !ok {
			return nil, errors.Errorf("malformed JWT token, missing data")
		}

		// Make sure the value is a string, anything else is not valid and should be thrown out.
		kid, ok = value.(string)
		if !ok {
			return nil, errors.Errorf("malformed JWT token, expected string")
		}

		// Make sure that string has some kind of non-whitespace value.
		if strings.TrimSpace(kid) == "" {
			return nil, errors.Errorf("malformed JWT token, empty data")
		}

		s.logger.Info("about to get verification key from in-memory-cache", zap.String("kid", kid), zap.Any("verification", s.inMemoryWebhookVerification))
		// exchange the kid for a public key and try to obtain the verification key function
		keyFunction, err := s.inMemoryWebhookVerification.GetVerificationKey(ctx, kid)
		if err != nil {
			s.logger.Info("failed to get veririfaction key")
			return nil, errors.Wrap(err, "failed to get verification key for webhook")
		}

		return keyFunction.Keyfunc(token)
	})

	if err != nil {
		s.ErrorResponse(w, req, err.Error(), http.StatusBadRequest)
		return
	}

	if !result.Valid {
		s.ErrorResponse(w, req, err.Error(), http.StatusUnauthorized)
		return
	}

	var hook = PlaidWebhook{}
	if err := json.Unmarshal(payload, &hook); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook error while parsing basic request. %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.processWebhookRequest(ctx, &hook); err != nil {
		s.ErrorResponse(w, req, err.Error(), http.StatusInternalServerError)
		return
	}
}

// processWebhookRequest processes a webhook request from Plaid. It first checks if
// the request is valid and then extracts relevant information from the request. It then retrieves the
// access token and user ID associated with the Plaid item, and based on the webhook type and code,
// triggers background jobs to sync transactions, holdings, or liabilities. It also handles various
// webhook codes related to item status, such as errors, pending expiration, and revocation of user
// permission. Finally, it returns an error if any issues arise during the processing of the webhook
// request.
func (s *Server) processWebhookRequest(ctx context.Context, req *PlaidWebhook) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "missing request")
	}

	{
		fields := map[string]interface{}{
			"type":   req.WebhookType,
			"code":   req.WebhookCode,
			"itemId": req.ItemId,
		}
		switch strings.ToUpper(fmt.Sprintf("%s.%s", req.WebhookType, req.WebhookCode)) {
		case "TRANSACTIONS.DEFAULT_UPDATE":
			fields["newTransactions"] = req.NewTransactions
		case "TRANSACTIONS.TRANSACTIONS_REMOVED":
			fields["removedTransactions"] = req.RemovedTransactions
		}
		s.logger.Info("Handling webhook from Plaid.", zap.Any("fields", fields))
	}

	// get the link fro the respective item
	link, err := s.conn.GetLinkByItemId(ctx, req.ItemId, false)
	if err != nil {
		return err
	}

	// decript the access token
	accessToken, err := encryptdecrypt.DecryptUserAccessToken(ctx, link.Token, s.kms, s.logger)
	if err != nil {
		return err
	}

	// get the user id from one of the connected bank accounts
	if len(link.BankAccounts) == 0 {
		return status.Error(codes.Internal, "no bank accounts found")
	}

	userId := link.BankAccounts[0].UserId

	switch req.WebhookType {
	case "INVESTMENTS_TRANSACTIONS":
		switch req.WebhookCode {
		// Fired when new or updated investment transactions have been detected on an investment account.
		// The webhook typically fires in response to any newly added transactions or price changes to existing
		// transactions, most commonly after market close.
		case "DEFAULT_UPDATE":
			// Trigger a background job to sync the plaid transactions
			// 1. Get investment transactions
			accountIds := make([]string, 0, len(link.InvestmentAccounts))
			for _, acct := range link.InvestmentAccounts {
				accountIds = append(accountIds, acct.PlaidAccountId)
			}

			if err := taskhandler.DispatchPullInvestmentTransactionsTask(ctx, s.taskprocessor, s.instrumentation, s.logger, userId, link.Id, *accessToken, accountIds); err != nil {
				return err
			}
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	case "HOLDINGS": // TODO: DO THIS TONIGHT
		switch req.WebhookCode {
		// Fired when new or updated holdings have been detected on an investment account.
		// The webhook typically fires in response to any newly added holdings or price changes to existing
		// holdings, most commonly after market close.
		case "DEFAULT_UPDATE":
			accountIds := make([]string, 0, len(link.InvestmentAccounts))
			for _, acct := range link.InvestmentAccounts {
				accountIds = append(accountIds, acct.PlaidAccountId)
			}
			// Trigger a background job to sync the plaid holdings
			if err := taskhandler.DispatchPullInvestmentHoldingsTask(ctx, s.taskprocessor, s.instrumentation, s.logger, userId, link.Id, *accessToken, accountIds); err != nil {
				return err
			}
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	case "LIABILITIES": // TODO: DO THIS TONIGHT
		switch req.WebhookCode {
		// Liabilities webhooks are sent to indicate that new loans or updated loan fields for existing accounts are available.
		// will be fired when new or updated liabilities have been detected on a liabilities item
		case "DEFAULT_UPDATE":
			// get account ids for all new liabilities
			accountIds := append(req.AccountIdsWithNewLiabilities, req.AccountIdsWithUpdatedLiabilities...)
			// sync any existing/updated liability account as well as save any new liability accounts
			// to datastore layer
			if err := taskhandler.DispatchSyncLiabilityAccountsTask(ctx, s.taskprocessor, s.instrumentation, s.logger, userId, link.Id, *accessToken, accountIds); err != nil {
				return err
			}
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	case "TRANSACTIONS":
		switch link.PlaidLink.UsePlaidSync {
		case true:
			// Trigger a background job to sync the plaid transactions
			if err := taskhandler.DispatchPlaidSyncTask(ctx, s.taskprocessor, s.instrumentation, s.logger, userId, link.Id, *accessToken); err != nil {
				return err
			}
		default:
			switch req.WebhookCode {
			/*
				Fired when recurring transactions data is updated. This includes when a new recurring
				 stream is detected or when a new transaction is added to an existing recurring stream.
				 The RECURRING_TRANSACTIONS_UPDATE webhook will also fire when one or more attributes
				 of the recurring stream changes, which is usually a result of the addition, update,
				 or removal of transactions to the stream.

				After receipt of this webhook, the updated data can be fetched from /transactions/recurring/get.

				```json
				{
					"webhook_type": "TRANSACTIONS",
					"webhook_code": "RECURRING_TRANSACTIONS_UPDATE",
					"item_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
					"account_ids": [
						"lPNjeW1nR6CDn5okmGQ6hEpMo4lLNoSrzqDje",
						"lPNjeW1nR6CDn5okmGQ6hEpMo4lLNoSrzqDff"
					],
					"environment": "production"
				}
				```
			*/
			case "RECURRING_TRANSACTIONS_UPDATE":
				// DISPATCH task to pull the recurring transactions for the given account ids and update the database
				// with the new recurring transactions
				if err := taskhandler.DispatchPullUpdatedReCurringTransactionsTask(ctx, s.taskprocessor, s.instrumentation, s.logger, userId, link.Id, *accessToken, req.AccountIds); err != nil {
					return err
				}

				/*
					Fired when an Item's transactions change. This can be due to any event resulting in new changes,
					such as an initial 30-day transactions fetch upon the initialization of an Item with transactions,
					the backfill of historical transactions that occurs shortly after, or when changes are populated from a
					regularly-scheduled transactions update job. It is recommended to listen for the SYNC_UPDATES_AVAILABLE
					webhook when using the /transactions/sync endpoint. Note that when using /transactions/sync
					the older webhooks INITIAL_UPDATE, HISTORICAL_UPDATE, DEFAULT_UPDATE, and TRANSACTIONS_REMOVED,
					which are intended for use with /transactions/get, will also continue to be sent in order
					to maintain backwards compatibility. It is not necessary to listen for and respond to those
					webhooks when using /transactions/sync.

					After receipt of this webhook, the new changes can be fetched for the Item from /transactions/sync.
					Note that to receive this webhook for an Item, /transactions/sync must have been called at least once
					 on that Item. This means that, unlike the INITIAL_UPDATE and HISTORICAL_UPDATE webhooks, it will not
					  fire immediately upon Item creation. If /transactions/sync is called on an Item that was not initialized
					   with Transactions, the webhook will fire twice: once the first 30 days of transactions data has been
					   fetched, and a second time when all available historical transactions data has been fetched.

					This webhook will typically not fire in the Sandbox environment, due to the lack of dynamic transactions data.
					To test this webhook in Sandbox, call /sandbox/item/fire_webhook.

					``json
						{
							"webhook_type": "TRANSACTIONS",
							"webhook_code": "SYNC_UPDATES_AVAILABLE",
							"item_id": "wz666MBjYWTp2PDzzggYhM6oWWmBb",
							"initial_update_complete": true,
							"historical_update_complete": false,
							"environment": "production"
						}
					``
				*/
			case "SYNC_UPDATES_AVAILABLE":
				// Trigger a background job to sync the plaid transactions
				if err := taskhandler.DispatchPlaidSyncTask(ctx, s.taskprocessor, s.instrumentation, s.logger, userId, link.Id, *accessToken); err != nil {
					return err
				}
			default:
				s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
			}
		}
	case "ITEM":
		switch req.WebhookCode {
		// 	Fired when an error is encountered with an Item. The error can be resolved by
		//  having the user go through Link’s update mode.
		// 	TODO: need to figure out how to properly handle a user going through
		// 	Link's update mode. This is a very rare case, but it can happen.
		case "ERROR":
			code, ok := req.Error["error_code"]
			if !ok {
				return errors.New("error code not found in webhook")
			}

			link.LinkStatus = proto.LinkStatus_LINK_STATUS_ERROR
			link.ErrorCode = fmt.Sprintf("%v", code)
			link.ShouldBeUpdated = true
			log.Warn("link is in an error state, updating")
			if err := s.conn.UpdateLink(ctx, link); err != nil {
				return err
			}
		// Fired when an Item’s access consent is expiring in 7 days. Some Items have explicit expiration
		// times and we try to relay this when possible to reduce service disruption. This can be resolved
		// by having the user go through Link’s update mode.
		case "ITEM_LOGIN_REQUIRED":
			link.LinkStatus = proto.LinkStatus_LINK_STATUS_ITEM_LOGIN_REQUIRED
			link.ShouldBeUpdated = true
			log.Warn("link required a new item login")
			if err := s.conn.UpdateLink(ctx, link); err != nil {
				return err
			}
		// Fired when an Item’s access consent is expiring in 7 days. Some Items have explicit expiration
		// times and we try to relay this when possible to reduce service disruption. This can be resolved
		// by having the user go through Link’s update mode.
		case "PENDING_EXPIRATION":
			link.LinkStatus = proto.LinkStatus_LINK_STATUS_PENDING_EXPIRATION
			link.ExpirationDate = req.ConsentExpirationTime.String()
			link.ShouldBeUpdated = true
			log.Warn("link is pending expiration")
			if err := s.conn.UpdateLink(ctx, link); err != nil {
				return err
			}
		// The USER_PERMISSION_REVOKED webhook is fired when an end user has used either the my.plaid.com portal
		// or the financial institution’s consent portal to revoke the permission that they previously granted to
		//  access an Item. Once access to an Item has been revoked, it cannot be restored. If the user subsequently
		//  returns to your application, a new Item must be created for the user.
		case "USER_PERMISSION_REVOKED":
			code := req.Error["error_code"]
			link.LinkStatus = proto.LinkStatus_LINK_STATUS_REVOKED
			link.ErrorCode = fmt.Sprintf("%v", code)
			link.ShouldBeUpdated = true
			if err := s.conn.UpdateLink(ctx, link); err != nil {
				return err
			}
		// Fired when an Item's webhook is updated. This will be sent to the newly specified webhook.
		case "WEBHOOK_UPDATE_ACKNOWLEDGED":
			// DISPATCH BACKGROUND JOB TO PULL TRANSACTIONS FOR THE PAST 7 DAYS
			// time.Now().Add(-7 * 24 * time.Hour)
		// Fired when Plaid detects a new account for Items created or updated with Account Select v2.
		// Upon receiving this webhook, you can prompt your users to share new accounts with you
		// through Account Select v2 update mode.
		case "NEW_ACCOUNTS_AVAILABLE":
			link.PlaidNewAccountsAvailable = true
			// If using Account Select v2, you can use update mode to request your users to share new accounts with you.
			// Receiving a NEW_ACCOUNTS_AVAILABLE webhook indicates that Plaid has detected new accounts that you may want
			// to ask your users to share
			link.ShouldBeUpdated = true
			if err := s.conn.UpdateLink(ctx, link); err != nil {
				return err
			}
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	default:
		s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
	}

	return nil
}
