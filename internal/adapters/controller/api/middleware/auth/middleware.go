package auth

import (
	"context"
	"github.com/google/uuid"
)

type jwtService interface {
	ParseToken(tokenString string) (userID uuid.UUID, jti string, err error)
}

type tokenService interface {
	ValidateToken(ctx context.Context, tokenStr string) (uuid.UUID, error)
}

type AuthMiddleware struct {
	jwtService   jwtService
	tokenService tokenService
}

func NewAuthMiddleware(jwtService jwtService, tokenService tokenService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService:   jwtService,
		tokenService: tokenService,
	}
}
