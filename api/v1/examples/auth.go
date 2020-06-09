package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
)

type AuthInput struct {

	// rules is validated using gopkg.in/go-playground/validator.v9

	Username string `form:"username" json:"username" rules:"required"`
	Password string `form:"password" json:"password" rules:"required"`
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
