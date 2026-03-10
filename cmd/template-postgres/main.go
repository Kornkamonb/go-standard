package main

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func GetPostgresVersion() {
	dsn := os.Getenv("POSTGRES_10_17_66_121_IOT")

	db, _ := sql.Open("pgx", dsn)
	defer db.Close()

	rows, _ := db.Query("SELECT VERSION();")
	defer rows.Close()

	for rows.Next() {
		var results string
		rows.Scan(&results)
		slog.Info(results)
	}
}

func main() {
	slog.Info("START")

	appBasePath := filepath.Join(os.Getenv("APP_BASE_PATH"), "configs", ".env")
	err := godotenv.Load(appBasePath)
	if err != nil {
		slog.Error("", "error", err)
		return
	}

	GetPostgresVersion()

	slog.Info("END")
}
