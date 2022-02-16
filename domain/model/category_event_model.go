package model

import "github.com/google/uuid"

const CategoryEventTopic = "CATEGORY"

type CreateCategoryEvent struct {
	ID   uuid.UUID
	Name string
}

func (CreateCategoryEvent) GetTopic() string {
	return CategoryEventTopic
}

func (CreateCategoryEvent) GetName() string {
	return "CreateCategoryEvent"
}

type UpdateCategoryNameEvent struct {
	ID   uuid.UUID
	Name string
}

func (UpdateCategoryNameEvent) GetTopic() string {
	return CategoryEventTopic
}

func (UpdateCategoryNameEvent) GetName() string {
	return "UpdateCategoryNameEvent"
}

type DeleteCategoryEvent struct {
	ID uuid.UUID
}

func (DeleteCategoryEvent) GetTopic() string {
	return CategoryEventTopic
}

func (DeleteCategoryEvent) GetName() string {
	return "DeleteCategoryEvent"
}

type GetCategoryByIDEvent struct {
	ID uuid.UUID
}

func (GetCategoryByIDEvent) GetTopic() string {
	return CategoryEventTopic
}

func (GetCategoryByIDEvent) GetName() string {
	return "GetCategoryByIDEvent"
}
