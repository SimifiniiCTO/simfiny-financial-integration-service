package database

import (
	"context"
	"fmt"

	"gorm.io/gen/field"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// LinkExistsForItem returns true if a link exists for the given item and user id
func (db *Db) LinkExistsForItem(ctx context.Context, userID uint64, itemID string) (bool, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-link-exists-for-item"); span != nil {
		defer span.End()
	}

	// ensure the user exists
	u := db.QueryOperator.UserProfileORM
	if _, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).First(); err != nil {
		return false, fmt.Errorf("user with id %d does not exist", userID)
	}

	pl := db.QueryOperator.PlaidLinkORM
	plaidLink, err := pl.WithContext(ctx).Where(pl.ItemId.Eq(itemID)).First()
	if err != nil {
		return false, fmt.Errorf("plaid link with item id %s does not exist", itemID)
	}

	// ensure the plaid link belonds to the user
	l := db.QueryOperator.LinkORM
	if _, err := l.WithContext(ctx).Where(l.Id.Eq(*plaidLink.LinkId), l.UserProfileId.Eq(userID)).First(); err != nil {
		return false, fmt.Errorf("link with id %d and user id %d does not exist", plaidLink.LinkId, userID)
	}

	return true, nil
}

// GetLink takes as input the userID and linkID of interest and returns the associated link
// if the link exists and is owned by the user. As part of this operation, we ensure all the associations
// a link returns are also populated
func (db *Db) GetLink(ctx context.Context, userID uint64, linkID uint64, clearAccessToken bool) (*schema.Link, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-link"); span != nil {
		defer span.End()
	}

	// ensure the user exists
	u := db.QueryOperator.UserProfileORM
	userProfile, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).First()
	if err != nil {
		return nil, fmt.Errorf("user with id %d does not exist", userID)
	}

	// ensure the link exists
	l := db.QueryOperator.LinkORM
	link, err := l.
		WithContext(ctx).
		Where(l.Id.Eq(linkID)).
		Preload(l.BankAccounts.Pockets.Goals.Forecasts).
		Preload(l.BankAccounts.Pockets.Goals.Milestones).
		Preload(l.CreditAccounts.Aprs).
		Preload(l.InvestmentAccounts.Holdings).
		Preload(l.InvestmentAccounts.Securities).
		Preload(l.MortgageAccounts).
		Preload(l.PlaidLink).
		Preload(l.StudentLoanAccounts).
		Preload(l.Token).
		Preload(l.PlaidSync).
		First()
	if err != nil {
		return nil, fmt.Errorf("link with id %d does not exist", linkID)
	}

	if link.UserProfileId == nil {
		return nil, fmt.Errorf("link with id %d does not belong to any user", linkID)
	}

	// ensure the link belongs to the user
	if *link.UserProfileId != userProfile.Id {
		return nil, fmt.Errorf("link with id %d does not belong to user %d", linkID, userID)
	}

	// convert to pb
	res, err := link.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	if clearAccessToken {
		res.Token = nil
	}

	// return the link
	return &res, nil
}

// DeleteLink takes as input the userID and linkID of interest and deletes the associated link and
// all the associations it has if the link exists and is owned by the user
func (db *Db) DeleteLink(ctx context.Context, userID uint64, linkID uint64) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-delete-link"); span != nil {
		defer span.End()
	}

	// ensure the user exists
	u := db.QueryOperator.UserProfileORM
	userProfile, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).First()
	if err != nil {
		return fmt.Errorf("user with id %d does not exist", userID)
	}

	// ensure the link exists
	l := db.QueryOperator.LinkORM
	link, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).First()
	if err != nil {
		return fmt.Errorf("link with id %d does not exist", linkID)
	}

	if link.UserProfileId == nil {
		return fmt.Errorf("link with id %d does not belong to any user", linkID)
	}

	// ensure the link belongs to the user
	if *link.UserProfileId != userProfile.Id {
		return fmt.Errorf("link with id %d does not belong to user with id %d", linkID, userID)
	}

	// delete the link and all associations
	result, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).
		Select(field.AssociationFields).
		Delete()
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

