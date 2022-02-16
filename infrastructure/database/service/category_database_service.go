package service

import (
	"context"
	model "event-bus-demo/domain/model"
	"event-bus-demo/infrastructure/database/mapper"
	"event-bus-demo/infrastructure/database/repository"
	"event-bus-demo/infrastructure/error"
	"github.com/google/uuid"
)

type CategoryDatabaseService interface {
	GetAllCategories() ([]model.Category, error.InfrastructureError)
	GetCategory(ID uuid.UUID) (model.Category, error.InfrastructureError)
	CreateCategory(category model.Category) error.InfrastructureError
	UpdateCategory(category model.Category) error.InfrastructureError
	DeleteCategory(ID uuid.UUID) error.InfrastructureError
}

type categoryDatabaseService struct {
	transactionalRepository repository.TransactionalRepository
	categoryRepository      repository.CategoryRepository
}

func NewCategoryDatabaseService(transactionalRepository repository.TransactionalRepository, categoryRepository repository.CategoryRepository) CategoryDatabaseService {
	return &categoryDatabaseService{
		transactionalRepository: transactionalRepository,
		categoryRepository:      categoryRepository,
	}
}

func (dbService *categoryDatabaseService) GetAllCategories() ([]model.Category, error.InfrastructureError) {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return nil, error.NewSQLError(err.Error())
	} else if entities, err := dbService.categoryRepository.FindCategoriesList(ctx, queries); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return nil, err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return nil, error.NewSQLError(err.Error())
	} else {
		return mapper.NewCategoriesListFromEntity(entities), nil
	}
}

func (dbService *categoryDatabaseService) GetCategory(ID uuid.UUID) (model.Category, error.InfrastructureError) {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return model.Category{}, error.NewSQLError(err.Error())
	} else if entity, err := dbService.categoryRepository.FindCategoryByID(ctx, queries, ID); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return model.Category{}, err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return model.Category{}, error.NewSQLError(err.Error())
	} else {
		return mapper.NewCategoryFromEntity(entity), nil
	}
}

func (dbService *categoryDatabaseService) CreateCategory(category model.Category) error.InfrastructureError {
	entity := mapper.NewCategoryEntityFromCategoryModel(category)
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.categoryRepository.CreateCategory(ctx, queries, entity); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *categoryDatabaseService) UpdateCategory(category model.Category) error.InfrastructureError {
	entity := mapper.NewCategoryEntityFromCategoryModel(category)
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.categoryRepository.UpdateCategoryName(ctx, queries, entity); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (dbService *categoryDatabaseService) DeleteCategory(ID uuid.UUID) error.InfrastructureError {
	ctx := context.Background()
	if queries, err := dbService.transactionalRepository.CreateNewTransaction(ctx); err != nil {
		return error.NewSQLError(err.Error())
	} else if err := dbService.categoryRepository.DeleteCategoryByID(ctx, queries, ID); err != nil {
		_ = dbService.transactionalRepository.RollbackTransaction(queries)
		return err
	} else if err = dbService.transactionalRepository.CommitTransaction(queries); err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}
