package inmemoryverifier

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/pkg/errors"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
)

func NewWebhookVerificationKeyFromPlaid(input plaid.JWKPublicKey) (plaidhandler.WebhookVerificationKey, error) {
	return plaidhandler.WebhookVerificationKey{
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

type WebhookVerification interface {
	GetVerificationKey(ctx context.Context, keyId string) (*keyfunc.JWKS, error)
	Close() error
}

var (
	_ WebhookVerification = &memoryWebhookVerification{}
)

func NewInMemoryWebhookVerification(log *zap.Logger, plaid *plaidhandler.PlaidWrapper, cleanupInterval time.Duration) WebhookVerification {
	verification := &memoryWebhookVerification{
		closed:        0,
		log:           log,
		plaid:         plaid,
		lock:          sync.Mutex{},
		cache:         map[string]*keyCacheItem{},
		cleanupTicker: time.NewTicker(cleanupInterval),
		closer:        make(chan chan error, 1),
	}

	go verification.cacheWorker() // Start the background worker.

	return verification
}

type keyCacheItem struct {
	expiration  time.Time
	keyFunction *keyfunc.JWKS
}

type memoryWebhookVerification struct {
	closed        uint32
	log           *zap.Logger
	plaid         *plaidhandler.PlaidWrapper
	lock          sync.Mutex
	cache         map[string]*keyCacheItem
	cleanupTicker *time.Ticker
	closer        chan chan error
}

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

	m.log.Info("retrieving jwk from plaid")

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
		Keys []plaidhandler.WebhookVerificationKey `json:"keys"`
	}{
		Keys: []plaidhandler.WebhookVerificationKey{
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

func (m *memoryWebhookVerification) cacheWorker() {
	for {
		select {
		case <-m.cleanupTicker.C:
			m.cleanup()
		case promise := <-m.closer:
			m.log.Debug("closing jwk cache, stopping background worker")
			promise <- nil
			return
		}
	}
}

func (m *memoryWebhookVerification) cleanup() {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(m.cache) == 0 {
		m.log.Info("no items in Plaid jwk cache, nothing to cleanup")
		return
	}

	m.log.Info("cleaning up Plaid jwk cache")

	itemsToRemove := make([]string, 0, len(m.cache))
	for key, item := range m.cache {
		// If the item expiration is not in the future, then we need to add it to our list to be removed.
		if !item.expiration.After(time.Now()) {
			itemsToRemove = append(itemsToRemove, key)
		}
	}

	if len(itemsToRemove) == 0 {
		m.log.Info("no items have expired in cache")
		return
	}

	m.log.Info("found %d expired item(s); cleaning them up", zap.Any("num items", len(itemsToRemove)))

	for _, key := range itemsToRemove {
		delete(m.cache, key)
	}
}

func (m *memoryWebhookVerification) Close() error {
	if ok := atomic.CompareAndSwapUint32(&m.closed, 0, 1); !ok {
		return errors.New("webhook verification is already closed")
	}

	promise := make(chan error)
	m.closer <- promise

	return <-promise
}
