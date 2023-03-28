-- name: CreateClient :one
INSERT INTO clients (
  name, email, phone, account_number, preferred_payment_type, createdby_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetClient :one
SELECT * FROM clients
WHERE id = $1 LIMIT 1;

-- name: GetClients :many
SELECT * FROM clients 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateClient :one
UPDATE clients 
SET
  name = COALESCE(sqlc.narg(name),name),
  email = COALESCE(sqlc.narg(email),email),
  phone = COALESCE(sqlc.narg(phone),phone),
  account_number = COALESCE(sqlc.narg(account_number),account_number),
  preferred_payment_type = COALESCE(sqlc.narg(preferred_payment_type),preferred_payment_type)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteClient :exec
DELETE 
FROM clients 
WHERE id = $1;