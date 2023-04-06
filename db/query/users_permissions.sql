-- name: CreateUserPermission :one
INSERT INTO users_permissions (
  user_id, permission_id, createdby_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserPermission :one
SELECT * FROM users_permissions
WHERE id = $1 LIMIT 1;

-- name: CreateUserPermissions :exec
INSERT INTO users_permissions 
(user_id, permission_id, createdby_id) 
VALUES (UNNEST(@user_id::varchar[]), UNNEST(@permission_id::BIGINT[]),UNNEST(@createdby_id::varchar[]));

-- name: GetUserPermissionByUserIdAndPermissionId :one
SELECT * 
FROM users_permissions
WHERE user_id = $1 
AND permission_id = $2 
LIMIT 1;

-- name: GetUsersPermissions :many
SELECT * FROM users_permissions 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserPermission :one
UPDATE users_permissions 
SET permission_id = COALESCE(sqlc.narg(permission_id),permission_id),
user_id = COALESCE(sqlc.narg(user_id),user_id)
WHERE id = $1
RETURNING *;

-- name: DeleteUserPermission :exec
DELETE 
FROM users_permissions 
WHERE id = $1;