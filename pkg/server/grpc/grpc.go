package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/lllllan02/iam/pkg/log"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	port int
	host string

	logger *log.Logger
}

type Option func(s *Server)

func NewServer(logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		Server: grpc.NewServer(),
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		s.logger.Sugar().Fatalf("Failed to listen: %v", err)
	}
	if err = s.Server.Serve(lis); err != nil {
		s.logger.Sugar().Fatalf("Failed to serve: %v", err)
	}
	return nil

}
func (s *Server) Stop(ctx context.Context) error {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	s.Server.GracefulStop()

	s.logger.Sugar().Info("Server exiting")

	return nil
}
