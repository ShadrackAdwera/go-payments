-- name: CreatePermission :one
INSERT INTO permissions (
  name, description, createdby_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPermission :one
SELECT * FROM permissions
WHERE id = $1 LIMIT 1;

-- name: GetPermissionByName :one
SELECT * FROM permissions
WHERE name = $1 LIMIT 1;

-- name: GetPermissions :many
SELECT * FROM permissions 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePermission :one
UPDATE permissions 
SET name = COALESCE(sqlc.narg(name),name),
description = COALESCE(sqlc.narg(description),description)
WHERE id = $1
RETURNING *;

-- name: DeletePermission :exec
DELETE 
FROM permissions 
WHERE id = $1;