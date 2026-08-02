package grpc

import (
	"context"

	authpb "github.com/misbahul-alam/banking-microservices/proto/auth"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/dto"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/service"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer

	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
	dtoReq := dto.RegisterRequest{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	res, err := h.authService.Register(ctx, dtoReq)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	dtoReq := dto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.authService.Login(ctx, dtoReq)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (h *AuthHandler) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.AuthResponse, error) {
	dtoReq := dto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	res, err := h.authService.RefreshToken(ctx, dtoReq.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthResponse{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
	}, nil
}
