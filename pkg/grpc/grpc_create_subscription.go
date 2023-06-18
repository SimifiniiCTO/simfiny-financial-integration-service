package grpc

import (
	"context"
	"fmt"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/sub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSubscription(ctx context.Context, req *proto.CreateSubscriptionRequest) (*proto.CreateSubscriptionResponse, error) {
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

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-create-subscription")
		defer span.End()
	}

	// get the user profile
	profile, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// from the profile extract the stripe customer id
	stripeCustomerId := profile.StripeCustomerId
	// Automatically save the payment method to the subscription
	// when the first payment is successful.
	paymentSettings := &stripe.SubscriptionPaymentSettingsParams{
		SaveDefaultPaymentMethod: stripe.String("on_subscription"),
	}

	// Create the subscription. Note we're expanding the Subscription's
	// latest invoice and that invoice's payment_intent
	// so we can pass it to the front end to confirm the payment
	subscriptionParams := &stripe.SubscriptionParams{
		Customer: stripe.String(stripeCustomerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(req.PriceId),
			},
		},
		PaymentSettings: paymentSettings,
		PaymentBehavior: stripe.String("default_incomplete"),
	}
	subscriptionParams.AddExpand("latest_invoice.payment_intent")
	res, err := sub.New(subscriptionParams)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var subscriptionStatus proto.StripeSubscriptionStatus
	switch res.Status {
	case "active":
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE
	case "incomplete":
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE
	case "incomplete_expired":
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED
	case "past_due":
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_PAST_DUE
	case "trialing":
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_TRIALING
	case "unpaid":
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_UNPAID
	default:
		subscriptionStatus = proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED
	}

	subscription := &proto.StripeSubscription{
		StripeSubscriptionId:          res.ID,
		StripeSubscriptionStatus:      subscriptionStatus,
		StripeSubscriptionActiveUntil: fmt.Sprintf("%d", res.CurrentPeriodEnd),
		StripeWebhookLatestTimestamp:  "",
	}

	// update the user profile with the saved subscription details
	if err := s.conn.UpdateUserProfileSubscription(ctx, req.UserId, subscription); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// We return the client_secret from the subscription’s first payment intent to the frontend to complete payment.
	// NOTE: At this point the Subscription is inactive and awaiting payment. Here’s an example response.
	// The minimum fields to store are highlighted, but store whatever your application frequently accesses.
	return &proto.CreateSubscriptionResponse{
		SubscriptionId:            res.ID,
		PaymentIntentClientSecret: res.LatestInvoice.PaymentIntent.ClientSecret,
	}, nil
}
