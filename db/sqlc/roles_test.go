package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomRole(t *testing.T) Role {
	name := utils.RandomString(10)
	user := CreateRandomUser(t)

	role, err := testQuery.CreateRole(context.Background(), CreateRoleParams{
		Name: name,
		CreatedbyID: sql.NullString{
			String: user.ID,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotZero(t, role.ID)
	require.Equal(t, role.Name, name)

	return role
}

func TestCreateRole(t *testing.T) {
	CreateRandomRole(t)
}

func TestGetRole(t *testing.T) {
	role := CreateRandomRole(t)

	foudRole, err := testQuery.GetRole(context.Background(), role.ID)

	require.NoError(t, err)
	require.NotEmpty(t, foudRole)
	require.Equal(t, role.ID, foudRole.ID)
	require.Equal(t, role.Name, foudRole.Name)
}

func TestGetRoles(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		CreateRandomRole(t)
	}

	roles, err := testQuery.GetRoles(context.Background(), GetRolesParams{
		Limit:  5,
		Offset: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, roles)
	require.Len(t, roles, n)
}

func TestUpdateRole(t *testing.T) {
	name := utils.RandomString(12)
	role := CreateRandomRole(t)

	updatedRole, err := testQuery.UpdateRole(context.Background(), UpdateRoleParams{
		ID: role.ID,
		Name: sql.NullString{
			String: name,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.Equal(t, role.ID, updatedRole.ID)
	require.Equal(t, updatedRole.Name, name)
	require.NotEqual(t, updatedRole.Name, role.Name)
}
