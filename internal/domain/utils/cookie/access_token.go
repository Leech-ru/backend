package cookie

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	accessCookieName = "user_auth_access_token"
)

// SetAccessTokenCookie creates and immediately sets Access-Token Cookie in response.
func SetAccessTokenCookie(c echo.Context, token string, ttl time.Duration, devMode bool) {
	cookie := &http.Cookie{
		Name:     accessCookieName,
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

// ReadAccessTokenCookie reads Cook from a request.
func ReadAccessTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(accessCookieName)
	switch {
	case errors.Is(err, http.ErrNoCookie):
		return "", errorz.NoCookie
	case err != nil:
		return "", err
	}
	return cookie.Value, nil
}

// ClearAccessTokenCookie sets an empty access-token Cook with an expired validity period.
func ClearAccessTokenCookie(c echo.Context, devMode bool) {
	cookie := &http.Cookie{
		Name:     accessCookieName,
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
