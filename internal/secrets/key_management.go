package secrets

import (
	"context"
)

// KeyManagement is the interface that defines the methods for encrypting and decrypting data
type KeyManagement interface {
	Encrypt(ctx context.Context, input []byte) (keyID, version string, result []byte, _ error)
	Decrypt(ctx context.Context, keyID, version string, input []byte) (result []byte, _ error)
}
