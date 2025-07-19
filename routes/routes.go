package routes

import "net/http"

func Routes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	return router

}
