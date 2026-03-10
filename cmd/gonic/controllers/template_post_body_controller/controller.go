package template_post_body_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Input1 string `json:"input1" binding:"required"`
	Input2 string `json:"input2" binding:"required"`
}

type Response struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    Request `json:"data"`
}

func Handle(c *gin.Context) {

	var req Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "OK",
		Message: "Get Hello",
		Data:    req,
	})
}
