package auth

import (
	"context"
	"github.com/google/uuid"
)

type tokenService interface {
	ParseAccessToken(ctx context.Context, token string) (uuid.UUID, error)
}

type AuthMiddleware struct {
	tokenService tokenService
}

func NewAuthMiddleware(tokenService tokenService) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: tokenService,
	}
}
