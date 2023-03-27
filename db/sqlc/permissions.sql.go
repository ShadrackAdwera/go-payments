// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: permissions.sql

package db

import (
	"context"
	"database/sql"
)

const createPermission = `-- name: CreatePermission :one
INSERT INTO permissions (
  name, description, role_id
) VALUES (
  $1, $2, $3
)
RETURNING id, name, description, role_id
`

type CreatePermissionParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	RoleID      int64  `json:"role_id"`
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, createPermission, arg.Name, arg.Description, arg.RoleID)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.RoleID,
	)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE 
FROM permissions 
WHERE id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePermission, id)
	return err
}

const getPermission = `-- name: GetPermission :one
SELECT id, name, description, role_id FROM permissions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPermission(ctx context.Context, id int64) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermission, id)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.RoleID,
	)
	return i, err
}

const getPermissions = `-- name: GetPermissions :many
SELECT id, name, description, role_id FROM permissions 
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetPermissionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPermissions(ctx context.Context, arg GetPermissionsParams) ([]Permission, error) {
	rows, err := q.db.QueryContext(ctx, getPermissions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.RoleID,
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

const updatePermission = `-- name: UpdatePermission :one
UPDATE permissions 
SET name = COALESCE($2,name),
description = COALESCE($3,description),
role_id = COALESCE($4,role_id)
WHERE id = $1
RETURNING id, name, description, role_id
`

type UpdatePermissionParams struct {
	ID          int64          `json:"id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	RoleID      sql.NullInt64  `json:"role_id"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, updatePermission,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.RoleID,
	)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.RoleID,
	)
	return i, err
}
