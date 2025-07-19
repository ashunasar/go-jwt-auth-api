package routes

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
			// w.Write([]byte("Empty Body Provided"))
			// json.NewEncoder(w).Encode(map[string]string{"error": "Empty body"})
			return
		}

		w.Write([]byte("from post api"))

	})

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world Bro !"))
	})

	return router

}
