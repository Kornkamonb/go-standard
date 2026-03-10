package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func GetPostgresVersion() {
	dsn := os.Getenv("POSTGRES_10_17_66_121_IOT")
	db, _ := sql.Open("pgx", dsn)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, _ := db.QueryContext(ctx, "SELECT VERSION();")
	defer rows.Close()

	for rows.Next() {
		var results string
		rows.Scan(&results)
		slog.Info(results)
	}
}

func main() {
	slog.Info("START")

	godotenv.Load(filepath.Join(os.Getenv("APP_BASE_PATH"), "configs", ".env"))

	GetPostgresVersion()

	slog.Info("END")
}
