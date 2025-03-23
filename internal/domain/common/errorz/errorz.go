package errorz

import "errors"

var (
	ValidationFailed = errors.New("Validation failed")
)
