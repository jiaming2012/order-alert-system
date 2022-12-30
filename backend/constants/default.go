package constants

import "os"

var TwillioUrl = "https://api.twilio.com/2010-04-01/Accounts/AC5703401886ac3503607af53f4dc8dd45/Messages.json"
var TwillioMessagingServiceSid = "MG64f62acb0bbfe442b46a8be8dcbd57c0"
var TwillioAccountSId = os.Getenv("TWILIO_ACCOUNT_SID")
var TwillioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
var TwilioPhoneNumber = os.Getenv("TWILIO_PHONE_NUMBER")
var SmsNotificationWelcomeMessage = "the is a test welcome message."

func init() {
	if len(TwillioAccountSId) == 0 {
		panic("TWILIO_ACCOUNT_SID not set")
	}

	if len(TwillioAuthToken) == 0 {
		panic("TWILIO_AUTH_TOKEN not set")
	}

	if len(TwilioPhoneNumber) == 0 {
		panic("TWILIO_PHONE_NUMBER not set")
	}
}
