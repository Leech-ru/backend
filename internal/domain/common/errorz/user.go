package errorz

import "errors"

var (
	EmailAlreadyExist    = errors.New("email already exist")
	PasswordMismatch     = errors.New("password mismatch")
	PasswordsCoincidence = errors.New("new password should differ from the old")
	UserNotFound         = errors.New("user not found")
)
