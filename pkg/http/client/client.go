package httpclient

import (
	"github.com/imroc/req/v3"
)

type HttpClient struct {
	*req.Client
}

func NewHttpClient(opts ...Option) *HttpClient {
	configs := defaultConfig()
	for _, opt := range opts {
		opt(configs)
	}

	client := req.C().
		SetBaseURL(configs.baseUrl).
		SetTimeout(configs.timeout)

	return &HttpClient{Client: client}
}
