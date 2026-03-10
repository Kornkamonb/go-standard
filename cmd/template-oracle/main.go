package main

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/godror/godror"
	"github.com/joho/godotenv"
)

func GetOracleVersion() {
	dsn := os.Getenv("ORACLE_FPC")

	db, _ := sql.Open("godror", dsn)
	defer db.Close()

	rows, _ := db.Query("SELECT banner FROM v$version")
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

	GetOracleVersion()

	slog.Info("END")
}
