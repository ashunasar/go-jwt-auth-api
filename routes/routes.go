package routes

import (
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/middleware"
	"github.com/ashunasar/go-jwt-auth-api/utils"
)

type SignUpBody struct {
	Name     string `json:"name" validate:"required,min=3,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("POST /api/auth/signup", middleware.ValidateRequest[SignUpBody](SignUpHandler))

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world Bro !"))
	})

	return router

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	body, ok := middleware.GetRequestBody[SignUpBody](r)
	if !ok {
		return
	}

	utils.WriteJson(w, http.StatusAccepted, body)

}
