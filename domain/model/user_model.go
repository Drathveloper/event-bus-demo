package model

import (
	"event-bus-demo/infrastructure/error"
	"event-bus-demo/infrastructure/util"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	password string
	Role     string
}

func (u User) GetPassword() string {
	return u.password
}

func (u User) SetNonHashedPassword(password string) error.InfrastructureError {
	password, err := util.HashPassword(password)
	if err != nil {
		return error.NewHashingError("error while hashing password")
	}
	u.password = password
	return nil
}

func (u User) SetHashedPassword(hashedPassword string) {
	u.password = hashedPassword
}
