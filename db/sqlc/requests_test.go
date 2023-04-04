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
		Status:       ApprovalStatusPending,
		Amount:       utils.RandomInteger(100, 1000000),
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

func TestGetRequests(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		CreateRandomRequest(t)
	}

	requests, err := testQuery.GetRequests(context.Background(), GetRequestsParams{
		Limit:  int32(n),
		Offset: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, requests)
	require.Len(t, requests, n)
}

func TestGetRequestsForApproval(t *testing.T) {
	client := CreateRandomClient(t)
	initiator := CreateRandomUser(t)
	approver := CreateRandomUser(t)

	req, err := testQuery.CreateRequest(context.Background(), CreateRequestParams{
		Title:        utils.RandomString(15),
		Status:       ApprovalStatusPending,
		Amount:       utils.RandomInteger(100, 1000000),
		PaidToID:     client.ID,
		CreatedbyID:  initiator.ID,
		ApprovedbyID: approver.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, req)
	require.Equal(t, req.PaidToID, client.ID)
	require.Equal(t, req.CreatedbyID, initiator.ID)
	require.Equal(t, approver.ID, req.ApprovedbyID)

	requests, err := testQuery.GetRequestsToApprove(context.Background(), GetRequestsToApproveParams{
		Status:       ApprovalStatusPending,
		ApprovedbyID: req.ApprovedbyID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, requests)
	require.Len(t, requests, 1)
}

func TestUpdateRequest(t *testing.T) {
	req := CreateRandomRequest(t)

	reviewedReq, err := testQuery.UpdateRequest(context.Background(), UpdateRequestParams{
		ID: req.ID,
		Status: NullApprovalStatus{
			ApprovalStatus: ApprovalStatusApproved,
			Valid:          true,
		},
	})

	require.NoError(t, err)
	require.Equal(t, req.ID, reviewedReq.ID)
	require.Equal(t, req.Amount, reviewedReq.Amount)
	require.NotEqual(t, req.Status, reviewedReq.Status)
}
