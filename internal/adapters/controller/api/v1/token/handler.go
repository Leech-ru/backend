package token

import (
	"Leech-ru/internal/adapters/controller/api/validator"
	"context"
	"github.com/go-playground/form"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

type tokenService interface {
	GenerateAccessToken(ctx context.Context, userID uuid.UUID) (string, error)
	ParseRefreshToken(ctx context.Context, token string) (uuid.UUID, error)
}

type jwtConfig interface {
	AccessTokenExpires() time.Duration
}

type handler struct {
	tokenService tokenService
	jwtConfig    jwtConfig
	validator    *validator.Validator
	formDecoder  *form.Decoder
}

func NewHandler(
	tokenService tokenService,
	jwtConfig jwtConfig,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		tokenService: tokenService,
		jwtConfig:    jwtConfig,
		validator:    validator,
		formDecoder:  formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.POST("/auth/refresh", h.Refresh)

}
