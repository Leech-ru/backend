package config

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JWTConfig interface {
	PrivateKey() *rsa.PrivateKey
	PublicKey() *rsa.PublicKey
	RefreshTokenExpires() time.Duration
	AccessTokenExpires() time.Duration
}

type jwtConfig struct {
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
	refreshTokenExpires time.Duration
	accessTokenExpires  time.Duration
}

func NewJWTConfig() (JWTConfig, error) {
	privateKeyPath := viper.GetString("service.jwt.path-to-private-key")
	if privateKeyPath == "" {
		return nil, errors.New("private key path not configured")
	}

	// Читаем файл с приватным ключом
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := &privateKey.PublicKey

	refreshExpiresRaw := viper.GetString("service.jwt.refresh-token-expires")
	refreshTokenExpires, err := time.ParseDuration(refreshExpiresRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid token_expires duration: %w", err)
	}

	accessExpiresRaw := viper.GetString("service.jwt.access-token-expires")
	accessTokenExpires, err := time.ParseDuration(accessExpiresRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid token_expires duration: %w", err)
	}

	return &jwtConfig{
		privateKey:          privateKey,
		publicKey:           publicKey,
		refreshTokenExpires: refreshTokenExpires,
		accessTokenExpires:  accessTokenExpires,
	}, nil
}

func (cfg *jwtConfig) PrivateKey() *rsa.PrivateKey {
	return cfg.privateKey
}

func (cfg *jwtConfig) PublicKey() *rsa.PublicKey {
	return cfg.publicKey
}

func (cfg *jwtConfig) RefreshTokenExpires() time.Duration {
	return cfg.refreshTokenExpires
}
func (cfg *jwtConfig) AccessTokenExpires() time.Duration {
	return cfg.accessTokenExpires
}
