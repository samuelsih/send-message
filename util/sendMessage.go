package util

import (
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
)

func SendMessage(message string) {
	param := openapi.CreateMessageParams{}
	param.SetTo(to)
	param.SetFrom(from)
	param.SetBody(message)

	res, err := client.ApiV2010.CreateMessage(&param)
	if err != nil {
		log.Fatal("Error in SendMessage function\n", err)
	}

	log.Printf("%s", *res.Status)
}