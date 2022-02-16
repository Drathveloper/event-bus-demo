package model

import "github.com/google/uuid"

type UserEntity struct {
	ID       uuid.UUID
	Username string
	Password string
	Role     string
}
