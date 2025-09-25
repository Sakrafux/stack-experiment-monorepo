package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/api/rest"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/db"
)

func main() {
	slog.Info("Starting server...")

	slog.Info("Loading config...")
	cfg := &config.Config{}
	cfg.Port = "8080"
	cfg.DbConnectionString = "postgres://postgres:password@localhost:5432/realworld"
	cfg.JWT.AccessSecret = os.Getenv("ACCESS_SECRET")
	cfg.JWT.RefreshSecret = os.Getenv("REFRESH_SECRET")
	slog.Info("Config: ", "config", cfg)

	sql := db.Connect(cfg.DbConnectionString)
	defer sql.Close()

	api := rest.NewApi(cfg, sql)

	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: api.CreateRouter(),
	}

	slog.Info("Listening on port " + cfg.Port + "...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
