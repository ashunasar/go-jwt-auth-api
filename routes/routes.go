package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/utils"
)

type SignUpBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

		w.Write([]byte("from post api"))

	})

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world Bro !"))
	})

	return router

}
