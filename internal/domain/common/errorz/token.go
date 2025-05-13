package errorz

import "errors"

var (
	TokenNotFound = errors.New("invalid token")
)
