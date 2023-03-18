package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AWSKMSConfig struct {
	Log       *zap.Logger
	KeyID     string
	Region    string
	SecretKey string
	Endpoint  *string
}

type AWSKMS struct {
	log    *zap.Logger
	config AWSKMSConfig
	client *kms.KMS
}

func NewAWSKMS(config AWSKMSConfig) (KeyManagement, error) {
	options := session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(config.KeyID, config.SecretKey, ""),
			Region:      aws.String(config.Region),
		},
	}

	if config.Endpoint != nil && *config.Endpoint != "" {
		options.Config.Endpoint = config.Endpoint
	}

	awsSession, err := session.NewSessionWithOptions(options)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws session")
	}

	client := kms.New(awsSession)
	return &AWSKMS{
		log:    config.Log,
		config: config,
		client: client,
	}, nil
}

func (a *AWSKMS) Encrypt(ctx context.Context, input []byte) (keyID string, version string, result []byte, _ error) {
	request := &kms.EncryptInput{
		EncryptionAlgorithm: aws.String("SYMMETRIC_DEFAULT"),
		EncryptionContext:   nil,
		GrantTokens:         []*string{},
		KeyId:               aws.String(a.config.KeyID),
		Plaintext:           input,
	}

	response, err := a.client.EncryptWithContext(ctx, request)
	if err != nil {
		return "", "", nil, errors.Wrap(err, "failed to encrypt data using AWS KMS")
	}

	return *response.KeyId, "", response.CiphertextBlob, nil
}

func (a *AWSKMS) Decrypt(ctx context.Context, keyID string, version string, input []byte) (result []byte, _ error) {

	request := &kms.DecryptInput{
		CiphertextBlob:      input,
		EncryptionAlgorithm: aws.String("SYMMETRIC_DEFAULT"), // TODO Maybe make this a config thing?
		EncryptionContext:   nil,
		GrantTokens:         []*string{},
		KeyId:               aws.String(keyID),
	}

	response, err := a.client.DecryptWithContext(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt data using AWS KMS")
	}

	return response.Plaintext, nil
}
