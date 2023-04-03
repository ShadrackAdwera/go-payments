package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// Username-Password-Authentication

type AuthRequest struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type CreatedUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func getProfileData(ctx *gin.Context) Profile {
	var profileData Profile
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	if profile == nil {
		return profileData
	}

	mapstructure.Decode(profile, &profileData)
	return profileData
}

type CreateUserArgs struct {
	Username   string `json:"username" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	Connection string `json:"connection" binding:"required"`
}

func (s *Server) createUser(ctx *gin.Context) {
	p := getProfileData(ctx)
	if p.Sub == "" {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("the request is not authenticated")))
		return
	}
	var createUserArgs CreateUserArgs

	if err := ctx.ShouldBindJSON(&createUserArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	_, err := s.IsAuthorized(ctx, p.Sub, utils.UsersCreate)

	if err != nil {
		ctx.JSON(http.StatusForbidden, errJSON(err))
		return
	}

	// request token from Auth0

	var authReq AuthRequest
	url := "https://" + os.Getenv("AUTH0_DOMAIN") + "/oauth/token"

	pld := fmt.Sprintf("{\"client_id\":\"%s\",\"client_secret\":\"%s\",\"grant_type\":\"client_credentials\"}", os.Getenv("AUTH0_CLIENT_ID"), os.Getenv("AUTH0_CLIENT_SECRET"))

	payload := strings.NewReader(pld)

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&authReq)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errJSON(fmt.Errorf("error decoding json")))
		return
	}

	jsonPayload, err := json.Marshal(createUserArgs)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errJSON(fmt.Errorf("error marshalling json")))
		return
	}

	createUserUrl := "https://" + os.Getenv("AUTH0_DOMAIN") + "/api/v2/users"

	req, err = http.NewRequest(http.MethodPost, createUserUrl, bytes.NewBuffer(jsonPayload))

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errJSON(fmt.Errorf("error creating new request")))
		return
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", authReq.AccessToken))
	req.Header.Add("content-type", "application/json")

	res, err = http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errJSON(fmt.Errorf("error client do")))
		return
	}

	defer res.Body.Close()

	var userRes CreatedUserResponse
	err = json.NewDecoder(res.Body).Decode(&userRes)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(fmt.Errorf("error decoding json into user res")))
		return
	}

	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		ID:       userRes.UserID,
		Username: userRes.Username,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})

}

type GetUserArgs struct {
	ID string `uri:"id" binding:"required"`
}

func (s *Server) getUserById(ctx *gin.Context) {
	var user GetUserArgs

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	foundUser, err := s.store.GetUser(ctx, user.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(errors.New("this user was not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, foundUser)
}
