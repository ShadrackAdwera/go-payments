package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ShadrackAdwera/go-payments/db/mocks"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func CreateNewClient() db.Client {
	name := utils.RandomString(8)
	email := fmt.Sprintf("%s@mail.com", name)
	phone := utils.RandomString(8)
	accNo := utils.RandomString(24)
	pType := db.PaymentTypes(utils.RandomPreferredPayment())

	return db.Client{
		ID:    utils.RandomInteger(1, 100),
		Name:  name,
		Email: email,
		Phone: phone,
		AccountNumber: sql.NullString{
			String: accNo,
			Valid:  true,
		},
		PreferredPaymentType: pType,
		CreatedbyID:          utils.RandomString(12),
	}
}

func setUpSession(ctx *gin.Context, t *testing.T) {

	pData := &Profile{
		Iss:      "https://dev-	kvocbgrhk3ci708w.us.auth0.com/",
		Sub:      "auth0|64186f3b8cca2db234b4f009",
		Aud:      []string{""},
		Iat:      1680186594,
		Exp:      1680222594,
		Azp:      "",
		Scope:    "",
		Name:     "vocek57061@oniecan.com",
		Nickname: "vocek57061",
		Picture:  "",
		Sid:      "TP7i5lfDgfjJe8HZmCpVQjWnSWsu_jEX",
	}

	b, err := json.Marshal(&pData)

	require.NoError(t, err)

	var profile map[string]interface{}

	err = json.Unmarshal(b, &profile)

	require.NoError(t, err)

	session := sessions.Default(ctx)
	session.Set("profile", profile)
}

func TestCreateClient(t *testing.T) {
	client := CreateNewClient()

	testCases := []struct {
		name       string
		body       CreateClientArgs
		addSession func(ctx *gin.Context)
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestOk",
			body: CreateClientArgs{
				Name:                 client.Name,
				Email:                client.Email,
				Phone:                client.Phone,
				AccountNumber:        client.AccountNumber.String,
				PreferredPaymentType: string(client.PreferredPaymentType),
			},
			addSession: func(ctx *gin.Context) {
				setUpSession(ctx, t)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				cl := db.CreateClientParams{
					Name:                 client.Name,
					Email:                client.Email,
					Phone:                client.Phone,
					AccountNumber:        client.AccountNumber,
					PreferredPaymentType: client.PreferredPaymentType,
					CreatedbyID:          "",
				}
				store.EXPECT().CreateClient(gomock.Any(), gomock.Eq(cl)).Times(1).Return(client, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				compareClientResponses(t, recorder.Body, client)
			},
		},
		{
			name: "TestBadRequest",
			body: CreateClientArgs{
				Name:                 client.Name,
				Phone:                client.Phone,
				AccountNumber:        client.AccountNumber.String,
				PreferredPaymentType: string(client.PreferredPaymentType),
			},
			addSession: func(ctx *gin.Context) {
				setUpSession(ctx, t)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				cl := db.CreateClientParams{
					Name:                 client.Name,
					Phone:                client.Phone,
					AccountNumber:        client.AccountNumber,
					PreferredPaymentType: client.PreferredPaymentType,
					CreatedbyID:          "",
				}
				store.EXPECT().CreateClient(gomock.Any(), gomock.Eq(cl)).Times(0)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TestInternalServerError",
			body: CreateClientArgs{
				Name:                 client.Name,
				Email:                client.Email,
				Phone:                client.Phone,
				AccountNumber:        client.AccountNumber.String,
				PreferredPaymentType: string(client.PreferredPaymentType),
			},
			addSession: func(ctx *gin.Context) {
				setUpSession(ctx, t)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				cl := db.CreateClientParams{
					Name:                 client.Name,
					Email:                client.Email,
					Phone:                client.Phone,
					AccountNumber:        client.AccountNumber,
					PreferredPaymentType: client.PreferredPaymentType,
					CreatedbyID:          "",
				}
				store.EXPECT().CreateClient(gomock.Any(), gomock.Eq(cl)).Times(1).Return(db.Client{}, sql.ErrConnDone)
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

			jsonData, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			//c, _ := gin.CreateTestContext(recorder)

			req, err := http.NewRequest(http.MethodPost, "/api/clients", bytes.NewReader(jsonData))
			require.NoError(t, err)

			srv := newServer(store)

			//testCase.addSession(c)
			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)

		})
	}
}

type CreatedClientResponse struct {
	Client db.Client `json:"client"`
}

func compareClientResponses(t *testing.T, res *bytes.Buffer, client db.Client) {
	b, err := io.ReadAll(res)
	require.NoError(t, err)

	var jsonClient CreatedClientResponse

	err = json.Unmarshal(b, &jsonClient)
	require.NoError(t, err)
	require.NotEmpty(t, jsonClient)
	require.Equal(t, jsonClient.Client.ID, client.ID)
	require.Equal(t, jsonClient.Client.Name, client.Name)
	require.Equal(t, jsonClient.Client.Email, client.Email)
	require.Equal(t, jsonClient.Client.Phone, client.Phone)
	require.Equal(t, jsonClient.Client.PreferredPaymentType, client.PreferredPaymentType)

}
