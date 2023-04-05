package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DarajaTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type RequestTokenResponse struct {
	AccessToken string
	ExpiresAt   time.Time
}

func RequestDarajaToken() (RequestTokenResponse, error) {
	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return RequestTokenResponse{}, err
	}

	clientName := os.Getenv("DARAJA_CONSUMER_KEY")
	clientSecret := os.Getenv("DARAJA_CONSUMER_SECRET")

	s := fmt.Sprintf("%s:%s", clientName, clientSecret)

	tkn := base64.StdEncoding.EncodeToString([]byte(s))
	bearerTkn := fmt.Sprintf("Basic %s", tkn)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearerTkn)

	res, err := client.Do(req)

	if err != nil {
		return RequestTokenResponse{}, err
	}
	defer res.Body.Close()

	var tknReq DarajaTokenResponse

	err = json.NewDecoder(res.Body).Decode(&tknReq)

	if err != nil {
		return RequestTokenResponse{}, err
	}

	sec, err := strconv.Atoi(tknReq.ExpiresIn)

	if err != nil {
		return RequestTokenResponse{}, err
	}

	expiry := time.Now().Add(time.Second * time.Duration(sec))

	tknResponse := RequestTokenResponse{
		AccessToken: tknReq.AccessToken,
		ExpiresAt:   expiry,
	}

	return tknResponse, nil
}
