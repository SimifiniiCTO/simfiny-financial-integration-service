package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func (s *Server) GetLink(ctx context.Context, req *proto.GetLinkRequest) (*proto.GetLinkResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-link")
		defer span.End()
	}

	// get the required link
	// NOTE: accesstoken will still be encrypted even here
	link, err := s.conn.GetLink(ctx, req.UserId, req.LinkId, true)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// remove the access token from the response
	NullifyAccessToken(link)

	return &proto.GetLinkResponse{
		Link: link,
	}, nil
}

func NullifyAccessToken(link *proto.Link) {
	link.Token = nil
}
