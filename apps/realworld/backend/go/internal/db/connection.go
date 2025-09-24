package db

import (
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(connectionString string) *sqlx.DB {
	slog.Info("Connecting to database...")
	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		panic(err)
	}

	slog.Info("Connected to database")
	return db
}
