package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// PlaidInitiateTokenExchange implements apiv1.FinancialServiceServer
func (*Server) PlaidInitiateTokenExchange(context.Context, *proto.PlaidInitiateTokenExchangeRequest) (*proto.PlaidInitiateTokenExchangeResponse, error) {
	panic("unimplemented")
}
