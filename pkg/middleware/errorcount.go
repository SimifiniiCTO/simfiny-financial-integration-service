package middleware

import (
	"context"
	"fmt"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/pkg/metrics"
	rkgrpcctx "github.com/rookie-ninja/rk-grpc/interceptor/context"
	rkgrpcmid "github.com/rookie-ninja/rk-grpc/v2/middleware"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ErrorCountUnaryServerInterceptor Create new unary server interceptor to capture number of errors.
func ErrorCountUnaryServerInterceptor(e *metrics.MetricsEngine, m *metrics.FinancialIntegrationServiceMetrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var opStatusCode codes.Code = codes.OK

		ctx = rkgrpcmid.WrapContextForServer(ctx)
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))
		resp, err := handler(ctx, req)

		if err != nil {
			opStatusCode = codes.Internal
			e.RecordCounterMetric(m.ErrorCountMetric, fmt.Sprintf("%s-%s", method, path), code.Code(opStatusCode))
		}

		return resp, err
	}
}

// ErrorCountStreamServerInterceptor Create new stream server interceptor to capture request latency.
func ErrorCountStreamServerInterceptor(e *metrics.MetricsEngine, m *metrics.FinancialIntegrationServiceMetrics) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)

		var opStatusCode codes.Code = codes.OK

		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(wrappedStream.WrappedContext))
		err := handler(srv, wrappedStream)

		if err != nil {
			opStatusCode = codes.Internal
			e.RecordCounterMetric(m.ErrorCountMetric, fmt.Sprintf("%s-%s", method, path), code.Code(opStatusCode))
		}

		return err
	}
}
