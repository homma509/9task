package interceptor

import (
	"context"
	"log"
	"time"

	"github.com/homma509/9task/shared/md"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// XTraceID コンテキストにTraceIDの追加
func XTraceID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		traceID := md.GetTraceIDFromContext(ctx)
		ctx = md.AddTraceIDToContext(ctx, traceID)
		return handler(ctx, req)
	}
}

const loggingFmt = "TraceID:%s\tFullMethod:%s\tElapsedTime:%s\tStatusCode:%s\tError:%s\n"

// Logging ログの出力
func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		h, err := handler(ctx, req)
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		log.Printf(loggingFmt,
			md.GetTraceIDFromContext(ctx),
			info.FullMethod,
			time.Since(start),
			status.Code(err), errMsg)
		return h, err
	}
}
