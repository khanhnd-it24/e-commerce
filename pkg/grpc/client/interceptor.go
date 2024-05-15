package grpcclient

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"github.com/khanhnd-it24/gobase/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"time"
)

func DefaultRetryUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return retry.UnaryClientInterceptor(
		retry.WithMax(5),
		retry.WithCodes(codes.Unavailable),
		retry.WithBackoff(retry.BackoffExponential(50*time.Millisecond)),
	)
}

func TimeoutUnaryClientInterceptor(dur time.Duration) grpc.UnaryClientInterceptor {
	return timeout.UnaryClientInterceptor(dur)
}

func TrackIdUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		trackId := ctxutil.GetTrackId(ctx)
		ctx = metadata.AppendToOutgoingContext(ctx, ctxutil.TrackIdKey, trackId)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
