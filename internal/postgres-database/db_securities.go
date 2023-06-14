package database

import (
	"context"
	"fmt"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func (db *Db) UpdateSecurities(ctx context.Context, investmentAccountId *uint64, securities []*schema.InvestmentSecurity) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-investment-securities"); span != nil {
		defer span.End()
	}

	if investmentAccountId == nil {
		return fmt.Errorf("investment account ID is nil")
	}

	if securities == nil {
		return fmt.Errorf("securities is nil")
	}

	securitiesOrm := make([]*schema.InvestmentSecurityORM, 0, len(securities))
	// validate all present securities
	for _, security := range securities {
		if security.Id != 0 {
			return fmt.Errorf("security ID must not be non-zero")
		}

		// validate the investment security object
		if err := security.Validate(); err != nil {
			return err
		}

		securityOrm, err := security.ToORM(ctx)
		if err != nil {
			return err
		}

		securitiesOrm = append(securitiesOrm, &securityOrm)
	}

	// get the investment account by id
	ia := db.QueryOperator.InvestmentAccountORM
	is := db.QueryOperator.InvestmentSecurityORM
	investmentAcct, err := ia.WithContext(ctx).Where(ia.Id.Eq(*investmentAccountId)).First()
	if err != nil {
		return fmt.Errorf("investment account with id %d does not exist", *investmentAccountId)
	}

	securityIds := make([]uint64, 0, len(securities))
	for _, security := range securities {
		securityIds = append(securityIds, security.Id)
	}

	// delete all the existing securities
	if _, err := is.WithContext(ctx).Where(is.Id.In(securityIds...)).Delete(); err != nil {
		return err
	}

	// perform the bulk insert operation
	if err := ia.Securities.Model(investmentAcct).Append(securitiesOrm...); err != nil {
		return err
	}

	return nil
}
