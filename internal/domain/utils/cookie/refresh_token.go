package cookie

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	refreshCookieName = "user_auth_refresh_token"
)

// SetRefreshTokenCookie creates and sets a secure HTTP-only refresh token cookie.
// It supports development mode by relaxing SameSite and Secure policies.
func SetRefreshTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool) {
	cookie := &http.Cookie{
		Name:     refreshCookieName,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(ttl),
		MaxAge:   int(ttl.Seconds()),
		HttpOnly: true,
		Secure:   !devMode,
	}

	if devMode {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.SameSite = http.SameSiteStrictMode
	}

	c.SetCookie(cookie)
}

// ReadRefreshTokenCookie extracts the refresh token value from the request cookie.
// Returns a domain-specific error if the cookie is missing.
func ReadRefreshTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(refreshCookieName)
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
func ClearRefreshTokenCookie(c echo.Context, devMode bool) {
	cookie := &http.Cookie{
		Name:     refreshCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   !devMode,
	}

	if devMode {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.SameSite = http.SameSiteStrictMode
	}

	c.SetCookie(cookie)
}
