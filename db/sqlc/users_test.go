package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	id := utils.RandomString(12)
	username := utils.RandomString(10)

	user, err := testQuery.CreateUser(context.Background(), CreateUserParams{
		ID:       id,
		Username: username,
	})

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, id)
	require.Equal(t, user.Username, username)
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

	foundUser, err = testQuery.GetUser(context.Background(), utils.RandomString(6))
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

func TestUpdateUser(t *testing.T) {
	user := CreateRandomUser(t)

	newUsername := utils.RandomString(15)

	updatedUser, err := testQuery.UpdateUser(context.Background(), UpdateUserParams{
		ID: user.ID,
		Username: sql.NullString{
			String: newUsername,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.Equal(t, newUsername, updatedUser.Username)
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
