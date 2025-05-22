package errorz

import "errors"

var (
	Unauthorized = errors.New("forbidden: not authorized")
	NoCookie     = errors.New("no mandatory cookies")
)
