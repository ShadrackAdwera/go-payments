package utils

const (
	UsersCreate       = "users:create"
	UsersRead         = "users:read"
	UsersDelete       = "users:delete"
	ClientsRead       = "clients:read"
	ClientsCreate     = "clients:create"
	PermissionsRead   = "permissions:read"
	PermissionsCreate = "permissions:create"
	RequestsCreate    = "requests:create"
	RequestsRead      = "requests:read"
	RequestsApprove   = "requests:approve"
	UserPaymentsRead  = "user:payments:read"
)

type PermissionData struct {
	PermissionName        string `json:"permission_name"`
	PermissionDescription string `json:"permission_description"`
}
