package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qazwsxedcrfvtgbyhnujmikolp"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomInteger(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomPreferredPayment() string {
	status := []string{"master_card", "visa", "mpesa"}
	return status[rand.Intn(len(status))]
}

func RandomRole() string {
	status := []string{"payment_initiator", "payment_approver", "admin"}
	return status[rand.Intn(len(status))]
}

func RandomStatus() string {
	status := []string{"pending", "approved", "rejected"}
	return status[rand.Intn(len(status))]
}
