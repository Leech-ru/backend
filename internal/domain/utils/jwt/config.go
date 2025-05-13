package jwt

import (
	"crypto/rsa"
	"time"
)

type jwtService struct {
	privateKey   *rsa.PrivateKey
	tokenExpires time.Duration
}

type jwtConfig interface {
	PrivateKey() *rsa.PrivateKey
	TokenExpires() time.Duration
}

// NewJWTService returns new jwt service.
func NewJWTService(config jwtConfig) *jwtService {
	return &jwtService{
		privateKey:   config.PrivateKey(),
		tokenExpires: config.TokenExpires(),
	}
}
