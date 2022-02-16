-- name: GetToDoById :one
SELECT * FROM todos WHERE id = $1;
-- name: GetToDoList :many
SELECT * FROM todos;
-- name: GetToDoCategories :many
SELECT c.id, c.name FROM todo_category tc INNER JOIN categories c on c.id = tc.category_id WHERE todo_id = $1;
-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1;
-- name: GetCategoriesList :many
SELECT * FROM categories;
-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;
-- name: CreateToDo :exec
INSERT INTO todos (id, title, description, created_at) VALUES ($1, $2, $3, $4);
-- name: UpdateToDoInformation :exec
UPDATE todos SET title = $2, description = $3, updated_at = $4 WHERE id = $1;
-- name: RemoveToDoFromCategory :exec
DELETE FROM todo_category WHERE todo_id = $1 AND category_id = $2;
-- name: RemoveAllCategoriesFromToDo :exec
DELETE FROM todo_category WHERE todo_id = $1;
-- name: AddToDoCategory :exec
INSERT INTO todo_category (todo_id, category_id) VALUES ($1, $2);
-- name: DeleteToDo :exec
DELETE FROM todos WHERE id = $1;
-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES ($1, $2);
-- name: UpdateCategoryName :exec
UPDATE categories SET name = $2 WHERE id = $1;
-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;
-- name: CreateUser :exec
INSERT INTO users (id, username, password, role) VALUES ($1, $2, $3, $4);
-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1;
-- name: UpdateUserRole :exec
UPDATE users SET role = $2 WHERE id = $1;
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;