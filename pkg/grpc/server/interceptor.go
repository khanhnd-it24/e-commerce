package grpcserver

import (
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/khanhnd-it24/e-commerce/pkg/ctxutil"
	httpconst "github.com/khanhnd-it24/e-commerce/pkg/http/const"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(func(ctx context.Context, p any) (err error) {
		//logx.Error(ctx, errors.New("panic grpc"), "panic with err %s and stack", p, debug.Stack())
		return status.Errorf(codes.Internal, "panic with err %s", p)
	}))
}

func TrackIdInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		trackId := metadata.ExtractIncoming(ctx).Get(ctxutil.TrackIdKey)

		if trackId == "" {
			requestId := metadata.ExtractIncoming(ctx).Get(httpconst.HeaderXRequestID)
			if requestId != "" {
				trackId = requestId
			} else {
				trackId = uuid.New().String()
			}
		}

		ctx = ctxutil.WithTrackId(ctx, trackId)
		return handler(ctx, req)
	}
}
