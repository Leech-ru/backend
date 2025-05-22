package errorz

import "errors"

var (
	UserAlreadyHasToken = errors.New("user already has token")
	TokenNotFound       = errors.New("token not found")
	TokenRevoked        = errors.New("token has been revoked")
)
