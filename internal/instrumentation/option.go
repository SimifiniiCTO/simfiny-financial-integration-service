package instrumentation

import newrelic "github.com/newrelic/go-agent"

// ServiceTelemetryOption is a function that configures a ServiceTelemetry
type ServiceTelemetryOption func(*ServiceTelemetry)

// WithServiceName configures the service name
func WithServiceName(name string) ServiceTelemetryOption {
	return func(t *ServiceTelemetry) {
		t.ServiceName = name
	}
}

// WithServiceVersion configures the service version
func WithServiceVersion(version string) ServiceTelemetryOption {
	return func(t *ServiceTelemetry) {
		t.ServiceVersion = version
	}
}

// WithServiceEnvironment configures the service environment
func WithServiceEnvironment(environment string) ServiceTelemetryOption {
	return func(t *ServiceTelemetry) {
		t.ServiceEnvironment = environment
	}
}

// WithTracingEnabled configures the service tracing
func WithTracingEnabled(enabled bool) ServiceTelemetryOption {
	return func(t *ServiceTelemetry) {
		t.TracingEnabled = enabled
	}
}

// WithMetricsEnabled configures the service metrics
func WithMetricsEnabled(enabled bool) ServiceTelemetryOption {
	return func(t *ServiceTelemetry) {
		t.MetricsEnabled = enabled
	}
}

// WithNewrelicSdk configures the service newrelic sdk
func WithNewrelicSdk(newrelicSdk newrelic.Application) ServiceTelemetryOption {
	return func(t *ServiceTelemetry) {
		t.NewrelicSdk = newrelicSdk
	}
}
