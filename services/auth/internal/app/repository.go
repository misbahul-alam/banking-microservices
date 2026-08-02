package app

import (
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database/sqlc"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/repository"
)

func (c *Container) registerRepositories() {
	c.Queries = sqlc.New(c.DB)

	c.UserRepository = repository.NewUserRepository(c.Queries)
}
