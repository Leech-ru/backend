package errorz

import "errors"

var (
	InvalidImageFormat = errors.New("invalid image format")
	ImageNotFound      = errors.New("image not found")
)
