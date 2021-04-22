package examples

import (
	"go_service/api/v1/response"
	"go_service/models"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	response.Response
	Data models.Task `json:"data" description:"The created task"`
}

func Success(ctx *gin.Context) (*SuccessResponse, error) {
	task := models.NewTask()
	r := &SuccessResponse{}
	r.Data = *task
	return r, nil
}
