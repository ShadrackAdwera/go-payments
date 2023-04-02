package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

func RandomPermission() db.Permission {
	return db.Permission{
		ID:          utils.RandomInteger(1, 100),
		Name:        utils.RandomString(12),
		Description: utils.RandomString(18),
		CreatedbyID: "",
	}
}

func TestCreatePermission(t *testing.T) {
	perm := RandomPermission()

	testCases := []struct {
		name       string
		body       CreatePermissionArgs
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestOK",
			body: CreatePermissionArgs{
				Name:        perm.Name,
				Description: perm.Description,
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				p := db.Permission{
					ID:          perm.ID,
					Name:        perm.Name,
					Description: perm.Description,
					CreatedbyID: perm.CreatedbyID,
				}
				store.EXPECT().CreatePermission(gomock.Any(), gomock.Eq(db.CreatePermissionParams{
					Name:        perm.Name,
					Description: perm.Description,
					CreatedbyID: "",
				})).Times(1).Return(p, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				comparePermissionResponses(t, recorder.Body, perm)
			},
		},
		{
			name: "TestBadRequest",
			body: CreatePermissionArgs{
				Name: perm.Name,
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().CreatePermission(gomock.Any(), gomock.Eq(db.CreatePermissionParams{
					Name:        perm.Name,
					Description: perm.Description,
					CreatedbyID: "",
				})).Times(0)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TestInternalServerError",
			body: CreatePermissionArgs{
				Name:        perm.Name,
				Description: perm.Description,
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().CreatePermission(gomock.Any(), gomock.Eq(db.CreatePermissionParams{
					Name:        perm.Name,
					Description: perm.Description,
					CreatedbyID: "",
				})).Times(1).Return(db.Permission{}, sql.ErrConnDone)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			req, err := http.NewRequest(http.MethodPost, "/api/permissions", bytes.NewReader(jsonBody))

			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			srv := newServer(store)

			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)
		})
	}
}

type CreatePermissionResponse struct {
	Data db.Permission `json:"data"`
}

func comparePermissionResponses(t *testing.T, body *bytes.Buffer, perm db.Permission) {

	var createPermResponse CreatePermissionResponse

	b, err := io.ReadAll(body)

	require.NoError(t, err)

	err = json.Unmarshal(b, &createPermResponse)

	require.NoError(t, err)

	require.NotEmpty(t, createPermResponse)
	require.Equal(t, createPermResponse.Data.ID, perm.ID)
	require.Equal(t, createPermResponse.Data.Name, perm.Name)
	require.Equal(t, createPermResponse.Data.Description, perm.Description)
}
