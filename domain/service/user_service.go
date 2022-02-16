package service

import (
	"event-bus-demo/application/dto"
	"event-bus-demo/domain/error"
	"event-bus-demo/domain/mapper"
	"event-bus-demo/domain/model"
	"event-bus-demo/infrastructure/database/service"
	"go.uber.org/zap"
)

type UserReadService interface {
	GetUserByID(event model.GetUserByIDEvent) (dto.GetUserResponse, error.DomainError)
}

type UserWriteService interface {
	AddUser(event model.CreateUserEvent) error.DomainError
	UpdateUserPassword(event model.UpdateUserPasswordEvent) error.DomainError
	DeleteUser(event model.DeleteUserEvent) error.DomainError
}

type userReadService struct {
	logger              *zap.Logger
	userDatabaseService service.UserDatabaseService
	domainAdvice        error.DomainAdvice
}

type userWriteService struct {
	logger              *zap.Logger
	userDatabaseService service.UserDatabaseService
	domainAdvice        error.DomainAdvice
}

func NewUserReadService(databaseService service.UserDatabaseService, advice error.DomainAdvice, logger *zap.Logger) UserReadService {
	return &userReadService{
		logger:              logger,
		userDatabaseService: databaseService,
		domainAdvice:        advice,
	}
}

func NewUserWriteService(databaseService service.UserDatabaseService, advice error.DomainAdvice, logger *zap.Logger) UserWriteService {
	return &userWriteService{
		logger:              logger,
		userDatabaseService: databaseService,
		domainAdvice:        advice,
	}
}

func (service *userReadService) GetUserByID(event model.GetUserByIDEvent) (dto.GetUserResponse, error.DomainError) {
	user, err := service.userDatabaseService.GetUser(event.ID)
	if err != nil {
		return dto.GetUserResponse{}, service.domainAdvice.TranslateError(err)
	}
	return mapper.NewGetUserResponseFromDomainModel(user), nil
}

func (service *userWriteService) AddUser(event model.CreateUserEvent) error.DomainError {
	user := model.User{
		ID:       event.ID,
		Username: event.Username,
	}
	if err := user.SetNonHashedPassword(event.Password); err != nil {
		return service.domainAdvice.TranslateError(err)
	} else if err := service.userDatabaseService.CreateUser(user); err != nil {
		return service.domainAdvice.TranslateError(err)
	}
	return nil
}

func (service *userWriteService) UpdateUserPassword(event model.UpdateUserPasswordEvent) error.DomainError {
	user := model.User{
		ID: event.ID,
	}
	if err := user.SetNonHashedPassword(event.Password); err != nil {
		return service.domainAdvice.TranslateError(err)
	} else if err := service.userDatabaseService.UpdateUserPassword(user); err != nil {
		return service.domainAdvice.TranslateError(err)
	}
	return nil
}

func (service *userWriteService) DeleteUser(event model.DeleteUserEvent) error.DomainError {
	err := service.userDatabaseService.DeleteUser(event.ID)
	if err != nil {
		return service.domainAdvice.TranslateError(err)
	}
	return nil
}
