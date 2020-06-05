package examples

import (
	"github.com/gin-gonic/gin"
	logy "github.com/sirupsen/logrus"
	"go_service/api/v1/response"
)

type AuthInput struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type AuthResponse struct {
	response.Response
	Data AuthInput
}

func Auth(ctx *gin.Context) {
	login := &AuthInput{}

	if err := ctx.ShouldBind(login); err != nil {
		logy.Info("pass params error", err)
		response.Error(ctx, "validation_error", err)
		return
	}

	validation := &response.ValidationErrorResponse{}

	if login.Username != "admin" {
		validation.Data.FieldName = "username"
		validation.Data.Message = "user_not_exist"
		validation.Error(ctx, "validation_error")
		return
	}

	if login.Password != "admin" {
		validation.Data.FieldName = "password"
		validation.Data.Message = "incorrect_password"
		validation.Error(ctx, "validation_error")
		return
	}

	r := &AuthResponse{}
	r.Data = *login
	r.Success(ctx)
}
