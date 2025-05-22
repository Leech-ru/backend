package cookie

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"net/http"
	"time"
)

const (
	refreshCookieName = "user_auth_refresh_token"
)

func NewRefreshTokenCookie(refreshToken string, ttl time.Duration, devMode bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = refreshCookieName
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(ttl)
	cookie.Path = "/"
	cookie.Secure = true
	if devMode {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.SameSite = http.SameSiteStrictMode
		cookie.HttpOnly = true
	}
	return cookie
}

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
