package util

import (
	"event-bus-demo/infrastructure/error"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error.InfrastructureError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	if err != nil {
		return "", error.NewHashingError(err.Error())
	}
	return string(bytes), nil
}
