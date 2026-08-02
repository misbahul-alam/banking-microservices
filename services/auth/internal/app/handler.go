package app

import "github.com/misbahul-alam/banking-microservices/services/auth/internal/handler/grpc"

func (c *Container) registerHandlers() {
	c.AuthHandler = grpc.NewAuthHandler(c.AuthService)
}
