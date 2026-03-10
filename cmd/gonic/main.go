package main

import (
	"log/slog"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"project/module/cmd/gonic/controllers/template_delete_body_controller"
	"project/module/cmd/gonic/controllers/template_delete_parameter_controller"
	"project/module/cmd/gonic/controllers/template_get_controller"
	"project/module/cmd/gonic/controllers/template_get_parameter_controller"
	"project/module/cmd/gonic/controllers/template_post_body_controller"
	"project/module/cmd/gonic/controllers/template_post_parameter_controller"
	"project/module/cmd/gonic/controllers/template_put_body_controller"
	"project/module/cmd/gonic/controllers/template_put_parameter_controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	slog.Info("START")

	godotenv.Load(filepath.Join(os.Getenv("APP_BASE_PATH"), "configs", ".env"))

	appBasePath := os.Getenv("APP_BASE_PATH")
	staticPath := filepath.Join(appBasePath, "static")
	tmpPath := filepath.Join(appBasePath, "tmp")

	router := gin.Default()
	router.MaxMultipartMemory = math.MaxInt64

	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Static("/static", staticPath)
	router.Static("/tmp", tmpPath)

	router.DELETE("/template/body", template_delete_body_controller.Handle)
	router.DELETE("/template/parameter", template_delete_parameter_controller.Handle)
	router.GET("/template", template_get_controller.Handle)
	router.GET("/template/parameter", template_get_parameter_controller.Handle)
	router.POST("/template/body", template_post_body_controller.Handle)
	router.POST("/template/parameter", template_post_parameter_controller.Handle)
	router.PUT("/template/body", template_put_body_controller.Handle)
	router.PUT("/template/parameter", template_put_parameter_controller.Handle)

	router.Run()
	slog.Info("END")
}
