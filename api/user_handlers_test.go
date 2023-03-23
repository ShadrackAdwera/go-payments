package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mockdb "github.com/ShadrackAdwera/go-payments/db/mocks"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg    db.CreateUserParams
	userId uuid.UUID
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	e.arg.ID = arg.ID
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and user id %v", e.arg, e.userId)
}

func EqCreateUserParams(arg db.CreateUserParams, userId uuid.UUID) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, userId}
}

func CreateMockUser(t *testing.T) (db.User, uuid.UUID) {
	userId, err := uuid.NewRandom()

	require.NoError(t, err)

	username := utils.RandomString(10)
	email := fmt.Sprintf("%s@mail.com", username)

	return db.User{
		ID:        userId,
		Username:  username,
		Email:     email,
		Role:      db.UserRoles(utils.RandomRole()),
		CreatedAt: time.Now(),
	}, userId
}

func TestCreateUserEndpoint(t *testing.T) {
	user, userId := CreateMockUser(t)

	testCases := []struct {
		name       string
		body       gin.H
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, body *httptest.ResponseRecorder)
	}{
		{
			name: "TestCreateUserOk",
			body: gin.H{
				"username": user.Username,
				"email":    user.Email,
				"role":     user.Role,
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				usr := db.CreateUserParams{
					Username: user.Username,
					Email:    user.Email,
					Role:     user.Role,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(usr, userId)).Times(1).Return(user, nil)
			},
			comparator: func(t *testing.T, body *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, body.Code)
				compareResponses(t, body.Body, db.User(user))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)

			store := mockdb.NewMockTxStore(ctlr)

			defer ctlr.Finish()

			testCase.buildStubs(store)

			jsonPayload, err := json.Marshal(testCase.body)

			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(jsonPayload))
			require.NoError(t, err)

			server := newServer(store)

			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, req)
			testCase.comparator(t, recorder)

		})
	}
}

func compareResponses(t *testing.T, body *bytes.Buffer, user db.User) {
	reqBody, err := io.ReadAll(body)

	require.NoError(t, err)

	var response UserResponse

	err = json.Unmarshal(reqBody, &response)
	require.NoError(t, err)

	require.NotEmpty(t, response)
	require.Equal(t, response.User.ID, user.ID)
	require.Equal(t, response.User.Username, user.Username)
	require.Equal(t, response.User.Email, user.Email)
	require.Equal(t, response.User.Role, user.Role)
}
