package main

import (
	"log"
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/config"
	"github.com/ashunasar/go-jwt-auth-api/database"
	"github.com/ashunasar/go-jwt-auth-api/routes"
)

func main() {

	cfg := config.LoadEnv()

	err := database.Connect(cfg.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	handler := routes.Routes()

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	server.ListenAndServe()

}
