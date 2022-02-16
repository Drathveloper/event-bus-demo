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
)

type CategoryRepository interface {
	FindCategoriesList(ctx context.Context, queries *sqlc.Queries) ([]model.CategoryEntity, error.InfrastructureError)
	FindCategoryByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) (model.CategoryEntity, error.InfrastructureError)
	CreateCategory(ctx context.Context, queries *sqlc.Queries, entity model.CategoryEntity) error.InfrastructureError
	DeleteCategoryByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) error.InfrastructureError
	UpdateCategoryName(ctx context.Context, queries *sqlc.Queries, entity model.CategoryEntity) error.InfrastructureError
}

type categoryRepository struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewCategoryRepository(logger *zap.Logger, db *sql.DB) CategoryRepository {
	return &categoryRepository{
		logger: logger,
		db:     db,
	}
}

func (repository *categoryRepository) FindCategoriesList(ctx context.Context, queries *sqlc.Queries) ([]model.CategoryEntity, error.InfrastructureError) {
	categories, err := queries.GetCategoriesList(ctx)
	if err != nil {
		return nil, error.NewSQLError(err.Error())
	}
	return mapper.NewCategoryEntityListFromSQLModelList(categories), nil
}

func (repository *categoryRepository) FindCategoryByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) (model.CategoryEntity, error.InfrastructureError) {
	category, err := queries.GetCategoryById(ctx, ID)
	if err != nil {
		if err.Error() == constants.NotFoundErrorMessage {
			return model.CategoryEntity{}, error.NewItemNotFoundError(fmt.Sprintf("category item with ID %s not found", ID))
		}
		return model.CategoryEntity{}, error.NewSQLError(err.Error())
	}
	return mapper.NewCategoryEntityFromSQLModel(category), nil
}

func (repository *categoryRepository) UpdateCategoryName(ctx context.Context, queries *sqlc.Queries, entity model.CategoryEntity) error.InfrastructureError {
	err := queries.UpdateCategoryName(ctx, sqlc.UpdateCategoryNameParams{
		ID:   entity.ID,
		Name: entity.Name,
	})
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repository *categoryRepository) CreateCategory(ctx context.Context, queries *sqlc.Queries, entity model.CategoryEntity) error.InfrastructureError {
	err := queries.CreateCategory(ctx, sqlc.CreateCategoryParams{
		ID:   entity.ID,
		Name: entity.Name,
	})
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repository *categoryRepository) DeleteCategoryByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) error.InfrastructureError {
	err := queries.DeleteCategory(ctx, ID)
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}
