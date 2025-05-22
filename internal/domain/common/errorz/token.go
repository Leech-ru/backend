package errorz

import "errors"

var (
	InvalidToken        = errors.New("invalid token format")
	TokenNotFound       = errors.New("token not found")
	TokenRevoked        = errors.New("token has been revoked")
	UserAlreadyHasToken = errors.New("user already has token")
)
