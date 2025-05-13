package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/domain/utils/jwt"
	"github.com/google/uuid"
)

type jwtService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(tokenString string) (uuid.UUID, error)
}

func (s *ServiceProvider) JWTService(jwtConfig config.JWTConfig) jwtService {
	if s.jwtService == nil {
		s.jwtService = jwt.NewJWTService(jwtConfig.PrivateKey(), jwtConfig.TokenExpires())
	}
	return s.jwtService
}
