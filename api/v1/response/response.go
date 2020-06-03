package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	message string
	data    interface{}
}

func Success(ctx *gin.Context, data interface{}) {
	response := &Response{}
	response.message = "success"
	response.data = data

	ctx.JSON(http.StatusOK, response)
}

func Error(ctx *gin.Context, message string, data interface{}) {
	response := &Response{}
	response.message = message
	response.data = data

	ctx.JSON(http.StatusOK, response)
}

func Exception(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, message)
}
