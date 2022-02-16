package model

import (
	"github.com/google/uuid"
	"time"
)

type ToDoEntity struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Categories  []CategoryEntity
}
