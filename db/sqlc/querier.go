// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateClient(ctx context.Context, arg CreateClientParams) (Client, error)
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error)
	CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserPayment(ctx context.Context, arg CreateUserPaymentParams) (UserPayment, error)
	CreateUserPermission(ctx context.Context, arg CreateUserPermissionParams) (UsersPermission, error)
	DeleteClient(ctx context.Context, id int64) error
	DeletePermission(ctx context.Context, id int64) error
	DeleteRequest(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id string) error
	DeleteUserPayment(ctx context.Context, id int64) error
	DeleteUserPermission(ctx context.Context, id int64) error
	GetClient(ctx context.Context, id int64) (Client, error)
	GetClients(ctx context.Context, arg GetClientsParams) ([]Client, error)
	GetPermission(ctx context.Context, id int64) (Permission, error)
	GetPermissionByName(ctx context.Context, name string) (Permission, error)
	GetPermissions(ctx context.Context, arg GetPermissionsParams) ([]Permission, error)
	GetRequest(ctx context.Context, id int64) (Request, error)
	GetRequests(ctx context.Context, arg GetRequestsParams) ([]Request, error)
	GetRequestsToApprove(ctx context.Context, arg GetRequestsToApproveParams) ([]Request, error)
	GetUser(ctx context.Context, id string) (User, error)
	GetUserPayment(ctx context.Context, id int64) (GetUserPaymentRow, error)
	GetUserPayments(ctx context.Context, arg GetUserPaymentsParams) ([]GetUserPaymentsRow, error)
	GetUserPermission(ctx context.Context, id int64) (UsersPermission, error)
	GetUserPermissionByUserIdAndPermissionId(ctx context.Context, arg GetUserPermissionByUserIdAndPermissionIdParams) (UsersPermission, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	GetUsersPermissions(ctx context.Context, arg GetUsersPermissionsParams) ([]UsersPermission, error)
	UpdateClient(ctx context.Context, arg UpdateClientParams) (Client, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error)
	UpdateRequest(ctx context.Context, arg UpdateRequestParams) (Request, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserPayment(ctx context.Context, arg UpdateUserPaymentParams) (UserPayment, error)
	UpdateUserPermission(ctx context.Context, arg UpdateUserPermissionParams) (UsersPermission, error)
}

var _ Querier = (*Queries)(nil)
