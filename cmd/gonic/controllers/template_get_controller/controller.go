package template_get_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `่json:"status"`
	Message string `่json:"message"`
}

func Handle(c *gin.Context) {

	c.JSON(http.StatusOK, Response{
		Status:  "OK",
		Message: "Get Hello",
	})
}
