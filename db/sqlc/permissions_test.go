package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomPermission(t *testing.T) Permission {
	user := CreateRandomUser(t)

	name := utils.RandomString(12)
	description := utils.RandomString(18)

	permission, err := testQuery.CreatePermission(context.Background(), CreatePermissionParams{
		Name:        name,
		Description: description,
		CreatedbyID: user.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, permission)
	require.Equal(t, permission.Name, name)
	require.Equal(t, permission.Description, description)

	return permission
}

func TestCreatePermission(t *testing.T) {
	CreateRandomPermission(t)
}

func TestGetPermission(t *testing.T) {
	p := CreateRandomPermission(t)

	perm, err := testQuery.GetPermission(context.Background(), p.ID)

	require.NoError(t, err)
	require.NotEmpty(t, perm)
	require.Equal(t, perm.Name, p.Name)
	require.Equal(t, perm.Description, p.Description)

	perm, err = testQuery.GetPermission(context.Background(), utils.RandomInteger(100, 1000))

	require.Error(t, err)
	require.Empty(t, perm)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestGetPermissionByName(t *testing.T) {
	p := CreateRandomPermission(t)

	perm, err := testQuery.GetPermissionByName(context.Background(), p.Name)

	require.NoError(t, err)
	require.NotEmpty(t, perm)
	require.Equal(t, perm.Name, p.Name)
	require.Equal(t, perm.Description, p.Description)

	perm, err = testQuery.GetPermissionByName(context.Background(), utils.RandomString(15))

	require.Error(t, err)
	require.Empty(t, perm)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestGetPermissions(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		CreateRandomPermission(t)
	}

	permissions, err := testQuery.GetPermissions(context.Background(), GetPermissionsParams{
		Limit:  5,
		Offset: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, permissions)
	require.Len(t, permissions, n)
}
