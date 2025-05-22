package cookie

import (
	"Leech-ru/internal/domain/common/errorz"
	"errors"
	"net/http"
	"time"
)

const (
	accessCookieName = "user_auth_access_token"
)

func NewAccessTokenCookie(refreshToken string, ttl time.Duration, devMode bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = accessCookieName
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(ttl)
	cookie.Path = "/"
	cookie.Secure = true
	cookie.HttpOnly = true
	if devMode {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.SameSite = http.SameSiteStrictMode
		cookie.HttpOnly = true
	}
	return cookie
}

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
