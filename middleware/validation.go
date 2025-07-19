package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/utils"
	"github.com/go-playground/validator/v10"
)

type contextKey string

const RequestBodyKey contextKey = "requestBody"

var validate = validator.New()

func ValidateRequest[T any](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body T

		// Decode JSON body
		err := json.NewDecoder(r.Body).Decode(&body)

		if errors.Is(err, io.EOF) {
			utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(fmt.Errorf("empty request body")))
			return
		} else if err != nil {
			utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
			return
		}

		// Validate the decoded body
		if err := validate.Struct(body); err != nil {
			validationErrs := err.(validator.ValidationErrors)
			utils.WriteJson(w, http.StatusBadRequest, utils.ValidationErrors(validationErrs))
			return
		}

		// Add validated body to context
		ctx := context.WithValue(r.Context(), RequestBodyKey, body)
		next(w, r.WithContext(ctx))
	}
}

// Helper function to get validated body from context
func GetRequestBody[T any](r *http.Request) (T, bool) {
	body, ok := r.Context().Value(RequestBodyKey).(T)
	return body, ok
}
