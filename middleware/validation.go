package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ashunasar/go-jwt-auth-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type contextKey string

const RequestBodyKey contextKey = "requestBody"

const UserIdKey contextKey = "UserIdKey "

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

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			utils.WriteJson(w, http.StatusUnauthorized, utils.GeneralError(fmt.Errorf("auth token not provided")))
			return
		}

		splitToken := strings.Split(authToken, "Bearer ")
		authToken = splitToken[1]

		id, err := utils.VerifyAccessToken(authToken)
		if err != nil {

			utils.WriteJson(w, http.StatusUnauthorized, utils.GeneralError(err))
			return
		}
		ctx := context.WithValue(r.Context(), UserIdKey, id)
		next(w, r.WithContext(ctx))
	}
}

// Helper function to get validated body from context
func GetRequestBody[T any](r *http.Request) (T, bool) {
	body, ok := r.Context().Value(RequestBodyKey).(T)
	return body, ok
}

func GetUserID(r *http.Request) (uuid.UUID, bool) {
	userID, ok := r.Context().Value(UserIdKey).(uuid.UUID)
	return userID, ok
}
