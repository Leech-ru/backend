package auth

import (
	"Leech-ru/internal/domain/common/errorz"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (m *AuthMiddleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header")
		}
		token := parts[1]

		ctx := c.Request().Context()

		userID, err := m.tokenService.ValidateToken(ctx, token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		}

		c.Set("user_id", userID)

		return next(c)
	}
}
