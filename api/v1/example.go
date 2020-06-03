package v1

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
	"go_service/models"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FormError struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Success(ctx *gin.Context) {
	task := models.NewTask()
	response.Success(ctx, task)
}

func Error(ctx *gin.Context) {
	validation := &FormError{}
	validation.Username = "user_not_exist"

	response.Error(ctx, "validation_error", validation)
}

func Auth(ctx *gin.Context) {
	login := &LoginForm{}

	if err := ctx.ShouldBindJSON(login); err != nil {
		response.Error(ctx, "validation_error", nil)
		return
	}

	validation := &FormError{}

	if login.Username != "admin" {
		validation.Username = "user_not_exist"
		response.Error(ctx, "validation_error", validation)
		return
	}

	if login.Password != "admin" {
		validation.Password = "incorrect_password"
		response.Error(ctx, "validation_error", validation)
		return
	}

	response.Success(ctx, login)
}

func Exception(ctx *gin.Context) {
	response.Exception(ctx, "internal_error")
}
