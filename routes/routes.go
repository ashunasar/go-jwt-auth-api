package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/utils"
	"github.com/go-playground/validator/v10"
)

type SignUpBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("POST /api/auth/signup", func(w http.ResponseWriter, r *http.Request) {

		var signUpBody SignUpBody

		err := json.NewDecoder(r.Body).Decode(&signUpBody)
		if errors.Is(err, io.EOF) {
			utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(fmt.Errorf("empty request body")))
			return
		} else if err != nil {
			utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
			return
		}
		if err := validator.New().Struct(signUpBody); err != nil {
			fmt.Printf("type of err %T", err)
			validationErrs := err.(validator.ValidationErrors)
			utils.WriteJson(w, http.StatusBadRequest, utils.ValidationErrors(validationErrs))
			return
		}

		utils.WriteJson(w, http.StatusAccepted, utils.GeneralResponse("everything looks good"))

	})

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world Bro !"))
	})

	return router

}
