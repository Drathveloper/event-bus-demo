package mapper

import (
	"event-bus-demo/application/dto"
	domainModel "event-bus-demo/domain/model"
)

func NewGetUserResponseFromDomainModel(user domainModel.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}
