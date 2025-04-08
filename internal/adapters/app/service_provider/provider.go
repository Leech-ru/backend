package service_provider

import (
	"LutiLeech/internal/adapters/controller/api/validator"
	"LutiLeech/pkg/logger"
)

type ServiceProvider struct {
	logger    *logger.Logger
	validator *validator.Validator

	orderService orderService
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
