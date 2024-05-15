package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type HttpServer struct {
	srv  *http.Server
	Root *gin.RouterGroup
}

func NewHttpServer(addr string, opts ...Option) *HttpServer {
	configs := defaultConfig(addr)
	for _, opt := range opts {
		opt(configs)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	engine := gin.New()
	root := engine.RouterGroup.Group(configs.prefix)

	srv := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  configs.readTimeout,
		WriteTimeout: configs.writeTimeout,
	}

	return &HttpServer{
		srv:  srv,
		Root: root,
	}
}

func (s *HttpServer) Start(ctx context.Context) error {
	addr := s.srv.Addr
	if addr == "" {
		addr = ":http"
	}
	if addr[0] != ':' {
		addr = fmt.Sprintf(":%s", addr)
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		_ = s.srv.Serve(ln)
	}()

	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	err := s.srv.Shutdown(ctx)
	if err != nil {
		wErr := fmt.Errorf("failed to stop http server: %w", err)
		return wErr
	}
	return nil
}
