package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/logger"
	"github.com/go-playground/form"
	"github.com/redis/go-redis/v9"
)

type ServiceProvider struct {
	loggerConfig config.LoggerConfig
	pgConfig     config.PGConfig
	redisConfig  config.RedisConfig
	httpConfig   config.HTTPConfig
	jwtConfig    config.JWTConfig
	mailConfig   config.MailConfig

	db    *ent.Client
	redis *redis.Client

	logger    *logger.Logger
	validator *validator.Validator
	decoder   *form.Decoder

	jwtService   jwtService
	tokenService tokenService
	orderService orderService
	userService  userService
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
