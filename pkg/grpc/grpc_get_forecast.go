package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetForecast implements apiv1.FinancialServiceServer
func (*Server) GetForecast(context.Context, *proto.GetForecastRequest) (*proto.GetForecastResponse, error) {
	panic("unimplemented")
}
