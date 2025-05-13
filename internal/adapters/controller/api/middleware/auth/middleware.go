package auth

import (
	user "Leech-ru/internal/adapters/repository/postgres/token"
	"Leech-ru/pkg/ent"
	"context"
)

type authRepo interface {
	GetByToken(ctx context.Context, inputToken string) (*ent.Token, error)
}

type AuthMiddleware struct {
	authRepo authRepo
}

func NewAuthMiddleware(cleint *ent.Client) *AuthMiddleware {
	return &AuthMiddleware{
		authRepo: user.NewTokenRepo(cleint),
	}
}
