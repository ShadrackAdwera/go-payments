package utils

func GetPermissionData() []PermissionData {
	p := []PermissionData{
		{
			PermissionName:        UsersCreate,
			PermissionDescription: "This permission allows creating users in the system",
		},
		{
			PermissionName:        UsersRead,
			PermissionDescription: "This permission allows reading users in the system",
		},
		{
			PermissionName:        UsersDelete,
			PermissionDescription: "This permission allows deleting users in the system",
		},
		{
			PermissionName:        ClientsRead,
			PermissionDescription: "This permission allows reading clients in the system",
		},
		{
			PermissionName:        ClientsCreate,
			PermissionDescription: "This permission allows creating clients in the system",
		},
		{
			PermissionName:        PermissionsRead,
			PermissionDescription: "This permission allows reading permissions in the system",
		},
		{
			PermissionName:        PermissionsCreate,
			PermissionDescription: "This permission allows creating permissions in the system",
		},
		{
			PermissionName:        RequestsCreate,
			PermissionDescription: "This permission creates payment requests",
		},
		{
			PermissionName:        RequestsRead,
			PermissionDescription: "This permission allows reading requests in the system",
		},
		{
			PermissionName:        RequestsApprove,
			PermissionDescription: "This permission allows approving requests in the system",
		},
		{
			PermissionName:        UserPaymentsRead,
			PermissionDescription: "This permission allows reading user payments in the system",
		},
	}

	return p
}
