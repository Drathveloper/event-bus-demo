package service

import (
	"context"
	"event-bus-demo/domain/model"
	"event-bus-demo/infrastructure/database/mapper"
	"event-bus-demo/infrastructure/database/repository"
	"event-bus-demo/infrastructure/error"
	"github.com/google/uuid"
)

type ToDoDatabaseService interface {
	GetAllToDoList() ([]model.ToDo, error.InfrastructureError)
	GetToDo(ID uuid.UUID) (model.ToDo, error.InfrastructureError)
	CreateToDo(toDo model.ToDo) error.InfrastructureError
	UpdateToDo(toDo model.ToDo) error.InfrastructureError
	DeleteToDo(ID uuid.UUID) error.InfrastructureError
	AddCategoriesIntoToDo(event model.AddCategoriesFromToDoEvent) error.InfrastructureError
	RemoveCategoriesFromToDo(event model.RemoveCategoriesFromToDoEvent) error.InfrastructureError
}

type toDoDatabaseService struct {
	transactionalRepository repository.TransactionalRepository
	toDoRepository          repository.ToDoRepository
}

func NewToDoDatabaseService(transactionalRepository repository.TransactionalRepository, toDoRepository repository.ToDoRepository) ToDoDatabaseService {
	return &toDoDatabaseService{
		transactionalRepository: transactionalRepository,
		toDoRepository:          toDoRepository,
	}
}

func (dbService *toDoDatabaseService) GetAllToDoList() ([]model.ToDo, error.InfrastructureError) {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return nil, error.NewSQLError(err.Error())
	} else if entities, err := dbService.toDoRepository.FindToDoList(ctx, queries); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return nil, err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return nil, error.NewSQLError(err.Error())
	} else {
		return mapper.NewToDoListFromEntityList(entities), nil
	}
}

func (dbService *toDoDatabaseService) GetToDo(ID uuid.UUID) (model.ToDo, error.InfrastructureError) {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return model.ToDo{}, error.NewSQLError(err.Error())
	} else if entity, err := dbService.toDoRepository.FindToDoByID(ctx, queries, ID); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return model.ToDo{}, err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return model.ToDo{}, error.NewSQLError(err.Error())
	} else {
		return mapper.NewToDoFromEntity(entity), nil
	}
}

func (dbService *toDoDatabaseService) CreateToDo(toDo model.ToDo) error.InfrastructureError {
	entity := mapper.NewToDoEntityFromToDoModel(toDo)
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.toDoRepository.CreateToDo(ctx, queries, entity); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *toDoDatabaseService) UpdateToDo(toDo model.ToDo) error.InfrastructureError {
	entity := mapper.NewToDoEntityFromToDoModel(toDo)
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.toDoRepository.UpdateToDoInformation(ctx, queries, entity); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *toDoDatabaseService) DeleteToDo(ID uuid.UUID) error.InfrastructureError {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.toDoRepository.DeleteToDoByID(ctx, queries, ID); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *toDoDatabaseService) AddCategoriesIntoToDo(event model.AddCategoriesFromToDoEvent) error.InfrastructureError {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.toDoRepository.AddToDoCategories(ctx, queries, event.ToDoID, event.Categories); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *toDoDatabaseService) RemoveCategoriesFromToDo(event model.RemoveCategoriesFromToDoEvent) error.InfrastructureError {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.toDoRepository.DeleteToDoCategories(ctx, queries, event.ToDoID, event.Categories); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}
