package token

import (
	"Leech-ru/internal/adapters/repository/postgres/refresh_token"
	"Leech-ru/internal/adapters/repository/valkey/access_token"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type accessTokenRepo interface {
	Set(ctx context.Context, userID uuid.UUID, value string, exp time.Time) error
	Get(ctx context.Context, userID uuid.UUID) (string, error)
}

type refreshTokenRepo interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*ent.RefreshToken, error)
	Update(ctx context.Context, entity ent.RefreshToken) (*ent.RefreshToken, error)
	Upsert(ctx context.Context, entity ent.RefreshToken) (*ent.RefreshToken, error)
}

type jwtService interface {
	GenerateToken(userID uuid.UUID, ttl time.Duration) (token string, jti string, err error)
	ParseToken(tokenString string) (userID uuid.UUID, jti string, err error)
}

type jwtConfig interface {
	RefreshTokenExpires() time.Duration
	AccessTokenExpires() time.Duration
}

// TODO настроить ошибки нормально (вместо nil errorz.Err...)
type tokenService struct {
	refreshTokenRepo refreshTokenRepo
	accessTokenRepo  accessTokenRepo
	jwtService       jwtService
	jwtConfig        jwtConfig
}

func NewTokenService(clientDB *ent.Client, clientRedis *redis.Client, jwtService jwtService, jwtConfig jwtConfig) *tokenService {
	return &tokenService{
		refreshTokenRepo: refresh_token.NewTokenRepo(clientDB),
		accessTokenRepo:  access_token.NewTokenRepo(clientRedis),
		jwtService:       jwtService,
		jwtConfig:        jwtConfig,
	}
}
