package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash password hash
func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", hashedPassword), nil
}

// VerifyPassword verify password with hash
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
