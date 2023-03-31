package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomRequest(t *testing.T) Request {
	client := CreateRandomClient(t)
	initiator := CreateRandomUser(t)
	approver := CreateRandomUser(t)

	req, err := testQuery.CreateRequest(context.Background(), CreateRequestParams{
		Title:        utils.RandomString(15),
		Status:       ApprovalStatus(utils.RandomStatus()),
		Amount:       utils.RandomInteger(1, 5),
		PaidToID:     client.ID,
		CreatedbyID:  initiator.ID,
		ApprovedbyID: approver.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, req)
	require.Equal(t, req.PaidToID, client.ID)
	require.Equal(t, req.CreatedbyID, initiator.ID)
	require.Equal(t, approver.ID, req.ApprovedbyID)

	return req
}

func TestCreateRequest(t *testing.T) {
	CreateRandomRequest(t)
}

func TestGetRequestById(t *testing.T) {
	req := CreateRandomRequest(t)

	foundReq, err := testQuery.GetRequest(context.Background(), req.ID)

	require.NoError(t, err)
	require.NotEmpty(t, foundReq)
	require.Equal(t, foundReq.ID, req.ID)
	require.Equal(t, foundReq.Amount, req.Amount)
	require.Equal(t, foundReq.ApprovedbyID, req.ApprovedbyID)
	require.Equal(t, foundReq.PaidToID, req.PaidToID)

	foundReq, err = testQuery.GetRequest(context.Background(), utils.RandomInteger(100, 1000))

	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, foundReq)
}