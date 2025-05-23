package auth

import (
	"context"
	"github.com/google/uuid"
)

type tokenService interface {
	ParseAccessToken(ctx context.Context, token string) (uuid.UUID, error)
}

type Middleware struct {
	tokenService tokenService
}

func NewAuthMiddleware(tokenService tokenService) *Middleware {
	return &Middleware{
		tokenService: tokenService,
	}
}
