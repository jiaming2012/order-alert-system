package api

import (
	"encoding/json"
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/jiaming2012/order-alert-system/backend/sms"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type NewOrderRequest struct {
	OrderId              string `json:"order_number"`
	PhoneNumber          string `json:"phone_number"`
	FormattedPhoneNumber string
}

func (req *NewOrderRequest) Validate() error {
	formattedPhoneNumber, err := sms.ValidatePhoneNumber(req.PhoneNumber)
	if err != nil {
		return err
	}

	req.FormattedPhoneNumber = formattedPhoneNumber

	if _, err = strconv.Atoi(req.OrderId); err != nil {
		return fmt.Errorf("OrderNumber %v must be a number", req.OrderId)
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

func PlaceOrderUpdate(w http.ResponseWriter, r *http.Request) {
	var updateOrderReq UpdateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&updateOrderReq); err != nil {
		sendBadRequestErrResponse("validation", fmt.Errorf("PlaceOrderUpdate: json: err: %w", err), w)
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

	order.Status = updateOrderReq.Status
	if err = order.Save(); err != nil {
		sendBadServerHtmlResponse(err, w)
		return
	}

	w.WriteHeader(200)
}

func PlaceNewOrder(w http.ResponseWriter, r *http.Request) {
	var newOrderReq NewOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&newOrderReq); err != nil {
		sendBadServerHtmlResponse(fmt.Errorf("PlaceNewOrder: json: err: %w", err), w)
		return
	}

	if err := newOrderReq.Validate(); err != nil {
		sendBadServerHtmlResponse(err, w)
		return
	}

	newOrder := models.Order{
		OrderNumber: newOrderReq.OrderId,
		PhoneNumber: newOrderReq.PhoneNumber,
		Status:      "open",
	}

	if err := newOrder.Create(); err != nil {
		sendBadServerHtmlResponse(err, w)
		return
	}

	w.WriteHeader(201)
}

func renderHomepage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("template/index.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("Error executing template :", err)
			sendBadServerHtmlResponse(err, w)
		}
	} else if r.Method == "POST" {
		fmt.Println("POST")
		err := r.ParseForm()
		if err != nil {
			log.Println("Error executing template :", err)
			sendBadServerHtmlResponse(err, w)
		}

		for k, v := range r.Form {
			fmt.Println(k, v)
		}
		fmt.Println("... values")
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

	err := models.BadResponseErr{Type: "invalid_input", Msg: errorMsg}
	renderTemplate("template/4s00-error.html", err, w)
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
		log.Println("Error executing template :", err)
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
