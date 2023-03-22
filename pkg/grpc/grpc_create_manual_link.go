package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateManualLink implements apiv1.FinancialServiceServer
// This endpoint is used to create a manual link for a user. A manual link is used to associate a manually created bank account.
// All links created via plaid should not be created via this endpoint.
func (s *Server) CreateManualLink(ctx context.Context, req *proto.CreateManualLinkRequest) (*proto.CreateManualLinkResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-create-manual-link")
		defer span.End()
	}

	// this endpoint is only used to create manual links
	// manual links are used to associate an abstraction of a manual bank account for the user
	manualLink := req.GetManualAccountLink()
	manualLink.LinkType = proto.LinkType_LINK_TYPE_MANUAL
	manualLink.LinkStatus = proto.LinkStatus_LINK_STATUS_SUCCESS

	// create the required manual link
	link, err := s.conn.CreateLink(ctx, req.UserId, req.ManualAccountLink)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateManualLinkResponse{
		LinkId: link.GetId(),
	}, nil
}
