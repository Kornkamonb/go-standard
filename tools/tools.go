//go:build tools
// +build tools

package tools

import (
	// Database
	_ "github.com/godror/godror"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"

	// Utility
	_ "github.com/fsnotify/fsnotify"
	_ "github.com/joho/godotenv"
	_ "github.com/robfig/cron/v3"
	_ "github.com/xuri/excelize/v2"
	_ "gopkg.in/yaml.v3"

	// API
	// _ "github.com/gin-gonic/gin"
	// _ "github.com/gin-contrib/cors"
)
