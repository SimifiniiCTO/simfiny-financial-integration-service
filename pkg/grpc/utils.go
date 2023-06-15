package grpc

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func (s *Server) UserHasActiveSubscription(ctx context.Context, userId *uint64) (bool, error) {
	var (
		conn = s.conn
	)

	if userId == nil {
		return false, fmt.Errorf("user id must be provided")
	}

	// check if the user has an active subscription
	profile, err := conn.GetUserProfileByUserID(ctx, *userId)
	if err != nil {
		return false, err
	}

	if profile == nil {
		return false, fmt.Errorf("user profile does not exist")
	}

	subscription := profile.StripeSubscriptions
	if subscription == nil {
		return false, fmt.Errorf("user profile does not have an active subscription")
	}

	if subscription.IsTrialing ||
		(subscription.StripeSubscriptionStatus == schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_TRIALING ||
			subscription.StripeSubscriptionStatus == schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_ACTIVE ||
			subscription.StripeSubscriptionStatus == schema.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_COMPLETE) {
		return true, nil
	}

	return false, fmt.Errorf("user profile does not have an active subscription")
}
