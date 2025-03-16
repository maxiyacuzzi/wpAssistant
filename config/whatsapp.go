package config

import (
	"fmt"
	"os"
	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/api/v2010"
)

// Cliente de Twilio
var TwilioClient *twilio.RestClient

func ConnectTwilio() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	TwilioClient = client
}

// Enviar mensaje por WhatsApp
func SendWhatsAppMessage(to, body string) {
	from := os.Getenv("TWILIO_WHATSAPP_NUMBER")
	message := &v2010.CreateMessageParams{
		To:   &to,
		From: &from,
		Body: &body,
	}

	_, err := TwilioClient.ApiV2010.CreateMessage(message)
	if err != nil {
		fmt.Println("Error enviando mensaje de WhatsApp:", err)
		return
	}
	fmt.Println("Mensaje enviado con éxito")
}
