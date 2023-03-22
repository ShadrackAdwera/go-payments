package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomClient() Client {
	name := utils.RandomString(8)
	email := fmt.Sprintf("%s@mail.com", name)
	return Client{
		ID:    utils.RandomInteger(1, 1000),
		Name:  name,
		Email: email,
		Phone: strconv.Itoa(int(utils.RandomInteger(1, 10))),
		AccountNumber: sql.NullString{
			String: strconv.Itoa(int(utils.RandomInteger(1, 16))),
			Valid:  true,
		},
		PreferredPaymentType: PaymentTypes(utils.RandomPreferredPayment()),
	}
}

func TestCreateClient(t *testing.T) {
	client := CreateRandomClient()

	newClient, err := testQuery.CreateClient(context.Background(), CreateClientParams{
		Name:                 client.Name,
		Email:                client.Email,
		Phone:                client.Phone,
		AccountNumber:        client.AccountNumber,
		PreferredPaymentType: client.PreferredPaymentType,
	})

	require.NoError(t, err)
	require.NotEmpty(t, newClient)
	require.NotZero(t, newClient.ID)
	require.Equal(t, newClient.Name, client.Name)
	require.Equal(t, newClient.Email, client.Email)
	require.Equal(t, newClient.Phone, client.Phone)
	require.Equal(t, newClient.AccountNumber, client.AccountNumber)
	require.Equal(t, newClient.PreferredPaymentType, client.PreferredPaymentType)
}
