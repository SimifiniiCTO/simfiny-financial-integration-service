package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// UpdateUserProfile implements apiv1.FinancialServiceServer
func (s *Server) UpdateUserProfile(ctx context.Context, req *proto.UpdateUserProfileRequest) (*proto.UpdateUserProfileResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-update-profile")
		defer span.End()
	}

	// perform an update operation in the database
	if err := s.conn.UpdateUserProfile(ctx, req.Profile); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateUserProfileResponse{
		Profile:        req.Profile,
		ProfileUpdated: true,
	}, nil
}
