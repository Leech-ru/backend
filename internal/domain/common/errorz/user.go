package errorz

import "errors"

var (
	EmailAlreadyExist      = errors.New("email already exist")
	InvalidEmailOrPassword = errors.New("invalid email or password")
	PasswordMismatch       = errors.New("password mismatch")
	PasswordsCoincidence   = errors.New("new password should differ from the old")
	UserNotFound           = errors.New("user not found")
)
