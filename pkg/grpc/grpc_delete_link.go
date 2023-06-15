package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func (s *Server) DeleteLink(ctx context.Context, req *proto.DeleteLinkRequest) (*proto.DeleteLinkResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-delete-link")
		defer span.End()
	}

	// get the link by id
	link, err := s.conn.GetLink(ctx, req.UserId, req.LinkId, true)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// ensure the link token exist for the link
	if link.Token == nil {
		return nil, status.Error(codes.NotFound, "link token not found")
	}

	// TODO: implement this as a workflow
	// decrypt the link token
	accessToken, err := s.DecryptUserAccessToken(ctx, link.Token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// only delete from plaid's vantage point if the link type is plaid
	if link.LinkType == proto.LinkType_LINK_TYPE_PLAID {
		// delete the link from plaids vantage point and any stripe
		// subscriptions associated with the linked accounts
		if err := s.plaidClient.DeleteItem(ctx, accessToken); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// delete the required link
	if err := s.conn.DeleteLink(ctx, req.UserId, req.LinkId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteLinkResponse{
		LinkId: req.LinkId,
	}, nil
}
