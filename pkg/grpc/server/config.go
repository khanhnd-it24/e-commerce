package grpcserver

import "google.golang.org/grpc"

type Config struct {
	addr string
	//tracer            bool
	serverOptions     []grpc.ServerOption
	unaryInterceptors []grpc.UnaryServerInterceptor
}

func getDefaultConfig(addr string) *Config {
	return &Config{
		addr:              addr,
		serverOptions:     []grpc.ServerOption{},
		unaryInterceptors: []grpc.UnaryServerInterceptor{},
	}
}

type Option func(conf *Config)

//func WithTracer(tracer bool) Option {
//	return func(conf *Config) {
//		conf.tracer = tracer
//	}
//}

func WithUnaryInterceptor(inters ...grpc.UnaryServerInterceptor) Option {
	return func(conf *Config) {
		if conf.unaryInterceptors == nil {
			conf.unaryInterceptors = make([]grpc.UnaryServerInterceptor, 0)
		}

		conf.unaryInterceptors = append(conf.unaryInterceptors, inters...)
	}
}

func WithServerOptions(opts ...grpc.ServerOption) Option {
	return func(conf *Config) {
		if conf.serverOptions == nil {
			conf.serverOptions = make([]grpc.ServerOption, 0)
		}
		conf.serverOptions = append(conf.serverOptions, opts...)
	}
}
