package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func MakeMobileMoneyPayment(clientPhone string, amount int64, initiator string, token string) (string, error) {
	url := "https://sandbox.safaricom.co.ke/mpesa/b2c/v1/paymentrequest"

	securityCredentials := os.Getenv("DARAJA_SECURITY_CREDENTIALS")
	paymentParty := os.Getenv("PAYMENT_PARTY")
	phone, err := strconv.Atoi(clientPhone)

	if err != nil {
		return "", err
	}

	requestBody := fmt.Sprintf("{\"InitiatorName\":\"%s\",\"SecurityCredential\":\"%s\",\"CommandID\":\"BusinessPayment\",\"Amount\": %d,\"PartyA\":\"%s\",\"PartyB\": %d,\"Remarks\":\"Title for request\",\"QueueTimeOutURL\":\"https://mydomain.com/b2c/queue\",\"ResultURL\":\"https://mydomain.com/b2c/result\",\"Occassion\":\"Sentences of up to 100 characters\"}", initiator, securityCredentials, amount, paymentParty, phone)

	payload := strings.NewReader(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	return "", nil
}
