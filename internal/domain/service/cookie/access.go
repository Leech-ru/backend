package cookie

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// SetAccessTokenCookie creates and immediately sets Access-Token Cookie in response.
func (s *cookieService) SetAccessTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool) {
	cookie := &http.Cookie{
		Name:     s.accessCookieName,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(ttl),
		MaxAge:   int(ttl.Seconds()),
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	if !devMode {
		cookie.HttpOnly = true
		cookie.Secure = true

	}

	c.SetCookie(cookie)
}

// ReadAccessTokenCookie reads Cook from a request.
func (s *cookieService) ReadAccessTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(s.accessCookieName)
	switch {
	case errors.Is(err, http.ErrNoCookie):
		return "", errorz.NoCookie
	case err != nil:
		return "", err
	}
	return cookie.Value, nil
}

// ClearAccessTokenCookie sets an empty access-token Cook with an expired validity period.
func (s *cookieService) ClearAccessTokenCookie(c echo.Context, devMode bool) {
	cookie := &http.Cookie{
		Name:     s.accessCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	if !devMode {
		cookie.HttpOnly = true
		cookie.Secure = true

	}

	c.SetCookie(cookie)
}
