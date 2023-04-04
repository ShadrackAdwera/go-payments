-- name: CreateRequest :one
INSERT INTO requests (
  title, status, amount, paid_to_id, createdby_id, approvedby_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetRequest :one
SELECT * FROM requests
WHERE id = $1 LIMIT 1;

-- name: GetRequests :many
SELECT * FROM requests 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetRequestsToApprove :many
SELECT * FROM requests 
WHERE status = $1
AND approvedby_id = $2
ORDER BY created_at DESC;

-- name: UpdateRequest :one
UPDATE requests 
SET
  title = COALESCE(sqlc.narg(title),title),
  status = COALESCE(sqlc.narg(status),status),
  amount = COALESCE(sqlc.narg(amount),amount),
  paid_to_id = COALESCE(sqlc.narg(paid_to_id),paid_to_id),
  approvedby_id = COALESCE(sqlc.narg(approvedby_id),approvedby_id),
  approved_at = COALESCE(sqlc.narg(approved_at),approved_at)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteRequest :exec
DELETE 
FROM requests 
WHERE id = $1;