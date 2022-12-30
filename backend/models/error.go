package models

type ApiErrorType string

var ClientError ApiErrorType = "client_error"
var ServerError ApiErrorType = "server_error"

type BadResponseError struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type ApiError struct {
	Type  ApiErrorType
	Error error
}
