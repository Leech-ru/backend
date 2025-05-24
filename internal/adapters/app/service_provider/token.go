package service_provider

import (
	"Leech-ru/internal/domain/service/token"
	"context"
	"github.com/google/uuid"
)

type tokenService interface {
	GenerateAccessToken(ctx context.Context, userID uuid.UUID) (string, error)
	GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error)
	ParseAccessToken(ctx context.Context, token string) (uuid.UUID, error)
	ParseRefreshToken(ctx context.Context, token string) (uuid.UUID, error)
	RevokeAccessToken(ctx context.Context, token string) (uuid.UUID, error)
	RevokeRefreshToken(ctx context.Context, token string) (uuid.UUID, error)
	LogoutAllSessions(ctx context.Context, userID uuid.UUID) error
}

func (s *ServiceProvider) TokenService() tokenService {
	if s.tokenService == nil {
		s.tokenService = token.NewTokenService(s.DB(), s.Redis(), s.JWTService(), s.JWTConfig())
	}
	return s.tokenService
}
