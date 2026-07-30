package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (sqlc.User, error)
	GetByEmail(ctx context.Context, email string) (sqlc.User, error)
	UpdateUser(ctx context.Context, params sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	ListUsers(ctx context.Context, params sqlc.ListUsersParams) ([]sqlc.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error) {
	return r.db.CreateUser(ctx, params)
}

func (r *userRepository) GetUserByID(ctx context.Context, id pgtype.UUID) (sqlc.User, error) {
	return r.db.GetUserByID(ctx, id)
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return r.db.GetUserByEmail(ctx, email)
}

func (r *userRepository) UpdateUser(ctx context.Context, params sqlc.UpdateUserParams) (sqlc.User, error) {
	return r.db.UpdateUser(ctx, params)
}

func (r *userRepository) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	return r.db.DeleteUser(ctx, id)
}

func (r *userRepository) ListUsers(ctx context.Context, params sqlc.ListUsersParams) ([]sqlc.User, error) {
	return r.db.ListUsers(ctx, params)
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return user.ID.Valid, nil
}
