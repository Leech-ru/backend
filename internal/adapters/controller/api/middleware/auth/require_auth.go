package auth

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"fmt"
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

		tokenEntity, err := m.authRepo.GetByToken(c.Request().Context(), token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		}
		switch {
		case errors.As(err, &errorz.TokenNotFound):
			echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		case err != nil:
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("internal server error: %w", err))
		}

		userID, err := m.jwtService.ParseToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		}
		if tokenEntity.UserID != userID {
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		}
		c.Set("user_id", userID)
		return next(c)
	}
}
