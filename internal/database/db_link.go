package database

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gen/field"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
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
func (db *Db) GetLink(ctx context.Context, userID uint64, linkID uint64) (*schema.Link, error) {
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
	link, err := l.WithContext(ctx).Where(l.Id.Eq(linkID)).First()
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

	// decrypt the access token
	decrypted, err := db.decryptAccessToken(ctx, link.Token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt access token")
	}
	link.Token.AccessToken = *decrypted

	// convert to pb
	res, err := link.ToPB(ctx)
	if err != nil {
		return nil, err
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
func (db *Db) CreateLink(ctx context.Context, userID uint64, link *schema.Link) (*schema.Link, error) {
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

	// ensure the link does not exist for the given user
	l := db.QueryOperator.LinkORM
	if _, err := l.WithContext(ctx).Where(l.UserProfileId.Eq(userProfile.Id)).First(); err == nil {
		return nil, fmt.Errorf("link with id %d already exists", link.Id)
	}

	// TODO: fascilitate token encryption
	// encrypt the token prior to storage in the database
	// if err := db.encryptAccessToken(ctx, link.Token); err != nil {
	//	return nil, err
	// }

	// convert the link to orm type
	linkORM, err := link.ToORM(ctx)
	if err != nil {
		return nil, err
	}

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

	return &createdLink, nil
}

// encryptAccessToken encrypts the access token using the KMS
func (db *Db) encryptAccessToken(ctx context.Context, token *schema.Token) error {
	keyId, version, encrypted, err := db.Kms.Encrypt(ctx, []byte(token.AccessToken))
	if err != nil {
		return errors.Wrap(err, "failed to encrypt access token")
	}

	token.AccessToken = hex.EncodeToString(encrypted)
	token.KeyId = keyId
	token.Version = version
	return nil
}

// decryptAccessToken decrypts the access token using the KMS
func (db *Db) decryptAccessToken(ctx context.Context, token *schema.TokenORM) (*string, error) {
	decryptionKey := token.KeyId
	decryptionVersion := token.Version

	// decrypt the access token
	encrypted, err := hex.DecodeString(token.AccessToken)
	if err != nil {
		return nil, err
	}

	decrypted, err := db.Kms.Decrypt(ctx, decryptionKey, decryptionVersion, encrypted)
	if err != nil {
		return nil, err
	}

	accessToken := string(decrypted)
	return &accessToken, nil
}
