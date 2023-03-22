// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateClient(ctx context.Context, arg CreateClientParams) (Client, error)
	CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserPayment(ctx context.Context, arg CreateUserPaymentParams) (UserPayment, error)
	DeleteClient(ctx context.Context, id int64) error
	DeleteRequest(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	DeleteUserPayment(ctx context.Context, id int64) error
	GetClient(ctx context.Context, id int64) (Client, error)
	GetClients(ctx context.Context, arg GetClientsParams) ([]Client, error)
	GetRequest(ctx context.Context, id int64) (Request, error)
	GetRequests(ctx context.Context, arg GetRequestsParams) ([]Request, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserPayment(ctx context.Context, id int64) (UserPayment, error)
	GetUserPayments(ctx context.Context, arg GetUserPaymentsParams) ([]UserPayment, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	UpdateClient(ctx context.Context, arg UpdateClientParams) (Client, error)
	UpdateRequest(ctx context.Context, arg UpdateRequestParams) (Request, error)
	UpdateUserPayment(ctx context.Context, arg UpdateUserPaymentParams) (UserPayment, error)
	UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error)
}

var _ Querier = (*Queries)(nil)