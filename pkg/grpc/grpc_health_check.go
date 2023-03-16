package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// HealthCheck implements apiv1.FinancialServiceServer
func (*Server) HealthCheck(context.Context, *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	panic("unimplemented")
}
