package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// PlaidExchangeToken implements apiv1.FinancialServiceServer
func (*Server) PlaidExchangeToken(context.Context, *proto.PlaidExchangeTokenRequest) (*proto.PlaidExchangeTokenResponse, error) {
	panic("unimplemented")
}
