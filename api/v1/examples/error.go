package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
)

func Error(ctx *gin.Context) error {
	validation := response.NewValidationErrorResponse()
	validation.Data.Message = "user_not_exist"
	validation.Data.FieldName = "username"

	return validation
}