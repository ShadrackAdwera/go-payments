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

func CreateRequest() db.Request {
	return db.Request{
		ID:           utils.RandomInteger(1, 100),
		Title:        utils.RandomString(24),
		Status:       db.ApprovalStatus(utils.RandomStatus()),
		Amount:       utils.RandomInteger(1000, 100000),
		PaidToID:     utils.RandomInteger(1, 50),
		CreatedbyID:  utils.RandomString(10),
		ApprovedbyID: utils.RandomString(8),
		CreatedAt:    time.Now(),
		ApprovedAt:   time.Now(),
	}
}

func TestCreateRequestEndpoint(t *testing.T) {
	request := CreateRequest()

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
