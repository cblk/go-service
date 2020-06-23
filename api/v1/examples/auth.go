package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
	"go_service/models"
)

type AuthInput struct {

	// validate is validated using https://github.com/go-playground/validator

	Username string `json:"username" form:"username" validate:"required" description:"The user's username"`
	Password string `json:"password" form:"password" validate:"required" description:"The user's password"`
}

type AuthResponse struct {
	response.Response
	Data models.User `json:"data"`
}

func Auth(ctx *gin.Context, in *AuthInput) (*AuthResponse, error) {

	validation := response.NewValidationErrorResponse()

	if in.Username != "admin" {
		validation.SetFieldName("username")
		validation.SetFieldMessage("user_not_exist")
		return nil, validation
	}

	if in.Password != "admin" {
		validation.SetFieldName("password")
		validation.SetFieldMessage("incorrect_password")
		return nil, validation
	}

	r := &AuthResponse{}
	r.Data = models.User{
		Username: in.Username,
		Password: in.Password,
	}

	return r, nil
}
