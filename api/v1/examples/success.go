package examples

import (
	response2 "go_service/api/response"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	response2.Response
}

func Success(ctx *gin.Context) (*SuccessResponse, error) {
	r := &SuccessResponse{}
	return r, nil
}
