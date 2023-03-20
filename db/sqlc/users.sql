-- name: CreateUser :one
INSERT INTO users (
  id, username, email, role
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users 
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUserRole :one
UPDATE users 
SET role = $2
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE 
FROM users 
WHERE id = $1;