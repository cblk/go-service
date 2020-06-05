package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
)

func Exception(ctx *gin.Context) {
	response.Exception(ctx, "internal_server_error")
}
