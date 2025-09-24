package main

import (
	"log/slog"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/api/rest"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/db"
)

func main() {
	cfg := config.Config{
		Port:               "8080",
		DbConnectionString: "postgres://postgres:password@localhost:5432/realworld",
	}

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
