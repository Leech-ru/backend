package token

import (
	"Leech-ru/internal/adapters/controller/api/validator"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/form"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type tokenService interface {
	GenerateAccessToken(ctx context.Context, userID uuid.UUID) (string, error)
	GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error)
	ParseRefreshToken(ctx context.Context, token string) (uuid.UUID, error)
}

type cookieService interface {
	// Access
	SetAccessTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool)

	// Refresh
	SetRefreshTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool)
	ReadRefreshTokenCookie(r *http.Request) (string, error)
}

type serverConfig interface {
	DevMode() bool
}

type jwtConfig interface {
	AccessTokenExpires() time.Duration
	RefreshTokenExpires() time.Duration
}

type handler struct {
	cookieService cookieService
	tokenService  tokenService
	jwtConfig     jwtConfig
	serverConfig  serverConfig
	validator     *validator.Validator
	formDecoder   *form.Decoder
}

func NewHandler(
	tokenService tokenService,
	cookieService cookieService,
	jwtConfig jwtConfig,
	serverConfig serverConfig,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		cookieService: cookieService,
		tokenService:  tokenService,
		jwtConfig:     jwtConfig,
		serverConfig:  serverConfig,
		validator:     validator,
		formDecoder:   formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.POST("/auth/refresh", h.Refresh)

}
