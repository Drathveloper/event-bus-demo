package dto

import (
	"event-bus-demo/domain/model"
	"github.com/google/uuid"
	"time"
)

type CreateToDoRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	CreatedAt   time.Time   `json:"createdAt" binding:"required"`
	Categories  []uuid.UUID `json:"categories"`
}

func (req CreateToDoRequest) ToEvent(ID uuid.UUID) model.CreateToDoEvent {
	return model.CreateToDoEvent{
		ID:          ID,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   req.CreatedAt,
		Categories:  req.Categories,
	}
}

type UpdateToDoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (req UpdateToDoRequest) ToEvent(ID uuid.UUID) model.UpdateToDoEvent {
	return model.UpdateToDoEvent{
		ID:          ID,
		Title:       req.Title,
		Description: req.Description,
	}
}

type GetToDoResponse struct {
	ID          uuid.UUID             `json:"id" binding:"required"`
	Title       string                `json:"title" binding:"required"`
	Description string                `json:"description" binding:"required"`
	CreatedAt   *time.Time            `json:"createdAt" binding:"required"`
	UpdatedAt   *time.Time            `json:"updatedAt,omitempty"`
	Categories  []GetCategoryResponse `json:"categories,omitempty"`
}

type GetAllToDoResponse struct {
	ToDos []GetToDoResponse `json:"ToDos" binding:"required"`
}

type CreateToDoResponse struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type RemoveCategoriesFromToDoRequest struct {
	CategoriesID []uuid.UUID `json:"categories_id" binding:"required"`
}

type AddCategoriesFromToDoRequest struct {
	CategoriesID []uuid.UUID `json:"categories_id" binding:"required"`
}
