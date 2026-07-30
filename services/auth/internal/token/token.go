package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	UserID    pgtype.UUID `json:"user_id"`
	IssuedAt  time.Time   `json:"issued_at"`
	ExpiredAt time.Time   `json:"expired_at"`
	jwt.RegisteredClaims
}

type Maker interface {
	CreateToken(userID pgtype.UUID, duration time.Duration) (string, *Payload, error)
	VerifyToken(tokenStr string) (*Payload, error)
}
