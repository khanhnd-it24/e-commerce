package httpclient

import "time"

type Config struct {
	baseUrl string
	timeout time.Duration
	//implement tracer
	//implement batching
}

func defaultConfig() *Config {
	return &Config{
		baseUrl: "",
		timeout: 5 * time.Second,
	}
}

type Option func(c *Config)

func WithBaseUrl(baseUrl string) Option {
	return func(c *Config) {
		c.baseUrl = baseUrl
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.timeout = timeout
	}
}
