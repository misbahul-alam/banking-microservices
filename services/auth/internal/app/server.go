package app

import "github.com/misbahul-alam/banking-microservices/services/auth/internal/handler/grpc"

func (c *Container) registerServers() {
	c.GRPCServer = grpc.NewServer(c.AuthHandler)
}
