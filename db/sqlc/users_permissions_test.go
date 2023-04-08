package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUserPermission(t *testing.T) UsersPermission {
	perm := CreateRandomPermission(t)
	usr := CreateRandomUser(t)

	userPerm, err := testQuery.CreateUserPermission(context.Background(), CreateUserPermissionParams{
		UserID:       usr.ID,
		PermissionID: perm.ID,
		CreatedbyID:  usr.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, userPerm)
	require.Equal(t, userPerm.PermissionID, perm.ID)
	require.Equal(t, userPerm.UserID, usr.ID)

	return userPerm
}

func TestCreateUserPermission(t *testing.T) {
	CreateRandomUserPermission(t)
}

func TestGetUserPermissionByUserIdAndPermissionId(t *testing.T) {
	upm := CreateRandomUserPermission(t)

	uperm, err := testQuery.GetUserPermissionByUserIdAndPermissionId(context.Background(), GetUserPermissionByUserIdAndPermissionIdParams{
		UserID:       upm.UserID,
		PermissionID: upm.PermissionID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, uperm)
	require.Equal(t, uperm.PermissionID, upm.PermissionID)
	require.Equal(t, uperm.UserID, upm.UserID)

	uperm, err = testQuery.GetUserPermissionByUserIdAndPermissionId(context.Background(), GetUserPermissionByUserIdAndPermissionIdParams{
		UserID:       utils.RandomString(18),
		PermissionID: utils.RandomInteger(100, 500),
	})

	require.Error(t, err)
	require.Empty(t, uperm)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestGetUsersPermissions(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		perm := CreateRandomPermission(t)
		usr := CreateRandomUser(t)
		_, _ = testQuery.CreateUserPermission(context.Background(), CreateUserPermissionParams{
			UserID:       usr.ID,
			PermissionID: perm.ID,
			CreatedbyID:  usr.ID,
		})
	}

	userPerms, err := testQuery.GetUsersPermissions(context.Background(), GetUsersPermissionsParams{
		Limit:  int32(n),
		Offset: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, userPerms)
	require.Len(t, userPerms, n)
}

func TestPermissionsByUserId(t *testing.T) {
	upm := CreateRandomUserPermission(t)

	userPerms, err := testQuery.GetPermissionsByUserId(context.Background(), upm.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, userPerms)
}

func TestCreateUserPermissions(t *testing.T) {
	user := CreateRandomUser(t)
	creator := CreateRandomUser(t)
	permIds := []int64{1, 5, 8, 9}

	pIds := []int64{}
	userIds := []string{}
	uCreatedByIds := []string{}

	for _, uPerm := range permIds {
		pIds = append(pIds, uPerm)
		userIds = append(userIds, user.ID)
		uCreatedByIds = append(uCreatedByIds, creator.ID)
	}

	err := testQuery.CreateUserPermissions(context.Background(), CreateUserPermissionsParams{
		UserID:       userIds,
		PermissionID: pIds,
		CreatedbyID:  uCreatedByIds,
	})

	require.NoError(t, err)
}
