package dto

import (
	"event-bus-demo/domain/model"
	"github.com/google/uuid"
)

type GetCategoriesResponse struct {
	Categories []GetCategoryResponse `json:"categories"`
}

type GetCategoryResponse struct {
	ID   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (req CreateCategoryRequest) ToEvent(ID uuid.UUID) model.CreateCategoryEvent {
	return model.CreateCategoryEvent{
		ID:   ID,
		Name: req.Name,
	}
}

type CreateCategoryResponse struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (req UpdateCategoryRequest) ToEvent(ID uuid.UUID) model.UpdateCategoryNameEvent {
	return model.UpdateCategoryNameEvent{
		ID:   ID,
		Name: req.Name,
	}
}
