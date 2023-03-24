package grpc

import (
	"context"
	"encoding/hex"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// EncryptAccessToken encrypts the access token using the KMS
func (s *Server) EncryptAccessToken(ctx context.Context, accessToken string) (*accessTokenMeta, error) {
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-encrypt-access-token")
		defer span.End()
	}

	keyId, version, encrypted, err := s.kms.Encrypt(ctx, []byte(accessToken))
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt access token")
	}

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				s.logger.Error(kms.ErrCodeNotFoundException, zap.Error(aerr))
			case kms.ErrCodeDisabledException:
				s.logger.Error(kms.ErrCodeDisabledException, zap.Error(aerr))
			case kms.ErrCodeKeyUnavailableException:
				s.logger.Error(kms.ErrCodeKeyUnavailableException, zap.Error(aerr))
			case kms.ErrCodeDependencyTimeoutException:
				s.logger.Error(kms.ErrCodeDependencyTimeoutException, zap.Error(aerr))
			case kms.ErrCodeInvalidKeyUsageException:
				s.logger.Error(kms.ErrCodeInvalidKeyUsageException, zap.Error(aerr))
			case kms.ErrCodeInvalidGrantTokenException:
				s.logger.Error(kms.ErrCodeInvalidGrantTokenException, zap.Error(aerr))
			case kms.ErrCodeInternalException:
				s.logger.Error(kms.ErrCodeInternalException, zap.Error(aerr))
			case kms.ErrCodeInvalidStateException:
				s.logger.Error(kms.ErrCodeInvalidStateException, zap.Error(aerr))
			default:
				s.logger.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			s.logger.Error(err.Error())
		}
		return nil, err
	}

	meta := &accessTokenMeta{}
	meta.accessToken = hex.EncodeToString(encrypted)
	meta.keyID = keyId
	meta.version = version
	return meta, nil
}

// DecryptUserAccessToken decrypts the access token using the KMS
func (s *Server) DecryptUserAccessToken(ctx context.Context, token *proto.Token) (*string, error) {
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-decrypt-access-token")
		defer span.End()
	}

	decryptionKey := token.KeyId
	decryptionVersion := token.Version

	// decrypt the access token
	encrypted, err := hex.DecodeString(token.AccessToken)
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
