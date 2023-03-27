-- name: CreateUserRole :one
INSERT INTO users_roles (
  user_id, role_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserRole :one
SELECT * FROM users_roles
WHERE id = $1 LIMIT 1;

-- name: GetUsersRoles :many
SELECT * FROM users_roles 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserRole :one
UPDATE users_roles 
SET role_id = COALESCE(sqlc.narg(role_id),role_id),
user_id = COALESCE(sqlc.narg(user_id),user_id)
WHERE id = $1
RETURNING *;

-- name: DeleteUserRole :exec
DELETE 
FROM users_roles 
WHERE id = $1;