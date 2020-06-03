package response

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Response struct {
	message string
	data    interface{}
}

func Success(ctx *gin.Context, data interface{}) {
	os.Exit(0)
}

func Error(ctx *gin.Context, data interface{}) {
	os.Exit(0)
}

func Exception(ctx *gin.Context, message string) {
	os.Exit(0)
}
