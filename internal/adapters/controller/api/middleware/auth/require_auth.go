package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (m *AuthMiddleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return nil // TODO new error
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header")
		}
		token := parts[1]

		tokenEntity, err := m.authRepo.GetByToken(c.Request().Context(), token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		userId := tokenEntity.UserID
		//if tokenEntity.UserID != token {
		//
		//}
		c.Set("user_id", userId)
		return next(c)
	}
}
