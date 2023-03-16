package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// DeleteSmartGoal implements apiv1.FinancialServiceServer
func (*Server) DeleteSmartGoal(context.Context, *proto.DeleteSmartGoalRequest) (*proto.DeleteSmartGoalResponse, error) {
	panic("unimplemented")
}
