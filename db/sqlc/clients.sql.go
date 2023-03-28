// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: clients.sql

package db

import (
	"context"
	"database/sql"
)

const createClient = `-- name: CreateClient :one
INSERT INTO clients (
  name, email, phone, account_number, preferred_payment_type, createdby_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, name, email, phone, account_number, preferred_payment_type, createdby_id
`

type CreateClientParams struct {
	Name                 string         `json:"name"`
	Email                string         `json:"email"`
	Phone                string         `json:"phone"`
	AccountNumber        sql.NullString `json:"account_number"`
	PreferredPaymentType PaymentTypes   `json:"preferred_payment_type"`
	CreatedbyID          string         `json:"createdby_id"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.AccountNumber,
		arg.PreferredPaymentType,
		arg.CreatedbyID,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.AccountNumber,
		&i.PreferredPaymentType,
		&i.CreatedbyID,
	)
	return i, err
}

const deleteClient = `-- name: DeleteClient :exec
DELETE 
FROM clients 
WHERE id = $1
`

func (q *Queries) DeleteClient(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteClient, id)
	return err
}

const getClient = `-- name: GetClient :one
SELECT id, name, email, phone, account_number, preferred_payment_type, createdby_id FROM clients
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, id int64) (Client, error) {
	row := q.db.QueryRowContext(ctx, getClient, id)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.AccountNumber,
		&i.PreferredPaymentType,
		&i.CreatedbyID,
	)
	return i, err
}

const getClients = `-- name: GetClients :many
SELECT id, name, email, phone, account_number, preferred_payment_type, createdby_id FROM clients 
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetClientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetClients(ctx context.Context, arg GetClientsParams) ([]Client, error) {
	rows, err := q.db.QueryContext(ctx, getClients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.AccountNumber,
			&i.PreferredPaymentType,
			&i.CreatedbyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClient = `-- name: UpdateClient :one
UPDATE clients 
SET
  name = COALESCE($1,name),
  email = COALESCE($2,email),
  phone = COALESCE($3,phone),
  account_number = COALESCE($4,account_number),
  preferred_payment_type = COALESCE($5,preferred_payment_type)
WHERE id = $6
RETURNING id, name, email, phone, account_number, preferred_payment_type, createdby_id
`

type UpdateClientParams struct {
	Name                 sql.NullString   `json:"name"`
	Email                sql.NullString   `json:"email"`
	Phone                sql.NullString   `json:"phone"`
	AccountNumber        sql.NullString   `json:"account_number"`
	PreferredPaymentType NullPaymentTypes `json:"preferred_payment_type"`
	ID                   int64            `json:"id"`
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, updateClient,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.AccountNumber,
		arg.PreferredPaymentType,
		arg.ID,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.AccountNumber,
		&i.PreferredPaymentType,
		&i.CreatedbyID,
	)
	return i, err
}
