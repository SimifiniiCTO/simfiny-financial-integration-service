package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// GetUserProfile implements apiv1.FinancialServiceServer
func (s *Server) GetUserProfile(ctx context.Context, req *proto.GetUserProfileRequest) (*proto.GetUserProfileResponse, error) {
	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	if err := req.ValidateAll(); err != nil {

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-get-profile")
		defer span.End()
	}

	// query db for user profile
	record, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// remove the access token from the response
	for _, link := range record.Link {
		NullifyAccessToken(link)
	}

	// query the financial context from the given user
	financialContext, err := s.clickhouseConn.
		GetFinancialContextForCurrentMonth(ctx, &req.UserId, 5)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUserProfileResponse{
		Profile:          record,
		FinancialContext: financialContext,
	}, nil
}
