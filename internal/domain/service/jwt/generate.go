package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

// GenerateToken generated new jwt token with jti.
func (s *jwtService) GenerateToken(userID uuid.UUID, ttl time.Duration) (string, string, error) {
	now := time.Now()
	jti := uuid.New().String()

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(now),
		Subject:   userID.String(),
		ID:        jti,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", "", err
	}

	return signedToken, jti, nil
}
