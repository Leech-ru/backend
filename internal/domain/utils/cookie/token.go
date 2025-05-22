package cookie

import (
	"net/http"
	"time"
)

func NewTokenCookie(refreshToken string, ttl time.Duration, devMode bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "user_auth_token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(ttl)
	cookie.Path = "/"
	cookie.Secure = true
	cookie.HttpOnly = true
	if devMode {
		cookie.SameSite = http.SameSiteStrictMode
	} else {
		cookie.SameSite = http.SameSiteNoneMode
	}
	return cookie
}

func ReadTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("user_auth_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
