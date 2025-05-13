package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/pkg/logger"
)

type ServiceProvider struct {
	loggerConfig config.LoggerConfig
	pgConfig     config.PGConfig
	httpConfig   config.HTTPConfig
	mailConfig   config.MailConfig

	logger    *logger.Logger
	validator *validator.Validator

	orderService orderService
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
