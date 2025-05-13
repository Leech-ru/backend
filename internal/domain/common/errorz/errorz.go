package errorz

import "errors"

var (
	Unauthorized = errors.New("forbidden: not authorized")
)
