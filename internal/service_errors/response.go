package service_errors

import (
	"net/http"
)

// APIResponse encapsulates some common features of a response
type APIResponse struct {
	// Header contain a map of all HTTP header keys to values. Its behavior and
	// caveats are identical to that of http.Header.
	Header http.Header

	// IdempotencyKey contains the idempotency key used with this request.
	// Idempotency keys are a financial service specific concept that helps guarantee that
	// requests that fail and need to be retried are not duplicated.
	IdempotencyKey string

	// RawJSON contains the response body as raw bytes.
	RawJSON []byte

	// RequestID contains a string that uniquely identifies the request.
	// Used for debugging or support purposes.
	RequestID string

	// Status is a status code and message. e.g. "200 OK"
	Status string

	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}

// APIResource is a type assigned to structs that may come from Stripe API
// endpoints and contains facilities common to all of them.
type APIResource struct {
	LastResponse *APIResponse `json:"-"`
}

// SetLastResponse sets the HTTP response that returned the API resource.
func (r *APIResource) SetLastResponse(response *APIResponse) {
	r.LastResponse = response
}

func newAPIResponse(res *http.Response, resBody []byte) *APIResponse {
	return &APIResponse{
		Header:         res.Header,
		IdempotencyKey: res.Header.Get("Idempotency-Key"),
		RawJSON:        resBody,
		RequestID:      res.Header.Get("Request-Id"),
		Status:         res.Status,
		StatusCode:     res.StatusCode,
	}
}
