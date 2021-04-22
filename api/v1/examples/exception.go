package examples

import (
	"errors"

	"go_service/api/v1/response"

	"github.com/gin-gonic/gin"
)

func Exception(ctx *gin.Context) error {
	return response.NewExceptionResponse(errors.New("internal server error"))
}
