package main

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS1(to string) {
	accountSid := "AC3840f9b8b415c39c89bae9f0bb6823aa"
	authToken := "d337752d9541ae66be34884b7bc52aed"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom("+18148139587")
	params.SetBody("Hello, evryone ")

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creating and sending message: %s\n", err.Error())
		return
	}

	fmt.Printf("message SID: %s\n", *response.Sid)
}

func SendSMS2(to string) {
	accountSid := "AC3840f9b8b415c39c89bae9f0bb6823aa"
	authToken := "d337752d9541ae66be34884b7bc52aed"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom("+18148139587")
	params.SetBody("this is message 2")

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creating and sending message: %s\n", err.Error())
		return
	}

	fmt.Printf("message SID: %s\n", *response.Sid)
}

func FatherSMS(mstype string, phonenum string) {
	if mstype == "SM1" {
		SendSMS1(phonenum)
	} else if mstype == "SM2" {
		SendSMS2(phonenum)
	}
}
