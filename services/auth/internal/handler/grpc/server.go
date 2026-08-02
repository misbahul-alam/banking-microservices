package grpc

import (
	"net"

	authpb "github.com/misbahul-alam/banking-microservices/proto/auth"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(authHandler *AuthHandler) *Server {
	server := grpc.NewServer()

	authpb.RegisterAuthServiceServer(server, authHandler)

	return &Server{
		grpcServer: server,
	}
}

func (s *Server) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}
