package cookie

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// SetRefreshTokenCookie creates and sets a secure HTTP-only refresh token cookie.
// It supports development mode by relaxing SameSite and Secure policies.
func (s *cookieService) SetRefreshTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool) {
	cookie := &http.Cookie{
		Name:     s.refreshCookieName,
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

// ReadRefreshTokenCookie extracts the refresh token value from the request cookie.
// Returns a domain-specific error if the cookie is missing.
func (s *cookieService) ReadRefreshTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(s.refreshCookieName)
	switch {
	case errors.Is(err, http.ErrNoCookie):
		return "", errorz.NoCookie
	case err != nil:
		return "", err
	}
	return cookie.Value, nil
}

// ClearRefreshTokenCookie invalidates the refresh token cookie by setting
// it with an expired timestamp and MaxAge=-1, forcing client removal.
func (s *cookieService) ClearRefreshTokenCookie(c echo.Context, devMode bool) {
	cookie := &http.Cookie{
		Name:     s.refreshCookieName,
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
