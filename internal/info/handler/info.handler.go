package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"project/module/internal/info/repository"
)

type Handler struct {
	DB *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Handle(c *gin.Context) {
	repo := repository.JobRepository{DB: h.DB}

	result := repo.GetTable()

	c.JSON(http.StatusOK, result)
}