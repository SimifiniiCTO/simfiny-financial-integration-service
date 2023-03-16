package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetSmartGoalsByPocketId implements apiv1.FinancialServiceServer
func (*Server) GetSmartGoalsByPocketId(context.Context, *proto.GetSmartGoalsByPocketIdRequest) (*proto.GetSmartGoalsByPocketIdResponse, error) {
	panic("unimplemented")
}
