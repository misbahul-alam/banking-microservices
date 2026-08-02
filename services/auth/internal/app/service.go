package app

import (
	"time"

	"github.com/misbahul-alam/banking-microservices/services/auth/internal/config"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/password"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/service"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/token"
)

func (c *Container) registerServices(cfg config.Config) {
	c.PasswordService = password.New(12)

	c.JWTService, _ = token.NewJWTMaker(cfg.JWTSecret)

	c.AuthService = service.NewAuthService(c.UserRepository, c.PasswordService, c.JWTService, time.Minute*15, time.Hour*24*30)
}
