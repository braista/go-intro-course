// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, name
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1
RETURNING id, created_at, updated_at, name
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteUser, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getUsers = `-- name: GetUsers :many
SELECT id, created_at, updated_at, name FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
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

const updateUser = `-- name: UpdateUser :execrows
UPDATE users
SET name = $2, updated_at = $3
WHERE id = $1
RETURNING id, created_at, updated_at, name
`

type UpdateUserParams struct {
	ID        uuid.UUID
	Name      string
	UpdatedAt time.Time
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updateUser, arg.ID, arg.Name, arg.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}