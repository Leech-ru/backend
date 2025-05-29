package role

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/types"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Middleware) RequireRole(minRole types.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			val := c.Get("user_id")
			userID, ok := val.(uuid.UUID)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
			}

			role, err := m.userService.GetRoleByID(c.Request().Context(), userID)
			switch {
			case errors.Is(err, errorz.UserNotFound):
				return echo.NewHTTPError(http.StatusForbidden, errorz.PermissionDenied)
			case err != nil:
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if role < minRole {
				return echo.NewHTTPError(http.StatusForbidden, errorz.PermissionDenied)
			}

			return next(c)
		}
	}
}
