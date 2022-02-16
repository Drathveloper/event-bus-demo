package model

import "github.com/google/uuid"

const UserEventTopic = "USERS"

type GetUserByIDEvent struct {
	ID uuid.UUID
}

func (GetUserByIDEvent) GetTopic() string {
	return UserEventTopic
}

func (GetUserByIDEvent) GetName() string {
	return "CreateUserEvent"
}

type CreateUserEvent struct {
	ID       uuid.UUID
	Username string
	Password string
}

func (CreateUserEvent) GetTopic() string {
	return UserEventTopic
}

func (CreateUserEvent) GetName() string {
	return "CreateUserEvent"
}

type UpdateUserPasswordEvent struct {
	ID       uuid.UUID
	Password string
}

func (UpdateUserPasswordEvent) GetTopic() string {
	return UserEventTopic
}

func (UpdateUserPasswordEvent) GetName() string {
	return "CreateUserEvent"
}

type DeleteUserEvent struct {
	ID uuid.UUID
}

func (DeleteUserEvent) GetTopic() string {
	return UserEventTopic
}

func (DeleteUserEvent) GetName() string {
	return "CreateUserEvent"
}
