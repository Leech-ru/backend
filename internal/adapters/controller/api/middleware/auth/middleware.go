package auth

import (
	user "Leech-ru/internal/adapters/repository/postgres/token"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type jwtService interface {
	ParseToken(tokenString string) (uuid.UUID, error)
}

type authRepo interface {
	GetByToken(ctx context.Context, inputToken string) (*ent.Token, error)
}

type AuthMiddleware struct {
	jwtService jwtService
	authRepo   authRepo
}

func NewAuthMiddleware(cleint *ent.Client, jwtService jwtService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
		authRepo:   user.NewTokenRepo(cleint),
	}
}
