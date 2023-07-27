package notification

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage(phoneNumber string, message string) {
	senderPhoneNumber := viper.GetString("TWILIO_PHONE_NUMBER")
	senderAccountID := viper.GetString("TWILIO_ACCOUNT_SID")
	senderAuthToken := viper.GetString("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: senderAccountID, Password: senderAuthToken})
	createMessageParams := &twilioApi.CreateMessageParams{}
	createMessageParams.SetTo(phoneNumber)
	createMessageParams.SetFrom(senderPhoneNumber)
	createMessageParams.SetBody(message)

	resp, err := client.Api.CreateMessage(createMessageParams)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Printf("Response: %s", string(response))
	}
}
