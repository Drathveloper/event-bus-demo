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

type UserRepository interface {
	FindUserByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) (model.UserEntity, error.InfrastructureError)
	CreateUser(ctx context.Context, queries *sqlc.Queries, entity model.UserEntity) error.InfrastructureError
	DeleteUserByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) error.InfrastructureError
	UpdateUserPassword(ctx context.Context, queries *sqlc.Queries, entity model.UserEntity) error.InfrastructureError
}

type userRepository struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewUserRepository(logger *zap.Logger, db *sql.DB) UserRepository {
	return &userRepository{
		logger: logger,
		db:     db,
	}
}

func (repo *userRepository) FindUserByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) (model.UserEntity, error.InfrastructureError) {
	user, err := queries.GetUserById(ctx, ID)
	if err != nil {
		if err.Error() == constants.NotFoundErrorMessage {
			return model.UserEntity{}, error.NewItemNotFoundError(fmt.Sprintf("toDo item with ID %s not found", ID))
		}
		return model.UserEntity{}, error.NewSQLError(err.Error())
	}

	return mapper.NewUserEntityFromSQLModel(user), nil
}

func (repo *userRepository) CreateUser(ctx context.Context, queries *sqlc.Queries, entity model.UserEntity) error.InfrastructureError {
	err := queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:       entity.ID,
		Username: entity.Username,
		Password: entity.Password,
	})
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repo *userRepository) DeleteUserByID(ctx context.Context, queries *sqlc.Queries, ID uuid.UUID) error.InfrastructureError {
	err := queries.DeleteUser(ctx, ID)
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repo *userRepository) UpdateUserPassword(ctx context.Context, queries *sqlc.Queries, entity model.UserEntity) error.InfrastructureError {
	err := queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:       entity.ID,
		Password: entity.Password,
	})
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}
