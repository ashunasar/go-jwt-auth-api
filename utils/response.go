package utils

import (
	"encoding/json"
	"net/http"
)

const (
	StatusOk    = "ok"
	StatusError = "error"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Messge string `json:"message,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, statutCode int, data any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statutCode)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}

}
