package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database/sqlc"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/dto"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/password"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/repository"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/token"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
	RefreshToken(ctx context.Context, tokenStr string) (*dto.AuthResponse, error)
}

type authService struct {
	userRepo        repository.UserRepository
	passwordService password.Service
	tokenMaker      token.Maker
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	passwordService password.Service,
	tokenMaker token.Maker,
	accessDuration time.Duration,
	refreshDuration time.Duration,
) AuthService {
	return &authService{
		userRepo:        userRepo,
		passwordService: passwordService,
		tokenMaker:      tokenMaker,
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := s.passwordService.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return s.generateTokens(user.ID)
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := s.passwordService.Compare(user.PasswordHash, req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return s.generateTokens(user.ID)
}

func (s *authService) RefreshToken(ctx context.Context, tokenStr string) (*dto.AuthResponse, error) {
	payload, err := s.tokenMaker.VerifyToken(tokenStr)
	if err != nil {
		return nil, ErrInvalidToken
	}

	user, err := s.userRepo.GetUserByID(ctx, payload.UserID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return s.generateTokens(user.ID)
}

func (s *authService) generateTokens(userID pgtype.UUID) (*dto.AuthResponse, error) {
	accessToken, _, err := s.tokenMaker.CreateToken(userID, s.accessDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := s.tokenMaker.CreateToken(userID, s.refreshDuration)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
