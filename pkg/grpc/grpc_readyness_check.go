package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// ReadynessCheck implements apiv1.FinancialServiceServer
func (*Server) ReadynessCheck(context.Context, *proto.ReadynessCheckRequest) (*proto.ReadynessCheckResponse, error) {
	panic("unimplemented")
}
