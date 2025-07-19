package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

func GeneralResponse(data any) Response {
	return Response{
		Status: StatusOk,
		Data:   data,
	}

}

func ValidationErrors(errs validator.ValidationErrors) Response {

	var errMsg strings.Builder

	for i, err := range errs {
		if i > 0 {
			errMsg.WriteString(", ")
		}

		field := strings.ToLower(err.Field())

		switch err.Tag() {
		case "required":
			errMsg.WriteString(fmt.Sprintf("field %s is required ", field))

		case "email":
			errMsg.WriteString(fmt.Sprintf("field %s must be a valid email address", field))

		case "min":
			errMsg.WriteString(fmt.Sprintf("field %s must be at least %s characters", field, err.Param()))

		case "max":
			errMsg.WriteString(fmt.Sprintf("field %s must be at most %s characters", field, err.Param()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  errMsg.String(),
	}
}
