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
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go/v74"
	"go.uber.org/zap"
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

	// // Replace this endpoint secret with your endpoint's unique secret
	// // If you are testing with the CLI, find the secret by running 'stripe listen'
	// // If you are using an endpoint defined with the API or dashboard, look in your webhook settings
	// // at https://dashboard.stripe.com/webhooks
	// var endpointSecret string
	// if s.config.Environment == "dev" {
	// 	endpointSecret = "whsec_21441814697a4a51dc01395a030498131d56ec4d7155bb216cc75f36548c86bf"
	// } else {
	// 	endpointSecret = s.config.StripeEndpointSigningSecretKey
	// }

	// signatureHeader := req.Header.Get("Stripe-Signature")
	// event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "⚠️  Webhook signature verification failed. %v\n", err)
	// 	w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
	// 	return
	// }
	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	// subscribption events - ref: https://stripe.com/docs/billing/subscriptions/overview#subscription-statuses
	// Events are triggered every time a subscription is created or changed.
	// Some events are sent immediately when a subscription is created, while others
	//  recur on regular billing intervals. We recommend listening for events with a webhook endpoint.
	// Make sure that your integration properly handles the events. For example, you may want to email
	//  a customer if a payment fails or revoke a customer’s access when a subscription is canceled.
	case "customer.subscription.created":
		// Sent when the subscription is created. The subscription status might
		// be incomplete if customer authentication is required to complete the payment
		// or if you set payment_behavior to default_incomplete
		// ref: https://stripe.com/docs/billing/subscriptions/overview#subscription-payment-behavior
		if err := s.handleSubscription(ctx, event, schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE); err != nil {
			log.Println("Error handling deleted subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "customer.subscription.deleted":
		if err := s.handleSubscription(ctx, event, schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_CANCELED); err != nil {
			log.Println("Error handling deleted subscription:", err)
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
		if err := s.handleSubscription(ctx, event, schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_PAUSED); err != nil {
			log.Println("Error handling deleted subscription:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "customer.subscription.resumed":
		// Sent when a subscription’s status changes to active after it has been paused.
		// ref: https://stripe.com/docs/billing/subscriptions/pause#unpausing
		// TODO: Perform the below operations
		// 2. We send an email to the user informing them that their subscription has been resumed
		if err := s.handleSubscription(ctx, event, schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE); err != nil {
			log.Println("Error handling deleted subscription:", err)
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
		// Nothing to do here for now
	case "checkout.session.completed":
		// the expectation upon receipt of this webhook is that the customer's email and customer id
		// will be present in the data object returned
		// ref: https://stripe.com/docs/api/events/types#event_types-checkout.session.completed
		if err := s.handleCheckoutSessionCompleted(ctx, event); err != nil {
			log.Println("Error handling checkout session completed:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleCheckoutSessionCompleted(ctx context.Context, event stripe.Event) error {
	checkoutSession, err := s.extractCheckoutSessionFromEvent(event)
	if err != nil {
		return err
	}

	if checkoutSession.Customer == nil || checkoutSession.Subscription == nil {
		log.Println("Error retrieving customer or subscription from checkout session")
		return errors.New("error retrieving customer or subscription from checkout session")
	}

	if err := s.updateCustomerSubscription(ctx, checkoutSession.Customer, checkoutSession.Subscription); err != nil {
		log.Println("Error updating customer subscription:", err)
		return err
	}

	return nil
}

// process delete subscription request
func (s *Server) handleSubscription(ctx context.Context, event stripe.Event, status schema.StripeSubscriptionStatus) error {
	subscription, err := s.getSubscriptionFromEvent(event)
	if err != nil {
		return err
	}

	customer := subscription.Customer

	if err := s.updateCustomerSubscription(ctx, customer, subscription); err != nil {
		log.Println("Error updating customer subscription:", err)
		return err
	}
	return nil
}

func (s *Server) updateCustomerSubscription(ctx context.Context, customer *stripe.Customer, subscription *stripe.Subscription) error {
	customerID := customer.ID
	customerEmail := customer.Email
	stripeNewSubscription := subscription
	if customerEmail == "" {
		log.Println("Error retrieving customer email from checkout session")
		return errors.New("error retrieving customer email from checkout session")
	}

	userProfile, err := s.conn.GetUserProfileByEmail(ctx, customerEmail)
	if err != nil {
		log.Println("Error retrieving user profile:", err)
		return nil
	}

	isTrialing := false
	if stripeNewSubscription.Status == "trialing" {
		isTrialing = true
	}

	t := time.Unix(stripeNewSubscription.TrialEnd, 0)

	activeUntil := t.Format("2006-01-02 15:04:05")

	currentSubscription := userProfile.StripeSubscriptions
	if currentSubscription == nil {
		currentSubscription = &schema.StripeSubscription{
			StripeSubscriptionId:          stripeNewSubscription.ID,
			StripeSubscriptionStatus:      schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE,
			StripeSubscriptionActiveUntil: activeUntil,
			StripeWebhookLatestTimestamp:  time.Now().UTC().Format(time.RFC3339),
			IsTrialing:                    isTrialing,
		}

		// associate the customer id with the user profile since
		// we can assume this is a new subscription being created
		userProfile.StripeCustomerId = customerID
	} else {
		currentSubscription.StripeSubscriptionId = stripeNewSubscription.ID
		currentSubscription.StripeSubscriptionStatus = schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE
		currentSubscription.StripeSubscriptionActiveUntil = activeUntil
		currentSubscription.StripeWebhookLatestTimestamp = time.Now().UTC().Format(time.RFC3339)
		currentSubscription.IsTrialing = isTrialing
	}

	userProfile.StripeSubscriptions = currentSubscription

	if err := s.conn.UpdateUserProfile(ctx, userProfile); err != nil {
		log.Println("Error updating user profile:", err)
		return err
	}

	return nil
}

func (s *Server) extractCheckoutSessionFromEvent(event stripe.Event) (*stripe.CheckoutSession, error) {
	var checkoutSession stripe.CheckoutSession
	s.logger.Info("checkout session completed", zap.Any("checkoutSession", checkoutSession))
	err := json.Unmarshal(event.Data.Raw, &checkoutSession)
	if err != nil {
		return nil, err
	}

	if checkoutSession.Customer == nil || checkoutSession.Subscription == nil {
		log.Println("Error retrieving customer or subscription from checkout session")
		return nil, err
	}

	return &checkoutSession, nil
}

func (s *Server) getSubscriptionFromEvent(event stripe.Event) (*stripe.Subscription, error) {
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}
