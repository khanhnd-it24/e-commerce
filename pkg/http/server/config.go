package httpserver

import "time"

type Config struct {
	addr   string
	prefix string
	// implement tracer
	serviceName  string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Option func(c *Config)

func defaultConfig(addr string) *Config {
	return &Config{
		addr:         addr,
		prefix:       "",
		readTimeout:  30 * time.Second,
		writeTimeout: 30 * time.Second,
	}
}

func WithServiceName(serviceName string) Option {
	return func(c *Config) {
		c.serviceName = serviceName
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(c *Config) {
		c.readTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(c *Config) {
		c.writeTimeout = writeTimeout
	}
}

func WithPrefix(prefix string) Option {
	return func(c *Config) {
		c.prefix = prefix
	}
}
