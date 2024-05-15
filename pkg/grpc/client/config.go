package grpcclient

import (
	"google.golang.org/grpc"
	"time"
)

type Config struct {
	target string
	//tracer             bool
	timeoutInterceptor grpc.UnaryClientInterceptor
	retryInterceptor   grpc.UnaryClientInterceptor
	clientOptions      []grpc.DialOption
	unaryInterceptors  []grpc.UnaryClientInterceptor
}

func getDefaultConfig(target string) *Config {
	return &Config{
		target:             target,
		timeoutInterceptor: TimeoutUnaryClientInterceptor(30 * time.Second),
		retryInterceptor:   DefaultRetryUnaryClientInterceptor(),
	}
}

type Option func(conf *Config)

func WithTimeout(dur time.Duration) Option {
	return func(conf *Config) {
		conf.timeoutInterceptor = TimeoutUnaryClientInterceptor(dur)
	}
}

//func WithTracer(tracer bool) Option {
//	return func(conf *Config) {
//		conf.tracer = tracer
//	}
//}

func WithUnaryInterceptor(inters ...grpc.UnaryClientInterceptor) Option {
	return func(conf *Config) {
		if conf.unaryInterceptors == nil {
			conf.unaryInterceptors = make([]grpc.UnaryClientInterceptor, 0)
		}

		conf.unaryInterceptors = append(conf.unaryInterceptors, inters...)
	}
}

func WithClientOptions(opts ...grpc.DialOption) Option {
	return func(conf *Config) {
		if conf.clientOptions == nil {
			conf.clientOptions = make([]grpc.DialOption, 0)
		}
		conf.clientOptions = append(conf.clientOptions, opts...)
	}
}
