package errorz

import "errors"

var (
	InvalidCategoryFormat = errors.New("invalid category format")
	CategoryNotFound      = errors.New("category not found")
)
