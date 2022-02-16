package service

import (
	"context"
	"event-bus-demo/domain/model"
	"event-bus-demo/infrastructure/database/mapper"
	"event-bus-demo/infrastructure/database/repository"
	"event-bus-demo/infrastructure/error"
	"github.com/google/uuid"
)

type UserDatabaseService interface {
	GetUser(ID uuid.UUID) (model.User, error.InfrastructureError)
	CreateUser(user model.User) error.InfrastructureError
	UpdateUserPassword(user model.User) error.InfrastructureError
	DeleteUser(ID uuid.UUID) error.InfrastructureError
}

type userDatabaseService struct {
	transactionalRepository repository.TransactionalRepository
	userRepository          repository.UserRepository
}

func NewUserDatabaseService(transactionalRepository repository.TransactionalRepository, userRepository repository.UserRepository) UserDatabaseService {
	return &userDatabaseService{
		transactionalRepository: transactionalRepository,
		userRepository:          userRepository,
	}
}

func (dbService *userDatabaseService) GetUser(ID uuid.UUID) (model.User, error.InfrastructureError) {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return model.User{}, error.NewSQLError(err.Error())
	} else if entity, err := dbService.userRepository.FindUserByID(ctx, queries, ID); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return model.User{}, err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return model.User{}, error.NewSQLError(err.Error())
	} else {
		return mapper.NewUserFromEntity(entity), nil
	}
}

func (dbService *userDatabaseService) CreateUser(user model.User) error.InfrastructureError {
	ctx := context.Background()
	entity := mapper.NewUserEntityFromUserModel(user)
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.userRepository.CreateUser(ctx, queries, entity); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *userDatabaseService) UpdateUserPassword(user model.User) error.InfrastructureError {
	ctx := context.Background()
	entity := mapper.NewUserEntityFromUserModel(user)
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.userRepository.UpdateUserPassword(ctx, queries, entity); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *userDatabaseService) DeleteUser(ID uuid.UUID) error.InfrastructureError {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.userRepository.DeleteUserByID(ctx, queries, ID); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}
