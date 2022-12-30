package models

import (
	"fmt"
	"strconv"
)

type SmsValidator interface {
	ValidatePhoneNumber(string) (string, *ApiError)
}

type NewOrderRequest struct {
	OrderNumber          string `json:"order_number"`
	PhoneNumber          string `json:"phone_number"`
	FormattedPhoneNumber string
}

func (req *NewOrderRequest) Validate(validator SmsValidator) *ApiError {
	formattedPhoneNumber, err := validator.ValidatePhoneNumber(req.PhoneNumber)
	if err != nil {
		return err
	}

	req.FormattedPhoneNumber = formattedPhoneNumber

	if _, strConvErr := strconv.Atoi(req.OrderNumber); strConvErr != nil {
		return &ApiError{
			Type:  ClientError,
			Error: fmt.Errorf("order # %v must be a number", req.OrderNumber),
		}
	}

	return nil
}

type UpdateOrderRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func (req *UpdateOrderRequest) Validate() error {
	if req.Status != "open" && req.Status != "awaiting_pickup" && req.Status != "closed" {
		return fmt.Errorf("UpdateOrderRequest: invalid order status %v", req.Status)
	}

	return nil
}
