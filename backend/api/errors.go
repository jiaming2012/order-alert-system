package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type badResponseErr struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func sendBadServerErrResponse(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Printf("Error: Bad server: %v\n", err)
}

func sendBadRequestErrResponse(errType string, err error, w http.ResponseWriter) {
	resp := badResponseErr{
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
