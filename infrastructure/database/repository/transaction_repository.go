package repository

import (
	"context"
	"database/sql"
	"event-bus-demo/infrastructure/database/sqlc"
	"event-bus-demo/infrastructure/error"
	"go.uber.org/zap"
)

type TransactionalRepository interface {
	CreateNewTransaction(ctx context.Context) (*sqlc.Queries, error.InfrastructureError)
	CommitTransaction(queries *sqlc.Queries) error.InfrastructureError
	RollbackTransaction(queries *sqlc.Queries) error.InfrastructureError
}

type transactionalRepository struct {
	db       *sql.DB
	logger   *zap.Logger
	registry map[*sqlc.Queries]*sql.Tx
}

func NewTransactionalRepository(logger *zap.Logger, db *sql.DB) TransactionalRepository {
	return &transactionalRepository{
		logger:   logger,
		db:       db,
		registry: make(map[*sqlc.Queries]*sql.Tx, 0),
	}
}

func (repo *transactionalRepository) CreateNewTransaction(ctx context.Context) (*sqlc.Queries, error.InfrastructureError) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, error.NewSQLError(err.Error())
	}
	queries := sqlc.New(repo.db).WithTx(tx)
	repo.registry[queries] = tx
	return queries, nil
}

func (repo *transactionalRepository) CommitTransaction(queries *sqlc.Queries) error.InfrastructureError {
	tx := repo.registry[queries]
	if tx == nil {
		return error.NewSQLError("error while getting transaction. Requested transaction not found")
	}
	err := tx.Commit()
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}

func (repo *transactionalRepository) RollbackTransaction(queries *sqlc.Queries) error.InfrastructureError {
	tx := repo.registry[queries]
	if tx == nil {
		return error.NewSQLError("error while getting transaction. Requested transaction not found")
	}
	err := tx.Rollback()
	if err != nil {
		return error.NewSQLError(err.Error())
	}
	return nil
}
