package interceptor

import (
	"context"

	"github.com/homma509/9task/shared/md"
	"google.golang.org/grpc"
)

func XTraceID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
			req interface{},
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler) (interface{}, error) {
				traceID := md.GetTraceIDFromContext(ctx)
				ctx := md.AddTraceIDToContext(ctx, traceID)
				return handler(ctx, req)
			}
	)
}