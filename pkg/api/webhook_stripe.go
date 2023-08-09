package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/subscription"
	"github.com/stripe/stripe-go/v74/webhook"
)

// The type SubscriptionData contains a single field for customer ID represented as a string in JSON
// format.
// @property {string} CustomerID - The CustomerID property is a string that represents the unique
// identifier of a customer who has subscribed to a service or product. It is tagged with
// `json:"customer"` which indicates that when this struct is serialized to JSON, the property name
// will be "customer".
type SubscriptionData struct {
	CustomerID string `json:"customer"`
}

// handleStripeWebhook handles Stripe webhook events. It reads the request
// body, verifies the signature of the event, and processes the event based on its type. The function
// updates the subscription status in the database based on the event type, and sends emails to the
// user informing them of changes to their subscription status. The function also returns appropriate
// HTTP status codes based on the success or failure of the event processing.
func (s *Server) handleStripeWebhook(w http.ResponseWriter, req *http.Request) {
	// ref: https://github.com/kazamori/stripe-webhook-sample/blob/main/handler/webhook.go

	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	ctx := req.Context()
	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook error while parsing basic request. %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Replace this endpoint secret with your endpoint's unique secret
	// If you are testing with the CLI, find the secret by running 'stripe listen'
	// If you are using an endpoint defined with the API or dashboard, look in your webhook settings
	// at https://dashboard.stripe.com/webhooks
	var endpointSecret string
	if s.config.Environment == "dev" {
		endpointSecret = "whsec_21441814697a4a51dc01395a030498131d56ec4d7155bb216cc75f36548c86bf"
	} else {
		endpointSecret = s.config.StripeEndpointSigningSecretKey
	}

	signatureHeader := req.Header.Get("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook signature verification failed. %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}
	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	// subscribption events - ref: https://stripe.com/docs/billing/subscriptions/overview#subscription-statuses
	// Events are triggered every time a subscription is created or changed.
	// Some events are sent immediately when a subscription is created, while others
	//  recur on regular billing intervals. We recommend listening for events with a webhook endpoint.
	// Make sure that your integration properly handles the events. For example, you may want to email
	//  a customer if a payment fails or revoke a customer’s access when a subscription is canceled.
	case "customer.subscription.created":
		// 	Sent when the subscription is created. The subscription status might
		// be incomplete if customer authentication is required to complete the payment
		// or if you set payment_behavior to default_incomplete
		// ref: https://stripe.com/docs/billing/subscriptions/overview#subscription-payment-behavior
		subscription, storedSubscription, shouldReturn := s.processSubscriptionEventAndQueryStoreSubscription(event, w, ctx)
		if shouldReturn {
			return
		}

		// update the subscription status in the database
		storedSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_CREATED
		if err := s.conn.UpdateSubscription(ctx, &subscription.ID, storedSubscription); err != nil {
			log.Println("Error updating subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "customer.subscription.deleted":
		// Sent when the subscription is deleted.
		subscription, storedSubscription, shouldReturn := s.processSubscriptionEventAndQueryStoreSubscription(event, w, ctx)
		if shouldReturn {
			return
		}

		// update the subscription status in the database
		storedSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_CANCELED
		if err := s.conn.UpdateSubscription(ctx, &subscription.ID, storedSubscription); err != nil {
			log.Println("Error updating subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "customer.subscription.paused":
		// 	Sent when a subscription’s status changes to paused. For example,
		// this is sent when a subscription is configured to pause when a free trial
		// ends without a payment method. Invoicing won’t occur until the subscription is resumed.
		// We don’t send this event if payment collection is paused because invoices continue to be created during that time period.
		// ref: https://stripe.com/docs/billing/subscriptions/pause
		// ref: https://stripe.com/docs/api/subscriptions/resume
		// ref: https://stripe.com/docs/billing/subscriptions/trials#create-free-trials-without-payment
		// ref: https://stripe.com/docs/api/subscriptions/create#create_subscription-trial_settings-end_behavior-missing_payment_method

		// TODO: Perform the below operations
		// 2. We send an email to the user informing them that their subscription has been paused
		subscription, storedSubscription, shouldReturn := s.processSubscriptionEventAndQueryStoreSubscription(event, w, ctx)
		if shouldReturn {
			return
		}

		// update the subscription status in the database
		storedSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_PAUSED
		if err := s.conn.UpdateSubscription(ctx, &subscription.ID, storedSubscription); err != nil {
			log.Println("Error updating subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "customer.subscription.resumed":
		// Sent when a subscription’s status changes to active after it has been paused.
		// ref: https://stripe.com/docs/billing/subscriptions/pause#unpausing
		// TODO: Perform the below operations
		// 2. We send an email to the user informing them that their subscription has been resumed
		subscription, storedSubscription, shouldReturn := s.processSubscriptionEventAndQueryStoreSubscription(event, w, ctx)
		if shouldReturn {
			return
		}

		// update the subscription status in the database
		storedSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE
		if err := s.conn.UpdateSubscription(ctx, &subscription.ID, storedSubscription); err != nil {
			log.Println("Error updating subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "customer.subscription.trial_will_end":
		// Sent three days before the trial period ends. If the trial is less than three days, this event is triggered.
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: 2. We send an email to the user informing them that their subscription has been cancelled
	case "customer.subscription.updated":
		// Sent when the subscription is successfully started, after the payment is confirmed. Also sent whenever a
		// subscription is changed. For example, adding a coupon, applying a discount, adding an invoice item,
		// and changing plans all trigger this event.
	case "invoice.paid": // Sent each billing interval when a payment succeeds.
		// Continue to provision the subscription as payments continue to be made.
		// Store the status in your database and check when a user accesses your service.
		// This approach helps you avoid hitting rate limits.
		// Sent when the invoice is successfully paid. You can provision access to your product when you receive this event
		// and the subscription status is active.
		// TODO: just update the subscription status in the database
		subscription, storedSubscription, shouldReturn := s.processSubscriptionEventAndQueryStoreSubscription(event, w, ctx)
		if shouldReturn {
			return
		}

		// update the subscription status in the database
		storedSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE
		if err := s.conn.UpdateSubscription(ctx, &subscription.ID, storedSubscription); err != nil {
			log.Println("Error updating subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "invoice.payment_failed": // Sent each billing interval if there is an issue with your customer’s payment method.
		// The payment failed or the customer does not have a valid payment method.
		// The subscription becomes past_due. Notify your customer and send them to the
		// customer portal to update their payment information.
		// A payment for an invoice failed. The PaymentIntent status changes to requires_action. The status of the subscription
		// continues to be incomplete only for the subscription’s first invoice. If a payment fails, there are several possible actions to take:
		// Notify the customer. Read about how you can configure subscription settings to enable Smart Retries and other revenue recovery features.
		// If you’re using PaymentIntents, collect new payment information and confirm the PaymentIntent.
		// Update the default payment method on the subscription.
		subscription, storedSubscription, shouldReturn := s.processSubscriptionEventAndQueryStoreSubscription(event, w, ctx)
		if shouldReturn {
			return
		}

		// update the subscription status in the database
		storedSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE
		if err := s.conn.UpdateSubscription(ctx, &subscription.ID, storedSubscription); err != nil {
			log.Println("Error updating subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "invoice.upcoming":
		// Sent a few days prior to the renewal of the subscription. The number of days is based on the number set for Upcoming renewal
		// events in the Dashboard. You can still add extra invoice items, if needed.
		// TODO: send an email to the user informing them that their subscription is about to be renewed
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Successful payment for %d.", paymentIntent.Amount)
	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "checkout.session.completed":
		// the expectation upon receipt of this webhook is that the customer's email and customer id
		// will be present in the data object returned
		// ref: https://stripe.com/docs/api/events/types#event_types-checkout.session.completed
		var checkoutSession stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &checkoutSession)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// extract the customer id from checkout session object as well as the customer email address
		customerID := checkoutSession.Customer.ID
		stripeNewSubscription := checkoutSession.Subscription
		userProfile, err := s.conn.GetUserProfileByCustomerId(ctx, customerID)
		if err != nil {
			log.Println("Error retrieving user profile:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// update the subscription object
		isTrialing := false
		if stripeNewSubscription.Status == "trialing" {
			isTrialing = true
		}

		// Convert the int64 timestamp to a time.Time value
		t := time.Unix(stripeNewSubscription.TrialEnd, 0)

		// Format the time.Time value as a string
		activeUntil := t.Format("2006-01-02 15:04:05")

		// update current subscription
		currentSubscription := userProfile.StripeSubscriptions
		if currentSubscription == nil {
			currentSubscription = &schema.StripeSubscription{
				StripeSubscriptionId:          stripeNewSubscription.ID,
				StripeSubscriptionStatus:      schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE,
				StripeSubscriptionActiveUntil: activeUntil,
				StripeWebhookLatestTimestamp:  time.Now().UTC().Format(time.RFC3339),
				IsTrialing:                    isTrialing,
			}
		} else {
			currentSubscription.StripeSubscriptionId = stripeNewSubscription.ID
			currentSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE
			currentSubscription.StripeSubscriptionActiveUntil = activeUntil
			currentSubscription.StripeWebhookLatestTimestamp = time.Now().UTC().Format(time.RFC3339)
			currentSubscription.IsTrialing = isTrialing
		}

		// update the user profile
		userProfile.StripeSubscriptions = currentSubscription
		if err := s.conn.UpdateUserProfile(ctx, userProfile); err != nil {
			log.Println("Error updating user profile:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}

// The processSubscriptionEventAndQueryStoreSubscription function  processes a Stripe subscription event and queries
// the database to retrieve the corresponding stored subscription. It takes in the Stripe event, a HTTP
// response writer, and a context as parameters. It first unmarshals the event data into a Stripe
// subscription object, and then queries the database using the subscription ID to retrieve the stored
// subscription. If there is an error in either of these steps, it returns an error response. The
// function returns the Stripe subscription object, the stored subscription object, and a boolean
// indicating whether there was an error.
func (s *Server) processSubscriptionEventAndQueryStoreSubscription(event stripe.Event, w http.ResponseWriter, ctx context.Context) (stripe.Subscription, *schema.StripeSubscription, bool) {
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return stripe.Subscription{}, nil, true
	}

	storedSubscription, err := s.conn.GetSubscriptionBySubscriptionId(ctx, &subscription.ID)
	if err != nil {
		log.Println("Error retrieving subscription:", err)
		w.WriteHeader(http.StatusBadRequest)
		return stripe.Subscription{}, nil, true
	}
	return subscription, storedSubscription, false
}

// handleSubscriptionCreated handles a Stripe event for a new subscription creation and retrieves the latest
// subscription for a customer.
func handleSubscriptionCreated(payload *stripe.Event) {
	var data SubscriptionData
	err := json.Unmarshal(payload.Data.Raw, &data)
	if err != nil {
		log.Println("Error unmarshaling subscription data:", err)
		return
	}

	customerID := data.CustomerID
	subscription, err := retrieveLatestSubscription(customerID)
	if err != nil {
		log.Println("Error retrieving subscription:", err)
		return
	}

	// Handle the new subscription
	fmt.Println("New subscription created for customer:", customerID)
	fmt.Println("Subscription ID:", subscription.ID)
}

// retrieveLatestSubscription retrieves the latest active subscription for a given customer ID using the Stripe API
// in Go.
func retrieveLatestSubscription(customerID string) (*stripe.Subscription, error) {
	activeSubscription := string(stripe.SubscriptionStatusActive)
	subscriptionListParams := &stripe.SubscriptionListParams{
		Customer: &customerID,
		Status:   &activeSubscription,
	}

	iter := subscription.List(subscriptionListParams)
	for iter.Next() {
		subscription := iter.Subscription()
		if subscription.Status == "active" {
			return subscription, nil
		}
	}

	if iter.Err() != nil {
		return nil, iter.Err()
	}

	return nil, fmt.Errorf("no active subscriptions found for customer: %s", customerID)
}
