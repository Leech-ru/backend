package errorz

import "errors"

var (
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrAccessDenied         = errors.New("access denied")
	ErrEmailAlreadyExists   = errors.New("email already exists")
)
