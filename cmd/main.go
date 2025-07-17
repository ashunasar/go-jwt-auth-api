package main

import (
	"log/slog"

	"github.com/ashunasar/go-jwt-auth-api/config"
)

func main() {

	cfg := config.LoadEnv()

	slog.Info("Db Path is values are ", slog.String("db path", cfg.DbPath))
}
