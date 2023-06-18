package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func (db *Db) GetSubscriptionBySubscriptionId(ctx context.Context, subscriptionId *string) (*schema.StripeSubscription, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-subscription-by-subscritionid"); span != nil {
		defer span.End()
	}

	// validate the request
	if subscriptionId == nil {
		return nil, fmt.Errorf("subscription id is nil")
	}

	// ensure the subscription exists
	s := db.QueryOperator.StripeSubscriptionORM
	subscription, err := s.WithContext(ctx).Where(s.StripeSubscriptionId.Eq(*subscriptionId)).First()
	if err != nil {
		return nil, fmt.Errorf("subscription with id %s does not exist", *subscriptionId)
	}

	// convert to pb type
	result, err := subscription.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *Db) UpdateSubscription(ctx context.Context, stripeSubscriptionId *string, subscription *schema.StripeSubscription) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-subscription"); span != nil {
		defer span.End()
	}

	// validate the request
	if stripeSubscriptionId == nil {
		return fmt.Errorf("stripe subscription id is nil")
	}

	// ensure the subscription exists
	s := db.QueryOperator.StripeSubscriptionORM
	if _, err := s.WithContext(ctx).Where(s.StripeSubscriptionId.Eq(*stripeSubscriptionId)).First(); err != nil {
		return fmt.Errorf("subscription with id %s does not exist", *stripeSubscriptionId)
	}

	// update the subscription
	if _, err := s.WithContext(ctx).Where(s.StripeSubscriptionId.Eq(*stripeSubscriptionId)).Updates(subscription); err != nil {
		return err
	}

	return nil
}
