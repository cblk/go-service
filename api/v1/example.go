package v1

import (
	"github.com/gin-gonic/gin"
	logy "github.com/sirupsen/logrus"
	"go_service/api/v1/response"
	"go_service/models"
)

type LoginForm struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
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

	if err := ctx.ShouldBind(login); err != nil {
		logy.Info("pass params error", err)
		response.Error(ctx, "validation_error", err)
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
