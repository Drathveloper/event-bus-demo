package model

import (
	"github.com/google/uuid"
	"time"
)

const ToDoEventTopic = "TODO"

type CreateToDoEvent struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
	Categories  []uuid.UUID
}

func (CreateToDoEvent) GetTopic() string {
	return ToDoEventTopic
}

func (CreateToDoEvent) GetName() string {
	return "CreateToDoEvent"
}

type UpdateToDoEvent struct {
	ID          uuid.UUID
	Title       string
	Description string
	UpdatedAt   time.Time
}

func (UpdateToDoEvent) GetTopic() string {
	return ToDoEventTopic
}

func (UpdateToDoEvent) GetName() string {
	return "UpdateToDoEvent"
}

type DeleteToDoEvent struct {
	ID uuid.UUID
}

func (DeleteToDoEvent) GetTopic() string {
	return ToDoEventTopic
}

func (DeleteToDoEvent) GetName() string {
	return "DeleteToDoEvent"
}

type GetToDoEvent struct {
	ID uuid.UUID
}

func (GetToDoEvent) GetTopic() string {
	return ToDoEventTopic
}

func (GetToDoEvent) GetName() string {
	return "GetToDoEvent"
}

type RemoveCategoriesFromToDoEvent struct {
	ToDoID     uuid.UUID
	Categories []uuid.UUID
}

func (RemoveCategoriesFromToDoEvent) GetTopic() string {
	return ToDoEventTopic
}

func (RemoveCategoriesFromToDoEvent) GetName() string {
	return "RemoveCategoriesFromToDoEvent"
}

type AddCategoriesFromToDoEvent struct {
	ToDoID     uuid.UUID
	Categories []uuid.UUID
}

func (AddCategoriesFromToDoEvent) GetTopic() string {
	return ToDoEventTopic
}

func (AddCategoriesFromToDoEvent) GetName() string {
	return "AddCategoriesFromToDoEvent"
}
