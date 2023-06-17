package plaidhandler

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/pkg/errors"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
)

// The WebhookVerificationKey type defines the properties of a JSON Web Key used for webhook
// verification.
// @property {string} Alg - identifies the cryptographic algorithm family used with the key.
// @property {string} Crv - The cryptographic curve used with the key. This is typically used in
// elliptic curve cryptography (ECC) to define the curve used for generating the key pair. Examples of
// curves include P-256 and P-384.
// @property {string} Kid - The Key ID parameter is used to match a specific key. It can be used to
// choose among a set of keys within the JWK during key rollover.
// @property {string} Kty - The "kty" property stands for "key type" and identifies the cryptographic
// algorithm family used with the key, such as RSA or EC.
// @property {string} Use - The "use" property in the WebhookVerificationKey struct identifies the
// intended use of the public key. It can have values like "sig" (signature), "enc" (encryption), or
// "wrapKey" (key wrapping). This property helps in determining the appropriate use of the key and can
// @property {string} X - The x member contains the x coordinate for the elliptic curve point. In the
// context of a WebhookVerificationKey, this likely refers to a point on an elliptic curve used for
// cryptographic purposes, such as for digital signatures. The x coordinate, along with the y
// coordinate, can be used to
// @property {string} Y - The "y" member contains the y coordinate for the elliptic curve point. This
// is used in the context of elliptic curve cryptography, which is a type of public key cryptography.
// The "y" coordinate, along with the "x" coordinate, is used to define a point on the ellipt
// @property {int32} CreatedAt - CreatedAt is a property of the WebhookVerificationKey struct that
// represents the timestamp (in seconds) when the key was created.
// @property {int32} ExpiredAt - ExpiredAt is a property of the WebhookVerificationKey struct that
// represents the expiration time of the key. It is an integer value that indicates the number of
// seconds since the Unix epoch (January 1, 1970, 00:00:00 UTC) at which the key will expire and
type WebhookVerificationKey struct {
	// The alg member identifies the cryptographic algorithm family used with the key.
	Alg string `json:"alg"`
	// The crv member identifies the cryptographic curve used with the key.
	Crv string `json:"crv"`
	// The kid (Key ID) member can be used to match a specific key. This can be used, for instance, to choose among a set of keys within the JWK during key rollover.
	Kid string `json:"kid"`
	// The kty (key type) parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC.
	Kty string `json:"kty"`
	// The use (public key use) parameter identifies the intended use of the public key.
	Use string `json:"use"`
	// The x member contains the x coordinate for the elliptic curve point.
	X string `json:"x"`
	// The y member contains the y coordinate for the elliptic curve point.
	Y         string `json:"y"`
	CreatedAt int32  `json:"created_at"`
	ExpiredAt int32  `json:"expired_at"`
}

// The WebhookVerification interface defines methods for retrieving a verification key and closing the
// interface.
// @property GetVerificationKey - GetVerificationKey is a method of the WebhookVerification interface
// that takes in a context and a keyId string as parameters and returns a pointer to a keyfunc.JWKS
// object and an error. This method is responsible for retrieving the verification key associated with
// the given keyId.
// @property {error} Close - Close is a method that closes any resources or connections used by the
// WebhookVerification implementation. It is typically used to clean up any resources that are no
// longer needed after the verification process is complete.
type WebhookVerification interface {
	GetVerificationKey(ctx context.Context, keyId string) (*keyfunc.JWKS, error)
	Close() error
}

// The type keyCacheItem represents an item in a cache that stores a key function and its expiration
// time.
// @property expiration - expiration is a property of the keyCacheItem struct that represents the time
// when the cached key will expire and need to be refreshed. It is of type time.Time.
// @property keyFunction - keyFunction is a pointer to a keyfunc.JWKS object, which is a struct that
// represents a JSON Web Key Set (JWKS) containing a collection of public keys used for verifying JSON
// Web Tokens (JWTs). The keyFunction object is used to retrieve the appropriate key from the JW
type keyCacheItem struct {
	expiration  time.Time
	keyFunction *keyfunc.JWKS
}

// The type `memoryWebhookVerification` is a struct that contains fields related to caching and
// verification of Plaid webhooks.
// @property {uint32} closed - closed is an unsigned integer that is used as a flag to indicate whether
// the webhook verification process has been closed or not. It is likely used to prevent any further
// processing of webhook events once the verification process has been closed.
// @property log - The `log` property is a pointer to a `zap.Logger` instance, which is a structured
// logging library for Go. It is used to log messages and errors throughout the codebase.
// @property lock - The `lock` property is a `sync.Mutex` type variable used for synchronizing access
// to the `cache` map. It ensures that only one goroutine can access the `cache` map at a time,
// preventing race conditions and ensuring data consistency.
// @property cache - `cache` is a map that stores key-value pairs where the key is a string and the
// value is a pointer to a `keyCacheItem` struct. This cache is used to store and retrieve webhook
// verification keys for Plaid API requests.
// @property cleanupTicker - `cleanupTicker` is a `time.Ticker` that is used to periodically trigger a
// cleanup function to remove expired items from the `cache` map. This helps to prevent the map from
// growing too large and consuming too much memory. The interval at which the cleanup function is
// triggered is determined by the
// @property closer - `closer` is a channel used to signal the closing of the
// `memoryWebhookVerification` instance. It is a channel of channels, where the inner channel is used
// to return an error (if any) during the closing process.
// @property plaid - The `plaid` property is a pointer to an instance of the `PlaidWrapper` struct,
// which is likely used to interact with the Plaid API.
type memoryWebhookVerification struct {
	closed        uint32
	log           *zap.Logger
	lock          sync.Mutex
	cache         map[string]*keyCacheItem
	cleanupTicker *time.Ticker
	closer        chan chan error
	plaid         *PlaidWrapper
}

