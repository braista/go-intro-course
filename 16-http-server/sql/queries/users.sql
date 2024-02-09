-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetCurrentUser :one
SELECT * FROM users
WHERE api_key = $1;

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: UpdateUser :execrows
UPDATE users
SET name = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1
RETURNING *;
