-- name: ListUsers :many
SELECT * FROM commerce_ai.users;

-- name: GetUser :one
SELECT * FROM commerce_ai.users WHERE id = ? ;

-- name: CreateUser :execresult
INSERT INTO commerce_ai.users (name, username, email, password, birth_date, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE commerce_ai.users
SET name = ?, username = ?, email = ?, birth_date = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteUser :exec
UPDATE commerce_ai.users
SET deleted_at = ?
WHERE id = ?;

-- name: CreateAddress :execresult
INSERT INTO commerce_ai.addresses (street, number, zip_code, city, country, state, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListAddresses :many
SELECT * FROM commerce_ai.addresses WHERE user_id = ? AND deleted_at IS NULL;

-- name: GetAddress :one
SELECT * FROM commerce_ai.addresses WHERE id = ? AND deleted_at IS NULL;

-- name: UpdateAddress :exec
UPDATE commerce_ai.addresses
SET street = ?, number = ?, zip_code = ?, city = ?, country = ?, state = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteAddress :exec
UPDATE commerce_ai.addresses
SET deleted_at = ?
WHERE id = ?;