-- name: CreateUserPayment :one
INSERT INTO user_payments (
  request_id, client_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserPayment :one
SELECT user_payments.id, user_payments.client_id, user_payments.request_id, clients.name, clients.preferred_payment_type, requests.title, requests.amount, requests.status 
FROM user_payments
JOIN clients ON user_payments.client_id = clients.id
JOIN requests ON user_payments.request_id = requests.id
WHERE user_payments.id = $1 LIMIT 1;

-- name: GetUserPayments :many
SELECT user_payments.id, user_payments.client_id, user_payments.request_id, clients.name, clients.preferred_payment_type, requests.title, requests.amount, requests.status 
FROM user_payments
JOIN clients ON user_payments.client_id = clients.id
JOIN requests ON user_payments.request_id = requests.id
ORDER BY user_payments.id
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