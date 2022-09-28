package metrics

import (
	"fmt"
)

func newRequestCountMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.count", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the number of request serviced by the service partitioned by name and status code",
		Subsystem:   GrpcSubSystem,
		Namespace:   RequestNamespace,
	}
}

func newRequestLatencyMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.latency", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the latency associated with various requests partitioned by service name, target name, status code, and latency",
		Subsystem:   GrpcSubSystem,
		Namespace:   RequestNamespace,
	}
}

func newErrorCountMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.error.count", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the numbr of errors encountered by the service",
		Subsystem:   ErrorSubSystem,
		Namespace:   ServiceNamespace,
	}
}

func newRequestStatusSummaryMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.summary", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the status of all requests serviced by the service",
		Subsystem:   Subsystem(RequestNamespace),
		Namespace:   RequestNamespace,
	}
}

func newDatabaseOperationStatusMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.db.operation.summary", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the status of a database operation",
		Subsystem:   Subsystem(DbSubSystem),
		Namespace:   DatabaseNamespace,
	}
}

func newDatabaseOperationLatencyCountMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.db.operation.latency", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the latency associated with a database operations partitioned by operation name",
		Subsystem:   Subsystem(DbSubSystem),
		Namespace:   DatabaseNamespace,
	}
}

func newPlaidOperationStatusMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.plaid.operation.status", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the status of plaid operations",
		Subsystem:   Subsystem(ThirdPartySubSystem),
		Namespace:   PlaidNamespace,
	}
}
