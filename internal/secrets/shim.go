package secrets

import "context"

// ShimKMS is a shim implementation of the KeyManagement interface useful for testing
type ShimKMS struct {
}

var _ KeyManagement = &ShimKMS{}

// NewShimKMS creates a new ShimKMS implementation of the KeyManagement interface
func (s *ShimKMS) Decrypt(ctx context.Context, keyID, version string, input []byte) (result []byte, _ error) {
	data := []byte("sldjfdshfljd")
	return data, nil
}

// Encrypt encrypts the input using the AWS KMS client
func (s *ShimKMS) Encrypt(ctx context.Context, input []byte) (keyID, version string, result []byte, _ error) {
	result = []byte("sldjfdshfljd")
	keyID = "kshdkfjsdfsdfd"
	version = "1"
	return keyID, version, result, nil
}
