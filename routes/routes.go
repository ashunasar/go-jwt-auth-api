package routes

import (
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/controllers"
	"github.com/ashunasar/go-jwt-auth-api/middleware"
	"github.com/ashunasar/go-jwt-auth-api/models"
)

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("POST /api/auth/signup", middleware.ValidateRequest[models.SignUpBody](controllers.SignUpHandler))

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world Bro !"))
	})

	return router

}
