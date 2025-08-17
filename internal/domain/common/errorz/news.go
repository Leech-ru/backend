package errorz

import "errors"

var (
	InvalidNewsFormat = errors.New("invalid news format")
	NewsNotFound      = errors.New("news not found")
)
