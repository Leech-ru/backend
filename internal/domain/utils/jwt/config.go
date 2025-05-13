package jwt

import (
	"crypto/rsa"
	"time"
)

type jwtService struct {
	privateKey   *rsa.PrivateKey
	tokenExpires time.Duration
}

// NewJWTService returns new jwt service.
func NewJWTService(privateKey *rsa.PrivateKey, tokenExpires time.Duration) *jwtService {
	return &jwtService{
		privateKey:   privateKey,
		tokenExpires: tokenExpires,
	}
}
