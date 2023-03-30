package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomClient(t *testing.T) Client {
	name := utils.RandomString(8)
	email := fmt.Sprintf("%s@mail.com", name)
	user := CreateRandomUser(t)
	phone := utils.RandomString(8)
	accNo := utils.RandomString(24)
	pType := PaymentTypes(utils.RandomPreferredPayment())

	newClient, err := testQuery.CreateClient(context.Background(), CreateClientParams{
		Name:  name,
		Email: email,
		Phone: phone,
		AccountNumber: sql.NullString{
			String: accNo,
			Valid:  true,
		},
		PreferredPaymentType: pType,
		CreatedbyID:          user.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, newClient)
	require.NotZero(t, newClient.ID)
	require.Equal(t, newClient.Name, name)
	require.Equal(t, newClient.Email, email)
	require.Equal(t, newClient.Phone, phone)
	require.Equal(t, newClient.AccountNumber.String, accNo)
	require.Equal(t, newClient.PreferredPaymentType, pType)
	return newClient
}

func TestCreateClient(t *testing.T) {
	CreateRandomClient(t)
}

func TestGetClient(t *testing.T) {
	client := CreateRandomClient(t)

	foundClient, err := testQuery.GetClient(context.Background(), client.ID)

	require.NoError(t, err)
	require.NotEmpty(t, foundClient)
	require.NotZero(t, foundClient.ID)
	require.Equal(t, foundClient.Name, client.Name)
	require.Equal(t, foundClient.Email, client.Email)
	require.Equal(t, foundClient.Phone, client.Phone)
	require.Equal(t, foundClient.AccountNumber.String, client.AccountNumber.String)
	require.Equal(t, foundClient.PreferredPaymentType, client.PreferredPaymentType)

	foundClient, err = testQuery.GetClient(context.Background(), utils.RandomInteger(100, 1000))

	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, foundClient)
}
