package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
)

func Error(ctx *gin.Context) {
	validation := &response.ValidationErrorResponse{}
	validation.Data.Message = "user_not_exist"
	validation.Data.FieldName = "username"
	validation.Error(ctx, "validation_error")
}
