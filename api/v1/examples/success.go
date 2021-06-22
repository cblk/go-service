package examples

import (
	"go_service/api/v1/response"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	response.Response
}

func Success(ctx *gin.Context) (*SuccessResponse, error) {
	r := &SuccessResponse{}
	return r, nil
}
