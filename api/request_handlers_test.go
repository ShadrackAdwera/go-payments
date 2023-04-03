package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/ShadrackAdwera/go-payments/db/mocks"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func createRequest(approvedById string, status string) db.Request {
	return db.Request{
		ID:           utils.RandomInteger(1, 100),
		Title:        utils.RandomString(24),
		Status:       db.ApprovalStatus(status),
		Amount:       utils.RandomInteger(1000, 100000),
		PaidToID:     utils.RandomInteger(1, 50),
		CreatedbyID:  utils.RandomString(10),
		ApprovedbyID: approvedById,
		CreatedAt:    time.Now(),
		ApprovedAt:   time.Now(),
	}
}

func TestCreateRequestEndpoint(t *testing.T) {
	approvedById := utils.RandomString(9)
	status := utils.RandomStatus()

	request := createRequest(approvedById, status)

	testCases := []struct {
		name       string
		body       NewRequestArgs
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestOK",
			body: NewRequestArgs{
				Title:        request.Title,
				Amount:       request.Amount,
				PaidToID:     request.PaidToID,
				ApprovedbyID: request.ApprovedbyID,
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				req := db.CreateRequestParams{
					Title:        request.Title,
					Status:       request.Status,
					Amount:       request.Amount,
					PaidToID:     request.PaidToID,
					CreatedbyID:  "",
					ApprovedbyID: "",
				}
				store.EXPECT().CreateRequest(gomock.Any(), gomock.Eq(req)).Times(1).Return(request, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				compareRequestResponse(t, recorder.Body, request)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)

			store := mockdb.NewMockTxStore(ctlr)

			defer ctlr.Finish()

			testCase.buildStubs(store)

			jsonBody, err := json.Marshal(testCase.body)

			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/api/requests", bytes.NewReader(jsonBody))

			require.NoError(t, err)

			srv := newServer(store)

			recorder := httptest.NewRecorder()

			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)
		})
	}
}

func TestGetRequestsToApproveEndpoint(t *testing.T) {
	n := 5

	requests := make([]db.Request, 5)
	status := utils.RandomStatus()

	for i := 0; n < 5; i++ {
		requests[i] = createRequest("approver", status)
	}

	type Query struct {
		status     string
		approverID string
	}

	testCases := []struct {
		name       string
		query      Query
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestOK",
			query: Query{
				status:     status,
				approverID: "approver",
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				q := db.GetRequestsToApproveParams{
					Status:       db.ApprovalStatus(status),
					ApprovedbyID: "approver",
				}
				store.EXPECT().GetRequestsToApprove(gomock.Any(), gomock.Eq(q)).Times(1).Return(requests, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareRequests(t, recorder.Body, requests)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)

			store := mockdb.NewMockTxStore(ctlr)

			defer ctlr.Finish()

			testCase.buildStubs(store)
			srv := newServer(store)
			recorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/api/requests/approval", nil)

			require.NoError(t, err)

			q := req.URL.Query()
			q.Add("status", testCase.query.status)
			q.Add("approver_id", testCase.query.approverID)
			req.URL.RawQuery = q.Encode()

			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)
		})
	}
}

type NewRequestResponse struct {
	Request db.Request `json:"request"`
}

func compareRequestResponse(t *testing.T, body *bytes.Buffer, request db.Request) {
	var reqRes NewRequestResponse

	payload, err := io.ReadAll(body)

	require.NoError(t, err)

	err = json.Unmarshal(payload, &reqRes)

	require.NoError(t, err)

	require.NotEmpty(t, reqRes)
	require.Equal(t, request.PaidToID, reqRes.Request.PaidToID)
	require.Equal(t, request.Amount, reqRes.Request.Amount)
	require.Equal(t, request.Status, reqRes.Request.Status)
	require.Equal(t, request.Title, reqRes.Request.Title)
}

type GetRequestsResponse struct {
	Requests []db.Request `json:"requests"`
}

func compareRequests(t *testing.T, body *bytes.Buffer, requests []db.Request) {
	var requestsFound GetRequestsResponse

	jsonBody, err := io.ReadAll(body)

	require.NoError(t, err)

	err = json.Unmarshal(jsonBody, &requestsFound)

	require.NoError(t, err)
	require.NotEmpty(t, requestsFound)
	require.Equal(t, requests, requestsFound.Requests)
}
