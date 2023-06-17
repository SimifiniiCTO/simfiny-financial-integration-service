package encryptdecrypt

import (
	"context"
	"encoding/hex"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AccessTokenMeta struct {
	AccessToken string
	KeyID       string
	Version     string
}

// EncryptAccessToken encrypts the access token using the KMS
func EncryptAccessToken(ctx context.Context, accessToken string, kmsClient secrets.KeyManagement, logger *zap.Logger) (*AccessTokenMeta, error) {
	keyId, version, encrypted, err := kmsClient.Encrypt(ctx, []byte(accessToken))
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt access token")
	}

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				logger.Error(kms.ErrCodeNotFoundException, zap.Error(aerr))
			case kms.ErrCodeDisabledException:
				logger.Error(kms.ErrCodeDisabledException, zap.Error(aerr))
			case kms.ErrCodeKeyUnavailableException:
				logger.Error(kms.ErrCodeKeyUnavailableException, zap.Error(aerr))
			case kms.ErrCodeDependencyTimeoutException:
				logger.Error(kms.ErrCodeDependencyTimeoutException, zap.Error(aerr))
			case kms.ErrCodeInvalidKeyUsageException:
				logger.Error(kms.ErrCodeInvalidKeyUsageException, zap.Error(aerr))
			case kms.ErrCodeInvalidGrantTokenException:
				logger.Error(kms.ErrCodeInvalidGrantTokenException, zap.Error(aerr))
			case kms.ErrCodeInternalException:
				logger.Error(kms.ErrCodeInternalException, zap.Error(aerr))
			case kms.ErrCodeInvalidStateException:
				logger.Error(kms.ErrCodeInvalidStateException, zap.Error(aerr))
			default:
				logger.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			logger.Error(err.Error())
		}
		return nil, err
	}

	meta := &AccessTokenMeta{}
	meta.AccessToken = hex.EncodeToString(encrypted)
	meta.KeyID = keyId
	meta.Version = version
	return meta, nil
}

// DecryptUserAccessToken decrypts the access token using the KMS
func DecryptUserAccessToken(ctx context.Context, token *proto.Token, kmsClient secrets.KeyManagement, logger *zap.Logger) (*string, error) {

	decryptionKey := token.KeyId
	decryptionVersion := token.Version

	// decrypt the access token
	encrypted, err := hex.DecodeString(token.AccessToken)
	if err != nil {
		return nil, err
	}

	decrypted, err := kmsClient.Decrypt(ctx, decryptionKey, decryptionVersion, encrypted)
	if err != nil {
		return nil, err
	}

	accessToken := string(decrypted)
	return &accessToken, nil
}
