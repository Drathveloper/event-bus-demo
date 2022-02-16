package model

import (
	"github.com/google/uuid"
)

type CategoryEntity struct {
	ID   uuid.UUID
	Name string
}
