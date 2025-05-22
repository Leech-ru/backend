package jwt

import (
	"crypto/rsa"
)

type jwtService struct {
	privateKey *rsa.PrivateKey
}

type jwtConfig interface {
	PrivateKey() *rsa.PrivateKey
	PublicKey() *rsa.PublicKey
}

// NewJWTService returns new jwt service.
func NewJWTService(config jwtConfig) *jwtService {
	return &jwtService{
		privateKey: config.PrivateKey(),
	}
}
