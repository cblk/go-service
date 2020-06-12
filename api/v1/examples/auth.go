package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
	"go_service/models"
)

type AuthInput struct {

	// rules is validated using gopkg.in/go-playground/validator.v9

	Username string `form:"username" json:"username" rules:"required" description:"The user's username"`
	Password string `form:"password" json:"password" rules:"required" description:"The user's password"`
}

type AuthResponse struct {
	response.Response
	Data models.User `json:"data"`
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
	r.Data = models.User{
		Username: in.Username,
		Password: in.Password,
	}

	return r, nil
}
