-- name: ListUsers :many
SELECT * FROM commerce_user.users;

-- name: GetUser :one
SELECT * FROM commerce_user.users WHERE id = ? ;

-- name: CreateUser :execresult
INSERT INTO commerce_user.users (name, username, email, password, birth_date, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE commerce_user.users
SET name = ?, username = ?, email = ?, birth_date = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteUser :exec
UPDATE commerce_user.users
SET deleted_at = ?
WHERE id = ?;