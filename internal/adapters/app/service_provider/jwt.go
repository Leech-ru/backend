package service_provider

import (
	"Leech-ru/internal/domain/utils/jwt"
	"github.com/google/uuid"
)

type jwtService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(tokenString string) (uuid.UUID, error)
}

func (s *ServiceProvider) JWTService() jwtService {
	if s.jwtService == nil {
		s.jwtService = jwt.NewJWTService(s.JWTConfig())
	}
	return s.jwtService
}
