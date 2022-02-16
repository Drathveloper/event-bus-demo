// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addToDoCategory = `-- name: AddToDoCategories :exec
INSERT INTO todo_category (todo_id, category_id) VALUES ($1, $2)
`

type AddToDoCategoryParams struct {
	TodoID     uuid.UUID
	CategoryID uuid.UUID
}

func (q *Queries) AddToDoCategory(ctx context.Context, arg AddToDoCategoryParams) error {
	_, err := q.db.ExecContext(ctx, addToDoCategory, arg.TodoID, arg.CategoryID)
	return err
}

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES ($1, $2)
`

type CreateCategoryParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name)
	return err
}

const createToDo = `-- name: CreateToDo :exec
INSERT INTO todos (id, title, description, created_at) VALUES ($1, $2, $3, $4)
`

type CreateToDoParams struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
}

func (q *Queries) CreateToDo(ctx context.Context, arg CreateToDoParams) error {
	_, err := q.db.ExecContext(ctx, createToDo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.CreatedAt,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, username, password, role) VALUES ($1, $2, $3, $4)
`

type CreateUserParams struct {
	ID       uuid.UUID
	Username string
	Password string
	Role     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Role,
	)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const deleteToDo = `-- name: DeleteToDo :exec
DELETE FROM todos WHERE id = $1
`

func (q *Queries) DeleteToDo(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteToDo, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getCategoriesList = `-- name: GetCategoriesList :many
SELECT id, name FROM categories
`

func (q *Queries) GetCategoriesList(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name FROM categories WHERE id = $1
`

func (q *Queries) GetCategoryById(ctx context.Context, id uuid.UUID) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryById, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getToDoById = `-- name: GetToDoById :one
SELECT id, title, description, created_at, updated_at FROM todos WHERE id = $1
`

func (q *Queries) GetToDoById(ctx context.Context, id uuid.UUID) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getToDoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getToDoCategories = `-- name: GetToDoCategories :many
SELECT c.id, c.name FROM todo_category tc INNER JOIN categories c on c.id = tc.category_id WHERE todo_id = $1
`

func (q *Queries) GetToDoCategories(ctx context.Context, todoID uuid.UUID) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getToDoCategories, todoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getToDoList = `-- name: GetToDoList :many
SELECT id, title, description, created_at, updated_at FROM todos
`

func (q *Queries) GetToDoList(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getToDoList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, password, role FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, role FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const removeAllCategoriesFromToDo = `-- name: RemoveAllCategoriesFromToDo :exec
DELETE FROM todo_category WHERE todo_id = $1
`

func (q *Queries) RemoveAllCategoriesFromToDo(ctx context.Context, todoID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, removeAllCategoriesFromToDo, todoID)
	return err
}

const removeToDoFromCategory = `-- name: RemoveToDoFromCategory :exec
DELETE FROM todo_category WHERE todo_id = $1 AND category_id = $2
`

type RemoveToDoFromCategoryParams struct {
	TodoID     uuid.UUID
	CategoryID uuid.UUID
}

func (q *Queries) RemoveToDoFromCategory(ctx context.Context, arg RemoveToDoFromCategoryParams) error {
	_, err := q.db.ExecContext(ctx, removeToDoFromCategory, arg.TodoID, arg.CategoryID)
	return err
}

const updateCategoryName = `-- name: UpdateCategoryName :exec
UPDATE categories SET name = $2 WHERE id = $1
`

type UpdateCategoryNameParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) UpdateCategoryName(ctx context.Context, arg UpdateCategoryNameParams) error {
	_, err := q.db.ExecContext(ctx, updateCategoryName, arg.ID, arg.Name)
	return err
}

const updateToDoInformation = `-- name: UpdateToDoInformation :exec
UPDATE todos SET title = $2, description = $3, updated_at = $4 WHERE id = $1
`

type UpdateToDoInformationParams struct {
	ID          uuid.UUID
	Title       string
	Description string
	UpdatedAt   sql.NullTime
}

func (q *Queries) UpdateToDoInformation(ctx context.Context, arg UpdateToDoInformationParams) error {
	_, err := q.db.ExecContext(ctx, updateToDoInformation,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.UpdatedAt,
	)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID       uuid.UUID
	Password string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.ID, arg.Password)
	return err
}

const updateUserRole = `-- name: UpdateUserRole :exec
UPDATE users SET role = $2 WHERE id = $1
`

type UpdateUserRoleParams struct {
	ID   uuid.UUID
	Role string
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateUserRole, arg.ID, arg.Role)
	return err
}