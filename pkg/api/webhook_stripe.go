package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/subscription"
	"github.com/stripe/stripe-go/v74/webhook"
)

type SubscriptionData struct {
	CustomerID string `json:"customer"`
}

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
	case "customer.created":
		// send when a customer is created
	case "customer.subscription.created":
		// 	Sent when the subscription is created. The subscription status might
		// be incomplete if customer authentication is required to complete the payment
		// or if you set payment_behavior to default_incomplete
		// ref: https://stripe.com/docs/billing/subscriptions/overview#subscription-payment-behavior
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Subscription created for %d.", subscription.ID)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handleSubscriptionCreated(subscription)
		handleSubscriptionCreated(&event)
	case "customer.subscription.deleted":
		// Sent when the subscription is deleted.
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Subscription deleted for %d.", subscription.ID)
		// Then define and call a func to handle the deleted subscription.
		// handleSubscriptionCanceled(subscription)

	case "customer.subscription.paused":
		// 	Sent when a subscription’s status changes to paused. For example,
		// this is sent when a subscription is configured to pause when a free trial
		// ends without a payment method. Invoicing won’t occur until the subscription is resumed.
		// We don’t send this event if payment collection is paused because invoices continue to be created during that time period.
		// ref: https://stripe.com/docs/billing/subscriptions/pause
		// ref: https://stripe.com/docs/api/subscriptions/resume
		// ref: https://stripe.com/docs/billing/subscriptions/trials#create-free-trials-without-payment
		// ref: https://stripe.com/docs/api/subscriptions/create#create_subscription-trial_settings-end_behavior-missing_payment_method
	case "customer.subscription.resumed":
		// Sent when a subscription’s status changes to active after it has been paused.
		// ref: https://stripe.com/docs/billing/subscriptions/pause#unpausing
	case "customer.subscription.trial_will_end":
		// Sent three days before the trial period ends. If the trial is less than three days, this event is triggered.
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Subscription trial will end for %d.", subscription.ID)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handleSubscriptionTrialWillEnd(subscription)
	case "customer.subscription.updated":
		// Sent when the subscription is successfully started, after the payment is confirmed. Also sent whenever a
		// subscription is changed. For example, adding a coupon, applying a discount, adding an invoice item,
		// and changing plans all trigger this event.
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Subscription updated for %d.", subscription.ID)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handleSubscriptionUpdated(subscription)
	case "invoice.created":
		// Sent when an invoice is created for a new or renewing subscription. If Stripe fails to receive a successful
		// response to invoice.created, then finalizing all invoices with automatic collection is delayed for up to 72 hours.
		//  Read more about finalizing invoices.
	case "invoice.finalized":
		// Sent when an invoice is successfully finalized and ready to be paid.
		// You can send the invoice to the customer. View invoice finalization to learn more.
		// Depending on your settings, we automatically charge the default payment method or attempt collection.
		// View emails after finalization to learn more.
		// ref: https://stripe.com/docs/invoicing/integration/workflow-transitions#finalized
		// ref: https://stripe.com/docs/invoicing/integration/workflow-transitions#emails
	case "invoice.finalization_failed":
		// The invoice couldn’t be finalized. Learn how to handle invoice finalization failures by reading the guide.
		// Learn more about invoice finalization in the invoices overview guide.
		// ref: https://stripe.com/docs/tax/customer-locations#finalizing-invoices-with-finalization-failures
		// ref: https://stripe.com/docs/invoicing/integration/workflow-transitions#finalized
		/*
			Inspect the Invoice’s last_finalization_error to determine the cause of the error.
			If you’re using Stripe Tax, check the Invoice object’s automatic_tax field.
			If automatic_tax[status]=requires_location_inputs, the invoice can’t be finalized and payments can’t be collected. Notify your customer and collect the required customer location.
			If automatic_tax[status]=failed, retry the request later.
		*/
	case "invoice.paid": // Sent each billing interval when a payment succeeds.
		// Continue to provision the subscription as payments continue to be made.
		// Store the status in your database and check when a user accesses your service.
		// This approach helps you avoid hitting rate limits.
		// Sent when the invoice is successfully paid. You can provision access to your product when you receive this event
		// and the subscription status is active.
		// TODO: just update the subscription status in the database
	case "invoice.payment_action_required":
		// Sent when the invoice requires customer authentication. Learn how to handle the subscription when the invoice requires action.
		// ref: https://stripe.com/docs/billing/subscriptions/overview#requires-action
	case "invoice.payment_failed": // Sent each billing interval if there is an issue with your customer’s payment method.
		// The payment failed or the customer does not have a valid payment method.
		// The subscription becomes past_due. Notify your customer and send them to the
		// customer portal to update their payment information.
		// A payment for an invoice failed. The PaymentIntent status changes to requires_action. The status of the subscription
		// continues to be incomplete only for the subscription’s first invoice. If a payment fails, there are several possible actions to take:
		// Notify the customer. Read about how you can configure subscription settings to enable Smart Retries and other revenue recovery features.
		// If you’re using PaymentIntents, collect new payment information and confirm the PaymentIntent.
		// Update the default payment method on the subscription.
	case "invoice.upcoming":
		// Sent a few days prior to the renewal of the subscription. The number of days is based on the number set for Upcoming renewal
		// events in the Dashboard. You can still add extra invoice items, if needed.
	case "invoice.updated":
		// Sent when a payment succeeds or fails. If payment is successful the paid attribute is set to true and the status is paid.
		// If payment fails, paid is set to false and the status remains open. Payment failures also trigger a invoice.payment_failed event.
	case "payment_intent.created":
		// Sent when a new PaymentIntent is created.
	case "charge.succeeded":
		// handleChargeSucceeded(payload)
	case "charge.failed":
		// handleChargeFailed(payload)
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Successful payment for %d.", paymentIntent.Amount)
		// Then define and call a func to handle the successful payment intent.
		// handlePaymentIntentSucceeded(paymentIntent) --> Update the subscription status in the database
	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handlePaymentMethodAttached(paymentMethod)
	case "checkout.session.completed": // Sent when a customer clicks the Pay or Subscribe button in Checkout, informing you of a new purchase.
		// Payment is successful and the subscription is created.
		// You should provision the subscription and save the customer ID to your database.
		// handleCheckoutSessionCompleted(payload)
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}

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
