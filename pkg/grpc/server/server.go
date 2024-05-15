package grpcserver

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	addr string
	*grpc.Server
}

func New(addr string, opts ...Option) *GrpcServer {
	conf := getDefaultConfig(addr)
	for _, opt := range opts {
		opt(conf)
	}

	defaultInterceptors := []grpc.UnaryServerInterceptor{
		RecoveryUnaryServerInterceptor(),
		TrackIdInterceptor(),
	}

	var unaryInterceptors = append(
		defaultInterceptors,
		conf.unaryInterceptors...,
	)

	defaultOptions := conf.serverOptions
	//if conf.tracer {
	//	defaultOptions = append(defaultOptions, grpc.StatsHandler(otelgrpc.NewServerHandler()))
	//}

	var serverOptions = append(defaultOptions, grpc.ChainUnaryInterceptor(
		unaryInterceptors...,
	))

	srv := grpc.NewServer(serverOptions...)

	server := &GrpcServer{
		addr:   addr,
		Server: srv,
	}

	return server
}

func (s *GrpcServer) Start(ctx context.Context) error {
	go func() {
		l, err := net.Listen("tcp", s.addr)
		if err != nil {
			return
		}
		if err = s.Serve(l); err != nil {
			return
		}
	}()
	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
