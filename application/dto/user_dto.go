package dto

import (
	"event-bus-demo/domain/model"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c CreateUserRequest) ToEvent(ID uuid.UUID) model.CreateUserEvent {
	return model.CreateUserEvent{
		ID:       ID,
		Username: c.Username,
		Password: c.Password,
	}
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

func (c UpdateUserPasswordRequest) ToEvent(ID uuid.UUID) model.UpdateUserPasswordEvent {
	return model.UpdateUserPasswordEvent{
		ID:       ID,
		Password: c.Password,
	}
}

type GetUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}
