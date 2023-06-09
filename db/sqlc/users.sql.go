// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id, username
) VALUES (
  $1, $2
)
RETURNING id, username
`

type CreateUserParams struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Username)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE 
FROM users 
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, username FROM users 
ORDER BY username
LIMIT $1
OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Username); err != nil {
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

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET username = COALESCE($2,username)
WHERE id = $1
RETURNING id, username
`

type UpdateUserParams struct {
	ID       string         `json:"id"`
	Username sql.NullString `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Username)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}
