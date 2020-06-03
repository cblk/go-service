package v1

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
	"go_service/models"
)

type FormError struct {
	username string
	password string
}

func Success(ctx *gin.Context) {
	task := models.NewTask()
	response.Success(ctx, task)
}

func Error(ctx *gin.Context) {
	validation := &FormError{}
	validation.username = "user_not_exist"

	response.Error(ctx, "validation_error", validation)
}

func Exception(ctx *gin.Context) {
	response.Exception(ctx, "internal_error")
}
