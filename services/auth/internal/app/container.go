package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/config"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/token"

	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database/sqlc"
	authgrpc "github.com/misbahul-alam/banking-microservices/services/auth/internal/handler/grpc"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/password"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/repository"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/service"
)

type Container struct {
	DB              *pgxpool.Pool
	Queries         *sqlc.Queries
	UserRepository  repository.UserRepository
	PasswordService password.Service
	JWTService      token.Maker
	AuthService     service.AuthService
	AuthHandler     *authgrpc.AuthHandler
	GRPCServer      *authgrpc.Server
}

func New(cfg config.Config, db *pgxpool.Pool) *Container {
	container := &Container{
		DB: db,
	}

	container.registerRepositories()
	container.registerServices(cfg)
	container.registerHandlers()
	container.registerServers()

	return container
}
