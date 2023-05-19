package grpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type PlaidClaims struct {
	jwt.StandardClaims
}

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

// ProcessWebhook processes a webhook from Plaid
func (s *Server) ProcessWebhook(ctx context.Context, req *proto.ProcessWebhookRequest) (*proto.ProcessWebhookResponse, error) {
	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-process-webhook-request")
		defer span.End()
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := status.Errorf(codes.Unauthenticated, "missing metadata")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	verification := md.Get("Plaid-Verification")[0]
	if strings.TrimSpace(verification) == "" {
		err := status.Errorf(codes.Unauthenticated, "unauthorized")
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
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

		// exchange the kid for a public key and try to obtain the verification key function
		keyFunction, err := s.InMemoryWebhookVerification.GetVerificationKey(ctx, kid)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get verification key for webhook")
		}

		return keyFunction.Keyfunc(token)
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if !result.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	if err := s.processWebhook(ctx, req); err != nil {
		return nil, err
	}

	return &proto.ProcessWebhookResponse{}, nil
}

func (s *Server) processWebhook(ctx context.Context, req *proto.ProcessWebhookRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "missing request")
	}

	if err := req.ValidateAll(); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
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
	link, err := s.conn.GetLinkByItemId(ctx, req.GetItemId())
	if err != nil {
		return err
	}

	switch req.WebhookType {
	case "INVESTMENTS_TRANSACTIONS":
		switch req.WebhookCode {
		// Fired when new or updated investment transactions have been detected on an investment account.
		// The webhook typically fires in response to any newly added transactions or price changes to existing
		// transactions, most commonly after market close.
		case "DEFAULT_UPDATE":
			// Trigger a background job to sync the plaid transactions
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	case "HOLDINGS":
		switch req.WebhookCode {
		// Fired when new or updated holdings have been detected on an investment account.
		// The webhook typically fires in response to any newly added holdings or price changes to existing
		// holdings, most commonly after market close.
		case "DEFAULT_UPDATE":
			// Trigger a background job to sync the plaid holdings
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	case "LIABILITIES":
		switch req.WebhookCode {
		// Liabilities webhooks are sent to indicate that new loans or updated loan fields for existing accounts are available.
		// will be fired when new or updated liabilities have been detected on a liabilities item
		case "DEFAULT_UPDATE":
			// Trigger a background job to sync the plaid liabilities
		default:
			s.logger.Error("Plaid webhook will not be handled, it is not implemented.")
		}
	case "TRANSACTIONS":
		switch link.PlaidLink.UsePlaidSync {
		case true:
			// Trigger a background job to sync the plaid transactions
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
				// if initial update field is present ... pull the past 7 days of tx
				// if historical update field is present ... pull the past 2 years of tx
				// trigger a backgrund job to pull the past 7 days of plaid transactions
				// start: time.Now().Add(-7 * 24 * time.Hour) | time.Now().Add(-2 * 365 * 24 * time.Hour)
				// end: time.Now()
			// Fired when transaction(s) for an Item are deleted. The deleted transaction IDs are
			// included in the webhook payload. Plaid will typically check for deleted transaction data several times a day.
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
			link.ErrorCode = code.String()
			log.Warn("link is in an error state, updating")
			if err := s.conn.UpdateLink(ctx, link); err != nil {
				return err
			}
		// Fired when an Item’s access consent is expiring in 7 days. Some Items have explicit expiration
		// times and we try to relay this when possible to reduce service disruption. This can be resolved
		// by having the user go through Link’s update mode.
		case "PENDING_EXPIRATION":
			link.LinkStatus = proto.LinkStatus_LINK_STATUS_PENDING_EXPIRATION
			link.ExpirationDate = req.ConsentExpirationTime
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
			link.ErrorCode = code.String()
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
