package instrumentation

import (
	"context"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type ServiceTelemetry struct {
	// The name of the service
	ServiceName string `json:"service_name"`
	// The version of the service
	ServiceVersion string `json:"service_version"`
	// The environment of the service
	ServiceEnvironment string `json:"service_environment"`
	// wether tracing is enabled
	TracingEnabled bool `json:"tracing_enabled"`
	// wether metrics are enabled
	MetricsEnabled bool `json:"metrics_enabled"`
	// The New Relic Sdk Object
	NewrelicSdk *newrelic.Application
}

type IServiceTelemetry interface {
	GetServiceName() string
	GetServiceVersion() string
	GetServiceEnvironment() string
	GetTracingEnabled() bool
	GetMetricsEnabled() bool
	GetTraceFromContext(ctx context.Context) *newrelic.Transaction
	GetTraceFromRequest(r *http.Request) *newrelic.Transaction
	WithContext(ctx context.Context, trace newrelic.Transaction) context.Context
	WithRequest(r *http.Request, trace newrelic.Transaction) *http.Request
	StartTransaction(name string) *newrelic.Transaction
	NewChildSpan(name string, parent newrelic.Transaction) *newrelic.Segment
	StartExternalSegment(txn *newrelic.Transaction, req *http.Request) *newrelic.ExternalSegment
	StartDatastoreSegment(txn *newrelic.Transaction, operation string) *newrelic.DatastoreSegment
	StartMessageQueueSegment(txn *newrelic.Transaction, queueName string) *newrelic.MessageProducerSegment
	NewRoundtripper(txn *newrelic.Transaction) *http.RoundTripper
}

var _ IServiceTelemetry = &ServiceTelemetry{}

// NewServiceTelemetry creates a new ServiceTelemetry
func NewServiceTelemetry(opts ...ServiceTelemetryOption) *ServiceTelemetry {
	telemetry := &ServiceTelemetry{}

	for _, opt := range opts {
		opt(telemetry)
	}

	return telemetry
}

// GetMetricsEnabled implements IServiceTelemetry
func (s *ServiceTelemetry) GetMetricsEnabled() bool {
	return s.MetricsEnabled
}

// GetServiceEnvironment implements IServiceTelemetry
func (s *ServiceTelemetry) GetServiceEnvironment() string {
	return s.ServiceEnvironment
}

// GetServiceName implements IServiceTelemetry
func (s *ServiceTelemetry) GetServiceName() string {
	return s.ServiceName
}

// GetServiceVersion implements IServiceTelemetry
func (s *ServiceTelemetry) GetServiceVersion() string {
	return s.ServiceVersion
}

// GetTraceFromContext implements IServiceTelemetry
func (s *ServiceTelemetry) GetTraceFromContext(ctx context.Context) *newrelic.Transaction {
	if s.TracingEnabled {
		return newrelic.FromContext(ctx)
	}

	return nil
}

// GetTraceFromRequest implements IServiceTelemetry
func (s *ServiceTelemetry) GetTraceFromRequest(r *http.Request) *newrelic.Transaction {
	if s.TracingEnabled {
		return newrelic.FromContext(r.Context())
	}

	return nil
}

// GetTracingEnabled implements IServiceTelemetry
func (s *ServiceTelemetry) GetTracingEnabled() bool {
	return s.TracingEnabled
}

// NewChildSpan implements IServiceTelemetry
func (s *ServiceTelemetry) NewChildSpan(name string, parent newrelic.Transaction) *newrelic.Segment {
	if s.TracingEnabled {
		return parent.StartSegment(name)
	}

	return nil
}

// StartTransaction implements IServiceTelemetry
func (s *ServiceTelemetry) StartTransaction(name string) *newrelic.Transaction {
	if s.TracingEnabled {
		return s.NewrelicSdk.StartTransaction(name)
	}

	return nil
}

// WithContext implements IServiceTelemetry
func (s *ServiceTelemetry) WithContext(ctx context.Context, trace newrelic.Transaction) context.Context {
	if s.TracingEnabled {
		return newrelic.NewContext(ctx, &trace)
	}

	return ctx
}

// WithRequest implements IServiceTelemetry
func (s *ServiceTelemetry) WithRequest(r *http.Request, trace newrelic.Transaction) *http.Request {
	if s.TracingEnabled {
		return nil
	}

	return r
}

// StartExternalSegment implements IServiceTelemetry
func (s *ServiceTelemetry) StartExternalSegment(txn *newrelic.Transaction, req *http.Request) *newrelic.ExternalSegment {
	if s.TracingEnabled {
		return newrelic.StartExternalSegment(txn, req)
	}

	return nil
}

// NewRoundTripper allows you to instrument external calls without
// calling StartExternalSegment by modifying the http.Client's Transport
// field.  If the Transaction parameter is nil, the RoundTripper
// returned will look for a Transaction in the request's context (using
// FromContext). This is recommended because it allows you to reuse the
// same client for multiple transactions.
func (s *ServiceTelemetry) NewRoundtripper(txn *newrelic.Transaction) *http.RoundTripper {
	client := &http.Client{}
	if s.TracingEnabled {
		h := newrelic.NewRoundTripper(client.Transport)
		return &h
	}

	return nil
}

// StartDatastoreSegment implements IServiceTelemetry
func (s *ServiceTelemetry) StartDatastoreSegment(txn *newrelic.Transaction, operation string) *newrelic.DatastoreSegment {
	if s.TracingEnabled {
		segment := newrelic.DatastoreSegment{
			StartTime: txn.StartSegmentNow(),
			// Product, Collection, and Operation are the most important
			// fields to populate because they are used in the breakdown
			// metrics.
			Product:   newrelic.DatastorePostgres,
			Operation: operation,
		}

		return &segment
	}

	return nil
}

// StartMessageQueueSegment implements IServiceTelemetry
func (s *ServiceTelemetry) StartMessageQueueSegment(txn *newrelic.Transaction, queueName string) *newrelic.MessageProducerSegment {
	if s.TracingEnabled {
		segment := newrelic.MessageProducerSegment{
			StartTime:       txn.StartSegmentNow(),
			Library:         "aws-sqs",
			DestinationType: newrelic.MessageQueue,
			DestinationName: queueName,
		}

		return &segment
	}

	return nil
}
