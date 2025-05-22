package service_provider

import (
	"Leech-ru/internal/domain/service/jwt"
	"github.com/google/uuid"
	"time"
)

type jwtService interface {
	GenerateToken(userID uuid.UUID, ttl time.Duration) (token string, jti string, err error)
	ParseToken(tokenString string) (userID uuid.UUID, jti string, err error)
}

func (s *ServiceProvider) JWTService() jwtService {
	if s.jwtService == nil {
		s.jwtService = jwt.NewJWTService(s.JWTConfig())
	}
	return s.jwtService
}
