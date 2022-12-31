package sms

import (
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/constants"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/twilio/twilio-go"
	twilioclient "github.com/twilio/twilio-go/client"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	lookups "github.com/twilio/twilio-go/rest/lookups/v1"
)

type smsErr struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (e smsErr) String() string {
	return fmt.Sprintf("%v Code=%v", e.Message, e.Code)
}

func ValidatePhoneNumber(phoneNumber string) (string, *models.ApiError) {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	client := twilio.NewRestClient()

	params := &lookups.FetchPhoneNumberParams{}
	params.SetCountryCode("US")

	resp, err := client.LookupsV1.FetchPhoneNumber(phoneNumber, params)
	if err != nil {
		switch e := err.(type) {
		case *twilioclient.TwilioRestError:
			return "", &models.ApiError{
				Type:  models.ServerError,
				Error: fmt.Errorf("twilio: %v, status: %v", e.Message, e.Status),
			}
		default:
			return "", &models.ApiError{Type: models.ServerError, Error: e}
		}
	} else {
		if resp.NationalFormat != nil {
			return *resp.NationalFormat, nil
		} else {
			return "", &models.ApiError{Type: models.ClientError, Error: fmt.Errorf("failed to validate number %v", phoneNumber)}
		}
	}
}

func SendSMS(phoneNumber string, msg string) error {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody(msg)
	params.SetFrom(constants.TwilioPhoneNumber)
	params.SetTo(phoneNumber)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	} else {
		if resp.Sid != nil {
			return nil
		} else {
			return fmt.Errorf("SendSMS: response does not have Sid %v", resp)
		}
	}
}
