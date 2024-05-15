package grpcclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	target string
	*grpc.ClientConn
}

func New(target string, options ...Option) (*GrpcClient, error) {
	conf := getDefaultConfig(target)
	for _, opt := range options {
		opt(conf)
	}

	var unaryInterceptors = append(
		[]grpc.UnaryClientInterceptor{
			conf.retryInterceptor,
			TrackIdUnaryClientInterceptor(),
		},
		conf.unaryInterceptors...,
	)

	var clientOptions = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	//if conf.tracer {
	//	clientOptions = append(clientOptions, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	//}

	clientOptions = append(clientOptions, conf.clientOptions...)
	clientOptions = append(clientOptions, grpc.WithChainUnaryInterceptor(unaryInterceptors...))

	conn, err := grpc.Dial(target, clientOptions...)

	if err != nil {
		return nil, err
	}
	return &GrpcClient{ClientConn: conn}, nil
}
