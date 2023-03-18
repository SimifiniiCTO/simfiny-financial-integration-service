package grpc

import (
	"context"
	"encoding/hex"

	"github.com/pkg/errors"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// EncryptAccessToken encrypts the access token using the KMS
func (s *Server) EncryptAccessToken(ctx context.Context, accessToken string) (*accessTokenMeta, error) {
	keyId, version, encrypted, err := s.kms.Encrypt(ctx, []byte(accessToken))
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt access token")
	}

	meta := &accessTokenMeta{}
	meta.accessToken = hex.EncodeToString(encrypted)
	meta.keyID = keyId
	meta.version = version
	return meta, nil
}

// DecryptUserAccessToken decrypts the access token using the KMS
func (s *Server) DecryptUserAccessToken(ctx context.Context, userProfile *proto.UserProfile) (*string, error) {
	decryptionKey := userProfile.DecryptionAccessTokenKey
	decryptionVersion := userProfile.DecryptionAccessTokenVersion
	// decrypt the access token
	encrypted, err := hex.DecodeString(userProfile.PlaidAccessToken)
	if err != nil {
		return nil, err
	}

	decrypted, err := s.kms.Decrypt(ctx, decryptionKey, decryptionVersion, encrypted)
	if err != nil {
		return nil, err
	}

	accessToken := string(decrypted)
	return &accessToken, nil
}
