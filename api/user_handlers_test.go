package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ShadrackAdwera/go-payments/db/mocks"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func RandomUser() db.User {
	return db.User{
		ID:       utils.RandomString(10),
		Username: utils.RandomString(15),
	}
}

func TestGetUserEndpoint(t *testing.T) {
	user := RandomUser()

	testCases := []struct {
		name       string
		userId     string
		buildStub  func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "TestOK",
			userId: user.ID,
			buildStub: func(store *mockdb.MockTxStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareResponses(t, recorder.Body, user)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)
			store := mockdb.NewMockTxStore(ctlr)

			defer ctlr.Finish()

			testCase.buildStub(store)

			url := fmt.Sprintf("/api/users/%s", testCase.userId)

			req, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			srv := newServer(store)

			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)
		})
	}
}

func compareResponses(t *testing.T, body *bytes.Buffer, user db.User) {
	var jsonUser db.User

	b, err := io.ReadAll(body)

	require.NoError(t, err)

	err = json.Unmarshal(b, &jsonUser)

	require.NoError(t, err)

	require.NotEmpty(t, jsonUser)
	require.Equal(t, jsonUser.ID, user.ID)
	require.Equal(t, jsonUser.Username, user.Username)
}
