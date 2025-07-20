package main

import (
	"log"
	"log/slog"
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

	slog.Info("Db Path is values are ", slog.String("db path", cfg.DbPath))

	server.ListenAndServe()

}
