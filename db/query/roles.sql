-- name: CreateRole :one
INSERT INTO roles (
  id, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: GetRoles :many
SELECT * FROM roles 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRole :one
UPDATE roles 
SET name = COALESCE(sqlc.narg(name),name)
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE 
FROM roles 
WHERE id = $1;