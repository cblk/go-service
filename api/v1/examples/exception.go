package examples

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
)

func Exception(ctx *gin.Context) error {
	return response.NewExceptionResponse(errors.New("internal server error"))
}
