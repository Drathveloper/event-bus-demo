package repository

import (
	"context"
	"database/sql"
	"event-bus-demo/infrastructure/constants"
	"event-bus-demo/infrastructure/database/mapper"
	"event-bus-demo/infrastructure/database/model"
	"event-bus-demo/infrastructure/database/sqlc"
	"event-bus-demo/infrastructure/error"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type ToDoRepository interface {
	FindToDoList(ctx context.Context, queries *sqlc.Queries) ([]model.ToDoEntity, error.InfrastructureError)
	FindToDoByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) (model.ToDoEntity, error.InfrastructureError)
	CreateToDo(ctx context.Context, queries *sqlc.Queries, entity model.ToDoEntity) error.InfrastructureError
	DeleteToDoByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) error.InfrastructureError
	UpdateToDoInformation(ctx context.Context, queries *sqlc.Queries, entity model.ToDoEntity) error.InfrastructureError
	DeleteToDoCategories(ctx context.Context, queries *sqlc.Queries, toDoID uuid.UUID, categoriesID []uuid.UUID) error.InfrastructureError
	AddToDoCategories(ctx context.Context, queries *sqlc.Queries, toDoID uuid.UUID, categoriesID []uuid.UUID) error.InfrastructureError
}

type toDoRepository struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewToDoRepository(logger *zap.Logger, db *sql.DB) ToDoRepository {
	return &toDoRepository{
		logger: logger,
		db:     db,
	}
}

func (repository *toDoRepository) FindToDoList(ctx context.Context, queries *sqlc.Queries) ([]model.ToDoEntity, error.InfrastructureError) {
	toDoList, err := queries.GetToDoList(ctx)
	if err != nil {
		return nil, error.NewSQLError(err.Error())
	}
	if toDoList == nil {
		return make([]model.ToDoEntity, 0), nil
	}
	entityList := mapper.NewToDoEntityListFromSQLModelList(toDoList)
	for _, entity := range entityList {
		categories, err := queries.GetToDoCategories(ctx, entity.ID)
		if err != nil {
			return nil, error.NewSQLError(err.Error())
		}
		entity.Categories = mapper.NewCategoryEntityListFromSQLModelList(categories)
	}
	return entityList, nil
}

func (repository *toDoRepository) FindToDoByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) (model.ToDoEntity, error.InfrastructureError) {
	toDo, err := queries.GetToDoById(ctx, ID)
	if err != nil {
		if err.Error() == constants.NotFoundErrorMessage {
			return model.ToDoEntity{}, error.NewItemNotFoundError(fmt.Sprintf("toDo item with ID %s not found", ID))
		}
		return model.ToDoEntity{}, error.NewSQLError(err.Error())
	}
	entity := mapper.NewToDoEntityFromSQLModel(toDo)
	categories, err := queries.GetToDoCategories(ctx, entity.ID)
	if err != nil {
		return model.ToDoEntity{}, error.NewSQLError(err.Error())
	}
	entity.Categories = mapper.NewCategoryEntityListFromSQLModelList(categories)
	return entity, nil
}

func (repository *toDoRepository) CreateToDo(ctx context.Context, queries *sqlc.Queries, entity model.ToDoEntity) error.InfrastructureError {
	err := queries.CreateToDo(ctx, sqlc.CreateToDoParams{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		CreatedAt:   *entity.CreatedAt,
	})
	for _, category := range entity.Categories {
		err := queries.AddToDoCategory(ctx, sqlc.AddToDoCategoryParams{
			TodoID:     entity.ID,
			CategoryID: category.ID,
		})
		if err != nil {
			return error.NewSQLError(err.Error())
		}
	}
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repository *toDoRepository) DeleteToDoByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) error.InfrastructureError {
	err := queries.RemoveAllCategoriesFromToDo(ctx, ID)
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	err = queries.DeleteToDo(ctx, ID)
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repository *toDoRepository) UpdateToDoInformation(ctx context.Context, queries *sqlc.Queries, entity model.ToDoEntity) error.InfrastructureError {
	err := queries.UpdateToDoInformation(ctx, sqlc.UpdateToDoInformationParams{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repository *toDoRepository) DeleteToDoCategories(ctx context.Context, queries *sqlc.Queries, toDoID uuid.UUID, categories []uuid.UUID) error.InfrastructureError {
	for _, categoryID := range categories {
		err := queries.RemoveToDoFromCategory(ctx, sqlc.RemoveToDoFromCategoryParams{
			TodoID:     toDoID,
			CategoryID: categoryID,
		})
		if err != nil {
			return error.NewSQLError(err.Error())
		}
	}
	return nil
}

func (repository *toDoRepository) AddToDoCategories(ctx context.Context, queries *sqlc.Queries, toDoID uuid.UUID, categories []uuid.UUID) error.InfrastructureError {
	for _, categoryID := range categories {
		err := queries.AddToDoCategory(ctx, sqlc.AddToDoCategoryParams{
			TodoID:     toDoID,
			CategoryID: categoryID,
		})
		if err != nil {
			return error.NewSQLError(err.Error())
		}
	}
	return nil
}
