package auth

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type tokenService interface {
	ParseAccessToken(ctx context.Context, token string) (uuid.UUID, error)
}

type cookieService interface {
	ReadAccessTokenCookie(r *http.Request) (string, error)
}

type Middleware struct {
	tokenService  tokenService
	cookieService cookieService
}

func NewAuthMiddleware(tokenService tokenService, cookieService cookieService) *Middleware {
	return &Middleware{
		tokenService:  tokenService,
		cookieService: cookieService,
	}
}
