package api

import (
	"encoding/json"
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/services"
	"github.com/jiaming2012/order-alert-system/backend/sms"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func welcomeMessage(orderNumber string) string {
	return fmt.Sprintf("Thank you for ordering at YumYums Smokin' Grille. You're all set! We'll text you once order #%s is ready to pick up.", orderNumber)
}

func pickUpMessage() string {
	return fmt.Sprintf("Hey there! We’ve got good news. Your order is ready for pickup! Come on by soon as you can.")
}

func HandlePlaceOrderUpdate(w http.ResponseWriter, r *http.Request) {
	var updateOrderReq models.UpdateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&updateOrderReq); err != nil {
		sendBadRequestErrResponse("validation", fmt.Errorf("HandlePlaceOrderUpdate: json: err: %w", err), w)
		return
	}

	if err := updateOrderReq.Validate(); err != nil {
		sendBadRequestErrResponse("validation", err, w)
		return
	}

	order, err := models.GetOrder(updateOrderReq.Id)
	if err != nil {
		fmt.Println(err)
		sendBadRequestErrResponse("validation", fmt.Errorf("failed to find order using id %v", updateOrderReq.Id), w)
		return
	}

	if updateOrderReq.Status == "awaiting_pickup" {
		order.NotifiedAt = time.Now()
		if err = sms.SendSMS(order.PhoneNumber, pickUpMessage()); err != nil {
			log.Println("Error sending SMS: ", err)
			sendBadServerHtmlResponse(err, w)
			return
		}
	}

	if updateOrderReq.Status == "picked_up" {
		order.PickedUpAt = time.Now()
	}

	order.Status = updateOrderReq.Status
	if err = order.Save(); err != nil {
		sendBadServerHtmlResponse(err, w)
		return
	}

	w.WriteHeader(200)
}

func HandlePlaceNewOrder(w http.ResponseWriter, r *http.Request) {
	var newOrderReq models.NewOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&newOrderReq); err != nil {
		sendBadServerHtmlResponse(fmt.Errorf("HandlePlaceNewOrder: json: err: %w", err), w)
		return
	}

	if !placeNewOrder(&newOrderReq, false, w, r) {
		return
	}

	w.WriteHeader(201)
}

func placeNewOrder(req *models.NewOrderRequest, isHtmlRequest bool, w http.ResponseWriter, r *http.Request) bool {
	if err := services.PlaceNewOrder(req); err != nil {
		if err.Type == models.ClientError {
			if isHtmlRequest {
				sendBadRequestHtmlResponse(err.Error, w, r)
			} else {
				sendBadRequestErrResponse("validation", err.Error, w)
			}

			return false
		} else {
			if isHtmlRequest {
				sendBadServerHtmlResponse(err.Error, w)
			} else {
				sendBadServerErrResponse(err.Error, w)
			}

			return false
		}
	}

	return true
}

func renderHomepage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/index.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("Error executing templates :", err)
			sendBadServerHtmlResponse(err, w)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error executing templates: ", err)
			sendBadServerHtmlResponse(err, w)
			return
		}

		orderNumber := r.FormValue("order_number")
		phoneNumber := r.FormValue("phone_number")

		req := &models.NewOrderRequest{OrderNumber: orderNumber, PhoneNumber: phoneNumber}
		if !placeNewOrder(req, true, w, r) {
			return
		}

		if err = sms.SendSMS(req.FormattedPhoneNumber, welcomeMessage(orderNumber)); err != nil {
			log.Println("Error sending SMS: ", err)
			sendBadServerHtmlResponse(err, w)
			return
		}

		http.Redirect(w, r, "/thank-you.html", http.StatusSeeOther)
	} else {
		sendBadRequestErrResponse("bad_request", fmt.Errorf("unknown http method %v", r.Method), w)
	}
}

func renderResponse(filename string, contentType string, w http.ResponseWriter) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
		sendBadServerHtmlResponse(err, w)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(buf)
}

func renderAsset(filename string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		renderResponse(filename, contentType, w)
	}
}

func renderTemplateWithParams(w http.ResponseWriter, r *http.Request) {
	errorMsg := r.URL.Query().Get("errorMsg")
	if len(errorMsg) == 0 {
		sendBadServerHtmlResponse(fmt.Errorf("expected errorMsg for client error response"), w)
		return
	}

	err := models.BadResponseError{Type: "invalid_input", Msg: errorMsg}
	renderTemplate("templates/400-error.html", err, w)
}

func renderTemplate(filename string, data any, w http.ResponseWriter) {
	parsedTemplate, err := template.ParseFiles(filename)
	if err != nil {
		log.Println("Error parsing file:", err)
		sendBadServerHtmlResponse(err, w)
		return
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		log.Println("Error executing templates :", err)
		sendBadServerHtmlResponse(err, w)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Email : ", r.PostForm.Get("email"))
	fmt.Fprintln(w, "Password : ", r.PostForm.Get("password"))
	fmt.Fprintln(w, "Remember Me : ", r.PostForm.Get("remember_check"))
}
