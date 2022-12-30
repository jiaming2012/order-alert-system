package models

type BadResponseErr struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}
