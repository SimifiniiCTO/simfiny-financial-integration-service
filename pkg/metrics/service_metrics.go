package metrics

import "fmt"

type FinancialIntegrationServiceMetrics struct {
	// Tracks the number of request serviced by the service partitioned by name and status code
	RequestCountMetric *Metric

	// Tracks the latency associatd with a various requests partitioned by service name, target name,
	// status code, and latency
	RequestLatencyMetric *Metric

	// Tracks the numbr of errors encountered by the service
	ErrorCountMetric *Metric

	// Tracks the status of all requests serviced by the service
	RequestStatusSummaryMetric *Metric

	// Tracks the status of a database operation
	DatabaseOperationStatusMetric *Metric

	// Tracks the latency associated with a database operations partitioned by operation name
	DatabaseOperationLatencyCountMetric *Metric

	// Tracks the status of plaid operations
	PlaidOperationStatusMetric *Metric
}

var (
	ErrInvalidServiceName error = fmt.Errorf("invalid input argument, service name cannot be nil")
)

func NewFinancialIntergrationServiceMetrics(serviceName *string) (*FinancialIntegrationServiceMetrics, error) {
	if serviceName == nil {
		return nil, ErrInvalidServiceName
	}

	return &FinancialIntegrationServiceMetrics{
		RequestCountMetric:                  newRequestCountMetric(serviceName),
		RequestLatencyMetric:                newRequestLatencyMetric(serviceName),
		ErrorCountMetric:                    newErrorCountMetric(serviceName),
		RequestStatusSummaryMetric:          newRequestStatusSummaryMetric(serviceName),
		DatabaseOperationStatusMetric:       newDatabaseOperationStatusMetric(serviceName),
		DatabaseOperationLatencyCountMetric: newDatabaseOperationLatencyCountMetric(serviceName),
		PlaidOperationStatusMetric:          newPlaidOperationStatusMetric(serviceName),
	}, nil
}
