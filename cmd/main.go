package main

import (
	"log"
	"log/slog"

	"github.com/ashunasar/go-jwt-auth-api/config"
	"github.com/ashunasar/go-jwt-auth-api/database"
)

func main() {

	cfg := config.LoadEnv()
	_, err := database.New(cfg.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Db Path is values are ", slog.String("db path", cfg.DbPath))
}
