package password

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash generates a secure hash of the password using bcrypt.
// The returned hash is in a format that can be directly stored in a database.
// An error is returned if the hashing fails.
func PasswordHash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword compares a plain-text password with a bcrypt hashed password.
// Returns nil on success, or an error if they don't match.
func VerifyPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
