package metrics

// Metric encompasses the relative metadata associated with a service level metric being emitted
type Metric struct {
	MetricName       string
	Name             string
	Help             string
	Subsystem        Subsystem
	Namespace        Namespace
	ServiceName      string
	MetricPartitions map[string]string
}

// Counter a metric type that only lineary increases
type Counter struct {
	Metric
}

// Summary summarizes numerous facets of a metric
type Summary struct {
	Metric
}

// Gauge a metric type that can increase and decrease
type Gauge struct {
	Metric
}

type Namespace string
type Subsystem string

const (
	RequestNamespace       Namespace = "request.namespace"
	DatabaseNamespace      Namespace = "database.namespace"
	DistributedTxNamespace Namespace = "distributed_transaction.namespace"
	ServiceNamespace       Namespace = "service.namespace"
	PlaidNamespace         Namespace = "plaid.namespace"
)

const (
	GrpcSubSystem       Subsystem = "grpc.subsystem"
	DbSubSystem         Subsystem = "database.subsystem"
	ThirdPartySubSystem Subsystem = "thirdparty.subsystem"
	ErrorSubSystem      Subsystem = "error.subsytem"
)
