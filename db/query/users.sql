-- name: CreateUser :one
INSERT INTO users (
  id, username
) VALUES (
  $1, $2
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

-- name: UpdateUser :one
UPDATE users 
SET username = COALESCE(sqlc.narg(username),username)
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE 
FROM users 
WHERE id = $1;