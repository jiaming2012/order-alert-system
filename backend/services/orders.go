package services

import "github.com/jiaming2012/order-alert-system/backend/models"

func PlaceNewOrder(newOrderReq *models.NewOrderRequest) *models.ApiError {
	if err := newOrderReq.Validate(Sms{}); err != nil {
		return err
	}

	newOrder := models.Order{
		OrderNumber: newOrderReq.OrderNumber,
		PhoneNumber: newOrderReq.PhoneNumber,
		Status:      "open",
	}

	if err := newOrder.Create(); err != nil {
		return &models.ApiError{
			Type:  models.ServerError,
			Error: err,
		}
	}

	return nil
}
