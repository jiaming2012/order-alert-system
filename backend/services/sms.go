package services

import (
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/sms"
)

type Sms struct{}

func (s Sms) ValidatePhoneNumber(phoneNumber string) (string, *models.ApiError) {
	return sms.ValidatePhoneNumber(phoneNumber)
}