// NewWebhookVerificationKeyFromPlaid creates a new WebhookVerificationKey object from a given Plaid JWK public key.
func NewWebhookVerificationKeyFromPlaid(input plaid.JWKPublicKey) (WebhookVerificationKey, error) {
	return WebhookVerificationKey{
		Alg:       input.GetAlg(),
		Crv:       input.GetCrv(),
		Kid:       input.GetKid(),
		Kty:       input.GetKty(),
		Use:       input.GetUse(),
		X:         input.GetX(),
		Y:         input.GetY(),
		CreatedAt: input.GetCreatedAt(),
		ExpiredAt: input.GetExpiredAt(),
	}, nil
}

// This code is creating a variable with an empty interface value that is used to assert that the
// `memoryWebhookVerification` struct implements the `WebhookVerification` interface. The underscore
// before the variable name is used to indicate that the variable is not actually used in the code, but
// is only there to perform the interface assertion.
var (
	_ WebhookVerification = &memoryWebhookVerification{}
)

// NewInMemoryWebhookVerification creates a new instance of an in-memory webhook verification system with a background
// worker for cache cleanup.
func NewInMemoryWebhookVerification(cleanupInterval time.Duration, logger *zap.Logger, plaidWrapper *PlaidWrapper) WebhookVerification {
	verification := &memoryWebhookVerification{
		closed:        0,
		log:           logger,
		lock:          sync.Mutex{},
		cache:         map[string]*keyCacheItem{},
		cleanupTicker: time.NewTicker(cleanupInterval),
		closer:        make(chan chan error, 1),
		plaid:         plaidWrapper,
	}
	go verification.cacheWorker() // Start the background worker.

	return verification
}

// GetVerificationKey retrieves the verification key associated with the given keyId. If the key is not present in the
// cache, then it will be retrieved from the Plaid API and added to the cache.
//
// NOTE: the cache of interest here is an in-memory one
// TODO: update to cache in redis instead
func (m *memoryWebhookVerification) GetVerificationKey(ctx context.Context, keyId string) (*keyfunc.JWKS, error) {
	if atomic.LoadUint32(&m.closed) > 0 {
		return nil, errors.New("webhook verification is closed")
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	item, ok := m.cache[keyId]
	if ok {
		if item.expiration.After(time.Now()) {
			m.log.Info("jwk function already present in cache, returning")
			return item.keyFunction, nil
		}

		m.log.Info("jwk function present in cache, but is expired; the cached function will be removed and a new one will be retrieved")
		delete(m.cache, keyId)
	}

	result, err := m.plaid.GetWebhookVerificationKey(ctx, keyId)
	if err != nil {
		return nil, err
	}

	var expiration time.Time
	if result.ExpiredAt != 0 {
		expiration = time.Unix(int64(result.ExpiredAt), 0)
	} else {
		// Making a huge assumption here, and this might end up causing problems later on. Maybe we should also add a
		// check here to make sure that items that are close to expiration even here should not be cached?
		expiration = time.Unix(int64(result.CreatedAt), 0).Add(30 * time.Minute)
	}

	var keys = struct {
		Keys []WebhookVerificationKey `json:"keys"`
	}{
		Keys: []WebhookVerificationKey{
			*result,
		},
	}

	encodedKeys, err := json.Marshal(keys)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert plaid verification key to json")
	}

	var jwksJSON json.RawMessage = encodedKeys

	jwksFunc, err := keyfunc.NewJSON(jwksJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create key function")
	}

	m.cache[keyId] = &keyCacheItem{
		expiration:  expiration,
		keyFunction: jwksFunc,
	}

	return jwksFunc, nil
}

// cacheWorker is a background worker that periodically cleans up the cache by removing expired items.
// It is started when the `memoryWebhookVerification` struct is created.
func (m *memoryWebhookVerification) cacheWorker() {
	for {
		select {
		case _ = <-m.cleanupTicker.C:
			m.cleanup()
		case promise := <-m.closer:
			m.log.Debug("closing jwk cache, stopping background worker")
			promise <- nil
			return
		}
	}
}

// cleanup removes expired items from the cache.
func (m *memoryWebhookVerification) cleanup() {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(m.cache) == 0 {
		m.log.Debug("no items in Plaid jwk cache, nothing to cleanup")
		return
	}

	m.log.Debug("cleaning up Plaid jwk cache")

	itemsToRemove := make([]string, 0, len(m.cache))
	for key, item := range m.cache {
		// If the item expiration is not in the future, then we need to add it to our list to be removed.
		if !item.expiration.After(time.Now()) {
			itemsToRemove = append(itemsToRemove, key)
		}
	}

	if len(itemsToRemove) == 0 {
		m.log.Debug("no items have expired in cache")
		return
	}

	m.log.Info("found %d expired item(s); cleaning them up", zap.Strings("items to remove", itemsToRemove))

	for _, key := range itemsToRemove {
		delete(m.cache, key)
	}
}

// Close closes the webhook verification system and cleans up any resources used by it.
func (m *memoryWebhookVerification) Close() error {
	if ok := atomic.CompareAndSwapUint32(&m.closed, 0, 1); !ok {
		return errors.New("webhook verification is already closed")
	}

	promise := make(chan error)
	m.closer <- promise

	return <-promise
}
