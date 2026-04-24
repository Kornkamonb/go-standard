package info

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"project/module/internal/info/handler"
)

func RegisterRoutes(r *gin.Engine, db *pgxpool.Pool) {
	h := handler.NewHandler(db)

	r.GET("/get-table-data", h.GetTableData)
	r.GET("/get-card-data", h.GetCardData)
}