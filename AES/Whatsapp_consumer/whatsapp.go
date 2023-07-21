package main

import (
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func sendWhatsApp1(num string) {
	accountSid := "AC3840f9b8b415c39c89bae9f0bb6823aa"
	authToken := "d337752d9541ae66be34884b7bc52aed"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello 1")
	params.SetTo(`whatsapp:` + num)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

func sendWhatsApp2(num string) {
	accountSid := "AC3840f9b8b415c39c89bae9f0bb6823aa"
	authToken := "d337752d9541ae66be34884b7bc52aed"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello2, this is another message guyyyssss")
	params.SetTo(`whatsapp:` + num)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

func FatherWhatsapp(mstype string, phone string) {
	if mstype == "WH1" {
		sendWhatsApp1(phone)
	} else if mstype == "WH2" {
		sendWhatsApp2(phone)
	}
}

// func main() {
// 	sendWhatsApp(MessageBody1)
// 	sendWhatsApp(MessageBody2)
// }
