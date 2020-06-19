package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
)

type AuthInput struct {

	// validate is validated using https://github.com/go-playground/validator

	Username string `query:"username" json:"username" validate:"required" description:"The user's username"`
	Password string `json:"password" validate:"required" description:"The user's password"`
}

type AuthResponse struct {
	response.Response
	Data AuthInput `json:"data"`
}

func Auth(ctx *gin.Context, in *AuthInput) (*AuthResponse, error) {

	validation := response.NewValidationErrorResponse()

	if in.Username != "admin" {
		validation.SetFieldName("username")
		validation.SetMessage("user_not_exist")
		return nil, validation
	}

	if in.Password != "admin" {
		validation.SetFieldName("password")
		validation.SetMessage("incorrect_password")
		return nil, validation
	}

	r := &AuthResponse{}
	r.Data = *in

	return r, nil
}
