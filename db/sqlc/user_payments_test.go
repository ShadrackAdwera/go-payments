package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomUserPayment(t *testing.T) UserPayment {
	request := CreateRandomRequest(t)
	client := CreateRandomClient(t)

	userPayment, err := testQuery.CreateUserPayment(context.Background(), CreateUserPaymentParams{
		RequestID: sql.NullInt64{
			Int64: request.ID,
			Valid: true,
		},
		ClientID: sql.NullInt64{
			Int64: client.ID,
			Valid: true,
		},
		Status: PaidStatusNotPaid,
	})

	require.NoError(t, err)
	require.NotEmpty(t, userPayment)
	require.Equal(t, request.ID, userPayment.RequestID.Int64)
	require.Equal(t, client.ID, userPayment.ClientID.Int64)

	return userPayment
}

func TestCreateUserPayment(t *testing.T) {
	CreateRandomUserPayment(t)
}

func TestGetUserPayment(t *testing.T) {
	uPayment := CreateRandomUserPayment(t)

	payment, err := testQuery.GetUserPayment(context.Background(), uPayment.ID)

	require.NoError(t, err)
	require.NotEmpty(t, payment)
	require.Equal(t, uPayment.ID, payment.ID)
	require.Equal(t, uPayment.ClientID, payment.ClientID)
	require.Equal(t, uPayment.RequestID, payment.RequestID)

	req, err := testQuery.GetRequest(context.Background(), uPayment.RequestID.Int64)

	require.NoError(t, err)
	require.NotEmpty(t, req)
	require.Equal(t, req.Amount, payment.Amount)
	require.Equal(t, req.Title, payment.Title)

	client, err := testQuery.GetClient(context.Background(), uPayment.ClientID.Int64)

	require.NoError(t, err)
	require.NotEmpty(t, client)
	require.Equal(t, client.Name, payment.Name)
	require.Equal(t, client.PreferredPaymentType, payment.PreferredPaymentType)
}

func TestUpdateUserPayment(t *testing.T) {
	uPayment := CreateRandomUserPayment(t)

	payment, err := testQuery.UpdateUserPayment(context.Background(), UpdateUserPaymentParams{
		ID:     uPayment.ID,
		Status: PaidStatusPaid,
	})

	require.NoError(t, err)
	require.NotEmpty(t, payment)
	require.Equal(t, payment.ID, uPayment.ID)
	require.Equal(t, uPayment.RequestID.Int64, payment.RequestID.Int64)
	require.Equal(t, uPayment.ClientID.Int64, payment.ClientID.Int64)
}
