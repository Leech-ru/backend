package auth

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := m.cookieService.ReadAccessTokenCookie(c.Request())
		switch {
		case errors.Is(err, errorz.NoCookie):
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		case err != nil:
			return err
		}
		userID, err := m.tokenService.ParseAccessToken(c.Request().Context(), token)
		switch {
		case errors.Is(err, errorz.InvalidToken):
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.InvalidToken)
		case errors.Is(err, errorz.Unauthorized):
			return echo.NewHTTPError(http.StatusUnauthorized, errorz.Unauthorized)
		case err != nil:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Set("user_id", userID)

		return next(c)
	}
}
