package examples

import (
	"go-service/api/response"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	response.Response
}

func Success(ctx *gin.Context) (*SuccessResponse, error) {
	r := &SuccessResponse{}
	return r, nil
}
