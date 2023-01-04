package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/services"
	"github.com/jiaming2012/order-alert-system/backend/sms"
	"net/http"
	"time"
)

func welcomeMessage(orderNumber string) string {
	return fmt.Sprintf("Thank you for ordering at YumYums Smokin' Grille. You're all set! We'll text you once order #%s is ready to pick up.", orderNumber)
}

func pickUpMessage() string {
	return fmt.Sprintf("Hey there! We’ve got good news. Your order is ready for pickup! Come on by soon as you can.")
}

func handlePlaceOrderUpdate(ctx *gin.Context) {
	var updateOrderReq models.UpdateOrderRequest

	if err := ctx.BindJSON(&updateOrderReq); err != nil {
		sendBadRequestErrResponse("validation", err, ctx)
	}

	if err := updateOrderReq.Validate(); err != nil {
		sendBadRequestErrResponse("validation", err, ctx)
		return
	}

	order, err := models.GetOrder(updateOrderReq.Id)
	if err != nil {
		sendBadRequestErrResponse("validation", fmt.Errorf("failed to find order using id %v", updateOrderReq.Id), ctx)
		return
	}

	if updateOrderReq.Status == "awaiting_pickup" {
		order.NotifiedAt = time.Now()
		if err = sms.SendSMS(order.PhoneNumber, pickUpMessage()); err != nil {
			sendBadServerHtmlResponse(err, ctx)
			return
		}
	}

	if updateOrderReq.Status == "picked_up" {
		order.PickedUpAt = time.Now()
	}

	order.Status = updateOrderReq.Status
	if err = order.Save(); err != nil {
		sendBadServerHtmlResponse(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func handlePlaceNewOrder(ctx *gin.Context) {
	var newOrderReq models.NewOrderRequest

	if err := ctx.Bind(&newOrderReq); err != nil {
		sendBadServerHtmlResponse(fmt.Errorf("handlePlaceNewOrder: json: err: %w", err), ctx)
		return
	}

	if !placeNewOrder(&newOrderReq, false, ctx) {
		return
	}

	ctx.Status(http.StatusCreated)
}

func placeNewOrder(req *models.NewOrderRequest, isHtmlRequest bool, ctx *gin.Context) bool {
	if err := services.PlaceNewOrder(req); err != nil {
		if err.Type == models.ClientError {
			if isHtmlRequest {
				sendBadRequestHtmlResponse(err.Error, ctx)
			} else {
				sendBadRequestErrResponse("validation", err.Error, ctx)
			}

			return false
		} else {
			if isHtmlRequest {
				sendBadServerHtmlResponse(err.Error, ctx)
			} else {
				sendBadServerErrResponse(err.Error, ctx)
			}

			return false
		}
	}

	return true
}

func postHomepageForm(c *gin.Context) {
	newOrderForm := &models.NewOrderRequest{}
	if err := c.Bind(newOrderForm); err != nil {
		sendBadServerHtmlResponse(err, c)
		return
	}

	if !placeNewOrder(newOrderForm, true, c) {
		return
	}

	if err := sms.SendSMS(newOrderForm.FormattedPhoneNumber, welcomeMessage(newOrderForm.OrderNumber)); err != nil {
		sendBadServerHtmlResponse(err, c)
	}

	c.Redirect(http.StatusSeeOther, "/thank-you.html")
}

func getHomepage(c *gin.Context) {
	c.File("templates/index.html")
}

func renderTemplateWithParams(ctx *gin.Context) {
	errorMsg, ok := ctx.GetQuery("errorMsg")
	if !ok || len(errorMsg) == 0 {
		sendBadServerHtmlResponse(fmt.Errorf("expected errorMsg for client error response"), ctx)
		return
	}

	ctx.HTML(400, "400-error.html", gin.H{
		"Msg": errorMsg,
	})
}
