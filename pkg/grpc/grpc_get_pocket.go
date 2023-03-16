package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetPocket implements apiv1.FinancialServiceServer
func (*Server) GetPocket(context.Context, *proto.GetPocketRequest) (*proto.GetPocketResponse, error) {
	panic("unimplemented")
}
