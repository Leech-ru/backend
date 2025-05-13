package token

import (
	"Leech-ru/internal/adapters/repository/postgres/token"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type tokenRepo interface {
	Create(ctx context.Context, entity ent.Token) (*ent.Token, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.Token, error)
	GetByToken(ctx context.Context, inputToken string) (*ent.Token, error)
	Update(ctx context.Context, tokenEntity ent.Token) (*ent.Token, error)
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}

type jwtService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(tokenString string) (uuid.UUID, error)
}

type tokenService struct {
	jwtService jwtService
	tokenRepo  tokenRepo
}

func NewTokenService(client *ent.Client, jwtService jwtService) *tokenService {
	return &tokenService{
		jwtService: jwtService,
		tokenRepo:  token.NewTokenRepo(client),
	}
}