// CreateLink takes as input the userID and link of interest and creates the associated link
func (db *Db) CreateLink(ctx context.Context, userID uint64, link *schema.Link, clearAccessToken bool) (*schema.Link, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-create-link"); span != nil {
		defer span.End()
	}

	// validate the link token
	if err := link.ValidateAll(); err != nil {
		return nil, err
	}

	// the token subfield must exist for a given link
	if link.Token == nil {
		return nil, fmt.Errorf("token must be provided")
	}

	// ensure the user exists
	u := db.QueryOperator.UserProfileORM
	userProfile, err := u.WithContext(ctx).Where(u.UserId.Eq(userID)).First()
	if err != nil {
		return nil, fmt.Errorf("user with id %d does not exist", userID)
	}

	// ensure the link does not exist for the given user with the defined institution name
	l := db.QueryOperator.LinkORM
	if _, err := l.WithContext(ctx).Where(l.InstitutionName.Eq(link.InstitutionName), l.UserProfileId.Eq(userProfile.Id)).First(); err == nil {
		return nil, fmt.Errorf("link with user profile id %d and instutition name %s already exists", userProfile.Id, link.InstitutionName)
	}

	// convert the link to orm type
	linkORM, err := link.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	// update link status
	linkORM.LinkStatus = schema.LinkStatus_LINK_STATUS_SUCCESS.String()

	// associate the link to the user profile
	if err := u.Link.Model(userProfile).Append(&linkORM); err != nil {
		return nil, err
	}

	// perform the create operation
	result, err := u.WithContext(ctx).Updates(userProfile)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	// the linkORM is now populated with the id
	createdLink, err := linkORM.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	if clearAccessToken {
		createdLink.Token = nil
	}

	return &createdLink, nil
}

// This function takes a context and an item ID as input and returns the associated PlaidLink object if
// it exists. It first starts a datastore span to instrument the operation. It then queries the
// PlaidLinkORM to find the link with the given item ID. If the link exists, it converts it to a
// protobuf object and returns it. If the link does not exist, it returns an error.
func (db *Db) GetLinkByItemId(ctx context.Context, itemId string, clearAccessToken bool) (*schema.Link, error) {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-get-link"); span != nil {
		defer span.End()
	}

	// ensure the link exists
	pl := db.QueryOperator.PlaidLinkORM
	l := db.QueryOperator.LinkORM
	plaidLink, err := pl.WithContext(ctx).Where(pl.ItemId.Eq(itemId)).First()
	if err != nil {
		return nil, fmt.Errorf("link with id %s does not exist", itemId)
	}

	linkId := plaidLink.LinkId
	link, err := l.
		WithContext(ctx).
		Where(l.Id.Eq(*linkId)).
		Preload(l.BankAccounts.Pockets.Goals.Forecasts).
		Preload(l.BankAccounts.Pockets.Goals.Milestones).
		Preload(l.CreditAccounts.Aprs).
		Preload(l.InvestmentAccounts.Holdings).
		Preload(l.InvestmentAccounts.Securities).
		Preload(l.MortgageAccounts).
		Preload(l.PlaidLink).
		Preload(l.StudentLoanAccounts).
		Preload(l.Token).
		Preload(l.PlaidSync).
		First()
	if err != nil {
		return nil, fmt.Errorf("link with id %d does not exist", *linkId)
	}

	// convert to pb
	res, err := link.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	if clearAccessToken {
		res.Token = nil
	}

	// return the link
	return &res, nil
}

func (db *Db) UpdateLink(ctx context.Context, link *schema.Link) error {
	// instrument operation
	if span := db.startDatastoreSpan(ctx, "dbtxn-update-pocket"); span != nil {
		defer span.End()
	}

	if link == nil {
		return fmt.Errorf("link must be provided")
	}

	if link.Id == 0 {
		return fmt.Errorf("link id must be provided. got: %d", link.Id)
	}

	l := db.QueryOperator.LinkORM
	if _, err := l.WithContext(ctx).Where(l.Id.Eq(link.Id)).First(); err != nil {
		return fmt.Errorf("link with id %d does not exist", link.Id)
	}

	// update the link
	linkORM, err := link.ToORM(ctx)
	if err != nil {
		return err
	}

	// perform the update operation
	result, err := l.WithContext(ctx).Updates(&linkORM)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
