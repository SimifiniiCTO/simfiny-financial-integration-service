package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// HealthCheck implements apiv1.FinancialServiceServer
func (s *Server) HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {

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

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-health-check")
		defer span.End()
	}

	// syncAllAccountsBatchJob, err := taskhandler.NewSyncAllPlatformConnectedPlaidAccounts()
	// if err != nil {
	// 	return nil, err
	// }

	// ctx = context.WithValue(ctx, "trigger", "grpc-health-check")
	// ctx, _ = context.WithTimeout(ctx, 60*time.Minute)
	// _, err = s.Taskprocessor.EnqueueTask(
	// 	ctx,
	// 	syncAllAccountsBatchJob)
	// if err != nil {
	// 	return nil, err
	// }

	trigger := "grpc-health-check"
	if err := s.th.ExecuteBatchSync(ctx, &trigger); err != nil {
		return nil, err
	}

	return &proto.HealthCheckResponse{
		Healthy: true,
	}, nil
}
