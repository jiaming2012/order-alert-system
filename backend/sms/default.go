package sms

import (
	"encoding/json"
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/constants"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/pubsub"
	"github.com/twilio/twilio-go"
	twilioclient "github.com/twilio/twilio-go/client"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	lookups "github.com/twilio/twilio-go/rest/lookups/v1"
	"io/ioutil"
	"net/http"
	"strings"
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

func Send(ev pubsub.NewOrderCreatedEvent) {
	fmt.Printf("send sms for %v\n", ev)
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

func SendSMS2() {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody("This is the ship that made the Kessel Run in fourteen parsecs?")
	params.SetFrom("+15017122661")
	params.SetTo("+15558675310")

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

func SendSMS(phoneNumber string, msg string) error {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf("To=%s&MessagingServiceSid=%s&Body=%s", phoneNumber, constants.TwillioMessagingServiceSid, msg))
	req, err := http.NewRequest("POST", fmt.Sprintf(constants.TwillioUrl), data)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(constants.TwillioAccountSId, constants.TwillioAuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var respErr smsErr
		bodyText, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			return ioErr
		}
		if jsonErr := json.Unmarshal(bodyText, &respErr); jsonErr != nil {
			return jsonErr
		}
		return fmt.Errorf("send failed: %v", respErr)
	}

	return nil
}
