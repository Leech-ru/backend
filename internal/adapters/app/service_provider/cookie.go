package service_provider

import (
	"Leech-ru/internal/domain/service/cookie"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type cookieService interface {
	// Access
	SetAccessTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool)
	ReadAccessTokenCookie(r *http.Request) (string, error)
	ClearAccessTokenCookie(c echo.Context, devMode bool)

	// Refresh
	SetRefreshTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool)
	ReadRefreshTokenCookie(r *http.Request) (string, error)
	ClearRefreshTokenCookie(c echo.Context, devMode bool)
}

func (s *ServiceProvider) CookieService() cookieService {
	if s.cookieService == nil {
		s.cookieService = cookie.NewService()
	}
	return s.cookieService
}
