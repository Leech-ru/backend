package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

// GenerateToken generated new jwt token.
func (s *jwtService) GenerateToken(userID uuid.UUID) (string, error) {
	// Стандартные claims с временем жизни
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenExpires)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userID.String(),
	}

	// Создаем токен с RSA256 подписью
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Подписываем токен приватным ключом
	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
