-- name: CreateUserPayment :one
INSERT INTO user_payments (
  request_id, client_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserPayment :one
SELECT * FROM user_payments
WHERE id = $1 LIMIT 1;

-- name: GetUserPayments :many
SELECT * FROM user_payments 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserPayment :one
UPDATE user_payments 
SET client_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUserPayment :exec
DELETE 
FROM user_payments 
WHERE id = $1;