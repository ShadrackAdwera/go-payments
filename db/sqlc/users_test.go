package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	userId, err := uuid.NewRandom()

	require.NoError(t, err)

	username := utils.RandomString(10)
	email := fmt.Sprintf("%s@mail.com", username)

	user, err := testQuery.CreateUser(context.Background(), CreateUserParams{
		ID:       userId,
		Username: username,
		Email:    email,
		Role:     UserRoles(utils.RandomRole()),
	})

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, userId)
	require.Equal(t, user.Username, username)
	require.Equal(t, user.Email, email)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)

	foundUser, err := testQuery.GetUser(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, foundUser)
	require.Equal(t, foundUser.ID, user.ID)
	require.Equal(t, foundUser.Username, user.Username)
	require.Equal(t, foundUser.Email, user.Email)
	require.Equal(t, foundUser.Role, user.Role)

	randomUUID, err := uuid.NewRandom()

	require.NoError(t, err)

	foundUser, err = testQuery.GetUser(context.Background(), randomUUID)
	require.Error(t, err)
	require.Empty(t, foundUser)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestGetUsers(t *testing.T) {
	n := 8

	for i := 0; i < n; i++ {
		CreateRandomUser(t)
	}

	users, err := testQuery.GetUsers(context.Background(), GetUsersParams{
		Limit:  5,
		Offset: 1,
	})

	require.NoError(t, err)
	require.Len(t, users, 5)
}

func TestUpdateUserRole(t *testing.T) {
	user := CreateRandomUser(t)

	updatedUser, err := testQuery.UpdateUserRole(context.Background(), UpdateUserRoleParams{
		ID:   user.ID,
		Role: UserRoles(utils.RandomRole()),
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
}

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)

	err := testQuery.DeleteUser(context.Background(), user.ID)

	require.NoError(t, err)
	foundUser, err := testQuery.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, foundUser)
	require.ErrorIs(t, err, sql.ErrNoRows)

}
