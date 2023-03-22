package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func (s *Server) GetLinks(ctx context.Context, req *proto.GetLinksRequest) (*proto.GetLinksResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-get-links")
		defer span.End()
	}

	// ensure that the user exists
	// note, the user will already have the associated links
	user, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetLinksResponse{
		Links: user.Link,
	}, nil
}
