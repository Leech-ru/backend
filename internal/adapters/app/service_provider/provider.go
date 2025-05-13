package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/logger"
	"github.com/go-playground/form"
)

type ServiceProvider struct {
	loggerConfig config.LoggerConfig
	pgConfig     config.PGConfig
	httpConfig   config.HTTPConfig
	jwtConfig    config.JWTConfig
	mailConfig   config.MailConfig

	db *ent.Client

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
