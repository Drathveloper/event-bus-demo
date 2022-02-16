package model

import (
	"github.com/google/uuid"
	"time"
)

type ToDo struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Categories  []Category
}
