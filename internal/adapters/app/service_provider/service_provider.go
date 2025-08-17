package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/logger"
	"github.com/go-playground/form"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type ServiceProvider struct {
	loggerConfig   config.LoggerConfig
	pgConfig       config.PGConfig
	redisConfig    config.RedisConfig
	httpConfig     config.ServerConfig
	jwtConfig      config.JWTConfig
	jsonInfoConfig config.JsonInfoConfig
	mailConfig     config.MailConfig
	minioConfig    config.MinIOConfig

	db    *ent.Client
	redis *redis.Client
	minio *minio.Client

	logger      *logger.Logger
	validator   *validator.Validator
	formDecoder *form.Decoder

	imageService     imageService
	jwtService       jwtService
	tokenService     tokenService
	orderService     orderService
	userService      userService
	cosmeticsService cosmeticsService
	categoryService  categoryService
	infoService      infoService
	partnerService   partnerService
	newsService      newsService
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
