package service_provider

import (
	"Leech-ru/internal/domain/service/token"
	"context"
	"github.com/google/uuid"
)

type tokenService interface {
	NewToken(ctx context.Context, id uuid.UUID) (string, error)
	UpdateToken(ctx context.Context, id uuid.UUID) (string, error)
	GetIDByToken(ctx context.Context, token string) (uuid.UUID, error)
	GetTokenByID(ctx context.Context, id uuid.UUID) (string, error)
}

func (s *ServiceProvider) TokenService() tokenService {
	if s.tokenService == nil {
		s.tokenService = token.NewTokenService(s.DB(), s.JWTService())
	}
	return s.tokenService
}
