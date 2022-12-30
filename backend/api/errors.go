package api

import (
	"encoding/json"
	"fmt"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"net/http"
)

func sendBadServerErrResponse(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Printf("Error: Bad server: %v\n", err)
}

func sendBadServerHtmlResponse(err error, w http.ResponseWriter) {
	sendBadServerErrResponse(err, w)
	renderResponse("template/500-error.html", "text/html", w)
}

func sendBadRequestErrResponse(errType string, err error, w http.ResponseWriter) {
	resp := models.BadResponseErr{
		Type: errType,
		Msg:  err.Error(),
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("Error: failed to marshall error %v, type=%v\n", err, errType)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(bytes)
}
