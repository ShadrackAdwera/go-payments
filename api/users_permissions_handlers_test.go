package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ShadrackAdwera/go-payments/db/mocks"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func userPermission() (db.UsersPermission, []int64, []string, []string) {
	permissionIds := []int64{}
	userIds := []string{}
	createdByIds := []string{}
	userId := utils.RandomString(12)
	n := 5

	for i := 0; i < n; i++ {
		permissionIds = append(permissionIds, utils.RandomInteger(1, 1000))
		userIds = append(userIds, userId)
		createdByIds = append(createdByIds, "")
	}

	return db.UsersPermission{
		ID:           utils.RandomInteger(1, 100),
		UserID:       userId,
		PermissionID: permissionIds[0],
		CreatedbyID:  "",
	}, permissionIds, userIds, createdByIds
}

func TestAddPermissionsToUserHandler(t *testing.T) {

	perm, permIds, userIds, cIds := userPermission()

	testCases := []struct {
		name       string
		body       AddPermissionsToUserArgs
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestOK",
			body: AddPermissionsToUserArgs{
				UserId:        perm.UserID,
				PermissionIds: permIds,
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				args := db.CreateUserPermissionsParams{
					UserID:       userIds,
					PermissionID: permIds,
					CreatedbyID:  cIds,
				}
				store.EXPECT().CreateUserPermissions(gomock.Any(), gomock.Eq(args)).Times(1)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
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

			req, err := http.NewRequest(http.MethodPost, "/api/user-permissions", bytes.NewReader(jsonBody))

			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			srv := newServer(store)

			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)
		})
	}
}
