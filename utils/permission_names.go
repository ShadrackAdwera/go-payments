package utils

const (
	UsersCreate       = "users:create"
	UsersRead         = "users:read"
	UsersDelete       = "users:delete"
	ClientsRead       = "clients:read"
	ClientsCreate     = "clients:create"
	PermissionsRead   = "permissions:read"
	PermissionsCreate = "permissions:create"
	PaymentInitiator  = "payment:initiator"
	PaymentApprover   = "payment:approver"
	RequestsRead      = "requests:read"
	RequestsApprove   = "requests:approve"
)

type PermissionData struct {
	PermissionName        string `json:"permission_name"`
	PermissionDescription string `json:"permission_description"`
}
