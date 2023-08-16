package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhook"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) StripeWebhook(ctx context.Context, req *schema.StripeWebhookRequest) (*schema.StripeWebhookResponse, error) {
	event := stripe.Event{}

	// Parse the request body into the event struct
	payload := req.GetBody()
	bytes, err := payload.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &event); err != nil {
		return nil, err
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
	return nil, status.Errorf(codes.Unimplemented, "method StripeWebhook not implemented")
}
