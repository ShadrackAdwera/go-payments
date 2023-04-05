-- name: CreateDarajaToken :one
INSERT INTO daraja_token (
  access_token, expires_at
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetDarajaToken :one
SELECT * 
FROM daraja_token 
LIMIT 1;

-- name: DeleteDarajaToken :exec
DELETE 
FROM daraja_token 
WHERE id = $1;