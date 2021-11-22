package util

import (
	"os"
	"github.com/twilio/twilio-go"
)

var (
	accountSID string
	authToken string
	from string
	to string
	client *twilio.RestClient
)

func CreateClientServer() {
	accountSID = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	from = os.Getenv("FROM")
	to = os.Getenv("TO")

	client = twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSID,
		Password: authToken,
	})
}