package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	response := &Response{}
	response.Message = "success"
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func Error(ctx *gin.Context, message string, data interface{}) {
	response := &Response{}
	response.Message = message
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func Exception(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, message)
}
