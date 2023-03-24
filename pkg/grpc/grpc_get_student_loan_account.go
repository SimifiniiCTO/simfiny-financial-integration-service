package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetStudentLoanAccount implements apiv1.FinancialServiceServer
func (*Server) GetStudentLoanAccount(context.Context, *proto.GetStudentLoanAccountRequest) (*proto.GetStudentLoanAccountResponse, error) {
	panic("unimplemented")
}
