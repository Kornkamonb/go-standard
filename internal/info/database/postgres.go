package database

import (
	"context"
	"log/slog"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() *pgxpool.Pool {
	dsn := os.Getenv("POSTGRE_10_17_66_145_QA")

	if dsn == "" {
		slog.Error("DSN is empty")
		panic("database connection string not set")
	}

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		slog.Error("cannot connect database", "error", err)
		panic(err)
	}

	slog.Info("Database connected")
	return db
}